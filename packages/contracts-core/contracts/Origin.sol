// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { InterfaceOrigin } from "./interfaces/InterfaceOrigin.sol";
import { DomainContext, StateHub } from "./hubs/StateHub.sol";
import { Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, SystemRegistry, InterfaceOrigin {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: move Events to a separate contract, once the old Origin is deprecated

    /**
     * @notice Emitted when a proof of invalid state in the signed attestation is submitted.
     * @param stateIndex    Index of invalid state in the snapshot
     * @param snapshot      Raw payload with snapshot data
     * @param attestation   Raw payload with Attestation data for snapshot
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

    // Old Event to ensure that go generation works with the existing Agents
    // TODO: remove once agents are updated to handle the new "Dispatched" event
    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    /**
     * @notice Emitted when a new message is dispatched.
     * @param messageHash   Hash of message; the leaf inserted to the Merkle tree for the message
     * @param nonce         Nonce of sent message (starts from 1)
     * @param destination   Destination domain
     * @param message       Raw bytes of message
     */
    event Dispatched(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes message
    );

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

    /// @inheritdoc InterfaceOrigin
    function verifyAttestation(
        bytes memory _snapPayload,
        uint256 _stateIndex,
        bytes memory _attPayload,
        bytes memory _attSignature
    ) external returns (bool isValid) {
        // This will revert if payload is not an attestation, or signer is not an active Notary
        (Attestation att, uint32 domain, address notary) = _verifyAttestation(
            _attPayload,
            _attSignature
        );
        // This will revert if payload is not a snapshot, or snapshot/attestation roots don't match
        Snapshot snapshot = _verifySnapshotRoot(att, _snapPayload);
        // This will revert, if state index is out of range, or state refers to another domain
        isValid = _isValidState(snapshot.state(_stateIndex));
        if (!isValid) {
            emit InvalidAttestationState(_stateIndex, _snapPayload, _attPayload, _attSignature);
            _slashAgent(domain, notary);
        }
    }

    /// @inheritdoc InterfaceOrigin
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
    ▏*║                          DISPATCH MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceOrigin
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        // TODO: implement
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
