import { useStats } from '@/hooks/useStats'
import { InterchainTransaction } from '@/types'

export const Stats = () => {
  const response = useStats()

  if (response.status !== 'success') {
    return null
  }

  const numSent = getMaxCount(response.data, 'interchainTransactionSent')
  const numReceived = getMaxCount(
    response.data,
    'interchainTransactionReceived'
  )

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
        <p className="text-xl">{numReceived}</p>
      </div>
      <div className="bg-gray-800 p-4 rounded">
        <p className="text-xs uppercase text-gray-400">Networks</p>
        <p className="text-xl">2</p>
      </div>
    </div>
  )
}

const getMaxCount = (
  data: InterchainTransaction[],
  key: 'interchainTransactionSent' | 'interchainTransactionReceived'
): number => {
  return (
    data
      .filter((t) => t[key] !== null)
      .map((t) => t[key])
      .sort((a, b) => Number(b?.count) - Number(a?.count))[0]?.count || 0
  )
}
