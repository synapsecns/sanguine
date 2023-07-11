# Submitter

Submitter is a module that submits transactions to an evm based-blockchain and bumps them/checks for completion until they are done. It is designed to abstract away gas bumping, confirmation checking, etc from the caller.


## Understanding `SubmitTransaction`

One of the main goals of submitter is for you to be able to call `SubmitTransaction` once and once you're returned a nonce, be confident that the transaction will eventually get through as long as run keeps running. It's important to understand how this process works and why it exists.

### Nonce Locking

The first thing you'll notice about the  `SubmitTransaction` method is it returns a nonce rather than a transaction hash. This is because once the transaction is submitted, it will be bumped (which requires changing the transaction hash) until the transaction is confirmed. The next nonce is generated in the following manner:

```mermaid
graph TB
    style Locker fill:#f9f,stroke:#333,stroke-width:2px;
    style Fetch DB, RPC, Stored, Unlock, Errored, C, D, E fill:#fff,stroke:#333,stroke-width:2px;
    style A fill:#eee,stroke:#333,stroke-width:2px;

    Locker[Lock Nonce Mutex] --> Fetch
    subgraph Fetch [External Fetches ]
        direction TB
        subgraph DB
            direction TB
            B[Get last used nonce in database] --> F[Increment Database Nonce]
        end
        style DB fill:#ccf,stroke:#333,stroke-width:2px;
        style RPC fill:#cfc,stroke:#333,stroke-width:2px;
        subgraph RPC
            A[Get last nonce on-chain]
        end
    end
    style Fetch fill:#ddf,stroke:#333,stroke-width:2px;
    Fetch --> Errored
    Errored{Errored?} -- Yes -->  Unlock
    Errored{Errored?} -- No -->  C
    E --> Stored
    D --> Stored
    Stored --> Unlock
    Stored[Store Association Between Database Nonce and Transaction]
    style Stored fill:#ffc,stroke:#333,stroke-width:2px;
    Unlock[Unlock Nonce Mutex]
    style Unlock fill:#f9f,stroke:#333,stroke-width:2px;
    C{Is on-chain nonce > database nonce?} -- Yes --> D[Use on-chain nonce]
    C -- No --> E[Use database nonce]
```

You'll now notice that there are two failure cases for this method: if either the db or the rpc url cannot be reached you'll have to resubmit the tx. But these failures occur atomically, so you can do this in a retry loop w/ a backoff.




<!-- TODO: mermade diagram of confirmation queue and process queue -->
<!-- aditionally, should describe cases in which submit transaction will return an error-->
<!-- To understand why this module is necessary, you first need to understand that the EVM does not specify anything about transaction submission or consensus. It merely refers to a set of instructions for executing byte-code, so different chains are free to change transaction submission and p2p semantics as they wish. This forms the first constraint on transaction submission. The second constraint is formed by rate-limits, if I send my transaction to an rpc node, even if that node has unlimited throughput, there's no gurantee it will send that to all it's peers and not be rate limited by then. We therefore need to be careful about submitting transactions too far ahead of the current nonce.-->
