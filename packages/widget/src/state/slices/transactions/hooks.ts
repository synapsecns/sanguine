import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'

export const useTransactionsState = (): RootState['transactions'] => {
  return useAppSelector((state) => state.transactions)
}
