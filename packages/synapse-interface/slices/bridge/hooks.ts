import { createAsyncThunk } from '@reduxjs/toolkit'

import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'
import {
  fetchBridgeQuote,
  fetchBridgeQuotes,
  BridgeQuoteRequest,
  BridgeQuoteResponse,
} from '@/utils/actions/fetchBridgeQuotes'

export const useBridgeState = (): RootState['bridge'] => {
  return useAppSelector((state) => state.bridge)
}

export const fetchAndStoreBridgeQuote = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuote',
  async ({
    request,
    synapseSDK,
  }: {
    request: BridgeQuoteRequest
    synapseSDK: any
  }) => {
    const bridgeQuote: BridgeQuoteResponse = await fetchBridgeQuote(
      request,
      synapseSDK
    )
    return bridgeQuote
  }
)

export const fetchAndStoreBridgeQuotes = createAsyncThunk(
  'bridge/fetchAndStoreBridgeQuotes',
  async ({
    requests,
    synapseSDK,
  }: {
    requests: BridgeQuoteRequest[]
    synapseSDK: any
  }) => {
    const bridgeQuotes: BridgeQuoteResponse[] = await fetchBridgeQuotes(
      requests,
      synapseSDK
    )
    return bridgeQuotes
  }
)
