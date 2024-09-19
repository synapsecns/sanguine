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
    const pendingTransactions =
      await resolvers.Query.pendingTransactionsMissingRelay()
    res.json(pendingTransactions)
  } catch (error) {
    console.error('Error fetching pending transactions missing relay:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/pending-transactions-missing-proof', async (req, res) => {
  try {
    const pendingTransactionsMissingProof =
      await resolvers.Query.pendingTransactionsMissingProof()
    res.json(pendingTransactionsMissingProof)
  } catch (error) {
    console.error('Error fetching pending transactions missing proof:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/pending-transactions-missing-claim', async (req, res) => {
  try {
    const pendingTransactionsMissingClaim =
      await resolvers.Query.pendingTransactionsMissingClaim()
    res.json(pendingTransactionsMissingClaim)
  } catch (error) {
    console.error('Error fetching pending transactions missing claim:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/recent-invalid-relays', async (req, res) => {
  try {
    const queryResult = await resolvers.Query.recentInvalidRelays()
    res.json(queryResult)
  } catch (error) {
    console.error('Error fetching recent invalid relays:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/conflicting-proofs', async (req, res) => {
  try {
    const conflictingProofs = await resolvers.Query.conflictingProofs()
    res.json(conflictingProofs)
  } catch (error) {
    console.error('Error fetching conflicting proofs:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/transaction/:transactionId', async (req, res) => {
  try {
    const transactionId = req.params.transactionId
    const transaction = await resolvers.Query.transactionById(null, {
      transactionId,
    })
    res.json(transaction)
  } catch (error) {
    console.error('Error fetching transaction by ID:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.get('/api/refunded-and-relayed-transactions', async (req, res) => {
  try {
    const transactions = await resolvers.Query.refundedAndRelayedTransactions()
    res.json(transactions)
  } catch (error) {
    console.error('Error fetching refunded and relayed transactions:', error)
    res.status(500).json({ error: 'Internal server error' })
  }
})

app.listen(process.env.PORT, () => {
  console.info('API server runs on http://localhost:3001')
  console.info('REST requests go through http://localhost:3001/api')
  console.info('GraphQL requests go through http://localhost:3001/graphql')
})
