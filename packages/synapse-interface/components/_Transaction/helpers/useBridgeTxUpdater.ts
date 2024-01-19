import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  updateTransactionKappa,
  completeTransaction,
} from '@/slices/_transactions/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { Chain } from '@/utils/types'

/**
 * Hook that updates bridge Tx in store for kappa + isComplete
 * Tx matched in store using originTxHash
 */
export const useBridgeTxUpdater = (
  connectedAddress: string,
  destinationChain: Chain,
  kappa: string,
  originTxHash: string,
  isTxComplete: boolean
) => {
  const dispatch = useAppDispatch()
  const { transactions } = use_TransactionsState()

  /** Update tx kappa in store when available */
  if (kappa && originTxHash) {
    dispatch(updateTransactionKappa({ originTxHash, kappa: kappa as string }))
  }

  /** Update tx for completion */
  /** Check that we have not already marked tx as complete */
  useEffect(() => {
    if (isTxComplete && originTxHash && kappa) {
      const txn = transactions.find((tx) => tx.originTxHash === originTxHash)
      if (!txn.isComplete) {
        dispatch(completeTransaction({ originTxHash, kappa }))
        /** Update Destination Chain token balances after tx is marked complete  */
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: connectedAddress,
            chainId: destinationChain.id,
          })
        )
      }
    }
  }, [isTxComplete, dispatch, transactions])
}
