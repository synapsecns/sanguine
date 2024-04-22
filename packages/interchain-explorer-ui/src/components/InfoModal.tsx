import { useRef } from 'react'

import { type InterchainTransaction } from '@/types'
import { TransactionInfo } from '@/components/TransactionInfo'
import { useCloseOnOutsideClick } from '@/hooks/useCloseOnOutsideClick'

export const InfoModal = ({
  isOpen,
  onClose,
  transaction,
}: {
  isOpen: boolean
  onClose: () => void
  transaction: InterchainTransaction | undefined
}) => {
  const ref = useRef(null)
  useCloseOnOutsideClick(ref, onClose)

  if (!isOpen || !transaction) {
    return null
  }

  return (
    <div
      ref={ref}
      className="relative mx-auto p-5 border w-full text-left z-50 shadow-lg rounded-md bg-gray-700"
    >
      <div className="mt-3 ">
        <TransactionInfo transaction={transaction} />
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
