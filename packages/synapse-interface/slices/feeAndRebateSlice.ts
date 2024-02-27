import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { Address } from 'viem'

import { ARB, parseTokenValue } from '@/components/Activity/AirdropRewards'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'

/** ARB STIP Rewarder */
export const Rewarder = {
  address: '0x48fa1ebda1af925898c826c566f5bb015e125ead' as Address,
  startBlock: 174234366n, // Start of STIP Rewards on Arbitrum
}

export interface FeeAndRebateState {
  toFromFeeAndRebateBps: {}
  isLoading: boolean
  transactions: any[]
  cumulativeRewards: any
  parsedCumulativeRewards: any
}

const initialState: FeeAndRebateState = {
  toFromFeeAndRebateBps: {},
  isLoading: false,
  transactions: [],
  cumulativeRewards: null,
  parsedCumulativeRewards: null,
}

const calculateTotalTransferValue = (data: any[]): bigint => {
  let total: bigint = 0n
  for (const item of data) {
    if (item.transferValue) {
      total += item.transferValue
    }
  }
  return total
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

export const fetchArbStipRewards = createAsyncThunk(
  'feeAndRebate/fetchArbStipRewards',
  async (connectedAddress: string, { rejectWithValue }) => {
    try {
      const { logs, data } = await getErc20TokenTransfers(
        ARB.tokenAddress,
        Rewarder.address,
        connectedAddress as Address,
        ARB.network,
        Rewarder.startBlock
      )

      const cumulativeRewards = calculateTotalTransferValue(data)

      return {
        logs: logs ?? [],
        transactions: data,
        cumulativeRewards,
      }
    } catch (error) {
      console.error('Error fetching ARB Stip Rewards:', error)
      return rejectWithValue(
        error.message || 'Failed to fetch ARB Stip Rewards'
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
      .addCase(fetchArbStipRewards.pending, (state) => {
        state.isLoading = true
      })
      .addCase(fetchArbStipRewards.fulfilled, (state, action) => {
        state.isLoading = false
        state.transactions = action.payload.transactions

        state.cumulativeRewards = action.payload.cumulativeRewards
        state.parsedCumulativeRewards = parseTokenValue(
          action.payload.cumulativeRewards,
          ARB.decimals
        )
      })
      .addCase(fetchArbStipRewards.rejected, (state) => {
        state.isLoading = false
        state.transactions = []
        state.cumulativeRewards = null
        state.parsedCumulativeRewards = null
      })
  },
})

export default feeAndRebateSlice.reducer
