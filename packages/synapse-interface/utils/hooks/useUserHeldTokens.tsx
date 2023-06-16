import { useState, useEffect, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { AddressZero } from '@ethersproject/constants'
import { multicall, Address } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import multicallABI from '@/constants/abis/multicall.json'
import erc20ABI from '@/constants/abis/erc20.json'
import { Token } from '../types'
import { Contract, BigNumber } from 'ethers'

interface TokenBalance {
  symbol: string
  balance: BigNumber
}

// Function to sort the tokens by priorityRank and alphabetically
function sortTokens(a, b) {
  if (a.priorityRank < b.priorityRank) {
    return -1
  } else if (a.priorityRank > b.priorityRank) {
    return 1
  } else {
    return a.symbol < b.symbol ? -1 : 1 // In case of a tie in priorityRank, sort alphabetically
  }
}

export function useSortedBridgableTokens(): Token[] {
  const userHeldTokens = useUserHeldTokens()
  const { chain } = useNetwork()

  const availableBridgableTokens: Token[] = BRIDGABLE_TOKENS[chain.id]
  const heldTokenSymbols = userHeldTokens.map(
    (token: TokenBalance) => token.symbol
  )

  const tokensWithBalance = availableBridgableTokens.filter((token) =>
    heldTokenSymbols.includes(token.symbol)
  )
  const tokensNoBalance = availableBridgableTokens.filter(
    (token) => !heldTokenSymbols.includes(token.symbol)
  )

  return [...tokensWithBalance, ...tokensNoBalance]
}

export function useUserHeldTokens(): TokenBalance[] {
  const [heldTokens, setHeldTokens] = useState<TokenBalance[]>([])
  const promise = fetchUserHeldTokens()

  promise.then((response) => setHeldTokens(response))

  return useMemo(() => {
    return heldTokens
  }, [heldTokens])
}

export function fetchUserHeldTokens(): Promise<TokenBalance[]> {
  const { address } = useAccount()
  const { chain } = useNetwork()

  return useMemo(async () => {
    if (address === undefined || chain === undefined) return []

    let heldTokens: TokenBalance[] = []
    const currentChainBridgableTokens: Token[] = BRIDGABLE_TOKENS[chain.id]
    let multicallInputs = []
    let multicallData: any

    currentChainBridgableTokens.map((token) => {
      const tokenAddress = token.addresses[chain.id as keyof Token['addresses']]
      const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11` //deterministic multicall3 ethereum address

      if (tokenAddress === undefined) return
      else if (tokenAddress === AddressZero) {
        multicallInputs.push({
          address: multicallAddress,
          abi: multicallABI,
          functionName: 'getEthBalance',
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
      heldTokens = await multicallData.map(
        (tokenBalance: BigNumber, index: number) => {
          return {
            symbol: currentChainBridgableTokens[index].symbol,
            balance: tokenBalance,
          } as TokenBalance
        }
      )
      return heldTokens.filter((token) => token.balance.gt(0))
    }

    return heldTokens
  }, [address, chain])
}
