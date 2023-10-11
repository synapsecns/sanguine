import { useState, useEffect, useMemo } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

export const useGasDropAmount = (
  destinationChainId: number
): {
  gasDrop: bigint | any
  loading: boolean
} => {
  const { synapseSDK } = useSynapseContext()

  const [gasDrop, setGasDrop] = useState<bigint | any>(null)
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    if (destinationChainId) {
      const fetchGasDrop = async () => {
        try {
          setGasDrop(await synapseSDK.getBridgeGas(destinationChainId))
          setLoading(false)
        } catch (error) {
          console.error(error)
          setLoading(false)
        }
      }
      fetchGasDrop()
    }

    return () => setLoading(true)
  }, [destinationChainId])

  return { gasDrop, loading }
}
