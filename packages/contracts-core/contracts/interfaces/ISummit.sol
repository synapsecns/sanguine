// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISnapshotHub } from "./ISnapshotHub.sol";

interface ISummit is ISnapshotHub {
    /**
     * @notice Submit a snapshot (list of states) signed by a Guard or a Notary.
     * Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * Notary-signed snapshots: Attestation Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * Notary could then proceed to sign the attestation for their submitted snapshot.
     * @dev Will revert if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     *  - Guard snapshot contains a state older then they have previously submitted
     *  - Notary snapshot contains a state that hasn't been previously submitted by a Guard.
     * Note that Notary will NOT be slashed for submitting such a snapshot.
     * @param _snapPayload      Raw payload with snapshot data
     * @param _snapSignature    Agent signature for the snapshot
     * @return wasAccepted      Whether the snapshot was accepted by the Summit contract
     */
    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Verifies an attestation signed by a Notary.
     *  - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     *  - Slashes the Notary otherwise (meaning the attestation is invalid).
     * @dev Will revert if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     * @param _attPayload       Raw payload with SnapAttestation data
     * @param _attSignature     Notary signature for the attestation
     * @return isValid          Whether the provided attestation is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid);
}
