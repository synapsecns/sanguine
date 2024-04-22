import { type InterchainTransaction } from '@/types'
import { ExplorerLink } from '@/components/ui/ExplorerLink'

export const TransactionInfo = ({
  transaction,
}: {
  transaction: InterchainTransaction
}) => {
  const { interchainTransactionSent, interchainTransactionReceived } =
    transaction

  return (
    <>
      <h3 className="text-lg leading-6 font-bold text-white">
        Transaction Details
      </h3>
      <div className="mb-2">
        <p>{transaction.id}</p>
        <p>{transaction.status ?? 'In-flight'}</p>
      </div>
      <div className="mb-4">
        <h4>Batch info</h4>
        <p>{transaction.interchainBatch.id}</p>
        <p>{transaction.interchainBatch.status}</p>
      </div>
      {interchainTransactionSent && (
        <div className="mb-4">
          <h2 className="text-lg">Sent</h2>
          <p>Chain ID: {interchainTransactionSent.chainId}</p>
          <p>dbNonce: {interchainTransactionSent.dbNonce.toString()}</p>
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
          <p>dbNonce: {interchainTransactionReceived.dbNonce.toString()}</p>
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
    </>
  )
}
