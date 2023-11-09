import { useMemo } from 'react'
import { createAsyncThunk } from '@reduxjs/toolkit'
import { useAccount, Address, useNetwork } from 'wagmi'
import { zeroAddress } from 'viem'

import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import {
  fetchBridgeQuote,
  fetchBridgeQuotes,
  BridgeQuoteRequest,
  BridgeQuoteResponse,
} from '@/utils/actions/fetchBridgeQuotes'
import { BridgeState, initialState } from './reducer'
import { usePortfolioState } from '../portfolio/hooks'
import { PortfolioState } from '../portfolio/reducer'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { TokenWithBalanceAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import { calculateEstimatedTransactionTime } from '@/utils/calculateEstimatedTransactionTime'
import { useGasDropAmount } from '@/utils/hooks/useGasDropAmount'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

// Bridge status based on user inputs
export const useBridgeStatus = (): {
  isConnected: boolean
  hasValidSelections: boolean
  hasValidRoute: boolean
  hasEnoughBalance: boolean
  hasInputAmount: boolean
  hasEnoughApproved: boolean
  hasSelectedNetwork: boolean
  onSelectedChain: boolean
} => {
  const { isConnected } = useAccount()
  const { chain } = useNetwork()
  const {
    debouncedFromValue,
    fromValue,
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    isLoading,
    bridgeQuote,
  }: BridgeState = useBridgeState()
  const { balancesAndAllowances }: PortfolioState = usePortfolioState()

  const hasValidSelections: boolean = useMemo(() => {
    return Boolean(fromChainId && toChainId && fromToken && toToken)
  }, [fromChainId, toChainId, fromToken, toToken])

  const hasValidRoute: boolean = useMemo(() => {
    const hasInput: boolean =
      debouncedFromValue !== initialState.debouncedFromValue
    const hasQuote: boolean = bridgeQuote !== initialState.bridgeQuote

    return Boolean(!isLoading && hasInput && hasQuote)
  }, [isLoading, bridgeQuote, debouncedFromValue])

  const hasEnoughBalance: boolean = useMemo(() => {
    const currentBalance: bigint = balancesAndAllowances[fromChainId]?.find(
      (token: TokenWithBalanceAndAllowances) => token.token === fromToken
    )?.balance
    const preciseFromValue: bigint = stringToBigInt(
      fromValue,
      fromToken?.decimals[fromChainId]
    )

    return currentBalance >= preciseFromValue
  }, [balancesAndAllowances, fromValue, fromToken, fromChainId])

  const hasInputAmount: boolean = useMemo(() => {
    const isEmpty: boolean = fromValue === initialState.fromValue

    return Boolean(!hasOnlyZeroes(fromValue) && !isEmpty)
  }, [fromValue])

  const hasEnoughApproved: boolean = useMemo(() => {
    if (fromToken?.addresses[fromChainId] === zeroAddress) {
      return true
    }
    const approved = balancesAndAllowances[fromChainId]?.find(
      (token: TokenWithBalanceAndAllowances) => token.token === fromToken
    )?.allowances[bridgeQuote?.routerAddress]

    return (
      fromToken &&
      approved >= stringToBigInt(fromValue, fromToken?.decimals[fromChainId])
    )
  }, [balancesAndAllowances, bridgeQuote, fromValue, fromToken])

  const hasSelectedNetwork: boolean = useMemo(() => {
    return fromChainId === chain?.id
  }, [fromChainId, chain])

  const onSelectedChain: boolean = useMemo(() => {
    return chain?.id === fromChainId
  }, [fromChainId, chain])

  return {
    isConnected,
    hasValidSelections,
    hasValidRoute,
    hasEnoughBalance,
    hasInputAmount,
    hasEnoughApproved,
    hasSelectedNetwork,
    onSelectedChain,
  }
}

/**
 * Bridge additional details based on current inputs
 * Includes gas fee, est duration, gas airdrop, received percentage
 */
export const useBridgeDetails = (): {
  estimatedBridgeDurationInSeconds: number
  isGasDropped: boolean
  isGasLoading: boolean
  gasDropAmount: bigint | any
  receivedPercentageOfInitial: number
  isUnderFee: boolean
  parsedExchangeRate: number
} => {
  const {
    fromChainId,
    fromToken,
    fromValue,
    debouncedFromValue,
    bridgeQuote,
    toChainId,
  }: BridgeState = useBridgeState()
  const { hasInputAmount } = useBridgeStatus()
  const { gasDrop: gasDropAmount, loading: isGasLoading } =
    useGasDropAmount(toChainId)

  const originTokenAddress: Address = useMemo(
    () => fromToken?.addresses[fromChainId] as Address,
    [fromChainId, fromToken]
  )

  const estimatedBridgeDurationInSeconds: number = useMemo(
    () =>
      calculateEstimatedTransactionTime({
        originChainId: fromChainId,
        originTokenAddress,
      }),
    [originTokenAddress, fromChainId]
  )

  const isGasDropped = useMemo(() => {
    if (gasDropAmount) {
      return gasDropAmount.gt(0)
    }
  }, [gasDropAmount])

  const receivedPercentageOfInitial: number = useMemo(() => {
    const outputAmount: string = bridgeQuote?.outputAmountString
    const inputAmount: string = fromValue

    return Number(inputAmount) / Number(outputAmount)
  }, [bridgeQuote, debouncedFromValue])

  const { isUnderFee, parsedExchangeRate } = useMemo(() => {
    const exchangeRate: bigint = bridgeQuote?.exchangeRate ?? 0n
    const formattedExchangeRate: string = formatBigIntToString(
      exchangeRate,
      18,
      4
    )
    const parsedExchangeRate: number = Number(formattedExchangeRate)

    return {
      isUnderFee: exchangeRate === 0n && hasInputAmount,
      parsedExchangeRate,
    }
  }, [bridgeQuote, hasInputAmount])

  return {
    estimatedBridgeDurationInSeconds,
    isGasDropped,
    isGasLoading,
    gasDropAmount,
    receivedPercentageOfInitial,
    isUnderFee,
    parsedExchangeRate,
  }
}

export const fetchAndStoreBridgeQuote = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuote',
  async ({
    request,
    synapseSDK,
  }: {
    request: BridgeQuoteRequest
    synapseSDK: any
  }) => {
    const bridgeQuote: BridgeQuoteResponse = await fetchBridgeQuote(
      request,
      synapseSDK
    )
    return bridgeQuote
  }
)

export const fetchAndStoreBridgeQuotes = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuotes',
  async ({
    requests,
    synapseSDK,
  }: {
    requests: BridgeQuoteRequest[]
    synapseSDK: any
  }) => {
    const bridgeQuotes: BridgeQuoteResponse[] = await fetchBridgeQuotes(
      requests,
      synapseSDK
    )
    return bridgeQuotes
  }
)
