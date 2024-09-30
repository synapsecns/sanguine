// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcTest} from "./FastBridgeV2.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcProtocolFeesTest is FastBridgeV2SrcTest {
    function configureFastBridge() public virtual override {
        super.configureFastBridge();
        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), address(this));
        fastBridge.setProtocolFeeRate(1e4); // 1%
    }

    function createFixtures() public virtual override {
        super.createFixtures();
        tokenTx.originFeeAmount = 0.01e6;
        tokenTx.originAmount = 0.99e6;
        tokenTx.destAmount = 0.98e6;
        tokenParams.destAmount = 0.98e6;
        ethTx.originFeeAmount = 0.01 ether;
        ethTx.originAmount = 0.99 ether;
        ethTx.destAmount = 0.98 ether;
        ethParams.destAmount = 0.98 ether;
    }
}
