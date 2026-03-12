---
title: Exclusivity
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

Starting with FastBridgeV2, specific Relayers can optionally be granted temporary exclusive rights to be the first to complete a [relay] for a given bridge.


## Benefits of Exclusivity

Utilizing exclusivity minimizes the wasteful aspects of competing relayers "racing" on the open-market to be the first to fill, allowing them to offer more efficient operations & pricing.


## How to Decide Exclusivity

Exclusivity settings are an optional tool offered to integrators and users - the functionality is available to be used in any way that suits them. For example, exclusivity can be used to implement an off-chain auction system, or to distribute relays to a private liquidity network.


## How to Set Exclusivity

Exclusivity can be activated and configured per bridge. Users and Integrators simply need to determine the following values off-chain and assign them as part of [BridgeParamsV2](https://rfq-contracts.synapseprotocol.com/contracts/interfaces/IFastBridgeV2.sol/interface.IFastBridgeV2.html#bridgeparamsv2) when the bridge deposit transaction is constructed.

| Parameter                  | Description                                                                                           |
|----------------------------|-------------------------------------------------------------------------------------------------------|
| quoteRelayer             | The address of the Relayer who will have temporarily exclusive fill rights.                           |
| quoteExclusivitySeconds  | The duration in seconds for which the `quoteRelayer` will have exclusive fill rights. The countdown begins when the bridge lands on the origin chain and is based on the transaction's block timestamp. |

Note: The seconds setting can be as granular and fine-tuned as desired. However, for simple implementations, thirty seconds is more than enough time for most relays.


## Relaying an Exclusive Fill

Exclusivity is enforced by the FastBridge contract according to two values that will be on the [BridgeTransactionV2] struct emitted on the [BridgeRequested] event.

| Parameter            | Description                                                                                           |
|----------------------|-------------------------------------------------------------------------------------------------------|
| exclusivityRelayer   | This will be set to the address of the [Relayer] who has been granted first-fill rights.              |
| exclusivityEndTime   | Once the Block timestamp exceeds this value, anyone can complete the relay if it has not yet been completed. |

Relayers who are *not* the assigned `exclusivityRelayer` can wait until the block timestamp of the destination has surpased the `exclusivityEndTime` and then attempt the relay for themselves, if it was not filled in time by the assignee.
