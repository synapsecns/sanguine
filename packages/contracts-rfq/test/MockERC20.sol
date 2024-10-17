// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// solhint-disable no-empty-blocks
contract MockERC20 is ERC20 {
    uint8 private _decimals;

    constructor(string memory name_, uint8 decimals_) ERC20(name_, name_) {
        _decimals = decimals_;
    }

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testMockERC20() external {}

    function burn(address account, uint256 amount) external {
        _burn(account, amount);
    }

    function mint(address account, uint256 amount) external {
        _mint(account, amount);
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }
}
