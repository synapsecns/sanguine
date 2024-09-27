// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
    struct BridgeRelay {
        uint48 blockNumber;
        uint48 blockTimestamp;
        address relayer;
    }

    /// @notice Relays destination side of bridge transaction by off-chain relayer
    /// @param request The encoded bridge transaction to relay on destination chain
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function relay(bytes memory request, address relayer) external payable;

    /// @notice Provides proof on origin side that relayer provided funds on destination side of bridge transaction
    /// @param transactionId The transaction id associated with the encoded bridge transaction to prove
    /// @param destTxHash The destination tx hash proving bridge transaction was relayed
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) external;

    /// @notice Completes bridge transaction on origin chain by claiming originally deposited capital.
    /// @notice Can only send funds to the relayer address on the proof.
    /// @param request The encoded bridge transaction to claim on origin chain
    function claim(bytes memory request) external;
    /// @notice Checks if a transaction has been relayed
    /// @param transactionId The ID of the transaction to check
    /// @return True if the transaction has been relayed, false otherwise
    function bridgeRelays(bytes32 transactionId) external view returns (bool);
    enum BridgeStatus {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED
    }

    struct BridgeTxDetails {
        BridgeStatus status;
        uint40 proofBlockTimestamp;
        uint48 proofBlockNumber;
        address proofRelayer;
    }

    /// @notice Returns the status of a bridge transaction
    /// @param transactionId The ID of the bridge transaction
    /// @return The status of the bridge transaction
    function bridgeStatuses(bytes32 transactionId) external view returns (BridgeStatus);

    /// @notice Returns the timestamp and relayer of a bridge proof
    /// @param transactionId The ID of the bridge transaction
    /// @return The timestamp and relayer address of the bridge proof
    function bridgeProofs(bytes32 transactionId) external view returns (BridgeProof memory);
}
