import express from 'express'
import swaggerUi from 'swagger-ui-express'
import { createProxyMiddleware } from 'http-proxy-middleware'

import { specs } from './swagger'
import routes from './routes'
import { logger } from './middleware/logger'
import fs from 'fs'
import path from 'path'

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())

const apiProxyRoutes = [
  {
    route: '/rfq-indexer-api',
    target:
      process.env.RFQ_INDEXER_API_URL ||
      'https://triumphant-magic-production.up.railway.app',
  },
  { route: '/rfq-api', target: 'https://rfq-api.omnirpc.io/' },
]

apiProxyRoutes.forEach(({ route, target }) => {
  app.use(
    route,
    createProxyMiddleware({
      target,
      changeOrigin: true,
      pathRewrite: {
        [`^${route}`]: '',
      },
    })
  )
})

app.use((req, res, next) => {
  logger.info({
    msg: `Incoming request ${req.path}`,
    method: req.method,
    path: req.path,
    query: req.query,
    body: req.method === 'POST' || req.method === 'PUT' ? req.body : undefined,
  })

  const originalPath = req.path

  const originalJson = res.json
  res.json = function (body) {
    logger.info({
      msg: `Outgoing response ${originalPath}`,
      method: req.method,
      path: originalPath,
      statusCode: res.statusCode,
      body:
        originalPath === '/' || originalPath.toLowerCase() === '/tokenlist'
          ? '[truncated for size]'
          : body,
    })
    return originalJson.call(this, body)
  }

  next()
})

app.listen(port, () => {
  logger.info(`Server is listening on port ${port}`)
})

app.use(
  '/api-docs',
  (_req, res, next) => {
    res.set(
      'Cache-Control',
      'no-store, no-cache, must-revalidate, proxy-revalidate'
    )
    res.set('Pragma', 'no-cache')
    res.set('Expires', '0')
    res.set('Surrogate-Control', 'no-store')
    next()
  },
  swaggerUi.serve,
  swaggerUi.setup(specs)
)

app.use('/', routes)

app.use((err, _req, res, _next) => {
  logger.error(`Express error: ${err.message}`, { stack: err.stack })
  res.status(500).json({ error: 'Something went wrong', details: err.message })
})

process.on('uncaughtException', (err) => {
  logger.error(`Uncaught Exception: ${err.message}`, { stack: err.stack })
  process.exit(1)
})

process.on('unhandledRejection', (reason, promise) => {
  logger.error('Unhandled Rejection at:', promise, 'reason:', reason)
})

const outputPath = path.join(__dirname, 'swagger-output.json')
fs.writeFileSync(outputPath, JSON.stringify(specs, null, 2), 'utf-8')
