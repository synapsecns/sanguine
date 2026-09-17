// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TestToken} from "./TestToken.sol";

contract MintableTestToken is TestToken {
    function mint(address to, uint256 amount) external {
        mintTestTokens(to, amount);
    }
}
