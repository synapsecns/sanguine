import _ from 'lodash'
import Fuse from 'fuse.js'

import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { hasBalance } from '@/utils/hasBalance'
import { sortByPriorityRank } from '@/utils/sortByPriorityRank'
import { CHAINS_BY_ID } from '@/constants/chains'

export const useSwapFromTokenListArray = (searchStr: string) => {
  const { swapFromTokens, swapChainId } = useSwapState()
  const portfolioBalances = usePortfolioBalances()

  const chain = CHAINS_BY_ID[swapChainId]

  let possibleTokens = sortByPriorityRank(swapFromTokens)

  possibleTokens = [
    ...possibleTokens.filter((t) =>
      hasBalance(t, swapChainId, portfolioBalances)
    ),
    ...possibleTokens.filter(
      (t) => !hasBalance(t, swapChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allSwapChainTokens } = getSwapPossibilities({
    fromChainId: swapChainId,
    fromToken: null,
    toChainId: swapChainId,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRank(
    _.difference(allSwapChainTokens, swapFromTokens)
  )

  remainingTokens = [
    ...remainingTokens.filter((t) =>
      hasBalance(t, swapChainId, portfolioBalances)
    ),
    ...remainingTokens.filter(
      (t) => !hasBalance(t, swapChainId, portfolioBalances)
    ),
  ]

  const possibleTokensWithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))
  const remainingTokensWithSource = remainingTokens.map((token) => ({
    ...token,
    source: 'remainingTokens',
  }))

  const masterList = [...possibleTokensWithSource, ...remainingTokensWithSource]

  const fuseOptions = {
    ignoreLocation: true,
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'symbol',
        weight: 2,
      },
      'routeSymbol',
      `addresses.${swapChainId}`,
      'name',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)
    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingTokens = results.filter(
      (item) => item.source === 'remainingTokens'
    )
  }

  return {
    'Swap…': possibleTokens,
    [`More on ${chain?.name}`]: remainingTokens,
  }
}
