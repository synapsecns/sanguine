import { Address } from 'wagmi'
import { useSynapseContext } from '../providers/SynapseProvider'
import { useEffect, useState } from 'react'
import { BigintIsh } from '@/../sdk-router/dist/constants'

interface useBridgeQuoteProps {
  fromChainId: number
  toChainId: number
  fromTokenAddress: string | Address
  toTokenAddress: string | Address
  inputValue: BigintIsh
}

/*
  useBridgeQuote hook
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  - Wraps SynapseSDK calls in Promise wrapper to be cancellable when new request is made
  */
export const useBridgeQuote = ({
  fromChainId,
  toChainId,
  fromTokenAddress,
  toTokenAddress,
  inputValue,
}: useBridgeQuoteProps) => {
  const SynapseSDK = useSynapseContext()

  const [quote, setQuote] = useState(null)
  const [error, setError] = useState(null)
  const [cancelToken, setCancelToken] = useState(null)

  useEffect(() => {
    let isMounted = true

    const fetchData = async () => {
      try {
        const newCancelToken = {}
        setCancelToken(newCancelToken)

        const response = await SynapseSDK.bridgeQuote(
          fromChainId,
          toChainId,
          fromTokenAddress,
          toTokenAddress,
          inputValue
        )

        if (isMounted) {
          setQuote(response)
          setError(null)
        }
      } catch (error) {
        if (isMounted) {
          setError(error)
          setQuote(null)
        }
      }
    }

    fetchData()

    return () => {
      isMounted = false
      if (cancelToken) {
        cancelToken.cancel() // Cancel the ongoing request when cleaning up
      }
    }
  }, [fromChainId, toChainId, fromTokenAddress, toTokenAddress, inputValue])

  return { quote, error }
}
