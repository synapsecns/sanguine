import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

// TODO: fill
const API_URL = 'API_URL'

/**
 * Hits Quoter API /quotes endpoint to get all quotes.
 *
 * @returns A promise that resolves to the list of quotes.
 */
export const getAllQuotes = async (): Promise<FastBridgeQuote[]> => {
  const response = await fetch(API_URL + '/quotes')
  // Return empty list if response is not ok
  if (!response.ok) {
    return []
  }
  // The response is a list of quotes in the FastBridgeQuoteAPI format
  const quotes: FastBridgeQuoteAPI[] = await response.json()
  return quotes.map(unmarshallFastBridgeQuote)
}
