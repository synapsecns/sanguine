import { BigNumber } from '@ethersproject/bignumber'

import { Ticker } from './ticker'
import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'
import { logger } from '../utils/logger'

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
  url: string,
  timeout: number,
  init?: RequestInit
): Promise<Response> => {
  const controller = new AbortController()
  const reason = `Timeout of ${timeout}ms exceeded for ${url}`
  const timeoutId = setTimeout(() => (controller.abort as any)(reason), timeout)
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
      logger.error({ response }, 'Error fetching quotes')
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    logger.info(
      { quotes },
      `Received ${quotes.length} quotes from ${API_URL}/quotes`
    )
    return quotes
      .map((quote) => {
        try {
          return unmarshallFastBridgeQuote(quote)
        } catch (error) {
          logger.error({ quote, error }, 'Error unmarshalling quote')
          return null
        }
      })
      .filter((quote): quote is FastBridgeQuote => quote !== null)
  } catch (error) {
    logger.error({ error }, 'Error fetching quotes')
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
  try {
    const rfqRequest: PutRFQRequestAPI = {
      // TODO: add active quotes once they are fixed
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
        // TODO: cleanup
        zap_data: options.zapData ?? '0x',
        zap_native: options.zapNative?.toString() ?? '0',
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
      logger.error({ rfqRequest, response }, 'Error fetching quote')
      return ZeroQuote
    }
    // Check that response is successful, contains non-zero dest amount, and has a relayer address
    const rfqResponse: PutRFQResponseAPI = await response.json()
    logger.info(
      { rfqRequest, rfqResponse },
      `Received quote from ${API_URL}/rfq`
    )
    if (!rfqResponse.success) {
      logger.error(
        {
          reason: rfqResponse.reason ?? 'Unknown reason',
        },
        'No RFQ quote returned'
      )
      return ZeroQuote
    }
    if (!rfqResponse.dest_amount || !rfqResponse.relayer_address) {
      logger.error(
        { rfqResponse },
        'Error fetching quote: missing dest_amount or relayer_address in response:'
      )
      return ZeroQuote
    }
    const destAmount = BigNumber.from(rfqResponse.dest_amount)
    if (destAmount.lte(0)) {
      logger.error({ rfqResponse }, 'No RFQ quote returned')
      return ZeroQuote
    }
    return {
      destAmount,
      relayerAddress: rfqResponse.relayer_address,
      quoteID: rfqResponse.quote_id,
    }
  } catch (error) {
    logger.error({ error }, 'Error fetching quote')
    return ZeroQuote
  }
}
