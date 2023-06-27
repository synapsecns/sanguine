import { BigNumber } from 'ethers'
import { multicall, Address } from '@wagmi/core'
import { Zero, AddressZero } from '@ethersproject/constants'

import multicallABI from '../constants/abis/multicall.json'
import erc20ABI from '../constants/abis/erc20.json'
import { Token } from '@/utils/types'

interface TokenAndBalance {
  token: Token
  balance: BigNumber
}

export const sortByVisibilityRank = (tokens: Token[]) => {
  if (tokens === undefined) {
    return []
  }

  return Object.values(tokens).sort(
    (a, b) => b.visibilityRank - a.visibilityRank
  )
}

const sortArrayByBalance = (array) => {
  return array.sort((a, b) => {
    const balanceA = BigInt(a.balance || '')
    const balanceB = BigInt(b.balance || '')

    if (balanceA < balanceB) {
      return 1
    } else if (balanceA > balanceB) {
      return -1
    } else {
      return 0
    }
  })
}

export const sortByTokenBalance = async (
  tokens: Token[],
  chainId: number,
  address: any
) => {
  const tokensWithBalances: any[] = []
  const multicallInputs = []
  let multicallData

  if (chainId === undefined || !address) {
    tokens.map((token) => {
      tokensWithBalances.push({
        token,
        balance: Zero,
      })
    })
  } else {
    tokens.map((token) => {
      const tokenAddress = token.addresses[chainId as keyof Token['addresses']]
      const tokenAbi = erc20ABI
      // deterministic multicall3 address on all eth chains
      const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11`

      if (tokenAddress === undefined) {
        return
      }

      if (tokenAddress === AddressZero || tokenAddress === '') {
        multicallInputs.push({
          address: multicallAddress,
          abi: multicallABI,
          functionName: 'getEthBalance',
          chainId,
          args: [address],
        })
      } else {
        multicallInputs.push({
          address: tokenAddress,
          abi: tokenAbi,
          functionName: 'balanceOf',
          chainId,
          args: [address],
        })
      }
    })
  }

  if (multicallInputs.length > 0) {
    multicallData = await multicall({
      contracts: multicallInputs,
      chainId,
    })
    return sortArrayByBalance(
      sortByVisibilityRank(
        multicallData.map((tokenBalance: BigNumber | undefined, index) => ({
          token: tokens[index],
          balance: tokenBalance,
        }))
      )
    )
  }

  return tokensWithBalances
}

// Function to sort the tokens by priorityRank and alphabetically
export const sortTokensByPriorityRankAndAlpha = (arr: Token[]): Token[] => {
  // Create a copy of the array to prevent modifying the original one
  const sortedArr = [...arr]

  return sortedArr.sort((a, b) => {
    // Sort by priorityRank first
    if (a.priorityRank !== b.priorityRank) {
      return a.priorityRank - b.priorityRank
    }

    // If priorityRank is the same, sort by symbol
    return a.symbol.localeCompare(b.symbol)
  })
}

export const separateAndSortTokensWithBalances = (
  tokensAndBalances: TokenAndBalance[]
): Token[] => {
  const hasTokensAndBalances = Object.keys(tokensAndBalances).length > 0

  if (hasTokensAndBalances) {
    const tokensWithBalances = tokensAndBalances
      .filter((t) => !t.balance.eq(Zero))
      .map((t) => t.token)

    const a = sortTokensByPriorityRankAndAlpha(tokensWithBalances)

    const tokensWithNoBalances = tokensAndBalances
      .filter((t) => t.balance.eq(Zero))
      .map((t) => t.token)

    const b = sortTokensByPriorityRankAndAlpha(tokensWithNoBalances)

    return [...a, ...b]
  } else {
    return []
  }
}
