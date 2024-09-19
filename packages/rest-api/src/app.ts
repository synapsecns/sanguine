import express from 'express'
import swaggerUi from 'swagger-ui-express'

import { specs } from './swagger'

import routes from './routes'

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs))

app.use('/', routes)

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
