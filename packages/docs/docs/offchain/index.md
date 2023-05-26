---
sidebar_position: 1
---

# Participating in the Network

Cover overview of what a guard, notary and executor is. Link to individual pages

## What is cross-chain messaging?

It is assumed that the reader of this document is already familiar with the current state of the universe when it comes to blockchain technology, which is that there are many different blockchains, and at times there is a need to communicate from one chain to another.

Cross-chain messaging aims to allow smart contracts to send and receive messages to and from different chains. This is done by enabling smart contracts to interact with each other through the cross-chain messaging system.

A cross-chain messaging system offers a fundamental building block that smart contract developers can use for sending any message (i.e. any arbitrary “string of bytes”) from one chain to another. It is up to the application to interpret what the “string of bytes” represents. For example, it could be a message indicating that tokens were locked on the sending chain, and that could mean that the receiving chain can unlock a corresponding number of tokens.

All cross-chain messaging systems aim to accomplish two goals, which can at times appear to compete with each other:

1. Liveness: If a message is sent on one chain destined for another chain, there is an expectation that the message will arrive on the destination chain within a “reasonable” amount of time.
2. Integrity: For all chains in the network, it must never be the case that a particular chain becomes convinced that there was a message sent to it from another chain that didn’t actually happen. THIS IS EXTREMELY CRITICAL!

Take the example of a cross-chain bridging application whereby tokens on one chain can be exchanged for tokens on another chain. In this example, we have a user Alice who has 10 tokens on chain A and would like to exchange them for 10 tokens on chain B.

Alice sends her 10 tokens to a special “Bridge” smart contract on chain A which locks the tokens.
The “Bridge” smart contract on chain A now needs to let chain B know about this transfer so chain B can give Alice 10 tokens.
The “Bridge” smart contract decides to use a dedicated “messaging system” for communicating from chain A to chain B. It formulates an application specific message as a “string of bytes” and tells the messaging system to send those bytes to the “Bridge” contract living on chain B.
The messaging system has no idea what the “string of bytes” means, however, it fulfills its job of delivering the message to the “Bridge” contract on chain B.
Upon receiving the “string of bytes”, the “Bridge” contract on chain B has code that knows how to interpret the message into an application specific action. In this case, it decodes the message as something like “Alice has locked 10 tokens on chain A, please give her 10 tokens on chain B”. Upon receiving this, the “Bridge” contract on chain B feels confident that it can safely release 10 tokens to Alice.

For the developers of the “Bridge” smart contract, things are greatly simplified by utilizing the “messaging system” as a black box. The only thing the Bridge smart contract needs to do is define what kinds of messages it sends and how to serialize and deserialize those messages as “strings of bytes”. .

Now, let’s dive deeper into what can go wrong if the “messaging system” fails at one of its two objectives of Liveness and Integrity.

Failure in delivering Liveness: In the Bridge example, let’s say that after Alice locks 10 tokens on chain A in step 1, the messaging system fails to deliver the message to chain B within a reasonable amount of time. Alice has given up her 10 tokens on chain A, and is waiting for her 10 tokens on chain B. During this time, she is incurring an opportunity cost because she has given away 10 tokens. If it takes years to get the tokens on chain B, this will certainly be the last time she uses this Bridge application and it will be a matter of time before the Bridge is out of business.
Failure in maintaining Integrity: Let’s say Alice does not actually send 10 tokens to chain A, but somehow fools chain B into accepting the message in step 5 above and now Chain B is fooled into thinking that Alice has locked 10 tokens on Chain A, and thus gives Alice 10 tokens on chain B. Alice now has gained 10 tokens on Chain B for nothing. If she did it once, why not do it again and again? Within a short period of time, Alice could drain Chain B’s store of tokens which could be valued at hundreds of millions of dollars. This has in fact happened in several high profile attacks on cross-chain Bridges, which is why Integrity is so critical. Chain B should NEVER be fooled into believing that a message was sent when it wasn’t actually sent.

Clearly, the developers of the “Bridge” application who choose to rely on a “messaging system” to communicate from one chain to another are putting a huge amount of trust in the messaging system. A poorly designed messaging system could spell disaster. It can’t be repeated enough that this is not just theoretical. It actually has happened to the tune of billions of US dollars lost in these [Cross-chain Bridge Attacks](https://www.coindesk.com/layer2/2022/10/14/blockchain-bridges-keep-getting-attacked-heres-how-to-prevent-it/).

This leads to the introduction of the Synapse Carbon messaging system that is designed to enable both Liveness and Integrity, with the priority of always being to maintain Integrity over Liveness since a breakdown of Integrity is what leads to the Black Swan disasters described above.

## How does Synapse Carbon work?

Synapse Carbon is a cross-chain messaging system designed to be used by Smart Contract developers who need to send messages from one chain to another. Synapse Carbon gives the Smart Contract developers the ability to choose the balance between Liveness and Integrity. If a message is not that important, then perhaps the application can tolerate the occasional risk to Integrity in preference to having the message delivered sooner rather than later. On the other hand, if there is a potential for a Black Swan disaster if a fraudulent message is accepted, then the application can choose to err on the side of having the message delayed (i.e. sacrifice Liveness) in order to ensure that it is correct.
<br/>
To accomplish this trade-off between Liveness and Integrity, Synapse Carbon uses a type of messaging called Optimistic Messaging. The “optimistic” in Optimistic Messaging means that the network uses an optimistic time period where fraud can be reported. This lets application developers determine their own optimistic periods, using longer ones for more important messages, and shorter ones for less important messages.

## Your Cross-Chain Application built with Carbon

From the perspective of the Smart Contract application developers, it is trivial to use the Synapse Carbon system to send a message from your smart contract on Chain A to one on Chain B. As already mentioned, the Smart Contract developer simply needs to define the sorts of messages that it intends to send and a way to serialize and deserialize those messages as strings of bytes. It is up to the Smart Contract on the Destination chain to decide what action to take upon receiving a given message.

## Network Composition

Before taking a deeper dive into the Synapse Carbon messaging system, it makes sense to step back once again and look at the big picture of what any messaging system needs to do.

Watch all the chains in the network and look for messages that are sent from one chain to another chain. For Liveness, there should not be any messages that are missed so it’s important that all messages are detected by the system.
Convince the destination chain that a given message was in fact sent such that if the Destination chain accepts the message, there is an extremely high probability that the message did in fact happen. The higher the probability, the higher the Integrity of the system.

Now let’s look at the specific mechanism employed by the Synapse Carbon messaging system:

There are dedicated agents called “Notaries” whose job it is to maintain Liveness in the system by observing all the chains in the network, and whenever a message is sent from an “Origin” chain to a “Destination” chain, all it takes is a single Notary to propose to the Destination chain that the particular message happened.
At this point, the message is not yet accepted by the application's Smart Contract on the Destination chain, but rather it is in a probationary state that lasts a certain amount of time specified by the application. This could be 1 second for a very unimportant message or it could be more than a day for a very important message.
During the probationary period, there are a different set of dedicated agents called “Guards” who observe all the messages that have been “proposed” on the various Destination chains, and if a fraudulent message is found, all it takes is one honest Guard to report that fraud and prevent the message from being accepted.
Clearly, the more time the Guards have from the time a message is proposed until the time when the message is accepted, the more likely it is that at least one honest Guard will detect the fraud and prevent it from happening.
After the “Optimistic Period” is over, the message can now move from the probationary state to actually being accepted by the Destination Smart Contract.
There is a final set of dedicated agents called “Executors” whose job it is to move messages from the probationary state to the accepted state.

In summary, the Synapse Carbon network has 3 different types of agents that do very important work in delivering messages from one chain to another while maintaining Liveness and Integrity.

[Notaries](notary): Maintain Liveness by observing messages from Origin chains and proposing them on Destination chains.

[Guards](guard): Maintain Integrity by observing proposed messages from Notaries on Destination chains and calling out if there is a proposed message that never happened.

[Executors](executor): Maintain Liveness by transitioning messages from the probationary state to the accepted state on the Destination chains so long as no Guard has reported the message fraudulent within the Optimistic Time Period.

What incentivizes each of these agents to do this work?
Agents in the system are doing work (i.e. running servers, paying electricity, etc), so they need to have the proper incentives to make it worthwhile. This means whoever is sending a cross-chain message needs to pay some sort of toll to have the message sent, and the agents all receive a fraction of the toll in proportion to the work they do to get the message delivered.

## Who is allowed to participate as one of the agents?

The Synapse Carbon network is designed to be decentralized and allow anyone to participate as any of the agent types. This means that anyone is able to be a Notary, Guard and/or Executor. However, the network needs to disincentivize fraudulent behavior to maintain integrity.
