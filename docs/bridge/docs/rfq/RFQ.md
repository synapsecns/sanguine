# RFQ

RFQ is a bridge module supported by the Synapse Router that allows different market makers to post quotes on different bridge routes. Users can take these quotes by submitting an on-chain bridge request. In the event these requests are not fulfilled, users can request a refund after a set period of time.

### Actors

With the exception of the smart contract itself, RFQ is agnostic to how users receive quotes and where Solvers choose to post quotes. Below, we explain who the general actors interacting with the contract are and then explain the canonical RFQ implementation.

- **Solvers -** Solvers (also known as fillers) are market makers that provide liquidity & post quotes to the API. They are then in charge of fulfilling requests on-chain.
- **Users** - End users observe quotes from solvers and post requests on chain. In the event these requests cannot be fulfilled, the user can reclaim their funds after the optimistic window has passed.
- **Quoter -** The quoter runs a service ran by different interfaces to the Synapse Bridge that allows market makers to post quotes and users to read them. The spec of RFQ does not require this to be an “API” in the traditional sense. Interfaces can use protocols like libp2p, irc and dht’s to communicate quotes.

Right now, RFQ consists of three-different components, with each of the two off-chain components being ran by different actors:

- **[API](./API) -** The RFQ api is an off-chain service ran by Quoters. user-interfaces that allows market makers/solvers to post quotes on different bridge routes. Solvers that have registered with the FastBridge contract can sign messages that post quotes signifying at what price they are willing to bridge tokens on a certain route.

  In the canonical implementation, users Solvers authenticated by signing requests with their private key in accordance with [EIP-191](https://eips.ethereum.org/EIPS/eip-191). The canonical implementation can be found [here](https://github.com/synapsecns/sanguine/tree/master/services/rfq).
- **Fast Bridge Contract -** The fast bridge contract is the core of the RFQ protocol and what allows solvers  to fulfill requests from users. A user deposits their funds into the FastBridge contract along with the lowest price they are willing to accept for a given route (a price they get by reading quotes from the Quoter).

  In the unlikely event no Solver is available to fulfill a users request, a user can permissionlessly  claim their funds back after waiting an optimistic period.

  Contract code level documentation can be found [here](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html).
- **Relayer** - The relayer is a service ran by the solvers. The relayer is responsible for posting quotes & fulfilling requests. While the relayer can be implemented in any way, the canonical implementation is a golang based relayer that provides a way to decide what chains/routes to quote on, how much to quote and which addresses not to relay for.

