import { createAsyncThunk } from '@reduxjs/toolkit'
import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { powBigInt } from '@/utils/powBigInt'
import { commify } from '@ethersproject/units'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { BridgeableToken } from 'types'

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
  }: {
    originChainId: number
    destinationChainId: number
    originToken: BridgeableToken
    destinationToken: BridgeableToken
    amount: bigint
    debouncedInputAmount: string
    synapseSDK: any
  }) => {
    const {
      feeAmount,
      routerAddress,
      maxAmountOut,
      originQuery,
      destQuery,
      estimatedTime,
      bridgeModuleName,
    } = await synapseSDK.bridgeQuote(
      originChainId,
      destinationChainId,
      originToken.addresses[originChainId],
      destinationToken.addresses[destinationChainId],
      amount
    )

    /** TODO: Handle when invalid quote returns */

    // if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
    //   // dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
    //   // dispatch(setIsLoading(false))
    //   return
    // }

    const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

    const originTokenDecimals = originToken?.decimals[originChainId]
    const adjustedFeeAmount =
      BigInt(feeAmount) <
      stringToBigInt(debouncedInputAmount, originToken?.decimals[originChainId])
        ? BigInt(feeAmount)
        : BigInt(feeAmount) / powBigInt(10n, BigInt(18 - originTokenDecimals))

    const { minAmountOut: originMinWithSlippage } =
      await synapseSDK.applySlippageInBips(originQuery, 10) // 10 bips = 0.1%

    const { minAmountOut: destMinWithSlippage } =
      await synapseSDK.applySlippageInBips(destQuery, 10) // 10 bips = 0.1%

    let newOriginQuery = { ...originQuery }
    newOriginQuery.minAmountOut = originMinWithSlippage

    let newDestQuery = { ...destQuery }
    newDestQuery.minAmountOut = destMinWithSlippage

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
      feeAmount: feeAmount,
      delta: BigInt(maxAmountOut.toString()),
      quotes: {
        originQuery: newOriginQuery,
        destQuery: newDestQuery,
      },
      estimatedTime: estimatedTime,
      bridgeModuleName: bridgeModuleName,
    }
  }
)
