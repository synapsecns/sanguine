import _ from 'lodash'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { BridgeableToken } from 'types'
import { TokenPopoverSelect } from '@/components/ui/TokenPopoverSelect'
import { useWalletState } from '@/state/slices/wallet/hooks'

type Props = {
  label: 'In' | 'Out'
  isOrigin: boolean
  onChange: (newToken: BridgeableToken) => void
  token: BridgeableToken
}

export function TokenSelect({ label, isOrigin, token, onChange }: Props) {
  const { originTokens, destinationTokens, targetTokens }: BridgeState =
    useBridgeState()

  const { balances } = useWalletState()

  let options
  let remaining

  if (isOrigin) {
    options = originTokens
    remaining = _.difference(targetTokens, originTokens)
  } else {
    options = destinationTokens
    remaining = _.difference(targetTokens, destinationTokens)
  }

  return (
    <TokenPopoverSelect
      options={options}
      remaining={remaining}
      balances={isOrigin ? balances : []}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
    />
  )
}
