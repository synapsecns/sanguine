import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  updateTransactionKappa,
  completeTransaction,
  revertTransaction,
} from '@/slices/_transactions/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { Chain } from '@/utils/types'

/**
 * Hook to update Tx store state based on returned SDK method calls
 *
 * @param connectedAddress address that executed tx
 * @param destinationChain dest. chain of executed tx
 * @param kappa fetched kappa from useBridgeTxStatus
 * @param originTxHash executed tx origin hash
 * @param isTxComplete fetched status from useBridgeTxStatus
 * @param isTxReverted fetched tx status on chain
 */
export const useBridgeTxUpdater = (
  connectedAddress: string,
  destinationChain: Chain,
  kappa: string,
  originTxHash: string,
  isTxComplete: boolean,
  isTxReverted: boolean
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

  /** Update tx for reverts in store */
  useEffect(() => {
    if (isTxReverted && !storedTx.isReverted) {
      dispatch(revertTransaction({ originTxHash }))
    }
  }, [isTxReverted])

  /** Update tx for completion in store */
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
