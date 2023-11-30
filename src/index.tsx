import './index.css'
import { useMemo, useState } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { fetchBridgeQuote } from '@/utils/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'

const originChainId = 1
const originTokenAddress = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
const destinationChainId = 42161
const destinationTokenAddress = '0xaf88d065e77c8cc2239327c5edb3a432268e5831'

export const Bridge = ({
  chainIds,
  providers,
}: {
  chainIds: number[]
  providers: any[]
}) => {
  const synapseSDK = new SynapseSDK(chainIds, providers)
  const [quote, setQuote] = useState<any>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [inputAmount, setInputAmount] = useState<string>('')

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
    if (!quote) {
      return null
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(max, 6, 4)
  }, [quote])

  return (
    <div className="w-[374px] bg-widget-primary p-2">
      <div className="mb-2 border rounded-md bg-widget-surface border-widget-separator">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 rounded-lg bg-widget-primary">
            <div>Ethereum</div>
            <ArrowSvg />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            placeholder=""
            className="text-xl"
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 rounded-lg bg-widget-primary">
            <div>USDC</div>
            <ArrowSvg />
          </div>
        </div>
      </div>
      <div className="mb-2 border rounded-md bg-widget-surface border-widget-separator">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 rounded-lg bg-widget-primary">
            <div>Arbitrum</div>
            <ArrowSvg />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            placeholder=""
            value={isLoading ? '...' : maxAmountOut}
            className="text-xl"
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 rounded-lg bg-widget-primary">
            <div>USDC</div>
            <ArrowSvg />
          </div>
        </div>
      </div>
      {quote ? <Receipt quote={quote} /> : null}
      <button
        className="h-[43px] rounded-md w-full bg-widget-surface border border-widget-separator mt-2"
        onClick={handleFetchQuote}
      >
        {isLoading ? 'Fetching' : 'Fetch Bridge Quote'}
      </button>
    </div>
  )
}

const Receipt = ({ quote }) => {
  const [isExpanded, setIsExpanded] = useState<boolean>(false)
  const estTime = useMemo(() => {
    return quote.estimatedTime / 60
  }, [quote])

  const handleToggle = () => {
    setIsExpanded(!isExpanded)
  }
  return (
    <div>
      <div className="flex items-center justify-end">
        <div className="text-sm">
          {estTime} min via <span className="text-widget-accent">Synapse</span>
        </div>
        <div onClick={handleToggle}>
          {isExpanded ? <DoubleArrowUp /> : <DoubleArrowDown />}
        </div>
      </div>
      {isExpanded && (
        <div className="p-2 mt-2 text-sm border border-widget-separator">
          <div className="flex items-center justify-between">
            <div>Router</div>
            <div className="text-widget-accent">{quote.bridgeModuleName}</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Origin</div>
            <div>Ethereum</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Destination</div>
            <div>Arbitrum</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Send</div>
            <div>
              {formatBigIntToString(
                quote.originQuery.minAmountOut.toString(),
                6,
                4
              )}
            </div>
          </div>
          <div className="flex items-center justify-between">
            <div>Receive</div>
            <div>
              {formatBigIntToString(
                quote.destQuery.minAmountOut.toString(),
                6,
                4
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

const ArrowSvg = () => {
  return (
    <svg
      width="10"
      height="9"
      viewBox="0 0 10 9"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M0.497159 0.727273H9.58807L5.04261 8.45455L0.497159 0.727273Z"
        fill="#696969"
      />
    </svg>
  )
}

const DoubleArrowDown = () => {
  return (
    <svg
      width="20"
      height="21"
      viewBox="0 0 20 21"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M6 5.5L10 9.5L14 5.5"
        stroke="#A9A9A9"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <path
        d="M6 11.5L10 15.5L14 11.5"
        stroke="#A9A9A9"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  )
}

const DoubleArrowUp = () => {
  return (
    <svg
      width="20"
      height="21"
      viewBox="0 0 20 21"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M6 15.5L10 11.5L14 15.5"
        stroke="#A9A9A9"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <path
        d="M6 9.5L10 5.5L14 9.5"
        stroke="#A9A9A9"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  )
}
