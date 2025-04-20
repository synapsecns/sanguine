import { getWithTimeout, logger } from '../utils'
import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'
import { marshallTicker } from './ticker'

const DEFAULT_API_URL = 'https://rfq-api.omnirpc.io'
const CUSTOM_API_URL = process.env.RFQ_API_URL
const API_TIMEOUT = 2000

/**
 * Get all API endpoints to query.
 * Uses both the default API endpoint and custom API endpoint if defined and different.
 *
 * @returns Array of API endpoints to query
 */
const getApiEndpoints = (): string[] => {
  return CUSTOM_API_URL && CUSTOM_API_URL !== DEFAULT_API_URL
    ? [DEFAULT_API_URL, CUSTOM_API_URL]
    : [DEFAULT_API_URL]
}

/**
 * Fetch quotes from a specific API endpoint
 *
 * @param endpoint The API endpoint URL to query
 * @returns Array of unmarshalled FastBridgeQuote objects or empty array if request fails
 */
const getQuotesFromEndpoint = async (
  endpoint: string
): Promise<FastBridgeQuote[]> => {
  try {
    const response = await getWithTimeout(
      'RFQ API',
      `${endpoint}/quotes`,
      API_TIMEOUT
    )
    if (!response) {
      logger.info({ endpoint }, 'No response from endpoint')
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    return quotes
      .map((quote) => {
        try {
          return unmarshallFastBridgeQuote(quote)
        } catch (error) {
          logger.error({ endpoint, quote, error }, 'Could not unmarshall quote')
          return null
        }
      })
      .filter((quote): quote is FastBridgeQuote => quote !== null)
  } catch (error) {
    logger.error({ endpoint, error }, 'Failed to fetch quotes from endpoint')
    return []
  }
}

/**
 * Create a unique key for a quote based on relayer address and ticker
 *
 * @param quote The quote to create a key for
 * @returns A unique string key
 */
const createQuoteKey = (quote: FastBridgeQuote): string => {
  return `${quote.relayerAddr}-${marshallTicker(quote.ticker)}`
}

/**
 * Merges quotes from multiple arrays, avoiding duplicates based on relayer + ticker.
 *
 * @param quotesArrays Arrays of quotes to merge
 * @returns Merged array of unique quotes
 */
const mergeQuotes = (quotesArrays: FastBridgeQuote[][]): FastBridgeQuote[] => {
  // Use a Map to track unique quotes by a string key representing the relayer + ticker key.
  const uniqueQuotes = new Map<string, FastBridgeQuote>()

  quotesArrays.forEach((quotes) => {
    quotes.forEach((quote) => {
      const key = createQuoteKey(quote)
      // If we haven't seen this ticker by this relayer before, or if this quote is newer, keep it
      const existingQuote = uniqueQuotes.get(key)
      if (!existingQuote || quote.updatedAt > existingQuote.updatedAt) {
        uniqueQuotes.set(key, quote)
      }
    })
  })

  return Array.from(uniqueQuotes.values())
}

/**
 * Hits all available Quoter API /quotes endpoints to get quotes.
 * Merges quotes from all endpoints, preferring the most recently updated quotes.
 *
 * @returns A promise that resolves to the list of quotes.
 * Will return an empty list if all requests fail or time out.
 */
export const getAllQuotes = async (): Promise<FastBridgeQuote[]> => {
  try {
    const endpoints = getApiEndpoints()
    // Fetch quotes from all endpoints in parallel
    const quotesArrays = await Promise.all(
      endpoints.map((endpoint) => getQuotesFromEndpoint(endpoint))
    )
    // Merge all quotes, keeping the most recent ones
    return mergeQuotes(quotesArrays)
  } catch (error) {
    logger.error({ error }, 'Failed to fetch all quotes')
    return []
  }
}
