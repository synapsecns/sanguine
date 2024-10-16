// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib} from "../contracts/libs/BridgeTransactionV2.sol";

import {FastBridgeV2SrcBaseTest} from "./FastBridgeV2.Src.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2GasBenchmarkEncodingTest is FastBridgeV2SrcBaseTest {
    function test_getBridgeTransaction() public view {
        bytes memory request = abi.encode(extractV1(tokenTx));
        fastBridge.getBridgeTransaction(request);
    }

    function test_getBridgeTransactionV2() public view {
        bytes memory request = BridgeTransactionV2Lib.encodeV2(tokenTx);
        fastBridge.getBridgeTransactionV2(request);
    }

    function test_getBridgeTransactionV2_withArbitraryCall() public {
        setTokenTestZapData({zapData: abi.encode(userA, keccak256("Random ID"))});
        test_getBridgeTransactionV2();
    }
}
