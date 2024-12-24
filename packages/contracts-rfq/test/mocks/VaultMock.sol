// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

// solhint-disable no-empty-blocks
/// @notice Vault mock for testing purposes. DO NOT USE IN PRODUCTION.
abstract contract VaultMock {
    using SafeERC20 for IERC20;

    address internal constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    mapping(address user => mapping(address token => uint256 amount)) public balanceOf;

    error VaultMock__AmountIncorrect();

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testVaultMock() external {}

    function _deposit(address user, address token, uint256 amount) internal {
        if (token == NATIVE_GAS_TOKEN) {
            if (msg.value != amount) revert VaultMock__AmountIncorrect();
        } else {
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        }
        balanceOf[user][token] += amount;
    }
}
