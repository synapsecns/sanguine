import { Address } from 'viem'
import { useEffect, useState } from 'react'

import { Chain } from '@/utils/types'
import { getTransactionReceipt } from './getTransactionReceipt'

export const useIsTxReverted = (
  txHash: Address,
  chain: Chain,
  checkForRevert: boolean,
  checkTime: number
) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)

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
  }, [checkForRevert, txHash, chain, checkTime])

  return isReverted
}
