---
title: Miscellaneous
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

## FastBridge Contract Utilities


### Call from Origin Chain:
<blockquote>
#### [canClaim](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#canclaim)

Checks if the [Dispute Period] has passed and the funds are now claimable

##### [bridgeStatuses](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgestatuses)

Returns the current [BridgeStatus](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgestatus) of the indicated `transactionId`

##### [bridgeProofs](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgeproofs)

Returns the details of a bridge proof, including the asserted `relayer`

#### [getBridgeTransaction](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#getbridgetransactionv2)

Decodes request bytes (such as those emitted from a [BridgeRequested] event) into a [BridgeTransactionV2] struct.
</blockquote>

### Call from Destination Chain:
<blockquote>
##### [bridgeRelays](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgerelays)

Checks if the indicated `transactionId` has already been relayed
</blockquote>

## RFQ Load Tester

The [`rfq-loadtest`](https://github.com/synapsecns/sanguine/tree/master/packages/rfq-loadtest) tool can be used to rapidly send ETH bridges for the purpose of load testing.
