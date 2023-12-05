import { useMemo, useState, useEffect, useContext } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Web3Context } from 'providers/Web3Provider'

import { fetchBridgeQuote } from '@/utils/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { DownArrow } from '@/components/DownArrow'
import { Receipt } from '@/components/Receipt'
import { WidgetProps } from 'types'
import { lightThemeVariables, darkThemeVariables } from '../constants'

const originChainId = 1
const destinationChainId = 42161

export const Widget = ({
  chainIds,
  web3Provider,
  networkProviders,
  theme,
  customTheme,
  tokens,
}: WidgetProps) => {
  const synapseSDK = new SynapseSDK(chainIds, networkProviders)
  const web3Context = useContext(Web3Context)
  const { connectedAddress, signer } = web3Context.web3Provider

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

  const handleBridge = async () => {
    try {
      const data = await synapseSDK.bridge(
        connectedAddress,
        quote.routerAddress,
        originToken.chainId,
        destinationToken.chainId,
        originToken?.tokenAddress,
        '1',
        quote.originQuery,
        quote.destQuery
      )
      const tx = await signer.sendTransaction(data)
    } catch (error) {
      console.error('handleBridge error: ', error)
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

  /** Fetch Web3 Provider Data */
  useEffect(() => {
    async function fetchData() {
      try {
        const signer = await web3Provider.getSigner()
        const address = await signer.getAddress()
        const network = await web3Provider.getNetwork()

        web3Context.setWeb3Provider({
          connectedAddress: address,
          networkId: Number(network?.chainId),
          signer,
          provider: web3Provider,
        })
      } catch (e) {
        console.log('Error', e)
      }
    }

    web3Provider && fetchData()
  }, [web3Provider])

  return (
    <div
      style={themeVariables}
      className="w-[374px] bg-[--background] p-2 text-[--primary] rounded-lg font-medium flex flex-col gap-2"
    >
      <div className="border rounded-md bg-[--surface] border-[--separator] p-2 flex flex-col gap-2">
        <div className="flex items-center justify-between">
          <Select label="Ethereum" />
        </div>
        <div className="flex pb-2">
          <input
            className="text-3xl w-full font-semibold bg-[--surface] focus:outline-none"
            placeholder="0"
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <div className="cursor-pointer items-center grid rounded-full bg-[--accent] border border-[--separator] hover:bg-[--separator] hover:border-[--secondary]">
            <span className="col-start-1 row-start-1 pr-3 text-xs h-min justify-self-end">
              <DownArrow />
            </span>
            <select
              className="col-start-1 row-start-1 py-1 pl-3 bg-transparent outline-none appearance-none cursor-pointer pr-7"
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
      <div className="border rounded-md bg-[--surface] border-[--separator] p-2 flex flex-col gap-2">
        <div className="flex items-center justify-between">
          <Select label="Arbitrum" />
        </div>
        <div className="flex items-center justify-between pb-1">
          <input
            className="text-3xl w-full font-semibold bg-[--surface] focus:outline-none cursor-not-allowed"
            disabled={true}
            placeholder="0"
            value={isLoading ? '...' : maxAmountOut}
          />
          <div className="cursor-pointer items-center grid rounded-full bg-[--accent] border border-[--separator] hover:bg-[--separator] hover:border-[--secondary]">
            <span className="col-start-1 row-start-1 pr-3 text-xs h-min justify-self-end">
              <DownArrow />
            </span>
            <select
              className="col-start-1 row-start-1 py-1 pl-3 bg-transparent outline-none appearance-none cursor-pointer pr-7"
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
      <Receipt
        quote={quote ?? null}
        send={formatBigIntToString(
          stringToBigInt(inputAmount, originToken.decimals),
          originToken.decimals,
          4
        )}
        receive={maxAmountOut}
      />
      <button
        className="rounded-md w-full bg-[--surface] font-semibold border border-[--separator] p-2 hover:bg-[--surface] hover:border-[--brand] active:opacity-40"
        onClick={handleFetchQuote}
      >
        {isLoading ? 'Fetching' : 'Fetch Bridge Quote'}
      </button>
    </div>
  )
}

type SelectProps = {
  label?: string
}

function Select({ label }: SelectProps) {
  return (
    <div className="cursor-pointer items-center grid rounded-full bg-[--accent] border border-[--separator] hover:bg-[--separator] hover:border-[--secondary]">
      <span className="col-start-1 row-start-1 pr-3 text-xs h-min justify-self-end">
        <DownArrow />
      </span>
      <select className="col-start-1 row-start-1 py-1 pl-3 bg-transparent outline-none appearance-none cursor-pointer pr-7">
        <option>{label ?? 'Network'}</option>
      </select>
    </div>
  )
}

function Input() {}
