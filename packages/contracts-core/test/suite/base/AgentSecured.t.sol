// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CallerNotAgentManager} from "../../../contracts/libs/Errors.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
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

    // ════════════════════════════════════════════ TESTS: OPEN DISPUTE ════════════════════════════════════════════════

    function openTestDispute() public {
        vm.prank(localAgentManager());
        localAgentSecured().openDispute(1, 2);
    }

    function test_openDispute() public {
        // TODO: add expectations
        openTestDispute();
    }

    function test_openDispute_revert_onlyAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert(CallerNotAgentManager.selector);
        vm.prank(caller);
        localAgentSecured().openDispute(1, 2);
    }

    // ══════════════════════════════════════════ TESTS: RESOLVE DISPUTE ═══════════════════════════════════════════════

    function test_resolveDispute() public {
        openTestDispute();
        // TODO: add expectations
        vm.prank(localAgentManager());
        localAgentSecured().resolveDispute(1, 2);
    }

    function test_resolveDispute_revert_onlyAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert(CallerNotAgentManager.selector);
        vm.prank(caller);
        localAgentSecured().resolveDispute(1, 2);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @dev Returns tested contract as `IAgentSecured`
    function localAgentSecured() internal view returns (IAgentSecured) {
        return IAgentSecured(localContract());
    }
}
