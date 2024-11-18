---
title: Relaying
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
[Dispute Period]: https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#dispute_period
[Cancel Delay]: https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#refund_delay
[Quoter API]: /docs/Routers/RFQ/Quoter%20API/
[User]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

# Relaying

In the Synapse RFQ System, [Relayer]s fulfill the intent of [User] [bridge] transactions by providing the liquidity and executing the [relay] transaction on the destination chain .

A valid relay will start by observing a [BridgeRequested] event on an origin chain, which will emit the `request` bytes of an encoded [BridgeTransactionV2] struct, `destChainId`, and other values.

To complete the relay, these same request bytes should be provided to the `request` parameter of the [relay] function and executed on the FastBridge contract of the destination chain (`destChainId`)

If the `destToken` is the native currency of the chain (`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`) then that amount must also be set as the `msg.value` (IE: Not wrapped gas, such as WETH)

If the `destToken` is some other address, this would represent an ERC20 on the destination Chain.
Prior to calling [relay], the [Relayer] must have already granted sufficient token approvals to FastBridge to allow transfers of this asset.

### Contract Functions

There are two overloaded versions of the relay function in FastBridge, [one](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#relay) which allows you to supply an arbitrary `relayer` address, and [another](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#relay)  which assigns the executing EOA as the `relayer`.

The address specified as the `relayer` on your [relay] will ultimately be reimbursed the funds when the [claim] occurs.
Note that you can utilize this feature to be reimbursed on different addresses than you are relaying from - as needed.

### Permissions

Although relaying and claiming can be performed permissionlessly, in the current system [Relayer]s will need to also operate a permissioned [Prover] role.
Note that this allows the use of different EOAs to [relay], [prove], and [claim] - which we recommend doing.
We also recommend that [Relayer]s operate a [Quoter] to compete on pricing and routes, but this is not a necessity.



Once the [User] has signed and submitted their deposit on-chain via a [bridge] transaction, a [BridgeRequested] event will be emitted.

[Relayer]s who observe this event can permissionlessly complete the bridge by calling [relay] on the FastBridge contract of the destination chain.

The [BridgeRequested] event emission includes the `request` bytes of an encoded [BridgeTransactionV2] struct. The [Relayer] must execute their [relay] with these exact bytes from the origin chain, or their relay will not be considered valid and could result in a loss of relayer funds. To reiterate, the destination contract must operate optimistically upon the [Relayer]'s supplied parameters - it has no capability to verify or protect the [Relayer] from loss of funds due to error or ReOrgs on the origin chain. All relays are final and all participating [Relayer]s must assume the full responsibility and risk of invalid relays.

Assuming the [Relayer] has sufficient funds and necessary ERC20 approvals on the destination chain, the [relay] transaction will facilitate the delivery of funds from the [Relayer] to the [User] via the FastBridge contract as intermediary. This will also emit a [BridgeRelayed] event and prevent any further attempts to relay that transaction.


At this point, the process is complete from the [User]'s perspective. However, the [Relayer] needs to be reimbursed for the funds they delivered.

