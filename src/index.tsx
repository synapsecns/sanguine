import './index.css'
import { useMemo, useState } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { fetchBridgeQuote } from '@/utils/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { DownArrow } from '@/components/DownArrow'
import { Receipt } from '@/components/Receipt'
import { useCustomTheme } from '@/hooks/useCustomTheme'
import { CustomTheme } from 'types'
import { nightTheme } from './constants'

const originChainId = 1
const originTokenAddress = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
const destinationChainId = 42161
const destinationTokenAddress = '0xaf88d065e77c8cc2239327c5edb3a432268e5831'

export const Bridge = ({
  chainIds,
  providers,
  theme,
  customTheme,
}: {
  chainIds: number[]
  providers: any[]
  theme?: 'day' | 'night'
  customTheme?: CustomTheme
}) => {
  const synapseSDK = new SynapseSDK(chainIds, providers)
  const [quote, setQuote] = useState<any>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [inputAmount, setInputAmount] = useState<string>('')

  if (theme === 'night') {
    useCustomTheme(nightTheme)
  } else if (customTheme) {
    useCustomTheme(customTheme)
  } else {
    useCustomTheme({})
  }

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
          amount: stringToBigInt(inputAmount, 6),
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

  const handleInputAmountChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const cleanedInput: string = cleanNumberInput(event.target.value)
    setInputAmount(cleanedInput)
  }

  const maxAmountOut = useMemo(() => {
    if (!quote || !quote.maxAmountOut) {
      return 0
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(max, 6, 4)
  }, [quote])

  return (
    <div className="w-[374px] bg-widget-primary p-2 text-widget-primary">
      <div className="mb-2 border rounded-md bg-widget-surface border-widget-background">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-xl bg-widget-primary border-widget-background">
            <div>Ethereum</div>
            <DownArrow />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            className="text-xl font-semibold bg-widget-surface focus:outline-none"
            placeholder=""
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-xl bg-widget-primary border-widget-background">
            <div>USDC</div>
            <DownArrow />
          </div>
        </div>
      </div>
      <div className="mb-2 border rounded-md bg-widget-surface border-widget-background">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-xl bg-widget-primary border-widget-background">
            <div>Arbitrum</div>
            <DownArrow />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            className="text-xl font-semibold bg-widget-surface"
            disabled={true}
            placeholder=""
            value={isLoading ? '...' : maxAmountOut}
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-xl bg-widget-primary border-widget-background">
            <div>USDC</div>
            <DownArrow />
          </div>
        </div>
      </div>
      {quote ? <Receipt quote={quote} /> : null}
      <button
        className="h-[43px] rounded-md w-full bg-widget-surface  border border-widget-background mt-2"
        onClick={handleFetchQuote}
      >
        {isLoading ? 'Fetching' : 'Fetch Bridge Quote'}
      </button>
    </div>
  )
}
