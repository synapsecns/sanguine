'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'

import { useInterchainTransactions } from '@/hooks/useInterchainTransactions'
import { type InterchainTransaction } from '@/types'
import { shortenHash } from '@/utils/shortenHash'
import { ExplorerLink } from '@/components/ui/ExplorerLink'
import { ChainImage } from '@/components/ui/ChainImage'
import { SearchInput } from '@/components/SearchInput'
import { Loader } from '@/components/ui/Loader'
import { InfoModal } from '@/components/InfoModal'

const Home = () => {
  const [activeModalTransactionId, setActiveModalTransactionId] =
    useState<string>('')

  const handleOpenModal = (transactionId: string) => {
    setActiveModalTransactionId(transactionId)
  }
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
                <th className="text-left text-transparent">pl</th>
                <th className="text-left">transactionId</th>
                <th className="text-left">source txn hash</th>
                <th className="text-left">srcSender</th>
                <th className="text-left">dest txn hash</th>
                <th className="text-left">dstReceiver</th>
                <th className="text-left">status</th>
              </tr>
            </thead>
            <tbody>
              {interchainTransactions?.map(
                (transaction: InterchainTransaction, index: number) => (
                  <tr key={index} className="border border-gray-700 m-2">
                    <td className="pl-2">
                      <button
                        onClick={() => handleOpenModal(transaction.id)}
                        className="text-xl hover:text-blue-500"
                      >
                        ⌕
                      </button>
                      {activeModalTransactionId !== '' && (
                        <div className="absolute">
                          <InfoModal
                            isOpen={
                              activeModalTransactionId !== '' &&
                              transaction.id === activeModalTransactionId
                            }
                            onClose={() => setActiveModalTransactionId('')}
                            transaction={interchainTransactions.find(
                              (t) => t.id === activeModalTransactionId
                            )}
                          />
                        </div>
                      )}
                    </td>
                    <td className="pl-2">
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
                      <span className="opacity-50 text-sm">
                        {new Date(
                          transaction.interchainTransactionSent.timestamp * 1000
                        ).toLocaleString()}
                      </span>
                    </td>
                    <td>
                      {shortenHash(
                        transaction.interchainTransactionSent.srcSender
                      )}
                    </td>
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
                      <span className="opacity-50 text-sm">
                        {transaction.interchainTransactionReceived &&
                          new Date(
                            transaction.interchainTransactionReceived
                              ?.timestamp * 1000
                          ).toLocaleString()}
                      </span>
                    </td>
                    <td>
                      {shortenHash(
                        transaction.interchainTransactionReceived?.srcSender
                      )}
                    </td>
                    <td>
                      {transaction.interchainTransactionReceived ? (
                        <span className="text-green-500">✓</span>
                      ) : Date.now() -
                          transaction.interchainTransactionSent.timestamp *
                            1000 >
                        3600000 ? (
                        <span className="text-red-300">x</span>
                      ) : (
                        <Loader />
                      )}
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

export default Home
