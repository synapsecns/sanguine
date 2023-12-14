import { useMemo, useState, useEffect, useContext, useCallback } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Web3Context } from 'providers/Web3Provider'

import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { Receipt } from '@/components/Receipt'

import { BridgeableToken, Chain, WidgetProps } from 'types'
import { ChainSelect } from '@/components/ui/ChainSelect'
import { TokenSelect } from '@/components/ui/TokenSelect'
import { useBridgeQuote, QuoteCallbackState } from '@/hooks/useBridgeQuote'
import { useAllowance } from '@/hooks/useAllowance'
import {
  ApproveCallbackState,
  UseApproveCallbackProps,
  useApproveCallback,
} from '@/hooks/useApproveCallback'
import {
  UseBridgeCallbackArgs,
  useBridgeCallback,
  BridgeCallbackState,
} from '@/hooks/useBridgeCallback'

import { useAppDispatch } from '@/state/hooks'
import {
  setDestinationChainId,
  setOriginChainId,
  setOriginToken,
  setDestinationToken,
  setTokens,
  setDebouncedInputAmount,
  setInputAmount,
} from '@/state/slices/bridge/reducer'
import {
  fetchAndStoreTokenBalances,
  useBridgeState,
} from '@/state/slices/bridge/hooks'
import { BridgeButton } from './BridgeButton'
import { isOnlyZeroes } from '@/utils/isOnlyZeroes'
import { switchNetwork } from '@/utils/actions/switchNetwork'

import { generateTheme } from '@/utils/generateTheme'
import { fetchTokenBalances } from '@/utils/actions/fetchTokenBalances'
import { AvailableBalance } from './AvailableBalance'

const chains = {
  1: {
    id: 1,
    name: 'Ethereum',
  },
  137: {
    id: 137,
    name: 'Polygon',
  },
  42161: {
    id: 42161,
    name: 'Arbitrum',
  },
  10: {
    id: 10,
    name: 'Optimism',
  },
}

export const Widget = ({
  chainIds,
  web3Provider,
  networkProviders,
  theme,
  customTheme,
  tokens,
  toChainId,
}: WidgetProps) => {
  const dispatch = useAppDispatch()
  const synapseSDK = new SynapseSDK(chainIds, networkProviders)
  const web3Context = useContext(Web3Context)
  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  console.log('web3Provider:', web3Provider)

  const {
    inputAmount,
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
    tokens: allTokens,
    balances,
  } = useBridgeState()

  /** Select Consumer networkProvider based on Origin ChainId */
  const originChainProvider = useMemo(() => {
    if (!Array.isArray(networkProviders)) return null
    if (!originChainId) return null
    const _provider = networkProviders.find(
      (provider) => Number(provider?._network?.chainId) === originChainId
    )
    return _provider
  }, [originChainId, networkProviders])

  const themeVariables = (() => {
    if (theme === 'dark') return generateTheme({ bgColor: 'dark' })
    if (customTheme) return generateTheme(customTheme)
    return generateTheme()
  })()

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

  useEffect(() => {
    dispatch(setTokens(tokens))
    dispatch(setDestinationChainId(toChainId))
    dispatch(setDestinationToken(tokens[0]))
  }, [tokens, toChainId])

  /** Debounce user input to fetch bridge quote (in ms) */
  useEffect(() => {
    const DEBOUNCE_DELAY = 300
    const debounceTimer = setTimeout(() => {
      dispatch(setDebouncedInputAmount(inputAmount))
    }, DEBOUNCE_DELAY)

    return () => {
      clearTimeout(debounceTimer)
    }
  }, [dispatch, inputAmount])

  /** Fetch token balances when signer/address connected */
  useEffect(() => {
    if (!signer && !originChainProvider) return
    if (originChainId && tokens && connectedAddress) {
      dispatch(
        fetchAndStoreTokenBalances({
          address: connectedAddress,
          chainId: originChainId,
          tokens: allTokens,
          signerOrProvider: originChainProvider ?? signer,
        })
      )
    }
  }, [originChainId, allTokens, connectedAddress, signer, originChainProvider])

  const originTokenDecimals = useMemo(() => {
    if (typeof originToken?.decimals === 'number') return originToken?.decimals

    return originToken?.decimals[originChainId]
  }, [originToken])

  const destinationTokenDecimals = useMemo(() => {
    if (typeof destinationToken?.decimals === 'number')
      return destinationToken?.decimals

    return destinationToken?.decimals[destinationChainId]
  }, [destinationToken])

  const {
    state: quoteState,
    callback: fetchQuoteCallback,
    reset: resetQuote,
    quote,
    error: quoteError,
  } = useBridgeQuote({
    originChainId: originChainId,
    originTokenAddress: originToken?.addresses[originChainId],
    destinationChainId: destinationChainId,
    destinationTokenAddress: destinationToken?.addresses[destinationChainId],
    amount: stringToBigInt(debouncedInputAmount, originTokenDecimals),
    synapseSDK: synapseSDK,
  })

  const routerAddress: string = quote?.routerAddress

  const { allowance, checkAllowanceCallback } = useAllowance({
    spenderAddress: routerAddress,
    tokenAddress: originToken?.addresses[originChainId],
    ownerAddress: connectedAddress,
    chainId: originToken?.addresses[originChainId],
    provider: originChainProvider ?? provider,
  })

  const useApproveCallbackArgs: UseApproveCallbackProps = {
    spenderAddress: routerAddress,
    tokenAddress: originToken?.addresses[originChainId],
    ownerAddress: connectedAddress,
    amount: stringToBigInt(inputAmount, originTokenDecimals),
    chainId: originChainId,
    onSuccess: checkAllowanceCallback,
    signer: signer,
  }
  const {
    state: approveState,
    callback: approveCallback,
    error: approveError,
  } = useApproveCallback(useApproveCallbackArgs)

  const useBridgeCallbackArgs: UseBridgeCallbackArgs = {
    destinationAddress: connectedAddress,
    originRouterAddress: routerAddress,
    originChainId: originChainId,
    destinationChainId: destinationChainId,
    tokenAddress: originToken?.addresses[originChainId],
    amount: stringToBigInt(debouncedInputAmount, originTokenDecimals),
    originQuery: quote?.originQuery,
    destinationQuery: quote?.destQuery,
    synapseSDK,
    signer,
  }
  const {
    state: bridgeState,
    callback: bridgeCallback,
    error: bridgeError,
  } = useBridgeCallback(useBridgeCallbackArgs)

  const isConnectedNetworkSelectedOriginNetwork: boolean = useMemo(() => {
    return networkId === originChainId
  }, [originChainId, networkId])

  const formattedInputAmount: bigint = useMemo(() => {
    return stringToBigInt(debouncedInputAmount ?? '0', originTokenDecimals)
  }, [debouncedInputAmount, originToken])

  const isInputValid: boolean = useMemo(() => {
    if (debouncedInputAmount === '') return false
    if (isOnlyZeroes(debouncedInputAmount)) return false
    return true
  }, [debouncedInputAmount])

  const isApproved: boolean = useMemo(() => {
    if (allowance === null) return true
    if (!formattedInputAmount) return true
    return formattedInputAmount <= allowance
  }, [formattedInputAmount, allowance])

  /** Handle refreshing quotes */
  useEffect(() => {
    if (
      isInputValid &&
      originToken &&
      destinationToken &&
      originChainId &&
      destinationChainId &&
      fetchQuoteCallback
    ) {
      fetchQuoteCallback()
    } else {
      resetQuote()
    }
  }, [
    debouncedInputAmount,
    originToken,
    destinationToken,
    originChainId,
    destinationChainId,
    isInputValid,
  ])

  const handleSwitchNetwork = useCallback(async () => {
    switchNetwork(originChainId, provider)
  }, [originChainId, provider])

  const maxAmountOut = useMemo(() => {
    if (!quote || !quote.maxAmountOut) {
      return 0
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(max, destinationTokenDecimals, 4)
  }, [quote])

  const handleUserInput = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const value = cleanNumberInput(event.target.value)
      dispatch(setInputAmount(value))
    },
    [dispatch]
  )

  const handleOriginChainSelection = useCallback(
    (newOriginChain: Chain) => {
      dispatch(setOriginChainId(newOriginChain.id))
    },
    [dispatch]
  )

  const handleDestinationChainSelection = useCallback(
    (newDestinationChain: Chain) => {
      dispatch(setDestinationChainId(newDestinationChain.id))
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
          chain={chains[originChainId]}
          onChange={handleOriginChainSelection}
        />
        <div className="flex pb-2">
          <div className="flex flex-col items-start">
            <input
              className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none"
              placeholder="0"
              value={inputAmount}
              onChange={handleUserInput}
            />
            <AvailableBalance
              originChainId={originChainId}
              originToken={originToken}
              inputAmount={inputAmount}
              connectedAddress={connectedAddress}
              balances={balances}
            />
          </div>
          <TokenSelect
            label="In"
            isOrigin={true}
            token={originToken}
            onChange={handleOriginTokenSelection}
          />
        </div>
      </div>
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="To"
          chain={chains[destinationChainId]}
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
            isOrigin={false}
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
      <BridgeButton
        originChain={chains[originChainId]}
        isApproved={isApproved}
        isDisabled={!isInputValid}
        isWrongNetwork={!isConnectedNetworkSelectedOriginNetwork}
        handleApprove={approveCallback}
        handleBridge={bridgeCallback}
        handleSwitchNetwork={handleSwitchNetwork}
        isApprovalPending={approveState === ApproveCallbackState.PENDING}
        isBridgePending={bridgeState === BridgeCallbackState.PENDING}
        approveError={approveError}
        bridgeError={bridgeError}
      />
    </div>
  )
}
