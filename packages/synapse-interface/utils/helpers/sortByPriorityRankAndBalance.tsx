import type { Token } from "@/utils/types"
import { hasBalance } from "./hasBalance"
import { sortByPriorityRank } from "./sortByPriorityRank"

export function sortByPriorityRankAndBalance({
  tokens,
  chainId,
  portfolioBalances,
  source
}: {
  tokens: Token[]
  chainId: number
  portfolioBalances: any
  source?: string
} ) {
  let sortedTokens = sortByPriorityRank(tokens)

  const sortedTokensWithSource = [
    ...sortedTokens.filter((t) =>
      hasBalance(t, chainId, portfolioBalances)
    ),
    ...sortedTokens.filter(
      (t) => !hasBalance(t, chainId, portfolioBalances)
    ),
  ].map(i => {
    return { ...i, source }
  })
  return sortedTokensWithSource
}