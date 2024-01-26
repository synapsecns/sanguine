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
    const url = `https://stip-api.omnirpc.io/fee-rebate-bps`

    // Configurable parameters
    const maxRetries = 5 // maximum number of retries
    const initialDelay = 1000 // initial delay in milliseconds

    // Helper function to delay for a given amount of time
    const delay = (duration) =>
      new Promise((resolve) => setTimeout(resolve, duration))

    for (let attempt = 0; attempt < maxRetries; attempt++) {
      try {
        const response = await fetch(url, { method: 'GET' })

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }

        const data = await response.json()
        return data
      } catch (error) {
        console.error(
          `Attempt ${attempt + 1}: Error fetching fee and rebate data`,
          error
        )

        // If it's the last attempt, reject the promise
        if (attempt === maxRetries - 1) {
          return rejectWithValue(
            error.message || 'Failed to fetch fee and rebate data'
          )
        }

        // Exponential backoff
        const delayDuration = initialDelay * 2 ** attempt
        console.log(`Retrying in ${delayDuration}ms...`)
        await delay(delayDuration)
      }
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
