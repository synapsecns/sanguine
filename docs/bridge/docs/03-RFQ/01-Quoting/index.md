---
title: Quoting
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
[Dispute Period]: /docs/RFQ/Security/#dispute-period
[Relaying]: /docs/RFQ/Relaying
[Proving]: /docs/RFQ/Proving
[Claiming]: /docs/RFQ/Claiming

[User]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

# Quoting

The Synapse RFQ systems allows [Quoter] entities (aka market makers / solvers / relayers) to post quotes via an off-chain [Quoter API]. These quotes are matched to `User` bridge inputs to achieve the optimal parameters (eg: the best price) for the [User]'s bridge transaction.

There are two types of quoting methods supported by the Synapse RFQ system:

##
### [Passive Quoting](/docs/Routers/RFQ/Quoter%20API/#passive-quotes)

Similar to an order book, Passive Quoting communicates a [Quoter]'s ongoing intention to fulfill any transaction that occurs upon specific routes and meets specific limits, pricing, and fee criteria.

### [Active Quoting](/docs/Routers/RFQ/Quoter%20API/#active-quotes)

Starting with [Fast Bridge V2](https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html), a new Active Quoting method has been introduced where a [Quoter] can listen and respond to live quote requests individually.

This supplements the existing Passive Quotes to create a hybrid system, where Active and Passive quoting can be utilized together by [Quoters](Quoter) in any desired combination to maximize their efficiency.

Active quoting is more complicated to implement and maintain, but allow for more granular & customized quotes that can improve efficiency among other benefits. Quoters who prefer a simpler approach are free to use nothing but Passive Quotes, if they choose.


##

Regardless of the method used, these quotes constitute a provisional commitment to fulfill the [User]'s bridge according to the quoted price and other parameters, once it is submitted on-chain.

To that end, integrators and users can utilize the data from these quotes to construct and submit a corresponding [bridge] transaction on-chain through the [Fast Bridge Contract](https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html). Once this transaction is finalized on-chain, `User`s can expect to receive their funds on the destination shortly after, as quoted.

[Quoters](Quoter) are responsible for keeping their quotes fresh and accurate. Likewise, they are responsible for completing their part of fulfillment for any transactions which act upon their quotes. To these effects, [Quoters](Quoter) should push updates as rapidly as possible in reaction to consequential changes in prices, balances, etc. By default, the [Canonical Relayer](../Relayer) continuously updates quotes by checking on-chain balances, in-flight requests, and gas prices - custom implementations should take a similar approach.

