// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AGENT_ROOT_PROPOSAL_TIMEOUT, FRESH_DATA_TIMEOUT} from "../../../contracts/libs/Constants.sol";
import {
    AgentRootNotProposed,
    AgentRootTimeoutNotOver,
    CallerNotDestination,
    IncorrectAgentProof,
    IncorrectAgentRoot,
    MustBeSynapseDomain,
    NotStuck,
    SynapseDomainForbidden,
    WithdrawTipsOptimisticPeriod
} from "../../../contracts/libs/Errors.sol";
import {AgentFlag, AgentStatus} from "../../../contracts/libs/Structures.sol";
import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";

import {AgentManagerTest} from "./AgentManager.t.sol";

import {AgentFlag, AgentStatus, LightManager, LightManagerHarness, SynapseTest} from "../../utils/SynapseTest.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract LightManagerTest is AgentManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        vm.chainId(domain);
        address caller = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        address inbox_ = random.nextAddress();
        address owner_ = random.nextAddress();
        LightManager cleanContract = new LightManager(DOMAIN_SYNAPSE);
        vm.prank(caller);
        cleanContract.initialize(origin_, destination_, inbox_, owner_);
        assertEq(cleanContract.localDomain(), domain);
        assertEq(cleanContract.owner(), owner_);
        assertEq(cleanContract.origin(), origin_);
        assertEq(cleanContract.destination(), destination_);
        assertEq(cleanContract.inbox(), inbox_);
    }

    function initializeLocalContract() public override {
        LightManager(localContract()).initialize(address(0), address(0), address(0), address(0));
    }

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new LightManager({synapseDomain_: 1});
    }

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.chainId(DOMAIN_SYNAPSE);
        vm.expectRevert(SynapseDomainForbidden.selector);
        new LightManagerHarness(DOMAIN_SYNAPSE);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(lightManager.version(), LATEST_VERSION, "!version");
    }

    // ══════════════════════════════════ TESTS: UNAUTHORIZED ACCESS (NOT OWNER) ═══════════════════════════════════════

    function test_setAgentRoot_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        vm.expectRevert(CallerNotDestination.selector);
        vm.prank(caller);
        lightManager.setAgentRoot(bytes32(uint256(1)));
    }

    // ═════════════════════════════════════════ TESTS: ADD/REMOVE AGENTS ══════════════════════════════════════════════

    function test_addAgent_new(address caller, uint32 domain, address agent) public {
        // Should not be an already added agent
        vm.assume(agent != address(0));
        vm.assume(lightManager.agentStatus(agent).flag == AgentFlag.Unknown);
        bytes32 root = addNewAgent(domain, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        expectStatusUpdated(AgentFlag.Active, domain, agent);
        // Anyone could add agents in Light Manager
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Active);
    }

    function test_updateAgentStatus_slashed(address caller, uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        // Set flag to Slashed in the Merkle Tree
        bytes32 root = updateAgent(AgentFlag.Slashed, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        expectStatusUpdated(AgentFlag.Slashed, domain, agent);
        expectDisputeResolved(0, agent, address(0), caller);
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Slashed);
    }

    // ═══════════════════════════════════════════ TESTS: SET AGENT ROOT ═══════════════════════════════════════════════

    function test_setAgentRoot(bytes32 root) public {
        bool isDifferent = root != lightManager.agentRoot();
        if (isDifferent) {
            vm.expectEmit();
            emit RootUpdated(root);
        }
        vm.recordLogs();
        vm.prank(destination);
        lightManager.setAgentRoot(root);
        if (!isDifferent) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't have");
        }
        assertEq(lightManager.agentRoot(), root, "!agentRoot");
    }

    function test_setAgentRoot_equal() public {
        test_setAgentRoot(lightManager.agentRoot());
    }

    // ════════════════════════════════════ TESTS: SET AGENT ROOT (WHEN STUCK) ═════════════════════════════════════════

    function checkProposedAgentData(bytes32 expectedAgentRoot, uint256 expectedProposedAt) public {
        (bytes32 agentRoot, uint256 proposedAt) = lightManager.proposedAgentRootData();
        assertEq(agentRoot, expectedAgentRoot, "!agentRoot");
        assertEq(proposedAt, expectedProposedAt, "!proposedAt");
    }

    function test_proposeAgentRootWhenStuck_revert_notOwner(address caller) public {
        vm.assume(caller != lightManager.owner());
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.proposeAgentRootWhenStuck("root");
    }

    function test_resolveProposedAgentRoot_revert_notOwner(address caller) public {
        vm.assume(caller != lightManager.owner());
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.resolveProposedAgentRoot();
    }

    function test_cancelProposedAgentRoot_revert_notOwner(address caller) public {
        vm.assume(caller != lightManager.owner());
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.cancelProposedAgentRoot();
    }

    function test_proposedAgentRootDataEmpty() public {
        checkProposedAgentData({expectedAgentRoot: 0, expectedProposedAt: 0});
    }

    function test_proposeAgentRootWhenStuck() public {
        bytes32 oldRoot = lightManager.agentRoot();
        vm.warp(1234);
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 expectedAgentRoot = keccak256("mock root");
        uint256 expectedProposedAt = block.timestamp;
        vm.expectEmit(address(lightManager));
        emit AgentRootProposed(expectedAgentRoot);
        lightManager.proposeAgentRootWhenStuck(expectedAgentRoot);
        checkProposedAgentData(expectedAgentRoot, expectedProposedAt);
        assertEq(lightManager.agentRoot(), oldRoot);
    }

    function test_proposeAgentRootWhenStuck_proposedTwice() public {
        bytes32 oldRoot = lightManager.agentRoot();
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("first root");
        skip(1 hours);
        bytes32 expectedAgentRoot = keccak256("second root");
        uint256 expectedProposedAt = block.timestamp;
        vm.expectEmit(address(lightManager));
        emit AgentRootProposed(expectedAgentRoot);
        lightManager.proposeAgentRootWhenStuck(expectedAgentRoot);
        checkProposedAgentData(expectedAgentRoot, expectedProposedAt);
        assertEq(lightManager.agentRoot(), oldRoot);
    }

    function test_proposeAgentRootWhenStuck_proposedTwice_revert_chainUnstuck() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("first root");
        skip(1 hours);
        mockSnapRootTime(0);
        vm.expectRevert(NotStuck.selector);
        lightManager.proposeAgentRootWhenStuck("second root");
    }

    function test_proposeAgentRootWhenStuck_revert_emptyRoot() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        vm.expectRevert(IncorrectAgentRoot.selector);
        lightManager.proposeAgentRootWhenStuck(0);
    }

    function test_proposeAgentRootWhenStuck_revert_notStuck() public {
        bytes32 newRoot = keccak256("mock root");
        mockSnapRootTime(FRESH_DATA_TIMEOUT - 1);
        vm.expectRevert(NotStuck.selector);
        lightManager.proposeAgentRootWhenStuck(newRoot);
    }

    function test_cancelProposedAgentRoot() public {
        bytes32 oldRoot = lightManager.agentRoot();
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 root = "mock root";
        lightManager.proposeAgentRootWhenStuck(root);
        skip(1 hours);
        // This should cancel the proposed agent root and the timestamp
        vm.expectEmit(address(lightManager));
        emit ProposedAgentRootCancelled(root);
        lightManager.cancelProposedAgentRoot();
        checkProposedAgentData(0, 0);
        assertEq(lightManager.agentRoot(), oldRoot);
    }

    function test_cancelProposedAgentRoot_chainUnstuck() public {
        bytes32 oldRoot = lightManager.agentRoot();
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 newRoot = keccak256("mock root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(1 hours);
        mockSnapRootTime(0);
        // This should cancel the proposed agent root and the timestamp
        vm.expectEmit(address(lightManager));
        emit ProposedAgentRootCancelled(newRoot);
        lightManager.cancelProposedAgentRoot();
        checkProposedAgentData(0, 0);
        assertEq(lightManager.agentRoot(), oldRoot);
    }

    function test_cancelProposedAgentRoot_revert_notProposed() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        vm.expectRevert(AgentRootNotProposed.selector);
        lightManager.cancelProposedAgentRoot();
    }

    function test_cancelProposedAgentRoot_revert_alreadyResolved() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("mock root");
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        lightManager.resolveProposedAgentRoot();
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        vm.expectRevert(AgentRootNotProposed.selector);
        lightManager.cancelProposedAgentRoot();
    }

    function test_resolveProposedAgentRoot() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 newRoot = keccak256("mock root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        // Should emit two events: one signaling the new root, and another one signaling the manual resolution
        vm.expectEmit(address(lightManager));
        emit RootUpdated(newRoot);
        vm.expectEmit(address(lightManager));
        emit ProposedAgentRootResolved(newRoot);
        lightManager.resolveProposedAgentRoot();
        checkProposedAgentData(0, 0);
        assertEq(lightManager.agentRoot(), newRoot);
    }

    /// @dev Should proceed with the proposed root, even if new Notary data is available.
    /// This is done to prevent rogue Notaries from going offline and then
    /// indefinitely blocking the agent root resolution.
    function test_resolveProposedAgentRoot_chainUnstuck() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 newRoot = keccak256("mock root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        mockSnapRootTime(0);
        // Should emit two events: one signaling the new root, and another one signaling the manual resolution
        vm.expectEmit(address(lightManager));
        emit RootUpdated(newRoot);
        vm.expectEmit(address(lightManager));
        emit ProposedAgentRootResolved(newRoot);
        lightManager.resolveProposedAgentRoot();
        checkProposedAgentData(0, 0);
        assertEq(lightManager.agentRoot(), newRoot);
    }

    function test_resolveProposedAgentRoot_revert_timeoutNotOver() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        bytes32 newRoot = keccak256("mock root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT - 1);
        vm.expectRevert(AgentRootTimeoutNotOver.selector);
        lightManager.resolveProposedAgentRoot();
    }

    function test_resolveProposedAgentRoot_proposedTwice() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("first root");
        skip(1 hours);
        bytes32 newRoot = keccak256("second root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        // Should emit two events: one signaling the new root, and another one signaling the manual resolution
        vm.expectEmit(address(lightManager));
        emit RootUpdated(newRoot);
        vm.expectEmit(address(lightManager));
        emit ProposedAgentRootResolved(newRoot);
        lightManager.resolveProposedAgentRoot();
        checkProposedAgentData(0, 0);
        assertEq(lightManager.agentRoot(), newRoot);
    }

    function test_resolveProposedAgentRoot_proposedTwice_revert_timeoutNotOver() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("first root");
        skip(1 hours);
        bytes32 newRoot = keccak256("second root");
        lightManager.proposeAgentRootWhenStuck(newRoot);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT - 1);
        vm.expectRevert(AgentRootTimeoutNotOver.selector);
        lightManager.resolveProposedAgentRoot();
    }

    function test_proposeAgentRootWhenStuck_proposedTwice_cancelled_revert_notProposed() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("first root");
        skip(1 hours);
        lightManager.cancelProposedAgentRoot();
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        vm.expectRevert(AgentRootNotProposed.selector);
        lightManager.resolveProposedAgentRoot();
    }

    function test_resolveProposedAgentRoot_revert_notProposed() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        vm.expectRevert(AgentRootNotProposed.selector);
        lightManager.resolveProposedAgentRoot();
    }

    function test_resolveProposedAgentRoot_revert_alreadyResolved() public {
        mockSnapRootTime(FRESH_DATA_TIMEOUT);
        lightManager.proposeAgentRootWhenStuck("mock root");
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        lightManager.resolveProposedAgentRoot();
        skip(AGENT_ROOT_PROPOSAL_TIMEOUT);
        vm.expectRevert(AgentRootNotProposed.selector);
        lightManager.resolveProposedAgentRoot();
    }

    // ═══════════════════════════════════════ TEST: UPDATE AGENTS (REVERTS) ═══════════════════════════════════════════

    function test_addAgent_revert_invalidProof(uint256 domainId, uint256 agentId) public {
        (, address agent) = getAgent(domainId, agentId);
        bytes32[] memory proof = getAgentProof(agent);
        AgentStatus memory status = getAgentStatus(agent);
        // This succeeds, but doesn't do anything, as agent was already added
        lightManager.updateAgentStatus(agent, status, proof);
        // Change agent root, so old proofs are no longer valid
        test_setAgentRoot(bytes32(0));
        assertEq(uint8(lightManager.agentStatus(agent).flag), uint8(AgentFlag.Unknown));
        vm.expectRevert(IncorrectAgentProof.selector);
        lightManager.updateAgentStatus(agent, status, proof);
    }

    // ════════════════════════════════════════════ TEST: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_remoteWithdrawTips(address actor, uint256 amount, uint32 proofMaturity) public {
        proofMaturity = uint32(bound(proofMaturity, BONDING_OPTIMISTIC_PERIOD, type(uint32).max));
        skip(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_SYNAPSE, remoteWithdrawTipsCalldata(actor, amount));
        bytes memory expectedCall = abi.encodeWithSelector(InterfaceOrigin.withdrawTips.selector, actor, amount);
        vm.expectCall(origin, expectedCall);
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        skip(BONDING_OPTIMISTIC_PERIOD);
        vm.expectRevert(CallerNotDestination.selector);
        vm.prank(caller);
        lightManager.remoteWithdrawTips(DOMAIN_SYNAPSE, BONDING_OPTIMISTIC_PERIOD, address(0), 0);
    }

    function test_remoteWithdrawTips_revert_notSynapseChain(uint32 msgOrigin) public {
        vm.assume(msgOrigin != DOMAIN_SYNAPSE);
        skip(BONDING_OPTIMISTIC_PERIOD);
        bytes memory msgPayload = managerMsgPayload(msgOrigin, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert(MustBeSynapseDomain.selector);
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_optimisticPeriodNotOver(uint32 proofMaturity) public {
        proofMaturity = proofMaturity % BONDING_OPTIMISTIC_PERIOD;
        skipPeriod(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_SYNAPSE, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert(WithdrawTipsOptimisticPeriod.selector);
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_optimisticPeriodMinus1Second() public {
        test_remoteWithdrawTips_revert_optimisticPeriodNotOver(BONDING_OPTIMISTIC_PERIOD - 1);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
