import './index.css'
import { useMemo, useState } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { fetchBridgeQuote } from '@/utils/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { DownArrow } from '@/components/DownArrow'
import { Receipt } from '@/components/Receipt'
import { CustomTheme } from 'types'
import { lightThemeVariables, darkThemeVariables } from './constants'

const originChainId = 1
const destinationChainId = 42161

interface TokenMetaData {
  tokenAddress: string
  symbol: string
  chainId: number
  decimals: number
}

export const Bridge = ({
  chainIds,
  providers,
  theme,
  customTheme,
  tokens,
}: {
  chainIds: number[]
  providers: any[]
  theme?: 'day' | 'night'
  customTheme?: CustomTheme
  tokens: TokenMetaData[]
}) => {
  const synapseSDK = new SynapseSDK(chainIds, providers)
  const [quote, setQuote] = useState<any>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [inputAmount, setInputAmount] = useState<string>('')

  const [originToken, setOriginToken] = useState(
    tokens.find((token) => token.chainId === originChainId)
  )
  const [destinationToken, setDestinationToken] = useState(
    tokens.find((token) => token.chainId === destinationChainId)
  )

  const themeVariables = (() => {
    if (theme === 'night') return darkThemeVariables as React.CSSProperties
    if (customTheme) return customTheme as React.CSSProperties
    return lightThemeVariables as React.CSSProperties
  })()

  const handleFetchQuote = async () => {
    setIsLoading(true)
    setQuote(null)
    try {
      const result = await fetchBridgeQuote(
        {
          originChainId: originToken.chainId,
          originTokenAddress: originToken.tokenAddress,
          destinationChainId: destinationToken.chainId,
          destinationTokenAddress: destinationToken.tokenAddress,
          amount: stringToBigInt(inputAmount, originToken.decimals),
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

    return formatBigIntToString(max, destinationToken.decimals, 4)
  }, [quote])

  return (
    <div
      style={themeVariables}
      className="w-[374px] bg-[--background] p-2 text-[--primary] rounded-lg"
    >
      <div className="mb-2 border rounded-md bg-[--surface] border-[--separator]">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-full bg-[--accent] border-[--separator]">
            <div>Ethereum</div>
            <DownArrow />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            className="text-xl font-semibold bg-[--surface] focus:outline-none"
            placeholder=""
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-full bg-[--accent] border-[--separator]">
            <select
              className="bg-transparent hover:cursor-pointer focus:outline-none"
              value={originToken.tokenAddress}
              onChange={(e) => {
                setOriginToken(
                  tokens.find((token) => token.tokenAddress === e.target.value)
                )
                setQuote(null)
              }}
            >
              {tokens
                .filter((token) => token.chainId === originChainId)
                .map((token, index) => (
                  <option key={index} value={token.tokenAddress}>
                    <div className="flex items-center">
                      <div>{token.symbol}</div>
                      <DownArrow />
                    </div>
                  </option>
                ))}
            </select>
          </div>
        </div>
      </div>
      <div className="mb-2 border rounded-md bg-[--surface] border-[--separator]">
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-full bg-[--accent] border-[--separator]">
            <div>Arbitrum</div>
            <DownArrow />
          </div>
          <div className="text-sm"></div>
        </div>
        <div className="flex items-center justify-between p-2">
          <input
            className="text-xl font-semibold bg-[--surface]"
            disabled={true}
            placeholder=""
            value={isLoading ? '...' : maxAmountOut}
          />
          <div className="flex items-center pt-1 pb-1 pl-2 pr-2 space-x-1 border rounded-full bg-[--accent] border-[--separator]">
            <select
              className="bg-transparent hover:cursor-pointer focus:outline-none"
              value={destinationToken.tokenAddress}
              onChange={(e) => {
                setDestinationToken(
                  tokens.find((token) => token.tokenAddress === e.target.value)
                )
                setQuote(null)
              }}
            >
              {tokens
                .filter((token) => token.chainId === destinationChainId)
                .map((token, index) => (
                  <option key={index} value={token.tokenAddress}>
                    <div className="flex items-center">
                      <div>{token.symbol}</div>
                      <DownArrow />
                    </div>
                  </option>
                ))}
            </select>
          </div>
        </div>
      </div>
      {quote ? (
        <Receipt
          quote={quote}
          send={formatBigIntToString(
            stringToBigInt(inputAmount, originToken.decimals),
            originToken.decimals,
            4
          )}
          receive={maxAmountOut}
        />
      ) : null}
      <button
        className="h-[43px] rounded-md w-full bg-[--surface] font-semibold border border-[--separator] mt-2"
        onClick={handleFetchQuote}
      >
        {isLoading ? 'Fetching' : 'Fetch Bridge Quote'}
      </button>
    </div>
  )
}
