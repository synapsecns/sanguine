// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice Minimal interface for Summit contract, required for sending messages.
interface ISummit {
    /**
     * @notice Emitted when a snapshot is accepted by the Summit contract.
     * @param domain        Domain whether the signed Agent is active (ZERO for Guards)
     * @param agent         Agent who signed the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param signature     Agent signature for the snapshot
     */
    event SnapshotAccepted(
        uint32 indexed domain,
        address indexed agent,
        bytes snapshot,
        bytes signature
    );

    /**
     * @notice Emitted when an attestation is created from an accepted Notary snapshot.
     * @param domain        Domain where the Notary is active
     * @param notary        Notary who submitted the accepted snapshot
     * @param attestation   Raw payload with attestation data
     */
    event AttestationCreated(uint32 indexed domain, address indexed notary, bytes attestation);

    /**
     * @notice Emitted when a proof of incorrect attestation is submitted.
     * @param domain        Domain where the Notary is active
     * @param attestation   Raw payload with attestation data
     * @param signature     Notary signature for the attestation
     */
    event IncorrectAttestation(uint32 indexed domain, bytes attestation, bytes signature);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               EXTERNAL                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Submit a snapshot (list of states) signed by a Guard or a Notary.
     * Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * Notary-signed snapshots: Attestation Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * Notary could then proceed to sign the attestation for their submitted snapshot.
     * @dev Will revert if either of these is true:
     * - Provided payload is not a formatted Snapshot.
     * - Snapshot signer is not an active Guard or an active Notary.
     * - Notary snapshot contains a state that hasn't been previously submitted by a Guard.
     * Note that Notary will NOT be slashed for submitting such a snapshot.
     * @param _payload      Raw payload with snapshot data
     * @param _signature    Agent signature for the snapshot
     * @return wasAccepted  Whether the snapshot was accepted by the Summit contract
     */
    function submitSnapshot(bytes memory _payload, bytes memory _signature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Verifies an attestation signed by a Notary.
     * - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     * - Slashes the Notary, if the attestation is invalid.
     * @dev Will revert if either of these is true:
     * - Provided payload is not a formatted Attestation.
     * - Attestation signer is not an active Notary.
     * @param _payload      Raw payload with attestation data
     * @param _signature    Notary signature for the attestation
     * @return isValid      Whether the provided attestation is valid.
     *                      Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory _payload, bytes memory _signature)
        external
        returns (bool isValid);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks if attestation was previously submitted as a snapshot by any Notary,
     * and returns such a snapshot.
     * @dev Will revert if provided payload is not a formatted Attestation.
     * @param _payload      Raw payload with attestation data
     * @return snapshotData Raw payload with snapshot data used for creating the attestation.
     *                      Empty payload, if attestation is not valid.
     */
    function decodeAttestation(bytes memory _payload)
        external
        view
        returns (bytes memory snapshotData);

    /**
     * @notice Retrieves the latest Origin state submitted by a given guard in a snapshot.
     * Notaries are expected to double-check this data before signing it.
     * @param _origin       Domain of origin chain
     * @param _guard        Guard address to check
     * @return stateData    Formatted State payload containing data about state of the Origin
     */
    function getLatestState(uint32 _origin, address _guard)
        external
        view
        returns (bytes memory stateData);

    /**
     * @notice Retrieves the latest Origin states submitted by any of the guards in a snapshot.
     * Notaries are expected to double-check this data before signing it.
     * @param _origin           Domain of origin chain
     * @return stateDataArray   List of formatted State payloads, one for every active Guard.
     *                          Empty states are removed from the list.
     */
    function getLatestStates(uint32 _origin) external view returns (bytes[] memory stateDataArray);
}
