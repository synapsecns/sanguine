import { type Address } from 'viem'
import { isNumber, isString } from 'lodash'
import { useEffect, useState } from 'react'
import { readContract } from '@wagmi/core'

import { type Chain } from '@/utils/types'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { wagmiConfig } from '@/wagmiConfig'
import fastBridgeAbi from '@/constants/abis/fastBridge.json'
import fastBridgeRouterAbi from '@/constants/abis/fastBridgeRouter.json'

enum BridgeStatus {
  NULL,
  REQUESTED,
  RELAYER_PROVED,
  RELAYER_CLAIMED,
  REFUNDED,
}

export const useTxRefundStatus = (
  txId: string | undefined,
  routerAddress: Address,
  chain: Chain,
  checkForRefund: boolean
) => {
  const [isRefunded, setIsRefunded] = useState<boolean>(false)
  const currentTime = useIntervalTimer(600000)

  const getTxRefundStatus = async () => {
    try {
      const bridgeContract = await getRFQBridgeContract(
        routerAddress,
        chain?.id
      )

      const status = await checkRFQTxBridgeStatus(
        txId,
        bridgeContract as Address,
        chain?.id
      )

      if (status === BridgeStatus.REFUNDED) {
        setIsRefunded(true)
      }
    } catch (error) {
      console.error('Failed to get transaction refund status:', error)
    }
  }

  useEffect(() => {
    if (checkForRefund) {
      getTxRefundStatus()
    }
  }, [checkForRefund, txId, chain, currentTime])

  return isRefunded
}

const getRFQBridgeContract = async (
  routerAddress: Address,
  chainId: number
): Promise<string | undefined> => {
  try {
    const fastBridgeAddress = await readContract(wagmiConfig, {
      abi: fastBridgeRouterAbi,
      address: routerAddress,
      functionName: 'fastBridge',
      chainId,
    })

    if (!isString(fastBridgeAddress)) {
      throw new Error('Invalid address')
    }

    return fastBridgeAddress
  } catch (error) {
    throw new Error(error)
  }
}

const checkRFQTxBridgeStatus = async (
  txId: string,
  bridgeContract: Address,
  chainId: number
): Promise<number | undefined> => {
  try {
    const status = await readContract(wagmiConfig, {
      abi: fastBridgeAbi,
      address: bridgeContract,
      functionName: 'bridgeStatuses',
      args: [txId],
      chainId,
    })

    if (!isNumber(status)) {
      throw new Error('Invalid status code')
    }

    return status
  } catch (error) {
    throw new Error(error)
  }
}
