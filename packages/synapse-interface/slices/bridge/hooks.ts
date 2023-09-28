import { createAsyncThunk } from '@reduxjs/toolkit'

import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import {
  fetchBridgeQuote,
  fetchBridgeQuotes,
  BridgeQuoteRequest,
} from '@/utils/actions/fetchBridgeQuotes'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

export const fetchAndStoreBridgeQuote = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuote',
  async (request: BridgeQuoteRequest, synapseSDK: any) => {
    const bridgeQuote = await fetchBridgeQuote(request, synapseSDK)
    return bridgeQuote
  }
)

export const fetchAndStoreBridgeQuotes = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuotes',
  async (requests: BridgeQuoteRequest[], synapseSDK: any) => {
    const bridgeQuotes = await fetchBridgeQuotes(requests, synapseSDK)
    return bridgeQuotes
  }
)
