// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseIntentPreviewerTest} from "./SynapseIntentPreviewer.t.sol";

contract SynapseIntentPreviewerWithSlippageTest is SynapseIntentPreviewerTest {
    function setUp() public virtual override {
        super.setUp();
        // 1% slippage
        swapMinAmountOut = SWAP_AMOUNT_OUT * 99 / 100;
        slippageWei = 0.01e18;
    }
}
