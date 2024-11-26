// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Interface for a contract that supports multiple calls from the same caller. Inspired by MulticallV3:
/// https://github.com/mds1/multicall/blob/master/src/Multicall3.sol
interface IMulticallTarget {
    struct Result {
        bool success;
        bytes returnData;
    }

    /// @notice Executes multiple calls to this contract in a single transaction while preserving msg.sender.
    /// Return data from the calls is discarded.
    /// @dev This method is non-payable, so only calls with msg.value of 0 can be batched.
    /// If ignoreReverts is set to true, reverted calls will be skipped.
    /// Otherwise, the entire batch will revert with the original revert reason.
    /// @param data             List of ABI-encoded calldata for the calls to execute
    /// @param ignoreReverts    Whether to skip calls that revert
    function multicallNoResults(bytes[] calldata data, bool ignoreReverts) external;

    /// @notice Executes multiple calls to this contract in a single transaction while preserving msg.sender.
    /// Return data from each call is preserved.
    /// @dev This method is non-payable, so only calls with msg.value of 0 can be batched.
    /// If ignoreReverts is set to true, reverted calls will be skipped.
    /// Otherwise, the entire batch will revert with the original revert reason.
    /// @param data             List of ABI-encoded calldata for the calls to execute
    /// @param ignoreReverts    Whether to skip calls that revert
    /// @return results         List of results from the calls, each containing (success, returnData)
    function multicallWithResults(
        bytes[] calldata data,
        bool ignoreReverts
    )
        external
        returns (Result[] memory results);
}
