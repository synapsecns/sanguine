import express from 'express'
import swaggerUi from 'swagger-ui-express'

import { specs } from './swagger'
import routes from './routes'

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())

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

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
