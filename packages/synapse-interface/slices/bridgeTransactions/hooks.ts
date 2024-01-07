import { useAppSelector } from '@/store/hooks'
import { RootState } from '@/store/store'

export const useBridgeTransactionsState =
  (): RootState['bridgeTransactions'] => {
    return useAppSelector((state) => state.bridgeTransactions)
  }
