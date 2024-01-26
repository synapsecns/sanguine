import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

export interface FeeAndRebateState {
  toFromFeeAndRebateBps: {}
  isLoading: boolean
}

const initialState: FeeAndRebateState = {
  toFromFeeAndRebateBps: {},
  isLoading: false,
}

export const fetchFeeAndRebate = createAsyncThunk(
  'feeAndRebate/fetchFeeAndRebate',
  async (_, { rejectWithValue }) => {
    /* TODO: UPDATE URL */
    const url = `https://api.url.com/api/v1/fee-rebate-bps`

    try {
      const response = await fetch(url, { method: 'GET' })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()

      return data
    } catch (error) {
      console.error('Error fetching fee and rebate data:', error)

      return rejectWithValue(
        error.message || 'Failed to fetch fee and rebate data'
      )
    }
  }
)

export const feeAndRebateSlice = createSlice({
  name: 'feeAndRebate',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchFeeAndRebate.pending, (state) => {
        state.isLoading = true
      })
      .addCase(fetchFeeAndRebate.fulfilled, (state, action) => {
        state.isLoading = false
        state.toFromFeeAndRebateBps = action.payload
      })
      .addCase(fetchFeeAndRebate.rejected, (state) => {
        state.isLoading = false
      })
  },
})

export default feeAndRebateSlice.reducer
