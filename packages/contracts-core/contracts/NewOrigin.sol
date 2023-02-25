// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DomainContext, StateHub } from "./hubs/StateHub.sol";
import { SnapAttestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract OriginNew is StatementHub, StateHub, SystemRegistry {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: move Events to a separate contract, once the old Origin is deprecated

    /**
     * @notice Emitted when a proof of invalid state in the signed attestation is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param attestation   Raw payload with SnapAttestation data for snapshot
     * @param attSignature  Notary signature for the attestation
     */
    event InvalidAttestationState(
        uint256 stateIndex,
        bytes snapshot,
        bytes attestation,
        bytes attSignature
    );

    /**
     * @notice Emitted when a proof of invalid state in the signed snapshot is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param snapSignature Agent signature for the snapshot
     */
    event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) {}

    /// @notice Initializes Origin contract:
    /// - msg.sender is set as contract owner
    /// - State of "empty merkle tree" is saved
    function initialize() external initializer {
        // Initialize SystemContract: msg.sender is set as "owner"
        __SystemContract_initialize();
        // Initialize "states": state of an "empty merkle tree" is saved
        _initializeStates();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the snapshot signer, if the state is invalid.
     * @dev Will revert if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     *  - Attestation root is not equal to root derived from the snapshot.
     *  - State index is out of range.
     *  - Snapshot state does not refer to this chain.
     * @param _snapPayload      Raw payload with snapshot data
     * @param _stateIndex       State index to check
     * @param _attPayload       Raw payload with SnapAttestation data
     * @param _attSignature     Notary signature for the attestation
     * @return isValid          Whether the requested state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(
        bytes memory _snapPayload,
        uint256 _stateIndex,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation, or signer is not an active Notary
        (SnapAttestation snapAtt, uint32 domain, address notary) = _verifyAttestation(
            _attPayload,
            _attSignature
        );
        // This will revert if payload is not a snapshot, or snapshot/attestation roots don't match
        Snapshot snapshot = _verifySnapshotRoot(snapAtt, _snapPayload);
        // This will revert, if state index is out of range, or state refers to another domain
        isValid = _isValidState(snapshot.state(_stateIndex));
        if (!isValid) {
            emit InvalidAttestationState(_stateIndex, _snapPayload, _attPayload, _attSignature);
            _slashAgent(domain, notary);
        }
    }

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * Does nothing, if the state is valid (matches the historical state of this contract).
     * Slashes the snapshot signer, if the state is invalid.
     * @dev Will revert if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     *  - State index is out of range.
     *  - Snapshot state does not refer to this chain.
     * @param _snapPayload      Raw payload with snapshot data
     * @param _stateIndex       State index to check
     * @param _snapSignature    Agent signature for the snapshot
     * @return isValid          Whether the requested state is valid.
     *                          Agent is slashed, if return value is FALSE.
     */
    function verifySnapshot(
        bytes memory _snapPayload,
        uint256 _stateIndex,
        bytes memory _snapSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not a snapshot, or signer is not an active Agent
        (Snapshot snapshot, uint32 domain, address agent) = _verifySnapshot(
            _snapPayload,
            _snapSignature
        );
        // This will revert, if state index is out of range, or state refers to another domain
        isValid = _isValidState(snapshot.state(_stateIndex));
        if (!isValid) {
            emit InvalidSnapshotState(_stateIndex, _snapPayload, _snapSignature);
            _slashAgent(domain, agent);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _slashAgent(uint32 _domain, address _account) internal {
        // TODO: Move to SystemRegistry?
        // TODO: send a system call indicating agent was slashed
        _removeAgent(_domain, _account);
    }

    function _isIgnoredAgent(uint32, address) internal view virtual override returns (bool) {
        // Origin keeps track of every agent
        return false;
    }
}
