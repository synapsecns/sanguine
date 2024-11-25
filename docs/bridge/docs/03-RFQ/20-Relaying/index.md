---
title: Relaying
---

<!-- Reference Links -->
[relay]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#relayv2
[prove]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#provev2
[dispute]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#dispute
[claim]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#claimv2
[cancel]: https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#cancelv2
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

# Relaying

In the Synapse RFQ System, Relayers fulfill the intent of [User] [bridge] transactions by providing the liquidity and executing the [relay] transaction on the destination chain .

:::info

If you are reading this documentation to become a participating Relayer, be sure to read all RFQ sections.

Also note - Currently, relaying involves actions which require explicit authorization.

If you are interested in participating as a Relayer, contact us.

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

There are two versions of the relay function in FastBridgeV2. Relayers can use whichever best suits their implementation.

<div style={{ marginLeft: '20px' }}>

<blockquote>
```solidity
    function relayV2(bytes memory request, address relayer) external payable;
```
This version allows an arbitrary `relayer` address to be supplied
</blockquote>
<br />
<blockquote>
```solidity
    function relay(bytes memory request) external payable;
```
This version will auto-assign the executing EOA (`msg.sender`) as the `relayer`
</blockquote>
</div>

Regardless of the method used, a [BridgeRelayed](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgedepositrefunded) event will be emitted.

### Setting the `relayer` parameter
The address which is specified as the `relayer` on the [relay] will have control of the reimbursed funds when the [claim] occurs later.

Note that Relayers can utilize this feature to be reimbursed on different addresses than they are actually relaying from. This can be useful for advanced relaying setups, but is not a necessity.

# Exclusivity

As of FastBridgeV2, it is possible for integrators to optionally assign temporary exclusive fill rights to certain relayers.

IE: For a temporary period of time, only the relayer chosen and assigned by the integrator will be able to execute the relay.

For details on Exclusivity and how to relay these types of bridges, see [Exclusivity]

# Multicalling

As of FastBridgeV2, it is possible to batch many [relay] transactions together with [Multicall]

However, the Multicall implementation is limited to non-payable transactions only, so native-gas bridges cannot be included in batches.

### Permissions

Although relaying and claiming can be performed permissionlessly, in the current system Relayers will need to also operate a permissioned [Prover] role.

Note that this allows the use of different EOAs to [relay], [prove], and [claim] - which we recommend doing.

We also recommend that Relayers operate a [Quoter] to compete on pricing and routes, but this is not a necessity.

## Next steps

After a [relay] has been completed, the process is fully complete from the perspective of the [User].

Next, the [Relayer] must proceed to [Proving] the relay so that they can be reimbursed for the liquidity they provided.
