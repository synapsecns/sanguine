---
sidebar_position: 3
---

# Running a Guard

Whereas Notaries are in a position to attack the system’s Integrity, Guards are in a position to attack the Liveness property of the system because they can report fraud that would prevent messages from being accepted. A malicious Guard therefore can perform a Denial of Service attack by claiming that an honest message is fraudulent, and thus causing delays while the system determines whether the fraud report is correct or not. We need to make sure the Guards are disincentivized from engaging in such Denial of Service attacks. Because of this, prior to becoming a Guard, you must post a stake and have that held in escrow by the Synapse Messaging System. The stake to become a Guard will be considerably less than that of a Notary due to the role each plays and the potential damage that could be inflicted by a Notary vs a Guard.

Unlike the Notary, the Guard can service all the chains, although it can be configured to only service a subset of the chains.
The Guard’s main job is to catch and report fraud, but it also has a less critical role in signing attestations so that it can earn tips during times when messages are just being sent routinely without fraud. Because of this, Guards should have an address used for paying Gas on all the chains it supports.
Like the Notary, there is another bonded Signing key that the Guard has. The Guard’s bonded Signing key should be kept secret so as to avoid having an attacker use it for getting the Guard slashed.
To run the Guard, you only need the executable written in Go, and provide the configuration file.

The reference implementation instructions for the Guard can be found here:

https://github.com/synapsecns/sanguine/blob/master/agents/agents/guard/cmd/cmd.md

The guard configuration file could look something like this:

```yaml

```

