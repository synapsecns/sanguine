import { createAsyncThunk } from '@reduxjs/toolkit'
import { ZeroAddress } from 'ethers'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { getTimeMinutesFromNow } from '@/utils/getTimeMinutesFromNow'

export const useBridgeTransactionState = (): RootState['bridgeTransaction'] => {
  return useAppSelector((state) => state.bridgeTransaction)
}

export const executeBridgeTxn = createAsyncThunk(
  'bridgeTransaction/executeBridgeTxn',
  async ({
    destinationAddress,
    originRouterAddress,
    originChainId,
    destinationChainId,
    tokenAddress,
    amount,
    parsedOriginAmount,
    originTokenSymbol,
    originQuery,
    destinationQuery,
    bridgeModuleName,
    estimatedTime,
    signer,
    synapseSDK,
  }: {
    destinationAddress: string
    originRouterAddress: string
    originChainId: number
    destinationChainId: number
    tokenAddress: string
    amount: bigint
    parsedOriginAmount: string
    originTokenSymbol: string
    originQuery: {}
    destinationQuery: {}
    estimatedTime: number
    bridgeModuleName: string
    signer: any
    synapseSDK: any
  }) => {
    const data = await synapseSDK.bridge(
      destinationAddress,
      originRouterAddress,
      originChainId,
      destinationChainId,
      tokenAddress,
      amount,
      originQuery,
      destinationQuery
    )

    const payload =
      tokenAddress === ZeroAddress
        ? {
            data: data.data,
            to: data.to,
            value: amount,
          }
        : {
            data: data.data,
            to: data.to,
          }

    const tx = await signer.sendTransaction(payload)

    const receipt = await tx.wait()

    const txHash = receipt?.hash

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
