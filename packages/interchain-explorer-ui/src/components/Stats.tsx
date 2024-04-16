import { useStats } from '@/hooks/useStats'

export const Stats = () => {
  const response = useStats()

  if (response.status !== 'success') {
    return null
  }

  const {
    interchainTransactionSent: { count: numSent },
    interchainTransactionReceived: { count: numRecived },
  } = response.data

  return (
    <div className="grid grid-cols-4 gap-4 mb-4">
      <div className="bg-gray-800 p-4 rounded">
        <p className="text-xs uppercase text-gray-400">Total Messages Sent</p>
        <p className="text-xl">{numSent}</p>
      </div>
      <div className="bg-gray-800 p-4 rounded">
        <p className="text-xs uppercase text-gray-400">
          Total Messages Received
        </p>
        <p className="text-xl">{numRecived}</p>
      </div>
      <div className="bg-gray-800 p-4 rounded">
        <p className="text-xs uppercase text-gray-400">Networks</p>
        <p className="text-xl">2</p>
      </div>
    </div>
  )
}
