import { getWithTimeout } from '../utils/api'
import { logger } from '../utils/logger'
import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

const API_URL = process.env.RFQ_API_URL || 'https://rfq-api.omnirpc.io'
const API_TIMEOUT = 2000

/**
 * Hits Quoter API /quotes endpoint to get all quotes.
 *
 * @returns A promise that resolves to the list of quotes.
 * Will return an empty list if the request fails or times out.
 */
export const getAllQuotes = async (): Promise<FastBridgeQuote[]> => {
  try {
    const response = await getWithTimeout(
      'RFQ API',
      `${API_URL}/quotes`,
      API_TIMEOUT
    )
    if (!response) {
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    return quotes
      .map((quote) => {
        try {
          return unmarshallFastBridgeQuote(quote)
        } catch (error) {
          logger.error({ quote, error }, 'Could not unmarshall quote')
          return null
        }
      })
      .filter((quote): quote is FastBridgeQuote => quote !== null)
  } catch (error) {
    logger.error({ error }, 'Failed to fetch all quotes')
    return []
  }
}
