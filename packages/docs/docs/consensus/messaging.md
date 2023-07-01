---
sidebar_position: 2
---

# Messaging

<!-- ## Sending a Message

1. To initiate a message, a client contract calls the `send message` method on the Synapse origin contract.
   This has the following effects:
   - The message is added to the origin's Message Merkle Tree
   - A state is added to the local chain with a new Message Merkle Root
2. On Synapse Chain, a guard submits a state snapshot that includes the new state of the origin chian.
3. On Synapse Chain, a notary submits a state snapshot that icnludes the new state of the origin chain.
4. Synapse Chain registers a new attestation that includes the snap and agent roots.
5. A notary on the destination chain proposes the attestation.
6. The security window begins.
7. The executor executres the message via a Merkle proofs for the state against the snap root in the attestation, and for the message against the message Merkle root in the state. -->

## Sending Messages

We will now illustrate sending a message from the [Client Sending Smart Contract](glossary.md/#client-sending-smart-contract)
on Chain A to the [Client Receiving Smart Contract](glossary.md/#client-receiving-smart-contract) on Chain B.

We will see how the [Synapse Messaging System](glossary.md/#synapse-messaging-system) supports both
[Liveness](glossary.md/#liveness) and [Integrity](glossary.md/#integrity).

1.  On Chain A, [Client Sending Smart Contract](glossary.md/#client-sending-smart-contract) calls the "send message" method on the [Origin Smart Contract](glossary.md/#origin-smart-contract).
    The result of this is that the message is added in the Origin's [Message Merkle Tree](glossary.md/#message-merkle-tree) and Chain A will have
    a new [state](glossary.md/#state) with a new [Message Merkle Root](glossary.md/#message-merkle-root).
    For this example, we will say that Chain A sends 4 messages to Chain B so we can illustrate what it's [Message Merkle Tree](glossary.md/#message-merkle-tree) looks like.
    This is what the message Merkle tree looks like, with the simplification of having a height of only 3 rather than 32:
    ![MessageMerkleRoot](../../static/img/MessageMerkleExampleForDeepDive.png 'Diagram of Message Merkle Root formed from Merkle Tree of Messages')
2.  On [Synapse Chain](glossary.md/#synapse-chain), the [Guard](glossary.md/#guard) will submit a [State Snapshot](glossary.md/#state-snapshot) that includes
    the new [state](glossary.md/#state) of Chain A. For this example, we will also claim that Chain B also has a [state](glossary.md/#state) included in the snapshot.
    This is what the [State Snapshot](glossary.md/#state-snapshot) would look like with the states from both chains:
    ![ExampleStateSnapshot](../../static/img/StateSnapshotExampleForDeepDive.png 'Diagram of Example of State Snapshot')
3.  On [Synapse Chain](glossary.md/#synapse-chain), either of the [Notaries](glossary.md/#notary) will submit a [State Snapshot](glossary.md/#state-snapshot) that includes
    the new [state](glossary.md/#state) of Chain A. This results in the [Synapse Chain](glossary.md/#synapse-chain) registering a
    new [Attestation](glossary.md/#attestation) that includes the [Snap Root](glossary.md/#snap-root) as well as the current [Agent Root](glossary.md/#agent-root).
    ![ExampleAttestation](../../static/img/ExampleAttestation.png 'Example Attestation')
4.  On Chain B, the [Notary](glossary.md/#notary) for Chain B proposes the [Attestation](glossary.md/#attestation).
5.  The system must wait for the [Optimistic Period](glossary.md/#optimistic-period) before executing the message. This is when
    the [Guard](glossary.md/#guard) has an opportunity to report [fraud](glossary.md/#fraud) if either the [Attestation](glossary.md/#attestation)
    was not registered on the [Synapse Chain](glossary.md/#synapse-chain) or if one of the [states](glossary.md/#state) included in the [State Snapshot](glossary.md/#state-snapshot)
    was not a valid state on the Origin chain.
6.  On Chain B after the [Optimistic Period](glossary.md/#optimistic-period) has passed without any [fraud reports](glossary.md/#fraud-report), the
    [Executor](glossary.md/#executor) will execute the message. This is done by a [Merkle Proof](glossary.md/#merkle-proof) for the [state](glossary.md/#state) against the [Snap Root](glossary.md/#snap-root)
    in the [Attestation](glossary.md/#attestation), and then with another [Merkle Proof](glossary.md/#merkle-proof) for the [message](glossary.md/#message) against the [Message Merkle Root](glossary.md/#message-merkle-root)
    in the [state](glossary.md/#state).

Below is a sequence diagram illustrating the steps involved in sending a message from Remote Chain A to Remote Chain B:

```mermaid
sequenceDiagram
    box On_Chain
    participant Synapse_Chain
    participant Chain_A
    participant Chain_B
    end
    box Off_Chain_Agents
    participant Guard_1
    participant Executor_1
    participant Notary_A
    participant Notary_B
    end
    Chain_A->>Chain_A: Client Smart Contract on Chain A calls "send message"<br/> on Chain A's Origin Smart Contract
    Chain_A->>Chain_A: Origin Smart Contract updates its<br/> "Message Merkle Root" and thus updates<br/> the state of Chain A.
    Guard_1->>Chain_A: Guard observes new state of Chain A<br/> after message was added.
    Guard_1->>Synapse_Chain: Guard posts a "State Snapshot"<br/> including new State of Chain A to the Synapse Chain.
    Notary_B->>Synapse_Chain: Any of the Notaries observes<br/> that there is a new State for Chain A registered on the<br/> Synapse Chain by a Guard.
    Notary_B->>Synapse_Chain: Any of the Notaries posts a "State Snapshot"<br/> including the new State for Chain A to the<br/> Synapse Chain.
    Synapse_Chain->>Synapse_Chain: Synapse Chain registers<br/> a new "Attestation" that can be<br/> used as a commitment <br/> to the message being sent<br/> from Chain A to Chain B. <br/>The Attestation is also a <br/>commitment to the <br/>Agent Set at that time.
    Notary_B->>Chain_B: Notary for Chain B proposes "Attestation".
    Chain_B->>Chain_B: Wait "optimistic period" before allowing<br/> "Attestation" to be used as<br/> proof to execute messages. <br/> Guards have an opportunity<br/> to report fraud during this time.
    Executor_1->>Chain_B: Executor executes the message using<br/> the Attestation as proof. <br/> Since the Optimistic Period<br/> passed without any fraud reported,<br/> the message is trusted.
    Chain_B->>Chain_B: A receipt is generated when the message<br/> is finally executed, which is used to give credit<br/> to the agents involved in sending the message.
    Notary_B->>Chain_B: The Notary observes the message receipt<br/> after the Executor executes the message.
    Notary_B->>Synapse_Chain: Notary signs the message receipt<br/> and submits it to the Synapse Chain.
    Synapse_Chain->>Synapse_Chain: Wait optimistic period to<br/> give Guards a chance to challenge<br/> the message receipt. <br/>After that, rewards are distributed<br/> to the agents involved<br/> in sending the message.
```

## Adding other Bonded Agents

The requirement to become a [Bonded Agent](glossary.md/#bonded-agent) is to post a [bond](glossary.md/#bond) on the [Bonding Manager Smart Contract](glossary.md/#bonding-manager-smart-contract) on the [Synapse Chain](glossary.md/#synapse-chain).
Upon adding the new [Bonded Agent](glossary.md/#bonded-agent) to the [Agent Set](glossary.md/#agent-set), the [Bonding Manager Smart Contract](glossary.md/#bonding-manager-smart-contract) on the [Synapse Chain](glossary.md/#synapse-chain) will
calculate a new [Agent Root](glossary.md/#agent-root).

Below is an illustration of how the [Agent Root](glossary.md/#agent-root) is calculated from the [Agent Set](glossary.md/#agent-set).

![AgentRoot](../../static/img/AgentRoot.png 'Diagram of AgentRoot formed from Merkle Tree of Agent Infos')

The tricky part is how to communicate this [Agent Root](glossary.md/#agent-root) to the [remote chains](glossary.md/#remote-chain).

Whenever a [Notary](glossary.md/#notary) submits a [State Snapshot](glossary.md/#state-snapshot) to the [Synapse Chain](glossary.md/#synapse-chain),
the Synapse Smart Contract that handles that transaction will register an [Attestation](glossary.md/#attestation)
that includes the [Snap Root](glossary.md/#snap-root) as well as the current [Agent Root](glossary.md/#agent-root).

Then, one of the [Notaries](glossary.md/#notary) for each [remote chain](glossary.md/#remote-chain) will take that
[Attestation](glossary.md/#attestation) and propose it to its [remote chain](glossary.md/#remote-chain).

Just as there is an [Optimistic Period](glossary.md/#optimistic-period) for each [message](glossary.md/#message), the [Agent Root](glossary.md/#agent-root)
also has an [Optimistic Period](glossary.md/#optimistic-period) defined by the system, during which any [Guard](glossary.md/#guard)
could call out [fraud](glossary.md/#fraud) if it believes the [Agent Root](glossary.md/#agent-root) is wrong.

If no guard submits a [fraud report](glossary.md/#fraud-report), the new [Agent Root](glossary.md/#agent-root) will become active
on the [remote chain](glossary.md/#remote-chain).

Whenever there is a change in the [Agent Root](glossary.md/#agent-root) on a [remote chain](glossary.md/#remote-chain),
each [Bonded Agent](glossary.md/#bonded-agent) must re-register themselves by providing a proof of inclusion.

Below is a sequence diagram illustrating the steps involved in adding another bonded agent (a new Guard):

```mermaid
sequenceDiagram
    box On_Chain
    participant Synapse_Chain
    participant Chain_A
    participant Chain_B
    end
    box Off_Chain_Agents
    participant Guard_1
    participant Guard_2
    participant Executor_1
    participant Notary_A
    participant Notary_B
    end
    Guard_2->>Synapse_Chain: New Bonded Agent (Guard_2) posts bond<br/> on Synapse Chain's Bonding Manager Smart Contract.
    Synapse_Chain->>Synapse_Chain: Bonding Manager on Synapse Chain<br/> updates its Agent Set,<br/> resulting in a new Agent Root.
    Notary_B->>Synapse_Chain: At any point in the future, whenever a<br/> Notary posts a "State Snapshot" to <br/> the Synapse Chain, the resulting<br/> "Attestation" will include the<br/> latest "Agent Root".
    Notary_A->>Chain_A: Notary for Chain A proposes Attestation on<br/> Chain A that includes the new Agent Root.
    Notary_B->>Chain_B: Notary for Chain B proposes Attestation on<br/> Chain B that includes the new Agent Root.
    Chain_A->>Chain_A: Wait optimistic period during which time any<br/> Guard can challenge new Agent Root on Chain A.
    Chain_B->>Chain_B: Wait optimistic period during which time any<br/> Guard can challenge new Agent Root on Chain B.
    Chain_A->>Chain_A: New Agent Root is accepted on Chain A.
    Chain_B->>Chain_B: New Agent Root is accepted on Chain B.
    Guard_1->>Chain_A: Guard_1 provides Chain A with proof of inclusion<br/> using the new Agent Root to re-register on Chain A.
    Guard_1->>Chain_B: Guard_1 provides Chain B with proof of inclusion<br/> using the new Agent Root to re-register on Chain B.
    Guard_2->>Chain_A: Guard_2 provides Chain A with proof of inclusion<br/> using the new Agent Root to register on Chain A.
    Guard_2->>Chain_B: Guard_2 provides Chain B with proof of inclusion<br/> using the new Agent Root to register on Chain B.
    Notary_A->>Chain_A: Notary for Chain A provides Chain A with proof <br/>of inclusion using the new Agent Root<br/> to re-register on Chain A.
    Notary_B->>Chain_B: Notary for Chain B provides Chain B with proof <br/>of inclusion using the new Agent Root<br/> to re-register on Chain B.
```
