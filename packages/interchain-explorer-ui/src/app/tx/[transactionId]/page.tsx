'use client'

import { TransactionInfo } from '@/components/TransactionInfo'
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

  return (
    <div className=" text-white p-5">
      <TransactionInfo transaction={interchainTransaction} />
    </div>
  )
}

export default TransactionIdPage
