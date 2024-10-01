// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcTest, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcExclusivityTest is FastBridgeV2SrcTest {
    uint256 public constant EXCLUSIVITY_PERIOD = 60 seconds;

    function createFixturesV2() public virtual override {
        tokenParamsV2.quoteRelayer = relayerA;
        tokenParamsV2.quoteExclusivitySeconds = int256(EXCLUSIVITY_PERIOD);
        tokenParamsV2.quoteId = bytes("Created by Relayer A");
        ethParamsV2.quoteRelayer = relayerB;
        ethParamsV2.quoteExclusivitySeconds = int256(EXCLUSIVITY_PERIOD);
        ethParamsV2.quoteId = bytes("Created by Relayer B");

        tokenTx.exclusivityRelayer = relayerA;
        tokenTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
        ethTx.exclusivityRelayer = relayerB;
        ethTx.exclusivityEndTime = block.timestamp + EXCLUSIVITY_PERIOD;
    }

    function bridge(address caller, uint256 msgValue, IFastBridge.BridgeParams memory params) public virtual override {
        IFastBridgeV2.BridgeParamsV2 memory paramsV2 = params.originToken == ETH_ADDRESS ? ethParamsV2 : tokenParamsV2;
        bridge(caller, msgValue, params, paramsV2);
    }
}
