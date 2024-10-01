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
        tokenParamsV2.quoteRelayer = relayerA;
        tokenParamsV2.quoteExclusivitySeconds = int256(EXCLUSIVITY_PERIOD);
        ethParamsV2.quoteRelayer = relayerA;
        ethParamsV2.quoteExclusivitySeconds = int256(EXCLUSIVITY_PERIOD);
        tokenTx.exclusivityRelayer = relayerA;
        tokenTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
        ethTx.exclusivityRelayer = relayerA;
        ethTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
    }
}
