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

> **Note**: `srcSender` and `dstReceiver` are the addresses of the App contracts on the source and destination chains, respectively. For EVM chains, only the lowest 20 bytes of the contract address are used, while the highest 12 bytes are zeroed out.

2. In the same transaction, a set of Interchain Modules are called to verify the batch on the destination chain. The `InterchainDB` contract on the source chain emits the `InterchainBatchVerificationRequested` event:

```solidity
    event InterchainBatchVerificationRequested(
        uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules
    );
```

3. Message could be executed on destination chain once enough modules have verified the batch. The amount of required verifications, as well as the module addresses are defined by the application config of `dstReceiver` contract.

```solidity
    /// @notice Returns the verification configuration of the Interchain App.
    /// @dev This configuration is used by the Interchain Client to verify that message has been confirmed
    /// by the Interchain Modules on the destination chain.
    /// Note: V1 version of AppConfig includes the required responses count, and optimistic period after which
    /// the message is considered confirmed by the module. Following versions may include additional fields.
    /// @return appConfig    The versioned configuration of the Interchain App, encoded as bytes.
    /// @return modules      The list of Interchain Modules that app is trusting to confirm the messages.
    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules)
```

> **Note**: there isn't currently an exposed method to decode the `appConfig` bytes into the struct. This will be added SOON (TM).

```solidity
struct AppConfigV1 {
  uint256 requiredResponses;
  uint256 optimisticPeriod;
}
```

4. The next step is the verification of batch that contains the message. Once the Interchain Module verifies the batch, the `InterchainDB` contract on the destination chain emits the `InterchainBatchVerified` event:

```solidity
    event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot);
```

> **Note**: the `module` address is the **destination chain** address of the Interchain Module that verified the batch.

5. The destination app only accepts the module verification once `optimisticPeriod` has passed since it happened. A message could be executed once at least `requiredResponses` modules have successfully verified the batch (as far the app is concerned).

6. Eventually, the message is executed on the destination chain. The `InterchainClientV1` contract on the destination chain emits the `InterchainTransactionReceived` event:

```solidity
    event InterchainTransactionReceived(
        bytes32 indexed transactionId,
        uint256 indexed dbNonce,
        uint64 indexed entryIndex,
        uint256 srcChainId,
        bytes32 srcSender,
        bytes32 dstReceiver
    );
```
