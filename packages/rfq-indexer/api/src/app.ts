import express from 'express'
import swaggerUi from 'swagger-ui-express'
import { createYoga } from 'graphql-yoga'

import { specs } from './swagger'
import routes from './routes'
import { schema } from './graphql/schema'
import { overrideJsonBigIntSerialization } from './utils/overrideJsonBigIntSerialization'
import fs from 'fs'
import path from 'path'

const app = express()
const port = process.env.PORT || 3001

overrideJsonBigIntSerialization()

app.use(express.json())

// Swagger UI setup
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs))

// REST API routes
app.use('/api', routes)

// GraphQL setup
const yoga = createYoga({ schema })
app.use('/graphql', yoga)

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
  console.info('API server runs on http://localhost:3001')
  console.info('REST requests go through http://localhost:3001/api')
  console.info('GraphQL requests go through http://localhost:3001/graphql')
})

// export swagger json to a file dynamically
const outputPath = path.join(__dirname, 'swagger-output.json')
fs.writeFileSync(outputPath, JSON.stringify(specs, null, 2), 'utf-8')
