// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkSrcTest} from "./FastBridgeV2.GasBench.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2GasBenchmarkSrcArbitraryCallTest is FastBridgeV2GasBenchmarkSrcTest {
    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        bytes memory mockCallParams = abi.encode(userA, keccak256("Random ID"));
        setTokenTestCallParams(mockCallParams);
        setEthTestCallParams(mockCallParams);
        bridgedTokenTx.callParams = mockCallParams;
        bridgedEthTx.callParams = mockCallParams;
        provenTokenTx.callParams = mockCallParams;
        provenEthTx.callParams = mockCallParams;
    }
}
