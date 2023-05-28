// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Origin} from "../../utils/SynapseTest.t.sol";
import {MessagingBaseTest} from "./MessagingBase.t.sol";

abstract contract AgentSecuredTest is MessagingBaseTest {
    // TODO: unit tests for AgentSecured
    // ═══════════════════════════════════════════ UPDATE IMPLEMENTATION ═══════════════════════════════════════════════

    function updateOrigin(uint32 domain, address agentManager, address inbox_, address gasOracle_) public {
        // Deploy new implementation with a different set of immutables
        Origin impl = new Origin(domain, agentManager, inbox_, gasOracle_);
        // Etch the implementation code to effectively update the values of immutables
        vm.etch(localOrigin(), address(impl).code);
    }
}
