import { useMemo, useState, useEffect, useContext, useCallback } from 'react'
import { Address } from 'viem'

import { SynapseSDK } from '@synapsecns/sdk-router'
import { Web3Context } from 'providers/Web3Provider'

import { fetchBridgeQuote } from '@/utils/actions/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { Receipt } from '@/components/Receipt'

import { Chain, TokenMetaData, WidgetProps } from 'types'
import { lightThemeVariables, darkThemeVariables } from '@/constants/index'
import { ChainSelect } from '@/components/ui/ChainSelect'
import { TokenSelect } from '@/components/ui/TokenSelect'
import {
  useBridgeQuoteCallback,
  QuoteCallbackState,
} from '@/hooks/useBridgeQuoteCallback'
import { useAllowance } from '@/hooks/useAllowance'
import { useApprove } from '@/hooks/useApproval'
import { useBridgeCallback } from '@/hooks/useBridgeCallback'

import { useAppDispatch } from '@/state/hooks'
import {
  setDestinationChain,
  setOriginChain,
  setOriginToken,
  setDestinationToken,
} from '@/state/slices/bridge/reducer'
import { useBridgeState } from '@/state/slices/bridge/hooks'

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

  const [inputAmount, setInputAmount] = useState<string>('')

  const { originChain, destinationChain, originToken, destinationToken } =
    useBridgeState()

  const themeVariables = (() => {
    if (theme === 'night') return darkThemeVariables as React.CSSProperties
    if (customTheme) return customTheme as React.CSSProperties
    return lightThemeVariables as React.CSSProperties
  })()

  const dispatch = useAppDispatch()

  const {
    state: quoteState,
    callback: fetchQuoteCallback,
    quote,
    error: quoteError,
  } = useBridgeQuoteCallback({
    originChainId: originToken.chainId,
    originTokenAddress: originToken.tokenAddress,
    destinationChainId: destinationToken.chainId,
    destinationTokenAddress: destinationToken.tokenAddress,
    amount: stringToBigInt(inputAmount, originToken.decimals),
    synapseSDK: synapseSDK,
  })

  console.log('quote:', quote)

  const routerAddress: Address = quote?.routerAddress as Address

  const {
    allowance,
    checkAllowanceCallback,
    error: allowanceError,
  } = useAllowance({
    spenderAddress: routerAddress as Address,
    tokenAddress: originToken.tokenAddress as Address,
    ownerAddress: connectedAddress as Address,
    chainId: originToken.chainId,
  })

  const approveCallback = useApprove({
    spenderAddress: routerAddress as Address,
    tokenAddress: originToken.tokenAddress as Address,
    ownerAddress: connectedAddress as Address,
    amount: stringToBigInt(inputAmount, originToken.decimals),
    chainId: originToken.chainId,
  })

  const bridgeCallback = useBridgeCallback({
    destinationAddress: connectedAddress as Address,
    originRouterAddress: routerAddress,
    originChainId: originToken.chainId,
    destinationChainId: destinationToken.chainId,
    tokenAddress: originToken.tokenAddress as Address,
    amount: stringToBigInt(inputAmount, originToken.decimals),
    originQuery: quote?.originQuery,
    destinationQuery: quote?.destQuery,
    synapseSDK,
    signer,
  })

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

  const handleOriginChainSelection = useCallback(
    (newOriginChain: Chain) => {
      dispatch(setOriginChain(newOriginChain))
    },
    [dispatch]
  )

  const handleDestinationChainSelection = useCallback(
    (newDestinationChain: Chain) => {
      dispatch(setDestinationChain(newDestinationChain))
    },
    [dispatch]
  )

  const handleOriginTokenSelection = useCallback(
    (newOriginToken: TokenMetaData) => {
      dispatch(setOriginToken(newOriginToken))
    },
    [dispatch]
  )

  const handleDestinationTokenSelection = useCallback(
    (newDestinationToken: TokenMetaData) => {
      dispatch(setDestinationToken(newDestinationToken))
    },
    [dispatch]
  )

  return (
    <div
      style={themeVariables}
      className="w-[374px] bg-[--synapse-bg-root] p-2 text-[--synapse-text-primary] rounded-lg font-medium flex flex-col gap-2"
    >
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="From"
          chain={originChain}
          onChange={handleOriginChainSelection}
        />
        <div className="flex pb-2">
          <input
            className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none"
            placeholder="0"
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <TokenSelect
            label="In"
            token={originToken}
            onChange={handleOriginTokenSelection}
          />
        </div>
      </div>
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="To"
          chain={destinationChain}
          onChange={handleDestinationChainSelection}
        />
        <div className="flex items-center justify-between pb-1">
          <input
            className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none cursor-not-allowed"
            disabled={true}
            placeholder=""
            value={
              quoteState === QuoteCallbackState.LOADING ? '...' : maxAmountOut
            }
          />
          <TokenSelect
            label="Out"
            token={destinationToken}
            onChange={handleDestinationTokenSelection}
          />
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
        onClick={fetchQuoteCallback}
        className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
      >
        {quoteState === QuoteCallbackState.LOADING
          ? 'Fetching...'
          : 'Fetch Bridge Quote'}
      </button>

      <button
        className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
        onClick={approveCallback}
      >
        {!quote ? 'Approve (Require Quote)' : 'Approve'}
      </button>

      <button
        className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
        onClick={bridgeCallback}
      >
        {!quote ? 'Bridge (Require Quote)' : 'Bridge'}
      </button>
    </div>
  )
}
