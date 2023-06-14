---
sidebar_position: 5
---

# Glossary

### Accused Agent
When a [fraud report](#fraud-report) is submitted by an [Accusing Guard](#accusing-guard), the report will include the
purported [fraud](#fraud) which is a false claim signed by the "Accused Agent".

### Accused Guard
When a [Guard](#guard) is the [Accused Agent](#accused-agent) in a [fraud report](#fraud-report) submitted by another [Accusing Guard](accusing-guard),
it is referred to as the "Accused Guard".

### Accused Notary
When a [Notary](#notary) is the [Accused Agent](#accused-agent) in a [fraud report](#fraud-report) submitted by an [Accusing Guard](accusing-guard),
it is referred to as the "Accused Notary".

### Accusing Guard
The [Guard](#guard) who submits a [fraud report](#fraud-report) is defined to be the Accusing Guard.

### Agent Root
The Agent Root is the root of the Merkle Tree formed from the [Agent Set](#agent-set). The list of all registered bonded agents
that have posted bond on the [Synapse Chain](#synapse-chain) are put into a Merkle Tree, and the root of this tree is the Agent Root.
This makes it easy for the Remote Chains to detect if there has been a change in the set of bonded agents. Because this Agent Root is
one of the fields in the [Attestation](#attestation), the Notaries are constantly attesting to the current Agent Root to its [Remote Chain](#remote-chain).
When the remote chain has a new Agent Root, there is an [Optimistic Period](#optimistic-period) that needs to pass before that new Agent Root is considered valid.
Once it is considered valid, then off-chain agents can submit a proof using that Agent Root that proves they are part of the new valid agent set.
The Agent Root is a critical security property because a malicious Agent Root could allow a chain to be convinced of a malicious agent, which could open the door
to a malicious message being executed.

**Agent Root formed from Merkle Tree of Agent Infos:**
![AgentRoot](../../static/img/AgentRoot.png 'Diagram of AgentRoot formed from Merkle Tree of Agent Infos')

### Agent Status
Every bonded agent will have a given status consisting of the following fields
1.  **Flag**: Indicates the status. The options are as follows:
    1.  Unknown: This agent was not found in the agent set.
    2.  Active: This is an active agent.
    3.  Unstaking: Agent has unstaked and is waiting for the [Unbonding Period](#unbonding-period).
    4.  Resting: Agent still has posted bond but currently down and not actively participating in the protocol for whatever reason.
    5.  Fraudulent: Agent has been accused of fraud by a [Guard](#guard).
    6.  Slashed: Agent was found guilty of fraud and slashed.
2.  **Domain**: Will be 0 if it is a [Guard](#guard) or set to a specific chain id if it is a [Notary](#notary).
3.  **Index**: The position in the list of agents states that this one was inserted into.

See the Diagram under the section explaining [Agent Root](#agent-root) which also depicts the [Agent Set](#agent-set) consisting of all the
Agent Statuses.

### Agent Set
The set of [bonded agents](#bonded-agent) that are currently active.

See the Diagram under the section explaining [Agent Root](#agent-root) which also depicts the Agent Set consisting of all the [Agent Statuses](#agent-status).

### Attestation
This is what [Notaries](#notary) sign and post to the chain that it is assigned to, and it contains crucial information that is used
to prove messages and also to prove the [Agent Set](#agent-set). The [Notary](#notary) signs an attestation and posts it to the [Destination](#destination) chain.
However, in order to be considered a valid attestation, it must first be registered on the [Synapse Chain](#synapse-chain)
as a result of a [Notary](#notary) submitting a [State Snapshot](#state-snapshot).
If an Attestation is not first submitted to the Synapse chain, it will not be considered valid by the Destination chain.
If the Attestation turns out to be [fraudulent](#fraud), the [Notary](#notary) will be [slashed](#slash) and removed as a valid [Notary](#notary). Thus, it is very important
that the [Notary](#notary) only sign Attestations that contain information that has been thoroughly confirmed. Below is the data contained in the attestation:
1. [Snap Root](#snap-root) is the Merkle root of the Origin [States](#state) that were grouped together in a [state snapshot](#state-snapshot) and made into a Merkle tree.
This snap root is used to prove that a particular Origin state did in fact occur on the [Origin](#origin).
2. [Agent Root](#agent-root) is the Merkle root of the bonded agent data that can be used to prove if a particular agent is part of the [Agent Set](#agent-set).
3. [Gas Data Snapshot](#gas-data-snapshot) contains the [Gas Data](#gas-data) of each of the chains being attested to.
4. Nonce is the total number of accepted Notary snapshots and serves to uniquely identify this attestation.
5. Block Number of the block on the [Synapse Chain](#synapse-chain) that the attestation was registered by a Notary on the [Synapse Chain](#synapse-chain), which does not have to be the same [Notary](#notary) posting to the [Destination](#destination-chain).
6. Timestamp is the time that the attestation was registered on the [Synapse Chain](#synapse-chain).

This is what the [Attestation](glossary.md/#attestation) might look like:

![ExampleAttestation](../../static/img/ExampleAttestation.png 'Example Attestation')

### Attestation Fraud Report
The "Attestation Fraud Report" is how a [Guard](#guard) protects a [remote chain](#remote-chain) when it detects a
[Fraudulent Attestation](#fraudulent-attestation). Because the [fraud resolution](#fraud-resolution) happens
on the [Synapse Chain](#synapse-chain) and takes time to propagate to the [remote chain](#remote-chain),
the [Guard](#guard) sumbits an "Attestation Fraud Report" so that the [remote chain](#remote-chain) disallows
any message using an [attestation](#attestation) from the [Accused Notary](#accused-notary).

### Bond
The cryptocurrency held in escrow in order to disincentivize fraud from one of the [Offchain Agents](#off-chain-agent) is
called the "Bond". If the agent is caught committing fraud, it will have it's bond [slashed](#slash). The idea is that
if the bond is high enough to make fraud unprofitable, then claims made by the agents can be trusted based on the
game theory assumption that the agents are rational.

### Bonded Agent
The [Root of Trust](#root-of-trust) of the [Synapse Messaging System](#synapse-messaging-system) comes from the [Off Chain Agents](#off-chain-agent)
posting a [bond](#bond). Within the [Synapse Messaging System](#synapse-messaging-system), there are two kinds of Bonded Agents: [Notaries](#notary)
and [Guards](#gaurd). The only kind of [Unbonded Agent](#unbonded-agent) is the [Executor](#executor).

### Bonding Manager Smart Contract
The Bonding Manager Smart Contract is deployed on the [Synapse Chain](#synapse-chain) along with the [Summit](#summit-smart-contract) and
the [Inbox](#inbox-smart-contract). See the source code of [BondingManager.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/manager/BondingManager.sol).
The BondingManager keeps track of all existing agents on the Synapse Chain.
It utilizes a dynamic Merkle Tree to store the agent information. This enables passing only the
latest merkle root of this tree (referenced as the Agent Merkle Root) to the remote chains,
so that the agents could "register" themselves by proving their current status against this root.
BondingManager is responsible for the following:
1.  Keeping track of all existing agents, as well as their statuses. In the MVP version there is no token staking,
which will be added in the future. Nonetheless, the agent statuses are still stored in the Merkle Tree, and
the agent slashing is still possible, though with no reward/penalty for the reporter/reported.
2.  Marking agents as "ready to be slashed" once their fraud is proven on the local or remote chain. Anyone could
complete the slashing by providing the proof of the current agent status against the current Agent Merkle Root.
3.  Sending Manager Message to remote `LightManager` to withdraw collected tips from the remote chain.
4.  Accepting Manager Message from remote `LightManager` to slash agents on the Synapse Chain, when their fraud
is proven on the remote chain.

### Canonical Source of Truth
In the world of [Cross Chain](#cross-chain) messaging, there is the problem of only being able
to perform atomic transactions on one blockchain at a time. For any piece of information stored, one of the blockchains
will serve as the canonical source of truth. If two or more chains disagree, the chain that is the canonical source of truth gets
to decide the real state, and the other chains will need to be corrected to match the correct state.

### Client Smart Contract Application Developer
The developer who is creating an application that requires [Cross Chain Messaging](#cross-chain-messaging-system) is
referred to as the "Client Smart Contract Application Developer".
These are the customers of the [Synapse Messaging System](#synapse-messaging-system).

### Client Receiving Smart Contract
If a Smart Contract application requires sending cross chain messages, it must develop and deploy
a Smart Contract to receive messages sent by the [Client Sending Smart Contract](#client-sending-smart-contract).
This is referred to as the "Client Receiving Smart Contract". This is the contract that
implements the interface [IMessageRecipient.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/interfaces/IMessageRecipient.sol)
in order to provide the "receiveBaseMessage" method. (Note that in theory the same Smart Contract could implement both the
sender and receiver, but of course it would be the sender of one chain sending to a different chain.)

Please see the examples here for reference (note that each combines the [sender](#client-sending-smart-contract) and receiver):
1.  [TestClient.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/TestClient.sol)
2.  [PingPong.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/PingPongClient.sol)

### Client Sending Smart Contract
If a Smart Contract application requires sending cross chain messages, it must develop and deploy
a Smart Contract to send messages. This is referred to as the "Client Sending Smart Contract". This is the contract that
interacts with the Synapse [Origin Smart Contract](#origin-smart-contract) by calling the Origin's "sendBaseMessage" method.

Please see the examples here for reference (note that each combines the sender and [receiver](#client-receiving-smart-contract)):
1.  [TestClient.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/TestClient.sol)
2.  [PingPong.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/client/PingPongClient.sol)

### Commitment
In cryptography, a commitment is often used when someone wants to commit to large amounts of data without having to pass around
all the data. For example, if there is a message that consists of many terabytes, that message can use a cryptographic hash
function to come up with a 32 byte hash representation of that message, and it would be probabilistically impossible for someone to come up with another message
that has the same 32 byte hash. It is a way for someone to say "I have a very large message that I will send you later, but I promise the message hashes to this small 32 bytes".
Later, when the message is given, it can be checked that it in fact hashes correctly and we know it was not altered.

### Cross Chain
As opposed to [On Chain](#on-chain), Cross Chain refers to communication between one blockchain and another blockchain. If a someone wants to
send a message from one chain to another chain, this cannot be done in a single blockchain transaction. The only way for Cross Chain
communication to occur is with the help of [Off Chain](#off-chain) agents who observe transactions on various chains and then submit necessary transactions
to other chains in order to communicate across chains.

### Cross Chain Messaging System
A system that allows a [message](#message) to be sent [Cross Chain](#cross-chain) from a [Client Sending Smart Contract](#client-sending-smart-contract)
on one chain to a [Client Receiving Smart Contract](#client-receiving-smart-contract) on another chain is a
"Cross Chain Messaging System".

### Destination Chain
The chain where the [Message](#message) is being sent to is known as the Destination Chain.

### Destination Smart Contract
The Destination Smart Contract is deployed on all the [Remote Chains](#remote-chain) along with the [Origin](#origin-smart-contract), [Light Inbox](#light-inbox-smart-contract) and
the [Light Manager](#light-manager-smart-contract).
(Note that since the [Synapse Chain](#synapse-chain) can also be a participant in the network just like any [Remote Chain](#remote-chain), it also will have Origin and Destination contracts deployed on it).
See the source code of [Destination.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/Destination.sol).
The Destination contract is used for receiving messages from other chains. It relies on
Notary-signed statements to get the truthful states of the remote chains. These states are then
used to verify the validity of the messages sent from the remote chains.
The Destination is responsible for the following:
1.  Accepting the Attestations from the local Inbox contract.
2.  Using these Attestations to execute the messages (see parent `ExecutionHub`).
3.  Passing the Agent Merkle Roots from the Attestations to the local LightManager contract, if deployed on a non-Synapse chain.
4.  Keeping track of the remote domains GasData submitted by Notaries, that could be later consumed by the local [GasOracle](#gas-oracle-smart-contract) contract.

### Digital Signatures
Digital Signatures are a fundamental tool used in cryptography that allows for proving that a specific actor signed
a specific message. More can be read [here](https://en.wikipedia.org/wiki/Digital_signature).

### Disputed Agent
If a [Guard](#guard) detects that another agent (either another [Guard](#guard) or a [Notary](#notary)) has signed a fraudulent claim, the accusing [Guard](#guard)
will first notify the [Victim Chain](#victim-chain) that the fraud happened and that it suspects the guilty agent. Because the [Fraud Resolution](#fraud-resolution)
happens on a different chain, the [fraud report](#fraud-report) is submitted on the [resolving chain](#resolving-chain), and the [resolution](#fraud-resolution) will be communicated to the other
chains through the combination of [System Messages](#system-message) and updating the [Agent Root](#agent-root). In the mean time,
the [Guard](#guard) will alert the [Victim Chain](#victim-chain) about the pending [Fraud Report](#fraud-report), and both the accusing [Guard](#guard) and the accused agent will be considered
in a state of dispute and will not be trusted. The exception is if the
accusing [Guard](#guard) submits another [fraud report](#fraud-report) and in the same transaction pays a [Just In Time Bond](#just-in-time-bond) which will be given back
once the [fraud report](#fraud-report) is [resolved](#fraud-resolution).

### Disputed Agent Set
Any time [fraud](#fraud) is reported to a [Victim Chain](#victim-chain), both the accusing Guard and the accused Bonded Agent will be placed in the
disputed agent set while the [Fraud Resolution](#fraud-resolution) process happens. During this time, even though the disputed agents are
in the active [Agent Set](#agent-set), the victim chain will not trust any claim signed by these agents.

### Executor
The Executor is an off-chain agent that does the work of delivering messages that have passed through the
[Optimistic Period](#optimistic-period). The Executors are not in a position to commit fraud and therefore do
not need to post a bond. They are eligible to collect [Tips](#tips) for the work they do in delivering messages.

### Fraud
In the Synapse Messaging, Fraud is any time a bonded [Off Chain Agent](#off-chain-agent) signs a claim of something that turns out to be false.
When this happens, the bonded agent will be [slashed](#slash). There are two properties that need to hold for detecting such fraud.
1. The agent will have digitally signed the claim so we know that the agent is guilty. This assumes the agent is the only one in possession of the bonded address's private key, so this is the responsibility of the agent.
2. The claim needs to be proven false by the appropriate Smart Contract on the [Resolving Chain](#resolving-chain).
For example, if there is a fraudulent claim that an [Origin](#origin) chain had a particular state some time in the past, the [Origin](#origin) chain will
be the [resolving chain](#resolving-chain) since it is able to decide whether that is true or not.

### Fraud Report
When a [Guard](#guard) discovers that a [Notary](#notary) or other [Guard](#guard) has committed [Fraud](#fraud), it will submit a Fraud report that includes proof of the fraud.
The proof of the fraud will be to show that the guilty agent signed something that is false. If the Fraud Report turns out to be wrong,
the [Guard](#guard) who submitted it will be slashed. Otherwise, if it is a valid fraud report, then the guilty agent will be slashed and
the reporting Guard will receive the reward.

### Fraud Resolution
When there is a fraud report, the resolution will be either that the Guard submitting the report is wrong or the accused agent is wrong.
Depending on the claim that is being reported as fraudulent, there will be a single chain that can decide if this is a true claim or not.
If the claim is in fact fraudulent, the agent who signed that claim will be slashed. Otherwise, if the claim is true, the Guard
submitting the fraud report will be slashed. Because the slashing needs to occur on the [Synapse chain](#synapse-chain), if the resolution
happens on a [Remote Chain](#remote-chain), then that remote chain will need to send a [System Message](#system-message) to the
Synapse chain. Because the resolution will always result in either the accuser or the accused agent being slashed, this will result
in that agent being removed from the agent set, and that means the [Agent Root](#agent-root) will be updated. This will eventually
propagate to all the chains in the network and they will learn about the new [Agent Set](#agent-set) this way.

### Fraudulent Attestation
This is [fraud](#fraud) committed by a [Notary](#notary) by submitting an [Attestation](#attestation) to its [remote chain](#remote-chain)
that is not registered on
the [Synapse Chain](#synapse-chain).

### Fraudulent Attestation Detection
[Fraudulent Attestation](#fraudulent-attestation) is easily detected by a [Guard](#guard) who observes the [Attestation](#attestation)
posted on the [remote chain](#remote-chain) and then doing a look-up on the [Synapse Chain](#synapse-chain) for that [Attestation](#attestation).
If the [Synapse Chain](#synapse-chain) has no record of that attestation, then the [Guard](#guard) will submit a
[fraud report](#fraud-report) to the [Synapse Chain](#synapse-chain)

### Fraudulent Attestation Report
A [fraud report](#fraud-report) submitted to the [Synapse Chain](#synapse-chain) by a [Guard](#guard) who
[detects a fraudulent attestation](#fraudulent-attestation-detection) commited by a [Notary](#notary).
1. Guard calls submitAttestationReport on remote chain so remote chain puts both in dispute.
    1. What happens if Guard just griefs remote chain with this?
    2. Iterate through guard reports by index with getGuardReport (all agents should do this, why not?)
    3. Check if its a valid attestation on syn chain
    4. If not, then call verifyAttestationReport on Inbox.
    5. If the guard is malicious, it gets slashed, otherwise sender just wastes gas.
    6. Sender should receive guard's bond?
2. Guard calls verifyAttestation on Inbox.sol
2. verifyAttestation ends up slashing Notary (or Guard)
3.

### Gas Data
The Gas Data holds important information about the current gas prices of a particular chain. This Gas Data is part of the [State](#state) of
each chain, and it will be updated whenever there is a change in gas prices above a certain threshold. The goal is to avoid updating the Gas Data
for small variations and only update if there is a significant change in prices. The Gas Data includes the following information:
1. Gas price for the chain (in Wei per gas unit).
2. Calldata price (in Wei per byte of content).
3. Transaction fee safety buffer for message execution (in Wei).
4. Amortized cost for attestation submission (in Wei).
5. Chain's Ether Price / Mainnet Ether Price (in BWAD).
6. Markup for the message execution (in BWAD).

### Gas Data Snapshot
A Gas Data Snapshot is nothing more than a list of Gas Data from more than one chain, and is used to batch together gas information from
multiple chains. Whenever a Notary posts an [Attestation](#attestation) to its [Destinaion Chain](#destination-chain), it will pass along the Gas Data Snapshot so that
chain can update its local Gas Oracle with the latest information about Gas Prices on the other chains.

**Gas Data Snapshot formed from the Gas Data from multiple chains:**
![GasSnapshot](../../static/img/GasSnapshot.png 'Diagram of Gas Data Snapshot formed from the Gas Data from multiple chains')

### Gas Oracle
The Gas Oracle is a Smart Contract deployed on each of the chains in the network that tracks the estimated gas prices on other chains.
This is needed to estimate the cost of gas to send a message in the messaging system. The sender of the message pays up front for the transactions
required on both the [Synapse Chain](#synapse-chain) and the [Destination Chain](#destination-chain), and the Gas Oracle is what allows
the [Origin Chain](#origin-chain) to estimate how much should be collected.

### Gas Oracle Smart Contract
The Gas Oracle Smart Contract is deployed on the chains and provides the service of the [Gas Oracle](#gas-oracle)
See the source code of [GasOracle.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/GasOracle.sol).
The GasOracle contract is responsible for tracking the gas data for both local and remote chains.
#### Local gas data tracking
1.  GasOracle is using the available tools such as "tx.gasprice" to track the time-averaged values
for different "gas statistics" (to be implemented in the future).
2.  These values are cached, so that the reported values are only changed when a big enough change is detected.
3.  In the MVP version the gas data is set manually by the owner of the contract.
4.  The reported values are included in [Origin's State](#state), whenever a new message is sent.
5.  This leads to cached "chain gas data" being included in the [Guard and Notary snapshots](#state-snapshot).
#### Remote gas data tracking
1.  To track gas data for the remote chains, GasOracle relies on the Notaries to pass the gas data alongside
their attestations.
2.  As the gas data is cached, this leads to a storage write only when the gas data for the remote chain changes significantly.
3.  GasOracle is in charge of enforcing the optimistic periods for the gas data it gets from [Destination](#destination).
4.  The optimistic period is smaller when the "gas statistics" are increasing, and bigger when they are decreasing. The reason for that is that the decrease of the gas price leads to lower execution/delivery tips, and we want the
Executors to be protected against that.

### Guard
The Guard is an off-chain agent that participates in delivering messages and more importantly in catching fraud committed by
Notaries and other Guards. If a Guard succeeds at catching fraud, it is eligible to receive the bond posted by the guilty agent.
As fraud happens more rarely, this is not the only way for the Guard to earn rewards. By submitting [state snapshots](#state-snapshot) (information about the states of
chains) to the [Synapse Chain](#synapse-chain), Guards are rewarded with [tips](#tips). The size of the bond of the Guard is significantly less than that of the Notary because
the primary fraud a Guard can do primarily just results in denial-of-service (attacking [Liveness](#liveness)).

### Inbox Smart Contract
The Inbox Smart Contract is deployed on the [Synapse Chain](#synapse-chain) along with the [Summit](#summit-smart-contract) and
the [Bonding Manager](#bonding-manager-smart-contract).
See the source code of [Inbox.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/inbox/Inbox.sol).
The Inbox
1.  Accepts Guard and Notary Snapshots and passes them to [Summit](#summit-smart-contract).
2.  Accepts Notary-signed Receipts and passes them to [Summit](#summit-smart-contract).
3.  Accepts Receipt Reports to initiate a dispute between Guard and Notary.
4.  Verifies Attestations and Attestation Reports, and slashes the signer if they are invalid.

### Integrity
Integrity is a property of the messaging system that means a chain cannot be fooled into thinking a message
was sent when it never was really sent.

### Just In Time Bond
Whereas normally the bond is escrowed on the [Synapse Chain](#synapse-chain), there are use cases for [Just In Time Gaurds](#just-in-time-guard) to
alert a [Victim Chain](#victim-chain) of fraud even though it is not in the active set of agents. The Just In Time Bond is collected in the same
transaction as the fraud report, and the bond will be released once the [Victim Chain](#victim-chain) receives the [fraud resolution](#fraud-resolution).
Note that this is not yet implemented in the current version.

### Just In Time Guard
There are edge cases when a chain either does not have any registered Guards, or there is a Guard that is momentarily in the
[Disputed Agent Set](#disputed-agent-set). If either the Disputed Guard or perhaps some other actor would like to submit a fraud report to
a [Victim Chain](#victim-chain), there is an option to submit the fraud report along with a [Just In Time Bond](#just-in-time-bond) that is collected within the same transaction
right on the spot. Note that this is not yet implemented in the current version.

### Light Inbox Smart Contract
The Light Inbox Smart Contract is deployed on all the [Remote Chains](#remote-chain) along with the [Origin](#origin-smart-contract), [Destination](#destination-smart-contract) and
the [Light Manager](#light-manager-smart-contract).
See the source code of [LightInbox.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/inbox/LightInbox.sol).
The LightInbox:
1.  Accepts Notary Attestations and passes them to the [Destination](#destination) contract.
2.  Accepts Attestation Reports and initiates a dispute between the [Notary](#notary) and the [Guard](#guard).

### Light Manager Smart Contract
The Light Manager Smart Contract is deployed on all the [Remote Chains](#remote-chain) along with the [Origin](#origin-smart-contract), [Destination](#destination-smart-contract) and
the [Light Inbox](#light-manager-smart-contract).
See the source code of [LightManager.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/manager/LightManager.sol).
The LightManager keeps track of all agents on chains other than Synapse Chain.
/// It uses the Agent Merkle Roots from the Notary-signed attestations to stay in sync with the [BondingManager](#bonding-manager-smart-contract).
/// The LightManager is responsible for the following:
1.  Accepting the Agent Merkle Roots (passing the optimistic period check) from the Destination contract.
2.  Using these roots to enable agents to register themselves by proving their status.
3.  Accepting Manager Message from [BondingManager](#bonding-manager-smart-contract) on the [Synapse Chain](#synapse-chain) to withdraw [tips](#tips).
4.  Sending Manager Messages to the [BondingManager](#bonding-manager-smart-contract) on [Synapse Chain](#synapse-chain) to [slash](#slash) agents, when their [fraud](#fraud) is proven.

### Liveness
Liveness is a property of the messaging system that means a message that is sent will be delivered within a
reasonable amount of time.

### Merkle Proof
The property of a Merkle Tree that makes it so handy is that a particular leaf node can be proven to exist in a Merkle tree
without providing the entire Merkle tree. All that is needed is the [Merkle Root](#merkle-root) and a Merkle Proof whose size is
logarithmic to the number of leaves in the Merkle Tree. Thus, if there are 2^32 number of leaves (over 4 billion leaves), the
Merkle Proof will only need 32 nodes. The nodes would be the ones along the path from the node we are trying to prove to the [Merkle Root](#merkle-root).

### Merkle Root
When a [Merkl Tree](#merkle-tree) is formed, it will have a unique root that is 32 bytes in size. This Merkle Root
can serve as a cryptographic [commitment](#commitment) to ALL the data contained in the Merkle Tree, which could be terabytes in size or more.

### Merkle Tree
A Merkle Tree is a fundamental building block used in cryptography that organizes a group of leaf nodes containing
arbitrary data in a way that allows for very small cryptographic [commitments](#commitment) and relatively short and fast proofs that a
particular node exists in the tree without sending the entire tree.
Please see [Merkle Trees](https://www.simplilearn.com/tutorials/blockchain-tutorial/merkle-tree-in-blockchain) for a description.

### Message
The Message is the raw payload that a sender wants delivered to the destination contract.

See the Diagram under the section explaining [Message Merkle Root](#message-merkle-root) which also depicts the [Message Merkle Tree](#message-merkle-tree) consisting of all the messages sent from a particular chain.

### Message Merkle Tree
Each chain has an [Origin Smart Contract](#origin-smart-contract) that keeps a list of ordered messages sent from it.
This list of ordered messages is put into a [Merkle Tree](#merkle-tree) with a capacity of (2^32 - 1) messages.

See the Diagram under the section explaining [Message Merkle Root](#message-merkle-root) which also depicts the Merkle Tree consisting of all the [Messages](#message).

### Message Merkle Root
The 32 byte [merkle root](#merkle-root) of the [Message Merkle Tree](#message-merkle-tree).

**Message Merkle Root formed from Merkle Tree of Messages, with the simplification of having a height of only 3 rather than 32::**
![MessageMerkleRoot](../../static/img/MessageMerkleRoot.png 'Diagram of Message Merkle Root formed from Merkle Tree of Messages')

### Notary
The Notary is an off-chain agent that is assigned to a specific chain and has the very important job of posting [attestations](#Attestation) to its chain
that can then be used to prove messages. If a fraudulent attestation is posted, an attacker could fool the destination into
executing a malicious message, so the Notary plays a crucial role in maintaining [Integrity](#integrity)

### Off Chain
Anything that happens outside of a blockchain is referred to as Off Chain. If a transaction happens [On Chain](#on-chain) on one
blockchain, there is no way for another blockchain to know about this without the help of Off Chain agents. These Off Chain agents
can look at the transactions on one chain and then submit transactions to other chains in order to convey this [Cross Chain](#cross-chain)
information.

### Off Chain Agent
When trying to do [Cross Chain](#cross-chain) communication, there needs to be an Off Chain Agent that observes transactions
on one chain and submits transactions to other chains in order to communicate what happened on the origin chain.
The term "Off Chain Agent" is very general and could be as simple as a human being looking at something that happened on
the first chain and then submitting a transaction to the second chain. More typically, this agent is software
written to do this job in an automated way. The important question is how can the second chain trust this agent.
In the specific case of the [Synapse Messaging System](#synapse-messaging-system), this trust is based on a bond that the agent must post and any fraud
can be detected and will result in that agent losing the bond.

### On Chain
The term "On Chain" refers to a transaction that happens on a single blockchain, which means it has all the security
guarantees of that particular chain. Within the context of discussing cross-chain communication, "on chain" transactioins
are assumed to be trustworthy so long as the probability of a chain reorg is extremely low.

### Optimistic Pause
If a [Guard](#guard) believes that a [Notary](#notary) has submitted a [fraudulent](#fraud) [attestation](#attestation) to its [Destination](#destination-chain),
the actual [fraud resolution](#fraud-resolution) needs to be decided on another chain, either the [Synapse Chain](#synapse-chain) or the
[Origin](#origin-chain). Because of this, we allow the Guard to optimisitcally pause the [Destination](#destination) chain which puts
both the accused Notary and the reporting Guard in dispute. Until the resolution is communicated to that destination chain,
that attestation and the Notary are not trusted by that destination, and the Guard would need to pay a significant amount to
submit additional reports on that chain (to avoid denial of service).

### Optimistic Period
This is a crucial property of messaging system that is set and enforced by the client Smart Contract of the messaging system.
It is the time that a message must wait to be executed in order to give the Guards time to catch potential fraud. The longer
this time means the Guards have more time to catch fraud and the less likely it is for an attacker to fool a Destination
into executing a fraudulent message. There is also an Optimistic Period when a new [Agent Root](#agent-root) is proposed, because
we want to give [Guards](#guard) enough time to challenge the proposed root before accepting it as valid.

### Origin Chain
The chain where the message is being sent from is known as the Origin Chain.

### Origin Smart Contract
The Origin Smart Contract is deployed on all the [Remote Chains](#remote-chain) along with the [Destination](#destination-smart-contract), [Light Inbox](#light-inbox-smart-contract) and
the [Light Manager](#light-manager-smart-contract).
(Note that since the [Synapse Chain](#synapse-chain) can also be a participant in the network just like any [Remote Chain](#remote-chain), it also will have Origin and Destination contracts deployed on it).
See the source code of [Origin.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/Origin.sol).
The Origin contract is used for sending messages to remote chains. It is done
by inserting the message hashes into the Origin Merkle, which makes it possible to
prove that message was sent using the Merkle proof against the Origin Merkle Root. This essentially
compresses the list of messages into a single 32-byte value that needs to be stored on the [destination chain](#destination-chain).
The Origin is responsible for the following:
1.  Formatting the sent message payloads, and inserting their hashes into the Origin Merkle Tree.
2.  Keeping track of its own historical [states](#state).
3.  Enforcing minimum [tip](#tips) values for sent base messages based on the provided execution requests.
4.  Distributing the collected tips upon request from a local AgentManager contract.

### Permissioned
As opposed to [permissionless](#permissionless), a permissioned system requires someone who is in a position of authority
to give permission to someone who wants to do something. For example, if becoming an agent required the signature of
a quorum of administrators, this would make it a permissioned system. The long term goal of the [Synapse Messaging System](#synapse-messaging-system) is
to become entirely permissionless once the ecosystem has been bootstrapped.

### Permissionless
The [Synapse Messaging System](#synapse-messaging-system) was designed to allow agents to participate without the need of gaining permission from any special
authority. The only requirement is to post a stake in the case of [guards](#guard) and [notaries](#notary). For [executors](#executor),
no stake is required and anyone can run a node to act as an executor.

### Receipt
When a message is delivered on the Destination chain, all of the agents who participated in delivering the message are owed
[Tips](#tips), which are to be handed out on the Synapse chain and not on the Destination chain. Upon delivering the message,
a receipt is produced on the Destination chain and a Notary for that chain can sign and submit this receipt to
the Synapse chain, and the tips will be distributed at that time. Of course, if the Notary signs a fake receipt, it
can be found guilty of fraud by a Guard and get slashed.

### Remote Chain
Any chain in the network that is not the Synapse chain is referred to as a "Remote Chain". Because the Synapse chain
is where Bonded Agents post their bond, that is the canonical source of truth regarding who is a registered agent.
Part of the protocol therefore requires this information to get propagated to the Remote Chains. The Synapse chain is also
where agents submit "Snapshots" containing information about other chains in the network, and its the job of the Notary agents
to communicate valid "Snap Roots" to the Remote Chains.

### Resolving Chain
By definition, [Fraud](#fraud) is when an [Off Chain Agent](#off-chain-agent) makes a claim to a chain about something that happened
on another chain. (It would be pointless to attempt to lie to a chain about itself because the chain would easily reject the transaction since it knows its own state).
Thus, the chain that is able to resolve a dispute about a claim made to another chain is called the "Resolving Chain". When the resolving chain
determines that fraud did in fact happen, it will need to communicate this to the [Victim Chain](#victim-chain).

### Root of Trust
A common problem in layered systems is that each layer must trust the layer before, and if the layer before is compromised, then
all trust is lost at all the layers after. At the very first layer, trust must be bootstrapped somehow in order to establish
the foundation of trust, or "Root of Trust". In the context of the [Synapse Messaging System](#synapse-messaging-system), a message receiver trusts in the
fact that a Bonded Agent must stake a large amount of value and will be certain to lose it if it signs something that is not true.
What's more, if an honest agent is alert to catching the fraud, it will receive the stake that is lost, and given a large
enough [Optimistic Period](#optimistic-period), the probability for a [Guard](#guard) to catch and repoprt cheating is high.
Thus, a message receiver trusts that:
1.  [Notaries](#notary) don't want to lose their stake since it is quite large.
2.  [Guards](#guard) want to catch a fraudulent Notary because they will receive a large reward for doing so.
3.  The [Optimistic Period](#optimistic-period) is in the control of the [Client Sending Smart Contract](#client-sending-smart-contract)
and the [Client Receiving Smart Contract](#client-receiving-smart-contract) and will be set high enough to give [Guards](#guard)
enough time to prevent the message from being executed.

In conclusion, the trust is based on the Bonded Agents, which means the "Root of Trust" must be establishing the set of
Bonded Agents. In the very beginning, Synapse Messaging bootstraps this root of trust through a [permissioned](#permissioned)
mechanism whereby when a new [remote chain](#remote-chain) is added, the [Notary](#notary) for that chain posts the bond to
the [Bonding Manager Smart Contract](#bonding-manager-smart-contract) deployed on the [Synapse chain](#synapse-chain).
This results in the Bonding Manager having a new [Agent Root](#agent-root) that must now propagate to all the remote chains.
Because the new chain does not have an active [Notary](#notary) under the old [Agent Root](#agent-root), there would be nobody
to let the new chain know about the new [Agent Root](#agent-root). To boostrap the "Root of Trust",
when the [Destination Smart Contract](#destination-smart-contract) is deployed on the [remote chain](#remote-chain),
it is initialized with the new [Agent Root](#agent-root). This will allow the new chain's [Notary](#notary) to provide a proof of inclusion
to the new chain's [Light Manager Smart Contract](#light-manager-smart-contract), and from then on that chain
will trust it's [Notary](#notary).

### Slash
If an agent is found guilty of fraud, the punishment is to slash the bond posted on the [Synapse Chain](#synapse-chain)

### Snap Root
When a bonded agent submits a [snapshot](#state-snapshot) (i.e. a list of States) to the Inbox contract on the Synapse chain,
the list of states will be inserted into a Merkle tree, where the leaves of the Merkle tree is
essentially a hash of the State. The root of the Merkle tree is known as the "Snap Root", which is short for
"Snapshot Merkle Root".

**Snap Root formed from Merkle Tree of Snapshot States:**
![SnapRoot](../../static/img/SnapRoot.png 'Diagram of SnapRoot formed from Snapshot of States')

### Stake
Synonym of [bond](#bond).

### State Snapshot
The Synapse messaging protocol uses the term snapshot to describe a list of Origin [States](#state).
The bonded agents (i.e. Notaries and Guards) periodically observe all the chains in the network and track
the latest "States" of those chains. For all the chains whose State has changed, the bonded agent will
update the Inbox contract on the Synapse chain by sending a list of all the new States. Thus, a snapshot is
just a way to batch the states of multiple chains in order to reduce the number of calls.

See the Diagram under the section explaining [Snap Root](#snap-root) which also depicts the Snapshot of States.

### State
Each chain in the network at a given point in time will have values set for the following properties that define its "state":
1. Root of merkle tree of messages that have been sent from this chain to another chain. Read more on [Merkle Trees](https://www.simplilearn.com/tutorials/blockchain-tutorial/merkle-tree-in-blockchain).
2. Origin Chain ID is the chain id that identifies this particular blockchain.
3. Nonce is the number of messages that have been sent from this chain to another chain.
4. Block Number is the current block number of this chain's tip.
5. Timestamp is the time when the current tip was added to the chain.
6. [Gas Data](#gas-data) contains information about recent gas rates on this chain so other chains can estimate gas costs
of performing necessary transactions on remote chains.

See the Diagram under the section explaining [Snap Root](#snap-root) which also depicts the [Snapshot](#state-snapshot) of States, with a
zoom in example State.

### Summit Smart Contract
The Summit Smart Contract is deployed on the [Synapse Chain](#synapse-chain) along with the [Bonding Manager](#bonding-manager-smart-contract) and
the [Inbox](#inbox-smart-contract). See the source code of [Summit.sol](https://github.com/synapsecns/sanguine/blob/master/packages/contracts-core/contracts/Summit.sol).
The Summit contract is the cornerstone of the Synapse messaging protocol. This is where the states of all the remote chains (provided collectively by the Guards and Notaries) are stored. This is
also the place where the tips are distributed among the off-chain actors.
Summit is responsible for the following:
1.  Accepting Guard and Notary snapshots from the local `Inbox` contract, and storing the states from these snapshots (see parent contract `SnapshotHub`).
2.  Accepting Notary Receipts from the local `Inbox` contract, and using them to distribute tips among the
off-chain actors that participated in the message lifecycle.

### Synapse Messaging System
The [Cross Chain Messaging System](#cross-chain-messaging-system) developed by Synapse.

### Synapse Chain
The [Synapse Chain](https://docs.synapseprotocol.com/protocol/synapse-chain) is a blockchain developed originally for the
[Synapse Bridge](https://docs.synapseprotocol.com/protocol/synapse-bridge).
In the [Synapse Messaging System](#synapse-messaging-system), the Synapse chain has special Smart Contracts deployed on it that serve as a
central hub when sending messages from one chain to another. Bonds are posted on the Synapse chain so this serves as the
canonical source of truth of who is a valid agent. An important part of the protocol is keeping the other [Remote Chains](#remote-chain) in sync
with what is on the Synapse chain. As another example of how the Synapse chain is special, the Bonded Agents observe the states of all the chains
in the network and first submit these states to the Inbox Smart Contract deployed on the Synapse chain.

### System Message
In the [Synapse Messaging System](#synapse-messaging-system), there are special "System Messages" that are not sent by a client but rather are for
communicating things like fraud resolution from one chain to another. System Messages go through the same path as normal
messages.

### Tips
Tips are the rewards that the off-chain agents earn for doing the work of delivering messages.

### Unbonded Agent
An [Offchain Agent](#off-chain-agent) that does valuable work in the [Synapse Messaging System](#synapse-messaging-system) but is not required
to post a [bond](#bond) because it is not in a position to sign any claims that need to be trusted. The only kind
of Unbonded Agent in the [Synapse Messaging System](#synapse-messaging-system) is the [Executor](#executor).

### Unbonding Period
When a bonded agent decides to unregister, it is of course eligible to receive its bond back so long as it has not been found
guilty of fraud. The bonded agent will submit the request to unbond on the [Synapse Chain](#synapse-chain), and it will be placed
in the unbonding state. After the Unbonding Period has passed and no fraud has been detected, the agent can claim its bond.
The purpose of the Unboding Period is to prevent a Bonded Agent from commiting [fraud](#fraud) and then escaping with its bond before there is enough
time to [slash](#slash) the agent.

### Victim Chain
If a malicious bonded agent signs a payload that is [fraudulent](#fraud), it could attempt to trick another chain or set of chains
into believing the false claim. We refer to these chains as "Victim Chains". Because the [fraud resolution](#fraud-resolution) process takes some time,
the Guard who detected the fraud will alert the "Victim Chains" about the pending [fraud report](#fraud-report) and the "Victim Chain"
will consider both the accused agent and the accuser to be in the [Disputed Agent Set](#disputed-agent-set). This means that any claim made by the accused agent will
not be trusted and also the accusing Guard will not be able to make additional accusation of fraud without posting a [Just In Time Bond](#just-in-time-bond)
on the Victim Chain at the time it reports the fraud.
