---
title: Bridging
---

<!-- Reference Links -->
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
[Cancel Delay]: https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#refund_delay
[Multicall]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IMulticallTarget.sol/interface.IMulticallTarget.html

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


Once a quote has been obtained via [Quoting], the details of the quote can be used to construct a bridge transaction for the user to sign and submit to the origin chain.

Bridges through Synapse RFQ utilize the [Synapse Router](/docs/Routers/Synapse-Router/) - Refer to those docs for more detail.


:::info
If you are interested in integrating with Synapse RFQ Bridging, refer to the [Synapse Bridge SDK](/docs/Bridge/SDK).

Alternatively, you can explore the [Bridge REST API](https://api.synapseprotocol.com/api-docs/).

:::


## Exclusivity

As of FastBridgeV2, it is possible for integrators to optionally assign temporary exclusive fill rights to certain relayers.

IE: For a temporary period of time, only the relayer chosen and assigned by the integrator will be able to execute the relay.

For details on Exclusivity and how to create these types of bridges, see [Exclusivity]


## Effects of a bridge Transaction

If sufficient funds and approvals exist, the bridging funds will be transferred from the [User] to the FastBridge contract.

The funds will remain with the contract in escrow until:

- [Claiming] occurs, which transfers the funds to the [Relayer] as reimbursement for completing the relay on the destination.

or

- [Canceling] occurs, which returns the funds to the `originSender`.


Additionally, a [BridgeRequested] event will be emitted which contains all instructions for the bridge to be completed by Relayers



## Next steps

Relayers will observe the bridge transaction via the [BridgeRequested] event and proceed to [Relaying] if it meets all of their criteria.
