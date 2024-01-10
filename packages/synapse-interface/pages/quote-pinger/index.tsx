import { useEffect, useState } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { USDC } from '@/constants/tokens/bridgeable'

const QuotePinger = () => {
  const { synapseSDK } = useSynapseContext()
  const chainIds = [1, 10, 42161]
  const amounts = [500, 5000, 10000]
  const [quotes, setQuotes] = useState({})

  useEffect(() => {
    amounts.forEach((amount) => {
      chainIds.forEach((originId) => {
        chainIds.forEach(async (destId) => {
          if (originId !== destId) {
            try {
              const allQuotes = await synapseSDK.allBridgeQuotes(
                originId,
                destId,
                USDC.addresses[originId],
                USDC.addresses[destId],
                stringToBigInt(amount.toString(), USDC.decimals[originId])
              )

              allQuotes.forEach((quote) => {
                setQuotes((prevQuotes) => ({
                  ...prevQuotes,
                  [amount]: {
                    ...prevQuotes[amount],
                    [`${originId}-${destId}-${quote.bridgeModuleName}`]: quote,
                  },
                }))
              })
            } catch (error) {
              console.log(
                `error fetching quotes for ${originId} to ${destId} with amount ${amount}`,
                error
              )
            }
          }
        })
      })
    })
  }, [])

  const renderQuotes = () => {
    return Object.entries(quotes).map(([amount, quotePairs]) => (
      <div key={amount} className="mb-2">
        <h3>${amount}</h3>
        {Object.entries(quotePairs).map(([key, quote]) => {
          const [originId, destId, bridgeModuleName] = key.split('-')
          const toValueBigInt = BigInt(quote.maxAmountOut?.toString()) ?? 0n

          let textColor
          switch (bridgeModuleName) {
            case 'SynapseRFQ':
              textColor = 'text-purple-500'
              break
            case 'SynapseCCTP':
              textColor = 'text-blue-500'
              break
            default:
              textColor = 'text-gray-500'
          }

          return (
            <div key={key} className="flex items-center space-x-3">
              <div className={`${textColor} w-[120px]`}>{bridgeModuleName}</div>
              <div className="w-[120px]">{`${originId} to ${destId}`}</div>
              <div>
                {formatBigIntToString(toValueBigInt, USDC.decimals[destId], 2)}
              </div>
            </div>
          )
        })}
      </div>
    ))
  }

  return (
    <div className="text-white">
      <div>Bridging USDC to USDC</div>
      {renderQuotes()}
    </div>
  )
}

export default QuotePinger
