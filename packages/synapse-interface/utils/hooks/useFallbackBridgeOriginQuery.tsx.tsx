import { useLazyGetOriginBridgeTxFallbackQuery } from '@/slices/api/generated'

export const useFallbackBridgeOriginQuery = () => {
  const [fetchFallbackBridgeOriginQuery, fetchedFallbackQueries] =
    useLazyGetOriginBridgeTxFallbackQuery({ pollingInterval: 30000 })

  return null
}
