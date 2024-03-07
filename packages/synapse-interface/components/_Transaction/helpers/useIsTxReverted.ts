import { Address } from 'viem'
import { useEffect, useState } from 'react'

import { Chain } from '@/utils/types'
import { getTransactionReceipt } from './getTransactionReceipt'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'

export const useIsTxReverted = (
  txHash: Address,
  chain: Chain,
  checkForRevert: boolean
) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)
  const currentTime = useIntervalTimer(60000)

  const getTxRevertStatus = async (txHash: Address, chain: Chain) => {
    try {
      const receipt = await getTransactionReceipt(txHash, chain)
      if (receipt?.status === 'reverted') {
        setIsReverted(true)
      }
    } catch (error) {
      console.error('Failed to fetch transaction receipt:', error)
    }
  }

  useEffect(() => {
    if (checkForRevert) {
      getTxRevertStatus(txHash, chain)
    }
  }, [checkForRevert, txHash, chain, currentTime])

  return isReverted
}
