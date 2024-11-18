---
title: RFQ
---

import { RFQFlow } from '@site/src/components/RFQFlow'

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
[CANCEL_DELAY]: https://vercel-rfq-docs.vercel.app/contracts/FastBridgeV2.sol/contract.FastBridgeV2.html#refund_delay
[Quoter API]: /docs/Routers/RFQ/Quoter%20API/
[User]: /docs/RFQ/#entities
[Relayer]: /docs/RFQ/#entities
[Guard]: /docs/RFQ/#entities
[Canceler]: /docs/RFQ/#entities

# Synapse RFQ

Synapse RFQ (Request-For-Quote) is an <abbr title="'Intent' refers to a user authorizing specific actions that they want to achieve, typically in very simple terms, such as a bridge or swap. Actual execution of the actions is then performed on the user's behalf by third parties known as solvers/relayers.">intent</abbr>-based bridging system that connects briding users to a network of relayers.

These relayers compete to provide the optimal bridge execution (eg: the best price) for the user's specific request.

<h2 style={{ textAlign: 'center' }}>Request-For-Quote End-to-End Flow</h2>
<figure>
    <RFQFlow />
</figure>
<br />
<div style={{ display: 'flex', justifyContent: 'center' }}>
  <table>
    <tbody>
      <tr>
        <td><strong>Order</strong></td>
        <td>User inputs their desired order into a bridge interface.</td>
      </tr>
      <tr>
        <td><strong>Quote</strong></td>
        <td>Quotes from [Relayer]s are evaluated and resolved to the optimal choice to match the user's input.
        <br/><br/>The resolved quote is used to construct a bridge transaction for the user to sign.</td>
      </tr>
      <tr>
        <td><strong>Request</strong></td>
        <td>User signs and submits the bridge transaction on-chain.
        <br/><br/>Their assets are deposited into a Bridge contract on the `originChain`.</td>
      </tr>
      <tr>
        <td><strong>Relay</strong></td>
        <td>A [Relayer] completes the user's bridge on the `destChain` by delivering the user's desired funds.
        <br/><br/>At this point, the bridge is complete from the user's perspective.</td>
      </tr>
      <tr>
        <td><strong>Prove</strong></td>
        <td>The [Relayer] submits an optimistic proof of their completed relay to the `originChain`.</td>
      </tr>
      <tr>
        <td><strong>Guard</strong></td>
        <td>After a proof is asserted for each bridge, a brief optimistic dispute period is initiated.
        <br/><br/>During this period, [Guard] entities will verify that the proof is valid and accurate.</td>
      </tr>
      <tr>
        <td><strong>Claim</strong></td>
        <td>After the dispute period has passed, the [Relayer] may claim the user's originally deposited funds.
        <br/><br/>This reimburses the [Relayer] for the funds which they delivered in the earlier [relay] step.</td>
      </tr>
    </tbody>
  </table>
</div>

## Entities

<b>Users</b> <span style={{color: 'darkgray'}}><i>(Permissionless Role)</i></span>
    <blockquote>
        Uses a bridge interface to submit their intent to the RFQ system for a quote, and then to the chain for fulfillment.
    </blockquote><br/>

<b>Quoters</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Posts Passive and/or Active Quotes to be matched against user bridge inputs.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating entity typically acts as their own Quoter, Relayer, and Prover.</i></span>

    </blockquote><br />

<b>Relayers</b> <span style={{color: 'darkgray'}}><i>(Permissionless Role)</i></span>
    <blockquote>
        Observes [bridge] deposits and submits [relay] transactions to fulfill them.
        <div></div>
        Also submits `claim` transactions to be reimbursed for their relays.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating entity typically acts as their own Quoter, Relayer, and Prover.</i></span>
    </blockquote><br />

<b>Provers</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Observes [relay]s and submits [prove] transactions to initiate relayer reimbursements.
       <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating entity typically acts as their own Quoter, Relayer, and Prover.</i></span>
    </blockquote><br />

<b>Guards</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Validates proofs during the [Dispute Period] and submits [dispute] transactions for any discrepancies found.
        <div></div>
        Currently, Synapse itself is currently the sole Guard operator of the Synapse RFQ system.
    </blockquote><br />

<b>Cancelers</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Allows pre-emptive cancellation of bridge requests which have already been deposited but have not yet been relayed.
        <div></div>
        Currently, Synapse itself is currently the sole Canceler operator of the Synapse RFQ system.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: Incomplete bridge requests can also be permissionlessly canceled after a cancellation window has passed.</i></span>
    </blockquote>



## Architecture

### Quoter API
The RESTful [Quoter API] allows Quoters to post Passive and/or Active Quotes.
These quotes are then used to create transactions for [User]s to sign and submit.

### FastBridge Smart Contracts
[Synapse FastBridge contracts](/docs/Contracts/RFQ) facilitate and enforce all of the on-chain functionality and security of the system.

This includes the [bridge] deposit, escrow, and eventual release of user funds on the origin chain.

Associated [relay], [prove], [claim], [dispute], and [cancel] actions are also facilitated by FastBridge contracts.

Additionally, the FastBridge contracts manage all relevant permissions and on-chain configuration settings of the system.

FastBridge contracts will be deployed on all supported chains.


## On-Chain Flow Detail
Once the [User] has signed and submitted their deposit on-chain via a [bridge] transaction, a [BridgeRequested] event will be emitted.

[Relayer]s who observe this event can permissionlessly complete the bridge by calling [relay] on the FastBridge contract of the destination chain.

The [BridgeRequested] event emission includes the `request` bytes of an encoded [BridgeTransactionV2] struct. The [Relayer] must execute their [relay] with these exact bytes from the origin chain, or their relay will not be considered valid and could result in a loss of relayer funds. To reiterate, the destination contract must operate optimistically upon the [Relayer]'s supplied parameters - it has no capability to verify or protect the [Relayer] from loss of funds due to error or ReOrgs on the origin chain. All relays are final and all participating [Relayer]s must assume the full responsibility and risk of invalid relays.

Assuming the [Relayer] has sufficient funds and necessary ERC20 approvals on the destination chain, the [relay] transaction will facilitate the delivery of funds from the [Relayer] to the [User] via the FastBridge contract as intermediary. This will also emit a [BridgeRelayed] event and prevent any further attempts to relay that transaction.


At this point, the process is complete from the [User]'s perspective. However, the [Relayer] needs to be reimbursed for the funds they delivered.

After the relay, a permissioned [Prover] (typically also the Relayer) who observed the [BridgeRelayed] event will submit a [prove] transaction to the origin chain's FastBridge contract. The [prove] parameters include the relay transaction hash along with `transactionId`, which is a 32-byte hash of the encoded [BridgeTransactionV2] struct also emitted on the [BridgeRelayed] event.

The [prove] transaction will emit a [BridgeProofProvided] event, store the [proof] in the bridgeTxDetails mapping, and initiate a [Dispute Period] for this transaction.

Shortly after the [Dispute Period] begins, a permissioned [Guard] entity who observed the [BridgeProofProvided] event will verify that it's asserted [relay] transaction is finalized and accurate, including whether it exactly matches the parameters from the user's original bridge transaction. If any discrepancies are discovered, the [Guard] will submit a [dispute] transaction on the origin chain which will effectively erase the [proof], at which time a new correct [proof] can be submitted.

Once the [Dispute Period] has passed without incident, a [claim] transaction can be executed on the origin chain which will release the user's originally deposited funds from escrow and deliver them to the indicated [Relayer] address. The [claim] function can be executed permissionlessly, but in practice the [Relayer] will typically execute it themselves.

## [Cancellations](docs/RFQ/Canceling/)

Due to various factors, it is possible but uncommon for a [bridge] to go unfilled.

In these situations, the user's funds will remain safely in escrow until the bridge can be [cancelled](docs/RFQ/Canceling/),  at which point the funds will be securely returned to the user.

For more detail, see the [Canceling](docs/RFQ/Canceling/) section.

<!-- :::note Signing quotes

Relayers authenticate quotes by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). See the canonical implementation [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq).

::: -->

<!-- RFQ consists of three components, with each of the two off-chain components being ran by different actors: -->

<!-- ### [API](API)

Off-chain service ran by Quoters. user-interfaces that allows market makers/relayers to post quotes on different bridge routes. Solvers that have registered with the FastBridge contract can sign messages that post quotes signifying at what price they are willing to bridge tokens on a certain route. -->

<!-- In the canonical implementation, Solvers authenticated by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). The canonical implementation can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq). -->

<!-- ### Fast Bridge Contract

The fast bridge contract is the core of the RFQ protocol and what allows solvers  to fulfill requests from users. A user deposits their funds into the FastBridge contract along with the lowest price they are willing to accept for a given route (a price they get by reading quotes from the Quoter). -->

<!-- In the unlikely event no Solver is available to fulfill a users request, a user can permissionlessly  claim their funds back after waiting an optimistic period. -->

<!-- Contract code level documentation can be found [here](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html). -->

<!-- ### Relayer

The relayer is a service ran by the solvers. The relayer is responsible for posting quotes & fulfilling requests. While the relayer can be implemented in any way, the canonical implementation is a golang based relayer that provides a way to decide what chains/routes to quote on, how much to quote and which addresses not to relay for. -->

## Dispute Period and Guards

The RFQ system optimistically assumes all [prove] transactions are accurate unless they are disputed. These disputes are effectuated during the `Dispute Window`  dispute window in which [Guard] entities may execute a dispute if they detect errors or fraudulent activity, such as incorrect fill amounts or proofs submitted by the wrong relayer.

In a successful dispute, the relayer loses their claimable funds. This design is intended to enforce honest behavior while also protecting honest relayers in cases of blockchain reorgs.

