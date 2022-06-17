// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Version0 } from "./Version0.sol";
import { ReplicaLib } from "./libs/Replica.sol";
import { MerkleLib } from "./libs/Merkle.sol";
import { Message } from "./libs/Message.sol";
// ============ External Imports ============
import { TypedMemView } from "./libs/TypedMemView.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title ReplicaManager
 * @notice Track root updates on Home,
 * prove and dispatch messages to end recipients.
 */
contract ReplicaManager is Version0, Initializable, OwnableUpgradeable {
    // ============ Libraries ============

    using ReplicaLib for ReplicaLib.Replica;
    using MerkleLib for MerkleLib.Tree;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Message for bytes29;

    // ============ Immutables ============

    // Domain of chain on which the contract is deployed
    uint32 public immutable localDomain;

    // Minimum gas for message processing
    uint256 public immutable PROCESS_GAS;
    // Reserved gas (to ensure tx completes in case message processing runs out)
    uint256 public immutable RESERVE_GAS;

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

    // Address of bonded Updater
    address public updater;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[44] private __GAP;

    // ============ Events ============

    /**
     * @notice Emitted when message is processed
     * @param messageHash The keccak256 hash of the message that was processed
     * @param success TRUE if the call was executed successfully,
     * FALSE if the call reverted or threw
     * @param returnData the return data from the external call
     */
    event Process(
        uint32 indexed remoteDomain,
        bytes32 indexed messageHash,
        bool indexed success,
        bytes returnData
    );

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

    /**
     * @notice Emitted when update is made on Home
     * or unconfirmed update root is submitted on Replica
     * @param homeDomain Domain of home contract
     * @param oldRoot Old merkle root
     * @param newRoot New merkle root
     * @param signature Updater's signature on `oldRoot` and `newRoot`
     */
    event Update(
        uint32 indexed homeDomain,
        bytes32 indexed oldRoot,
        bytes32 indexed newRoot,
        bytes signature
    );

    /**
     * @notice Emitted when Updater is rotated
     * @param oldUpdater The address of the old updater
     * @param newUpdater The address of the new updater
     */
    event NewUpdater(address oldUpdater, address newUpdater);

    // ============ Constructor ============

    constructor(
        uint32 _localDomain,
        uint256 _processGas,
        uint256 _reserveGas
    ) {
        localDomain = _localDomain;
        require(_processGas >= 850_000, "!process gas");
        require(_reserveGas >= 15_000, "!reserve gas");
        PROCESS_GAS = _processGas;
        RESERVE_GAS = _reserveGas;
    }

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
        __Ownable_init();
        _setUpdater(_updater);
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
        returns (ReplicaLib.MessageStatus)
    {
        return allReplicas[activeReplicas[_remoteDomain]].messages[_messageId];
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

    function _isUpdaterSignature(
        uint32 _remoteDomain,
        bytes32 _oldRoot,
        bytes32 _newRoot,
        bytes memory _signature
    ) internal view returns (bool) {
        bytes32 _digest = keccak256(
            abi.encodePacked(_homeDomainHash(_remoteDomain), _oldRoot, _newRoot)
        );
        _digest = ECDSA.toEthSignedMessageHash(_digest);
        return (ECDSA.recover(_digest, _signature) == updater);
    }

    /**
     * @notice First attempts to prove the validity of provided formatted
     * `message`. If the message is successfully proven, then tries to process
     * message.
     * @dev Reverts if `prove` call returns false
     * @param _message Formatted message (refer to SynapseBase.sol Message library)
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
     * or if not enough gas is provided for the dispatch transaction.
     * @param _message Formatted message
     * @return _success TRUE iff dispatch transaction succeeded
     */
    function process(bytes memory _message) public returns (bool _success) {
        bytes29 _m = _message.ref(0);
        uint32 _remoteDomain = _m.origin();
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure message was meant for this domain
        require(_m.destination() == localDomain, "!destination");
        // ensure message has been proven
        bytes32 _messageHash = _m.keccak();
        require(replica.messages[_messageHash] == ReplicaLib.MessageStatus.Proven, "!proven");
        // check re-entrancy guard
        require(entered == 1, "!reentrant");
        entered = 0;
        // update message status as processed
        replica.setMessageStatus(_messageHash, ReplicaLib.MessageStatus.Processed);
        // A call running out of gas TYPICALLY errors the whole tx. We want to
        // a) ensure the call has a sufficient amount of gas to make a
        //    meaningful state change.
        // b) ensure that if the subcall runs out of gas, that the tx as a whole
        //    does not revert (i.e. we still mark the message processed)
        // To do this, we require that we have enough gas to process
        // and still return. We then delegate only the minimum processing gas.
        require(gasleft() >= PROCESS_GAS + RESERVE_GAS, "!gas");
        // get the message recipient
        address _recipient = _m.recipientAddress();
        // set up for assembly call
        uint256 _toCopy;
        uint256 _maxCopy = 256;
        uint256 _gas = PROCESS_GAS;
        // allocate memory for returndata
        bytes memory _returnData = new bytes(_maxCopy);
        bytes memory _calldata = abi.encodeWithSignature(
            "handle(uint32,uint32,bytes32,bytes)",
            _m.origin(),
            _m.nonce(),
            _m.sender(),
            _m.optimisticSeconds(),
            _m.body().clone()
        );
        // dispatch message to recipient
        // by assembly calling "handle" function
        // we call via assembly to avoid memcopying a very large returndata
        // returned by a malicious contract
        assembly {
            _success := call(
                _gas, // gas
                _recipient, // recipient
                0, // ether value
                add(_calldata, 0x20), // inloc
                mload(_calldata), // inlen
                0, // outloc
                0 // outlen
            )
            // limit our copy to 256 bytes
            _toCopy := returndatasize()
            if gt(_toCopy, _maxCopy) {
                _toCopy := _maxCopy
            }
            // Store the length of the copied bytes
            mstore(_returnData, _toCopy)
            // copy the bytes from returndata[0:_toCopy]
            returndatacopy(add(_returnData, 0x20), 0, _toCopy)
        }
        // emit process results
        emit Process(_remoteDomain, _messageHash, _success, _returnData);
        // reset re-entrancy guard
        entered = 1;
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
        return block.timestamp + _optimisticSeconds >= _time;
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
        uint32 optimisticSeconds = _message.ref(0).optimisticSeconds();
        bytes32 _leaf = keccak256(_message);
        ReplicaLib.Replica storage replica = allReplicas[activeReplicas[_remoteDomain]];
        // ensure that message has not been proven or processed
        require(replica.messages[_leaf] == ReplicaLib.MessageStatus.None, "!MessageStatus.None");
        // calculate the expected root based on the proof
        bytes32 _calculatedRoot = MerkleLib.branchRoot(_leaf, _proof, _index);
        // if the root is valid, change status to Proven
        if (acceptableRoot(_remoteDomain, optimisticSeconds, _calculatedRoot)) {
            replica.setMessageStatus(_leaf, ReplicaLib.MessageStatus.Proven);
            return true;
        }
        return false;
    }

    /**
     * @notice Hash of Home domain concatenated with "SYN"
     * @param _homeDomain the Home domain to hash
     */
    function homeDomainHash(uint32 _homeDomain) public pure returns (bytes32) {
        return _homeDomainHash(_homeDomain);
    }

    // ============ Internal Functions ============

    function _createReplica(uint32 _remoteDomain) internal returns (uint256 replicaIndex) {
        replicaIndex = replicaCount;
        allReplicas[replicaIndex].setupReplica(_remoteDomain);
        unchecked {
            replicaCount = replicaIndex + 1;
        }
    }

    /**
     * @notice Hash of Home domain concatenated with "SYN"
     * @param _homeDomain the Home domain to hash
     */
    function _homeDomainHash(uint32 _homeDomain) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_homeDomain, "SYN"));
    }

    /**
     * @notice Set the Updater
     * @param _newUpdater Address of the new Updater
     */
    function _setUpdater(address _newUpdater) internal {
        address _oldUpdater = updater;
        updater = _newUpdater;
        emit NewUpdater(_oldUpdater, _newUpdater);
    }

    /// @notice Hook for potential future use
    // solhint-disable-next-line no-empty-blocks
    function _beforeUpdate() internal {}
}
