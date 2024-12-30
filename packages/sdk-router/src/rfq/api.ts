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
    origin_sender?: string
    dest_recipient?: string
    zap_data?: string
    zap_native?: string
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

export type QuoteRequestOptions = {
  originSender?: string
  destRecipient?: string
  zapData?: string
  zapNative?: BigNumber
}

const ZeroQuote: RelayerQuote = {
  destAmount: BigNumber.from(0),
}

export const fetchWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any,
  init?: RequestInit
): Promise<Response | null> => {
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)
  try {
    const response = await fetch(url, {
      signal: controller.signal,
      ...init,
    })
    clearTimeout(timeoutId)
    if (!response.ok) {
      console.info({ name, url, params, response }, `${name}: not OK`)
      return null
    }
    return response
  } catch (error) {
    console.info({ name, url, timeout, params, error }, `${name}: timed out`)
    return null
  }
}

export const postWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any
): Promise<Response | null> => {
  try {
    const response = await fetchWithTimeout(name, url, timeout, {
      method: 'POST',
      body: JSON.stringify(params),
      headers: {
        'Content-Type': 'application/json',
      },
    })
    if (!response.ok) {
      const text = await response.text()
      console.info(
        { url, timeout, params, response, text },
        `${name}: response was not OK`
      )
      return null
    }
    return response
  } catch (error) {
    console.info(
      { url, timeout, params, error },
      `${name}: was not able to get response in time`
    )
    return null
  }
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
      const text = await response.text()
      console.info(
        { response, text },
        'Response was not OK for getting all quotes'
      )
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    return quotes
      .map((quote) => {
        try {
          return unmarshallFastBridgeQuote(quote)
        } catch (error) {
          console.error({ quote, error }, 'Could not unmarshall quote')
          return null
        }
      })
      .filter((quote): quote is FastBridgeQuote => quote !== null)
  } catch (error) {
    console.info({ error }, 'Was not able to get all quotes in time')
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
  options: QuoteRequestOptions = {}
): Promise<RelayerQuote> => {
  const rfqRequest: PutRFQRequestAPI = {
    quote_types: ['active', 'passive'],
    data: {
      origin_chain_id: ticker.originToken.chainId,
      dest_chain_id: ticker.destToken.chainId,
      origin_token_addr: ticker.originToken.token,
      dest_token_addr: ticker.destToken.token,
      origin_amount_exact: originAmount.toString(),
      expiration_window: EXPIRATION_WINDOW,
      origin_sender: options.originSender,
      dest_recipient: options.destRecipient,
      zap_data: options.zapData ?? '0x',
      zap_native: options.zapNative?.toString() ?? '0',
    },
  }
  try {
    const response = await fetchWithTimeout(`${API_URL}/rfq`, API_TIMEOUT, {
      method: 'PUT',
      body: JSON.stringify(rfqRequest),
      headers: {
        'Content-Type': 'application/json',
      },
    })
    if (!response.ok) {
      const text = await response.text()
      console.info(
        {
          rfqRequest,
          response,
          text,
        },
        'Response was not OK for RFQ quote'
      )
      return ZeroQuote
    }
    // Check that response is successful, contains non-zero dest amount, and has a relayer address
    const rfqResponse: PutRFQResponseAPI = await response.json()
    if (!rfqResponse.success) {
      console.info(
        {
          rfqRequest,
          rfqResponse,
          reason: rfqResponse.reason ?? 'Unknown reason',
        },
        'No RFQ quote returned'
      )
      return ZeroQuote
    }
    if (!rfqResponse.dest_amount || !rfqResponse.relayer_address) {
      console.error(
        { rfqRequest, rfqResponse },
        'Error getting RFQ quote: missing dest_amount or relayer_address in response'
      )
      return ZeroQuote
    }
    const destAmount = BigNumber.from(rfqResponse.dest_amount)
    if (destAmount.lte(0)) {
      console.info({ rfqRequest, rfqResponse }, 'No RFQ quote returned')
      return ZeroQuote
    }
    return {
      destAmount,
      relayerAddress: rfqResponse.relayer_address,
      quoteID: rfqResponse.quote_id,
    }
  } catch (error) {
    console.info({ rfqRequest, error }, 'Was not able to get RFQ quote in time')
    return ZeroQuote
  }
}
