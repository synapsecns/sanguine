import _ from 'lodash'
import Fuse from 'fuse.js'

import * as ALL_CHAINS from '@/constants/chains/master'
import {
  CHAINS_BY_ID,
  sortChains,
  PAUSED_FROM_CHAIN_IDS,
} from '@/constants/chains'
import { useBridgeState } from '@/slices/bridge/hooks'

export const useFromChainListArray = (searchStr: string = '') => {
  const { fromChainIds } = useBridgeState()

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(fromChainIds, value.id))
    .values()
    .value()
    .filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      fromChainIds?.map((id) => CHAINS_BY_ID[id])
    )
  ).filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  const possibleChainsWithSource = possibleChains.map((chain) => ({
    ...chain,
    source: 'possibleChains',
  }))

  const remainingChainsWithSource = remainingChains.map((chain) => ({
    ...chain,
    source: 'remainingChains',
  }))

  const masterList = [...possibleChainsWithSource, ...remainingChainsWithSource]

  const fuseOptions = {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'id',
      'nativeCurrency.symbol',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

  return { 'From…': possibleChains, 'All chains': remainingChains }
}
