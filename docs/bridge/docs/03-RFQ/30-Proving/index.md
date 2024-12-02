---
title: Proving
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

After [Relaying] successfully and observing that the [relay] transaction has finalized on the destination chain, [prove] transactions are executed by authorized Provers (who are typically also the [Relayer])

Through these transactions, the Prover is asserting that the indicated `relayer` completed a [relay] and can rightfully [claim] the escrowed bridge funds as a reimbursement.

Each [prove] transaction sets the [proof] data for the bridge and initiates a [Dispute Period].

### Function Options

There are two versions of the prove function in FastBridgeV2. Relayers can use whichever best suits their implementation.

<div style={{ marginLeft: '20px' }}>
<blockquote>
```solidity
    function proveV2(bytes32 transactionId, bytes32 destTxHash, address relayer) external;
```
This version allows an arbitrary `relayer` address to be supplied
</blockquote>

<br />
<blockquote>
```solidity
    function prove(bytes memory request, bytes32 destTxHash) external;
```
This version will auto-assign the executing EOA (`msg.sender`) as the `relayer`
</blockquote>
</div>

Regardless of the method used, a [BridgeProofProvided](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeproofprovided) event will be emitted.

Note that during the [Dispute Period], the `relayer` indicated on this function will be required to match the `relayer` on the actual [relay].

:::warning

prove should not be called until the Relayer is confident that the relay transaction is finalized and will not be reorganized.

:::

### Multicalling

As of FastBridgeV2, it is possible to batch many [prove] transactions together with [Multicall]


## Next steps

Following the [prove] transaction, a [Dispute Period] begins - after which the [Relayer] may proceed to [Claiming]

Although disputes are unlikely, monitoring should occur during the [Dispute Period] to verify there are no issues.
