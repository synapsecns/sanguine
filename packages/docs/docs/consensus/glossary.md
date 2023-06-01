---
sidebar_position: 8
---

# Glossary

### Agent Root
The Agent Root is the root of the Merkle Tree of formed from the [Agent Set](#agent-set). The list of all registered bonded agents
that have posted bond on the [SYN Chain](#synapse-chain) are put into a Merkle Tree, and the root of this tree is the Agent Root.
This makes it easy for the Remote Chains to detect if there has been a change in the set of bonded agents. Because this Agent Root is
one of the fields in the [Attestation](#attestation), the Notaries are constantly attesting to the current Agent Root to its chain.
When the remote chain has a new Agent Root, there is an Optimistic Period that needs to pass before that new Agent Root is considered valid.
Once it is considered valid, then off-chain agents can submit a proof using that Agent Root that they are part of the new valid agent set.

### Agent Set
The set of bonded agents ([Guards](#guard) and [Notaries](#notary)) that are currently active.

### Attestation
This is what Notaries sign and post to the chain that it is assigned to, and it contains crucial information that is used
to prove messages and also to prove the state of agents. The Notary signs an attestation and posts it to the destination chain.
However, in order to be considered a valid attestation, it must first be registered on the [SYN Chain](#synapse-chain)
as a result of a Notary submitting a [State Snapshot](#state-snapshot) there.
If an Attestation is not first submitted to the SYN chain, it will not be considered valid by the Destination chain.
If the Attestation turns out to be fraudulent, the Notary will be slashed and removed as a valid Notary. Thus, it is very important
that the Notary only sign Attestations that contain information that has been thoroughly confirmed. Below is the data contained in the attestation:
1. [Snap Root](#snap-root) is the Merkle root of the Origin [States](#state) that were grouped together in a [state snapshot](#state-snapshot) and made into a Merkle tree.
This snap root is used to prove that a particular Origin state did in fact occur.
2. [Agent Root](#agent-root) is the Merkle root of the bonded agent data that can be used to prove if a particular agent is part of the set of valid agents.
3. [Gas Data Snapshot](#gas-data-snapshot) contains the [Gas Data](#gas-data) of each of the chains being attested to.
4. Nonce is the total number of accepted Notary snapshots and serves to uniquely identify this attestation.
5. Block Number of the block on the [SYN Chain](#synapse-chain) that the attestation was registered by a Notary on the Summit chain, which does not have to be the same Notary posting to the destination.
6. Timestamp is the time that the attestation was registered on the [SYN Chain](#synapse-chain).

### Commitment
In cryptography, a commitment is often used when someone wants to commit to large amounts of data without having to pass around
all the data. For example, if there is a message that consists of many terabytes, that message can use a cryptographic hash
function to come up with a 32 byte hash representation of that message, and it would be probabilistically impossible for someone to come up with another message
that has the same 32 byte hash. It is a way for someone to say "I have a very large message that I will send you later, but I promise the message hashes to this small 32 bytes".
Later, when the message is given, it can be checked that it in fact hashes correctly and we know it was not altered.

### Cross Chain
As opposed to [On Chain](#on-chain), Cross Chain refers to communication between one blockchain and another. If a someone wants to
send a message from one chain to another chain, this cannot be done in a single blockchain transaction. The only way for Cross Chain
communication to occur is with the help of [Off Chain](#off-chain) agents who observe transactions on various chains and then submit necessary transactions
to other chains in order to communicate across chains.

### Destination Chain
The chain where the message is being sent to is known as the Destination Chain.

### Executor
The Executor is an off-chain agent that does the work of delivering messages that have passed through the
[Optimistic Period](#optimistic-period). The Executors are not in a position to commit fraud and therefore do
not need to post a bond. The are elligible to collect [Tips](#tips) for the work they do in delivering messages.

### Fraud
In the Synapse Messaging, Fraud is any time a bonded  [Off Chain Agent](#off-chain-agent) signs a claim of something that turns out to be false.
When this happens, the bonded agent will be [slashed](#slash). There are two properties that need to hold for detecting such fraud.
1. The agent will be guilty of digitally signing the claim so we know that the agent is the one guilty. This assumes the agent is the only one in possession of the bonded address.
2. The claim needs to be proven to be false by the appropriate Smart Contract on the appropriate chain.
For example, if there is a fraudulent claim that an Origin chain had a particular state some time in the past, that chain will
be able to decide whether that is true or not.

### Fraud Report
When a Guard discovers that a Notary or other Guard has committed Fraud, it will submit a Fraud report that includes proof of the fraud.
The proof of the fraud will be to show that the guilty agent signed something that is false. If the Fraud Report turns out to be wrong,
the Guard who submitted it will be slashed. Otherwise, if it is a valid fraud report, then the guilty agent will be slashed and
the reporting Guard will receive the reward.

### Fraud Resolution
When there is a fraud report, the resolution will be either that the Guard submitting the report is wrong or the accused agent is wrong.
Depending on the claim that is being reported as fraudulent, there will be a single chain that can decide if this is a true claim or not.
If the claim is in fact fraudulent, the agent who signed that claim will be slashed. Otherwise, if the claim is true, the Guard
submitting the fraud report will be slashed. Because the slashing needs to occur on the [SYN chain](#synapse-chain), if the resolution
happens on a [Remote Chain](#remote-chain), then that remote chain will need to send a [System Message](#system-message) to the
SYN chain. Because the resolution will always result in either the accuser or the accused agent being slashed, this will result
in that agent being removed from the agent set, and that means the [Agent Root](#agent-root) will be updated. This will eventually
propagate to all the chains in the network and they will learn about the new [Agent Set](#agent-set) this way.

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

### Gas Oracle
The Gas Oracle is a Smart Contract deployed on each of the chains in the network that tracks the estimated gas prices on other chains.
This is needed to estimate the cost of gas to send a message in the messaging system. The sender of the message pays up front for the transactions
required on both the [SYN Chain](#synapse-chain) and the [Destination Chain](#destination-chain), and the Gas Oracle is what allows
the [Origin Chain](#origin-chain) to estimate how much should be collected.

### Guard
The Guard is an off-chain agent that participates in delivering messages and more importantly in catching fraud committed by
Notaries and other Guards. If a Guard succeeds at catching fraud, it is elligible to receive the bond posted by the guilty agent.
As fraud happens more rarely, this is not the only way for the Guard to earn rewards. By submitting [state snapshots](#state-snapshot) (information about the states of
chains) to the [SYN Chain](#synapse-chain), Guards can receive [tips](#tips) for doing this required step in the protocol for
normal happy path message sending. The size of the bond of the Guard is significantly less than that of the Notary because
the primary fraud a Guard can do primarily just results in denial-of-service (attacking [Liveness](#liveness)).

### Integrity
Integrity is a property of the messaging system that means a chain cannot be fooled into thinking a message
was sent when it never was really sent.

### Liveness
Liveness is a property of the messaging system that means a message that is sent will be delivered within a
reasonable amount of time.

### Merkle Proof
The property of a Merkle Tree that makes it so handy is that a particular leaf node can be proven to exist in a Merkle tree
without providing the entire Merkle tree. All that is needed is the [Merkle Root](#merkle-root) and a Merkle Proof whose size is
logarithmic to the number of leaves in the Merkle Tree. Thus, if there are 2^32 number of leaves (over 4 billion leaves), the
Merkle Proof will only need 32 nodes along the path from the node we are trying to prove to the [Merkle Root](#merkle-root).

### Merkle Root
When a [Merkl Tree](#merkle-tree) is formed, it will have a unique root that is 32 bytes in size. This Merkle Root
can serve as a cryptographic commitment to ALL the data contained in the Merkle Tree, which could be terabytes in size or more.

### Merkle Tree
A Merkle Tree is a fundamental building block used in cryptography that organizes a group of leaf nodes containing
arbitrary data in a way that allows for very small cryptographic commitments and relatively short and fast proofs that a
particular node exists in the tree without sending the entire tree.
Please see [Merkle Trees](https://www.simplilearn.com/tutorials/blockchain-tutorial/merkle-tree-in-blockchain) for a description.

### Message
The Message is the raw payload that a sender wants delivered to the destination contract.

### Notary
The Notary is an off-chain agent that is assigned to a specific chain and has the very important job of posting attestations to its chain
that can then be used to prove messages. If a fraudulent attestation is posted, an attacker could fool the destination into
executing a malicious message, so the Notary plays a crucial role in maintaining [Integrity](#integrity)

### Optimistic Pause
If a [Guard](#guard) believes that a [Notary](#notary) has submitted a [fraudulent](#fraud) [attestation](#attestation) to its [Destination](#destination-chain),
the actual [fraud resolution](#fraud-resolution) needs to be decided on another chain, either the [SYN Chain](#synapse-chain) or the
[Origin](#origin-chain). Because of this, we allow the Guard to optimisitcally pause the [Destination](#destination) chain which puts
both the accused Notary and the reporting Guard in dispute. Until the resolution is communicated to that destination chain,
that attestation and the Notary are not truested by that destination, and the Guard would need to pay a significant amount to
submit additional reports on that chain (to avoid denial of service).

### Optimistic Period
This is a crucial property of messaging system that is set and enforced by the client Smart Contract of the messaging system.
It is the time that a message must wait to be executed in order to give the Guards time to catch potential fraud. The longer
this time means the Guards have more time to catch fraud and the less likely it is for an attacker to fool a Destination
into executing a fraudulent message.

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
In the specific case of the Synapse messaging system, this trust is based on a bond that the agent must post and any fraud
can be detected and will result in that agent losing the bond. So long as the [Optimistic Period](#optimistic-period) is high
enough to allow [Guards](#guard) to detect the fraud in time, the
opportunity to commit fraud that yields more reward than the bond that is slashed is low enough such that the second chain
can trust what the offchain agent says happened on the first chain.

### On Chain
The term "On Chain" refers to a transaction that happens on a single blockchain, which means it has all the security
guarantees of that particular chain. Within the context of discussing cross-chain communication, "on chain" transactioins
are assumed to be trustworthy so long as the probability of a chain reorg is extremely low.

### Origin Chain
The chain where the message is being sent from is known as the Origin Chain.

### Receipt
When a message is delivered on the Destination chain, all of the agents who participated in delivering the message are owed
[Tips](#tips), which are to be handed out on the SYN chain and not on the Destination chain. Upon delivering the message,
a receipt is produced on the Destination chain and a Notary for that chain can sign and sumbit this receipt to
the SYN chain, and the tips will be distributed at that time. Of course, if the Notary signs a fake receipt, it
can be found guilty of fraud by a Guard and get slashed.

### Remote Chain
Any chain in the network that is not the SYN chain is referred to as a "Remote Chain". Because the SYN chain
is where Bonded Agents post their bond, that is the canonical source of truth regarding who is a registered agent.
Part of the protocol therefore requires this information to get propagated to the Remote Chains. The SYN chain is also
where agents submit "Snapshots" containing information about other chains in the network, and its the job of the Notary agents
to communicate valid "Snap Roots" to the Remote Chains.

### Slash
If an agent is found guilty of fraud, the punishment is to slash the bond posted on the [SYN Chain](#synapse-chain)

### Snap Root
When a bonded agent submits a [snapshot](#state-snapshot) (i.e. a list of States) to the Inbox contract on the Synapse chain,
the list of states will be inserted into a Merkle tree, where the leaves of the Merkle tree is
essentially a hash of the State. The root of the Merkle tree is known as the "Snap Root", which is short for
"Snapshot Merkle Root".

### State Snapshot
The Synapse messaging protocol uses the term snapshot to describe a list of Origin [States](#state).
The bonded agents (i.e. Notaries and Guards) periodically observe all the chains in the network and track
the latest "States" of those chains. For all the chains whose State has changed, the bonded agent will
update the Inbox contract on the Synapse chain by sending a list of all the new States. Thus, a snapshot is
just a way to batch the states of multiple chains in order to reduce the number of calls.

### State
Each chain in the network at a given point in time will have values set for the following properties that define its "state":
1. Root of merkle tree of messages that have been sent from this chain to another chain. Read more on [Merkle Trees](https://www.simplilearn.com/tutorials/blockchain-tutorial/merkle-tree-in-blockchain).
2. Origin Chain ID is the chain id that identifies this particular blockchain.
3. Nonce is the number of messages that have been sent from this chain to another chain.
4. Block Number is the current block number of this chain's tip.
5. Timestamp is the time when the current tip was added to the chain.
6. [Gas Data](#gas-data) contains information about recent gas rates on this chain so other chains can estimate gas costs
of performing necessary transactions on remote chains.

### Synapse Chain
The [Synapse Chain](https://docs.synapseprotocol.com/protocol/synapse-chain) (aka SYN chain) is a blockchain developeed originally for the
[Synapse Bridge](https://docs.synapseprotocol.com/protocol/synapse-bridge).
In the new Synapse Messaging System, the SYN chain has special Smart Contracts deployed on it that serve as a
central hub when sending messages from one chain to another. Bonds are posted on the SYN chain so this serves as the
canonical source of truth of who is a valid agent. An important part of the protocol is keeping the other [Remote Chains](#remote-chain) in sync
with what is on the SYN chain. As another example of how the SYN chain is special, the Bonded Agents observe the states of all the chains
in the network and first submit these states to the Inbox Smart Contract deployed on the SYN chain.

### System Message
In the Synapse Messaging system, there are special "System Messages" that are not sent by a client but rather are for
communicating things like fraud resolution from one chain to another. System Messages go through the same path as normal
messages.

### Tips
Tips are the rewards that the off-chain agents earn for doing the work of delivering messages.
