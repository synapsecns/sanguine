import {
  FastBridgeQuote,
  FastBridgeQuoteAPI,
  unmarshallFastBridgeQuote,
} from './quote'

const API_URL = 'https://rfq-api-stage.omnirpc.io'
const API_TIMEOUT = 2000

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
