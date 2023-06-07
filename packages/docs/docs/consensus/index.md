---
sidebar_position: 1
---

# Protocol Overview

The **Synapse Messaging System** allows for messages to be sent from a Smart Contract on one blockchain to a Smart Contract on another blockchain.

The protocol relies on a concept called **Optimistic Messaging** that results in minimizing the amount of work performed unless fraud is detected.

**Synapse Optimistic Messaging:**
![SynapseOptimisticMessaging](../../static/img/SynapseMessagingMileHigh.png 'Synapse Optimistic Messaging')
- **Step 1**: Client Smart Contract on Chain A tells Synapse Origin Contract to say "hello" to Chain B.
- **Step 2**: Off Chain Agents observe that Chain A client wants to say "hello" to Chain B.
- **Step 3**: Off Chain Agents tell the Synapse Destination Contract on Chain B that Chain A says "hello".
- **Step 4**: Destination waits for a period of time (in this case 1 hour) to give other Off Chain Agents time to object
- **Step 5**: After 1 hour with nobody objecting, Off Chain Agents tells the Client Smart Contract on Chain B that Chain A says "hello".

During the normal path of sending a message, the flow looks like this:
1. Client Smart Contract submits a blockchain transaction on the Origin blockchain, posting the following:
    1. **Message** to send.
    2. **Destination Blockchain ID** to send the message to.
    3. **Smart Contract Address** to deliver the message on the destination chain.
2. Off Chain Agents observe the posted message and propose the message to be executed on the Destination Smart Contract.
3. The system waits for a period of time to allow other Off Chain Agents to report fraud.
4. Assuming no fraud is reported, another Off Chain Agent executes the message by submitting a transaction on the Destination blockchain.

During this normal Happy Path, the only blockchain transactions that happen on a per-message basis is at steps 1 and 4, first sending the message on the Origin chain and second executing the message on the Destination chain.
Step 2 above does require blockchain transactions, however it can be amortized across a batch of messages. This is what is meant by "Optimistic" messaging because
in the event of no fraud, the amount of work done On-Chain is minimized.

If fraud is detected by an Off-Chain Agent during step 3, the steps look like this:
1. A fraud report is submitted by an Off-Chain Agent to the Destination chain, which "optimistically" pauses the ability to execute messages.
2. The same fraud report is submitted by the Off-Chain Agent to the Origin chain which decides if it was fraud or not.
3. Assuming there was fraud, the Origin lets the rest of the Network know by propagating a special system message.
4. Once the message from step 3 is received, the Fraudulent Agent is slashed and the Fraudulent claim will be discarded.
5. Messages can continue being sent by the remaining honest Off Chain Agents.

Under normal conditions, in order to send a message, there needs to be 3 different Off Chain Agents doing different jobs to get the message from one chain to another.

For cases when fraud occurs, all it takes is one honest Off Chain Agent to stop the fraud.


