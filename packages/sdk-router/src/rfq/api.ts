import { BigNumber } from '@ethersproject/bignumber'

import { Ticker } from './ticker'
import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

const API_URL = 'https://rfq-api-stage.omnirpc.io'
const API_TIMEOUT = 2000

/**
 * The expiration window for active quotes in milliseconds to be used by the RFQ API.
 * Relayers will have to respond with a quote within this time window.
 */
const EXPIRATION_WINDOW = 1000

export type PutRFQRequestAPI = {
  user_address?: string
  // TODO: make integrator_id required
  integrator_id?: string
  quote_types: string[]
  data: {
    origin_chain_id: number
    dest_chain_id: number
    origin_token_addr: string
    dest_token_addr: string
    origin_amount_exact: string
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

export type RelayerQuote = {
  destAmount: BigNumber
  relayerAddress?: string
  quoteID?: string
}

const ZeroQuote: RelayerQuote = {
  destAmount: BigNumber.from(0),
}

const fetchWithTimeout = async (
  url: string,
  timeout: number,
  init?: RequestInit
): Promise<Response> => {
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)
  return fetch(url, { signal: controller.signal, ...init }).finally(() =>
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
 * Hits Quoter API /rfq PUT endpoint to get the best quote for a given ticker and origin amount.
 *
 * @returns A promise that resolves to the best quote.
 * Will return a zero quote if the request fails or times out.
 */
export const getBestRelayerQuote = async (
  ticker: Ticker,
  originAmount: BigNumber,
  originUserAddress?: string
): Promise<RelayerQuote> => {
  try {
    const rfqRequest: PutRFQRequestAPI = {
      user_address: originUserAddress,
      // TODO: add active quotes once they are fixed
      quote_types: ['passive'],
      data: {
        origin_chain_id: ticker.originToken.chainId,
        dest_chain_id: ticker.destToken.chainId,
        origin_token_addr: ticker.originToken.token,
        dest_token_addr: ticker.destToken.token,
        origin_amount_exact: originAmount.toString(),
        expiration_window: EXPIRATION_WINDOW,
      },
    }
    const response = await fetchWithTimeout(`${API_URL}/rfq`, API_TIMEOUT, {
      method: 'PUT',
      body: JSON.stringify(rfqRequest),
      headers: {
        'Content-Type': 'application/json',
      },
    })
    if (!response.ok) {
      console.error('Error fetching quote:', response.statusText)
      return ZeroQuote
    }
    // Check that response is successful, contains non-zero dest amount, and has a relayer address
    const rfqResponse: PutRFQResponseAPI = await response.json()
    if (!rfqResponse.success) {
      console.error(
        'No RFQ quote returned:',
        rfqResponse.reason ?? 'Unknown reason'
      )
      return ZeroQuote
    }
    if (!rfqResponse.dest_amount || !rfqResponse.relayer_address) {
      console.error(
        'Error fetching quote: missing dest_amount or relayer_address in response:',
        rfqResponse
      )
      return ZeroQuote
    }
    const destAmount = BigNumber.from(rfqResponse.dest_amount)
    if (destAmount.lte(0)) {
      console.error('No RFQ quote returned')
      return ZeroQuote
    }
    return {
      destAmount,
      relayerAddress: rfqResponse.relayer_address,
      quoteID: rfqResponse.quote_id,
    }
  } catch (error) {
    console.error('Error fetching quote:', error)
    return ZeroQuote
  }
}
