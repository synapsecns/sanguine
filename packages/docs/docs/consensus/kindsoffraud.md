---
sidebar_position: 4
---

# Fraud in Synapse Messaging

The two kinds of bonded agents are Guards and Notaries. Notaries require a much larger stake due to the potential harm they could cause.

## How a Notary commits Fraud
1. Notary can sign and attest to a Merkle Root of States that never happened on the SYN chain.
This is easy to catch because a Guard can simply observe the attestation on the Receiving Chain and check if it has already been registered on the SYN chain.
2. Notary can sign and attest to a Merkle Root of States that is registered on the SYN chain,
but that Merkle Root includes an invalid state.
For this, the Guard needs to prove that the invalid state is part of that Merkle Tree of States
and then the Sending Chain will quickly know that the state is not valid and be able to slash the Notary that signed the attestation.
3. Note that for the Merkle Root of States to be registered on SYN chain,
there must have been a Guard who signed for the state followed by some
other Notary (could be same or could be different Notary) who also submitted that state to the SYN chain.
4. To Summarize, in the event that a Receiving chain has an bad attestation posted for a particular Merkle Root of States,
the Notary for the Receiving chain must have signed it and will get slashed.
If the bad Merkle Root of States is registered on the SYN chain, the Notary who submitted the list of states
that included that bad state will get slashed AND also the Guard who originally registered the state will get slashed.
5. Who receives the slashed bond? The Guard who reports each of the agents. Note there must be a fraud report for each agent being slashed.

## How a Guard commits Fraud
1. A Guard can submit an invalid state to the SYN chain.
2. A Guard can submit an invalid fraud report for the purposes of waging a denial of service attack.
3. A Guard has the power to put any attestation posted in dispute, and resolving the dispute requires cross-chain coordination and takes time.
4. While a Guard submitted a bad fraud report will eventually get slashed, if the Guard's stake is too small, it might fail to disincentive griefing from malicious Guards.
