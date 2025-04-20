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

type Status = 'SEEN' | 'PENDING' | 'CONFIRMED' | 'PRIORITY' | 'CANCELLED'

interface TransactionStatusData {
  deposit: Deposit
  txs: Transaction[]
}

interface Deposit {
  block: number
  chain: number
  hash: string
  log: number
  sender: string
  shorts: number[]
  status: Status
  time: number
  to: string
  usd: number
  value: string
}

interface Transaction {
  chain: number
  hash: string
  nonce: number
  refund: boolean
  cancelled: boolean
  signer: string
  status: Status
  time: number
  to: string
  usd: number
  value: number
}

enum GasZipStatus {
  CONFIRMED = 'CONFIRMED', // user received funds on dst chain
  CANCELLED = 'CANCELLED', // tx will not be processed, user hasn't received their origin funds back yet
  REFUNDED = 'REFUNDED', // user received funds back on origin chain after tx was cancelled
  OTHER = 'OTHER', // other status
}

const GAS_ZIP_API_URL = 'https://backend.gas.zip/v2/deposit'
const GAS_ZIP_DEPOSIT_ADDRESS = '0x391E7C679d29bD940d63be94AD22A25d25b5A604'

export const useTxRefundStatus = (
  txId: string,
  routerAddress: Address,
  chain: Chain,
  checkForRefund: boolean
) => {
  const [isRefunded, setIsRefunded] = useState<boolean>(false)
  const currentTime = useIntervalTimer(600000)

  const getTxRefundStatus = async () => {
    try {
      if (
        routerAddress.toLowerCase() === GAS_ZIP_DEPOSIT_ADDRESS.toLowerCase()
      ) {
        const status = await checkGasZipTxStatus(txId)
        if (status === GasZipStatus.REFUNDED) {
          setIsRefunded(true)
        }
        return
      }
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
// TODO: this logic could live in the sdk-router
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

const checkGasZipTxStatus = async (txId: string): Promise<GasZipStatus> => {
  try {
    const res = await fetch(`${GAS_ZIP_API_URL}/${txId}`, { method: 'GET' })
    const data = (await res.json()) as TransactionStatusData
    if (!data.txs || !data.txs.length) {
      return GasZipStatus.OTHER
    }
    if (
      data.txs[0].status === GasZipStatus.CANCELLED ||
      data.txs[0].cancelled
    ) {
      // Check if there is a refund tx in the list
      return data.txs.find((tx) => tx.refund)
        ? GasZipStatus.REFUNDED
        : GasZipStatus.CANCELLED
    }
    if (data.txs[0].status === GasZipStatus.CONFIRMED) {
      return GasZipStatus.CONFIRMED
    }
    return GasZipStatus.OTHER
  } catch (error) {
    throw new Error(error)
  }
}
