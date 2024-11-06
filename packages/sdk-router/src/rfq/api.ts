import { BigNumber } from '@ethersproject/bignumber'

import { Ticker } from './ticker'
import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

const API_URL = 'https://rfq-api.omnirpc.io'
const API_TIMEOUT = 2000

export type PutRFQRequestAPI = {
  user_address: string
  // TODO: make integrator_id required
  integrator_id?: string
  quote_types: string[]
  data: {
    origin_chain_id: number
    dest_chain_id: number
    origin_token_addr: string
    dest_token_addr: string
    origin_amount: string
    expiration_window: number
  }
}

export type PutRFQResponseAPI = {
  success: boolean
  reason?: string
  quote_type?: string
  quote_id?: string
  dest_amount?: string
  relayer_address?: string
}

export type Quote = {
  destAmount: BigNumber
  relayerAddress: string
  quoteID?: string
}

const fetchWithTimeout = async (
  url: string,
  timeout: number
): Promise<Response> => {
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)
  return fetch(url, { signal: controller.signal }).finally(() =>
    clearTimeout(timeoutId)
  )
}

/**
 * Hits Quoter API /quotes endpoint to get all quotes.
 *
 * @returns A promise that resolves to the list of quotes.
 * Will return an empty list if the request fails or times out.
 */
export const getAllQuotes = async (): Promise<FastBridgeQuote[]> => {
  try {
    const response = await fetchWithTimeout(`${API_URL}/quotes`, API_TIMEOUT)
    if (!response.ok) {
      console.error('Error fetching quotes:', response.statusText)
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    return quotes
      .map((quote) => {
        try {
          return unmarshallFastBridgeQuote(quote)
        } catch (error) {
          console.error('Error unmarshalling quote:', error)
          return null
        }
      })
      .filter((quote): quote is FastBridgeQuote => quote !== null)
  } catch (error) {
    console.error('Error fetching quotes:', error)
    return []
  }
}

/**
 * Hits Quoter API /rfq endpoint to get the best quote for a given ticker and origin amount.
 *
 * @returns A promise that resolves to the best quote.
 * Will return null if the request fails or times out.
 */
export const getBestRFQQuote = async (
  ticker: Ticker,
  originAmount: BigNumber,
  originUserAddress?: string
): Promise<Quote | null> => {}
