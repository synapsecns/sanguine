// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {VaultMock} from "./VaultMock.sol";

/// @notice Vault mock for testing purposes. DO NOT USE IN PRODUCTION.
contract VaultManyArguments is VaultMock {
    error VaultManyArguments__SomeError();

    function deposit(
        address token,
        bytes memory encodedToken,
        uint256 amount,
        address user,
        bytes memory encodedUser
    )
        external
        payable
    {
        // Make sure the data is not malformed
        _validateBytes(token, encodedToken);
        _validateBytes(user, encodedUser);
        _deposit(user, token, amount);
    }

    function depositNoAmount(address user) external payable {
        _deposit(user, NATIVE_GAS_TOKEN, msg.value);
    }

    function depositWithRevert() external payable {
        revert VaultManyArguments__SomeError();
    }

    function _validateBytes(address addr, bytes memory encoded) internal pure {
        if (keccak256(abi.encode(addr)) != keccak256(encoded)) revert VaultManyArguments__SomeError();
    }
}
