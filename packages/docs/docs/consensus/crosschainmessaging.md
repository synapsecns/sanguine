---
sidebar_position: 1
---

# Ecosystem

The Synapse Messaging System is built from the ground up to establish trust in each layer.

## Summit Contracts

Synapse Messaging is governed by the [Summit.sol](#) smart contract on the Synapse Chain, which stores remote chain states from bonded off-chain agents, and distributes tips among the off-chain agents.

```
Summit
|- Bonding Manager
|- Inbox
|- Gas Oracle
```

### Supporting contracts

- **BondingManager.sol**:
  Off-chain actors post escrowed funds to the [BondingManager.sol](#) smart contract to interact with the Summit Inbox, allowing them to submit snapshots, receipts, and disputes.
- **Inbox.sol**:
  Off-chain agents' snapshots and receipts are sent to the Summit via [Inbox.sol](#), which also accepts and adjudicates disputes between guards and notaries.
- **GasOracle.sol**:
  Gas data for local and remote chains is tracked via [GasOracle.sol](#)

## Registered Agents

Agents are off-chain actors that relay messages between client-controlled smart contracts and the Synapse origin and destination contracts.

They register themselves by posting a bond to the BondingManager.sol smart contract, which they risk forfeiting if they are found to have acted dishonestly.

Agents receive tips from the originating contract via the Summit for each successfully executed message.

### Types of Agents

- **Notaries**:
  Agents assigned to specific chains which post attestations used to prove messages.
- **Guards**:
  Guards watch for Notary attestations and may file a dispute in the event of a suspected fraud.
- **Executors**:
  Executors are unbonded agents which execute messages after the security window has passed.

## Setting up a Remote Chain

Remote chains are governed by the LightManager.sol smart contract, which keeps track of its on-chain agents. The Light Manager stays in sync with the Bonding Manager via agent merkle roots from notary-signed attestations, allowing agents to register themselves, and receive tips, and also be subject to reward or penalization following a dispute.

```
Light Manager
|- Origin Contract
   |- Gas Oracle
|- Destination Contract
   |- Light Inbox
```

### Suppporting contracts

- **Origin.sol**:
  [Origin.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/Origin.sol) formats and sends message payloads, inserts message hashes into the Origin Merkle Tree, and enforces and distributes tip values.
- **Destination.sol**:
  [Destination.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/Destination.sol) receives and executes messages from its local Inbox, passes Agent Merkle roots to its local Light Manager, and tracks gas data submitted by notaries.
- **LightInbox.sol**:
  On-chain agents' snapshots and receipts are sent to the local LightManager via [LightInbox.sol](#), which also passes notary attestations to the Synapse Destination contract and accepts and initaites disputes between guards and notaries.
  Inbox for the local LightManager contract.
- **GasOracle.sol**:
  Tracks gas data

## New Chains and Agents

When a chain is first set up, its destination contract is initialized through permissioned mechanism using the latest Agent Root from the Bonding Manager contract.

After initialization, new agents can then add themselves permissionlessly via the Bonding Manager on the Synapse chain, which sends an updated Agent Root to the remote chains.

Once the Agent Root is updated, all bonded agents must submit a proof of inclusion to be added to or remain in the Agent Set for their chain.

## Client Apps

To use the Synapse Messaging System, clients must deploy smart contracts for sending and receiving messages. The sending and receiving functions may be combined in a single contract, as seen in the examples here:

1.  [TestClient.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/TestClient.sol)
2.  [PingPong.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/PingPongClient.sol)
