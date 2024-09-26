// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    uint8 private _decimals;

    constructor(string memory name_, uint8 decimals_) ERC20(name_, name_) {
        _decimals = decimals_;
    }

    function burn(address account, uint256 amount) external {
        // TODO: remove
        assert(balanceOf(account) >= amount);
        _burn(account, amount);
    }

    function mint(address account, uint256 amount) external {
        _mint(account, amount);
        // TODO: remove
        assert(balanceOf(account) >= amount);
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }
}
