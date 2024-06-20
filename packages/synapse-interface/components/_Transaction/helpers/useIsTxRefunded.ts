import { type Address } from 'viem'
import { useEffect, useState } from 'react'
import { readContract } from '@wagmi/core'

import { Chain } from '@/utils/types'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import fastBridgeAbi from '@/constants/abis/fastBridge.json'
import { wagmiConfig } from '@/wagmiConfig'

enum BridgeStatus {
  NULL,
  REQUESTED,
  RELAYER_PROVED,
  RELAYER_CLAIMED,
  REFUNDED,
}

export const useIsTxRefunded = (
  txId: Address | undefined,
  bridgeContract: Address,
  chain: Chain,
  checkForRefund: boolean
) => {
  const [isRefunded, setIsRefunded] = useState<boolean>(false)
  const currentTime = useIntervalTimer(600000)

  const getTxRefundStatus = async () => {
    try {
      const status = await checkRFQTxBridgeStatus(
        txId,
        bridgeContract,
        chain?.id
      )
      console.log('status: ', status)
      if (status === BridgeStatus.REFUNDED) {
        setIsRefunded(true)
      }
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

const checkRFQTxBridgeStatus = async (
  txId: Address,
  bridgeContract: Address,
  chainId: number
) => {
  const result = await readContract(wagmiConfig, {
    abi: fastBridgeAbi,
    address: bridgeContract,
    functionName: 'bridgeStatuses',
    args: [txId],
    chainId,
  })
  return result
}
