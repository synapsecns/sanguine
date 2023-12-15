import { createAsyncThunk } from '@reduxjs/toolkit'
import { RootState } from '@/state/store'
import { useAppSelector } from '@/state/hooks'
import {
  fetchTokenBalances,
  TokenBalance,
} from '@/utils/actions/fetchTokenBalances'

export const useWalletState = (): RootState['wallet'] => {
  return useAppSelector((state) => state.wallet)
}

export const fetchAndStoreTokenBalances = createAsyncThunk(
  'wallet/fetchAndStoreTokenBalances',
  async ({
    address,
    chainId,
    tokens,
    signerOrProvider,
  }: {
    address: string
    chainId: number
    tokens: any[]
    signerOrProvider: any
  }) => {
    const balances: TokenBalance[] = await fetchTokenBalances({
      address,
      chainId,
      tokens,
      signerOrProvider,
    })
    return balances
  }
)
