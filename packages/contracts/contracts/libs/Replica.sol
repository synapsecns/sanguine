// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

library ReplicaLib {
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

    // TODO: optimize read/writes by further packing?
    struct Replica {
        // The latest root that has been signed by the Updater for this given Replica
        bytes32 committedRoot; // 256 bits
        // Domain of home chain
        uint256 optimisticSeconds;
        // Status of Replica based on the Home remote domain
        uint32 remoteDomain;
        // Optimistic seconds per remote domain  (E.g specifies optimistic seconds on a remote domain basis to wait)
        ReplicaStatus status;
        // Mapping of roots to allowable confirmation times
        // TODO: confirmAt doesn't need to be uint256 necessarily
        mapping(bytes32 => uint256) confirmAt;
        // Mapping of message leaves to MessageStatus
        mapping(bytes32 => MessageStatus) messages;
    }

    function setupReplica(
        Replica storage replica,
        uint32 _remoteDomain,
        uint256 _optimisticSeconds
    ) internal {
        replica.remoteDomain = _remoteDomain;
        replica.optimisticSeconds = _optimisticSeconds;
        replica.status = ReplicaStatus.Active;
    }

    function setCommittedRoot(Replica storage replica, bytes32 _committedRoot) internal {
        replica.committedRoot = _committedRoot;
    }

    function setConfirmAt(
        Replica storage replica,
        bytes32 _root,
        uint256 _confirmAt
    ) internal {
        replica.confirmAt[_root] = _confirmAt;
    }

    function setMessageStatus(
        Replica storage replica,
        bytes32 _messageHash,
        MessageStatus _status
    ) internal {
        replica.messages[_messageHash] = _status;
    }

    function setOptimisticTimeout(Replica storage replica, uint256 _optimisticSeconds) internal {
        replica.optimisticSeconds = _optimisticSeconds;
    }

    function setStatus(Replica storage replica, ReplicaStatus _status) internal {
        replica.status = _status;
    }
}
