'use client'

import { useInterchainTransaction } from '@/hooks/useInterchainTransaction'

export default function TransactionDetailsPage({
  params,
}: {
  params: { transactionId: string }
}) {
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

  return (
    <div className=" text-white p-5">
      <h1 className="text-xl font-bold mb-4">Transaction Details</h1>
      <div className="mb-4">
        <p>{interchainTransaction.id}</p>
      </div>
      {interchainTransaction.interchainTransactionSent && (
        <div className="mb-4">
          <h2 className="text-lg font-semibold">Sent</h2>
          <p>
            Address: {interchainTransaction.interchainTransactionSent.address}
          </p>
          <p>
            Chain ID: {interchainTransaction.interchainTransactionSent.chainId}
          </p>
          <p>
            Destination Chain ID:{' '}
            {interchainTransaction.interchainTransactionSent.dstChainId}
          </p>
          <p>
            Destination Receiver:{' '}
            {interchainTransaction.interchainTransactionSent.dstReceiver}
          </p>
          <p>
            Source Sender:{' '}
            {interchainTransaction.interchainTransactionSent.srcSender}
          </p>
          <p>
            Timestamp:{' '}
            {new Date(
              interchainTransaction.interchainTransactionSent.timestamp * 1000
            ).toLocaleString()}
          </p>
          <p>
            Transaction Hash:{' '}
            {interchainTransaction.interchainTransactionSent.transactionHash}
          </p>
        </div>
      )}
      {interchainTransaction.interchainTransactionReceived && (
        <div>
          <h2 className="text-lg font-semibold">Received</h2>
          <p>
            Address:{' '}
            {interchainTransaction.interchainTransactionReceived.address}
          </p>
          <p>
            Chain ID:{' '}
            {interchainTransaction.interchainTransactionReceived.chainId}
          </p>
          <p>
            Source Chain ID:{' '}
            {interchainTransaction.interchainTransactionReceived.srcChainId}
          </p>
          <p>
            Destination Receiver:{' '}
            {interchainTransaction.interchainTransactionReceived.dstReceiver}
          </p>
          <p>
            Source Sender:{' '}
            {interchainTransaction.interchainTransactionReceived.srcSender}
          </p>
          <p>
            Timestamp:{' '}
            {new Date(
              interchainTransaction.interchainTransactionReceived.timestamp *
                1000
            ).toLocaleString()}
          </p>
          <p>
            Transaction Hash:{' '}
            {
              interchainTransaction.interchainTransactionReceived
                .transactionHash
            }
          </p>
        </div>
      )}
    </div>
  )
}
