import { useState, useEffect } from 'react'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { HYPERLIQUID } from '@/constants/chains/master'

interface UseBridgeTxStatusProps {
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName?: string
  kappa?: string
  checkStatus: boolean
  currentTime: number
}

/**
 * Hook that queries SDK for Synapse transaction ID (kappa) and transaction completion status.
 */
export const useBridgeTxStatus = ({
  originChainId,
  destinationChainId,
  originTxHash,
  bridgeModuleName,
  kappa,
  checkStatus = false,
  currentTime,
}: UseBridgeTxStatusProps): [boolean, string, number?] => {
  const { synapseSDK } = useSynapseContext()
  const [isComplete, setIsComplete] = useState<boolean>(false)
  const [fetchedKappa, setFetchedKappa] = useState<string>(kappa ?? null)
  const [deliveryChainId, setDeliveryChainId] = useState<number>()

  useEffect(() => {
    if (!checkStatus) return
    if (isComplete) return
    ;(async () => {
      if (fetchedKappa === null) {
        const _kappa = await getKappa(
          synapseSDK,
          originChainId,
          bridgeModuleName,
          originTxHash
        )
        setFetchedKappa(_kappa)
      }

      if (fetchedKappa) {
        if (
          bridgeModuleName === 'SYN' &&
          destinationChainId === HYPERLIQUID.id
        ) {
          const deliveredToChainId = await getSynDeliveryChainId(
            synapseSDK,
            destinationChainId,
            fetchedKappa
          )
          setDeliveryChainId(deliveredToChainId)
          setIsComplete(deliveredToChainId !== undefined)
          return
        }
        const txStatus = await getBridgeTxStatus(
          synapseSDK,
          destinationChainId,
          bridgeModuleName,
          fetchedKappa
        )

        if (txStatus !== null && txStatus === true && fetchedKappa !== null) {
          setIsComplete(true)
        } else {
          setIsComplete(false)
        }
      }
    })()
  }, [currentTime, checkStatus, fetchedKappa])

  return [isComplete, fetchedKappa, deliveryChainId]
}

const getSynDeliveryChainId = async (
  synapseSDK: any,
  destinationChainId: number,
  txHash: string
): Promise<number | undefined> => {
  try {
    return await synapseSDK?.synModuleSet?.getBridgeDeliveryChainId(
      destinationChainId,
      txHash
    )
  } catch (error) {
    console.error('Error in getBridgeDeliveryChainId:', error)
    return undefined
  }
}

const getKappa = async (
  synapseSDK: any,
  originChainId: number,
  bridgeModuleName: string,
  originTxHash: string
): Promise<string> => {
  if (!synapseSDK) return null
  if (!bridgeModuleName || !originChainId || !originTxHash) return null
  try {
    const kappa = await synapseSDK.getSynapseTxId(
      originChainId,
      bridgeModuleName,
      originTxHash
    )
    return kappa
  } catch (error) {
    console.error('Error in getKappa:', error)
    return null
  }
}

const getBridgeTxStatus = async (
  synapseSDK: any,
  destinationChainId: number,
  bridgeModuleName: string,
  kappa: string
) => {
  if (!synapseSDK) return null
  if (!destinationChainId || !bridgeModuleName || !kappa) return null
  try {
    const status = await synapseSDK.getBridgeTxStatus(
      destinationChainId,
      bridgeModuleName,
      kappa
    )

    return status
  } catch (error) {
    console.error('Error in getBridgeTxStatus:', error)
    return null
  }
}
