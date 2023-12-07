import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { TokenMetaData } from 'types'
import { TokenPopoverSelect } from './TokenPopoverSelect'

type Props = {
  label: 'Out' | 'In'
  onChange: (newToken: TokenMetaData) => void
  token: TokenMetaData
}

export function TokenSelect({ label, token, onChange }: Props) {
  const { originChain, destinationChain, tokens }: BridgeState =
    useBridgeState()

  let options

  if (label === 'In') {
    options = tokens.filter((token) => token.chainId === originChain.id)
  } else {
    options = tokens.filter((token) => token.chainId === destinationChain.id)
  }

  return (
    <TokenPopoverSelect
      selectedChain={label === 'In' ? originChain : destinationChain}
      options={options}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
      label={label}
    />
  )
}
