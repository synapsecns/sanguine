import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

const API_URL = 'https://rfq-api.omnirpc.io'
const API_TIMEOUT = 2000

const fetchWithTimeout = async (
  url: string,
  timeout: number
): Promise<Response> => {
  const controller = new AbortController()
  setTimeout(() => controller.abort(), timeout)
  return fetch(url, { signal: controller.signal })
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
      return []
    }
    // The response is a list of quotes in the FastBridgeQuoteAPI format
    const quotes: FastBridgeQuoteAPI[] = await response.json()
    return quotes.map(unmarshallFastBridgeQuote)
  } catch (error) {
    return []
  }
}
