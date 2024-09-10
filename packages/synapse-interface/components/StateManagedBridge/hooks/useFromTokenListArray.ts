import _ from 'lodash'
import Fuse from 'fuse.js'
import { useTranslations } from 'next-intl'

import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'
import { hasBalance } from '@/utils/hasBalance'
import { sortByPriorityRank } from '@/utils/sortByPriorityRank'

export const useFromTokenListArray = (searchStr: string = '') => {
  const { fromTokens, fromChainId } = useBridgeState()
  const portfolioBalances = usePortfolioBalances()

  const t = useTranslations('Bridge')

  let possibleTokens = sortByPriorityRank(fromTokens)

  possibleTokens = [
    ...possibleTokens.filter((t) =>
      hasBalance(t, fromChainId, portfolioBalances)
    ),
    ...possibleTokens.filter(
      (t) => !hasBalance(t, fromChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allFromChainTokens } = getRoutePossibilities({
    fromChainId,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let remainingTokens = sortByPriorityRank(
    _.difference(allFromChainTokens, fromTokens)
  )

  remainingTokens = [
    ...remainingTokens.filter((t) =>
      hasBalance(t, fromChainId, portfolioBalances)
    ),
    ...remainingTokens.filter(
      (t) => !hasBalance(t, fromChainId, portfolioBalances)
    ),
  ]

  const { fromTokens: allTokens } = getRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherFromTokens = sortByPriorityRank(
    _.difference(allTokens, allFromChainTokens)
  )

  const possibleTokensWithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))
  const remainingTokensWithSource = remainingTokens.map((token) => ({
    ...token,
    source: 'remainingTokens',
  }))
  const allOtherFromTokensWithSource = allOtherFromTokens.map((token) => ({
    ...token,
    source: 'allOtherFromTokens',
  }))

  const masterList = [
    ...possibleTokensWithSource,
    ...remainingTokensWithSource,
    ...allOtherFromTokensWithSource,
  ]

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
      `addresses.${fromChainId}`,
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
    allOtherFromTokens = results.filter(
      (item) => item.source === 'allOtherFromTokens'
    )
  }

  return {
    [t('SendWithEllipsis')]: possibleTokens,
    [t('All sendable tokens')]: remainingTokens,
    [t('All other tokens')]: allOtherFromTokens,
  }
}
