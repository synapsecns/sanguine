import './index.css'
import { useMemo, useState } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { fetchBridgeQuote } from '@/utils/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'

const originChainId = 1
const originTokenAddress = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
const destinationChainId = 42161
const destinationTokenAddress = '0xaf88d065e77c8cc2239327c5edb3a432268e5831'

export const Bridge = ({ chainIds, providers }) => {
  const synapseSDK = new SynapseSDK(chainIds, providers)
  const [quote, setQuote] = useState<any>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [inputAmount, setInputAmount] = useState<number>(0)

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
          amount: stringToBigInt(inputAmount.toString(), 6),
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

  const maxAmountOut = useMemo(() => {
    if (!quote) {
      return null
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(max, 6, 4)
  }, [quote])

  return (
    <div className="w-[374px] bg-[#F5F5F5] p-2">
      <div className="mb-2 bg-white border border-[#DCDCDC] rounded-md">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center space-x-1 rounded-lg bg-[#F5F5F5] pb-1 pl-2 pr-2 pt-1">
            <div>Ethereum</div>
            <ArrowSvg />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            placeholder="Amount"
            className="text-xl"
            value={inputAmount}
            onChange={(e) => setInputAmount(Number(e.target.value))}
          />
          <div className="flex items-center space-x-1 rounded-lg bg-[#F5F5F5] pb-1 pl-2 pr-2 pt-1">
            <div>USDC</div>
            <ArrowSvg />
          </div>
        </div>
      </div>
      <div className="mb-2 bg-white border border-[#DCDCDC] rounded-md">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center space-x-1 rounded-lg bg-[#F5F5F5] pb-1 pl-2 pr-2 pt-1">
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
          <div className="flex items-center space-x-1 rounded-lg bg-[#F5F5F5] pb-1 pl-2 pr-2 pt-1">
            <div>USDC</div>
            <ArrowSvg />
          </div>
        </div>
      </div>
      <button
        className="h-[43px] rounded-md w-full bg-white border border-[#DCDCDC]"
        onClick={handleFetchQuote}
      >
        {isLoading ? 'Fetching' : 'Fetch Bridge Quote'}
      </button>
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
