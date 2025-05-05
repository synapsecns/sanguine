// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {FastBridgeInterceptorTest} from "./FastBridgeInterceptor.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeInterceptorSrc6Dst18Test is FastBridgeInterceptorTest {
    function getOriginValue(uint256 valueWei) internal pure override returns (uint256) {
        return valueWei / 1e12;
    }
}
