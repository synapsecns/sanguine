import { createAsyncThunk } from '@reduxjs/toolkit'
import { commify } from '@ethersproject/units'
import { BridgeableToken } from 'types'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { parseBigIntValue } from '@/utils/parseBigIntValue'
import { selectBridgeQuote } from '@/utils/selectBridgeQuote'

export const useBridgeQuoteState = (): RootState['bridgeQuote'] => {
  return useAppSelector((state) => state.bridgeQuote)
}

type SelectedBridgeQuote = {
  expectedToAmount: string
  estimatedTime: number | null
  id?: string | null
  moduleNames: string[]
  nativeFee?: unknown
  routerAddress: string
  tx?: {
    data?: string
    to?: string
    value?: string | null
  } | null
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
    const quote = selectBridgeQuote<SelectedBridgeQuote>({
      quotes: allQuotes,
      originChainId,
      pausedModules,
    })

    if (!quote) {
      return rejectWithValue('No active bridge quotes available')
    }

    const bridgeModuleName = quote.moduleNames[quote.moduleNames.length - 1]
    const toValueBigInt = BigInt(quote.expectedToAmount)
    const hasExecutableQuoteTx = Boolean(quote.tx?.to && quote.tx?.data)
    const normalizedQuoteTx = hasExecutableQuoteTx
      ? {
          data: quote.tx!.data!,
          to: quote.tx!.to!,
          value: quote.tx?.value ?? null,
        }
      : null

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
      nativeFee: parseBigIntValue(quote.nativeFee) ?? 0n,
      delta: toValueBigInt,
      estimatedTime: quote.estimatedTime,
      bridgeModuleName,
      tx: normalizedQuoteTx,
      quoteAddress:
        hasExecutableQuoteTx && connectedAddress ? connectedAddress : null,
      requestId,
      timestamp,
    }
  }
)
