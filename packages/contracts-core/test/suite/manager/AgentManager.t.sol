// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {FRESH_DATA_TIMEOUT} from "../../../contracts/libs/Constants.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "../../../contracts/libs/Structures.sol";

import {InterfaceDestination} from "../../../contracts/interfaces/InterfaceDestination.sol";
import {IStatementInbox} from "../../../contracts/interfaces/IStatementInbox.sol";

import {MessagingBaseTest} from "../base/MessagingBase.t.sol";
import {AgentManagerHarness} from "../../harnesses/manager/AgentManagerHarness.t.sol";

import {RawCallData, RawManagerCall, RawSnapshot, RawState} from "../../utils/libs/SynapseStructs.t.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Random} from "../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract AgentManagerTest is MessagingBaseTest {
    using Address for address;

    uint256 internal rootSubmittedAt;

    function test_setup() public virtual {
        assertEq(address(testedAM().destination()), localDestination());
        assertEq(address(testedAM().origin()), localOrigin());
        assertEq(testedAM().agentRoot(), getAgentRoot());
    }

    function mockSnapRootTime(uint256 timePassed) public {
        // Force destStatus() to return (timestamp, timestamp, 1) as (snapRootTime, agentRootTime, notaryIndex)
        vm.mockCall(
            localDestination(),
            abi.encodeWithSelector(InterfaceDestination.destStatus.selector),
            abi.encode(block.timestamp, block.timestamp, 1)
        );
        skip(timePassed);
    }

    // ══════════════════════════════════════════════ TESTS: DISPUTES ══════════════════════════════════════════════════

    function test_openDispute(Random memory random) public {
        address guard = randomGuard(random);
        address notary = randomNotary(random);
        RawState memory rs = random.nextState();
        RawSnapshot memory rawSnap;
        rawSnap.states = new RawState[](1);
        rawSnap.states[0] = rs;
        (bytes memory snapPayload, bytes memory snapSignature) = signSnapshot(notary, rawSnap);
        (bytes memory statePayload, bytes memory srSignature) = signStateReport(guard, rs);
        assertEq(testedAM().getDisputesAmount(), 0);
        expectDisputeOpened(0, guard, notary);
        IStatementInbox(localInbox()).submitStateReportWithSnapshot(0, srSignature, snapPayload, snapSignature);
        assertEq(testedAM().getDisputesAmount(), 1);
        // Scope to get around stack too deep error
        {
            (
                address guard_,
                address notary_,
                address slashedAgent,
                address fraudProver,
                bytes memory reportPayload,
                bytes memory reportSignature
            ) = testedAM().getDispute(0);
            assertEq(guard_, guard);
            assertEq(notary_, notary);
            assertEq(slashedAgent, address(0));
            assertEq(fraudProver, address(0));
            assertEq(reportPayload, statePayload);
            assertEq(reportSignature, srSignature);
        }
        checkDisputeStatus(guard, DisputeFlag.Pending, notary, address(0), 1);
        checkDisputeStatus(notary, DisputeFlag.Pending, guard, address(0), 1);
    }

    function test_resolveDispute(Random memory random) public {
        test_openDispute(random);
        (address guard, address notary,,, bytes memory reportPayload, bytes memory reportSignature) =
            testedAM().getDispute(0);
        // Pick a random agent in Dispute to slash
        address slashedAgent = random.nextUint256() % 2 == 0 ? guard : notary;
        address rival = slashedAgent == guard ? notary : guard;
        address fraudProver = random.nextAddress();
        expectStatusUpdated(AgentFlag.Fraudulent, agentDomain[slashedAgent], slashedAgent);
        expectDisputeResolved(1, slashedAgent, rival, fraudProver);
        vm.prank(localInbox());
        testedAM().slashAgent(agentDomain[slashedAgent], slashedAgent, fraudProver);
        assertEq(testedAM().getDisputesAmount(), 1);
        (
            address guard_,
            address notary_,
            address slashedAgent_,
            address fraudProver_,
            bytes memory reportPayload_,
            bytes memory reportSignature_
        ) = testedAM().getDispute(0);
        assertEq(guard_, guard);
        assertEq(notary_, notary);
        assertEq(slashedAgent_, slashedAgent);
        assertEq(fraudProver_, fraudProver);
        assertEq(reportPayload_, reportPayload);
        assertEq(reportSignature_, reportSignature);
        checkDisputeStatus(slashedAgent, DisputeFlag.Slashed, rival, fraudProver, 1);
        checkDisputeStatus(rival, DisputeFlag.None, address(0), address(0), 0);
    }

    function test_slashAgentWithoutDispute(Random memory random) public {
        address slashedAgent = randomAgent(random);
        address fraudProver = random.nextAddress();
        expectStatusUpdated(AgentFlag.Fraudulent, agentDomain[slashedAgent], slashedAgent);
        expectDisputeResolved(0, slashedAgent, address(0), fraudProver);
        vm.recordLogs();
        vm.prank(localInbox());
        testedAM().slashAgent(agentDomain[slashedAgent], slashedAgent, fraudProver);
        // Should only emit StatusUpdated
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(testedAM().getDisputesAmount(), 0);
        checkDisputeStatus(slashedAgent, DisputeFlag.Slashed, address(0), fraudProver, 0);
    }

    function checkDisputeStatus(address agent, DisputeFlag flag, address rival, address prover, uint256 disputePtr)
        public
    {
        (DisputeFlag flag_, address rival_, address prover_, uint256 disputePtr_) = testedAM().disputeStatus(agent);
        assertEq(uint8(flag_), uint8(flag), "!flag");
        assertEq(rival_, rival, "!rival");
        assertEq(prover_, prover, "!fraudProver");
        assertEq(disputePtr_, disputePtr, "!disputePtr");
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
        localContract().functionCall(payload);
    }

    function managerMsgPayload(uint32 msgOrigin, RawCallData memory rcd) public view returns (bytes memory) {
        RawManagerCall memory rmc =
            RawManagerCall({origin: msgOrigin, proofMaturity: block.timestamp - rootSubmittedAt, callData: rcd});
        return rmc.callPayload();
    }

    function remoteSlashAgentCalldata(uint32 domain, address agent, address prover)
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

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return localAgentManager();
    }

    /// @notice Returns tested contract as AgentManager
    function testedAM() public view returns (AgentManagerHarness) {
        return AgentManagerHarness(localAgentManager());
    }
}
