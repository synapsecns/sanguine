<!-- TODO, add a proper intro -->

# Interchain Communication Contracts

[See the Docs](https://communication-docs.vercel.app/)

## Event lifecycle of a sent message

> **Note:** the event signatures are subject to change in the testnet phase.

1. Message is sent through the `InterchainClientV1` contract. The `InterchainClientV1` contract on source chain emits the `InterchainTransactionSent` event:

```solidity
    event InterchainTransactionSent(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint64 indexed entryIndex,
        uint256 dstChainId,
        bytes32 srcSender,
        bytes32 dstReceiver,
        uint256 verificationFee,
        uint256 executionFee,
        bytes options,
        bytes message
    );
```

> The sent message is added to the current batch in `InterchainDB` contract on the source chain:
>
> - `dbNonce` is the nonce of the batch
> - `entryIndex` is the index of the message in the batch

> **Note:** in the testnet phase, the batching is effectively disabled, so the `entryIndex` will always be 0.

2. In the same transaction, a set of Interchain Modules are called to verify the batch on the destination chain. The `InterchainDB` contract on the source chain emits the `InterchainBatchVerificationRequested` event:

```solidity
    event InterchainBatchVerificationRequested(
        uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules
    );
```

3. The next step is the verification of batch that contains the message. Once the Interchain Module verifies the batch, the `InterchainDB` contract on the destination chain emits the `InterchainBatchVerified` event:

```solidity
    event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot);
```

> **Note**: the `module` address is the **destination chain** address of the Interchain Module that verified the batch.
