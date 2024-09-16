import { makeExecutableSchema } from '@graphql-tools/schema'
import { loadFilesSync } from '@graphql-tools/load-files'
import { mergeTypeDefs } from '@graphql-tools/merge'

import { resolvers } from './resolvers'

const typesArray = loadFilesSync(`${__dirname}/types/**/*.graphql`)
const queriesArray = loadFilesSync(`${__dirname}/queries/**/*.graphql`)

// Define a union type for BridgeEvent
const additionalTypes = `
  union BridgeEvent = BridgeRequestEvent | BridgeRelayedEvent | BridgeProofProvidedEvent | BridgeDepositRefundedEvent | BridgeDepositClaimedEvent

  type Query {
    events: [BridgeEvent!]!
  }
`

const typeDefs = mergeTypeDefs([...typesArray, ...queriesArray, additionalTypes])

export const schema = makeExecutableSchema({
  typeDefs,
  resolvers,
})