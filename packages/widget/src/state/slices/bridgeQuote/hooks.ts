import { createAsyncThunk } from '@reduxjs/toolkit'
import { commify } from '@ethersproject/units'
import { BridgeableToken } from 'types'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'

export const useBridgeQuoteState = (): RootState['bridgeQuote'] => {
  return useAppSelector((state) => state.bridgeQuote)
}

const DECIMAL_BIGINT_PATTERN = /^-?\d+$/

const parseNativeFee = (nativeFee: unknown): bigint => {
  if (typeof nativeFee === 'bigint') {
    return nativeFee
  }

  if (typeof nativeFee === 'string') {
    const normalizedNativeFee = nativeFee.trim()

    if (DECIMAL_BIGINT_PATTERN.test(normalizedNativeFee)) {
      return BigInt(normalizedNativeFee)
    }
  }

  return 0n
}

export const fetchBridgeQuote = createAsyncThunk(
  'bridgeQuote/fetchBridgeQuote',
  async (
    {
      originChainId,
      destinationChainId,
      originToken,
      destinationToken,
      amount,
      debouncedInputAmount,
      synapseSDK,
      requestId,
      pausedModules,
      timestamp,
      connectedAddress,
    }: {
      originChainId: number
      destinationChainId: number
      originToken: BridgeableToken
      destinationToken: BridgeableToken
      amount: bigint
      debouncedInputAmount: string
      synapseSDK: any
      requestId: number
      pausedModules: any
      timestamp: number
      connectedAddress?: string
    },
    { rejectWithValue }
  ) => {
    const quoteParams: {
      fromChainId: number
      toChainId: number
      fromToken: string
      toToken: string
      fromAmount: string
      slippagePercentage: number
      fromSender?: string
      toRecipient?: string
    } = {
      fromChainId: originChainId,
      toChainId: destinationChainId,
      fromToken: originToken.addresses[originChainId],
      toToken: destinationToken.addresses[destinationChainId],
      fromAmount: amount.toString(),
      slippagePercentage: 0.1,
    }

    if (connectedAddress) {
      quoteParams.fromSender = connectedAddress
      quoteParams.toRecipient = connectedAddress
    }

    const allQuotes = await synapseSDK.bridgeV2(quoteParams)

    const pausedBridgeModules = new Set(
      pausedModules
        .filter((module) =>
          module.chainId ? module.chainId === originChainId : true
        )
        .flatMap(getBridgeModuleNames)
    )

    const activeQuotes = allQuotes.filter(
      (fetchedQuote) =>
        !fetchedQuote.moduleNames.some((moduleName) =>
          pausedBridgeModules.has(moduleName)
        )
    )

    const rfqQuote = activeQuotes.find((activeQuote) =>
      activeQuote.moduleNames.includes('SynapseRFQ')
    )

    const quote = rfqQuote ?? activeQuotes[0]

    if (!quote) {
      return rejectWithValue('No active bridge quotes available')
    }

    const bridgeModuleName = quote.moduleNames[quote.moduleNames.length - 1]
    const toValueBigInt = BigInt(quote.expectedToAmount)
    const hasExecutableQuoteTx = Boolean(quote.tx?.to && quote.tx?.data)

    return {
      id: quote.id ?? null,
      outputAmount: toValueBigInt,
      outputAmountString: commify(
        formatBigIntToString(
          toValueBigInt,
          destinationToken.decimals[destinationChainId],
          8
        )
      ),
      routerAddress: quote.routerAddress,
      exchangeRate: calculateExchangeRate(
        stringToBigInt(
          debouncedInputAmount,
          originToken?.decimals[originChainId]
        ),
        originToken?.decimals[originChainId],
        toValueBigInt,
        destinationToken.decimals[destinationChainId]
      ),
      feeAmount: 0n,
      nativeFee: parseNativeFee(quote.nativeFee),
      delta: toValueBigInt,
      estimatedTime: quote.estimatedTime,
      bridgeModuleName,
      tx: quote.tx ?? null,
      quoteAddress:
        hasExecutableQuoteTx && connectedAddress ? connectedAddress : null,
      requestId,
      timestamp,
    }
  }
)
