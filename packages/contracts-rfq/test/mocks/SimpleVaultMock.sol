// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {VaultMock} from "./VaultMock.sol";

// solhint-disable no-empty-blocks
/// @notice Vault mock for testing purposes. DO NOT USE IN PRODUCTION.
contract SimpleVaultMock is VaultMock {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testSimpleVaultMock() external {}

    function deposit(address token, uint256 amount, address user) external payable {
        _deposit(user, token, amount);
    }
}
