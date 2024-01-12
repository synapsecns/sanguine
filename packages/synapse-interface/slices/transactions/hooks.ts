import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useTransactionsState = (): RootState['transactions'] => {
  return useAppSelector((state) => state.transactions)
}
