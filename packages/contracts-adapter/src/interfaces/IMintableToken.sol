// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IMintableToken {
    function mint(address to, uint256 amount) external;
}
