---
sidebar_position: 6
---

# Rewarding Agents

As mentioned there are 3 kinds of agents:
1. Notaries who post a large bond and are assigned to one Receiving chain.
2. Guards who post a smaller bond and can report fraud and also submit valid Origin states to the SYN chain that can be used by Notaries when attesting to a list of states.
3. Executors actually execute messages after the optimistic period has passed for an attested state that includes that message in the Merkle Tree of Messages.

## Working for tips.
When a client wants to send a message, it must obviously pay the gas for performing the transaction on the Sending chain.
<br/>
However, for the agents to be incentivized to do any work to move that message to the Receiving Chain, the sender needs to provide enough money to pay for the gas on the other chains and also an additional amount to incentivize the agents to do the work.
### Where do we pay Gas?
1. The Sending Chain will obviously cost a certain amount of Gas.
2. A Guard then needs to pay Gas to submit a list of states to SYN chain, although a given state can be amortized for all the messages that came before.
3. A Notary also needs to submit a list of states on the SYN chain, again amortized for all messages coming before.
4. For each Receiving chain that has a message sent to it, a Notary assigned to that chain must submit an attesation for a Merkle Root of States which can be amortized for all messages prior.
5. For each message to be executed, the Executor needs to pay Gas on the Receiving chain for every message exucuted.

### Who pays the Gas?
The Sender of the message must pay gas, but how does this work in terms of paying for gas on the SYN chain and the Receiving chain?
<br/>
<br/>
The next section will discuss how Synapse uses "Gas Oracles" to estimate the amount of Gas needed to perform the work on the other chains so that the gas money plus the extra tips can be collected at the time the message is Sent on the Sending chain.

### Where are rewards paid?
Note the tips will be given to the agents on the SYN chain. This means that periodically, the gas collected on the Origin chains will be bridged to the SYN chain because the sender pays on the Sending chain.
<br/>
<br/>
When a message is executed on th Destination chain, all the agents who had a hand in delivering that message are elligible to receive a reward that should cover the gas as well as additional rewards to keep them incentivized.
<br/>
Upon having the message executed, the Destination chain produces a receipt of the message being executed.
The Notary will need to take that receipt, sign it and submit on the SYN chain in order for it and the other agents to receive tips.
<br/>
Note that this is yet another potential for fraud from the Notary, because the receipt could in theory be fraudulent, but the Notary would be risking a very large stake for a small amount of tips.
This is another payload that a Guard can potentially look for fraud and report.
