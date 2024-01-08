import { useAppSelector } from '@/store/hooks'
import { RootState } from '@/store/store'

export const use_TransactionsState = (): RootState['_transactions'] => {
  return useAppSelector((state) => state._transactions)
}
