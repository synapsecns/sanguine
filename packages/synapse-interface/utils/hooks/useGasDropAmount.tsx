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
          setLoading(false)
        } catch (error) {
          console.error(error)
          setLoading(false)
        }
      })()
    }

    return () => setLoading(true)
  }, [chainId])

  return { gasDrop, loading }
}
