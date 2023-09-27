import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from './reducer'

export default function Updater(): null {
  const { fromChainId, toChainId, fromToken, toTokens }: BridgeState =
    useBridgeState()

  return null
}
