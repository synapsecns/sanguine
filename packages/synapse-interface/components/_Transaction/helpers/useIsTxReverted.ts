import { Address } from 'viem'
import { useEffect, useState } from 'react'

import { Chain } from '@/utils/types'
import { getTransactionReceipt } from './getTransactionReceipt'
import { useIntervalTimer } from './useIntervalTimer'

export const useIsTxReverted = (
  txHash: Address,
  chain: Chain,
  checkForRevert: boolean
) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)
  const currentTime = useIntervalTimer(60000)

  const getTxRevertStatus = async (txHash: Address, chain: Chain) => {
    /** Remove after testing */
    console.log('fetching for revert: ', txHash)
    /** Remove after testing */

    const receipt = await getTransactionReceipt(txHash, chain)

    if (receipt?.status === 'reverted') {
      setIsReverted(true)
    }
  }

  useEffect(() => {
    if (checkForRevert) {
      getTxRevertStatus(txHash, chain)
    }
  }, [checkForRevert, txHash, chain, currentTime])

  return isReverted
}
