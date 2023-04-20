// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManager} from "../../../contracts/manager/AgentManager.sol";
import {ISystemRegistry} from "../../../contracts/interfaces/ISystemRegistry.sol";
import {AgentFlag, AgentStatus, SlashStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";

import {SystemContractTest} from "../system/SystemContract.t.sol";
import {RawCallData, RawManagerCall} from "../../utils/libs/SynapseStructs.t.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract AgentManagerTest is SystemContractTest {
    using Address for address;

    uint256 internal rootSubmittedAt;

    function test_setup() public virtual {
        assertEq(address(testedAM().destination()), localDestination());
        assertEq(address(testedAM().origin()), localOrigin());
        assertEq(testedAM().agentRoot(), getAgentRoot());
    }

    // ═══════════════════════════════════════════ TESTS: REGISTRY SLASH ═══════════════════════════════════════════════

    function test_registrySlash_revertUnauthorized(address caller) public {
        vm.assume(!isLocalSystemRegistry(caller));
        vm.expectRevert("Unauthorized caller");
        vm.prank(caller);
        // Try to slash an existing agent
        testedAM().registrySlash(0, domains[0].agent, address(0));
    }

    function test_registrySlash_origin(uint256 domainId, uint256 agentId, address prover) public {
        _test_registrySlash(true, domainId, agentId, prover);
    }

    function test_registrySlash_destination(uint256 domainId, uint256 agentId, address prover) public {
        _test_registrySlash(false, domainId, agentId, prover);
    }

    function _test_registrySlash(bool onOrigin, uint256 domainId, uint256 agentId, address prover) internal {
        address registryF = onOrigin ? localOrigin() : localDestination();
        address registryS = onOrigin ? localDestination() : localOrigin();
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        vm.expectEmit();
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        // Expect call to the second registry
        vm.expectCall(registryS, abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent));
        // Prank call from the first registry
        vm.prank(registryF);
        testedAM().registrySlash(domain, agent, prover);
        assertEq(uint8(testedAM().agentStatus(agent).flag), uint8(AgentFlag.Fraudulent));
        (bool isSlashed, address prover_) = testedAM().slashStatus(agent);
        assertTrue(isSlashed);
        assertEq(prover_, prover);
    }

    // ═══════════════════════════════════════════════ TESTS: VIEWS ════════════════════════════════════════════════════

    function test_getAgent_notExistingIndex() public {
        (address agent, AgentStatus memory status) = testedAM().getAgent(0);
        assertEq(agent, address(0));
        assertEq(uint8(status.flag), 0);
        assertEq(status.domain, 0);
        assertEq(status.index, 0);
        // Last agent has index DOMAIN_AGENTS * allDomains.length
        (agent, status) = testedAM().getAgent(DOMAIN_AGENTS * allDomains.length + 1);
        assertEq(agent, address(0));
        assertEq(uint8(status.flag), 0);
        assertEq(status.domain, 0);
        assertEq(status.index, 0);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkAgentStatus(address agent, AgentStatus memory status, AgentFlag flag) public virtual override {
        super.checkAgentStatus(agent, status, flag);
        (address agent_, AgentStatus memory status_) = testedAM().getAgent(status.index);
        assertEq(agent_, agent, "!agent");
        super.checkAgentStatus(agent, status_, flag);
    }

    function skipBondingOptimisticPeriod() public {
        skipPeriod(BONDING_OPTIMISTIC_PERIOD);
    }

    function skipPeriod(uint256 period) public {
        rootSubmittedAt = block.timestamp;
        skip(period);
    }

    function managerMsgPrank(bytes memory payload) public {
        vm.prank(localDestination());
        systemContract().functionCall(payload);
    }

    function managerMsgPayload(uint32 msgOrigin, RawCallData memory rcd) public view returns (bytes memory) {
        RawManagerCall memory rmc =
            RawManagerCall({origin: msgOrigin, proofMaturity: block.timestamp - rootSubmittedAt, callData: rcd});
        return rmc.callPayload();
    }

    function remoteRegistrySlashCalldata(uint32 domain, address agent, address prover)
        public
        view
        returns (RawCallData memory)
    {
        // (msgOrigin, proofMaturity) are omitted => (domain, agent, prover)
        return RawCallData({
            selector: bondingManager.remoteRegistrySlash.selector,
            args: abi.encode(domain, agent, prover)
        });
    }

    function remoteWithdrawTipsCalldata(address actor, uint256 amount) public view returns (RawCallData memory) {
        // (msgOrigin, proofMaturity) are omitted => (actor, amount)
        return RawCallData({selector: lightManager.remoteWithdrawTips.selector, args: abi.encode(actor, amount)});
    }

    /// @notice Returns address of the tested system contract
    function systemContract() public view override returns (address) {
        return localAgentManager();
    }

    /// @notice Returns tested system contract as AgentManager
    function testedAM() public view returns (AgentManager) {
        return AgentManager(localAgentManager());
    }
}
