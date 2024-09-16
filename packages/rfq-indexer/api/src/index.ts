import express from 'express'
import { createYoga } from 'graphql-yoga'
import { schema } from './graphql/schema'
import { overrideJsonBigIntSerialization } from './utils/overrideJsonBigIntSerialization'
import { resolvers } from './graphql/resolvers'

overrideJsonBigIntSerialization()

const app = express()

const yoga = createYoga({ schema })

app.use(yoga.graphqlEndpoint, yoga)

app.get('/api/hello', (req, res) => {
    res.json({ message: 'Hello World!' })
  })  

app.get('/api/pending-transactions-missing-relay', async (req, res) => {
  try {
    const pendingTransactions = await resolvers.Query.pendingTransactionsMissingRelay();
    res.json(pendingTransactions);
  } catch (error) {
    console.error('Error fetching pending transactions:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

app.get('/api/pending-transactions-missing-proof', async (req, res) => {
  try {
    const pendingTransactionsMissingProof = await resolvers.Query.pendingTransactionsMissingProof();
    res.json(pendingTransactionsMissingProof);
  } catch (error) {
    console.error('Error fetching pending transactions missing proof:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

app.get('/api/pending-transactions-missing-claim', async (req, res) => {
  try {
    const pendingTransactionsMissingClaim = await resolvers.Query.pendingTransactionsMissingClaim();
    res.json(pendingTransactionsMissingClaim);
  } catch (error) {
    console.error('Error fetching pending transactions missing claim:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});


app.listen(process.env.PORT, () => {
  console.info('API server runs on http://localhost:3001')
  console.info('REST requests go through http://localhost:3001/api')
  console.info('GraphQL requests go through http://localhost:3001/graphql')
})