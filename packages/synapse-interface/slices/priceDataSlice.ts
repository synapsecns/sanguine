import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import {
  getSynPrices,
  getAvaxPrice,
  getEthPrice,
  getMetisPrice,
} from '@/utils/actions/getPrices'

export interface PriceDataState {
  synPrices: {
    ethBalanceNumber: number
    ethPrice: number
    synBalanceNumber: number
    synPrice: number
  }
  ethPrice: number
  avaxPrice: number
  metisPrice: number
  isLoadingSynPrices: boolean
  isLoadingEthPrice: boolean
  isLoadingAvaxPrice: boolean
  isLoadingMetisPrice: boolean
}

const initialState: PriceDataState = {
  synPrices: {
    ethBalanceNumber: null,
    ethPrice: null,
    synBalanceNumber: null,
    synPrice: null,
  },
  ethPrice: null,
  avaxPrice: null,
  metisPrice: null,
  isLoadingSynPrices: false,
  isLoadingEthPrice: false,
  isLoadingAvaxPrice: false,
  isLoadingMetisPrice: false,
}

export const fetchSynPrices = createAsyncThunk(
  'priceData/fetchSynPrices',
  async () => {
    const synPrices = await getSynPrices()
    return synPrices
  }
)

export const fetchEthPrice = createAsyncThunk(
  'priceData/fetchEthPrice',
  async () => {
    const ethPrice = await getEthPrice()
    return ethPrice
  }
)

export const fetchAvaxPrice = createAsyncThunk(
  'priceData/fetchAvaxPrice',
  async () => {
    const avaxPrice = await getAvaxPrice()
    return avaxPrice
  }
)

export const fetchMetisPrice = createAsyncThunk(
  'priceData/fetchMetisPrice',
  async () => {
    const metisPrice = await getMetisPrice()
    return metisPrice
  }
)

export const priceDataSlice = createSlice({
  name: 'priceData',
  initialState,
  reducers: {
    resetPriceData: () => initialState,
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchSynPrices.pending, (state) => {
        state.isLoadingSynPrices = true
      })
      .addCase(fetchSynPrices.fulfilled, (state, action) => {
        state.isLoadingSynPrices = false
        state.synPrices = action.payload
      })
      .addCase(fetchSynPrices.rejected, (state) => {
        state.isLoadingSynPrices = false
        console.error('Error fetching Syn prices')
      })
      .addCase(fetchEthPrice.pending, (state) => {
        state.isLoadingEthPrice = true
      })
      .addCase(fetchEthPrice.fulfilled, (state, action) => {
        state.isLoadingEthPrice = false
        state.ethPrice = action.payload
      })
      .addCase(fetchEthPrice.rejected, (state) => {
        state.isLoadingEthPrice = false
        console.error('Error fetching Eth price')
      })
      .addCase(fetchAvaxPrice.pending, (state) => {
        state.isLoadingAvaxPrice = true
      })
      .addCase(fetchAvaxPrice.fulfilled, (state, action) => {
        state.isLoadingAvaxPrice = false
        state.avaxPrice = action.payload
      })
      .addCase(fetchAvaxPrice.rejected, (state) => {
        state.isLoadingAvaxPrice = false
        console.error('Error fetching Avax price')
      })
      .addCase(fetchMetisPrice.pending, (state) => {
        state.isLoadingMetisPrice = true
      })
      .addCase(fetchMetisPrice.fulfilled, (state, action) => {
        state.isLoadingMetisPrice = false
        state.metisPrice = action.payload
      })
      .addCase(fetchMetisPrice.rejected, (state) => {
        state.isLoadingMetisPrice = false
        console.error('Error fetching Metis price')
      })
  },
})

export const { resetPriceData } = priceDataSlice.actions

export default priceDataSlice.reducer
