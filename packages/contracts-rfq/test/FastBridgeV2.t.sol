// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeTest} from "./FastBridge.t.sol";

contract FastBridgeV2Test is FastBridgeTest {
    function deployFastBridge() internal virtual override returns (address) {
        // TODO: change to FastBridgeV2
        return super.deployFastBridge();
    }
}