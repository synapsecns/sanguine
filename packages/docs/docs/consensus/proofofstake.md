---
sidebar_position: 5
---

# Proving Bonded Agents

Thus far, we have glossed over the very important detail of how do we know a Notary is a Notary and a Guard is Guard.
We mentioned that a Notary posts a LARGE stake and a Guard posts a smaller stake that is large enough to disincentivize griefing.
<br/>
This section will explain the details of how the stake is collected and how slashing occurs.
Additionally, we need to talk about how the list of Notaries and Guards are kept in sync across all the remote chains.

## SYN Chain as Central Hub
The SYN chain is where each agent must post bond which will be held in Escrow by a special Smart Contract deployed there.
The name of this smart contract on SYN chain is called the Bonding Manager. Upon posting a bond, the agent will be registered as either a Guard or Notary.
Each Notary is assigned to a specific chain, while each Guard can service all the chains.
<br/>
These are the steps of becoming an agent.
1. Bond is posted to the BondingManager smart contract on SYN chain along with what type of agent it will be.
2. Upon receiving the bond, the BondingManager adds the new agent to the list of agents.
3. The list of agents if Merklized and a Merkle Root of Agents is registered.
4. The 32 bytes Merkle Root of Agents is now the current latest state.
5. Note that upon receiving the stake, the BondingManager on SYN chain immediately knows who is a Notary and who is a Guard.
6. The challenge now is communicating this to the remote chains.

### Updating the Remote Chains with new Agent Set
1. The Merkle Root of Agents will be stored with every attestation from the previous section where we produce Merkle Roots of Origin States.
2. Thus, any time a Notary posts an attestation for a Merkle Root of Origin States on its particular Receiving chain, it will also attest to the Merkle Root of Agents at the time the Merkle Root of States was posted on SYN chain.
3. Just as there is an optimistic period before the Merkle Root of States can be used to execute a message, there will be an optimistic period before an agent can prove itself as a member of the latest Merkle Root of Agents.
4. After the optimistic period, the Remote Chain will allow any agent to prove itself as being part of the set of agents.
5. Once a new Merkle Root of Agents is accepted by a remote chain, any agent must re-register itself against that root in order for the remote chain to trust it.

### What happens during a fraud dispute?
1. If a Guard accuses a Notary of fraud, the remote chain will temporarily put both the Notary and Guard on probation while the dispute is resolved.
2. For the dipsute to be resolved, decisions made by other chains will need to be submitted to the remote chain that is in dispute.
3. The message about who was right and who was wrong will arrive similar to any other message, which means there needs to be a non-disputed Guard and Notary able to send messages.
4. Once the dispute is resolved, either the Guard or Notary will have been removed from the list of valid agents and their stake will go to the Guard reporting the fraud.
