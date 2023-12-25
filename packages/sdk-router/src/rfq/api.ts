import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

// TODO: change to mainnet API URL once it's enabled
const API_URL = 'https://rfq-api-testnet.omnirpc.io'

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
