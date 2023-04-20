// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_SALT, RECEIPT_BODY_LENGTH, RECEIPT_LENGTH, TIPS_LENGTH} from "./Constants.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// Receipt is a memory view over a formatted "full receipt" payload.
type Receipt is uint256;

using ReceiptLib for Receipt global;

/// ReceiptBody is a memory view over a formatted "receipt body" payload.
type ReceiptBody is uint256;

using ReceiptLib for ReceiptBody global;

/// Receipt structure represents a Notary statement that a certain message has been executed in `ExecutionHub`.
/// Receipt contains two fields:
/// - Receipt body, which is stored in `ExecutionHub` for every executed message
/// - Receipt tips, which are emitted whenever a message is executed.
/// Note: It is possible to prove the correctness of the tips payload using the message hash.
/// # Memory layout of ReceiptBody
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
///
/// # Memory layout of Receipt fields
///
/// | Position   | Field | Type   | Bytes | Description                       |
/// | ---------- | ----- | ------ | ----- | --------------------------------- |
/// | [000..133) | body  | bytes  | 133   | Receipt body (see above)          |
/// | [133..165) | tips  | uint32 | 32    | Encoded tips paid on origin chain |

/// @dev Receipt could be signed by a Notary and submitted to `Summit` in order to initiate the tips
/// distribution for an executed message.
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

    uint256 private constant OFFSET_BODY = 0;
    uint256 private constant OFFSET_TIPS = RECEIPT_BODY_LENGTH;

    // ═══════════════════════════════════════════════ RECEIPT BODY ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted ReceiptBody payload with provided fields
     * @param origin_           Domain where message originated
     * @param destination_      Domain where message was executed
     * @param messageHash_      Hash of the message
     * @param snapshotRoot_     Snapshot root used for proving the message
     * @param stateIndex_       Index of state used for the snapshot proof
     * @param attNotary_        Notary who posted attestation with snapshot root
     * @param firstExecutor_    Executor who performed first valid execution attempt
     * @param finalExecutor_    Executor who successfully executed the message
     * @return Formatted receipt body
     */
    function formatReceiptBody(
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
     * @notice Returns a ReceiptBody view over the given payload.
     * @dev Will revert if the payload is not a receipt body.
     */
    function castToReceiptBody(bytes memory payload) internal pure returns (ReceiptBody) {
        return castToReceiptBody(payload.ref());
    }

    /**
     * @notice Casts a memory view to a ReceiptBody view.
     * @dev Will revert if the memory view is not over a receipt body.
     */
    function castToReceiptBody(MemView memView) internal pure returns (ReceiptBody) {
        require(isReceiptBody(memView), "Not a receipt body");
        return ReceiptBody.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted ReceiptBody.
    function isReceiptBody(MemView memView) internal pure returns (bool) {
        // Check payload length
        return memView.len() == RECEIPT_BODY_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(ReceiptBody receiptBody) internal pure returns (MemView) {
        return MemView.wrap(ReceiptBody.unwrap(receiptBody));
    }

    /// @notice Compares two ReceiptBody structures.
    function equals(ReceiptBody a, ReceiptBody b) internal pure returns (bool) {
        // Length of a ReceiptBody payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Receipt payload with provided fields
     * @param bodyPayload       Formatted payload with receipt body
     * @param tips_             Encoded tips information
     * @return Formatted receipt
     */
    function formatReceipt(bytes memory bodyPayload, Tips tips_) internal pure returns (bytes memory) {
        return abi.encodePacked(bodyPayload, tips_);
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
        require(isReceipt(memView), "Not a receipt");
        return Receipt.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Receipt.
    function isReceipt(MemView memView) internal pure returns (bool) {
        // Check payload length
        if (memView.len() != RECEIPT_LENGTH) return false;
        // Check that body payload is formatted
        return isReceiptBody(_body(memView));
    }

    /// @notice Returns the hash of an Receipt, that could be later signed by a Notary.
    function hash(Receipt receipt) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptSalt, keccak(receipt))
        return receipt.unwrap().keccakSalted(RECEIPT_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Receipt receipt) internal pure returns (MemView) {
        return MemView.wrap(Receipt.unwrap(receipt));
    }

    // ═══════════════════════════════════════════ RECEIPT BODY SLICING ════════════════════════════════════════════════

    /// @notice Returns receipt's origin field
    function origin(ReceiptBody receiptBody) internal pure returns (uint32) {
        return uint32(unwrap(receiptBody).indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns receipt's destination field
    function destination(ReceiptBody receiptBody) internal pure returns (uint32) {
        return uint32(unwrap(receiptBody).indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(ReceiptBody receiptBody) internal pure returns (bytes32) {
        return unwrap(receiptBody).index({index_: OFFSET_MESSAGE_HASH, bytes_: 32});
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(ReceiptBody receiptBody) internal pure returns (bytes32) {
        return unwrap(receiptBody).index({index_: OFFSET_SNAPSHOT_ROOT, bytes_: 32});
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(ReceiptBody receiptBody) internal pure returns (uint8) {
        return uint8(unwrap(receiptBody).indexUint({index_: OFFSET_STATE_INDEX, bytes_: 1}));
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(ReceiptBody receiptBody) internal pure returns (address) {
        return unwrap(receiptBody).indexAddress({index_: OFFSET_ATT_NOTARY});
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(ReceiptBody receiptBody) internal pure returns (address) {
        return unwrap(receiptBody).indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(ReceiptBody receiptBody) internal pure returns (address) {
        return unwrap(receiptBody).indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }

    // ═════════════════════════════════════════════ RECEIPT SLICING ═════════════════════════════════════════════════

    /// @notice Returns a typed memory view over the payload with receipt body.
    function body(Receipt receipt) internal pure returns (ReceiptBody) {
        // We check that body payload is properly formatted, when the whole payload is wrapped
        // into Receipt, so this never reverts.
        return castToReceiptBody(_body(receipt.unwrap()));
    }

    /// @notice Returns encoded tips paid on origin chain.
    function tips(Receipt receipt) internal pure returns (Tips) {
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into Receipt, so this never reverts.
        return TipsLib.wrapPadded((receipt.unwrap().indexUint({index_: OFFSET_TIPS, bytes_: TIPS_LENGTH})));
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns an untyped memory view over the body field without checking
    /// if the whole payload or the body are properly formatted.
    function _body(MemView memView) private pure returns (MemView) {
        return memView.slice({index_: OFFSET_BODY, len_: RECEIPT_BODY_LENGTH});
    }
}
