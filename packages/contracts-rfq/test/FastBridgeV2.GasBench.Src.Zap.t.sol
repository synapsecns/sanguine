// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2GasBenchmarkSrcTest} from "./FastBridgeV2.GasBench.Src.t.sol";

// solhint-disable func-name-mixedcase, no-empty-blocks
contract FastBridgeV2GasBenchmarkSrcZapTest is FastBridgeV2GasBenchmarkSrcTest {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testFastBridgeV2GasBenchmarkSrcZapTest() external {}

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        bytes memory mockZapData = abi.encode(userA, keccak256("Random ID"));
        setTokenTestZapData(mockZapData);
        setEthTestZapData(mockZapData);
        bridgedTokenTx.zapData = mockZapData;
        bridgedEthTx.zapData = mockZapData;
        provenTokenTx.zapData = mockZapData;
        provenEthTx.zapData = mockZapData;
    }
}
