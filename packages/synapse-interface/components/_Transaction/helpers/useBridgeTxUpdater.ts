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
  const storedTx = transactions.find((tx) => tx.originTxHash === originTxHash)

  /** Update stored tx kappa if not updated with fetched kappa */
  useEffect(() => {
    if (!storedTx.kappa && kappa) {
      dispatch(updateTransactionKappa({ originTxHash, kappa }))
    }
  }, [kappa, storedTx])

  /** Update tx status in store */
  useEffect(() => {
    if (isTxComplete && originTxHash && kappa) {
      /** Check that we have not already marked tx as complete */
      if (!storedTx.isComplete) {
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
  }, [isTxComplete, dispatch, transactions, originTxHash, kappa])
}
