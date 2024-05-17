// URL for RFQ quotes
const RFQ_URL = 'https://rfq-api.omnirpc.io/quotes'

const fetchRfqData = async () => {
  try {
    const response = await fetch(RFQ_URL)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    return await response.json()
  } catch (error) {
    console.error('Failed to fetch RFQ data:', error)
    return []
  }
}

module.exports = { fetchRfqData }
