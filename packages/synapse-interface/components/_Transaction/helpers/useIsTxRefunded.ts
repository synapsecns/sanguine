import { Address } from 'viem'
import { useEffect, useState } from 'react'

import { Chain } from '@/utils/types'
import { getTransactionRefundLogs } from './getTransactionReceipt'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useIsTxRefunded = (
  txId: Address | undefined,
  bridgeContract: Address,
  chain: Chain,
  checkForRefund: boolean
) => {
  const [isRefunded, setIsRefunded] = useState<boolean>(false)
  const currentTime = useIntervalTimer(60000)

  const getTxRefundStatus = async () => {
    try {
      await getTransactionRefundLogs(bridgeContract, chain?.id)
    } catch (error) {
      console.error('Failed to fetch transaction receipt:', error)
    }
  }

  useEffect(() => {
    if (checkForRefund) {
      getTxRefundStatus()
    }
  }, [checkForRefund, txId, chain, currentTime])

  return isRefunded
}
