// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @notice Simple ERC20 token for testing purposes with a public mint function
contract MockERC20Decimals is ERC20 {
    uint8 internal _decimals;

    constructor(string memory name_, string memory symbol_, uint8 decimals_) ERC20(name_, symbol_) {
        _decimals = decimals_;
    }

    /// @notice For testing purposes, mints tokens to the given account
    function mintPublic(address account, uint256 amount) external {
        _mint(account, amount);
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }

    // solhint-disable-next-line no-empty-blocks
    function testMockERC20Decimals() external pure {
        // This function is only used to remove MockERC20Decimals from coverage reports
    }
}
