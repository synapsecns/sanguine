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
function sortTokensArray(arr: TokenBalance[]): TokenBalance[] {
  return arr.sort((a, b) => {
    const tokenA: Token = a.token
    const tokenB: Token = b.token
    if (tokenA.priorityRank < tokenB.priorityRank) {
      return -1
    } else if (tokenA.priorityRank > tokenB.priorityRank) {
      return 1
    } else {
      return tokenA.symbol < tokenB.symbol ? -1 : 1
    }
  })
}

export function useUserHeldTokens(bridgeTxHash: string): TokenBalance[] {
  const [heldTokens, setHeldTokens] = useState<TokenBalance[]>([])
  const { address } = useAccount()
  const { chain } = useNetwork()

  useEffect(() => {
    if (address === undefined || chain === undefined) return

    async function fetchUserHeldTokens() {
      let multicallInputs: any[] = []
      let multicallData: any

      const currentChainBridgableTokens: Token[] = BRIDGABLE_TOKENS[chain?.id]

      currentChainBridgableTokens.forEach((token) => {
        const tokenAddress =
          token.addresses[chain.id as keyof Token['addresses']]
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
            chainId: chain.id as number,
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
  }, [address, chain, bridgeTxHash])

  return heldTokens
}

export function getSortedBridgableTokens(
  chainId: number,
  bridgeTxHash: string
): TokenBalance[] {
  const userHeldTokens: TokenBalance[] = useUserHeldTokens(bridgeTxHash)

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
    ...sortTokensArray(userHeldTokens),
    ...sortTokensArray(noBalanceTokens),
  ]
}
