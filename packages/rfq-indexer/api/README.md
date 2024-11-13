# RFQ Indexer API

This API provides access to the indexed bridge event data.

To make requests, use: https://api.synapseprotocol.com
Swagger docs can be found [here](https://api.synapseprotocol.com/api-docs)

## API Calls
All API calls can be viewed in Swagger:

[Swagger Documentation](http://localhost:3001/api-docs)

1. GET /api/pending-transactions-missing-relay
   - Description: Retrieves pending transactions that are missing relay events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions/missing-relay
     ```

2. GET /api/pending-transactions-missing-proof
   - Description: Retrieves pending transactions that are missing proof events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions/missing-proof
     ```

3. GET /api/pending-transactions-missing-claim
   - Description: Retrieves pending transactions that are missing claim events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions/missing-claim
     ```

4. GraphQL endpoint: /graphql
   - Description: Provides a GraphQL interface for querying indexed data, the user is surfaced an interface to query the data via GraphiQL

## Env Vars

- **NODE_ENV**: Set to `"development"` for localhost testing.
- **DATABASE_URL**: PostgreSQL connection URL for the ponder index.

## Important Scripts

- `yarn dev:local`: Runs the API in development mode using local environment variables
- `yarn dev:prod`: Runs the API in development mode using production environment variables
- `yarn start`: Starts the API in production mode

