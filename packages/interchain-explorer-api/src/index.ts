import express from 'express'
import { createYoga } from 'graphql-yoga'
import { schema } from './schema'
import { overrideJsonBigIntSerialization } from '@/utils/overrideJsonBigIntSerialization'

overrideJsonBigIntSerialization()

const app = express()

const yoga = createYoga({ schema })

app.use(yoga.graphqlEndpoint, yoga)

app.get('/api/hello', (req, res) => {
  res.json({ message: 'Hello World!' })
})

app.listen(process.env.PORT, () => {
  console.info('API server runs on http://localhost:3001')
  console.info('REST requests go through http://localhost:3001/api')
  console.info('GraphQL requests go through http://localhost:3001/graphql')
})
