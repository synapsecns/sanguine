// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IBurnableToken {
    function burn(uint256 amount) external;

    function burnFrom(address from, uint256 amount) external;
}
