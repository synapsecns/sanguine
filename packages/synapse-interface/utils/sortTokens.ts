import { multicall, Address, erc20ABI } from '@wagmi/core'
import { zeroAddress } from 'viem'

import multicallABI from '@/constants/abis/multicall.json'
import type { Token } from '@/utils/types'
import { formatBigIntToString } from './bigint/format'
import { TokenAndBalance } from './actions/fetchPortfolioBalances'
import { CHAINS_BY_ID } from '@/constants/chains'

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
): Promise<TokenAndBalance[]> => {
  const tokensWithBalances: any[] = []
  const multicallInputs = []

  if (chainId === undefined || !address) {
    tokens.forEach((token) => {
      tokensWithBalances.push({
        token,
        balance: 0n,
      })
    })
  } else {
    tokens.forEach((token) => {
      // deterministic multicall3 address on all eth chains
      const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11`
      const tokenAddress = token?.addresses[chainId as keyof Token['addresses']]

      if (tokenAddress === zeroAddress || tokenAddress === undefined) {
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
          abi: erc20ABI,
          functionName: 'balanceOf',
          chainId,
          args: [address],
        })
      }
    })
  }

  let multicallData: any[] | any
  if (multicallInputs.length > 0) {
    multicallData = await multicall({
      contracts: multicallInputs,
      chainId,
    })
    return sortArrayByBalance(
      sortByVisibilityRank(
        multicallData.map(
          (
            tokenBalance: { result: bigint; status: string } | undefined,
            index: number
          ) => ({
            queriedChain: CHAINS_BY_ID[chainId],
            token: tokens[index],
            tokenAddress: tokens[index].addresses[chainId],
            balance: tokenBalance.result,
            parsedBalance: formatBigIntToString(
              tokenBalance.result,
              tokens[index]?.decimals[chainId],
              4
            ),
          })
        )
      )
    )
  }

  return tokensWithBalances
}

// Function to sort the tokens by priorityRank and alphabetically
export const sortTokensByPriorityRankAndAlpha = (arr: Token[]): Token[] => {
  // Create a copy of the array to prevent modifying the original one
  const sortedArr = arr && [...arr]

  return sortedArr?.sort((a, b) => {
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
  const hasTokensAndBalances = tokensAndBalances.length > 0
  if (hasTokensAndBalances) {
    const tokensWithBalances = tokensAndBalances
      .filter((t) => !(t.balance === 0n))
      .map((t) => t.token)

    const a = sortTokensByPriorityRankAndAlpha(tokensWithBalances)

    const tokensWithNoBalances = tokensAndBalances
      .filter((t) => t.balance === 0n)
      .map((t) => t.token)

    const b = sortTokensByPriorityRankAndAlpha(tokensWithNoBalances)
    return [...a, ...b]
  } else {
    return []
  }
}
