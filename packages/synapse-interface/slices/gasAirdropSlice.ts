import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

type CoingeckoData = {
  symbol: string
  price: number
}

export interface GasAirdropState {
  prices: CoingeckoData[]
  isLoadingPrices: boolean
}

const initialState: GasAirdropState = {
  prices: [],
  isLoadingPrices: false,
}

export const fetchGasAirdropPrices = createAsyncThunk(
  'gasAirdrop/fetchGasAirdropPrices',
  async () => {
    const ids = [
      'ethereum',
      'avalanche-2',
      'defi-kingdoms',
      'moonriver',
      'moonbeam',
      'canto',
      'fantom',
      'metis-token',
      'binancecoin',
      'matic-network',
      'klay-token',
    ]

    try {
      const prices = await fetch(
        `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=${ids.join(
          ','
        )}`
      )
      return prices.json()
    } catch {
      return []
    }
  }
)

export const gasAirdropSlice = createSlice({
  name: 'gasAirdrop',
  initialState,
  reducers: {
    resetPriceData: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchGasAirdropPrices.pending, (state) => {
        state.isLoadingPrices = true
      })
      .addCase(fetchGasAirdropPrices.fulfilled, (state, action) => {
        state.isLoadingPrices = false

        const prices = action.payload.map((price) => ({
          symbol: price.symbol.toUpperCase(),
          price: price.current_price,
        }))

        state.prices = prices
      })
      .addCase(fetchGasAirdropPrices.rejected, (state) => {
        state.isLoadingPrices = false
        console.error('Error fetching  prices')
      })
  },
})

export default gasAirdropSlice.reducer
