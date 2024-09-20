### Synapse RFQ

The Synapse RFQ contract source code can be found [here](https://github.com/synapsecns/sanguine/tree/master/packages/contracts-rfq) along with generated documentation [here](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html)

| Chain    | Address                                                                                                                          |
| -------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Arbitrum | `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://arbiscan.io/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E)             |
| Base     | `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://basescan.org/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E)            |
| Ethereum | `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://etherscan.io/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E)            |
| Optimism | `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://optimistic.etherscan.io/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E) |
| Scroll   | `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://scrollscan.com/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E)          |
| Linea    | `0x34F52752975222d5994C206cE08C1d5B329f24dD` [↗](https://lineascan.build/address/0x34F52752975222d5994C206cE08C1d5B329f24dD)         |
| BNB Chain| `0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E` [↗](https://bscscan.com/address/0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E)             |
| Blast    | `0x34F52752975222d5994C206cE08C1d5B329f24dD` [↗](https://blastscan.io/address/0x34F52752975222d5994C206cE08C1d5B329f24dD)            |

### On-Chain Architecture & Transaction Flow

The RFQ contract allows users to post bridge requests based on quotes they have received from the solvers. At a high level, the contract works as follows:

1. **User calls bridge**: The user calls the bridge contract with the quote they have received from the RFQ API and passing in origin, destination and other paramaters as a [BridgeParam](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgeparams).
2. **Bridge emits event**: The bridge contract emits a [`BridgeRequested`](https://vercel-rfq-docs.vercel.app/contracts/interfaces/IFastBridge.sol/interface.IFastBridge.html#bridgerequested) event.
3. **Solver relays request**: The solver relays the request by calling the [`relay`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#relay) function on the RFQ contract. The contract then pulls the tokens from the solvers wallet (or [msg.value](https://ethereum.stackexchange.com/questions/43362/what-is-msg-value) in the case of eth) and sends them to the user filling their order.
4. **Solver Calls Prove**: The solver then calls the [`prove`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#prove) function on the RFQ contract to prove they have filled the order. In the current implementation, the function must be called from the solver address.
5. **User Claims**: If the solver does not call prove within the optimistic window, the user can call the [`claim`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#claim) function to claim their funds back.

### On-Chain Statuses

Like the relayer, each transaction in the RFQ contract has a status. The statuses are as follows:

| Status          | Int | Meaning                                                                                                                                                               |
|-----------------|-----|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Null            | 0   | Bridge transaction doesn't exist yet on the origin chain.                                                                                                             |
| Requested       | 1   | A bridge has been requested, but the [`prove`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#prove) has not yet been called    |
| Relayer Proved  | 2   | The relayer has tried to [`prove`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#prove) the transaction, but cannot claim yet. |
| Relayer Claimed | 3   | The relayer has called [`claim`](https://vercel-rfq-docs.vercel.app/contracts/FastBridge.sol/contract.FastBridge.html#claim) and gotten the funds.                    |
| Refunded        | 4   | The relayer has not called `claim` within the optimistic period or a dispute has been decided in favor of the user and the users been refunded.                       |


### Dispute Period and Guards

The RFQ system includes a dispute period and guards to ensure the integrity of bridge transactions. Here's how it works:

After a relayer submits a proof for a bridge transaction, there's a set period during which the transaction can be disputed. This period allows for detection and correction of any errors or fraudulent activities. Guards are responsible for monitoring bridge transactions and initiating disputes if they detect issues such as incorrect fill amounts and proofs submitted by the wrong relayer. A successful dispute would end up with the relayer losing their claimable funds.

The current implementation is architectured to enforce honest behavior and also protect honest relayers in cases of blockchain reorgs.

### Other Functionalities

**ChainGas**

`sendChainGas` is a field that is populated by the bridge user, and it's a simple bool flag. If `sendChainGas=true` the amount is specified in the FastBridge contract on the destination chain as `chainGasAmount`. This is currently set to zero in all the contracts, and can thus be ignored by filling orders with either no `sendChainGas` option (or to chains with `chainGasAmount==0`)
