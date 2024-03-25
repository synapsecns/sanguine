import _ from 'lodash'
import { useState } from 'react'
import Fuse from 'fuse.js'

import { useSwapState } from '@/slices/swap/hooks'
import { getSwapPossibilities } from '@/utils/swapFinder/generateSwapPossibilities'
import { sortByPriorityRank } from '@/utils/sortByPriorityRank'
import { CHAINS_BY_ID } from '@/constants/chains'

export const useSwapToTokenListArray = () => {
  const { swapChainId, swapToTokens } = useSwapState()

  const chain = CHAINS_BY_ID[swapChainId]

  const [searchStr] = useState('')

  let possibleTokens = sortByPriorityRank(swapToTokens)

  const { toTokens: allToChainTokens } = getSwapPossibilities({
    fromChainId: swapChainId,
    fromToken: null,
    toChainId: swapChainId,
    toToken: null,
  })

  let remainingChainTokens = swapChainId
    ? sortByPriorityRank(_.difference(allToChainTokens, swapToTokens))
    : []

  const { toTokens: allTokens } = getSwapPossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: null,
    toToken: null,
  })

  let allOtherToTokens = swapChainId
    ? sortByPriorityRank(_.difference(allTokens, allToChainTokens))
    : sortByPriorityRank(allTokens)

  const possibleTokenswithSource = possibleTokens.map((token) => ({
    ...token,
    source: 'possibleTokens',
  }))

  const remainingChainTokensWithSource = remainingChainTokens.map((token) => ({
    ...token,
    source: 'remainingChainTokens',
  }))

  const allOtherToTokensWithSource = allOtherToTokens.map((token) => ({
    ...token,
    source: 'allOtherToTokens',
  }))

  const masterList = [
    ...possibleTokenswithSource,
    ...remainingChainTokensWithSource,
    ...allOtherToTokensWithSource,
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
      `addresses.${swapChainId}`,
      'name',
    ],
  }
  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleTokens = results.filter((item) => item.source === 'possibleTokens')
    remainingChainTokens = results.filter(
      (item) => item.source === 'remainingChainTokens'
    )
    allOtherToTokens = results.filter(
      (item) => item.source === 'allOtherToTokens'
    )
  }

  return {
    'Receive…': possibleTokens,
    [`More on ${chain?.name}`]: remainingChainTokens,
    'All swappable tokens': allOtherToTokens,
  }
}
