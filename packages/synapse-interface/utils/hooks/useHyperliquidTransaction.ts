import { useBridgeState } from '@/slices/bridge/hooks'

export const useHyperliquidTransaction = () => {
  const { toChainId } = useBridgeState()
}
