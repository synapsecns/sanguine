import { Chain } from 'types'
import { ChainPopoverSelect } from './ChainPopoverSelect'
import { ARBITRUM, ETHEREUM, POLYGON, OPTIMISM } from '@/constants/chains'

type Props = {
  label: 'To' | 'From'
  onChange: (newChain: Chain) => void
  chain: Chain
}

export function ChainSelect({ label, chain, onChange }: Props) {
  const chains = [ETHEREUM, ARBITRUM, POLYGON, OPTIMISM]

  return (
    <ChainPopoverSelect
      options={chains}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={chain}
      label={label}
    />
  )
}
