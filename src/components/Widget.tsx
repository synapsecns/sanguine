import { useMemo, useEffect, useContext, useCallback } from 'react'
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
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { fetchAndStoreTokenBalances } from '@/state/slices/wallet/hooks'
import { BridgeButton } from './BridgeButton'

import { generateTheme } from '@/utils/generateTheme'
import { AvailableBalance } from './AvailableBalance'
import { ZeroAddress } from 'ethers'
import { checkExists } from '@/utils/checkExists'
import { useCurrentTokenBalance } from '@/hooks/useCurrentTokenBalance'
import { useValidations } from '@/hooks/useValidations'

import { Transaction } from './Transaction'

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
  networkProviders,
  theme,
  customTheme,
  tokens,
  toChainId,
}) => {
  const dispatch = useAppDispatch()
  const synapseSDK = new SynapseSDK(chainIds, networkProviders)
  const web3Context = useContext(Web3Context)
  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  const {
    inputAmount,
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
    tokens: allTokens,
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
    amount: stringToBigInt(
      debouncedInputAmount,
      originToken?.decimals[originChainId]
    ),
    synapseSDK: synapseSDK,
  })

  const { allowance, checkAllowanceCallback } = useAllowance({
    spenderAddress: quote?.routerAddress,
    tokenAddress: originToken?.addresses[originChainId],
    ownerAddress: connectedAddress,
    chainId: originToken?.addresses[originChainId],
    provider: originChainProvider ?? provider,
  })

  const useApproveCallbackArgs: UseApproveCallbackProps = {
    spenderAddress: quote?.routerAddress,
    tokenAddress: originToken?.addresses[originChainId],
    ownerAddress: connectedAddress,
    amount: stringToBigInt(inputAmount, originToken?.decimals[originChainId]),
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
    originRouterAddress: quote?.routerAddress,
    originChainId: originChainId,
    destinationChainId: destinationChainId,
    tokenAddress: originToken?.addresses[originChainId],
    amount: stringToBigInt(
      debouncedInputAmount,
      originToken?.decimals[originChainId]
    ),
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

  const formattedInputAmount: bigint = useMemo(() => {
    return stringToBigInt(
      debouncedInputAmount ?? '0',
      originToken?.decimals[originChainId]
    )
  }, [debouncedInputAmount, originToken])

  const isApproved: boolean = useMemo(() => {
    if (originToken?.addresses[originChainId] === ZeroAddress) {
      return true
    }
    if (!checkExists(allowance)) return true
    if (!formattedInputAmount) return true
    return formattedInputAmount <= allowance
  }, [formattedInputAmount, allowance, originToken, originChainId])

  const currentTokenBalance = useCurrentTokenBalance()

  const { hasEnoughBalance, isInputValid } = useValidations()

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

  const maxAmountOut = useMemo(() => {
    if (!quote || !quote.maxAmountOut) {
      return 0
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(
      max,
      destinationToken?.decimals[destinationChainId],
      4
    )
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
      <Transaction
        originChainId={42161}
        destinationChainId={137}
        originTxHash="0x6c25a451f4fe26742eeafe2475a190a5c9a6cf6b6ab9cecd10348be506402f66"
        destinationTxHash="0x2e6d03f06b3ca74a681e48a1d3cba3fa62172f3a00f1385e1084602838154540"
        kappa="6cb14bf1a4914aac28ef173dc00427ed815306f15c495688921e8648176bb2a4"
      />
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="From"
          chain={chains[originChainId]}
          onChange={handleOriginChainSelection}
        />
        <div className="flex">
          <input
            className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none"
            placeholder="0"
            value={inputAmount}
            onChange={handleUserInput}
          />
          <div className="flex flex-col items-end space-y-1">
            <TokenSelect
              label="In"
              isOrigin={true}
              token={originToken}
              onChange={handleOriginTokenSelection}
            />
            <AvailableBalance
              originChainId={originChainId}
              originToken={originToken}
              tokenBalance={currentTokenBalance}
              connectedAddress={connectedAddress}
              hasEnoughBalance={hasEnoughBalance}
            />
          </div>
        </div>
      </div>
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="To"
          chain={chains[destinationChainId]}
          onChange={handleDestinationChainSelection}
        />
        <div className="flex items-center justify-between">
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
          stringToBigInt(inputAmount, originToken?.decimals[originChainId]),
          originToken?.decimals[originChainId],
          4
        )}
        receive={maxAmountOut}
      />
      <BridgeButton
        originChain={chains[originChainId]}
        isApproved={isApproved}
        isValidQuote={Boolean(quote)}
        handleApprove={approveCallback}
        handleBridge={bridgeCallback}
        isApprovalPending={approveState === ApproveCallbackState.PENDING}
        isBridgePending={bridgeState === BridgeCallbackState.PENDING}
        approveError={approveError}
        bridgeError={bridgeError}
      />
    </div>
  )
}
