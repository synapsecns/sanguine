---
sidebar_position: 2
---

# Running a Guard

## What is a Guard

Whereas Notaries are in a position to attack the system’s Integrity, Guards are in a position to attack the Liveness property of the system because they can report fraud that would prevent messages from being accepted. A malicious Guard therefore can perform a Denial of Service attack by claiming that an honest message is fraudulent, and thus causing delays while the system determines whether the fraud report is correct or not. We need to make sure the Guards are disincentivized from engaging in such Denial of Service attacks. Because of this, prior to becoming a Guard, you must post a stake and have that held in escrow by the Synapse Carbon system. The stake to become a Guard will be considerably less than that of a Notary due to the role each plays and the potential damage that could be inflicted by a Notary vs a Guard.

## Running a Guard

```
codeblock!
```
