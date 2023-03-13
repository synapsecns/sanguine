// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { MAX_MESSAGE_BODY_BYTES, SYSTEM_ROUTER } from "./libs/Constants.sol";
import { HeaderLib, MessageLib } from "./libs/Message.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Snapshot } from "./libs/Snapshot.sol";
import { State, StateLib, TypedMemView } from "./libs/State.sol";
import { Tips, TipsLib } from "./libs/Tips.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { OriginEvents } from "./events/OriginEvents.sol";
import { InterfaceOrigin } from "./interfaces/InterfaceOrigin.sol";
import { DomainContext, StateHub } from "./hubs/StateHub.sol";
import { Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
import { SystemRegistry } from "./system/SystemRegistry.sol";

contract Origin is StatementHub, StateHub, SystemRegistry, OriginEvents, InterfaceOrigin {
    using MerkleLib for MerkleLib.Tree;
    using TipsLib for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    MerkleLib.Tree private tree;

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
        uint256 _stateIndex,
        bytes memory _snapPayload,
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
        // This will revert if state index is out of range
        State state = snapshot.state(_stateIndex);
        // This will revert if  state refers to another domain
        isValid = _isValidState(state);
        if (!isValid) {
            emit InvalidAttestationState(
                _stateIndex,
                state.unwrap().clone(),
                _attPayload,
                _attSignature
            );
            // Slash Notary and trigger a hook to send a slashAgent system call
            _slashAgent(domain, notary, true);
        }
    }

    /// @inheritdoc InterfaceOrigin
    function verifySnapshot(
        uint256 _stateIndex,
        bytes memory _snapPayload,
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
            // Slash Agent and trigger a hook to send a slashAgent system call
            _slashAgent(domain, agent, true);
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
        // Modifiers are removed because they prevent from slashing the last active Guard/Notary
        // haveActiveGuard
        // haveActiveNotary(_destination)
        // TODO: figure out a way to filter out unknown domains once Agent Merkle Tree is implemented
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        // This will revert if payload is not a formatted tips payload
        Tips tips = _tips.castToTips();
        // Total tips must exactly match msg.value
        require(tips.totalTips() == msg.value, "!tips: totalTips");
        // Format the message header
        messageNonce = _nextNonce();
        bytes memory header = HeaderLib.formatHeader({
            _origin: localDomain,
            _sender: _checkForSystemRouter(_recipient),
            _nonce: messageNonce,
            _destination: _destination,
            _recipient: _recipient,
            _optimisticSeconds: _optimisticSeconds
        });
        // Format the full message payload
        bytes memory message = MessageLib.formatMessage(header, _tips, _messageBody);

        // Insert new leaf into the Origin Merkle Tree
        messageHash = keccak256(message);
        /// @dev Before insertion: messageNonce == tree.count() - 1
        /// tree.insert() requires amount of leaves AFTER the leaf insertion
        tree.insert(messageNonce, messageHash);

        // Save new State of Origin contract
        /// @dev After insertion: messageNonce == tree.count()
        /// tree.root() requires current amount of leaves
        bytes32 newRoot = tree.root(messageNonce);
        _saveState(StateLib.originState(newRoot));

        // Emit Dispatched event with message information
        emit Dispatched(messageHash, messageNonce, _destination, message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Hook that is called after an existing agent was slashed,
    /// when verification of an invalid agent statement was done in this contract.
    function _afterAgentSlashed(uint32 _domain, address _agent) internal virtual override {
        /// @dev We send a "slashAgent" system message
        /// after the Agent is slashed by submitting an invalid statement.
        _callLocalBondingManager(_dataSlashAgent(_domain, _agent));
    }

    /**
     * @notice Returns adjusted "sender" field.
     * @dev By default, "sender" field is msg.sender address casted to bytes32.
     * However, if SYSTEM_ROUTER is used for "recipient" field, and msg.sender is SystemRouter,
     * SYSTEM_ROUTER is also used as "sender" field.
     * Note: tx will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _checkForSystemRouter(bytes32 _recipient) internal view returns (bytes32 sender) {
        if (_recipient != SYSTEM_ROUTER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_ROUTER has only the highest 12 bytes set,
             * whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             * Thus, in this branch: sender != SYSTEM_ROUTER
             */
        } else {
            // Check that SystemRouter specified SYSTEM_ROUTER as recipient, revert otherwise.
            _assertSystemRouter();
            // Adjust "sender" field for correct processing on remote chain.
            sender = SYSTEM_ROUTER;
        }
    }

    function _isIgnoredAgent(uint32, address) internal view virtual override returns (bool) {
        // Origin keeps track of every agent
        return false;
    }
}
