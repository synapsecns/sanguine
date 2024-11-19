---
title: Claiming
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

After [Proving] and waiting for the [Dispute Period] to end, [Relayers] can then execute a [claim] transaction which will release the funds which have been escrowed since the original [bridge].

The funds will be transferred to the rightful [Relayer] as a reimbursement for the liquidity they provided on the [relay].


### Function Options

There are two overloaded versions of the claim function in FastBridgeV2. Relayers can use whichever best suits their implementation.

<div style={{ marginLeft: '20px' }}>
1)
```solidity
    function claim(bytes memory request, address to) external;
```
This version can only be executed by the `relayer` address on the proof and allows an arbitrary `to` address to be the recipient of the funds.

2)
```solidity
    function claim(bytes memory request) external;
```
This version can be executed permissionlessly, and will transfer the funds only to the `relayer` address on the proof.
</div>

Regardless of the method used, a [BridgeDepositClaimed](hhttps://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgedepositclaimed) event will be emitted.

### Multicalling

As of FastBridgeV2, it is possible to batch many [claim] transactions together with [Multicall]


## Next steps

Following the [claim] transation, the bridge deposit funds have been reimbursed to the [Relayer] and are ready to be used for another relay. The bridge cycle is fully complete.
