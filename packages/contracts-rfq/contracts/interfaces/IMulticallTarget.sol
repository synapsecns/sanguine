// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice Interface for a contract that can be called multiple times by the same caller. Inspired by MulticallV3:
/// https://github.com/mds1/multicall/blob/master/src/Multicall3.sol
interface IMulticallTarget {
    struct Result {
        bool success;
        bytes returnData;
    }

    function multicallNoResults(bytes[] calldata data, bool ignoreReverts) external;
    function multicallWithResults(
        bytes[] calldata data,
        bool ignoreReverts
    )
        external
        returns (Result[] memory results);
}
