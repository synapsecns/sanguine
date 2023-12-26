import _ from 'lodash'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { BridgeableToken } from 'types'
import { TokenPopoverSelect } from './TokenPopoverSelect'
import { useWalletState } from '@/state/slices/wallet/hooks'

type Props = {
  label: 'In' | 'Out'
  isOrigin: boolean
  onChange: (newToken: BridgeableToken) => void
  token: BridgeableToken
}

export function TokenSelect({ label, isOrigin, token, onChange }: Props) {
  const { originTokens, destinationTokens, tokens }: BridgeState =
    useBridgeState()

  const { balances } = useWalletState()

  const options = isOrigin
    ? originTokens
    : filterTokens(destinationTokens, tokens)
  const remainingOptions = _.difference(tokens, options)

  return (
    <TokenPopoverSelect
      options={options}
      remaining={remainingOptions}
      balances={isOrigin ? balances : []}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
    />
  )
}

const filterTokens = (
  destinationTokens: BridgeableToken[],
  tokens: BridgeableToken[]
) => {
  return destinationTokens.filter((destinationToken) =>
    tokens.some((token) => token.routeSymbol === destinationToken.routeSymbol)
  )
}
