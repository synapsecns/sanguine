import { createAsyncThunk } from '@reduxjs/toolkit'
import { RootState } from '@/state/store'
import { useAppSelector } from '@/state/hooks'
import {
  fetchTokenBalances,
  TokenBalance,
} from '@/utils/actions/fetchTokenBalances'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

export const fetchAndStoreTokenBalances = createAsyncThunk(
  'bridge/fetchAndStoreTokenBalances',
  async ({
    address,
    chainId,
    tokens,
    signer,
  }: {
    address: string
    chainId: number
    tokens: any[]
    signer: any
  }) => {
    const balances: TokenBalance[] = await fetchTokenBalances({
      address,
      chainId,
      tokens,
      signer,
    })
    return balances
  }
)
