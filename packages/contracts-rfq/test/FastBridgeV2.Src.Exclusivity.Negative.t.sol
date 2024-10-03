// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcTest, IFastBridge, IFastBridgeV2} from "./FastBridgeV2.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcExclusivityNegativeTest is FastBridgeV2SrcTest {
    uint256 public constant EXCLUSIVITY_PERIOD_ABS = 60 seconds;

    function createFixturesV2() public virtual override {
        // Populate the fields using the absolute exclusivity period
        setTokenTestExclusivityParams(relayerA, EXCLUSIVITY_PERIOD_ABS);
        setEthTestExclusivityParams(relayerB, EXCLUSIVITY_PERIOD_ABS);
        // Override with negative exclusivity period
        tokenParamsV2.quoteExclusivitySeconds = -int256(EXCLUSIVITY_PERIOD_ABS);
        ethParamsV2.quoteExclusivitySeconds = -int256(EXCLUSIVITY_PERIOD_ABS);
        tokenTx.exclusivityEndTime = block.timestamp - EXCLUSIVITY_PERIOD_ABS;
        ethTx.exclusivityEndTime = block.timestamp - EXCLUSIVITY_PERIOD_ABS;
    }

    function bridge(address caller, uint256 msgValue, IFastBridge.BridgeParams memory params) public virtual override {
        IFastBridgeV2.BridgeParamsV2 memory paramsV2 = params.originToken == ETH_ADDRESS ? ethParamsV2 : tokenParamsV2;
        bridge(caller, msgValue, params, paramsV2);
    }

    function test_bridge_revert_exclusivityEndTimeZero() public {
        tokenParamsV2.quoteExclusivitySeconds = -int256(block.timestamp);
        vm.expectRevert(ExclusivityParamsIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
    }

    function test_bridge_revert_exclusivityPeriodUnderflow() public {
        tokenParamsV2.quoteExclusivitySeconds = -int256(block.timestamp + 1);
        vm.expectRevert(ExclusivityParamsIncorrect.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
    }

    function test_bridge_eth_revert_exclusivityEndTimeZero() public {
        ethParamsV2.quoteExclusivitySeconds = -int256(block.timestamp);
        vm.expectRevert(ExclusivityParamsIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    function test_bridge_eth_revert_exclusivityPeriodUnderflow() public {
        ethParamsV2.quoteExclusivitySeconds = -int256(block.timestamp + 1);
        vm.expectRevert(ExclusivityParamsIncorrect.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }
}
