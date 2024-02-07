import { createAsyncThunk } from '@reduxjs/toolkit'

import { useAppSelector } from '@/state/hooks'
import { RootState } from '@/state/store'
import { approveErc20Token } from '@/utils/actions/approveErc20Token'

export const useApproveTransactionState =
  (): RootState['approveTransaction'] => {
    return useAppSelector((state) => state.approveTransaction)
  }

export const executeApproveTxn = createAsyncThunk(
  'bridgeTransaction/executeApproveTxn',
  async ({
    spenderAddress,
    tokenAddress,
    amount,
    signer,
  }: {
    spenderAddress: string
    tokenAddress: string
    amount: bigint
    signer: any
  }) => {
    const approval = await approveErc20Token({
      spenderAddress,
      tokenAddress,
      amount,
      signer,
    })

    return approval
  }
)
