import { type InterchainTransaction } from '@/types'
import { ExplorerLink } from '@/components/ui/ExplorerLink'

export const InfoModal = ({
  isOpen,
  onClose,
  transaction,
}: {
  isOpen: boolean
  onClose: () => void
  transaction: InterchainTransaction | undefined
}) => {
  if (!isOpen || !transaction) {
    return null
  }

  const { interchainTransactionSent, interchainTransactionReceived } =
    transaction

  return (
    <div className="relative mx-auto p-5 border w-full text-left z-50 shadow-lg rounded-md bg-gray-700">
      <div className="mt-3 ">
        <h3 className="text-lg leading-6 font-bold text-white">
          Transaction Details
        </h3>
        <div className="mb-4">
          <p>{transaction.id}</p>
          <p>{transaction.status ?? 'In-flight'}</p>
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
        <div className="items-center px-4 py-3">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-gray-800 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-gray-500"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  )
}
