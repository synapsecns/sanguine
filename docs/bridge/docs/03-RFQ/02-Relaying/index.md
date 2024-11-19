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

# Relaying

In the Synapse RFQ System, [Relayers](Relayer) fulfill the intent of [User] [bridge] transactions by providing the liquidity and executing the [relay] transaction on the destination chain .

:::note

If you are interested in participating as a Relayer, it is important to read all sections of the RFQ documentation

:::

## Detecting a Bridge Request

A relay will start by observing a [BridgeRequested] event on an origin chain, which will emit the `request` bytes of an encoded [BridgeTransactionV2] struct, `destChainId`, and other values.

These are the bridge instructions that the [Relayer] is *relaying* to the FastBridge contract on the indicated destination chain (`destChainId`) - which will then utilize the [Relayer]'s liquidity to complete the bridge.

## Executing the [relay] function

To complete the relay, the [Relayer] should provide these same `request` bytes emitted by [BridgeRequested] to the `request` parameter of the [relay] function on the FastBridge contract of the destination chain.

Additionally , if the `destToken` is the native currency of the chain (indicated by placeholder `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`) then that amount must also be set as the `msg.value`

If the `destToken` is some other address, this would indicate that the destination asset to deliver is an ERC20.
Prior to calling [relay], the [Relayer] must have already granted sufficient token approvals to FastBridge to allow transfers of this ERC20.

If the [Relayer] has sufficient funds and approvals, and the relay has not already been completed by another relayer, FastBridge will then facilitate the delivery of the funds to the [User] and emit a [BridgeRelayed] event.

### Function Options

There are two overloaded versions of the relay function in FastBridge:

<div style={{ marginLeft: '20px' }}>
1)
```solidity
    function relay(bytes memory request, address relayer) external payable;
```
This version allows arbitrary `relayer` address to be supplied

2)
```solidity
    function relay(bytes memory request) external payable;
```
This version will auto-assign the executing EOA (`msg.sender`) as the `relayer`
</div>

### Setting the `relayer` parameter
The address which is specified as the `relayer` on the [relay] will ultimately be reimbursed the funds when the [claim] occurs later.

Note that [Relayer]s can utilize this feature to be reimbursed on different addresses than they are actually relaying from. This can be useful for advanced relaying setups, but is not a necessity.

### Permissions

Although relaying and claiming can be performed permissionlessly, in the current system [Relayer]s will need to also operate a permissioned [Prover] role.

Note that this allows the use of different EOAs to [relay], [prove], and [claim] - which we recommend doing.

We also recommend that [Relayer]s operate a [Quoter] to compete on pricing and routes, but this is not a necessity.

## Next steps

After a [relay] has been completed, the next step is [Proving] the relay.
