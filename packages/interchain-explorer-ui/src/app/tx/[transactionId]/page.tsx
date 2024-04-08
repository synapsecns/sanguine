'use client'

import { ExplorerLink } from '@/components/ui/ExplorerLink'
import { useInterchainTransaction } from '@/hooks/useInterchainTransaction'

const TransactionIdPage = ({
  params,
}: {
  params: { transactionId: string }
}) => {
  const { data: interchainTransaction, status } = useInterchainTransaction(
    params.transactionId
  )

  if (status === 'pending') {
    return <div>Loading...</div>
  }

  if (status === 'error') {
    return <div>Error fetching transaction details.</div>
  }

  if (!interchainTransaction) {
    return <div>Transaction not found.</div>
  }

  const { interchainTransactionSent, interchainTransactionReceived } =
    interchainTransaction

  return (
    <div className=" text-white p-5">
      <h1 className="text-xl font-bold mb-4">Transaction Details</h1>
      <div className="mb-4">
        <p>{interchainTransaction.id}</p>
      </div>
      {interchainTransactionSent && (
        <div className="mb-4">
          <h2 className="text-lg font-semibold">Sent</h2>
          <p>Chain ID: {interchainTransactionSent.chainId}</p>
          <p>Address: {interchainTransactionSent.address}</p>
          <p>dstChainId: {interchainTransactionSent.dstChainId}</p>
          <p>dstReceiver: {interchainTransactionSent.dstReceiver}</p>
          <p>srcSender: {interchainTransactionSent.srcSender}</p>
          <p>
            Timestamp:{' '}
            {new Date(
              interchainTransactionSent.timestamp * 1000
            ).toLocaleString()}
          </p>
          <p>
            Transaction Hash:{' '}
            <ExplorerLink short={false} {...interchainTransactionSent} />
          </p>
        </div>
      )}
      {interchainTransactionReceived && (
        <div>
          <h2 className="text-lg font-semibold">Received</h2>
          <p>Chain ID: {interchainTransactionReceived.chainId}</p>
          <p>Address: {interchainTransactionReceived.address}</p>
          <p>srcChainId: {interchainTransactionReceived.srcChainId}</p>
          <p>dstReceiver: {interchainTransactionReceived.dstReceiver}</p>
          <p>srcSender: {interchainTransactionReceived.srcSender}</p>
          <p>
            Timestamp:{' '}
            {new Date(
              interchainTransactionReceived.timestamp * 1000
            ).toLocaleString()}
          </p>
          <p>
            Transaction Hash:{' '}
            <ExplorerLink short={false} {...interchainTransactionReceived} />
          </p>
        </div>
      )}
    </div>
  )
}

export default TransactionIdPage
