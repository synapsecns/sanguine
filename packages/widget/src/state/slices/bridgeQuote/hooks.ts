import { createAsyncThunk } from '@reduxjs/toolkit'
import { commify } from '@ethersproject/units'
import { BridgeableToken } from 'types'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { powBigInt } from '@/utils/powBigInt'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'

export const useBridgeQuoteState = (): RootState['bridgeQuote'] => {
  return useAppSelector((state) => state.bridgeQuote)
}

export const fetchBridgeQuote = createAsyncThunk(
  'bridgeQuote/fetchBridgeQuote',
  async ({
    originChainId,
    destinationChainId,
    originToken,
    destinationToken,
    amount,
    debouncedInputAmount,
    synapseSDK,
    requestId,
  }: {
    originChainId: number
    destinationChainId: number
    originToken: BridgeableToken
    destinationToken: BridgeableToken
    amount: bigint
    debouncedInputAmount: string
    synapseSDK: any
    requestId: number
  }) => {
    const allQuotes = await synapseSDK.allBridgeQuotes(
      originChainId,
      destinationChainId,
      originToken.addresses[originChainId],
      destinationToken.addresses[destinationChainId],
      amount
    )

    /** TODO: Handle when invalid quote returns or no quotes */

    // if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
    //   // dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
    //   // dispatch(setIsLoading(false))
    //   return
    // }

    const rfqQuote = allQuotes.find((q) => q.bridgeModuleName === 'SynapseRFQ')

    let quote

    if (rfqQuote) {
      quote = rfqQuote
    } else {
      /* allBridgeQuotes returns sorted quotes by maxAmountOut descending */
      quote = allQuotes[0]
    }

    const {
      feeAmount,
      routerAddress,
      maxAmountOut,
      originQuery,
      destQuery,
      estimatedTime,
      bridgeModuleName,
    } = quote

    const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

    const originTokenDecimals = originToken?.decimals[originChainId]
    const adjustedFeeAmount =
      BigInt(feeAmount) <
      stringToBigInt(debouncedInputAmount, originToken?.decimals[originChainId])
        ? BigInt(feeAmount)
        : BigInt(feeAmount) / powBigInt(10n, BigInt(18 - originTokenDecimals))

    const {
      originQuery: originQueryWithSlippage,
      destQuery: destQueryWithSlippage,
    } = synapseSDK.applyBridgeSlippage(bridgeModuleName, originQuery, destQuery)

    return {
      outputAmount: toValueBigInt,
      outputAmountString: commify(
        formatBigIntToString(
          toValueBigInt,
          destinationToken.decimals[destinationChainId],
          8
        )
      ),
      routerAddress,
      exchangeRate: calculateExchangeRate(
        stringToBigInt(
          debouncedInputAmount,
          originToken?.decimals[originChainId]
        ) - BigInt(adjustedFeeAmount),
        originToken?.decimals[originChainId],
        toValueBigInt,
        destinationToken.decimals[destinationChainId]
      ),
      feeAmount,
      delta: BigInt(maxAmountOut.toString()),
      quotes: {
        originQuery: originQueryWithSlippage,
        destQuery: destQueryWithSlippage,
      },
      estimatedTime,
      bridgeModuleName,
      requestId,
    }
  }
)
