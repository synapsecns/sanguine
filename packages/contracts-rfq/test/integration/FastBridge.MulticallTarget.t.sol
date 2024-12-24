// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridge} from "../../contracts/FastBridge.sol";

import {IFastBridge, MulticallTargetIntegrationTest} from "./MulticallTarget.t.sol";

// solhint-disable no-empty-blocks
contract FastBridgeMulticallTargetTest is MulticallTargetIntegrationTest {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testFastBridgeMulticallTargetTest() external {}

    function deployAndConfigureFastBridge() public override returns (address) {
        FastBridge fastBridge = new FastBridge(address(this));
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayer);
        return address(fastBridge);
    }

    function getEncodedBridgeTx(IFastBridge.BridgeTransaction memory bridgeTx)
        public
        pure
        override
        returns (bytes memory)
    {
        return abi.encode(bridgeTx);
    }
}
