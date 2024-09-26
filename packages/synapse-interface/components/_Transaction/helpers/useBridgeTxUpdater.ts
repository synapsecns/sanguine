import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  updateTransactionKappa,
  completeTransaction,
  revertTransaction,
  refundTransaction,
  _TransactionDetails,
} from '@/slices/_transactions/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { type Chain } from '@/utils/types'

/**
 * Hook that updates bridge transaction in state.
 *
 * @param {string} connectedAddress - The signer executing bridge transactions.
 * @param {Chain} destinationChain - The destination chain of bridge transaction.
 * @param {string} kappa - The Synapse transaction ID queried from SDK.
 * @param {string} originTxHash - The transaction hash returned when initiating the bridge transaction.
 * @param {boolean} isTxComplete - Whether bridge transaction has completed, queried from SDK.
 * @param {boolean} isTxReverted - Whether bridge transaction was reverted, queried on-chain.
 */
export const useBridgeTxUpdater = (
  connectedAddress: string,
  destinationChain: Chain,
  kappa: string,
  originTxHash: string,
  isTxComplete: boolean,
  isTxReverted: boolean,
  isTxRefunded: boolean
) => {
  const dispatch = useAppDispatch()
  const { transactions } = use_TransactionsState()
  const storedTx: _TransactionDetails = transactions.find(
    (tx) => tx.originTxHash === originTxHash
  )

  /** Update stored tx kappa if not updated with fetched kappa */
  useEffect(() => {
    if (!storedTx.kappa && kappa) {
      dispatch(updateTransactionKappa({ originTxHash, kappa }))
    }
  }, [kappa, storedTx])

  /** Update tx for reverts in store */
  useEffect(() => {
    if (isTxReverted && storedTx.status !== 'reverted') {
      dispatch(revertTransaction({ originTxHash }))
    }
  }, [isTxReverted])

  /** Update tx for refunds in store */
  useEffect(() => {
    if (isTxRefunded && storedTx.status !== 'refunded') {
      dispatch(refundTransaction({ originTxHash }))
    }
  }, [isTxRefunded])

  /** Update tx for completion in store */
  useEffect(() => {
    if (isTxComplete && originTxHash && kappa) {
      /** Check that we have not already marked tx as complete */
      if (storedTx.status !== 'completed') {
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
