// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";
import {AGENT_TREE_HEIGHT} from "../../../contracts/libs/Constants.sol";
import {
    AgentCantBeAdded,
    AgentNotActive,
    AgentNotUnstaking,
    CallerNotSummit,
    MustBeSynapseDomain,
    SlashAgentOptimisticPeriod,
    SynapseDomainForbidden
} from "../../../contracts/libs/Errors.sol";
import {MerkleMath} from "../../../contracts/libs/merkle/MerkleMath.sol";
import {AgentFlag} from "../../../contracts/libs/Structures.sol";
import {AgentManagerTest} from "./AgentManager.t.sol";

import {BondingManager, BondingManagerHarness, SynapseTest} from "../../utils/SynapseTest.t.sol";

import {Random} from "../../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract BondingManagerTest is AgentManagerTest {
    // Deploy mocks for everything except BondingManager and Inbox
    constructor() SynapseTest(0) {}

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        address caller = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        address summit_ = random.nextAddress();
        address inbox_ = random.nextAddress();
        BondingManager cleanContract = new BondingManager(domain);
        vm.prank(caller);
        cleanContract.initialize(origin_, destination_, inbox_, summit_);
        assertEq(cleanContract.localDomain(), domain);
        assertEq(cleanContract.owner(), caller);
        assertEq(cleanContract.origin(), origin_);
        assertEq(cleanContract.destination(), destination_);
        assertEq(cleanContract.inbox(), inbox_);
        assertEq(cleanContract.summit(), summit_);
        assertEq(cleanContract.leafsAmount(), 1);
    }

    function initializeLocalContract() public override {
        BondingManager(localContract()).initialize(address(0), address(0), address(0), address(0));
    }

    function test_constructor_revert_notOnSynapseChain(uint32 domain) public {
        vm.assume(domain != DOMAIN_SYNAPSE);
        vm.expectRevert(MustBeSynapseDomain.selector);
        new BondingManager(domain);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(bondingManager.summit(), localSummit(), "!summit");
        assertEq(bondingManager.version(), LATEST_VERSION, "!version");
    }

    // ══════════════════════════════════ TESTS: UNAUTHORIZED ACCESS (NOT OWNER) ═══════════════════════════════════════

    function test_addAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.addAgent(1, address(1), new bytes32[](0));
    }

    function test_initiateUnstaking_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.initiateUnstaking(1, address(1), new bytes32[](0));
    }

    function test_completeUnstaking_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.completeUnstaking(1, address(1), new bytes32[](0));
    }

    // ═════════════════════════════════════════ TESTS: ADD/REMOVE AGENTS ══════════════════════════════════════════════

    function test_addAgent_fromScratch() public {
        // Deploy fresh instance of BondingManager
        bondingManager = new BondingManagerHarness(DOMAIN_SYNAPSE);
        bondingManager.initialize(originSynapse, destinationSynapse, address(inbox), summit);
        // Try to add all agents one by one
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
                address agent = domains[domain].agents[i];
                bytes32[] memory proof = bondingManager.getProof(agent);
                bondingManager.addAgent(domain, agent, proof);
                checkAgentStatus(agent, bondingManager.agentStatus(agent), AgentFlag.Active);
            }
        }
    }

    function test_addAgent_new(uint32 domain, address agent) public {
        // Notaries on Syn Chain could nto be added
        vm.assume(domain != DOMAIN_SYNAPSE);
        // Should not be an already added agent
        vm.assume(bondingManager.agentStatus(agent).flag == AgentFlag.Unknown);
        vm.assume(agent != address(0));
        bytes32[] memory proof = getZeroProof();
        bytes32 newRoot = addNewAgent(domain, agent);
        expectStatusUpdated(AgentFlag.Active, domain, agent);
        vm.expectEmit();
        emit RootUpdated(newRoot);
        bondingManager.addAgent(domain, agent, proof);
        checkAgentStatus(agent, bondingManager.agentStatus(agent), AgentFlag.Active);
        assertEq(bondingManager.agentRoot(), newRoot, "!agentRoot");
    }

    function test_addAgent_revert_synapseDomain(address agent) public {
        bytes32[] memory proof = getZeroProof();
        vm.expectRevert(SynapseDomainForbidden.selector);
        bondingManager.addAgent(DOMAIN_SYNAPSE, agent, proof);
    }

    function test_addAgent_resting(uint256 domainId, uint256 agentId) public {
        // Full lifecycle for a live agent:
        // Active -> Unstaking -> Resting -> Active
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatus(AgentFlag.Resting, domain, agent);
        updateStatus(AgentFlag.Active, domain, agent);
    }

    function test_initiateUnstaking(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
    }

    function test_completeUnstaking(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatus(AgentFlag.Resting, domain, agent);
    }

    function updateStatus(AgentFlag flag, uint32 domain, address agent) public {
        updateStatus(address(this), flag, domain, agent);
    }

    function updateStatus(address caller, AgentFlag flag, uint32 domain, address agent) public {
        bytes32[] memory proof = getAgentProof(agent);
        bytes32 newRoot = updateAgent(flag, agent);
        expectStatusUpdated(flag, domain, agent);
        vm.expectEmit();
        emit RootUpdated(newRoot);
        vm.prank(caller);
        updateStatusWithProof(flag, domain, agent, proof);
        assertEq(bondingManager.agentRoot(), newRoot, "!agentRoot");
        checkAgentStatus(agent, bondingManager.agentStatus(agent), flag);
    }

    function updateStatusWithProof(AgentFlag flag, uint32 domain, address agent, bytes32[] memory proof) public {
        if (flag == AgentFlag.Unstaking) {
            bondingManager.initiateUnstaking(domain, agent, proof);
        } else if (flag == AgentFlag.Resting) {
            bondingManager.completeUnstaking(domain, agent, proof);
        } else if (flag == AgentFlag.Active) {
            bondingManager.addAgent(domain, agent, proof);
        } else if (flag == AgentFlag.Slashed) {
            bondingManager.completeSlashing(domain, agent, proof);
        }
    }

    // ═══════════════════════════════════════ TEST: UPDATE AGENTS (REVERTS) ═══════════════════════════════════════════

    function test_addAgent_revert_active(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatusWithRevert(AgentFlag.Active, domain, agent, AgentCantBeAdded.selector);
    }

    function test_addAgent_revert_unstaking(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatusWithRevert(AgentFlag.Active, domain, agent, AgentCantBeAdded.selector);
    }

    function test_initiateUnstaking_revert_unstaking(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatusWithRevert(AgentFlag.Unstaking, domain, agent, AgentNotActive.selector);
    }

    function test_initiateUnstaking_revert_resting(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatus(AgentFlag.Resting, domain, agent);
        updateStatusWithRevert(AgentFlag.Unstaking, domain, agent, AgentNotActive.selector);
    }

    function test_completeUnstaking_revert_active(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatusWithRevert(AgentFlag.Resting, domain, agent, AgentNotUnstaking.selector);
    }

    function test_completeUnstaking_revert_resting(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        updateStatus(AgentFlag.Resting, domain, agent);
        updateStatusWithRevert(AgentFlag.Resting, domain, agent, AgentNotUnstaking.selector);
    }

    function updateStatusWithRevert(AgentFlag flag, uint32 domain, address agent, bytes4 err) public {
        bytes32[] memory proof = getAgentProof(agent);
        vm.expectRevert(err);
        updateStatusWithProof(flag, domain, agent, proof);
    }

    // ═══════════════════════════════════════════ TEST: SLASHING AGENTS ═══════════════════════════════════════════════

    // TODO: test_initiateSlashing

    function test_remoteSlashAgent(uint32 msgOrigin, uint256 domainId, uint256 agentId, address prover) public {
        // Needs to be a REMOTE call
        vm.assume(msgOrigin != DOMAIN_SYNAPSE);
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        skipBondingOptimisticPeriod();
        bytes memory msgPayload = managerMsgPayload(msgOrigin, remoteSlashAgentCalldata(domain, agent, prover));
        expectStatusUpdated(AgentFlag.Fraudulent, domain, agent);
        expectDisputeResolved(0, agent, address(0), prover);
        managerMsgPrank(msgPayload);
        assertEq(uint8(bondingManager.agentStatus(agent).flag), uint8(AgentFlag.Fraudulent));
        // (bool isSlashed, address prover_) = bondingManager.slashStatus(agent);
        // assertTrue(isSlashed);
        // assertEq(prover_, prover);
    }

    function test_remoteSlashAgent_revert_optimisticPeriodNotOver(uint32 proofMaturity) public {
        proofMaturity = proofMaturity % BONDING_OPTIMISTIC_PERIOD;
        skip(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_REMOTE, remoteSlashAgentCalldata(0, address(0), address(0)));
        vm.expectRevert(SlashAgentOptimisticPeriod.selector);
        managerMsgPrank(msgPayload);
    }

    function test_completeSlashing_active(uint256 domainId, uint256 agentId, address slasher) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        // Initiate slashing
        test_remoteSlashAgent(DOMAIN_REMOTE, domainId, agentId, address(1));
        updateStatus(slasher, AgentFlag.Slashed, domain, agent);
        checkAgentStatus(agent, bondingManager.agentStatus(agent), AgentFlag.Slashed);
    }

    function test_completeSlashing_unstaking(uint256 domainId, uint256 agentId, address slasher) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        updateStatus(AgentFlag.Unstaking, domain, agent);
        // Initiate slashing
        test_remoteSlashAgent(DOMAIN_REMOTE, domainId, agentId, address(1));
        updateStatus(slasher, AgentFlag.Slashed, domain, agent);
        checkAgentStatus(agent, bondingManager.agentStatus(agent), AgentFlag.Slashed);
    }

    // ════════════════════════════════════════════ TEST: WITHDRAW TIPS ════════════════════════════════════════════════

    function test_withdrawTips_local(address recipient, uint256 amount) public {
        bytes memory expectedCall = abi.encodeWithSelector(InterfaceOrigin.withdrawTips.selector, recipient, amount);
        vm.expectCall(originSynapse, expectedCall);
        vm.prank(summit);
        bondingManager.withdrawTips(recipient, DOMAIN_SYNAPSE, amount);
    }

    function test_withdrawTips_remote(address recipient, uint32 domain, uint256 amount) public {
        vm.assume(domain != DOMAIN_SYNAPSE);
        // remoteWithdrawTips(msgOrigin, proofMaturity, recipient, amount), but first two are omitted
        bytes memory payload = abi.encodeWithSelector(lightManager.remoteWithdrawTips.selector, recipient, amount);
        // sendManagerMessage(destination, optimisticPeriod, payload)
        bytes memory expectedCall = abi.encodeWithSelector(
            InterfaceOrigin.sendManagerMessage.selector, domain, BONDING_OPTIMISTIC_PERIOD, payload
        );
        vm.expectCall(address(originSynapse), expectedCall);
        vm.prank(summit);
        bondingManager.withdrawTips(recipient, domain, amount);
    }

    function test_withdrawTips_revert_notSummit(address caller) public {
        vm.assume(caller != summit);
        vm.expectRevert(CallerNotSummit.selector);
        vm.prank(caller);
        bondingManager.withdrawTips(address(0), 0, 0);
    }

    // ════════════════════════════════════════════════ TEST: VIEWS ════════════════════════════════════════════════════

    function test_agentLeaf_knownAgent(uint256 domainId, uint256 agentId) public {
        (, address agent) = getAgent(domainId, agentId);
        assertEq(bondingManager.agentLeaf(agent), getAgentLeaf(agentIndex[agent]));
    }

    function test_agentLeaf_unknownAgent(address agent) public {
        // Should not be an already added agent
        vm.assume(bondingManager.agentStatus(agent).flag == AgentFlag.Unknown);
        assertEq(bondingManager.agentLeaf(agent), bytes32(0));
    }

    function test_getActiveAgents() public {
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            address[] memory agents = bondingManager.getActiveAgents(domain);
            assertEq(agents.length, DOMAIN_AGENTS);
            for (uint256 i = 0; i < agents.length; ++i) {
                assertEq(agents[i], domains[domain].agents[i]);
            }
        }
    }

    function test_getActiveAgents_agentsRemoved() public {
        // Change status of four agents into Unstaking, Resting, Fraudulent and Slashed - one for each domain
        test_initiateUnstaking(0, 0);
        test_completeUnstaking(1, 1);
        test_remoteSlashAgent(DOMAIN_REMOTE, 2, 2, address(1));
        test_completeSlashing_active(3, 3, address(1));
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            address[] memory agents = bondingManager.getActiveAgents(domain);
            assertEq(agents.length, DOMAIN_AGENTS - 1);
            for (uint256 i = 0; i < agents.length; ++i) {
                // Agent with index `d` was removed
                assertEq(agents[i], domains[domain].agents[i < d ? i : i + 1]);
            }
        }
    }

    function test_getProof_knownAgent(uint256 domainId, uint256 agentId) public {
        (, address agent) = getAgent(domainId, agentId);
        bytes32[] memory proof = bondingManager.getProof(agent);
        uint256 index = agentIndex[agent];
        checkProof(index, proof);
    }

    function test_getProof_unknownAgent(address agent) public {
        // Should not be an already added agent
        vm.assume(bondingManager.agentStatus(agent).flag == AgentFlag.Unknown);
        bytes32[] memory proof = bondingManager.getProof(agent);
        // Use the next index
        uint256 index = totalAgents + 1;
        checkProof(index, proof);
    }

    function checkProof(uint256 index, bytes32[] memory proof) public {
        assertEq(MerkleMath.proofRoot(index, getAgentLeaf(index), proof, AGENT_TREE_HEIGHT), getAgentRoot());
    }

    function test_allLeafs() public {
        assertEq(bondingManager.leafsAmount(), totalAgents + 1, "!leafsAmount");
        bytes32[] memory leafs = bondingManager.allLeafs();
        for (uint256 i = 0; i < leafs.length; ++i) {
            assertEq(leafs[i], getAgentLeaf(i));
        }
    }

    function test_getLeafs(uint256 indexFrom, uint256 amount) public {
        uint256 totalLeafs = totalAgents + 1;
        indexFrom = indexFrom % totalLeafs;
        // Allow index overrun
        amount = amount % (totalLeafs + 10);
        bytes32[] memory leafs = bondingManager.getLeafs(indexFrom, amount);
        if (indexFrom + amount <= totalLeafs) {
            assertEq(leafs.length, amount);
        } else {
            assertEq(leafs.length, totalLeafs - indexFrom);
        }
        for (uint256 i = 0; i < leafs.length; ++i) {
            assertEq(leafs[i], getAgentLeaf(indexFrom + i));
        }
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
