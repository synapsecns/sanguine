// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TestToken} from "./TestToken.sol";

contract TestTokenDecimals is TestToken {
    uint8 private immutable tokenDecimals;

    constructor(uint8 decimals_) {
        tokenDecimals = decimals_;
    }

    function decimals() public view override returns (uint8) {
        return tokenDecimals;
    }
}
