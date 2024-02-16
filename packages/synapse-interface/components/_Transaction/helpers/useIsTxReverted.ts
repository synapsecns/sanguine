import { Address } from 'viem'
import { useEffect, useState } from 'react'

import { Chain } from '@/utils/types'
import { getTransactionReceipt } from './getTransactionReceipt'

export const useIsTxReverted = (
  txHash: Address,
  chain: Chain,
  checkForRevert: boolean
) => {
  const [isReverted, setIsReverted] = useState<boolean>(false)

  const getTxRevertStatus = async (txHash: Address, chain: Chain) => {
    console.log('fetching for revert: ', txHash)
    const receipt = await getTransactionReceipt(txHash, chain)

    if (receipt.status === 'reverted') {
      setIsReverted(true)
    }
  }

  useEffect(() => {
    if (checkForRevert) {
      getTxRevertStatus(txHash, chain)
    }
  }, [checkForRevert, txHash, chain])

  return isReverted
}
