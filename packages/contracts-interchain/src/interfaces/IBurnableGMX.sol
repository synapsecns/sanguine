// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBurnableGMX is IERC20 {
    function burn(address account, uint256 amount) external;

    function mint(address account, uint256 amount) external;
}
