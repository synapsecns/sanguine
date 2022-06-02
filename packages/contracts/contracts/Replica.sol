// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

contract ReplicaStorage is Ownable {
    address replicaManager;
    // ============ Enums ============
    // Status of Message:
    //   0 - None - message has not been proven or processed
    //   1 - Proven - message inclusion proof has been validated
    //   2 - Processed - message has been dispatched to recipient
    enum MessageStatus {
        None,
        Proven,
        Processed
    }

    // States:
    //   0 - UnInitialized - before initialize function is called
    //   note: the contract is initialized at deploy time, so it should never be in this state
    //   1 - Active - as long as the contract has not become fraudulent
    //   2 - Failed - after a valid fraud proof has been submitted;
    //   contract will no longer accept updates or new messages
    enum ReplicaStatus {
        UnInitialized,
        Active,
        Failed
    }

    // The latest root that has been signed by the Updater for this given Replica
    bytes32 public committedRoot; // 256 bits
    // Domain of home chain
    //TODO: Could/should be immutable
    uint32 public remoteDomain;
    // Optimistic seconds per remote domain  (E.g specifies optimistic seconds on a remote domain basis to wait)
    uint256 public optimisticSeconds;
    // Status of Replica based on the Home remote domain
    ReplicaStatus public status;
    // Mapping of roots to allowable confirmation times
    // TODO: confirmAt doesn't need to be uint256 necessarily
    mapping(bytes32 => uint256) public confirmAt;
    // Mapping of message leaves to MessageStatus
    mapping(bytes32 => MessageStatus) public messages;

    constructor(
        address _replicaManager,
        uint32 _remoteDomain,
        uint256 _optimisticSeconds
    ) {
        replicaManager = _replicaManager;
        remoteDomain = _remoteDomain;
        optimisticSeconds = _optimisticSeconds;
        committedRoot = bytes32("");
        confirmAt[bytes32("")] = 0;
        status = ReplicaStatus.Active;
    }

    modifier onlyReplicaManager() {
        require(msg.sender == address(replicaManager), "!replica");
        _;
    }

    function setCommittedRoot(bytes32 _committedRoot) public onlyReplicaManager {
        committedRoot = _committedRoot;
    }

    function setConfirmAt(bytes32 _root, uint256 _confirmAt) public onlyReplicaManager {
        confirmAt[_root] = _confirmAt;
    }

    function setMessageStatus(bytes32 _messageHash, MessageStatus _status)
        public
        onlyReplicaManager
    {
        messages[_messageHash] = _status;
    }

    function setOptimisticTimeout(uint256 _optimisticSeconds) public onlyReplicaManager {
        optimisticSeconds = _optimisticSeconds;
    }

    function setStatus(ReplicaStatus _status) public onlyReplicaManager {
        status = _status;
    }

    function setReplica(address _newReplicaManager) public onlyOwner {
        replicaManager = _newReplicaManager;
    }
}
