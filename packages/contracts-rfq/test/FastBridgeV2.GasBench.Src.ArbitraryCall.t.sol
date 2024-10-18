// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkSrcTest} from "./FastBridgeV2.GasBench.Src.t.sol";

// solhint-disable func-name-mixedcase, no-empty-blocks
contract FastBridgeV2GasBenchmarkSrcArbitraryCallTest is FastBridgeV2GasBenchmarkSrcTest {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testFastBridgeV2GasBenchmarkSrcArbitraryCallTest() external {}

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
