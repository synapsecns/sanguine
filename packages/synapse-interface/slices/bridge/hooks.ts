import { createAsyncThunk } from '@reduxjs/toolkit'

import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import {
  fetchBridgeQuote,
  BridgeQuoteRequest,
} from '@/utils/actions/fetchBridgeQuotes'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

export const fetchAndStoreBridgeQuotes = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuotes',
  async () => {}
)
