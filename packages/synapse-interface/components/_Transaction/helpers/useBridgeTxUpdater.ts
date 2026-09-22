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
import { CHAINS_BY_ID } from '@/constants/chains'

interface UseBridgeTxUpdaterProps {
  connectedAddress: string
  destinationChain: Chain
  kappa: string
  originTxHash: string
  isTxComplete: boolean
  isTxReverted: boolean
  isTxRefunded: boolean
  deliveryChainId?: number
}

/** Updates the stored bridge transaction and destination balances. */
export const useBridgeTxUpdater = ({
  connectedAddress,
  destinationChain,
  kappa,
  originTxHash,
  isTxComplete,
  isTxReverted,
  isTxRefunded,
  deliveryChainId,
}: UseBridgeTxUpdaterProps) => {
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
        const deliveredToChain = deliveryChainId
          ? CHAINS_BY_ID[deliveryChainId]
          : destinationChain
        dispatch(
          completeTransaction({
            originTxHash,
            kappa,
            destinationChain: deliveredToChain,
          })
        )

        /** Update Destination Chain token balances after tx is marked complete  */
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: connectedAddress,
            chainId: deliveredToChain.id,
          })
        )
      }
    }
  }, [
    isTxComplete,
    dispatch,
    transactions,
    originTxHash,
    kappa,
    destinationChain,
    deliveryChainId,
  ])
}
