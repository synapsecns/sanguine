---
sidebar_position: 3
---

# Message Flow Overview

Before diving deeper into the specifics of each of the offchain agents, it is useful to learn how a message flows from one chain to another with Synapse.
<br/>
At the highest level, we have a smart contract on the chain that we are sending the message from, known hereafter as the "Sending Chain",
and there is a smart contract on the chain that we are sending the message to, known hereafter as the "Receiving Chain".
<br/>
Note that the "Client Smart Contracts" that we are talking about will be whatever smart contract that wants to use the Synapse Messaging System as a client.
For example, if someone would like to build a cross-chain bridge to exchange tokens from one chain to tokens of another chain,
the Smart Contract Developers building such a bridge can leverage the Synapse Messaging System to send important messages used in this applicaton.
<br/>
<br/>
We are now going to take a look at the system starting from the very highest vantage point, then diving increasingly deeper into how the Synapse Messaging System works.

### Mile High View of the Messaging System

1. Client Smart Contract sends a message on the "Sending Chain".
2. Client Smart Contract receives the message on the "Receiving Chain" within a "reasonable" amount of time.
<br/>
<br/>
This is literally the highest possible level to view the system, and it is the goal of the Synapse team to have it be this simple from the point of view of the Smart Contract Developers.

### Pulling Back the Curtains

Now how does it actually work?
1. The Client Smart Contract gives the Synapse "Origin" contract the message as raw bytes and lets it know the "Receiving Chain" and the address of the "Receiving Client Smart Contract".
2. The Synapse "Origin" Smart contract keeps a Merkle tree of messages with a capacity of 2^32 messages.
3. Each time a message is inserted into the Synapse Origin Smart Contract, it updates its Merkle tree state with 2 important pieces of information. First is the number of total messages stored and the second is the Merkle root of the tree after that last message was inserted.
4. Note that given the Merkle root of the Merkle tree, it is easy to prove any of the messages that have been inserted up to that point in time.
5. This means that if the "Receiving Chain" can be convinced that the "Sending Chain" at one time was in a state of having a given number of messages sent and a particular Merkle root at that time, then any message that happened prior can be proven to have been inserted.
6. Therefore, it is the role of the offline agents first to convince the "Receiving Chain" of a valid state of the "Sending Chain", and once such a state is accepted as valid, any message prior can be delivered to that "Receiving Chain".

### Attesting to States of Sending Chain
The goal of the offline agents is to convince the Receiving Chain that a particular state occured on the Sending chain, where the state consists of:
1. Total number of messages sent.
2. The Merkle root of the Merkle tree at the time that number of messages was sent.

Let's say for example that we had utter confidence that when the 1 millionth message was sent from Chain A, the Merkle root was 0xDEADBEEF.
What can we do with this knowledge? We can submit a Merkle Proof (of size roughly 1 Kilobyte) for any of the messages between 1 and a million.

#### Synapse Optimistic Messaging
What does Synapse do to achieve the state of "utter confidence" that the Sending Chain had a particular state in its Merkle tree?
Synapse uses staked agents and a strategy of optimistic attestation. This is how it works at a high level:


1.  A Notary who is a special agent with a LARGE stake held in escrow submits a signed attestation and submits to the "Receiving Chain" something along the lines of "This Sending Chain has sent 1 million messages and has a Merkle Root of 0xDEADBEEF".
2.  If this Notary is lying, it will lose its LARGE stake so it is incentived to be honest.
3.  Any message between 1 and a million can now be proved against that Merkle Root of 0xDEADBEEF, but each message has been assigned an "optimistic time" so it must wait that long from when the Notary attested to that state and when the message is executed.
4.  Other agents called Guards with a significantly less stake held in escrow can challenge what the Notaries attest to, and if a Guard catches a Notary lying, the Guard is elligible to receive the LARGE Notary stake.

### Diving into the Technical Details
At this point, we are ready to look at the specific steps that happen. It is good to look at this flow and then we will be ready to dig even deeper into the specifics of each agent in future pages.
1. Client Smart Contract dispatches a message to the Synapse Origin Smart contract, providing the raw message payload, the Receiving Chain ID and the Receiving Chain Smart Contract Address.
2. The Sending Chain has a Synapse Origin Smart Contract and inserts the message into its Merkle tree, and thus obtains a new Merkle state consisting of the number of total messages and the Merkle root.
3. Each Guard agent plays a role in attesting to the Origin Chain states. This is done primarily to give the Guards something to do in addition to detecting fraud so they can receive rewards as part of the normal message flow.
4. Each Guard observes the current "state" of each Origin Chain in the network, and puts together a list of all the origin states from each chain that it wants to attest to.
5. The Guard signs the list of states and Submits this to a special Smart Contract on the Synapse blockchain. The Synapse blockchain serves as a central hub for all the other chains in the Network.
6. After at least one Guard has attested to the states of the various Origin chains, a Notary can put together a list of those states which have been registered on SYN chain.
7. When a Notary attests to a list of states, it also submits to the Smart Contract on SYN chain, and those list of states are automatically Merklized to produce a unique Merkle root as a 32 byte string.
8. This Merkle Root of states can now be used by any of the Notaries who are assigned to specific Receiving chains. A Notary can only submit attestations on a single chain.
9. A Notary can take one of the registered Merkle Roots of states and sign and attest to that on its Destination chain.
10. There is another special agent called Executor that observes all the Origin chains for all the messages and it will try to execute those messages.
11. To execute a message, the Executor must identify the Receiving chain, and it must wait until of the Notaries for that Receiving Chain submits an attestation that includes the proper state for the Sending Chain.
12. The proper state is any state where the desired message was before the number of messages in the Merkle tree at the time of the state.
13. The Executor must first prove that the Sending Chain's state is part of the Merkle Root of States.
14. After proving that the Origin state was part of that Merkle Root of States, the Executor then needs to prove that the message was part of the State's Merkle Root of Messages.
15. Assuming the message being submitted has an optimistic period that is less than the time when the attestation happened and the current time, then the Receiving Smart contract can execute the message.
<br/>
<br/>
Next, we can talk about the various kinds of "fraud" that can get the agent's slashed.
