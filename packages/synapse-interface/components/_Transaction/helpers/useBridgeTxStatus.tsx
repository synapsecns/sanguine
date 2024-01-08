import { useState, useEffect } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

interface UseBridgeTxStatusProps {
  synapseSDK: any
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName?: string
  kappa?: string
  checkStatus: boolean
  currentTime: number // used as trigger to refetch status
}

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
  const [isComplete, setIsComplete] = useState<boolean>(kappa ? true : false)
  const [fetchedKappa, setFetchedKappa] = useState<string>(kappa ?? null)

  const getKappa = async (): Promise<string> => {
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

  useEffect(() => {
    // if (!checkStatus) return
    if (isComplete) return
    ;(async () => {
      if (fetchedKappa === null) {
        let _kappa = await getKappa()
        setFetchedKappa(_kappa)
      }

      if (fetchedKappa) {
        const txStatus = await getBridgeTxStatus(
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
