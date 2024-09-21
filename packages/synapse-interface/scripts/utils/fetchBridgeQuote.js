// URL for Synapse bridge quotes
const BRIDGE_URL = 'https://api.synapseprotocol.com/bridge'

// Function to fetch bridge quote
const fetchBridgeQuote = async (
  fromChain,
  toChain,
  fromToken,
  toToken,
  amount
) => {
  try {
    const url = new URL(BRIDGE_URL)

    // Add query parameters
    url.searchParams.append('fromChain', fromChain)
    url.searchParams.append('toChain', toChain)
    url.searchParams.append('fromToken', fromToken)
    url.searchParams.append('toToken', toToken)
    url.searchParams.append('amount', amount)

    const response = await fetch(url)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const quotes = await response.json()
    return quotes
  } catch (error) {
    console.error('Failed to fetch bridge quote:', error)
    return []
  }
}

module.exports = { fetchBridgeQuote }
