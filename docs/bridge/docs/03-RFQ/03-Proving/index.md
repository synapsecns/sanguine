---
title: Proving
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
[Multicall]: https://vercel-rfq-docs.vercel.app/contracts/utils/MulticallTarget.sol/abstract.MulticallTarget.html

[Quoter API]: /docs/Routers/RFQ/Quoter%20API/
[Dispute Period]: /docs/RFQ/Security/#dispute-period
[Relaying]: /docs/RFQ/Relaying
[Proving]: /docs/RFQ/Proving
[Claiming]: /docs/RFQ/Claiming

[User]: /docs/RFQ/#entities
[Quoter]: /docs/RFQ/#entities
[Prover]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

After [Relaying], [prove] transactions are executed by authorized [Provers](Prover) (who are typically also the [Relayer])

Through these transactions, the Prover is asserting that the indicated `relayer` completed a [relay] and can rightfully [claim] the escrowed [bridge] funds as a reimbursement.

Each [prove] transaction sets the [proof] data for the bridge and initiates a [Dispute Period].

### Function Options

There are two overloaded versions of the prove function in FastBridgeV2. Relayers can use whichever best suits their implementation.

<div style={{ marginLeft: '20px' }}>
1)
```solidity
    function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) external;
```
This version allows an arbitrary `relayer` address to be supplied

2)
```solidity
    function prove(bytes memory request, bytes32 destTxHash) external;
```
This version will auto-assign the executing EOA (`msg.sender`) as the `relayer`
</div>

Regardless of the method used, a [BridgeProofProvided](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofprovided) event will be emitted.

Note that during the [Dispute Period], the `relayer` indicated on this function will be required to match the `relayer` on the actual [relay].

### Multicalling

As of FastBridgeV2, it is possible to batch many [prove] transactions together with [Multicall]


## Next steps

Following the [prove] transation, a [Dispute Period] begins - after which the [Relayer] may proceed to [Claiming]

Although disputes are unlikely, monitoring should occur during the [Dispute Period] to verify there are no issues.
