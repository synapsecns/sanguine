import { useEffect, useState } from 'react'

import { getTransactionReceipt } from '@/utils/actions/getTransactionReceipt'
import { useIntervalTimer } from './useIntervalTimer'

enum TransactionStatus {
  REVERT = 0,
  SUCCESS = 1,
}

export const useTxRevertCheck = (
  txHash: string,
  chainId: number,
  provider: any,
  checkForRevert: boolean
) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)
  const currentTime = useIntervalTimer(60000)

  const getTxRevertStatus = async () => {
    try {
      const receipt = await getTransactionReceipt(txHash, chainId, provider)
      console.log('receipt from tx: ', receipt)

      if (receipt?.status === TransactionStatus.REVERT) {
        setIsReverted(true)
      }
    } catch (error) {
      console.error('Failed to fetch transaction receipt:', error)
    }
  }

  useEffect(() => {
    if (checkForRevert) {
      getTxRevertStatus()
    }
  }, [checkForRevert, txHash, chainId, currentTime])

  return isReverted
}
