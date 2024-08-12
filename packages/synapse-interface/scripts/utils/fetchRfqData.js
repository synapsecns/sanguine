// URL for RFQ quotes
const RFQ_URL = 'https://rfq-api.omnirpc.io/quotes'

const fetchRfqData = async () => {
  try {
    const response = await fetch(RFQ_URL)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    quotes = await response.json()
    // Filter out quotes older than a day ago
    const updatedAtThreshold = Date.now() - 24 * 60 * 60 * 1000
    return quotes.filter(
      (quote) => new Date(quote.updated_at) > updatedAtThreshold
    )
  } catch (error) {
    console.error('Failed to fetch RFQ data:', error)
    return []
  }
}

module.exports = { fetchRfqData }
