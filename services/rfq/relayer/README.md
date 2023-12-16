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
   2. If these chekcs pass, it's stored as `CommittedPending` because we're still not sure it's pending on chain, but we have committed our liquidity to it.
1.
### Quote Posting

TODO
