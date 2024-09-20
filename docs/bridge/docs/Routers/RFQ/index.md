# RFQ

A [Synapse Router](../Synapse-Router) bridge module which matches on-chain user requests against posted bridge quotes.

## Agents

[Synapse Fast Bridge contracts](/docs/Contracts/RFQ) coordinate decentralized agents to match user requests with the best quote for a given route, and secure user funds until their transaction is fulfilled.

<!-- https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html -->

| Agents  | Description |
|---------|-------------|
| Quoters | Quote distribution services run through traditional [APIs](API), or protocols like libp2p, irc, or dht
| Solvers | Posts, then fulfills, route quotes through a [Relayer](Relayer), when matched by the Fast Bridge contract against a user request.
| Users   | Uses a route quote to form a bridge request which is matched on-chain by the solver who posted the quote.
| Guards  | Raises a dispute if errors or fraudulent activity are detected

## Transaction Flow

### Status

| Code | Status          | Description |
|------|-----------------|-------------|
| 0    | Null      | Bridge transaction does not exist yet on origin chain |
| 1    | Requested | Bridge requested. Waiting for Relayer |
| 2    | Proved    | Relayer called [`prove`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#prove), and is waiting for the optimistic period to end. |
| 3    | Claimed   | Relayer called [`claim`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#claim) and received their funds. |
| 4    | Refunded  | Relayer did not claim within the optimistic period, or a dispute was decided in favor of the user. |

### Behavior

1. **User requests bridge**: A user signs a [`BridgeRequest`](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeparams) with a received quote and required [`BridgeParams`](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeparams), broadcast as a [`BridgeRequested`](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested) event.
3. **Solver accepts bridge**: A solver calls [`relay`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#relay) on the Bridge contract.
3. **Bridge relays funds**: The Bridge contract sends the requested funds ([`msg.value`](https://ethereum.stackexchange.com/questions/43362/what-is-msg-value) in the case of ETH) from solver to user, on the destination chain.
4. **Solver submits proof**: The solver that accepted the bridge calls [`prove`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#prove) on the Bridge contract.
5. **Solver receives payment**: The Solver receives their funds from the Bridge contract at the end of the optimistic period

<!-- :::note Signing quotes

Solvers authenticate quotes by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). See the canonical implementation [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq).

::: -->

<!-- RFQ consists of three components, with each of the two off-chain components being ran by different actors: -->

<!-- ### [API](API)

Off-chain service ran by Quoters. user-interfaces that allows market makers/solvers to post quotes on different bridge routes. Solvers that have registered with the FastBridge contract can sign messages that post quotes signifying at what price they are willing to bridge tokens on a certain route. -->

<!-- In the canonical implementation, users Solvers authenticated by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). The canonical implementation can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq). -->

<!-- ### Fast Bridge Contract

The fast bridge contract is the core of the RFQ protocol and what allows solvers  to fulfill requests from users. A user deposits their funds into the FastBridge contract along with the lowest price they are willing to accept for a given route (a price they get by reading quotes from the Quoter). -->

<!-- In the unlikely event no Solver is available to fulfill a users request, a user can permissionlessly  claim their funds back after waiting an optimistic period. -->

<!-- Contract code level documentation can be found [here](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html). -->

<!-- ### Relayer

The relayer is a service ran by the solvers. The relayer is responsible for posting quotes & fulfilling requests. While the relayer can be implemented in any way, the canonical implementation is a golang based relayer that provides a way to decide what chains/routes to quote on, how much to quote and which addresses not to relay for. -->

## Dispute Period and Guards

The RFQ system includes an optimistic dispute window in which Guard contracts may initiate a dispute if they detect errors or fraudulent activity, such as incorrect fill amounts or proofs submitted by the wrong relayer.

In a successful dispute the relayer loses their claimable funds. This design is intended to enforce honest behavior while also protecting honest relayers in cases of blockchain reorgs.

## Unfulfilled requests

If a request is not fulfilled, users can reclaim their funds by using the [`claim`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#claim) function once the optimistic window has passed.

## Other Functionality

<!-- #### ChainGas

`sendChainGas` is a field that is populated by the bridge user, and it's a simple bool flag. If `sendChainGas=true` the amount is specified in the FastBridge contract on the destination chain as `chainGasAmount`. This is currently set to zero in all the contracts, and can thus be ignored by filling orders with either no `sendChainGas` option (or to chains with `chainGasAmount==0`) -->
