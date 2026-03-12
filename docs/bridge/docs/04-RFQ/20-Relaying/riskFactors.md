---
title: Risk Factors
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


:::warning

As is the case with many cross-chain systems, there are inherent risks involved with providing liquidity as a [Relayer].

Although the system is designed to minimize these risks, it is ultimately the sole responsibility and liability of the [Relayer] to fully understand and manage the risks involved with participating as a [Relayer] in the Synapse Intent Network (SIN) system.

:::

### Missing Relays

If a bridge transaction occurs that acts upon a quote issued by a particular [Relayer], and that [Relayer] fails to complete the [relay] transaction - the bridge deposit can be come "stuck" if there is no one to complete it.

Although funds are safe in these instances and will eventually either be cancelled & refunded or completed late, these are adverse UX incidents nonetheless.

All participating Relayers must ensure they are able to complete all bridge transactions that act upon their quotes and will be asked to quickly rectify any shortfalls.

This may include completing the [relay] at an unexpectedly unfavorable price - within a reasonable threshold. For example, if the gas price sharply increased beyond what was anticipated.

Refer to the [RFQ Indexer](/docs/Services/RFQ-Indexer-API) for endpoints that can aid with tracking missing relays.

### Missing Proofs / Claims

Currently, it is the responsibility of the [Relayer] to track and submit timely [prove] transactions for their own finalized [relay] transactions.
Failure to do so can result in a delayed reimbursement at best, or an eventual loss of funds at worst.

Likewise, Relayers are expected to track and submit their own [claim] transactions once their proofs are past the [Dispute Period].
Missing claims are less urgent than missing proofs, but the funds can remain unclaimed in escrow indefinitely if the responsible [Relayer] loses track of them.

Refer to the [RFQ Indexer](/docs/Services/RFQ-Indexer-API) for endpoints that can aid with tracking missing proofs and claims.

### Invalid Relays

If a [relay] occurs for an incorrect output amount (or any other incorrect [BridgeTransactionV2] data) and is submitted via [prove] as the fulfillment a bridge transaction, the [proof] will be disputed by a [Guard].

In other terms, although a [relay] transaction may have truly taken place and delivered funds from the [Relayer] to the [User], if the [relay] request parameter does not exactly match the encoded [BridgeTransactionV2] data on the bridge transaction, then it cannot be accepted as the completion of that bridge.

The incorrect [relay] in this example cannot be reversed or reimbursed and would effectively be a loss of funds for the [Relayer] who performed it.

Reorganizations of the bridge transaction on the origin chain are the most likely vector for this scenario.

### Pricing Failures

Relayers should never assume that all bridges are profitable or properly priced, even those that came from the Synapse Frontend, and instead should independently verify using trusted mechanisms and oracles.

Similarly, it is important to remember that bridge is a permissionless function. As with any similar system, bad actors *can* and *will* submit exploitative bridges to see if any Relayers make the fatal mistake of filling them.

Failure to implement appropriate pricing protections could result in a loss of Relayer funds.

### Self-Monitoring

If you are a participating [Relayer], it is highly recommended that you monitor proofs, disputes, and other critical functions that involve your quotes and transactions.

Refer to the [RFQ Indexer](/docs/Services/RFQ-Indexer-API) for endpoints that can aid with monitoring and alerting.

