import { createAsyncThunk } from '@reduxjs/toolkit'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { getTimeMinutesFromNow } from '@/utils/getTimeMinutesFromNow'
import { type BridgeQuoteTransaction } from '@/state/slices/bridgeQuote/reducer'

export const useBridgeTransactionState = (): RootState['bridgeTransaction'] => {
  return useAppSelector((state) => state.bridgeTransaction)
}

export const executeBridgeTxn = createAsyncThunk(
  'bridgeTransaction/executeBridgeTxn',
  async ({
    originChainId,
    destinationChainId,
    parsedOriginAmount,
    originTokenSymbol,
    bridgeModuleName,
    estimatedTime,
    quoteTx,
    signer,
  }: {
    originChainId: number
    destinationChainId: number
    parsedOriginAmount: string
    originTokenSymbol: string
    estimatedTime: number
    bridgeModuleName: string
    quoteTx: BridgeQuoteTransaction
    signer: any
  }) => {
    if (!quoteTx?.to || !quoteTx?.data) {
      throw new Error('Bridge quote transaction missing')
    }

    const payload = {
      data: quoteTx.data,
      to: quoteTx.to,
      ...(quoteTx.value !== undefined && quoteTx.value !== null
        ? { value: quoteTx.value }
        : {}),
    }

    const tx = await signer.sendTransaction(payload)

    const receipt = await tx.wait()

    const txHash = receipt?.hash ?? receipt?.transactionHash

    const timestamp = getTimeMinutesFromNow(0)

    return {
      txHash,
      bridgeModuleName,
      parsedOriginAmount,
      originTokenSymbol,
      originChainId,
      destinationChainId,
      estimatedTime,
      timestamp,
    }
  }
)
