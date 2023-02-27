// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { ISummit } from "./interfaces/ISummit.sol";
import { SnapAttestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SnapshotHub } from "./hubs/SnapshotHub.sol";

/**
 * @notice Accepts snapshots signed by Guards and Notaries. Verifies Notaries attestations.
 */
contract Summit is StatementHub, SnapshotHub, ISummit {
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

    /// @inheritdoc ISummit
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

    /// @inheritdoc ISummit
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
