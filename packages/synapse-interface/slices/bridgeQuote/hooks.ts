import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useBridgeQuoteState = (): RootState['bridgeQuote'] => {
  return useAppSelector((state) => state.bridgeQuote)
}
