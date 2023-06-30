---
sidebar_position: 1
---

# Protocol Overview

## How It Works

Decentralized agents follow a simple and secure four-step process to relay cross-chain messages on behalf of client-controlled smart contracts:

1. **Broadcast**:
   Cross-chain messages are broadcast by client-controlled smart contracts and sent to the Synapse origin chain contract.

```
DEST_CHAIN_ID:    123456
MESSAGE:          "Hello, World!"
DEST_ADDRESS:     0x1234…cdef
SECURITY_SECONDS: 60
TIP_USDC:         10
```

2. **Submit**:
   Off-chain agents watching the origin-chain contract compete to validate and submit messages to the Synapse destination-chain contract.

3. **Challenge**:
   Off-chain agents may watch for, and object to, fraudulently submitted messages during a security window set by the initiating contract.

4. **Execute**:
   Once the security window has passed, off-chain agents compete to relay the message to the client smart contract on the destination chain.

## Blockchain Execution

On the origin and destination chains, individual blockchain transactions are made on a per-message basis. The Synapse origin and destination contracts also record blockchain transactions, but in a batched method.

### Rewards

To reward participation, broadcast messages include a tip to be split among the submission and execution agents once the broadcast is sucessfully executed.

```
Recommended tip: $10
```

## Optimistic Security

To discourage fraudulent messages from being submitted, participating agents must register and post a bond which is forfeit in the event of a fraudulent message or report.

During the challenge window, reporting agents may detect and report fraud to the Synapse origin and destination contracts. Agents found to have acted dishonestly forfeit their bond to the reporting agent.

```
Submission bond: $2,000,000
Reporting bond:  $20,000
```

### Fraud Reporting

If a fraudulent message is observed, detecting agents submit a fraud report to the Synapse origin and destination contracts. The destination contract pauses message execution while the origin contract evaluates the report:

- If upheld, submitting agents forfeit their bond to the reporting agent
- If rejected, reporting agents forefit their bond to the Synapse origin contract
- Dishonest agents are removed from the system, in addition to forfeiting their bonds

Fraud evaluation takes approximately 30 seconds, after which message execution resumes among the remaining honest agents.

<!-- **Synapse Optimistic Messaging:**
![SynapseOptimisticMessaging](../../static/img/SynapseMessagingMileHigh.png 'Synapse Optimistic Messaging')

- **Step 1**: Client Smart Contract on Chain A tells Synapse Origin Contract to say "hello" to Chain B.
- **Step 2**: Off Chain Agents observe that Chain A client wants to say "hello" to Chain B.
- **Step 3**: Off Chain Agents tell the Synapse Destination Contract on Chain B that Chain A says "hello".
- **Step 4**: Destination waits for a period of time (in this case 1 hour) to give other Off Chain Agents time to object.
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

**Synapse Messaging Fraud Protection:**
![SynapseMessagingFraudProtection](../../static/img/SynapseMessagingFraudProtection.png 'Synapse Messaging Fraud Protection')

- **Step 1**: Malicious Off Chain Agent lies to Chain B and says Chain A said "goodbye".
- **Step 2**: Destination waits for a period of time (in this case 1 hour) to give other Off Chain Agents time to object.
- **Step 3**: An honest Off Chain Agent double checks with Chain A and checks if Chain A did in fact say "goodbye" to Chain B.
- **Step 4**: Chain A lets the honest Off Chain Agent know that it never said goodbye to Chain B.
- **Step 5**: The honest Off Chain Agent tells Chain B that the Malicious agent is lying, and Chain B will not allow the fraudulent message. -->
