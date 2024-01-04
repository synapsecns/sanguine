import { createAsyncThunk } from '@reduxjs/toolkit'
import { RootState } from '@/state/store'
import { useAppSelector } from '@/state/hooks'
import {
  fetchTokenBalances,
  TokenBalance,
} from '@/utils/actions/fetchTokenBalances'
import { fetchErc20TokenAllowance } from '@/utils/actions/fetchErc20TokenAllowance'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { BridgeableToken } from 'types'

export const useWalletState = (): RootState['wallet'] => {
  return useAppSelector((state) => state.wallet)
}

export const fetchAndStoreTokenBalances = createAsyncThunk(
  'wallet/fetchAndStoreTokenBalances',
  async (
    {
      address,
      chainId,
      tokens,
      signerOrProvider,
    }: {
      address: string
      chainId: number
      tokens: BridgeableToken[]
      signerOrProvider: any
    },
    { rejectWithValue }
  ) => {
    try {
      const balances: TokenBalance[] = await fetchTokenBalances({
        address,
        chainId,
        tokens,
        signerOrProvider,
      })

      /** Throw and store error when response invalid */
      if (!Array.isArray(balances)) {
        throw new Error(balances)
      }

      return balances
    } catch (error) {
      return rejectWithValue(error?.message)
    }
  }
)

export const fetchAndStoreAllowance = createAsyncThunk(
  'wallet/fetchAndStoreAllowance',
  async ({
    spenderAddress,
    ownerAddress,
    provider,
    token,
    chainId,
  }: {
    spenderAddress: string
    ownerAddress: string
    provider: any
    token: BridgeableToken
    chainId: number
  }) => {
    const tokenAddress = token?.addresses[chainId]

    const allowance: bigint = await fetchErc20TokenAllowance({
      spenderAddress,
      tokenAddress,
      ownerAddress,
      provider,
    })

    return formatBigIntToString(allowance, token.decimals[chainId])
  }
)
