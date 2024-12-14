// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseIntentPreviewerTest} from "./SynapseIntentPreviewer.t.sol";

contract SynapseIntentPreviewerStrictOutTest is SynapseIntentPreviewerTest {
    function setUp() public virtual override {
        super.setUp();
        strictOut = true;
    }
}
