'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'

import { useInterchainTransactions } from '@/hooks/useInterchainTransactions'
import { type PageInfo, type InterchainTransaction } from '@/types'
import { shortenHash } from '@/utils/shortenHash'
import { ExplorerLink } from '@/components/ui/ExplorerLink'
import { ChainImage } from '@/components/ui/ChainImage'
import { Loader } from '@/components/ui/Loader'
import { InfoModal } from '@/components/InfoModal'

export const TransactionsTable = () => {
  const [pageInfo, setPageInfo] = useState<PageInfo>({
    endCursor: null,
    startCursor: null,
    hasNextPage: true,
    hasPreviousPage: false,
  })

  const [activeModalTransactionId, setActiveModalTransactionId] =
    useState<string>('')

  const handleOpenModal = (transactionId: string) => {
    setActiveModalTransactionId(transactionId)
  }
  const router = useRouter()

  const response = useInterchainTransactions({
    limit: 50,
    after: pageInfo.endCursor,
    before: pageInfo.startCursor,
  })

  if (response.status !== 'success') {
    return null
  }

  const interchainTransactions = response.data.items

  const handleTransactionClick = (transactionId: string) => {
    router.push(`/tx/${transactionId}`)
  }

  const newPageInfo = response.data.pageInfo

  const handleNextPage = () => {
    setPageInfo({
      startCursor: null,
      endCursor: newPageInfo.endCursor,
      hasNextPage: newPageInfo.hasNextPage,
      hasPreviousPage: true,
    })
  }

  const handlePrevPage = () => {
    setPageInfo({
      startCursor: newPageInfo.startCursor,
      endCursor: null,
      hasNextPage: true,
      hasPreviousPage: newPageInfo.hasPreviousPage,
    })
  }

  return (
    <>
      <div className="flex justify-between">
        <button
          className="hover:cursor-pointer"
          onClick={handlePrevPage}
          disabled={!pageInfo.hasPreviousPage}
        >
          Previous
        </button>
        <button
          className="hover:cursor-pointer"
          onClick={handleNextPage}
          disabled={!pageInfo.hasNextPage}
        >
          Next
        </button>
      </div>
      <div className="bg-gray-800 p-4 rounded">
        <table className="w-full">
          <thead>
            <tr>
              <th className="text-left text-transparent">pl</th>
              <th className="text-left">transactionId</th>
              <th className="text-left">source txn hash</th>
              <th className="text-left">dest txn hash</th>
              <th className="text-left">batch status</th>
              <th className="text-left">txn status</th>
              <th className="text-left">counter</th>
              <th className="text-left text-transparent">pl</th>
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
                  <td className="pl-2 max-w-[90px]">
                    <span
                      className="hover:underline hover:text-blue-500 cursor-pointer"
                      onClick={() => handleTransactionClick(transaction.id)}
                    >
                      {shortenHash(transaction.id)}
                    </span>
                  </td>
                  <td className="max-w-[90px]">
                    {transaction.interchainTransactionSent && (
                      <div className="flex items-center space-x-2">
                        <ChainImage
                          {...transaction.interchainTransactionSent}
                        />
                        <ExplorerLink
                          {...transaction.interchainTransactionSent}
                        />
                      </div>
                    )}
                    <span className="opacity-50 text-sm">
                      {transaction.interchainTransactionSent &&
                        new Date(
                          transaction.interchainTransactionSent?.timestamp *
                            1000
                        ).toLocaleString()}
                    </span>
                  </td>
                  <td className="max-w-[50px]">
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
                          transaction.interchainTransactionReceived?.timestamp *
                            1000
                        ).toLocaleString()}
                    </span>
                  </td>
                  <td className="max-w-[89px] break-words">
                    {transaction.interchainBatch?.status}
                  </td>
                  <td className="max-w-[89px] break-words">
                    {transaction.status}
                  </td>
                  <td>
                    <OptimisticCountdown transaction={transaction} />
                  </td>
                  <td>
                    {transaction.interchainTransactionReceived ? (
                      <span className="text-green-500">✓</span>
                    ) : Date.now() -
                        transaction.interchainTransactionSent?.timestamp *
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
    </>
  )
}

const OptimisticCountdown = ({
  transaction,
}: {
  transaction: InterchainTransaction
}) => {
  const [countdown, setCountdown] = useState<number | string>('')

  useEffect(() => {
    if (!transaction.interchainBatch?.verifiedAt) {
      setCountdown('')
      return
    }

    const updateCountdown = () => {
      const timePassedMs =
        Date.now() - Number(transaction.interchainBatch.verifiedAt) * 1000
      const timePassedSec = timePassedMs / 1000
      const remainingSec = Math.max(0, 30 - timePassedSec)

      if (remainingSec > 0) {
        setCountdown(Math.floor(remainingSec))
      } else {
        if (!transaction.interchainTransactionReceived) {
          setCountdown('...')
        } else {
          setCountdown('P')
        }
      }
    }

    updateCountdown()

    const intervalId = setInterval(updateCountdown, 1000)

    return () => clearInterval(intervalId)
  }, [transaction])

  return <>{countdown}</>
}
