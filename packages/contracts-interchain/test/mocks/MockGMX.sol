// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {MockERC20} from "./MockERC20.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @notice MockGMX is a mock ERC20 token that follows GMX's mint-burn interface on Avalanche
contract MockGMX is MockERC20, Ownable {
    constructor(string memory name_, address owner_) MockERC20(name_) Ownable(owner_) {}

    // Follows GMX's mint-burn interface on Avalanche

    function burn(address account, uint256 amount) external onlyOwner {
        _burn(account, amount);
    }

    function mint(address account, uint256 amount) external onlyOwner {
        _mint(account, amount);
    }

    // solhint-disable-next-line no-empty-blocks
    function testMockGMX() external pure {
        // This function is only used to remove MockGMX from coverage reports
    }
}
