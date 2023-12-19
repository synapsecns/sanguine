import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { createAsyncThunk } from '@reduxjs/toolkit'
import { ZeroAddress } from 'ethers'

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
    originQuery,
    destinationQuery,
    signer,
    synapseSDK,
  }: {
    destinationAddress: string
    originRouterAddress: string
    originChainId: number
    destinationChainId: number
    tokenAddress: string
    amount: bigint
    originQuery: {}
    destinationQuery: {}
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
        : data

    const transactionHash = await signer.sendTransaction(payload)

    return transactionHash
  }
)
