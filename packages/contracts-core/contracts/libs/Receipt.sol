// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {RECEIPT_SALT, RECEIPT_LENGTH, TIPS_LENGTH} from "./Constants.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Receipt is a memory view over a formatted receipt payload.
type Receipt is bytes29;

/// @dev Attach library functions to Receipt
using ReceiptLib for Receipt global;

library ReceiptLib {
    using ByteString for bytes;
    using TipsLib for bytes29;
    using TypedMemView for bytes29;

    /**
     *
     * @dev Memory layout of Receipt fields
     * [000 .. 004): origin             uint32   4 bytes    Domain where message originated
     * [004 .. 008): destination        uint32   4 bytes    Domain where message was executed
     * [008 .. 040): messageHash        bytes32 32 bytes    Hash of the message
     * [040 .. 072): snapshotRoot       bytes32 32 bytes    Snapshot root used for proving the message
     * [072 .. 092): notary             address 20 bytes    Notary who posted attestation with snapshot root
     * [092 .. 112): firstExecutor      address 20 bytes    Executor who performed first valid execution attempt
     * [112 .. 132): finalExecutor      address 20 bytes    Executor who successfully executed the message
     * [132 .. 180): tips               bytes   48 bytes    Tips paid on origin chain
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_ORIGIN = 0;
    uint256 private constant OFFSET_DESTINATION = 4;
    uint256 private constant OFFSET_MESSAGE_HASH = 8;
    uint256 private constant OFFSET_SNAPSHOT_ROOT = 40;
    uint256 private constant OFFSET_NOTARY = 72;
    uint256 private constant OFFSET_FIRST_EXECUTOR = 92;
    uint256 private constant OFFSET_FINAL_EXECUTOR = 112;
    uint256 private constant OFFSET_TIPS = 132;

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Receipt payload with provided fields
     * @param origin_           Domain where message originated
     * @param destination_      Domain where message was executed
     * @param messageHash_      Hash of the message
     * @param snapshotRoot_     Snapshot root used for proving the message
     * @param notary_           Notary who posted attestation with snapshot root
     * @param firstExecutor_    Executor who performed first valid execution attempt
     * @param finalExecutor_    Executor who successfully executed the message
     * @param tipsPayload       Formatted payload with tips information
     * @return Formatted receipt
     */
    function formatReceipt(
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        address notary_,
        address firstExecutor_,
        address finalExecutor_,
        bytes memory tipsPayload
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            origin_, destination_, messageHash_, snapshotRoot_, notary_, firstExecutor_, finalExecutor_, tipsPayload
        );
    }

    /**
     * @notice Returns a Receipt view over the given payload.
     * @dev Will revert if the payload is not a receipt.
     */
    function castToReceipt(bytes memory payload) internal pure returns (Receipt) {
        return castToReceipt(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Receipt view.
     * @dev Will revert if the memory view is not over a receipt.
     */
    function castToReceipt(bytes29 view_) internal pure returns (Receipt) {
        require(isReceipt(view_), "Not a receipt");
        return Receipt.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted Receipt.
    function isReceipt(bytes29 view_) internal pure returns (bool) {
        // Check payload length
        if (view_.len() != RECEIPT_LENGTH) return false;
        // Check that tips payload is formatted
        return _tips(view_).isTips();
    }

    /// @notice Returns the hash of an Receipt, that could be later signed by a Notary.
    function hash(Receipt receipt) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 view_ = receipt.unwrap();
        // The final hash to sign is keccak(receiptSalt, keccak(receipt))
        return keccak256(bytes.concat(RECEIPT_SALT, view_.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Receipt receipt) internal pure returns (bytes29) {
        return Receipt.unwrap(receipt);
    }

    // ═════════════════════════════════════════════ RECEIPT SLICING ═════════════════════════════════════════════════

    /// @notice Returns receipt's origin field
    function origin(Receipt receipt) internal pure returns (uint32) {
        bytes29 view_ = unwrap(receipt);
        return uint32(view_.indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns receipt's destination field
    function destination(Receipt receipt) internal pure returns (uint32) {
        bytes29 view_ = unwrap(receipt);
        return uint32(view_.indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(Receipt receipt) internal pure returns (bytes32) {
        bytes29 view_ = unwrap(receipt);
        return view_.index({index_: OFFSET_MESSAGE_HASH, bytes_: 32});
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(Receipt receipt) internal pure returns (bytes32) {
        bytes29 view_ = unwrap(receipt);
        return view_.index({index_: OFFSET_SNAPSHOT_ROOT, bytes_: 32});
    }

    /// @notice Returns receipt's notary field
    function notary(Receipt receipt) internal pure returns (address) {
        bytes29 view_ = unwrap(receipt);
        return view_.indexAddress({index_: OFFSET_NOTARY});
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(Receipt receipt) internal pure returns (address) {
        bytes29 view_ = unwrap(receipt);
        return view_.indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(Receipt receipt) internal pure returns (address) {
        bytes29 view_ = unwrap(receipt);
        return view_.indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }

    /// @notice Returns a typed memory view over the payload with tips paid on origin chain.
    function tips(Receipt receipt) internal pure returns (Tips) {
        bytes29 view_ = receipt.unwrap();
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into Receipt, so this never reverts.
        return _tips(view_).castToTips();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns an untyped memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _tips(bytes29 view_) private pure returns (bytes29) {
        return view_.slice({index_: OFFSET_TIPS, len_: TIPS_LENGTH, newType: 0});
    }
}
