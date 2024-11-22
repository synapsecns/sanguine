---
title: Exclusivity
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

[User]: /docs/RFQ/#entities
[Quoter]: /docs/RFQ/#entities
[Prover]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

Starting with FastBridgeV2, specific Relayers can optionally be granted temporary exclusive rights to be the first to complete a [relay] for a given bridge.

Utilizing exclusivity minimizes the wasteful aspects of competing relayers "racing" to be the first to fill, allowing for more efficient operations & pricing.


Exclusivity is enforced by the FastBridge contract according to two values that will be on the [BridgeTransactionV2] struct emitted on the [BridgeRequested] event.


- exclusivityRelayer - This will be set to the address of the [Relayer] who has been granted first-fill rights.

- exclusivityEndTime - Once the Block timestamp exceeds this value, anyone can complete the relay if it has not yet been completed.

