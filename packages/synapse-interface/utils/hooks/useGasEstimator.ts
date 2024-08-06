import { isNumber } from 'lodash'
import { zeroAddress, Address } from 'viem'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { estimateGas } from '@wagmi/core'
import { PayloadAction } from '@reduxjs/toolkit'

import { useAppDispatch, useAppSelector } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { calculateGasCost } from '../calculateGasCost'
import { stringToBigInt, formatBigIntToString } from '../bigint/format'
import { Token } from '../types'
import { wagmiConfig } from '@/wagmiConfig'
import {
  fetchGasData,
  setGasLimit,
  resetGasLimit,
  setIsLoadingGasLimit,
  GasDataState,
} from '@/slices/gasDataSlice'

export const useGasEstimator = () => {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken } = useBridgeState()
  const { gasData, gasLimit, isLoadingGasLimit } = useAppSelector(
    (state) => state.gasData
  )
  const { maxFeePerGas } = gasData?.formatted
  const { rawGasCost, parsedGasCost } = calculateGasCost(
    maxFeePerGas,
    gasLimit?.toString(),
    fromChainId
  )

  const { addresses, decimals } = fromToken || {}
  const tokenAddress = addresses?.[fromChainId]
  const tokenDecimals = isNumber(decimals) ? decimals : decimals?.[fromChainId]
  const isGasToken: boolean = tokenAddress === zeroAddress
  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === tokenAddress
  )
  const parsedBalance = formatBigIntToString(
    selectedFromToken?.balance,
    tokenDecimals
  )

  const gasFeeExceedsBalance: boolean =
    isGasToken &&
    parsedGasCost &&
    parsedBalance &&
    parseFloat(parsedGasCost) > parseFloat(parsedBalance)

  const maxBridgeableGas: number | null =
    isGasToken && parsedGasCost
      ? calculateMaxBridgeableGas(
          parseFloat(parsedBalance),
          parseFloat(parsedGasCost)
        )
      : null

  const hasValidGasEstimateInputs = (): boolean => {
    if (!fromChainId || !toChainId) return false
    if (!fromToken || !toToken) return false
    if (!isGasToken) return false
    if (!selectedFromToken || !parsedBalance) return false
    return true
  }

  const estimateGasLimit = async () => {
    if (hasValidGasEstimateInputs()) {
      dispatch(resetGasLimit())
      dispatch(setIsLoadingGasLimit(true))
      try {
        const gasLimit = await queryEstimatedBridgeGasLimit(
          synapseSDK,
          address,
          address,
          fromChainId,
          toChainId,
          fromToken,
          toToken,
          parsedBalance
        )
        dispatch(setGasLimit(gasLimit))
        return gasLimit
      } catch (error) {
        console.error('Error estimating gas limit:', error)
        dispatch(resetGasLimit())
      }
    }
  }

  const estimateBridgeableBalance = async () => {
    const gasData = (await dispatch(
      fetchGasData(fromChainId)
    )) as PayloadAction<GasDataState>
    const { maxFeePerGas } = gasData?.payload?.gasData.formatted || {}

    const gasLimit = await estimateGasLimit()
    const { parsedGasCost } = calculateGasCost(
      maxFeePerGas,
      gasLimit?.toString(),
      fromChainId
    )
    const maxBridgeableGas: number | null =
      isGasToken && parsedGasCost
        ? calculateMaxBridgeableGas(
            parseFloat(parsedBalance),
            parseFloat(parsedGasCost)
          )
        : null

    return maxBridgeableGas
  }

  // Reset gas limit when chainId changes
  useEffect(() => {
    dispatch(resetGasLimit())
  }, [fromChainId])

  return {
    isLoading: isLoadingGasLimit,
    isGasToken,
    rawGasCost,
    parsedGasCost,
    maxBridgeableGas,
    gasFeeExceedsBalance,
    hasValidGasEstimateInputs,
    estimateGasLimitCallback: estimateGasLimit,
    estimateBridgeableBalanceCallback: estimateBridgeableBalance,
  }
}

const calculateMaxBridgeableGas = (
  parsedGasBalance: number,
  parsedGasCost: number
): number => {
  const maxBridgeable = parsedGasBalance - parsedGasCost
  return maxBridgeable > 0 ? maxBridgeable : 0
}

const getBridgeQuote = async (
  synapseSDK: any,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  toToken: Token,
  amount: string,
  userAddress: string
) => {
  try {
    return await synapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId],
      toToken.addresses[toChainId],
      stringToBigInt(amount, fromToken?.decimals[fromChainId]),
      {
        originUserAddress: userAddress,
      }
    )
  } catch (error) {
    console.error('getBridgeQuote: ', error)
    return null
  }
}

const getBridgePayload = async (
  synapseSDK: any,
  bridgeQuote: any | null,
  address: string,
  toAddress: string,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  amount: string
) => {
  if (!bridgeQuote) return null

  try {
    const data = await synapseSDK.bridge(
      address,
      bridgeQuote.routerAddress,
      fromChainId,
      toChainId,
      fromToken?.addresses[fromChainId as keyof Token['addresses']],
      stringToBigInt(amount, fromToken?.decimals[fromChainId]),
      bridgeQuote.originQuery,
      bridgeQuote.destQuery
    )

    const payload =
      fromToken?.addresses[fromChainId as keyof Token['addresses']] ===
        zeroAddress ||
      fromToken?.addresses[fromChainId as keyof Token['addresses']] === ''
        ? {
            data: data.data,
            to: data.to,
            value: stringToBigInt(amount, fromToken?.decimals[fromChainId]),
          }
        : data

    return payload
  } catch (error) {
    console.error('getBridgePayload: ', error)
    return null
  }
}

const getBridgeGasLimitEstimate = async (
  bridgePayload: any,
  fromChainId: number,
  address: string
) => {
  if (!bridgePayload) return null

  try {
    const gasEstimate = await estimateGas(wagmiConfig, {
      value: bridgePayload.value,
      to: bridgePayload.to,
      account: address as Address,
      data: bridgePayload.data,
      chainId: fromChainId,
    })

    return gasEstimate
  } catch (error) {
    console.error('calculateEstimatedBridgeGasLimit: ', error)
    return null
  }
}

const queryEstimatedBridgeGasLimit = async (
  synapseSDK: any,
  address: string,
  toAddress: string,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  toToken: Token,
  amount: string
) => {
  const bridgeQuote = await getBridgeQuote(
    synapseSDK,
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    amount,
    address // userAddress
  )

  const bridgePayload = await getBridgePayload(
    synapseSDK,
    bridgeQuote,
    address,
    toAddress,
    fromChainId,
    toChainId,
    fromToken,
    amount
  )

  const gasLimit = await getBridgeGasLimitEstimate(
    bridgePayload,
    fromChainId,
    address
  )

  return gasLimit
}
