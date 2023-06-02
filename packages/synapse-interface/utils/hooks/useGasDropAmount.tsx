import { useState, useEffect } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

export const useGasDropAmount = (chainId: number) => {
  const [gasDrop, setGasDrop] = useState<BigInt | any>(null)
  const { synapseSDK } = useSynapseContext()

  useEffect(() => {
    if (chainId) {
      ;(async () => {
        try {
          setGasDrop(await synapseSDK.getBridgeGas(chainId))
        } catch (error) {
          //remove after testing
          console.error('Error from useGasDropAmount hook: ', error)
        }
      })()
    }
  }, [chainId])

  return gasDrop
}
