---
title: Security and Risk Factors
---

<!-- Reference Links -->
[bridge]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridge
[relay]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#relay
[prove]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#prove
[dispute]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#dispute
[claim]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#claim
[cancel]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#cancel
[proof]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetxdetails
[BridgeRequested]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested
[BridgeTransactionV2]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgetransactionv2
[BridgeRelayed]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerelayed
[BridgeProofProvided]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofprovided
[Cancel Delay]: https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#refund_delay
[Quoter API]: /docs/Routers/RFQ/Quoter%20API/
[User]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

[Dispute Period]: https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#dispute_period

# Dispute Period

The RFQ system optimistically assumes all [prove] transactions are accurate unless they are disputed. These disputes can occur during the [Dispute Period], where [Guard] entities may execute a [dispute] transaction if they detect errors or potentially fraudulent activity, such as incorrect fill amounts or proofs submitted by the wrong [Relayer].

After the dispute period has passed, the funds in escrow from the original [bridge] transaction can be released via a [claim] transaction, which will reimburse the rightful [Relayer].

Currently, [Guard] entities are operated solely by Synapse itself and monitor all FastBridge activity across all deployments.


# Self-Monitoring

If you are a participating [Relayer], it is highly recommended that you monitor proofs, disputes, and other critical functions that involve your quotes and transactions.

Refer to the [RFQ Indexer] for endpoints that can aid with monitoring and alerting.

# Relaying Risk Factors

### Missing Relays

If a [bridge] transaction occurs that acts upon a quote issued by a particular [Relayer], and that [Relayer] fails to complete the [relay] transaction - the [bridge] deposit can be come "stuck" if there is no one to complete it.

Although funds are safe in these instances and will eventually either be cancelled & refunded or completed late, these are adverse UX incidents nonetheless.

All participating [Relayer]s must ensure they are able to complete all [bridge] transactions that act upon their quotes and will be asked to quickly rectify any shortfalls.

Refer to the [RFQ Indexer] for endpoints that can aid with tracking missing relays.

### Missing Proofs / Claims

Currently, it is the responsibility of the [Relayer] to track and submit timely [prove] transactions for their own finalized [relay] transactions.
Failure to do so can result in a delayed reimbursement at best, or an eventual loss of funds at worst.

Likewise, [Relayer]s are expected to track and submit their own [claim] transactions once their [proofs] are past the [Dispute Period].
Missing claims are less urgent than missing proofs, but the funds can remain unclaimed in escrow indefinitely if the responsible [Relayer] loses track of them.

Refer to the [RFQ Indexer] for endpoints that can aid with tracking missing proofs and claims.

### Invalid Relays

As is the case with many cross-chain optimistic systems, there are risks involved with providing liquidity as a [Relayer].

If a [relay] occurs for an incorrect output amount (or any other incorrect [BridgeTransactionV2] data) and is submitted via [prove] as the fulfillment a [bridge] transaction, the [proof] will be disputed by a [Guard]. The dispute action will effectively ignore & erase the [proof] and allow for a new correct [proof] to be submitted.

In other terms, although a [relay] transaction may have truly taken place and delivered funds from the [Relayer] to the [User], if the [relay] request parameter does not perfectly match the encoded [BridgeTransactionV2] data on the [bridge] transaction, then it cannot be accepted as the completion of that bridge.

The incorrect [relay] in this example cannot be reversed or reimbursed and would effectively be a loss of funds for the [Relayer] who performed it.

Reorganizations of the [bridge] transaction on the origin chain are the most likely vector for this scenario.

:::note

<span style={{ color: 'red' }}>**IMPORTANT**</span>

It is the sole responsibility and liability of the [Relayer] to fully understand and manage the risks involved with participating as a [Relayer] in the Synapse RFQ system, including the risk of invalid relays.

:::