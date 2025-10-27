import { createAsyncThunk } from '@reduxjs/toolkit'
import { commify } from '@ethersproject/units'
import { Address, isAddress, zeroAddress } from 'viem'

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
      destinationAddress,
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
      destinationAddress?: Address
      pausedModulesList: BridgeModulePause[]
    },
    { rejectWithValue }
  ) => {
    const toRecipient =
      destinationAddress && isAddress(destinationAddress)
        ? destinationAddress
        : address

    const allQuotes = await synapseSDK.bridgeV2({
      fromChainId,
      toChainId,
      fromToken: fromToken.addresses[fromChainId],
      toToken: toToken.addresses[toChainId],
      fromAmount: stringToBigInt(
        debouncedFromValue,
        fromToken?.decimals[fromChainId]
      ).toString(),
      fromSender: address,
      toRecipient,
      slippagePercentage: 0.1,
    })

    const pausedBridgeModules = new Set(
      pausedModulesList
        .filter((module) =>
          module.chainId ? module.chainId === fromChainId : true
        )
        .flatMap(getBridgeModuleNames)
    )
    const activeQuotes = allQuotes.filter(
      (quote) => !quote.moduleNames.some((m) => pausedBridgeModules.has(m))
    )

    if (activeQuotes.length === 0) {
      const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
      return rejectWithValue(msg)
    }

    const rfqQuote = activeQuotes.find((q) =>
      q.moduleNames.includes('SynapseRFQ')
    )

    const nonRfqQuote = activeQuotes.find(
      (quote) => !quote.moduleNames.includes('SynapseRFQ')
    )

    let quote

    if (rfqQuote && nonRfqQuote) {
      const rfqMaxAmountOut = BigInt(rfqQuote.expectedToAmount)
      const nonRfqMaxAmountOut = BigInt(nonRfqQuote.expectedToAmount)

      const allowedPercentileDifference = 30n
      const maxDifference =
        (nonRfqMaxAmountOut * allowedPercentileDifference) / 100n

      if (rfqMaxAmountOut > nonRfqMaxAmountOut - maxDifference) {
        quote = rfqQuote
      } else {
        quote = nonRfqQuote

        const nonRfqBridgeModule =
          nonRfqQuote.moduleNames[nonRfqQuote.moduleNames.length - 1]
        segmentAnalyticsEvent(`[Bridge] use non-RFQ quote over RFQ`, {
          bridgeModuleName: nonRfqBridgeModule,
          originChainId: fromChainId,
          originToken: fromToken.symbol,
          originTokenAddress: fromToken.addresses[fromChainId],
          destinationChainId: toChainId,
          destinationToken: toToken.symbol,
          destinationTokenAddress: toToken.addresses[toChainId],
          rfqQuoteAmountOut: rfqQuote.expectedToAmount,
          nonRfqMaxAmountOut: nonRfqQuote.expectedToAmount,
        })
      }
    } else {
      quote = rfqQuote ?? nonRfqQuote
    }

    const {
      id,
      routerAddress,
      expectedToAmount,
      minToAmount,
      estimatedTime,
      moduleNames,
      gasDropAmount,
      fromChainId: originChainId,
      toChainId: destChainId,
      tx,
    } = quote

    if (!(expectedToAmount && minToAmount && toChainId !== HYPERLIQUID.id)) {
      const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
      return rejectWithValue(msg)
    }

    const toValueBigInt = BigInt(expectedToAmount) ?? 0n
    const bridgeModuleName = moduleNames[moduleNames.length - 1]

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
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
        fromToken?.decimals[fromChainId],
        toValueBigInt,
        toToken.decimals[toChainId]
      ),
      delta: toValueBigInt,
      estimatedTime,
      bridgeModuleName,
      gasDropAmount: BigInt(gasDropAmount),
      timestamp: currentTimestamp,
      originChainId,
      destChainId,
      requestId,
      id,
      tx,
    }
  }
)
