import _ from 'lodash'
import { Chain } from 'types'

import { ChainPopoverSelect } from '@/components/ui/ChainPopoverSelect'
import { CHAINS_ARRAY, CHAINS_BY_ID } from '@/constants/chains'
import { useBridgeState } from '@/state/slices/bridge/hooks'

type Props = {
  label: 'To' | 'From'
  isOrigin: boolean
  onChange: (newChain: Chain) => void
  chain: Chain
}

export const ChainSelect = ({ label, isOrigin, chain, onChange }: Props) => {
  const { originChainIds, destinationChainIds } = useBridgeState()

  const allChainIds = CHAINS_ARRAY.map((c) => c.id)

  let options
  let remaining

  if (isOrigin) {
    options = originChainIds.map((chainId) => CHAINS_BY_ID[chainId])
    remaining = _.difference(allChainIds, originChainIds).map(
      (chainId) => CHAINS_BY_ID[chainId]
    )
  } else {
    options = destinationChainIds.map((chainId) => CHAINS_BY_ID[chainId])
    remaining = _.difference(allChainIds, destinationChainIds).map(
      (chainId) => CHAINS_BY_ID[chainId]
    )
  }

  return (
    <ChainPopoverSelect
      options={options}
      remaining={remaining}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={chain}
      label={label}
    />
  )
}
