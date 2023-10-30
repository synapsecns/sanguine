// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice Library for accessing chain context variables as tightly packed integers.
/// Messaging contracts should rely on this library for accessing chain context variables
/// instead of doing the casting themselves.
library ChainContext {
    using SafeCast for uint256;

    /// @notice Returns the current block number as uint40.
    /// @dev Reverts if block number is greater than 40 bits, which is not supposed to happen
    /// until the block.timestamp overflows (assuming block time is at least 1 second).
    function blockNumber() internal view returns (uint40) {
        return block.number.toUint40();
    }

    /// @notice Returns the current block timestamp as uint40.
    /// @dev Reverts if block timestamp is greater than 40 bits, which is
    /// supposed to happen approximately in year 36835.
    function blockTimestamp() internal view returns (uint40) {
        return block.timestamp.toUint40();
    }

    /// @notice Returns the chain id as uint32.
    /// @dev Reverts if chain id is greater than 32 bits, which is not
    /// supposed to happen in production.
    function chainId() internal view returns (uint32) {
        return block.chainid.toUint32();
    }
}
