// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { OriginEvents } from "./events/OriginEvents.sol";
import { InterfaceOrigin } from "./interfaces/InterfaceOrigin.sol";
import { DomainContext, StateHub } from "./hubs/StateHub.sol";
import { Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, SystemRegistry, OriginEvents, InterfaceOrigin {
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
