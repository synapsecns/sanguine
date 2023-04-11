// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManagerMock, AgentFlag, AgentStatus} from "./AgentManagerMock.t.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract AgentManagerMockTest is Test {
    AgentManagerMock internal agentManager;

    function setUp() public {
        agentManager = new AgentManagerMock();
    }

    // ═════════════════════════════════════════════════ ADD AGENT ═════════════════════════════════════════════════════

    function test_addAgent(uint32 domain, address agent) public {
        agentManager.addAgent(domain, agent);
        AgentStatus memory status = agentManager.agentStatus(agent);
        assertEq(uint8(status.flag), uint8(AgentFlag.Active));
        assertEq(status.domain, domain);
        assertEq(status.index, 1);
        assertEq(agentManager.totalAgents(), 1);
        (bool isSlashed, address prover) = agentManager.slashStatus(agent);
        assertFalse(isSlashed);
        assertEq(prover, address(0));
    }

    function test_addAgent_revert_alreadyActive(uint32 domain, address agent, uint32 otherDomain) public {
        test_addAgent(domain, agent);
        vm.expectRevert("Agent already active");
        agentManager.addAgent(otherDomain, agent);
    }

    // ═══════════════════════════════════════════════ REMOVE AGENT ════════════════════════════════════════════════════

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        agentManager.removeAgent(domain, agent);
        AgentStatus memory status = agentManager.agentStatus(agent);
        assertEq(uint8(status.flag), uint8(AgentFlag.Unknown));
        assertEq(status.domain, 0);
        assertEq(status.index, 0);
        (bool isSlashed, address prover) = agentManager.slashStatus(agent);
        assertFalse(isSlashed);
        assertEq(prover, address(0));
    }

    function test_removeAgent_revert_agentNotActive(uint32 domain, address agent) public {
        test_removeAgent(domain, agent);
        vm.expectRevert("Agent not active");
        agentManager.removeAgent(domain, agent);
    }

    function test_removeAgent_revert_incorrectDomain(uint32 domain, address agent, uint32 otherDomain) public {
        vm.assume(domain != otherDomain);
        test_addAgent(domain, agent);
        vm.expectRevert("Incorrect domain");
        agentManager.removeAgent(otherDomain, agent);
    }

    // ══════════════════════════════════════════════ REGISTRY SLASH ═══════════════════════════════════════════════════

    function test_registrySlash(uint32 domain, address agent, address prover) public {
        test_addAgent(domain, agent);
        agentManager.registrySlash(domain, agent, prover);
        AgentStatus memory status = agentManager.agentStatus(agent);
        assertEq(uint8(status.flag), uint8(AgentFlag.Slashed));
        assertEq(status.domain, domain);
        assertEq(status.index, 1);
        (bool isSlashed, address prover_) = agentManager.slashStatus(agent);
        assertTrue(isSlashed);
        assertEq(prover_, prover);
    }

    function test_registrySlash_revert_agentNotActive(uint32 domain, address agent) public {
        test_removeAgent(domain, agent);
        vm.expectRevert("Agent not active");
        agentManager.registrySlash(domain, agent, address(0));
    }

    function test_registrySlash_revert_incorrectDomain(uint32 domain, address agent, uint32 otherDomain) public {
        vm.assume(domain != otherDomain);
        test_addAgent(domain, agent);
        vm.expectRevert("Incorrect domain");
        agentManager.registrySlash(otherDomain, agent, address(0));
    }

    // ══════════════════════════════════════════════ SET AGENT ROOT ═══════════════════════════════════════════════════

    function test_setAgentRoot(bytes32 root) public {
        agentManager.setAgentRoot(root);
        assertEq(agentManager.agentRoot(), root);
    }
}
