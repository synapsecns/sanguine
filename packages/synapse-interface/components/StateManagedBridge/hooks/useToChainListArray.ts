import _ from 'lodash'
import Fuse from 'fuse.js'
import { useTranslations } from 'next-intl'

import { ACTIVE_CHAINS_BY_ID, sortChains } from '@/constants/chains'
import { useBridgeState } from '@/slices/bridge/hooks'

export const useToChainListArray = (searchStr: string = '') => {
  const { toChainIds } = useBridgeState()
  const t = useTranslations('Bridge')

  let possibleChains = _(ACTIVE_CHAINS_BY_ID)
    .pickBy((value) => _.includes(toChainIds, value.id))
    .values()
    .value()

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(ACTIVE_CHAINS_BY_ID).map((id) => ACTIVE_CHAINS_BY_ID[id]),
      toChainIds?.map((id) => ACTIVE_CHAINS_BY_ID[id])
    )
  )

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

  return {
    [t('ToWithEllipsis')]: possibleChains,
    [t('All chains')]: remainingChains,
  }

  return { 'To…': possibleChains, 'All chains': remainingChains }
}
