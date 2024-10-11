// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcBaseTest} from "./FastBridgeV2.Src.Base.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2GasBenchmarkEncodingTest is FastBridgeV2SrcBaseTest {
    // TODO: add more tests with variable length requests once arbitrary call is done

    function test_getBridgeTransaction() public view {
        bytes memory request = abi.encode(extractV1(tokenTx));
        fastBridge.getBridgeTransaction(request);
    }

    function test_getBridgeTransactionV2() public view {
        bytes memory request = abi.encode(tokenTx);
        fastBridge.getBridgeTransactionV2(request);
    }
}
