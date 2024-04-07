'use client'

import { useInterchainTransactions } from '@/hooks/useInterchainTransactions'
import { type InterchainTransaction } from '@/types'
import { shortenHash } from '@/utils/shortenHash'
import { ExplorerLink } from '@/components/ui/ExplorerLink'
import { ChainImage } from '@/components/ui/ChainImage'
import { useRouter } from 'next/navigation'
import SearchInput from '@/components/SearchInput'

export default function Home() {
  const router = useRouter()

  const response = useInterchainTransactions()

  if (response.status !== 'success') {
    return null
  }

  const interchainTransactions = response.data

  const numSent = interchainTransactions.filter(
    (i) => i.interchainTransactionSent
  ).length

  const numRecived = interchainTransactions.filter(
    (i) => i.interchainTransactionReceived
  ).length

  const handleTransactionClick = (transactionId: string) => {
    router.push(`/tx/${transactionId}`)
  }

  return (
    <div className="bg-gray-900 text-white min-h-screen">
      <div className="container mx-auto p-4">
        <div className="flex justify-between items-center mb-4">
          <SearchInput />
        </div>

        <div className="grid grid-cols-4 gap-4 mb-4">
          <div className="bg-gray-800 p-4 rounded">
            <p className="text-xs uppercase text-gray-400">
              Total Messages Sent
            </p>
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

        <div className="bg-gray-800 p-4 rounded">
          <table className="w-full">
            <thead>
              <tr>
                <th className="text-left">transactionId</th>
                <th className="text-left">source txn hash</th>
                <th className="text-left">timestamp</th>
                <th className="text-left">dest txn hash</th>
                <th className="text-left">timestamp</th>
              </tr>
            </thead>
            <tbody>
              {interchainTransactions?.map(
                (transaction: InterchainTransaction, index: number) => (
                  <tr key={index}>
                    <td>
                      <span
                        className="hover:underline hover:text-blue-500 cursor-pointer"
                        onClick={() => handleTransactionClick(transaction.id)}
                      >
                        {shortenHash(transaction.id)}
                      </span>
                    </td>
                    <td className="">
                      <div className="flex items-center space-x-2">
                        <ChainImage
                          {...transaction.interchainTransactionSent}
                        />
                        <ExplorerLink
                          {...transaction.interchainTransactionSent}
                        />
                      </div>
                    </td>
                    <td>{transaction.interchainTransactionSent.timestamp}</td>
                    <td className="">
                      {transaction.interchainTransactionReceived && (
                        <div className="flex items-center space-x-2">
                          <ChainImage
                            {...transaction.interchainTransactionReceived}
                          />
                          <ExplorerLink
                            {...transaction.interchainTransactionReceived}
                          />
                        </div>
                      )}
                    </td>
                    <td>
                      {transaction.interchainTransactionReceived?.timestamp}
                    </td>
                  </tr>
                )
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}
