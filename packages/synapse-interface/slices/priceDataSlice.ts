import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import {
  getSynPrices,
  getAvaxPrice,
  getEthPrice,
  getMetisPrice,
  getArbPrice,
  getGmxPrice,
  getAllEthStablecoinPrices,
  getCoingeckoPrices,
  getMusdcPrice,
  getDaiePrice,
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
  arbPrice: number
  gmxPrice: number
  fraxPrice: number
  usdtPrice: number
  usdcPrice: number
  crvUsdPrice: number
  daiPrice: number
  lusdPrice: number
  notePrice: number
  susdPrice: number
  usdbcPrice: number
  usdcePrice: number
  usdtePrice: number
  musdcPrice: number
  daiePrice: number
  isLoadingSynPrices: boolean
  isLoadingEthPrice: boolean
  isLoadingAvaxPrice: boolean
  isLoadingMetisPrice: boolean
  isLoadingArbPrice: boolean
  isLoadingGmxPrice: boolean
  isLoadingAllEthPrices: boolean
  isLoadingCoingeckoPrices: boolean
  isLoadingMusdcPrice: boolean
  isLoadingDaiePrice: boolean
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
  arbPrice: null,
  gmxPrice: null,
  fraxPrice: null,
  usdtPrice: null,
  usdcPrice: null,
  crvUsdPrice: null,
  daiPrice: null,
  lusdPrice: null,
  notePrice: null,
  susdPrice: null,
  usdbcPrice: null,
  usdcePrice: null,
  usdtePrice: null,
  musdcPrice: null,
  daiePrice: null,
  isLoadingSynPrices: false,
  isLoadingEthPrice: false,
  isLoadingAvaxPrice: false,
  isLoadingMetisPrice: false,
  isLoadingArbPrice: false,
  isLoadingGmxPrice: false,
  isLoadingAllEthPrices: false,
  isLoadingCoingeckoPrices: false,
  isLoadingMusdcPrice: false,
  isLoadingDaiePrice: false,
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

export const fetchArbPrice = createAsyncThunk(
  'priceData/fetchArbPrice',
  async () => {
    const arbPrice = await getArbPrice()
    return arbPrice
  }
)

export const fetchGmxPrice = createAsyncThunk(
  'priceData/fetchGmxPrice',
  async () => {
    const gmxPrice = await getGmxPrice()
    return gmxPrice
  }
)

export const fetchAllEthStablecoinPrices = createAsyncThunk(
  'priceData/fetchAllEthStablecoinPrices',
  async () => {
    const prices = await getAllEthStablecoinPrices()
    return prices
  }
)

export const fetchCoingeckoPrices = createAsyncThunk(
  'priceData/fetchCoingeckoPrices',
  async () => {
    const prices = await getCoingeckoPrices()
    return prices
  }
)

export const fetchMusdcPrice = createAsyncThunk(
  'priceData/fetchMusdcPrice',
  async () => {
    const price = await getMusdcPrice()
    return price
  }
)

export const fetchDaiePrice = createAsyncThunk(
  'priceData/fetchDaiePrice',
  async () => {
    const price = await getDaiePrice()
    return price
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
      .addCase(fetchArbPrice.pending, (state) => {
        state.isLoadingArbPrice = true
      })
      .addCase(fetchArbPrice.fulfilled, (state, action) => {
        state.isLoadingArbPrice = false
        state.arbPrice = action.payload
      })
      .addCase(fetchArbPrice.rejected, (state) => {
        state.isLoadingArbPrice = false
        console.error('Error fetching Arb price')
      })
      .addCase(fetchGmxPrice.pending, (state) => {
        state.isLoadingGmxPrice = true
      })
      .addCase(fetchGmxPrice.fulfilled, (state, action) => {
        state.isLoadingGmxPrice = false
        state.gmxPrice = action.payload
      })
      .addCase(fetchGmxPrice.rejected, (state) => {
        state.isLoadingGmxPrice = false
        console.error('Error fetching GMX price')
      })
      .addCase(fetchAllEthStablecoinPrices.pending, (state) => {
        state.isLoadingAllEthPrices = true
      })
      .addCase(fetchAllEthStablecoinPrices.fulfilled, (state, action) => {
        state.isLoadingAllEthPrices = false
        const { usdcPrice, usdtPrice, fraxPrice, daiPrice, crvUsdPrice } =
          action.payload

        state.usdcPrice = usdcPrice
        state.usdtPrice = usdtPrice
        state.fraxPrice = fraxPrice
        state.daiPrice = daiPrice
        state.crvUsdPrice = crvUsdPrice
      })
      .addCase(fetchAllEthStablecoinPrices.rejected, (state) => {
        state.isLoadingAllEthPrices = false
        console.error('Error fetching prices on Ethereum')
      })
      .addCase(fetchCoingeckoPrices.pending, (state) => {
        state.isLoadingCoingeckoPrices = true
      })
      .addCase(fetchCoingeckoPrices.fulfilled, (state, action) => {
        state.isLoadingCoingeckoPrices = false

        const {
          notePrice,
          susdPrice,
          lusdPrice,
          usdbcPrice,
          usdcePrice,
          usdtePrice,
        } = action.payload

        state.notePrice = notePrice
        state.susdPrice = susdPrice
        state.lusdPrice = lusdPrice
        state.usdbcPrice = usdbcPrice
        state.usdcePrice = usdcePrice
        state.usdtePrice = usdtePrice
      })
      .addCase(fetchCoingeckoPrices.rejected, (state) => {
        state.isLoadingCoingeckoPrices = false
        console.error('Error fetching prices from Coingecko')
      })
      .addCase(fetchMusdcPrice.pending, (state) => {
        state.isLoadingMusdcPrice = true
      })
      .addCase(fetchMusdcPrice.fulfilled, (state, action) => {
        state.isLoadingMusdcPrice = false
        state.musdcPrice = action.payload
      })
      .addCase(fetchMusdcPrice.rejected, (state) => {
        state.isLoadingMusdcPrice = false
        console.error('Error fetching mUSDC price')
      })
      .addCase(fetchDaiePrice.pending, (state) => {
        state.isLoadingDaiePrice = true
      })
      .addCase(fetchDaiePrice.fulfilled, (state, action) => {
        state.isLoadingDaiePrice = false
        state.daiePrice = action.payload
      })
      .addCase(fetchDaiePrice.rejected, (state) => {
        state.isLoadingDaiePrice = false
        console.error('Error fetching dai.e price')
      })
  },
})

export const { resetPriceData } = priceDataSlice.actions

export default priceDataSlice.reducer
