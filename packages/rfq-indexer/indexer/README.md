# RFQ Indexer

This indexer captures and stores FastBridgeV2 events from various blockchain networks.

## Important Scripts

- `yarn dev:local`: Runs the indexer in development mode, clearing previous data
- `yarn dev`: Runs the indexer in development mode
- `yarn start`: Starts the indexer in production mode

To run these scripts, use `yarn <script-name>` in the terminal from the indexer directory.

## Main Files for Contributors

1. ponder.schema.ts
   - Description: Defines the database schema for indexed events
2. ponder.config.ts
   - Description: Configures the indexer, including network details and contract addresses
3. src/index.ts
   - Description: Contains the main indexing logic for different event types

4. abis/FastBridgeV2.ts
   - Description: Contains the ABI (Application Binary Interface) for the FastBridgeV2 contract

When contributing, focus on these files for making changes to the indexing logic, adding new event types, or modifying the database schema.
