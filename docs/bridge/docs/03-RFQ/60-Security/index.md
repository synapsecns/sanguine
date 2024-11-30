---
title: Security
---

<!-- Reference Links -->
[relay]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#relayv2
[prove]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#provev2
[dispute]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#dispute
[claim]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#claimv2
[cancel]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#cancelv2
[proof]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetxdetails
[BridgeRequested]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested
[BridgeTransactionV2]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetransactionv2
[BridgeRelayed]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerelayed
[BridgeProofProvided]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofprovided
[Cancel Delay]: https://rfq-contracts.synapseprotocol.com/contracts/FastBridge.sol/contract.FastBridge.html#refund_delay
[Multicall]: https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IMulticallTarget.sol/interface.IMulticallTarget.html

[Quoter API]: /docs/RFQ/Quoting/Quoter%20API/
[Dispute Period]: /docs/RFQ/Security/#dispute-period
[Quoting]: /docs/RFQ/Quoting
[Bridging]: /docs/RFQ/Bridging
[Relaying]: /docs/RFQ/Relaying
[Proving]: /docs/RFQ/Proving
[Claiming]: /docs/RFQ/Claiming
[Canceling]: /docs/RFQ/Canceling
[Security]: /docs/RFQ/Security
[Exclusivity]: /docs/RFQ/Exclusivity

[User]: /docs/RFQ/#entities
[Quoter]: /docs/RFQ/#entities
[Prover]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities


Synapse RFQ is an optimistic cross-chain system. This means that any ambiguous actions in the system are assumed to be accurate and honest by default unless they are challenged/disputed within a short timeframe.


### Proofs

With this system in particular, [prove] transactions are the focal point of this optimistic mechanism - whereby a [Relayer] is asserting that they completed a [relay] and can rightfully [claim] the escrowed bridge funds as a reimbursement.

Each [prove] transaction sets the [proof] data for the bridge and initiates a dispute period.


### Dispute Period

After a [prove] transaction is posted and the [proof] data is set, a window of time called the [Dispute Period](https://rfq-contracts.synapseprotocol.com/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#dispute_period) begins.

During this time, the prove/proof is eligible to be dispuated by [Guard] entities.

After the [Dispute Period](https://rfq-contracts.synapseprotocol.com/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#dispute_period) has passed without any disputes, the funds in escrow from the original bridge transaction can be released via a [claim] transaction, which will reimburse the rightful [Relayer].


### Guards

During the dispute period, off-chain [Guard] entities provide the security function of evaluating the [relay] asserted by the [prove]:

- Does the asserted [relay] transaction exist in a finalized state on the destination chain?

- Do all [BridgeTransactionV2] parameters of the destination [relay] match the origin bridge?

- Is the `relayer` asserted by the [prove] the same as the `relayer` of the [relay]?

If any discrepancies are found, the [Guard] will execute a [dispute].


### Effects of a [dispute] Transaction

When a [dispute] is executed, it effectively negates/erases the [proof] which was asserted by the disputed [prove] transaction.

This allows for a new corrected [prove] to be submitted and the process begins again.

If a [relay] truly did occur for the disputed [prove], but it was not for the correct bridge parameters, this constitutes an [invalid relay](/docs/RFQ/Relaying/riskFactors#invalid-relays).

Additionally, a [BridgeProofDisputed](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofdisputed) event will be emitted.

This event can be useful for monitoring / alerting.
