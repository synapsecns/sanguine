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
  const {
    originChainId,
    destinationChainId,
    originTokens,
    destinationTokens,
    balances,
    tokens,
  }: BridgeState = useBridgeState()

  let options

  if (label === 'In') {
    options = originTokens
  } else {
    options = filterTokens(destinationTokens, tokens)
  }

  return (
    <TokenPopoverSelect
      selectedChainId={label === 'In' ? originChainId : destinationChainId}
      options={options}
      balances={balances}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
      label={label}
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
