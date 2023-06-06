---
sidebar_position: 2
---

# Synapse Messaging Deep Dive

Synapse Messaging System is a cross-chain messaging system designed to be used by Smart Contract developers who need to send messages from one chain to another.
Synapse Messaging gives the Smart Contract developers the ability to choose the balance between Liveness and Integrity.
If a message is not that important, then perhaps the application can tolerate the occasional risk to Integrity in preference to having the message delivered sooner rather than later.
On the other hand, if there is a potential for a Black Swan disaster if a fraudulent message is accepted, then the application can choose to err on the side of having the message delayed (i.e. sacrifice Liveness) in order to ensure that it is correct.

To accomplish this trade-off between Liveness and Integrity, Synapse Messaging uses a type of messaging called Optimistic Messaging. The “optimistic” in Optimistic Messaging means that the network uses an optimistic time period where fraud can be reported.
This lets application developers determine their own optimistic periods, using longer ones for more important messages, and shorter ones for less important messages.

## Your Cross-Chain Application built with Synapse Messaging

From the perspective of the Smart Contract application developers, it is trivial to use the Synapse Messaging System to send a message from your smart contract on Chain A to one on Chain B.
As already mentioned, the Smart Contract developer simply needs to define the sorts of messages that it intends to send and a way to serialize and deserialize those messages as strings of bytes.
It is up to the Smart Contract on the Destination chain to decide what action to take upon receiving a given message.

## Network Composition

Before taking a deeper dive into the Synapse Messaging System, it makes sense to step back once again and look at the big picture of what any messaging system needs to do.

1. Watch all the chains in the network and look for messages that are sent from one chain to another chain. For Liveness, there should not be any messages that are missed so it’s important that all messages are detected by the system.
2. Convince the destination chain that a given message was in fact sent such that if the Destination chain accepts the message, there is an extremely high probability that the message did in fact happen. The higher the probability, the higher the Integrity of the system.

Now let’s look at the specific mechanism employed by the Synapse Messaging System:

1. There are dedicated agents called “Notaries” whose job it is to maintain Liveness in the system by observing all the chains in the network, and whenever a message is sent from an “Origin” chain to a “Destination” chain, all it takes is a single Notary to propose to the Destination chain that the particular message happened.
2. At this point, the message is not yet accepted by the application's Smart Contract on the Destination chain, but rather it is in a probationary state that lasts a certain amount of time specified by the application. This could be 1 second for a very unimportant message or it could be more than a day for a very important message.
3. During the probationary period, there are a different set of dedicated agents called “Guards” who observe all the messages that have been “proposed” on the various Destination chains, and if a fraudulent message is found, all it takes is one honest Guard to report that fraud and prevent the message from being accepted.
4. Clearly, the more time the Guards have from the time a message is proposed until the time when the message is accepted, the more likely it is that at least one honest Guard will detect the fraud and prevent it from happening.
5. After the “Optimistic Period” is over, the message can now move from the probationary state to actually being accepted by the Destination Smart Contract.
6. There is a final set of dedicated agents called “Executors” whose job it is to move messages from the probationary state to the accepted state.

In summary, the Synapse Messaging network has 3 different types of agents that do very important work in delivering messages from one chain to another while maintaining Liveness and Integrity.

1. [Notaries](../offchain/notary): Maintain Liveness by observing messages from Origin chains and proposing them on Destination chains.

2. [Guards](../offchain/guard): Maintain Integrity by observing proposed messages from Notaries on Destination chains and calling out if there is a proposed message that never happened.

3. [Executors](../offchain/executor): Maintain Liveness by transitioning messages from the probationary state to the accepted state on the Destination chains so long as no Guard has reported the message fraudulent within the Optimistic Time Period.

What incentivizes each of these agents to do this work?
Agents in the system are doing work (i.e. running servers, paying electricity, etc), so they need to have the proper incentives to make it worthwhile. This means whoever is sending a cross-chain message needs to pay some sort of toll to have the message sent, and the agents all receive a fraction of the toll in proportion to the work they do to get the message delivered.

## Who is allowed to participate as one of the agents?

The Synapse Messaging network is designed to be permissionless and allow anyone to participate as any of the agent types. This means that anyone is able to be a Notary, Guard and/or Executor. However, the network needs to disincentivize fraudulent behavior to maintain integrity.

# Message Flow Overview

Before diving deeper into the specifics of each of the offchain agents, it is useful to learn how a message flows from one chain to another with Synapse.
<br/>
At the highest level, there is a smart contract on the chain that we are sending the message from, known hereafter as the "Sending Chain",
and there is a smart contract on the chain that we are sending the message to, known hereafter as the "Receiving Chain".
<br/>
Note that the "Client Smart Contracts" that we are talking about will be whatever smart contract that wants to use the Synapse Messaging System as a client.
For example, if someone would like to build a cross-chain bridge to exchange tokens from one chain to tokens of another chain,
the Smart Contract Developers building such a bridge can leverage the Synapse Messaging System to send important messages used in this application.
<br/>
<br/>
We are now going to take a look at the system starting from the very highest vantage point, then diving increasingly deeper into how the Synapse Messaging System works.

### Mile High View of the Messaging System

1. Client Smart Contract sends a message on the "Sending Chain".
2. Client Smart Contract receives the message on the "Receiving Chain" within a "reasonable" amount of time.
   <br/>
   <br/>
   This is literally the highest possible level to view the system, and it is the goal to have it be this simple from the point of view of the Smart Contract Developers.

### Pulling Back the Curtains

Now how does it actually work?
1. The Client Smart Contract gives the Synapse "Origin" contract the message as raw bytes and lets it know the "Receiving Chain" and the address of the "Receiving Client Smart Contract".
2. The Synapse "Origin" Smart contract keeps a Merkle tree of messages with a capacity of 2^32 messages.
3. Each time a message is inserted into the Synapse Origin Smart Contract, it updates its Merkle tree state with 2 important pieces of information. First is the number of total messages stored and the second is the Merkle root of the tree after that last message was inserted.
4. Note that given the Merkle root of the Merkle tree, it is easy to prove any of the messages that have been inserted up to that point in time.
5. This means that if the "Receiving Chain" can be convinced that the "Sending Chain" at one time was in a state of having a given number of messages sent and a particular Merkle root at that time, then any message that happened prior can be proven to have been inserted.
6. Therefore, it is the role of the off-chain agents first to convince the "Receiving Chain" of a valid state of the "Sending Chain", and once such a state is accepted as valid, any message prior can be delivered to that "Receiving Chain".

### Attesting to States of Sending Chain
The goal of the off-chain agents is to convince the Receiving Chain that a particular state occurred on the Sending chain, where the state consists of:
1. Total number of messages sent.
2. The Merkle root of the Merkle tree at the time that number of messages was sent.

Let's say for example that we had utter confidence that when the 1 millionth message was sent from Chain A, the Merkle root was 0xDEADBEEF.
What can we do with this knowledge? We can submit a Merkle Proof (of size roughly 1 Kilobyte) for any of the messages between 1 and a million.

#### Synapse Optimistic Messaging
What does Synapse do to achieve the state of "utter confidence" that the Sending Chain had a particular state in its Merkle tree?
Synapse uses staked agents and a strategy of optimistic attestation. This is how it works at a high level:

1.  A Notary who is a special agent with a LARGE stake held in escrow submits a signed attestation and submits to the "Receiving Chain" something along the lines of "This Sending Chain has sent 1 million messages and has a Merkle Root of 0xDEADBEEF".
2.  If this Notary is lying, it will lose its LARGE stake so it is incentived to be honest.
3.  Any message between 1 and a million can now be proven against that Merkle Root of 0xDEADBEEF, but each message has been assigned an "optimistic time" so it must wait that long from when the Notary attested to that state and when the message is executed.
4.  Other agents called Guards with a significantly less stake held in escrow can challenge what the Notaries attest to, and if a Guard catches a Notary lying, the Guard is eligible to receive the LARGE Notary stake.

### Diving into the Technical Details
At this point, we are ready to look at the specific steps that happen. It is good to look at this flow and then we will be ready to dig even deeper into the specifics of each agent in future pages.
1. Client Smart Contract dispatches a message to the Synapse Origin Smart contract, providing the raw message payload, the Receiving Chain ID and the Receiving Chain Smart Contract Address.
2. The Sending Chain has a Synapse Origin Smart Contract and inserts the message into its Merkle tree, and thus obtains a new Merkle state consisting of the number of total messages and the Merkle root.
3. Each Guard agent plays a role in attesting to the Origin Chain states. This is done primarily to give the Guards something to do in addition to detecting fraud so they can receive rewards as part of the normal message flow.
4. Each Guard observes the current "state" of each Origin Chain in the network, and puts together a list of all the origin states from each chain that it wants to attest to.
5. The Guard signs the list of states and Submits this to a special Smart Contract on the Synapse blockchain. The Synapse blockchain serves as a central hub for all the other chains in the Network.
6. After at least one Guard has attested to the states of the various Origin chains, a Notary can put together a list of those states which have been registered on the Synapse chain.
7. When a Notary attests to a list of states, it also submits to the Smart Contract on the Synapse chain, and those list of states are automatically Merklized to produce a unique Merkle root as a 32 byte string.
8. This Merkle Root of states can now be used by any of the Notaries who are assigned to specific Receiving chains. A Notary can only submit attestations on a single chain.
9. A Notary can take one of the registered Merkle Roots of states and sign and attest to that on its Destination chain.
10. There is another special agent called Executor that observes all the Origin chains for all the messages and it will try to execute those messages.
11. To execute a message, the Executor must identify the Receiving chain, and it must wait until one of the Notaries for that Receiving Chain submits an attestation that includes the proper state for the Sending Chain.
12. The proper state is any state where the desired message was before the number of messages in the Merkle tree at the time of the state.
13. The Executor must first prove that the Sending Chain's state is part of the Merkle Root of States.
14. After proving that the Origin state was part of that Merkle Root of States, the Executor then needs to prove that the message was part of the State's Merkle Root of Messages.
15. Assuming the message being submitted has an optimistic period that is less than the time when the attestation happened and the current time, then the Receiving Smart contract can execute the message.

Next, we can talk about the various kinds of "fraud" that can get the agent's slashed.

# Fraud in Synapse Messaging

The two kinds of bonded agents are Guards and Notaries. Notaries require a much larger stake due to the potential harm they could cause.

## How a Notary commits Fraud
1. Notary can sign and attest to a Merkle Root of States that never happened on the Synapse chain.
   This is easy to catch because a Guard can simply observe the attestation on the Receiving Chain and check if it has already been registered on the Synapse chain.
2. Notary can sign and attest to a Merkle Root of States that is registered on the Synapse chain,
   but that Merkle Root includes an invalid state.
   For this, the Guard needs to prove that the invalid state is part of that Merkle Tree of States
   and then the Sending Chain will quickly know that the state is not valid and be able to slash the Notary that signed the attestation.
3. Note that for the Merkle Root of States to be registered on Synapse chain,
   there must have been a Guard who signed for the state followed by some
   other Notary (could be same or could be different Notary) who also submitted that state to the Synapse chain.
4. To Summarize, in the event that a Receiving chain has an bad attestation posted for a particular Merkle Root of States,
   the Notary for the Receiving chain must have signed it and will get slashed.
   If the bad Merkle Root of States is registered on the Synapse chain, the Notary who submitted the list of states
   that included that bad state will get slashed AND also the Guard who originally registered the state will get slashed.
5.  Who receives the slashed bond? The Guard who reports each of the agents. Note there must be a fraud report for each agent being slashed.

## How a Guard commits Fraud
1. A Guard can submit an invalid state to the Synapse chain.
2. A Guard can submit an invalid fraud report for the purposes of waging a denial of service attack.
3. A Guard has the power to put any attestation posted in dispute, and resolving the dispute requires cross-chain coordination and takes time.
4. While a Guard submitted a bad fraud report will eventually get slashed, if the Guard's stake is too small, it might fail to disincentive griefing from malicious Guards.

# Proving Bonded Agents

Thus far, we have glossed over the very important detail of how do we know a Notary is a Notary and a Guard is Guard.
We mentioned that a Notary posts a LARGE stake and a Guard posts a smaller stake that is large enough to disincentivize griefing.

This section will explain the details of how the stake is collected and how slashing occurs.
Additionally, we need to talk about how the list of Notaries and Guards are kept in sync across all the remote chains.

## Synapse Chain as Central Hub
The Synapse chain is where each agent must post bond which will be held in Escrow by a special Smart Contract deployed there.
The name of this smart contract on Synapse chain is called the Bonding Manager. Upon posting a bond, the agent will be registered as either a Guard or Notary.
Each Notary is assigned to a specific chain, while each Guard can service all the chains.

These are the steps of becoming an agent.
1. Bond is posted to the BondingManager smart contract on Synapse chain along with what type of agent it will be.
2. Upon receiving the bond, the BondingManager adds the new agent to the list of agents.
3. The list of agents if Merklized and a Merkle Root of Agents is registered.
4. The 32 bytes Merkle Root of Agents is now the current latest state.
5. Note that upon receiving the stake, the BondingManager on Synapse chain immediately knows who is a Notary and who is a Guard.
6. The challenge now is communicating this to the remote chains.

### Purpose of Synapse Chain as Central Hub
When sending a message from chain A to chain B, it seems superfluous to position the Synapse Chain as an intermediary when
attesting to the state of chain A. It is a
fair question to ask why is this necessary. The decision to architect the Synapse Messaging System this way came about to
facilitate centralizing where the bonded agents place their bonds and also it is the central place to decide when an agent
is slashed. Furthermore, it makes it more difficult for an attacker to commit fraud because the Synapse Chain as a central hub
creates an added layer that the attacker would need to deceive.

### Updating the Remote Chains with new Agent Set
1. The Merkle Root of Agents will be stored with every attestation from the previous section where we produce Merkle Roots of Origin States.
2. Thus, any time a Notary posts an attestation for a Merkle Root of Origin States on its particular Receiving chain, it will also attest to the Merkle Root of Agents at the time the Merkle Root of States was posted on Synapse chain.
3. Just as there is an optimistic period before the Merkle Root of States can be used to execute a message, there will be an optimistic period before an agent can prove itself as a member of the latest Merkle Root of Agents.
4. After the optimistic period, the Remote Chain will allow any agent to prove itself as being part of the set of agents.
5. Once a new Merkle Root of Agents is accepted by a remote chain, any agent must re-register itself against that root in order for the remote chain to trust it.

### What happens during a fraud dispute?
1. If a Guard accuses a Notary of fraud, the remote chain will temporarily put both the Notary and Guard on probation while the dispute is resolved.
2. For the dispute to be resolved, decisions made by other chains will need to be submitted to the remote chain that is in dispute.
3. The message about who was right and who was wrong will arrive similar to any other message, which means there needs to be a non-disputed Guard and Notary able to send messages.
4. Once the dispute is resolved, either the Guard or Notary will have been removed from the list of valid agents and their stake will go to the Guard reporting the fraud.

# Rewarding Agents

As mentioned there are 3 kinds of agents:
1. Notaries who post a large bond and are assigned to one Receiving chain.
2. Guards who post a smaller bond and can report fraud and also submit valid Origin states to the Synapse chain that can be used by Notaries when attesting to a list of states.
3. Executors actually execute messages after the optimistic period has passed for an attested state that includes that message in the Merkle Tree of Messages.

## Working for tips.
When a client wants to send a message, it must obviously pay the gas for performing the transaction on the Sending chain.

However, for the agents to be incentived to do any work to move that message to the Receiving Chain, the sender needs to provide enough money to pay for the gas on the other chains and also an additional amount to incentivize the agents to do the work.

### Where do we pay Gas?
1.  The Sending Chain will obviously cost a certain amount of Gas.
2.  A Guard then needs to pay Gas to submit a list of states to Synapse chain, although a given state can be amortized for all the messages that came before.
3.  A Notary also needs to submit a list of states on the Synapse chain, again amortized for all messages coming before.
4.  For each Receiving chain that has a message sent to it, a Notary assigned to that chain must submit an attestation for a Merkle Root of States which can be amortized for all messages prior.
5.  For each message to be executed, the Executor needs to pay Gas on the Receiving chain for every message executed.

### Who pays the Gas?
The Sender of the message must pay gas, but how does this work in terms of paying for gas on the Synapse chain and the Receiving chain?

Synapse uses "Gas Oracles" to estimate the amount of Gas needed to perform the work on the other chains so that the gas money plus the extra tips can be collected at the time the message is Sent on the Sending chain.

### Where are rewards paid?
Note the tips will be given to the agents on the Synapse chain. This means that periodically, the gas collected on the Origin chains will be bridged to the Synapse chain because the sender pays on the Sending chain.

When a message is executed on th Destination chain, all the agents who had a hand in delivering that message are eligible to receive a reward that should cover the gas as well as additional rewards to keep them incentivized.

Upon having the message executed, the Destination chain produces a receipt of the message being executed.
The Notary will need to take that receipt, sign it and submit on the Synapse chain in order for it and the other agents to receive tips.

Note that this is yet another potential for fraud from the Notary, because the receipt could in theory be fraudulent, but the Notary would be risking a very large stake for a small amount of tips.
This is another payload that a Guard can potentially look for fraud and report.

# Estimating Total Gas
Each chain in the Synapse network has a special Gas Oracle contract that is used to determine the current price of Gas on that chain.

What this means is that part of the Origin "state" will include the current cost of gas on that chain.

This means that when a remote chain hears about the state of another chain, it will get a rough idea of the going rate of gas.

Thus, when a message is sent from chain A to chain B, chain A will estimate the cost of gas on the Syanpse chain and also the cost of gas on chain B, and it will collect a sufficient amount to cover the gas required to execute the message.


