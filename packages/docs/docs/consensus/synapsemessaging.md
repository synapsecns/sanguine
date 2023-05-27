---
sidebar_position: 2
---

# Synapse Messaging Overview

Synapse Carbon is a cross-chain messaging system designed to be used by Smart Contract developers who need to send messages from one chain to another. Synapse Carbon gives the Smart Contract developers the ability to choose the balance between Liveness and Integrity. If a message is not that important, then perhaps the application can tolerate the occasional risk to Integrity in preference to having the message delivered sooner rather than later. On the other hand, if there is a potential for a Black Swan disaster if a fraudulent message is accepted, then the application can choose to err on the side of having the message delayed (i.e. sacrifice Liveness) in order to ensure that it is correct.
<br/>
To accomplish this trade-off between Liveness and Integrity, Synapse Carbon uses a type of messaging called Optimistic Messaging. The “optimistic” in Optimistic Messaging means that the network uses an optimistic time period where fraud can be reported. This lets application developers determine their own optimistic periods, using longer ones for more important messages, and shorter ones for less important messages.

## Your Cross-Chain Application built with Carbon

From the perspective of the Smart Contract application developers, it is trivial to use the Synapse Carbon system to send a message from your smart contract on Chain A to one on Chain B. As already mentioned, the Smart Contract developer simply needs to define the sorts of messages that it intends to send and a way to serialize and deserialize those messages as strings of bytes. It is up to the Smart Contract on the Destination chain to decide what action to take upon receiving a given message.

## Network Composition

Before taking a deeper dive into the Synapse Carbon messaging system, it makes sense to step back once again and look at the big picture of what any messaging system needs to do.

1. Watch all the chains in the network and look for messages that are sent from one chain to another chain. For Liveness, there should not be any messages that are missed so it’s important that all messages are detected by the system.
2. Convince the destination chain that a given message was in fact sent such that if the Destination chain accepts the message, there is an extremely high probability that the message did in fact happen. The higher the probability, the higher the Integrity of the system.

Now let’s look at the specific mechanism employed by the Synapse Carbon messaging system:

1. There are dedicated agents called “Notaries” whose job it is to maintain Liveness in the system by observing all the chains in the network, and whenever a message is sent from an “Origin” chain to a “Destination” chain, all it takes is a single Notary to propose to the Destination chain that the particular message happened.
2. At this point, the message is not yet accepted by the application's Smart Contract on the Destination chain, but rather it is in a probationary state that lasts a certain amount of time specified by the application. This could be 1 second for a very unimportant message or it could be more than a day for a very important message.
3. During the probationary period, there are a different set of dedicated agents called “Guards” who observe all the messages that have been “proposed” on the various Destination chains, and if a fraudulent message is found, all it takes is one honest Guard to report that fraud and prevent the message from being accepted.
4. Clearly, the more time the Guards have from the time a message is proposed until the time when the message is accepted, the more likely it is that at least one honest Guard will detect the fraud and prevent it from happening.
5. After the “Optimistic Period” is over, the message can now move from the probationary state to actually being accepted by the Destination Smart Contract.
6. There is a final set of dedicated agents called “Executors” whose job it is to move messages from the probationary state to the accepted state.

In summary, the Synapse Carbon network has 3 different types of agents that do very important work in delivering messages from one chain to another while maintaining Liveness and Integrity.

1. [Notaries](notary): Maintain Liveness by observing messages from Origin chains and proposing them on Destination chains.

2. [Guards](guard): Maintain Integrity by observing proposed messages from Notaries on Destination chains and calling out if there is a proposed message that never happened.

3. [Executors](executor): Maintain Liveness by transitioning messages from the probationary state to the accepted state on the Destination chains so long as no Guard has reported the message fraudulent within the Optimistic Time Period.

What incentivizes each of these agents to do this work?
Agents in the system are doing work (i.e. running servers, paying electricity, etc), so they need to have the proper incentives to make it worthwhile. This means whoever is sending a cross-chain message needs to pay some sort of toll to have the message sent, and the agents all receive a fraction of the toll in proportion to the work they do to get the message delivered.

## Who is allowed to participate as one of the agents?

The Synapse Carbon network is designed to be decentralized and allow anyone to participate as any of the agent types. This means that anyone is able to be a Notary, Guard and/or Executor. However, the network needs to disincentivize fraudulent behavior to maintain integrity.


