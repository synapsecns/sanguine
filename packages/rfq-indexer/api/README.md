# RFQ Indexer API

This API provides access to the indexed bridge event data.

## API Calls

1. GET /api/hello
   - Description: A simple hello world endpoint
   - Example: `curl http://localhost:3001/api/hello`

2. GET /api/pending-transactions-missing-relay
   - Description: Retrieves pending transactions that are missing relay events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions-missing-relay
     ```

3. GET /api/pending-transactions-missing-proof
   - Description: Retrieves pending transactions that are missing proof events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions-missing-proof
     ```

4. GET /api/pending-transactions-missing-claim
   - Description: Retrieves pending transactions that are missing claim events
   - Example:
     ```
     curl http://localhost:3001/api/pending-transactions-missing-claim
     ```

5. GraphQL endpoint: /graphql
   - Description: Provides a GraphQL interface for querying indexed data, the user is surfaced an interface to query the data via GraphiQL

## Important Scripts

- `yarn dev:local`: Runs the API in development mode using local environment variables
- `yarn dev:prod`: Runs the API in development mode using production environment variables
- `yarn start`: Starts the API in production mode

