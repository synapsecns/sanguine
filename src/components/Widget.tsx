import { useMemo, useState, useEffect, useContext, useCallback } from 'react'
import { Address } from 'viem'

import { SynapseSDK } from '@synapsecns/sdk-router'
import { Web3Context } from 'providers/Web3Provider'

import { fetchBridgeQuote } from '@/utils/actions/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { Receipt } from '@/components/Receipt'

import { BridgeableToken, Chain, WidgetProps } from 'types'
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
  setTokens,
} from '@/state/slices/bridge/reducer'
import { useBridgeState } from '@/state/slices/bridge/hooks'

import { generateTheme } from '@/utils/generateTheme'

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
    if (theme === 'dark') return generateTheme({ bgColor: 'dark' })
    if (customTheme) return generateTheme(customTheme)
    return generateTheme()
  })()

  const dispatch = useAppDispatch()

  useEffect(() => {
    dispatch(setTokens(tokens))
  }, [tokens])

  const originTokenDecimals = useMemo(() => {
    if (typeof originToken.decimals === 'number') return originToken.decimals

    return originToken.decimals[originChain.id]
  }, [originToken])

  const destinationTokenDecimals = useMemo(() => {
    if (typeof destinationToken.decimals === 'number')
      return destinationToken.decimals

    return destinationToken.decimals[destinationChain.id]
  }, [destinationToken])

  const {
    state: quoteState,
    callback: fetchQuoteCallback,
    quote,
    error: quoteError,
  } = useBridgeQuoteCallback({
    originChainId: originChain.id,
    originTokenAddress: originToken.addresses[originChain.id],
    destinationChainId: destinationChain.id,
    destinationTokenAddress: destinationToken.addresses[destinationChain.id],
    amount: stringToBigInt(inputAmount, originTokenDecimals),
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
    tokenAddress: originToken.addresses[originChain.id] as Address,
    ownerAddress: connectedAddress as Address,
    chainId: originToken.addresses[originChain.id],
  })

  const approveCallback = useApprove({
    spenderAddress: routerAddress as Address,
    tokenAddress: originToken.addresses[originChain.id] as Address,
    ownerAddress: connectedAddress as Address,
    amount: stringToBigInt(inputAmount, originTokenDecimals),
    chainId: originChain.id,
  })

  const bridgeCallback = useBridgeCallback({
    destinationAddress: connectedAddress as Address,
    originRouterAddress: routerAddress,
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    tokenAddress: originToken.addresses[originChain.id] as Address,
    amount: stringToBigInt(inputAmount, originTokenDecimals),
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

    return formatBigIntToString(max, destinationTokenDecimals, 4)
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
    (newOriginToken: BridgeableToken) => {
      dispatch(setOriginToken(newOriginToken))
    },
    [dispatch]
  )

  const handleDestinationTokenSelection = useCallback(
    (newDestinationToken: BridgeableToken) => {
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
          stringToBigInt(inputAmount, originTokenDecimals),
          originTokenDecimals,
          4
        )}
        receive={maxAmountOut}
      />
      <button
        onClick={fetchQuoteCallback}
        className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-border-hover] active:opacity-40"
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
