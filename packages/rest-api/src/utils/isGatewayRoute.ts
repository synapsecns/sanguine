import { createProxyMiddleware } from 'http-proxy-middleware'

// Environment variables for RFQ API and Indexer URLs
const RFQ_API_URL = process.env.RFQ_API_URL || 'https://rfq-api.omnirpc.io'
const RFQ_INDEXER_URL =
  process.env.RFQ_INDEXER_URL || 'https://rfq-indexer.synapseprotocol.com/api'

export const isRFQIndexerRequest = (route: string): boolean => {
  return (
    route.includes('/conflicting-proofs') ||
    route.includes('/disputes') ||
    route.includes('/invalid-relaus') ||
    route.includes('/pending-transactions') ||
    route.includes('/refunded-and-relayed') ||
    route.includes('/transaction-id')
  )
}

export const isRFQAPIRequest = (route: string): boolean => {
  return (
    route.includes('/ack') ||
    route.includes('/bulk_quotes') ||
    route.includes('/contracts') ||
    route.includes('/open_quote_requests') ||
    route.includes('/quotes') ||
    route.includes('/rfq') ||
    route.includes('/rfq_stream')
  )
}

export const rfqApiProxy = createProxyMiddleware({
  target: RFQ_API_URL,
  changeOrigin: true,
})

export const rfqIndexerProxy = createProxyMiddleware({
  target: RFQ_INDEXER_URL,
  changeOrigin: true,
})
