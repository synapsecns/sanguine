const { prettyPrintTS } = require('./utils/prettyPrintTs')

const RFQ_URL = 'https://rfq-api.omnirpc.io/quotes'

const fetchRFQData = async (url) => {
  try {
    const response = await fetch(url)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    return await response.json()
  } catch (error) {
    console.error('Failed to fetch RFQ data:', error)
    return []
  }
}

const convertDataToMap = (data) => {
  const result = {}

  data.forEach((item) => {
    const originKey = `${item.origin_token_addr.toLowerCase()}-${
      item.origin_chain_id
    }`
    const destinationValue = `${item.dest_token_addr.toLowerCase()}-${
      item.dest_chain_id
    }`

    if (!result[originKey]) {
      result[originKey] = []
    }

    if (!result[originKey].includes(destinationValue)) {
      result[originKey].push(destinationValue)
    }
  })

  return result
}

const printMaps = async () => {
  const rfqData = await fetchRFQData(RFQ_URL)

  const data = convertDataToMap(rfqData)

  prettyPrintTS(data, 'RFQ_MAP', './constants/rfqMap.ts')
}

printMaps()
