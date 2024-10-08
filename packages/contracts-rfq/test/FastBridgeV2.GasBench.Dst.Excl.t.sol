// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2DstGasBenchmarkTest} from "./FastBridgeV2.GasBench.Dst.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2DstExclusivityTest is FastBridgeV2DstGasBenchmarkTest {
    uint256 public constant EXCLUSIVITY_PERIOD = 60 seconds;

    function setUp() public virtual override {
        super.setUp();
        skip({time: EXCLUSIVITY_PERIOD / 2});
    }

    function createFixturesV2() public virtual override {
        setTokenTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
        setEthTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD);
    }
}
