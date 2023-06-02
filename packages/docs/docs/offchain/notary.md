---
sidebar_position: 2
---

# Running a Notary

The potential for fraud by a malicious Notary is enormous and there needs to be a HUGE penalty for any Notary who proposes a fraudulent message to a Destination chain. We described how a single honest Guard can report a Notary for fraud, but this means nothing if the Notary doesn’t receive punishment for the fraud. Thus, prior to acting as a Notary, you must post a sizable stake and have it held in escrow by the Synapse Messaging System. A Notary who commits fraud will lose the escrowed stake, so this stake is what keeps Notaries honest.

Each Notary can actually only attest to messages destined to a particular chain, so if you want to run a Notary for more than one destination chain, you will need to run separate Notaries for each.
Every Notary must post a bond prior to running, and it will need to designate an address to act as the Signing key for the very important attestations that it needs to sign. This should be a different address from the one that pays gas, and there should be great care to keep the bonded signing key safe. If an adversary took possession of that key, they could commit fraud with it just for the purposes of getting that Notary slashed. Usually the Guard who reports the Notary receives the slashed bond, so this is one way an attacker could obtain the bonded amount for itself.

The Synapse team can provide a reference implementation and deployment instructions if you would like to run a Notary.
To run the Notary, you only need the executable written in Go, and provide the configuration file.

