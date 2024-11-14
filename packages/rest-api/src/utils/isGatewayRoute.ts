import { createProxyMiddleware } from 'http-proxy-middleware'

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
  target: 'https://rfq-api.omnirpc.io',
  changeOrigin: true,
})

export const rfqIndexerProxy = createProxyMiddleware({
  target: 'https://rfq-indexer.synapseprotocol.com/api',
  changeOrigin: true,
})
