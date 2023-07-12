import { useState, useEffect, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { AddressZero, Zero } from '@ethersproject/constants'
import { multicall, Address } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import multicallABI from '@/constants/abis/multicall.json'
import erc20ABI from '@/constants/abis/erc20.json'
import { Token } from '../types'
import { Contract, BigNumber } from 'ethers'

interface TokenBalance {
  token: Token
  symbol: string
  balance: BigNumber
}

// Function to sort the tokens by priorityRank and alphabetically
function sortTokensArray(arr: TokenBalance[], chainId: number): TokenBalance[] {
  // Create a copy of the array to prevent modifying the original one
  const sortedArr = [...arr]

  return sortedArr.sort((a, b) => {
    const tokenA: Token = a.token
    const tokenB: Token = b.token

    // Sort by priorityRank first
    if (tokenA.priorityRank !== tokenB.priorityRank) {
      return tokenA.priorityRank - tokenB.priorityRank
    }

    // If priorityRank is the same, sort by balance, taking into account decimals
    const balanceA = a.balance
      .div(BigNumber.from(10).pow(BigNumber.from(tokenA.decimals[chainId])))
      .toNumber()
    const balanceB = b.balance
      .div(BigNumber.from(10).pow(BigNumber.from(tokenB.decimals[chainId])))
      .toNumber()

    if (balanceA !== balanceB) {
      return balanceB - balanceA // For descending order
    }

    // If balance is the same, sort by symbol
    return tokenA.symbol.localeCompare(tokenB.symbol)
  })
}

/**
 *  Returns an array of token balances held by current connected wallet
 *  based on specified chainId
 */
export function useUserHeldTokens(
  address: string,
  chainId: number,
  bridgeTxHash: string
): TokenBalance[] {
  const [heldTokens, setHeldTokens] = useState<TokenBalance[]>([])

  useEffect(() => {
    if (address === undefined || chainId === undefined) return

    async function fetchUserHeldTokens() {
      let multicallInputs: any[] = []
      let multicallData: any

      const currentChainBridgableTokens: Token[] = BRIDGABLE_TOKENS[chainId]

      currentChainBridgableTokens.forEach((token) => {
        const tokenAddress =
          token.addresses[chainId as keyof Token['addresses']]
        const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11` //deterministic multicall3 ethereum address

        if (tokenAddress === undefined) return
        else if (tokenAddress == AddressZero) {
          multicallInputs.push({
            address: multicallAddress,
            abi: multicallABI,
            functionName: 'getEthBalance',
            args: [address],
          } as Partial<Contract>)
        } else {
          const formattedTokenAddress: Address = `0x${tokenAddress.slice(2)}`
          multicallInputs.push({
            address: formattedTokenAddress,
            abi: erc20ABI,
            functionName: 'balanceOf',
            chainId: chainId as number,
            args: [address],
          } as Partial<Contract>)
        }
      })

      if (multicallInputs.length > 0) {
        multicallData = await multicall({ contracts: multicallInputs })
        const newHeldTokens = multicallData.map(
          (tokenBalance: BigNumber, index: number) => {
            return {
              token: currentChainBridgableTokens[index],
              symbol: currentChainBridgableTokens[index].symbol,
              balance: tokenBalance,
            } as TokenBalance
          }
        )
        setHeldTokens(
          newHeldTokens.filter((token: TokenBalance) => token.balance.gt(0))
        )
      }
    }

    fetchUserHeldTokens()
  }, [address, chainId, bridgeTxHash])

  return heldTokens
}

/** Returns sorted array of tokens, ordered by Held / Unheld tokens */
export function getSortedBridgableTokens(
  address: string,
  chainId: number,
  bridgeTxHash?: string
): TokenBalance[] {
  const userHeldTokens: TokenBalance[] = useUserHeldTokens(
    address,
    chainId,
    bridgeTxHash
  )

  if (chainId === undefined) return []

  const availableBridgableTokens: Token[] = BRIDGABLE_TOKENS[chainId]
  const heldTokenSymbols = userHeldTokens.map(
    (token: TokenBalance) => token.symbol
  )

  const noBalanceTokens = availableBridgableTokens
    .filter((token) => !heldTokenSymbols.includes(token.symbol))
    .map((token) => {
      return {
        token: token,
        symbol: token.symbol,
        balance: Zero,
      } as TokenBalance
    })

  return [
    ...sortTokensArray(userHeldTokens, chainId),
    ...sortTokensArray(noBalanceTokens, chainId),
  ]
}
