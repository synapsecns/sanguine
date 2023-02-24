// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { SnapAttestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SnapshotHub } from "./hubs/SnapshotHub.sol";

/**
 * @notice Accepts snapshots signed by Guards and Notaries. Verifies Notaries attestations.
 */
contract Summit is StatementHub, SnapshotHub {
    /**
     * @notice Emitted when a proof of invalid attestation is submitted.
     * @param attestation   Raw payload with attestation data
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidAttestation(bytes attestation, bytes attSignature);

    /**
     * @notice Emitted when a snapshot is accepted by the Summit contract.
     * @param domain        Domain where the signed Agent is active (ZERO for Guards)
     * @param agent         Agent who signed the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event SnapshotAccepted(
        uint32 indexed domain,
        address indexed agent,
        bytes snapshot,
        bytes snapSignature
    );

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Submit a snapshot (list of states) signed by a Guard or a Notary.
     * Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * Notary-signed snapshots: Attestation Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * Notary could then proceed to sign the attestation for their submitted snapshot.
     * @dev Will revert if either of these is true:
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
        returns (bool wasAccepted)
    {
        // This will revert if payload is not a snapshot, or signer is not an active Agent
        (Snapshot snapshot, uint32 domain, address agent) = _verifySnapshot(
            _snapPayload,
            _snapSignature
        );
        if (domain == 0) {
            // This will revert if Guard has previously submitted
            // a fresher state than one in the snapshot.
            _acceptGuardSnapshot(snapshot, agent);
        } else {
            // This will revert if any of the states from the Notary snapshot
            // haven't been submitted by any of the Guards before.
            _acceptNotarySnapshot(snapshot, agent);
        }
        emit SnapshotAccepted(domain, agent, _snapPayload, _snapSignature);
        return true;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Verifies an attestation signed by a Notary.
     *  - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     *  - Slashes the Notary otherwise (meaning the attestation is invalid).
     * @dev Will revert if either of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     * @param _attPayload       Raw payload with SnapAttestation data
     * @param _attSignature     Notary signature for the attestation
     * @return isValid          Whether the provided attestation is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not an attestation, or signer is not an active Notary
        (SnapAttestation snapAtt, uint32 domain, address notary) = _verifyAttestation(
            _attPayload,
            _attSignature
        );
        isValid = _isValidAttestation(snapAtt);
        if (!isValid) {
            emit InvalidAttestation(_attPayload, _attSignature);
            _slashAgent(domain, notary);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _slashAgent(uint32 _domain, address _account) internal {
        // TODO: Move somewhere else?
        // TODO: send a system call indicating agent was slashed
        _removeAgent(_domain, _account);
    }

    function _isIgnoredAgent(uint32, address) internal view virtual override returns (bool) {
        // Summit keeps track of every agent
        return false;
    }
}
