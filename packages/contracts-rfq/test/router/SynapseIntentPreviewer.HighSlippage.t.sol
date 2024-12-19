// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseIntentPreviewerTest} from "./SynapseIntentPreviewer.t.sol";

contract SynapseIntentPreviewerWithSlippageTest is SynapseIntentPreviewerTest {
    function setUp() public virtual override {
        super.setUp();
        // slippage out of bounds - should be capped at 100%
        swapMinAmountOut = 0;
        slippageWei = 1e20;
    }
}
