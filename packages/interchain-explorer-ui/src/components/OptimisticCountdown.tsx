import { useEffect, useState } from 'react'

import { type InterchainTransaction } from '@/types'

export const OptimisticCountdown = ({
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
          setCountdown('Exec...')
        } else {
          setCountdown('')
        }
      }
    }

    updateCountdown()

    const intervalId = setInterval(updateCountdown, 1000)

    return () => clearInterval(intervalId)
  }, [transaction])

  return <>{countdown}</>
}
