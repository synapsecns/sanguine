// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { GlobalNotaryRegistry } from "./registry/GlobalNotaryRegistry.sol";
import { GuardRegistry } from "./registry/GuardRegistry.sol";
import { Attestation } from "./libs/Attestation.sol";
import { Version0 } from "./Version0.sol";
import { ReplicaLib } from "./libs/Replica.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { Header } from "./libs/Header.sol";
import { Tips } from "./libs/Tips.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
import { SystemMessage } from "./system/SystemMessage.sol";
import { SystemContract } from "./system/SystemContract.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
// ============ External Imports ============
import { TypedMemView } from "./libs/TypedMemView.sol";

/**
 * @title ReplicaManager
 * @notice Track merkle root state of Home contracts on other chains,
 * prove and dispatch messages to end recipients.
 */
contract ReplicaManager is Version0, SystemContract, GlobalNotaryRegistry, GuardRegistry {
    // ============ Libraries ============

    using ReplicaLib for ReplicaLib.Replica;
    using MerkleLib for MerkleLib.Tree;
    using Message for bytes;
    using TypedMemView for bytes29;
    using Attestation for bytes29;
    using Message for bytes29;
    using Header for bytes29;

    // ============ Public Storage ============

    // re-entrancy guard
    uint8 private entered;

    uint256 internal replicaCount;

    // all Replicas: both active and archived
    mapping(uint256 => ReplicaLib.Replica) internal allReplicas;

    // (domain => replica index): index of the active replica in allReplicas
    mapping(uint32 => uint256) internal activeReplicas;

    //TODO: Handle fail-over replicas and modify activeReplicas
    // (domain => [replica indexes]): array of indexes of archived replicas in allReplicas
    mapping(uint32 => uint256[]) internal archivedReplicas;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[45] private __GAP;

    // ============ Events ============

    /**
     * @notice Emitted when message is processed
     * @param messageHash The keccak256 hash of the message that was processed
     */
    event Process(uint32 indexed remoteDomain, bytes32 indexed messageHash);

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
        uint32 indexed homeDomain,
        uint32 indexed nonce,
        bytes32 indexed root,
        bytes signature
    );

    // ============ Constructor ============

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _localDomain) SystemContract(_localDomain) {}

    // ============ Initializer ============

    /**
     * @notice Initialize the replica
     * @dev Performs the following action:
     *      - initializes inherited contracts
     *      - initializes re-entrancy guard
     *      - sets remote domain
     *      - sets a trusted root, and pre-approves messages under it
     *      - sets the optimistic timer
     * @param _remoteDomain The domain of the Home contract this follows
     * @param _notary The EVM id of the notary
     */
    function initialize(uint32 _remoteDomain, address _notary) public initializer {
        __SystemContract_initialize();
        _addNotary(_remoteDomain, _notary);
        // set storage variables
        entered = 1;
        activeReplicas[_remoteDomain] = _createReplica(_remoteDomain);
    }

    // ============ Active Replica Views ============

    function activeReplicaNonce(uint32 _remoteDomain) external view returns (uint32) {
        return allReplicas[activeReplicas[_remoteDomain]].nonce;
    }

    function activeReplicaConfirmedAt(uint32 _remoteDomain, bytes32 _root)
        external
        view
        returns (uint256)
    {
        return allReplicas[activeReplicas[_remoteDomain]].confirmAt[_root];
    }

    function activeReplicaMessageStatus(uint32 _remoteDomain, bytes32 _messageId)
        external
        view
        returns (bytes32)
    {
        return allReplicas[activeReplicas[_remoteDomain]].messageStatus[_messageId];
    }

    // ============ Archived Replica Views ============

    // TODO: getters for archived replicas

    // ============ External Functions ============

    /**
     * @notice Called by external agent. Submits the signed attestation,
     * marks root's allowable confirmation time, and emits an `AttestationAccepted` event.
     * @dev Reverts if signature is invalid.
     * @param _attestation  Attestation data and signature
     */
    function submitAttestation(bytes memory _attestation) external {
        (, bytes29 _view) = _checkNotaryAuth(_attestation);
        uint32 remoteDomain = _view.attestationDomain();
        require(remoteDomain != localDomain, "Attestation refers to local chain");
        uint32 nonce = _view.attestationNonce();
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[remoteDomain]];
        require(nonce > replica.nonce, "Attestation older than current state");
        bytes32 newRoot = _view.attestationRoot();
        replica.setConfirmAt(newRoot, block.timestamp);
        // update nonce
        replica.setNonce(nonce);
        emit AttestationAccepted(
            remoteDomain,
            nonce,
            newRoot,
            _view.attestationSignature().clone()
        );
    }

    /**
     * @notice First attempts to prove the validity of provided formatted
     * `message`. If the message is successfully proven, then tries to process
     * message.
     * @dev Reverts if `prove` call returns false
     * @param _message Formatted message (refer to Message library)
     * @param _proof Merkle proof of inclusion for message's leaf
     * @param _index Index of leaf in home's merkle tree
     */
    function proveAndProcess(
        uint32 _remoteDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) external {
        require(prove(_remoteDomain, _message, _proof, _index), "!prove");
        process(_message);
    }

    /**
     * @notice Given formatted message, attempts to dispatch
     * message payload to end recipient.
     * @dev Recipient must implement a `handle` method (refer to IMessageRecipient.sol)
     * Reverts if formatted message's destination domain is not the Replica's domain,
     * if message has not been proven,
     * or if recipient reverted upon receiving the message.
     * @param _message Formatted message
     */
    function process(bytes memory _message) public {
        bytes29 _m = _message.messageView();
        bytes29 _header = _m.header();
        uint32 _remoteDomain = _header.origin();
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure message was meant for this domain
        require(_header.destination() == localDomain, "!destination");
        // ensure message has been proven
        bytes32 _messageHash = _m.keccak();
        bytes32 _root = replica.messageStatus[_messageHash];
        require(ReplicaLib.isPotentialRoot(_root), "!exists || processed");
        require(
            acceptableRoot(_remoteDomain, _header.optimisticSeconds(), _root),
            "!optimisticSeconds"
        );
        // check re-entrancy guard
        require(entered == 1, "!reentrant");
        entered = 0;
        _storeTips(_m.tips());
        // update message status as processed
        replica.setMessageStatus(_messageHash, ReplicaLib.MESSAGE_STATUS_PROCESSED);
        address recipient = _checkForSystemMessage(_header.recipient());
        IMessageRecipient(recipient).handle(
            _remoteDomain,
            _header.nonce(),
            _header.sender(),
            replica.confirmAt[_root],
            _m.body().clone()
        );
        emit Process(_remoteDomain, _messageHash);
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
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        uint256 _previousConfirmAt = replica.confirmAt[_root];
        replica.setConfirmAt(_root, _confirmAt);
        emit SetConfirmation(_remoteDomain, _root, _previousConfirmAt, _confirmAt);
    }

    // ============ Public Functions ============

    /**
     * @notice Check that the root has been submitted
     * and that the optimistic timeout period has expired,
     * meaning the root can be processed
     * @param _root the Merkle root, submitted in an attestation, to check
     * @return TRUE iff root has been submitted & timeout has expired
     */
    function acceptableRoot(
        uint32 _remoteDomain,
        uint32 _optimisticSeconds,
        bytes32 _root
    ) public view returns (bool) {
        uint256 _time = allReplicas[activeReplicas[_remoteDomain]].confirmAt[_root];
        if (_time == 0) {
            return false;
        }
        return block.timestamp >= _time + _optimisticSeconds;
    }

    /**
     * @notice Attempts to prove the validity of message given its leaf, the
     * merkle proof of inclusion for the leaf, and the index of the leaf.
     * @dev Reverts if message's MessageStatus != None (i.e. if message was
     * already proven or processed)
     * @dev For convenience, we allow proving against any previous root.
     * This means that witnesses never need to be updated for the new root
     * @param _message Formatted message
     * @param _proof Merkle proof of inclusion for leaf
     * @param _index Index of leaf in home's merkle tree
     * @return Returns true if proof was valid and `prove` call succeeded
     **/
    function prove(
        uint32 _remoteDomain,
        bytes memory _message,
        bytes32[32] calldata _proof,
        uint256 _index
    ) public returns (bool) {
        bytes32 _leaf = keccak256(_message);
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure that replica is active
        require(replica.status == ReplicaLib.ReplicaStatus.Active, "Replica not active");
        // ensure that message has not been proven or processed
        require(
            replica.messageStatus[_leaf] == ReplicaLib.MESSAGE_STATUS_NONE,
            "!MessageStatus.None"
        );
        // calculate the expected root based on the proof
        bytes32 _calculatedRoot = MerkleLib.branchRoot(_leaf, _proof, _index);
        // if the root is valid, save it for later optimistic period checking
        if (replica.confirmAt[_calculatedRoot] != 0) {
            replica.setMessageStatus(_leaf, _calculatedRoot);
            return true;
        }
        return false;
    }

    // ============ Internal Functions ============

    function _createReplica(uint32 _remoteDomain) internal returns (uint256 replicaIndex) {
        // Start indexing from 1, so default replica (allReplicas[0]) will be forever inactive
        unchecked {
            replicaIndex = replicaCount + 1;
        }
        allReplicas[replicaIndex].setupReplica(_remoteDomain);
        replicaCount = replicaIndex;
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
        // Check if SYSTEM_SENDER was specified as message recipient
        if (_recipient == SystemMessage.SYSTEM_SENDER) {
            /**
             * @dev Route message to SystemMessenger.
             *      Note: Only SystemMessenger contract on origin chain
             *      can send such a message (enforced in Home.sol).
             */
            recipient = address(systemMessenger);
        } else {
            // Cast bytes32 to address otherwise
            recipient = TypeCasts.bytes32ToAddress(_recipient);
        }
    }

    function _storeTips(bytes29 _tips) internal virtual {
        // TODO: implement storing & claiming logic
    }
}
