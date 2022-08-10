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
    //   contract will no longer accept attestations or new messages
    enum ReplicaStatus {
        UnInitialized,
        Active,
        Failed
    }

    // ============ Constants ============
    /// @dev Should not be possible to have 0x0 or 0x1 as valid Merkle root,
    /// so it's safe to use those values as NONE/PROCESSED
    bytes32 public constant MESSAGE_STATUS_NONE = bytes32(0);
    bytes32 public constant MESSAGE_STATUS_PROCESSED = bytes32(uint256(1));

    // TODO: optimize read/writes by further packing?
    struct Replica {
        // The latest nonce that has been signed by the Notary for this given Replica
        uint32 nonce; // 32 bits
        // Domain of home chain
        uint32 remoteDomain; // 32 bits
        // Status of Replica based on the Home remote domain
        ReplicaStatus status; // 8 bits
        // Mapping of roots to time at which Broadcaster submitted on-chain. Latency period begins here.
        // TODO: confirmAt doesn't need to be uint256 necessarily
        mapping(bytes32 => uint256) confirmAt;
        // Mapping of message leaves to status:
        // - NONE: message not yet submitted
        // - PROCESSED: message was proven and processed
        // bytes32 root: message was proven against `root`, but not yet processed
        mapping(bytes32 => bytes32) messageStatus;
    }

    function setupReplica(Replica storage replica, uint32 _remoteDomain) internal {
        replica.remoteDomain = _remoteDomain;
        replica.status = ReplicaStatus.Active;
    }

    function setNonce(Replica storage replica, uint32 _nonce) internal {
        replica.nonce = _nonce;
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
        bytes32 _status
    ) internal {
        replica.messageStatus[_messageHash] = _status;
    }

    function setStatus(Replica storage replica, ReplicaStatus _status) internal {
        replica.status = _status;
    }

    function isPotentialRoot(bytes32 messageStatus) internal pure returns (bool) {
        return messageStatus != MESSAGE_STATUS_NONE && messageStatus != MESSAGE_STATUS_PROCESSED;
    }
}
