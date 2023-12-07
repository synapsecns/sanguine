import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { BridgeableToken } from 'types'
import { TokenPopoverSelect } from './TokenPopoverSelect'

type Props = {
  label: 'Out' | 'In'
  onChange: (newToken: BridgeableToken) => void
  token: BridgeableToken
}

export function TokenSelect({ label, token, onChange }: Props) {
  const { originChain, destinationChain, tokens }: BridgeState =
    useBridgeState()

  let options

  if (label === 'In') {
    options = filterObjectsWithAddressKey(tokens, originChain.id)
  } else {
    options = filterObjectsWithAddressKey(tokens, destinationChain.id)
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

export const filterObjectsWithAddressKey = (
  array: BridgeableToken[],
  chainId: number
) => {
  return array.filter((item) => {
    if (item.addresses) {
      return Object.keys(item.addresses).some((key) => Number(key) === chainId)
    }
    return false
  })
}
