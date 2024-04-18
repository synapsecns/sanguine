// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MathLib} from "../../contracts/libs/Math.sol";

contract MathLibHarness {
    function roundUpToWord(uint256 x) external pure returns (uint256) {
        return MathLib.roundUpToWord(x);
    }
}
