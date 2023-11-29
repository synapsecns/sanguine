import { useEffect } from 'react'
import { Button } from '@/components/Button/Button'
import { SynapseSDK } from '@synapsecns/sdk-router'

export const Bridge = ({
  chainIds,
  providers,
}: {
  chainIds: number[]
  providers: any[]
}) => {
  const synapse = new SynapseSDK(chainIds, providers)

  console.log(synapse)

  useEffect(() => {
    const tryit = async () => {
      return await synapse.bridgeQuote(
        1, // From Chain
        42161, // To Chain
        '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48', // From token Address
        '0xaf88d065e77c8cc2239327c5edb3a432268e5831', // To token Address
        2000000000
      )
    }

    tryit()
      .then((result) => {
        console.log(result)
      })
      .catch((error) => {
        console.error('Error:', error)
      })
  }, [])

  return (
    <div>
      <div>Bridge</div>
      <Button />
    </div>
  )
}
