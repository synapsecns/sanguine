import _ from 'lodash'
import { type Chain } from '@/utils/types'

import { ChainPopoverSelect } from './ChainPopoverSelect'
import { ETH, OPTIMISM } from '@/constants/chains/master'

type Props = {
  label: 'To' | 'From'
  isOrigin: boolean
  onChange: (newChain: Chain) => void
  chain: Chain
}

export const ChainSelect = ({ label, isOrigin, chain, onChange }: Props) => {
  const options = [OPTIMISM]

  return (
    <ChainPopoverSelect
      options={options}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={chain}
      label={label}
      isOrigin={isOrigin}
    />
  )
}
