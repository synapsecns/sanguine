// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// solhint-disable-next-line no-empty-blocks
contract NoOpContract {
    /// @notice Mock needs to accept ETH
    receive() external payable {}
}
