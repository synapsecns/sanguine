import { useState } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { fetchBridgeQuote } from './utils/fetchBridgeQuote'

const originChainId = 1
const originTokenAddress = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
const destinationChainId = 42161
const destinationTokenAddress = '0xaf88d065e77c8cc2239327c5edb3a432268e5831'
const amount = 2000000000n

export const Bridge = ({ chainIds, providers }) => {
  const synapseSDK = new SynapseSDK(chainIds, providers)
  const [quote, setQuote] = useState<any>()
  const [isLoading, setIsLoading] = useState<boolean>(false)

  const handleFetchQuote = async () => {
    setIsLoading(true)
    setQuote(null)
    try {
      const result = await fetchBridgeQuote(
        {
          originChainId,
          originTokenAddress,
          destinationChainId,
          destinationTokenAddress,
          amount,
        },
        synapseSDK
      )
      console.log('result', result)
      setQuote(result)
    } catch (error) {
      setQuote(null)
      console.error('Error:', error)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div>
      <div>Bridge</div>
      <div>origin chainId: {originChainId}</div>
      <div>origin originTokenAddress: {originTokenAddress}</div>
      <div>destination chainId: {destinationChainId}</div>
      <div>destinationTokenAddress: {destinationTokenAddress}</div>
      <button onClick={handleFetchQuote}>Fetch Bridge Quote</button>
      {isLoading && <div>Loading...</div>}
      {quote && <div>{quote.routerAddress}</div>}
    </div>
  )
}
