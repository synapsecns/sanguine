// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "./MemView.sol";
import {RECEIPT_LENGTH, RECEIPT_VALID_SALT, RECEIPT_INVALID_SALT} from "../Constants.sol";
import {UnformattedReceipt} from "../Errors.sol";

/// Receipt is a memory view over a formatted "full receipt" payload.
type Receipt is uint256;

using ReceiptLib for Receipt global;

/// Receipt structure represents a Notary statement that a certain message has been executed in `ExecutionHub`.
/// - It is possible to prove the correctness of the tips payload using the message hash, therefore tips are not
///   included in the receipt.
/// - Receipt is signed by a Notary and submitted to `Summit` in order to initiate the tips distribution for an
///   executed message.
/// - If a message execution fails the first time, the `finalExecutor` field will be set to zero address. In this
///   case, when the message is finally executed successfully, the `finalExecutor` field will be updated. Both
///   receipts will be considered valid.
/// # Memory layout of Receipt fields
///
/// | Position   | Field         | Type    | Bytes | Description                                      |
/// | ---------- | ------------- | ------- | ----- | ------------------------------------------------ |
/// | [000..004) | origin        | uint32  | 4     | Domain where message originated                  |
/// | [004..008) | destination   | uint32  | 4     | Domain where message was executed                |
/// | [008..040) | messageHash   | bytes32 | 32    | Hash of the message                              |
/// | [040..072) | snapshotRoot  | bytes32 | 32    | Snapshot root used for proving the message       |
/// | [072..073) | stateIndex    | uint8   | 1     | Index of state used for the snapshot proof       |
/// | [073..093) | attNotary     | address | 20    | Notary who posted attestation with snapshot root |
/// | [093..113) | firstExecutor | address | 20    | Executor who performed first valid execution     |
/// | [113..133) | finalExecutor | address | 20    | Executor who successfully executed the message   |
library ReceiptLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_ORIGIN = 0;
    uint256 private constant OFFSET_DESTINATION = 4;
    uint256 private constant OFFSET_MESSAGE_HASH = 8;
    uint256 private constant OFFSET_SNAPSHOT_ROOT = 40;
    uint256 private constant OFFSET_STATE_INDEX = 72;
    uint256 private constant OFFSET_ATT_NOTARY = 73;
    uint256 private constant OFFSET_FIRST_EXECUTOR = 93;
    uint256 private constant OFFSET_FINAL_EXECUTOR = 113;

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Receipt payload with provided fields.
     * @param origin_           Domain where message originated
     * @param destination_      Domain where message was executed
     * @param messageHash_      Hash of the message
     * @param snapshotRoot_     Snapshot root used for proving the message
     * @param stateIndex_       Index of state used for the snapshot proof
     * @param attNotary_        Notary who posted attestation with snapshot root
     * @param firstExecutor_    Executor who performed first valid execution attempt
     * @param finalExecutor_    Executor who successfully executed the message
     * @return Formatted receipt
     */
    function formatReceipt(
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        uint8 stateIndex_,
        address attNotary_,
        address firstExecutor_,
        address finalExecutor_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_
        );
    }

    /**
     * @notice Returns a Receipt view over the given payload.
     * @dev Will revert if the payload is not a receipt.
     */
    function castToReceipt(bytes memory payload) internal pure returns (Receipt) {
        return castToReceipt(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Receipt view.
     * @dev Will revert if the memory view is not over a receipt.
     */
    function castToReceipt(MemView memView) internal pure returns (Receipt) {
        if (!isReceipt(memView)) revert UnformattedReceipt();
        return Receipt.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Receipt.
    function isReceipt(MemView memView) internal pure returns (bool) {
        // Check payload length
        return memView.len() == RECEIPT_LENGTH;
    }

    /// @notice Returns the hash of an Receipt, that could be later signed by a Notary to signal
    /// that the receipt is valid.
    function hashValid(Receipt receipt) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptSalt, keccak(receipt))
        return receipt.unwrap().keccakSalted(RECEIPT_VALID_SALT);
    }

    /// @notice Returns the hash of a Receipt, that could be later signed by a Guard to signal
    /// that the receipt is invalid.
    function hashInvalid(Receipt receipt) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptBodyInvalidSalt, keccak(receipt))
        return receipt.unwrap().keccakSalted(RECEIPT_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Receipt receipt) internal pure returns (MemView) {
        return MemView.wrap(Receipt.unwrap(receipt));
    }

    /// @notice Compares two Receipt structures.
    function equals(Receipt a, Receipt b) internal pure returns (bool) {
        // Length of a Receipt payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    // ═════════════════════════════════════════════ RECEIPT SLICING ═════════════════════════════════════════════════

    /// @notice Returns receipt's origin field
    function origin(Receipt receipt) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(receipt.unwrap().indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns receipt's destination field
    function destination(Receipt receipt) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(receipt.unwrap().indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(Receipt receipt) internal pure returns (bytes32) {
        return receipt.unwrap().index({index_: OFFSET_MESSAGE_HASH, bytes_: 32});
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(Receipt receipt) internal pure returns (bytes32) {
        return receipt.unwrap().index({index_: OFFSET_SNAPSHOT_ROOT, bytes_: 32});
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(Receipt receipt) internal pure returns (uint8) {
        // Can be safely casted to uint8, since we index a single byte
        return uint8(receipt.unwrap().indexUint({index_: OFFSET_STATE_INDEX, bytes_: 1}));
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_ATT_NOTARY});
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }
}
