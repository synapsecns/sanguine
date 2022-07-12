// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { UpdaterStorage } from "./UpdaterStorage.sol";
import { Version0 } from "./Version0.sol";
import { ReplicaLib } from "./libs/Replica.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
import { IMessageRecipient } from "./interfaces/IMessageRecipient.sol";
// ============ External Imports ============
import { TypedMemView } from "./libs/TypedMemView.sol";

/**
 * @title ReplicaManager
 * @notice Track root updates on Home,
 * prove and dispatch messages to end recipients.
 */
contract ReplicaManager is Version0, UpdaterStorage {
    // ============ Libraries ============

    using ReplicaLib for ReplicaLib.Replica;
    using MerkleLib for MerkleLib.Tree;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

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

    // ============ Constructor ============

    constructor(uint32 _localDomain) UpdaterStorage(_localDomain) {}

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
     * @param _updater The EVM id of the updater
     */
    function initialize(uint32 _remoteDomain, address _updater) public initializer {
        __SynapseBase_initialize(_updater);
        // set storage variables
        entered = 1;
        activeReplicas[_remoteDomain] = _createReplica(_remoteDomain);
    }

    // ============ Active Replica Views ============

    function activeReplicaCommittedRoot(uint32 _remoteDomain) external view returns (bytes32) {
        return allReplicas[activeReplicas[_remoteDomain]].committedRoot;
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
     * @notice Called by external agent. Submits the signed update's new root,
     * marks root's allowable confirmation time, and emits an `Update` event.
     * @dev Reverts if update doesn't build off latest committedRoot
     * or if signature is invalid.
     * @param _oldRoot Old merkle root
     * @param _newRoot New merkle root
     * @param _signature Updater's signature on `_oldRoot` and `_newRoot` = `keccak256(message, optimisticSeconds)`
     */
    function update(
        uint32 _remoteDomain,
        bytes32 _oldRoot,
        bytes32 _newRoot,
        bytes memory _signature
    ) external {
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure that update is building off the last submitted root
        require(_oldRoot == replica.committedRoot, "not current update");
        // validate updater signature
        require(_isUpdaterSignature(_remoteDomain, _oldRoot, _newRoot, _signature), "!updater sig");
        // Hook for future use
        _beforeUpdate();
        // set the new root's confirmation timer
        replica.setConfirmAt(_newRoot, block.timestamp);
        // update committedRoot
        replica.setCommittedRoot(_newRoot);
        emit Update(_remoteDomain, _oldRoot, _newRoot, _signature);
    }

    /**
     * @notice First attempts to prove the validity of provided formatted
     * `message`. If the message is successfully proven, then tries to process
     * message.
     * @dev Reverts if `prove` call returns false
     * @param _message Formatted message (refer to UpdaterStorage.sol Message library)
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
        bytes29 _m = _message.ref(0);
        uint32 _remoteDomain = _m.origin();
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure message was meant for this domain
        require(_m.destination() == localDomain, "!destination");
        // ensure message has been proven
        bytes32 _messageHash = _m.keccak();
        bytes32 _root = replica.messageStatus[_messageHash];
        require(ReplicaLib.isPotentialRoot(_root), "!exists || processed");
        require(acceptableRoot(_remoteDomain, _m.optimisticSeconds(), _root), "!optimisticSeconds");
        // check re-entrancy guard
        require(entered == 1, "!reentrant");
        entered = 0;
        // update message status as processed
        replica.setMessageStatus(_messageHash, ReplicaLib.MESSAGE_STATUS_PROCESSED);
        address recipient = _m.recipientAddress();
        IMessageRecipient(recipient).handle(
            _remoteDomain,
            _m.nonce(),
            _m.sender(),
            replica.confirmAt[_root],
            _m.body().clone()
        );
        emit Process(_remoteDomain, _messageHash);
        // reset re-entrancy guard
        entered = 1;
    }

    /**
     * @notice Hash of Home domain concatenated with "SYN"
     * @param _homeDomain the Home domain to hash
     */
    function homeDomainHash(uint32 _homeDomain) external pure returns (bytes32) {
        return _domainHash(_homeDomain);
    }

    // ============ External Owner Functions ============

    /**
     * @notice Set Updater role
     * @dev MUST ensure that all roots signed by previous Updater have
     * been relayed before calling. Only callable by owner (Governance)
     * @param _updater New Updater
     */
    function setUpdater(address _updater) external onlyOwner {
        _setUpdater(_updater);
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
     * @param _root the Merkle root, submitted in an update, to check
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
        replicaIndex = replicaCount;
        allReplicas[replicaIndex].setupReplica(_remoteDomain);
    unchecked {
        replicaCount = replicaIndex + 1;
    }
    }

    /// @notice Hook for potential future use
    // solhint-disable-next-line no-empty-blocks
    function _beforeUpdate() internal {}

    function _getRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        // If the _res length is less than 68, then the transaction failed silently (without a revert message)
        if (_returnData.length < 68) return "Transaction reverted silently";

        assembly {
        // Slice the sighash.
            _returnData := add(_returnData, 0x04)
        }
        return abi.decode(_returnData, (string)); // All that remains is the revert string
    }
}
