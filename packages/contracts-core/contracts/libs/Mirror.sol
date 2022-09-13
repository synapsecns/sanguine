// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

library MirrorLib {
    // ============ Enums ============
    // Status of Message:
    //   0 - None - message has not been proven or executed
    //   1 - Proven - message inclusion proof has been validated
    //   2 - Executed - message has been dispatched to recipient
    enum MessageStatus {
        None,
        Proven,
        Executed
    }

    // States:
    //   0 - UnInitialized - before initialize function is called
    //   note: the contract is initialized at deploy time, so it should never be in this state
    //   1 - Active - as long as the contract has not become fraudulent
    //   2 - Failed - after a valid fraud proof has been submitted;
    //   contract will no longer accept attestations or new messages
    enum MirrorStatus {
        UnInitialized,
        Active,
        Failed
    }

    // ============ Constants ============
    /// @dev Should not be possible to have 0x0 or 0x1 as valid Merkle root,
    /// so it's safe to use those values as NONE/EXECUTED
    bytes32 public constant MESSAGE_STATUS_NONE = bytes32(0);
    bytes32 public constant MESSAGE_STATUS_EXECUTED = bytes32(uint256(1));

    // TODO: optimize read/writes by further packing?
    struct Mirror {
        // The latest nonce that has been signed by the Notary for this given Mirror
        uint32 nonce; // 32 bits
        // Domain of origin chain
        uint32 remoteDomain; // 32 bits
        // Status of Mirror based on the Origin remote domain
        MirrorStatus status; // 8 bits
        // Mapping of roots to time at which Broadcaster submitted on-chain. Latency period begins here.
        // TODO: confirmAt doesn't need to be uint256 necessarily
        mapping(bytes32 => uint256) confirmAt;
        // Mapping of message leaves to status:
        // - NONE: message not yet submitted
        // - EXECUTED: message was proven and executed
        // bytes32 root: message was proven against `root`, but not yet executed
        mapping(bytes32 => bytes32) messageStatus;
    }

    function setupMirror(Mirror storage mirror, uint32 _remoteDomain) internal {
        mirror.remoteDomain = _remoteDomain;
        mirror.status = MirrorStatus.Active;
    }

    function setNonce(Mirror storage mirror, uint32 _nonce) internal {
        mirror.nonce = _nonce;
    }

    function setConfirmAt(
        Mirror storage mirror,
        bytes32 _root,
        uint256 _confirmAt
    ) internal {
        mirror.confirmAt[_root] = _confirmAt;
    }

    function setMessageStatus(
        Mirror storage mirror,
        bytes32 _messageHash,
        bytes32 _status
    ) internal {
        mirror.messageStatus[_messageHash] = _status;
    }

    function setStatus(Mirror storage mirror, MirrorStatus _status) internal {
        mirror.status = _status;
    }

    function isPotentialRoot(bytes32 messageStatus) internal pure returns (bool) {
        return messageStatus != MESSAGE_STATUS_NONE && messageStatus != MESSAGE_STATUS_EXECUTED;
    }
}
