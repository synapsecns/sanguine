// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkSrcTest} from "./FastBridgeV2.GasBench.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2GasBenchmarkSrcProtocolFeesTest is FastBridgeV2GasBenchmarkSrcTest {
    function configureFastBridge() public virtual override {
        super.configureFastBridge();
        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), address(this));
        fastBridge.setProtocolFeeRate(1e4); // 1%
    }

    function createFixtures() public virtual override {
        super.createFixtures();
        tokenTx.txV1.originFeeAmount = 0.01e6;
        tokenTx.txV1.originAmount = 0.99e6;
        tokenTx.txV1.destAmount = 0.98e6;
        tokenParams.destAmount = 0.98e6;
        ethTx.txV1.originFeeAmount = 0.01 ether;
        ethTx.txV1.originAmount = 0.99 ether;
        ethTx.txV1.destAmount = 0.98 ether;
        ethParams.destAmount = 0.98 ether;

        // Copy txs to bridged and proven with different nonce
        bridgedTokenTx = tokenTx;
        provenTokenTx = tokenTx;
        bridgedEthTx = ethTx;
        provenEthTx = ethTx;

        bridgedTokenTx.txV1.nonce = 0;
        bridgedEthTx.txV1.nonce = 1;
        provenTokenTx.txV1.nonce = 2;
        provenEthTx.txV1.nonce = 3;
    }
}
