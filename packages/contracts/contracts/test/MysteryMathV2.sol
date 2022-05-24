// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "./MysteryMath.sol";

contract MysteryMathV2 is MysteryMath {
    uint32 public immutable version;

    constructor() {
        version = 2;
    }

    function doMath(uint256 a, uint256 b)
        external
        pure
        override
        returns (uint256 _result)
    {
        _result = a * b;
    }
}
