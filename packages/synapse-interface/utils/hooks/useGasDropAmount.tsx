import { useState, useEffect, useMemo } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

export const useGasDropAmount = (chainId: number) => {
  const [gasDrop, setGasDrop] = useState<BigInt | any>(null)
  const [loading, setLoading] = useState(true)

  const { synapseSDK } = useSynapseContext()

  useEffect(() => {
    if (chainId) {
      ;(async () => {
        try {
          setGasDrop(await synapseSDK.getBridgeGas(chainId))
          setLoading(true)
        } catch (error) {
          //remove after testing
          console.error('Error from useGasDropAmount hook: ', error)
        }
      })()
    }

    return () => setLoading(false)
  }, [chainId])

  return { gasDrop, loading }
}
