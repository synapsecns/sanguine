<!-- TODO, add a proper intro -->

# Interchain Communication Contracts

[See the Docs](https://communication-docs.vercel.app/)

## Event lifecycle of a sent message

> **Note:** the event signatures are subject to change in the testnet phase.

1. Message is sent through the `InterchainClientV1` contract. The `InterchainClientV1` contract on source chain emits the `InterchainTransactionSent` event:

https://github.com/synapsecns/sanguine/blob/10afc7a61561ff39a988470252e165b4fe7f6a0f/packages/contracts-communication/contracts/events/InterchainClientV1Events.sol#L14-L37

> The sent message is added to the current batch in `InterchainDB` contract on the source chain:
>
> - `dbNonce` is the nonce of the batch
> - `entryIndex` is the index of the message in the batch

> **Note:** in the testnet phase, the batching is effectively disabled, so the `entryIndex` will always be 0.

> **Note**: `srcSender` and `dstReceiver` are the addresses of the App contracts on the source and destination chains, respectively. For EVM chains, only the lowest 20 bytes of the contract address are used, while the highest 12 bytes are zeroed out.

2. In the same transaction, a set of Interchain Modules are called to verify the batch on the destination chain. The `InterchainDB` contract on the source chain emits the `InterchainBatchVerificationRequested` event:

https://github.com/synapsecns/sanguine/blob/10afc7a61561ff39a988470252e165b4fe7f6a0f/packages/contracts-communication/contracts/events/InterchainDBEvents.sol#L28-L36

3. Message could be executed on destination chain once enough modules have verified the batch. The amount of required verifications, as well as the module addresses are defined by the application config of `dstReceiver` contract. This is exposed for the off-chain agents in the InterchainClientV1 contract:

https://github.com/synapsecns/sanguine/blob/7cda26dcdec16637aaa6ec653c073b0b398c1850/packages/contracts-communication/contracts/InterchainClientV1.sol#L212-L216

https://github.com/synapsecns/sanguine/blob/7cda26dcdec16637aaa6ec653c073b0b398c1850/packages/contracts-communication/contracts/libs/AppConfig.sol#L6-L9

4. The next step is the verification of batch that contains the message. Once the Interchain Module verifies the batch, the `InterchainDB` contract on the destination chain emits the `InterchainBatchVerified` event:

https://github.com/synapsecns/sanguine/blob/10afc7a61561ff39a988470252e165b4fe7f6a0f/packages/contracts-communication/contracts/events/InterchainDBEvents.sol#L19-L26

> **Note**: the `module` address is the **destination chain** address of the Interchain Module that verified the batch.

5. The destination app only accepts the module verification once `optimisticPeriod` has passed since it happened. A message could be executed once at least `requiredResponses` modules have successfully verified the batch (as far the app is concerned).

6. Once enough module verifications are received, the message could be executed on the destination chain by any executor. The `InterchainClientV1` contract on the destination chain emits the `InterchainTransactionReceived` event:

https://github.com/synapsecns/sanguine/blob/10afc7a61561ff39a988470252e165b4fe7f6a0f/packages/contracts-communication/contracts/events/InterchainClientV1Events.sol#L39-L54

To sum up, the message status lifecycle is currently:

- Sent on the source chain (with verification requested).
- Waiting for at least `requiredResponses` module verifications on the destination chain.
- Waiting for `optimisticPeriod` to pass since the last verification.
- Waiting for execution on the destination chain (anyone could trigger execution at this point).
- Executed on the destination chain.

> **Note**: once the batching is implemented, there will be an additional step of waiting for the source chain batch to be finalized before it could be verified on the destination chain. The event spec for this is not yet finalized, so it's not included in this document.
