import { useState, useEffect } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

interface UseBridgeTxStatusProps {
  originChainId: number
  destinationChainId: number
  transactionHash: string
  bridgeModuleName?: string
  kappa?: string
}

export const useBridgeTxStatus = ({
  originChainId,
  destinationChainId,
  transactionHash,
  bridgeModuleName,
  kappa,
}: UseBridgeTxStatusProps) => {
  const [isComplete, setIsComplete] = useState<boolean>(false)
  const { synapseSDK } = useSynapseContext()

  const getKappa = async (): Promise<string> => {
    if (!bridgeModuleName || !originChainId || !transactionHash) return
    return await synapseSDK.getSynapseTxId(
      originChainId,
      bridgeModuleName,
      transactionHash
    )
  }

  const getBridgeTxStatus = async (
    destinationChainId: number,
    bridgeModuleName: string,
    kappa: string
  ): Promise<boolean> => {
    if (!destinationChainId || !bridgeModuleName || !kappa) return
    return await synapseSDK.getBridgeTxStatus(
      destinationChainId,
      bridgeModuleName,
      kappa
    )
  }

  useEffect(() => {
    console.log('this ran')
  }, [])
}
