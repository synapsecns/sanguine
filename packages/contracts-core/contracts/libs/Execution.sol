// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ByteString} from "./ByteString.sol";
import {EXECUTION_SALT, EXECUTION_LENGTH, TIPS_LENGTH} from "./Constants.sol";
import {MessageStatus} from "./Structures.sol";
import {Tips, TipsLib} from "./Tips.sol";
import {TypedMemView} from "./TypedMemView.sol";

/// @dev Execution is a memory view over a formatted execution payload.
type Execution is bytes29;

/// @dev Attach library functions to Execution
using ExecutionLib for Execution global;

library ExecutionLib {
    using ByteString for bytes;
    using TipsLib for bytes29;
    using TypedMemView for bytes29;

    /**
     *
     * @dev Memory layout of Execution fields
     * [000 .. 001): status             uint8    1 byte     Message execution status
     * [001 .. 005): origin             uint32   4 bytes    Domain where message originated
     * [005 .. 009): nonce              uint32   4 bytes    Message nonce on the origin domain
     * [009 .. 013): destination        uint32   4 bytes    Domain where message was executed
     * [013 .. 017): attNonce           uint32   4 bytes    Nonce of the execution used for proving the message
     * [017 .. 037): firstExecutor      address 20 bytes    Executor who performed first valid execution attempt
     * [037 .. 057): finalExecutor      address 20 bytes    Executor who successfully executed the message
     * [057 .. 115): tips               bytes   48 bytes    Tips paid on origin chain
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_STATUS = 0;
    uint256 private constant OFFSET_ORIGIN = 1;
    uint256 private constant OFFSET_NONCE = 5;
    uint256 private constant OFFSET_DESTINATION = 9;
    uint256 private constant OFFSET_ATT_NONCE = 13;
    uint256 private constant OFFSET_FIRST_EXECUTOR = 17;
    uint256 private constant OFFSET_FINAL_EXECUTOR = 37;
    uint256 private constant OFFSET_TIPS = 57;

    // ═════════════════════════════════════════════════ EXECUTION ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Execution payload with provided fields
     * @param status_           Message execution status
     * @param origin_           Domain where message originated
     * @param nonce_            Message nonce on origin chain
     * @param destination_      Domain where message was executed
     * @param attNonce_         Nonce of the execution used for proving the message
     * @param firstExecutor_    Executor who performed first valid execution attempt
     * @param finalExecutor_    Executor who successfully executed the message
     * @param tipsPayload       Formatted payload with tips information
     * @return Formatted execution
     */
    function formatExecution(
        MessageStatus status_,
        uint32 origin_,
        uint32 nonce_,
        uint32 destination_,
        uint32 attNonce_,
        address firstExecutor_,
        address finalExecutor_,
        bytes memory tipsPayload
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            status_, origin_, nonce_, destination_, attNonce_, firstExecutor_, finalExecutor_, tipsPayload
        );
    }

    /**
     * @notice Returns a Execution view over the given payload.
     * @dev Will revert if the payload is not a execution.
     */
    function castToExecution(bytes memory payload) internal pure returns (Execution) {
        return castToExecution(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Execution view.
     * @dev Will revert if the memory view is not over a execution.
     */
    function castToExecution(bytes29 view_) internal pure returns (Execution) {
        require(isExecution(view_), "Not a execution");
        return Execution.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted Execution.
    function isExecution(bytes29 view_) internal pure returns (bool) {
        // Check payload length
        if (view_.len() != EXECUTION_LENGTH) return false;
        // Check that status in in range of MessageStatus enum
        if (_status(view_) > uint8(type(MessageStatus).max)) return false;
        // Check that tips payload is formatted
        return _tips(view_).isTips();
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Execution execution) internal pure returns (bytes29) {
        return Execution.unwrap(execution);
    }

    // ═════════════════════════════════════════════ EXECUTION HASHING ═════════════════════════════════════════════════

    /// @notice Returns the hash of an Execution, that could be later signed by a Notary.
    function hash(Execution execution) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 view_ = execution.unwrap();
        // The final hash to sign is keccak(executionSalt, keccak(execution))
        return keccak256(bytes.concat(EXECUTION_SALT, view_.keccak()));
    }

    // ═════════════════════════════════════════════ EXECUTION SLICING ═════════════════════════════════════════════════

    /// @notice Returns execution's status.
    function status(Execution execution) internal pure returns (MessageStatus) {
        bytes29 view_ = execution.unwrap();
        // We check that status fits into enum, when payload is wrapped
        // into Execution, so this never reverts
        return MessageStatus(_status(view_));
    }

    /// @notice Returns execution's origin field
    function origin(Execution execution) internal pure returns (uint32) {
        bytes29 view_ = unwrap(execution);
        return uint32(view_.indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns execution's nonce field
    function nonce(Execution execution) internal pure returns (uint32) {
        bytes29 view_ = unwrap(execution);
        return uint32(view_.indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns execution's origin and nonce fields combined in a composite key.
    function originAndNonce(Execution execution) internal pure returns (uint64) {
        bytes29 view_ = unwrap(execution);
        return uint64(view_.indexUint({index_: OFFSET_ORIGIN, bytes_: 8}));
    }

    /// @notice Returns execution's destination field
    function destination(Execution execution) internal pure returns (uint32) {
        bytes29 view_ = unwrap(execution);
        return uint32(view_.indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns execution's "attestation nonce" field
    function attNonce(Execution execution) internal pure returns (uint32) {
        bytes29 view_ = unwrap(execution);
        return uint32(view_.indexUint({index_: OFFSET_ATT_NONCE, bytes_: 4}));
    }

    /// @notice Returns execution's "first executor" field
    function firstExecutor(Execution execution) internal pure returns (address) {
        bytes29 view_ = unwrap(execution);
        return view_.indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns execution's "final executor" field
    function finalExecutor(Execution execution) internal pure returns (address) {
        bytes29 view_ = unwrap(execution);
        return view_.indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }

    /// @notice Returns a typed memory view over the payload with tips paid on origin chain.
    function tips(Execution execution) internal pure returns (Tips) {
        bytes29 view_ = execution.unwrap();
        // We check that tips payload is properly formatted, when the whole payload is wrapped
        // into Execution, so this never reverts.
        return _tips(view_).castToTips();
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns execution's status without checking that it fits into MessageStatus enum.
    function _status(bytes29 view_) private pure returns (uint8) {
        return uint8(view_.indexUint({index_: OFFSET_STATUS, bytes_: 1}));
    }

    /// @dev Returns an untyped memory view over the tips field without checking
    /// if the whole payload or the tips are properly formatted.
    function _tips(bytes29 view_) private pure returns (bytes29) {
        return view_.slice({index_: OFFSET_TIPS, len_: TIPS_LENGTH, newType: 0});
    }
}
