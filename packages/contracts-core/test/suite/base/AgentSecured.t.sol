// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CallerNotAgentManager} from "../../../contracts/libs/Errors.sol";
import {DisputeFlag, DisputeStatus} from "../../../contracts/libs/Structures.sol";
import {IAgentSecured} from "../../../contracts/interfaces/IAgentSecured.sol";
import {Origin} from "../../utils/SynapseTest.t.sol";
import {MessagingBaseTest} from "./MessagingBase.t.sol";

abstract contract AgentSecuredTest is MessagingBaseTest {
    // ═══════════════════════════════════════════ UPDATE IMPLEMENTATION ═══════════════════════════════════════════════

    function updateOrigin(uint32 domain, address agentManager, address inbox_, address gasOracle_) public {
        // Deploy new implementation with a different set of immutables
        Origin impl = new Origin(domain, agentManager, inbox_, gasOracle_);
        // Etch the implementation code to effectively update the values of immutables
        vm.etch(localOrigin(), address(impl).code);
    }

    // ════════════════════════════════════════════ TESTS: OPEN DISPUTE ════════════════════════════════════════════════

    function openTestDispute(uint32 guardIndex, uint32 notaryIndex) public {
        vm.prank(localAgentManager());
        localAgentSecured().openDispute(guardIndex, notaryIndex);
    }

    function resolveTestDispute(uint32 slashedIndex, uint32 rivalIndex) public {
        vm.prank(localAgentManager());
        localAgentSecured().resolveDispute(slashedIndex, rivalIndex);
    }

    function checkLatestDisputeStatus(uint32 index, DisputeStatus memory expected) public {
        DisputeStatus memory actual = localAgentSecured().latestDisputeStatus(index);
        assertEq(uint8(actual.flag), uint8(expected.flag), "!flag");
        assertEq(actual.openedAt, expected.openedAt, "!openedAt");
        assertEq(actual.resolvedAt, expected.resolvedAt, "!resolvedAt");
    }

    function test_openDispute() public {
        DisputeStatus memory expected = DisputeStatus({flag: DisputeFlag.Pending, openedAt: 1234, resolvedAt: 0});
        vm.warp(1234);
        openTestDispute({guardIndex: 1, notaryIndex: 2});
        checkLatestDisputeStatus(1, expected);
        checkLatestDisputeStatus(2, expected);
    }

    function test_openDispute_revert_onlyAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert(CallerNotAgentManager.selector);
        vm.prank(caller);
        localAgentSecured().openDispute(1, 2);
    }

    // ══════════════════════════════════════════ TESTS: RESOLVE DISPUTE ═══════════════════════════════════════════════

    function test_resolveDispute() public {
        vm.warp(1234);
        openTestDispute({guardIndex: 1, notaryIndex: 2});
        vm.warp(5678);
        resolveTestDispute({slashedIndex: 1, rivalIndex: 2});
        checkLatestDisputeStatus(1, DisputeStatus({flag: DisputeFlag.Slashed, openedAt: 1234, resolvedAt: 5678}));
        checkLatestDisputeStatus(2, DisputeStatus({flag: DisputeFlag.None, openedAt: 1234, resolvedAt: 5678}));
    }

    function test_resolveDispute_revert_onlyAgentManager(address caller) public {
        openTestDispute({guardIndex: 1, notaryIndex: 2});
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
