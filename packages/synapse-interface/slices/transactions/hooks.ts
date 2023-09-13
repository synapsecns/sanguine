import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import { BridgeTransaction } from '../api/generated'
import { TransactionsState } from './reducer'

export const useTransactionsState = (): RootState['transactions'] => {
  return useAppSelector((state) => state.transactions)
}

export const useUserHistoricalTransactions = ({
  recentTimeframe = 600, //default 10 mins to be considered Recent
}: {
  recentTimeframe?: number
}) => {
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
  }: TransactionsState = useTransactionsState()

  const getCurrentUnixTimestamp = (): number => Math.floor(Date.now() / 1000)

  const filterRecentlyCompleted = (
    transactions: BridgeTransaction[],
    thresholdSeconds: number
  ): BridgeTransaction[] => {
    const currentTimestamp: number = getCurrentUnixTimestamp()
    return transactions.filter((transaction) => {
      const timestamp: number = Number(transaction.fromInfo?.formattedTime)
      return currentTimestamp - timestamp <= thresholdSeconds
    })
  }

  const recentlyCompletedTransactions: BridgeTransaction[] =
    filterRecentlyCompleted(userHistoricalTransactions, recentTimeframe)

  const historicalCompletedTransactions: BridgeTransaction[] =
    userHistoricalTransactions.filter(
      (transaction) => !recentlyCompletedTransactions.includes(transaction)
    )

  return {
    recentlyCompletedTransactions,
    historicalCompletedTransactions,
    allCompletedTransactions: userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
  }
}
