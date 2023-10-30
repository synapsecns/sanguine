// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ISnapshotHub {
    /**
     * @notice Check that a given attestation is valid: matches the historical attestation
     * derived from an accepted Notary snapshot.
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     * @param attPayload    Raw payload with attestation data
     * @return isValid      Whether the provided attestation is valid
     */
    function isValidAttestation(bytes memory attPayload) external view returns (bool isValid);

    /**
     * @notice Returns saved attestation with the given nonce.
     * @dev Reverts if attestation with given nonce hasn't been created yet.
     * @param attNonce      Nonce for the attestation
     * @return attPayload   Raw payload with formatted Attestation data
     * @return agentRoot    Agent root hash used for the attestation
     * @return snapGas      Snapshot gas data used for the attestation
     */
    function getAttestation(uint32 attNonce)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas);

    /**
     * @notice Returns the state with the highest known nonce submitted by a given Agent.
     * @param origin        Domain of origin chain
     * @param agent         Agent address
     * @return statePayload Raw payload with agent's latest state for origin
     */
    function getLatestAgentState(uint32 origin, address agent) external view returns (bytes memory statePayload);

    /**
     * @notice Returns latest saved attestation for a Notary.
     * @param notary        Notary address
     * @return attPayload   Raw payload with formatted Attestation data
     * @return agentRoot    Agent root hash used for the attestation
     * @return snapGas      Snapshot gas data used for the attestation
     */
    function getLatestNotaryAttestation(address notary)
        external
        view
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas);

    /**
     * @notice Returns Guard snapshot from the list of all accepted Guard snapshots.
     * @dev Reverts if snapshot with given index hasn't been accepted yet.
     * @param index             Snapshot index in the list of all Guard snapshots
     * @return snapPayload      Raw payload with Guard snapshot
     * @return snapSignature    Raw payload with Guard signature for snapshot
     */
    function getGuardSnapshot(uint256 index)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature);

    /**
     * @notice Returns Notary snapshot from the list of all accepted Guard snapshots.
     * @dev Reverts if snapshot with given index hasn't been accepted yet.
     * @param index             Snapshot index in the list of all Notary snapshots
     * @return snapPayload      Raw payload with Notary snapshot
     * @return snapSignature    Raw payload with Notary signature for snapshot
     */
    function getNotarySnapshot(uint256 index)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature);

    /**
     * @notice Returns Notary snapshot that was used for creating a given attestation.
     * @dev Reverts if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation is invalid (doesn't have a matching Notary snapshot).
     * @param attPayload        Raw payload with attestation data
     * @return snapPayload      Raw payload with Notary snapshot
     * @return snapSignature    Raw payload with Notary signature for snapshot
     */
    function getNotarySnapshot(bytes memory attPayload)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature);

    /**
     * @notice Returns proof of inclusion of (root, origin) fields of a given snapshot's state
     * into the Snapshot Merkle Tree for a given attestation.
     * @dev Reverts if any of these is true:
     *  - Attestation with given nonce hasn't been created yet.
     *  - State index is out of range of snapshot list.
     * @param attNonce      Nonce for the attestation
     * @param stateIndex    Index of state in the attestation's snapshot
     * @return snapProof    The snapshot proof
     */
    function getSnapshotProof(uint32 attNonce, uint8 stateIndex) external view returns (bytes32[] memory snapProof);
}
