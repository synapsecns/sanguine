// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
    /// @notice Relays destination side of bridge transaction by off-chain relayer
    /// @param request The encoded bridge transaction to relay on destination chain
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function relay(bytes memory request, address relayer) external payable;

    /// @notice Provides proof on origin side that relayer provided funds on destination side of bridge transaction
    /// @param request The encoded bridge transaction to prove on origin chain
    /// @param destTxHash The destination tx hash proving bridge transaction was relayed
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function prove(bytes memory request, bytes32 destTxHash, address relayer) external;

    /// @notice Completes bridge transaction on origin chain by claiming originally deposited capital. Can only send funds to the relayer address on the proof.
    /// @param request The encoded bridge transaction to claim on origin chain
    function claim(bytes memory request) external;

    enum BridgeStatus {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED
    }

    struct ProofDetail {
        uint40 blockTimestamp;
        uint48 blockNumber;
        address relayer;
    }

    struct BridgeTxDetails {
        BridgeStatus status;
        ProofDetail proof;
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
