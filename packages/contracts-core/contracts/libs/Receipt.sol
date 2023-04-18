// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RECEIPT_SALT, RECEIPT_LENGTH, TIPS_LENGTH} from "./Constants.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {MemView, MemViewLib} from "./MemView.sol";

/// @dev Receipt is a memory view over a formatted receipt payload.
type Receipt is uint256;

/// @dev Attach library functions to Receipt
using ReceiptLib for Receipt global;

library ReceiptLib {
    using MemViewLib for bytes;
    using TipsLib for MemView;

    /**
     *
     * @dev Memory layout of Receipt fields
     * [000 .. 004): origin             uint32   4 bytes    Domain where message originated
     * [004 .. 008): destination        uint32   4 bytes    Domain where message was executed
     * [008 .. 040): messageHash        bytes32 32 bytes    Hash of the message
     * [040 .. 072): snapshotRoot       bytes32 32 bytes    Snapshot root used for proving the message
     * [072 .. 073): stateIndex         uint8    1 byte     Index of state used for the snapshot proof
     * [073 .. 093): attNotary          address 20 bytes    Notary who posted attestation with snapshot root
     * [093 .. 113): firstExecutor      address 20 bytes    Executor who performed first valid execution attempt
     * [113 .. 133): finalExecutor      address 20 bytes    Executor who successfully executed the message
     * [133 .. 181): tips               bytes   48 bytes    Tips paid on origin chain
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_ORIGIN = 0;
    uint256 private constant OFFSET_DESTINATION = 4;
    uint256 private constant OFFSET_MESSAGE_HASH = 8;
    uint256 private constant OFFSET_SNAPSHOT_ROOT = 40;
    uint256 private constant OFFSET_STATE_INDEX = 72;
    uint256 private constant OFFSET_ATT_NOTARY = 73;
    uint256 private constant OFFSET_FIRST_EXECUTOR = 93;
    uint256 private constant OFFSET_FINAL_EXECUTOR = 113;
    uint256 private constant OFFSET_TIPS = 133;

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Receipt payload with provided fields
     * @param origin_           Domain where message originated
     * @param destination_      Domain where message was executed
     * @param messageHash_      Hash of the message
     * @param snapshotRoot_     Snapshot root used for proving the message
     * @param stateIndex_       Index of state used for the snapshot proof
     * @param attNotary_        Notary who posted attestation with snapshot root
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
        uint8 stateIndex_,
        address attNotary_,
        address firstExecutor_,
        address finalExecutor_,
        bytes memory tipsPayload
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            origin_,
            destination_,
            messageHash_,
            snapshotRoot_,
            stateIndex_,
            attNotary_,
            firstExecutor_,
            finalExecutor_,
            tipsPayload
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
        require(isReceipt(memView), "Not a receipt");
        return Receipt.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Receipt.
    function isReceipt(MemView memView) internal pure returns (bool) {
        // Check payload length
        if (memView.len() != RECEIPT_LENGTH) return false;
        // Check that tips payload is formatted
        return _tips(memView).isTips();
    }

    /// @notice Returns the hash of an Receipt, that could be later signed by a Notary.
    function hash(Receipt receipt) internal pure returns (bytes32) {
        // Get the underlying memory view
        MemView memView = receipt.unwrap();
        // The final hash to sign is keccak(receiptSalt, keccak(receipt))
        return keccak256(bytes.concat(RECEIPT_SALT, memView.keccak()));
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Receipt receipt) internal pure returns (MemView) {
        return MemView.wrap(Receipt.unwrap(receipt));
    }

    // ═════════════════════════════════════════════ RECEIPT SLICING ═════════════════════════════════════════════════

    /// @notice Returns receipt's origin field
    function origin(Receipt receipt) internal pure returns (uint32) {
        MemView memView = unwrap(receipt);
        return uint32(memView.indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns receipt's destination field
    function destination(Receipt receipt) internal pure returns (uint32) {
        MemView memView = unwrap(receipt);
        return uint32(memView.indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(Receipt receipt) internal pure returns (bytes32) {
        MemView memView = unwrap(receipt);
        return memView.index({index_: OFFSET_MESSAGE_HASH, bytes_: 32});
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(Receipt receipt) internal pure returns (bytes32) {
        MemView memView = unwrap(receipt);
        return memView.index({index_: OFFSET_SNAPSHOT_ROOT, bytes_: 32});
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(Receipt receipt) internal pure returns (uint8) {
        MemView memView = unwrap(receipt);
        return uint8(memView.indexUint({index_: OFFSET_STATE_INDEX, bytes_: 1}));
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(Receipt receipt) internal pure returns (address) {
        MemView memView = unwrap(receipt);
        return memView.indexAddress({index_: OFFSET_ATT_NOTARY});
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(Receipt receipt) internal pure returns (address) {
        MemView memView = unwrap(receipt);
        return memView.indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(Receipt receipt) internal pure returns (address) {
        MemView memView = unwrap(receipt);
        return memView.indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }

    /// @notice Returns a typed memory view over the payload with tips paid on origin chain.
    function tips(Receipt receipt) internal pure returns (Tips) {
        MemView memView = receipt.unwrap();
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into Receipt, so this never reverts.
        return _tips(memView).castToTips();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns an untyped memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _tips(MemView memView) private pure returns (MemView) {
        return memView.slice({index_: OFFSET_TIPS, len_: TIPS_LENGTH});
    }
}
