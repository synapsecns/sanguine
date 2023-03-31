// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice A collection of events emitted by the Origin contract
abstract contract OriginEvents {
    // Old Event to ensure that go generation works with the existing Agents
    // TODO: remove once agents are updated to handle the new "Dispatched" event
    event Dispatch(
        bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message
    );

    /**
     * @notice Emitted when a new message is dispatched.
     * @param messageHash   Hash of message; the leaf inserted to the Merkle tree for the message
     * @param nonce         Nonce of sent message (starts from 1)
     * @param destination   Destination domain
     * @param message       Raw bytes of message
     */
    event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message);

    /**
     * @notice Emitted when a proof of invalid state in the signed attestation is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param state         Raw payload with state data
     * @param attestation   Raw payload with Attestation data for snapshot
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidAttestationState(uint256 stateIndex, bytes state, bytes attestation, bytes attSignature);

    /**
     * @notice Emitted when a proof of invalid state in the signed snapshot is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature);

    /**
     * @notice Emitted when a proof of invalid state report is submitted.
     * @param srPayload     Raw payload with report data
     * @param srSignature   Guard signature for the report
     */
    event InvalidStateReport(bytes srPayload, bytes srSignature);
}
