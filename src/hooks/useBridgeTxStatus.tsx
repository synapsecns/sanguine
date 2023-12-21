import { useState, useEffect } from 'react'

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
  synapseSDK,
  originChainId,
  destinationChainId,
  originTxHash,
  bridgeModuleName,
  kappa,
  checkStatus = false,
  currentTime,
}: UseBridgeTxStatusProps) => {
  const [isComplete, setIsComplete] = useState<boolean>(false)
  const [fetchedKappa, setFetchedKappa] = useState<string>(null)

  const getKappa = async (): Promise<string> => {
    if (!synapseSDK) return null
    if (!bridgeModuleName || !originChainId || !originTxHash) return null
    return await synapseSDK.getSynapseTxId(
      originChainId,
      bridgeModuleName,
      originTxHash
    )
  }

  const getBridgeTxStatus = async (
    destinationChainId: number,
    bridgeModuleName: string,
    kappa: string
  ) => {
    if (!synapseSDK) return null
    if (!destinationChainId || !bridgeModuleName || !kappa) return null
    return await synapseSDK.getBridgeTxStatus(
      destinationChainId,
      bridgeModuleName,
      kappa
    )
  }

  useEffect(() => {
    if (!checkStatus) return
    if (isComplete) return
    ;(async () => {
      let _kappa

      if (!kappa) {
        console.log('fetching kappa')
        _kappa = await getKappa()
        setFetchedKappa(_kappa)
      } else {
        _kappa = kappa
      }

      console.log('fetching tx status')
      const txStatus = await getBridgeTxStatus(
        destinationChainId,
        bridgeModuleName,
        _kappa
      )

      if (txStatus !== null) {
        setIsComplete(txStatus)
      }
    })()
  }, [currentTime, checkStatus])

  return [isComplete, fetchedKappa]
}
