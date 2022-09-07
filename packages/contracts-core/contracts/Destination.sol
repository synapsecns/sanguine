// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { GlobalNotaryRegistry } from "./registry/GlobalNotaryRegistry.sol";
import { GuardRegistry } from "./registry/GuardRegistry.sol";
import { ReportHub } from "./hubs/ReportHub.sol";
import { Attestation } from "./libs/Attestation.sol";
import { Report } from "./libs/Report.sol";
import { Version0 } from "./Version0.sol";
import { MirrorLib } from "./libs/Mirror.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { Header } from "./libs/Header.sol";
import { Tips } from "./libs/Tips.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
import { SystemMessage } from "./libs/SystemMessage.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
// ============ External Imports ============
import { TypedMemView } from "./libs/TypedMemView.sol";

/**
 * @title Destination
 * @notice Track merkle root state of Origin contracts on other chains,
 * prove and dispatch messages to end recipients.
 */
contract Destination is Version0, SystemContract, ReportHub, GlobalNotaryRegistry, GuardRegistry {
    // ============ Libraries ============

    using MirrorLib for MirrorLib.Mirror;
    using MerkleLib for MerkleLib.Tree;
    using Message for bytes;
    using TypedMemView for bytes29;
    using Attestation for bytes29;
    using Message for bytes29;
    using Header for bytes29;
    using Report for bytes29;

    // ============ Public Storage ============

    // re-entrancy guard
    uint8 private entered;

    uint256 internal mirrorCount;

    // all Mirrors: both active and archived
    mapping(uint256 => MirrorLib.Mirror) internal allMirrors;

    // (domain => mirror index): index of the active mirror in allMirrors
    mapping(uint32 => uint256) internal activeMirrors;

    //TODO: Handle fail-over mirrors and modify activeMirrors
    // (domain => [mirror indexes]): array of indexes of archived mirrors in allMirrors
    mapping(uint32 => uint256[]) internal archivedMirrors;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[45] private __GAP;

    // ============ Events ============

    /**
     * @notice Emitted when message is executed
     * @param messageHash The keccak256 hash of the message that was executed
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash);

    /**
     * @notice Emitted when a root's confirmation is modified by governance
     * @param root The root for which confirmAt has been set
     * @param previousConfirmAt The previous value of confirmAt
     * @param newConfirmAt The new value of confirmAt
     */
    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    event AttestationAccepted(
        uint32 indexed origin,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );

    event NotaryBlacklisted(
        address indexed notary,
        address indexed guard,
        address indexed reporter,
        bytes report
    );

    // ============ Constructor ============

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _localDomain) SystemContract(_localDomain) {}

    // ============ Initializer ============

    /**
     * @notice Initialize the mirror
     * @dev Performs the following action:
     *      - initializes inherited contracts
     *      - initializes re-entrancy guard
     *      - sets remote domain
     *      - sets a trusted root, and pre-approves messages under it
     *      - sets the optimistic timer
     * @param _remoteDomain The domain of the Origin contract this follows
     * @param _notary The EVM id of the notary
     */
    function initialize(uint32 _remoteDomain, address _notary) public initializer {
        __SystemContract_initialize();
        _addNotary(_remoteDomain, _notary);
        // set storage variables
        entered = 1;
        activeMirrors[_remoteDomain] = _createMirror(_remoteDomain);
    }

    // ============ Active Mirror Views ============

    function activeMirrorNonce(uint32 _remoteDomain) external view returns (uint32) {
        return allMirrors[activeMirrors[_remoteDomain]].nonce;
    }

    function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root)
        external
        view
        returns (uint256)
    {
        return allMirrors[activeMirrors[_remoteDomain]].confirmAt[_root];
    }

    function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId)
        external
        view
        returns (bytes32)
    {
        return allMirrors[activeMirrors[_remoteDomain]].messageStatus[_messageId];
    }

    // ============ Archived Mirror Views ============

    // TODO: getters for archived mirrors

    // ============ External Functions ============

    /**
     * @notice Called by external agent. Submits the signed attestation,
     * marks root's allowable confirmation time, and emits an `AttestationAccepted` event.
     * @dev Reverts if signature is invalid.
     * @param _attestation  Attestation data and signature
     */
    function submitAttestation(bytes memory _attestation) external {
        (, bytes29 _view) = _checkNotaryAuth(_attestation);
        uint32 remoteDomain = _view.attestedDomain();
        require(remoteDomain != localDomain, "Attestation refers to local chain");
        uint32 nonce = _view.attestedNonce();
        MirrorLib.Mirror storage mirror = allMirrors[activeMirrors[remoteDomain]];
        require(nonce > mirror.nonce, "Attestation older than current state");
        bytes32 newRoot = _view.attestedRoot();
        mirror.setConfirmAt(newRoot, block.timestamp);
        // update nonce
        mirror.setNonce(nonce);
        emit AttestationAccepted(remoteDomain, nonce, newRoot, _view.notarySignature().clone());
    }

    /**
     * @notice First attempts to prove the validity of provided formatted
     * `message`. If the message is successfully proven, then tries to execute
     * message.
     * @dev Reverts if `prove` call returns false
     * @param _message Formatted message (refer to Message library)
     * @param _proof Merkle proof of inclusion for message's leaf
     * @param _index Index of leaf in origin's merkle tree
     */
    function proveAndExecute(
        uint32 _remoteDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) external {
        require(prove(_remoteDomain, _message, _proof, _index), "!prove");
        execute(_message);
    }

    /**
     * @notice Given formatted message, attempts to dispatch
     * message payload to end recipient.
     * @dev Recipient must implement a `handle` method (refer to IMessageRecipient.sol)
     * Reverts if formatted message's destination domain is not the Mirror's domain,
     * if message has not been proven,
     * or if recipient reverted upon receiving the message.
     * @param _message Formatted message
     */
    function execute(bytes memory _message) public {
        bytes29 messageView = _message.castToMessage();
        bytes29 header = messageView.header();
        uint32 remoteDomain = header.origin();
        MirrorLib.Mirror storage mirror = allMirrors[activeMirrors[remoteDomain]];
        // ensure message was meant for this domain
        require(header.destination() == localDomain, "!destination");
        // ensure message has been proven
        bytes32 messageHash = messageView.keccak();
        bytes32 root = mirror.messageStatus[messageHash];
        require(MirrorLib.isPotentialRoot(root), "!exists || executed");
        require(
            acceptableRoot(remoteDomain, header.optimisticSeconds(), root),
            "!optimisticSeconds"
        );
        // check re-entrancy guard
        require(entered == 1, "!reentrant");
        entered = 0;
        _storeTips(messageView.tips());
        // update message status as executed
        mirror.setMessageStatus(messageHash, MirrorLib.MESSAGE_STATUS_EXECUTED);
        address recipient = _checkForSystemMessage(header.recipient());
        IMessageRecipient(recipient).handle(
            remoteDomain,
            header.nonce(),
            header.sender(),
            mirror.confirmAt[root],
            messageView.body().clone()
        );
        emit Executed(remoteDomain, messageHash);
        // reset re-entrancy guard
        entered = 1;
    }

    // ============ External Owner Functions ============

    /**
     * @notice Set Notary role
     * @dev MUST ensure that all roots signed by previous Notary have
     * been relayed before calling. Only callable by owner (Governance)
     * @param _notary New Notary
     */
    function setNotary(uint32 _domain, address _notary) external onlyOwner {
        // TODO: proper implementation
        _addNotary(_domain, _notary);
    }

    /**
     * @notice Set confirmAt for a given root
     * @dev To be used if in the case that fraud is proven
     * and roots need to be deleted / added. Only callable by owner (Governance)
     * @param _root The root for which to modify confirm time
     * @param _confirmAt The new confirmation time. Set to 0 to "delete" a root.
     */
    function setConfirmation(
        uint32 _remoteDomain,
        bytes32 _root,
        uint256 _confirmAt
    ) external onlyOwner {
        MirrorLib.Mirror storage mirror = allMirrors[activeMirrors[_remoteDomain]];
        uint256 _previousConfirmAt = mirror.confirmAt[_root];
        mirror.setConfirmAt(_root, _confirmAt);
        emit SetConfirmation(_remoteDomain, _root, _previousConfirmAt, _confirmAt);
    }

    // ============ Public Functions ============

    /**
     * @notice Check that the root has been submitted
     * and that the root's optimistic timeout period has expired,
     * meaning message proven against the root can be executed.
     * @param _root the Merkle root, submitted in an attestation, to check
     * @return TRUE iff root has been submitted & timeout has expired
     */
    function acceptableRoot(
        uint32 _remoteDomain,
        uint32 _optimisticSeconds,
        bytes32 _root
    ) public view returns (bool) {
        uint256 _time = allMirrors[activeMirrors[_remoteDomain]].confirmAt[_root];
        if (_time == 0) {
            return false;
        }
        return block.timestamp >= _time + _optimisticSeconds;
    }

    /**
     * @notice Attempts to prove the validity of message given its leaf, the
     * merkle proof of inclusion for the leaf, and the index of the leaf.
     * @dev Reverts if message's MessageStatus != None (i.e. if message was
     * already proven or executed)
     * @dev For convenience, we allow proving against any previous root.
     * This means that witnesses never need to be updated for the new root
     * @param _message Formatted message
     * @param _proof Merkle proof of inclusion for leaf
     * @param _index Index of leaf in origin's merkle tree
     * @return Returns true if proof was valid and `prove` call succeeded
     **/
    function prove(
        uint32 _remoteDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) public returns (bool) {
        bytes32 _leaf = keccak256(_message);
        MirrorLib.Mirror storage mirror = allMirrors[activeMirrors[_remoteDomain]];
        // ensure that mirror is active
        require(mirror.status == MirrorLib.MirrorStatus.Active, "Mirror not active");
        // ensure that message has not been proven or executed
        require(
            mirror.messageStatus[_leaf] == MirrorLib.MESSAGE_STATUS_NONE,
            "!MessageStatus.None"
        );
        // calculate the expected root based on the proof
        bytes32 _calculatedRoot = MerkleLib.branchRoot(_leaf, _proof, _index);
        // if the root is valid, save it for later optimistic period checking
        if (mirror.confirmAt[_calculatedRoot] != 0) {
            mirror.setMessageStatus(_leaf, _calculatedRoot);
            return true;
        }
        return false;
    }

    // ============ Internal Functions ============

    function _createMirror(uint32 _remoteDomain) internal returns (uint256 mirrorIndex) {
        // Start indexing from 1, so default mirror (allMirrors[0]) will be forever inactive
        unchecked {
            mirrorIndex = mirrorCount + 1;
        }
        allMirrors[mirrorIndex].setupMirror(_remoteDomain);
        mirrorCount = mirrorIndex;
    }

    function _getRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        // If the _res length is less than 68, then the transaction failed silently (without a revert message)
        if (_returnData.length < 68) return "Transaction reverted silently";

        assembly {
            // Slice the sighash.
            _returnData := add(_returnData, 0x04)
        }
        return abi.decode(_returnData, (string)); // All that remains is the revert string
    }

    function _checkForSystemMessage(bytes32 _recipient) internal view returns (address recipient) {
        // Check if SYSTEM_ROUTER was specified as message recipient
        if (_recipient == SystemMessage.SYSTEM_ROUTER) {
            /**
             * @dev Route message to SystemRouter.
             *      Note: Only SystemRouter contract on origin chain
             *      can send such a message (enforced in Origin.sol).
             */
            recipient = address(systemRouter);
        } else {
            // Cast bytes32 to address otherwise
            recipient = TypeCasts.bytes32ToAddress(_recipient);
        }
    }

    /**
     * @notice Applies submitted Report to blacklist reported Notary,
     * and all roots signed by this Notary. An honest Notary is incentivized to sign
     * a valid Attestation to collect tips from the pending messages,
     * which prevents downtime caused by root blacklisting.
     *
     * @dev Both Notary and Guard signatures
     * have been checked at this point (see ReportHub.sol).
     *
     * @param _guard            Guard address
     * @param _notary           Notary address
     * @param _attestationView  Memory view over reported Attestation
     * @param _reportView       Memory view over Report
     * @param _report           Payload with Report data and signature
     * @return blacklisted      TRUE if Notary was blacklisted as a result,
     *                          FALSE if Notary has been blacklisted earlier.
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool blacklisted) {
        require(_reportView.reportedFraud(), "Not a fraud report");
        blacklisted = _blacklistNotary(_attestationView.attestedDomain(), _notary);
        if (blacklisted) {
            emit NotaryBlacklisted(_notary, _guard, msg.sender, _report);
        }
    }

    function _storeTips(bytes29 _tips) internal virtual {
        // TODO: implement storing & claiming logic
    }

    function _blacklistNotary(uint32 _domain, address _notary) internal returns (bool blacklisted) {
        blacklisted = _isNotary(_domain, _notary);
        if (blacklisted) {
            // TODO: implement actual blacklisting for the roots
            // TODO: remove records about Notary, if it was active on other domains
            // assuming being a Notary for more than one domain is possible
            _removeNotary(_domain, _notary);
        }
    }
}
