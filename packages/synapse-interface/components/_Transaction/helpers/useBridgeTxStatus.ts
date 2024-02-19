import { useState, useEffect } from 'react'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'

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
 * Hook will return bridge Tx status via Synapse SDK
 * returns bridge Tx completion status and fetched kappa
 */
export const useBridgeTxStatus = ({
  originChainId,
  destinationChainId,
  originTxHash,
  bridgeModuleName,
  kappa,
  checkStatus = false,
  currentTime,
}: UseBridgeTxStatusProps): [boolean, string] => {
  const { synapseSDK } = useSynapseContext()
  const [isComplete, setIsComplete] = useState<boolean>(false)
  const [fetchedKappa, setFetchedKappa] = useState<string>(kappa ?? null)

  useEffect(() => {
    if (!checkStatus) return
    if (isComplete) return
    ;(async () => {
      /** Remove after testing */
      console.log('fetching bridge tx status:', originTxHash)
      /** Remove after testing */

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

  return [isComplete, fetchedKappa]
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
