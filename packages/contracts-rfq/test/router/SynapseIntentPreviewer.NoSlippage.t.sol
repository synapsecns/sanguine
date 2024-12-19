// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseIntentPreviewerTest} from "./SynapseIntentPreviewer.t.sol";

contract SynapseIntentPreviewerNoSlippageTest is SynapseIntentPreviewerTest {
    function setUp() public virtual override {
        super.setUp();
        swapMinAmountOut = SWAP_AMOUNT_OUT;
        slippageWei = 0;
    }
}
