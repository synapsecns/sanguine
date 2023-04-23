// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ISystemRegistry} from "../../../contracts/interfaces/ISystemRegistry.sol";
import {AgentFlag, AgentStatus, SlashStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";

import {SystemContractTest} from "../system/SystemContract.t.sol";
import {AgentManagerHarness} from "../../harnesses/manager/AgentManagerHarness.t.sol";
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
        return
            RawCallData({selector: bondingManager.remoteSlashAgent.selector, args: abi.encode(domain, agent, prover)});
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
    function testedAM() public view returns (AgentManagerHarness) {
        return AgentManagerHarness(localAgentManager());
    }
}
