import { isNumber } from 'lodash'
import { zeroAddress, Address } from 'viem'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'
import { estimateGas } from '@wagmi/core'

import { useAppSelector } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { calculateGasCost } from '../calculateGasCost'
import { stringToBigInt, formatBigIntToString } from '../bigint/format'
import { Token } from '../types'
import { wagmiConfig } from '@/wagmiConfig'

export const useGasEstimator = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const { balances } = usePortfolioState()
  const { fromChainId, toChainId, fromToken, toToken } = useBridgeState()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [estimatedGasLimit, setEstimatedGasLimit] = useState<bigint>(0n)
  const { gasData } = useAppSelector((state) => state.gasData)
  const { maxFeePerGas } = gasData?.formatted
  const { rawGasCost, parsedGasCost } = calculateGasCost(
    maxFeePerGas,
    estimatedGasLimit.toString()
  )

  const { addresses, decimals } = fromToken || {}
  const tokenAddress = addresses?.[fromChainId]
  const tokenDecimals = isNumber(decimals) ? decimals : decimals?.[fromChainId]
  const isGasToken: boolean = tokenAddress === zeroAddress
  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
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
    if (!selectedFromToken) return false
    if (!parsedBalance) return false
    return true
  }

  const estimateGasLimit = async () => {
    if (hasValidGasEstimateInputs()) {
      setEstimatedGasLimit(0n)
      setIsLoading(true)
      try {
        const gasLimit = await queryEstimatedBridgeGasLimit(
          synapseSDK,
          address,
          address,
          fromChainId,
          toChainId,
          fromToken,
          toToken,
          selectedFromToken?.parsedBalance
        )
        setEstimatedGasLimit(gasLimit ?? 0n)
      } catch (error) {
        console.error('Error estimating gas limit:', error)
        setEstimatedGasLimit(0n)
      } finally {
        setIsLoading(false)
      }
    }
  }

  /** Fetch gasLimit using Wallet's gas balance */
  useEffect(() => {
    estimateGasLimit()
  }, [
    fromChainId,
    toChainId,
    isGasToken,
    selectedFromToken,
    fromToken,
    toToken,
    address,
  ])

  return {
    rawGasCost,
    parsedGasCost,
    maxBridgeableGas,
    isLoading,
    gasFeeExceedsBalance,
    estimateGasLimitCallback: estimateGasLimit,
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
  amount: string
) => {
  try {
    return await synapseSDK.bridgeQuote(
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId],
      toToken.addresses[toChainId],
      stringToBigInt(amount, fromToken?.decimals[fromChainId])
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

const calculateEstimatedBridgeGasLimit = async (
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
    amount
  )

  const bridgePayload = await getBridgePayload(
    synapseSDK,
    bridgeQuote,
    address,
    address,
    fromChainId,
    toChainId,
    fromToken,
    amount
  )

  const gasLimit = await calculateEstimatedBridgeGasLimit(
    bridgePayload,
    fromChainId,
    address
  )

  return gasLimit
}
