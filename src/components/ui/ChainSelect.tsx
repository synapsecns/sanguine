import { Chain } from 'types'
import { ChainPopoverSelect } from './ChainPopoverSelect'

type Props = {
  label: 'To' | 'From'
  onChange: (newChain: Chain) => void
  chain: Chain
}

const ETH = {
  id: 1,
  name: 'Ethereum',
}

const ARBITRUM = {
  id: 42161,
  name: 'Arbitrum',
}

export function ChainSelect({ label, chain, onChange }: Props) {
  const chains = [ETH, ARBITRUM]

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
