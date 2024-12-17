import { createAsyncThunk } from '@reduxjs/toolkit'
import { commify } from '@ethersproject/units'
import { Address, zeroAddress } from 'viem'

import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { AcceptedChainId, CHAINS_BY_ID } from '@/constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { stringToBigInt, formatBigIntToString } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'
import { Token } from '@/utils/types'
import { BridgeModulePause } from '@/components/Maintenance/Maintenance'
import { HYPERLIQUID } from '@/constants/chains/master'

export const fetchBridgeQuote = createAsyncThunk(
  'bridgeQuote/fetchBridgeQuote',
  async (
    {
      synapseSDK,
      fromChainId,
      toChainId,
      fromToken,
      toToken,
      debouncedFromValue,
      requestId,
      currentTimestamp,
      address,
      pausedModulesList,
    }: {
      synapseSDK: any
      fromChainId: number
      toChainId: number
      fromToken: Token
      toToken: Token
      debouncedFromValue: string
      requestId: number
      currentTimestamp: number
      address: Address
      pausedModulesList: BridgeModulePause[]
    },
    { rejectWithValue }
  ) => {
    const allQuotes = await synapseSDK.allBridgeQuotes(
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId],
      toToken.addresses[toChainId],
      stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
      {
        originUserAddress: address,
      }
    )

    const pausedBridgeModules = new Set(
      pausedModulesList
        .filter((module) =>
          module.chainId ? module.chainId === fromChainId : true
        )
        .flatMap(getBridgeModuleNames)
    )
    const activeQuotes = allQuotes.filter(
      (quote) => !pausedBridgeModules.has(quote.bridgeModuleName)
    )

    if (activeQuotes.length === 0) {
      const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
      return rejectWithValue(msg)
    }

    const rfqQuote = activeQuotes.find(
      (q) => q.bridgeModuleName === 'SynapseIntents'
    )

    const nonRfqQuote = activeQuotes.find(
      (quote) => quote.bridgeModuleName !== 'SynapseRFQ'
    )

    let quote

    if (rfqQuote && nonRfqQuote) {
      const rfqMaxAmountOut = BigInt(rfqQuote.maxAmountOut.toString())
      const nonRfqMaxAmountOut = BigInt(nonRfqQuote.maxAmountOut.toString())

      const allowedPercentileDifference = 30n
      const maxDifference =
        (nonRfqMaxAmountOut * allowedPercentileDifference) / 100n

      if (rfqMaxAmountOut > nonRfqMaxAmountOut - maxDifference) {
        quote = rfqQuote
      } else {
        quote = nonRfqQuote

        segmentAnalyticsEvent(`[Bridge] use non-RFQ quote over RFQ`, {
          bridgeModuleName: nonRfqQuote.bridgeModuleName,
          originChainId: fromChainId,
          originToken: fromToken.symbol,
          originTokenAddress: fromToken.addresses[fromChainId],
          destinationChainId: toChainId,
          destinationToken: toToken.symbol,
          destinationTokenAddress: toToken.addresses[toChainId],
          rfqQuoteAmountOut: rfqQuote.maxAmountOut.toString(),
          nonRfqMaxAmountOut: nonRfqQuote.maxAmountOut.toString(),
        })
      }
    } else {
      quote = rfqQuote ?? nonRfqQuote
    }

    const {
      id,
      feeAmount,
      routerAddress,
      maxAmountOut,
      originQuery,
      destQuery,
      estimatedTime,
      bridgeModuleName,
      gasDropAmount,
      originChainId,
      destChainId,
    } = quote

    if (
      !(
        originQuery &&
        maxAmountOut &&
        destQuery &&
        feeAmount &&
        toChainId !== HYPERLIQUID.id
      )
    ) {
      const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
      return rejectWithValue(msg)
    }

    const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

    // Bridge Lifecycle: originToken -> bridgeToken -> destToken
    // debouncedFromValue is in originToken decimals
    // originQuery.minAmountOut and feeAmount is in bridgeToken decimals
    // Adjust feeAmount to be in originToken decimals
    const adjustedFeeAmount =
      (BigInt(feeAmount) *
        stringToBigInt(
          `${debouncedFromValue}`,
          fromToken?.decimals[fromChainId]
        )) /
      BigInt(originQuery.minAmountOut)

    const isUnsupported = AcceptedChainId[fromChainId] ? false : true

    const allowance =
      fromToken?.addresses[fromChainId] === zeroAddress ||
      address === undefined ||
      isUnsupported
        ? 0n
        : await getErc20TokenAllowance({
            address,
            chainId: fromChainId,
            tokenAddress: fromToken?.addresses[fromChainId] as Address,
            spender: routerAddress,
          })

    const {
      originQuery: originQueryWithSlippage,
      destQuery: destQueryWithSlippage,
    } = synapseSDK.applyBridgeSlippage(bridgeModuleName, originQuery, destQuery)

    return {
      inputAmountForQuote: debouncedFromValue,
      originTokenForQuote: fromToken,
      destTokenForQuote: toToken,
      outputAmount: toValueBigInt,
      outputAmountString: commify(
        formatBigIntToString(toValueBigInt, toToken.decimals[toChainId], 8)
      ),
      routerAddress,
      allowance,
      exchangeRate: calculateExchangeRate(
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]) -
          BigInt(adjustedFeeAmount),
        fromToken?.decimals[fromChainId],
        toValueBigInt,
        toToken.decimals[toChainId]
      ),
      feeAmount,
      delta: BigInt(maxAmountOut.toString()),
      originQuery: originQueryWithSlippage,
      destQuery: destQueryWithSlippage,
      estimatedTime,
      bridgeModuleName,
      gasDropAmount: BigInt(gasDropAmount.toString()),
      timestamp: currentTimestamp,
      originChainId,
      destChainId,
      requestId,
      id,
    }
  }
)
