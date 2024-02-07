// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// Request is encoded data with "message execution request".
type Request is uint192;

using RequestLib for Request global;

/// Library for formatting _the request part_ of _the base messages_.
/// - Request represents a message sender requirements for the message execution on the destination chain.
/// - Request occupies a single storage word, and thus is stored on stack instead of being stored in memory.
/// > gasDrop field is included for future compatibility and is ignored at the moment.
///
/// # Request stack layout (from highest bits to lowest)
///
/// | Position   | Field    | Type   | Bytes | Description                                          |
/// | ---------- | -------- | ------ | ----- | ---------------------------------------------------- |
/// | (024..012] | gasDrop  | uint96 | 12    | Minimum amount of gas token to drop to the recipient |
/// | (012..004] | gasLimit | uint64 | 8     | Minimum amount of gas units to supply for execution  |
/// | (004..000] | version  | uint32 | 4     | Base message version to pass to the recipient        |

library RequestLib {
    /// @dev Amount of bits to shift to gasDrop field
    uint192 private constant SHIFT_GAS_DROP = 12 * 8;
    /// @dev Amount of bits to shift to gasLimit field
    uint192 private constant SHIFT_GAS_LIMIT = 4 * 8;

    /// @notice Returns an encoded request with the given fields
    /// @param gasDrop_     Minimum amount of gas token to drop to the recipient (ignored at the moment)
    /// @param gasLimit_    Minimum amount of gas units to supply for execution
    /// @param version_     Base message version to pass to the recipient
    function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) internal pure returns (Request) {
        // Casts below are upcasts, so they are safe
        return Request.wrap(uint192(gasDrop_) << SHIFT_GAS_DROP | uint192(gasLimit_) << SHIFT_GAS_LIMIT | version_);
    }

    /// @notice Wraps the padded encoded request into a Request-typed value.
    /// @dev The "padded" request is simply an encoded request casted to uint256 (highest bits are set to zero).
    /// Casting to uint256 is done automatically in Solidity, so no extra actions from consumers are needed.
    /// The highest bits are discarded, so that the contracts dealing with encoded requests
    /// don't need to be updated, if a new field is added.
    function wrapPadded(uint256 paddedRequest) internal pure returns (Request) {
        // Casting to uint192 will truncate the highest bits, which is the behavior we want
        return Request.wrap(uint192(paddedRequest));
    }

    /// @notice Returns the requested of gas token to drop to the recipient.
    function gasDrop(Request request) internal pure returns (uint96) {
        // Casting to uint96 will truncate the highest bits, which is the behavior we want
        return uint96(Request.unwrap(request) >> SHIFT_GAS_DROP);
    }

    /// @notice Returns the requested minimum amount of gas units to supply for execution.
    function gasLimit(Request request) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Request.unwrap(request) >> SHIFT_GAS_LIMIT);
    }

    /// @notice Returns the requested base message version to pass to the recipient.
    function version(Request request) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Request.unwrap(request));
    }
}
