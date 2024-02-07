# Relayer

The relayer is a service that listens to the RFQ contract for new RFQs and optimistically relays requests to the destination chain. The relayer can then claim the funds on the origin chain after waiting a short time. The relayer is also responsible for continuously posting api quotes with it's available balances so clients know how much liquidity is available.

## How It Works

The relayer is responsible for two things:

1. Continuously posting api quotes with it's available balances so clients know how much liquidity is available.
1. Listening to the RFQ contract for new Bridge transactions and optimistically relaying requests to the destination chain.

In many ways, these can be considered as two seperate programs with a shared state, so we'll treat them as such in the docs.

### Transaction Flow

The relayer consists of two main loops that contain the entire business logic of the relayer. The first loop is the `chainParser` loop which listens for events from on-chain and processes them. The second loop continously pools the db for events in the db which are unfinalized by their status and checks if they can be moved through the queue. Below, we'll go through the full lifecycle of a transaction and the statuses that different paths.

1. Pre-Status: An on-chain transaction emits the event [`BridgeRequested`](https://vercel-rfq-docs-trajan0x-synapsecns.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested). We store this event in the db with the status `Seen` & wait for the db processor to pick it up.
1. `Seen`: The db processor picks up the event
   1. From this step until `Relay()` is called, we check to make sure the deadline has not been exceeded.
   2. It checks with the quoter if it is valid. If not, it is marked as `WillNotProcess`
   1. It checks with the inventory manger that there's enough inventory. If not, it is marked as `NotEnoughInventory` and can be tried later.
   2. If these checks pass, it's stored as `CommittedPending` because we're still not sure it's pending on chain, but we have committed our liquidity to it.
1. `CommitPending`: check the chain to see if transaction is finalized yet, if not wait until it is.
1. `CommitConfirmed`: The transaction is finalized on chain, we can now relay it to the destination chain.
1. `RelayPending`: We now listen (don't poll)  for the relay in the logs. Once we get it we mark the transaction as `RelayComplete`
2. `RelayComplete`: We now call Prove() on the contract to prove that we relayed the transaction. Once this is done, we mark the transaction as `ProvePosting`
3. `ProvePosting`: We've called`Prove()` event from the contract. It's now time to start waiting for it to confirm and give us a log.
4. `ProofPosted`: The proof has been sucessfully submitted to the contract. We now wait for the claim period to expire. Once it does, we mark the transaction as `ClaimPending`
4. `ClaimComplete`: We now wait for the claim period to expire. Once it does, we mark the transaction as `ClaimComplete`

### Quote Posting

The quote posting process is rather rudimentary. The relayer continously fetches a list of its on chain balances and subtraces any open commitments.

Currently, the quotes are standalone; that is, they are not responding to any client requests. The quoter specifies a `FixedFee` parameter that is meant to account for the gas costs associated with executing transactions on the origin and destinations chains.

In a future version, quotes may be issued in a more classic RFQ-style, where they are posted in response to a client request. In that case we can incorporate more precise pricing logic.
