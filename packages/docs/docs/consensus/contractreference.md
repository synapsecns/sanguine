---
sidebar_position: 3
---

# Smart Contract Reference

## Synapse Chain

- **Summit.sol**: Stores remote chain states and distributes tips to off-chain actors
- **BondingManager.sol**: Current Synapse Chain agents
- **Inbox.sol**: Receives snapshots and receipts, and adjudicates disputes
- **GasOracle.sol**: Tracks gas data for local and remote chains

## Remote Chain

- **LightManager.sol**: Tracks agents on chains other than Synapse Chain
- **Origin.sol**: Origin smart contract
- **Destination.sol**: Destination smart contract
- **LightInbox.sol**: Communicates with LightManager and Destination contracts

## Client contracts

- **TestClient.sol**: Test client (combines sender and receiver)
- **PingPing.sol**: Test message (combines sender and receiver)
- **IMessageRecipient.sol**: Interface to receive base messages
