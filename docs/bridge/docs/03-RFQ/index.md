---
title: RFQ
---

import { RFQFlow } from '@site/src/components/RFQFlow'

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

# Synapse RFQ Summary

Synapse RFQ (Request-For-Quote) is an <abbr title="'Intent' refers to a user authorizing specific actions that they want to achieve, typically in very simple terms, such as a bridge or swap. Actual execution of the actions is then performed on the user's behalf by third parties known as solvers/relayers.">intent</abbr>-based bridging system that connects bridging users to a network of relayers.

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
        <td>Quotes from Relayers are evaluated and resolved to the optimal choice to match the user's input.
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
        Posts Passive and/or Active Quotes via the [Quoter API] to be matched against user bridge inputs.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating Quoter typically acts as their own Relayer and Prover.</i></span>

    </blockquote><br />

<b>Relayers</b> <span style={{color: 'darkgray'}}><i>(Permissionless Role)</i></span>
    <blockquote>
        Observes bridge deposits and submits [relay] transactions to fulfill them.
        <div></div>
        Also submits `claim` transactions to be reimbursed for their relays.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating Relayer typically acts as their own Quoter and Prover.</i></span>
    </blockquote><br />

<b>Provers</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Observes [relay]s and submits [prove] transactions to initiate relayer reimbursements.
       <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: In practice, each participating Prover typically acts as their own Quoter and Relayer.</i></span>
    </blockquote><br />

<b>Guards</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Validates proofs during the [Dispute Period] and submits [dispute] transactions for any discrepancies found.
        <div></div>
        Currently, Synapse itself is the sole Guard operator of the Synapse RFQ system.
    </blockquote><br />

<b>Cancelers</b> <span style={{color: 'darkgray'}}><i>(Permissioned Role)</i></span>
    <blockquote>
        Able to manually pre-emptively cancel bridge requests which have been deposited, have not yet been relayed, and are past their relay deadline.
        <div></div>
        Currently, Synapse itself is the sole Canceler operator of the Synapse RFQ system.
        <div></div>
        <span style={{color: 'darkgray', fontSize: '0.9em'}}><i>Note: Incomplete bridge requests can also be canceled permissionlessly (without any involvement from a Canceler) after a longer cancellation window has passed.</i></span>
    </blockquote>



## Architecture

### Quoter API
The RESTful [Quoter API] allows Quoters to post Passive and/or Active Quotes.
These quotes are then used to create transactions for Users to sign and submit.

### FastBridge Smart Contracts
[Synapse FastBridge contracts](/docs/Contracts/RFQ) facilitate and enforce all of the on-chain functionality and security of the system.

This includes the bridge, deposit, escrow, and eventual release of user funds on the origin chain.

Associated [relay], [prove], [claim], [dispute], and [cancel] actions are also facilitated by FastBridge contracts.

Additionally, the FastBridge contracts manage all relevant permissions and on-chain configuration settings of the system.

FastBridge contracts will be deployed on all supported chains.


##  Flow Summary

#### [Quoting]
<blockquote>
A [User] inputs their desired bridge information into an bridge interface such as [Synapse](https://synapseprotocol.com/?fromChainId=1).

Quotes to complete the bridge are collected from [Quoters]. The best of these is resolved and presented to the user as a bridge transaction to sign & submit on-chain.
</blockquote>

#### [Bridging]
<blockquote>
Once the [User] has signed and submitted their deposit on-chain via a bridge transaction, the bridging funds will be transferred to a FastBridge contract and held in escrow.

</blockquote>

#### [Relaying]

<blockquote>
Relayers who observe the bridge deposit event can permissionlessly complete the bridge by calling [relay] on the FastBridge contract of the destination chain. Typically the Relayer is the same as the Quoter.

This will transfer the bridge destination funds from the [Relayer] to the [User], with FastBridge acting as intermediary.

At this point, the process is complete from the [User]'s perspective. However, the [Relayer] needs to be reimbursed for the funds they delivered.
</blockquote>

#### [Proving]

<blockquote>

After the relay, a [Prover] (typically also the Relayer) will submit a [prove] transaction with the relay transaction's details to the origin chain's FastBridge contract.

This asserts that the [relay] for the bridge was completed according to the specifications of the [User], and that the [Relayer] can rightfully [claim] the escrowed funds.

A [Dispute Period] begins for the transaction before the [claim] can occur.

</blockquote>

#### [Dispute Period]

<blockquote>
During the [Dispute Period], [Guard] entities will verify that the proof's assertion is accurate.

IE: They will confirm that the [relay] being asserted was completed exactly as specified by the user.

If any discrepancies are found, the guards will [dispute] the proof
</blockquote>

#### [Claiming]

<blockquote>
Once the [Dispute Period] has passed without incident, a [claim] transaction can be executed by the [Relayer] on the origin chain.

This wil release the deposit funds from escrow and deliver them to the rightful [Relayer] as a reimbursement for the liquidity they provided on the [relay].
</blockquote>
