// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, AgentStatus, SystemEntity} from "../../../contracts/libs/Structures.sol";
import {InterfaceDestination} from "../../../contracts/interfaces/InterfaceDestination.sol";
import {InterfaceOrigin} from "../../../contracts/interfaces/InterfaceOrigin.sol";

import {AgentManagerTest} from "./AgentManager.t.sol";

import {
    AgentFlag,
    AgentStatus,
    LightManager,
    LightManagerHarness,
    IAgentSecured,
    SynapseTest
} from "../../utils/SynapseTest.t.sol";
import {Random} from "../../utils/libs/Random.t.sol";
import {RawAttestation} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract LightManagerTest is AgentManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        address caller = random.nextAddress();
        address origin_ = random.nextAddress();
        address destination_ = random.nextAddress();
        LightManager cleanContract = new LightManager(domain);
        vm.prank(caller);
        cleanContract.initialize(origin_, destination_);
        assertEq(cleanContract.localDomain(), domain);
        assertEq(cleanContract.owner(), caller);
        assertEq(cleanContract.origin(), origin_);
        assertEq(cleanContract.destination(), destination_);
    }

    function initializeLocalContract() public override {
        LightManager(localContract()).initialize(address(0), address(0));
    }

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.expectRevert("Can't be deployed on SynChain");
        new LightManagerHarness(DOMAIN_SYNAPSE);
    }

    function test_setup() public override {
        super.test_setup();
        assertEq(lightManager.version(), LATEST_VERSION, "!version");
    }

    // ══════════════════════════════════ TESTS: UNAUTHORIZED ACCESS (NOT OWNER) ═══════════════════════════════════════

    function test_setAgentRoot_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        vm.expectRevert("Only Destination sets agent root");
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
        expectDisputeResolved(agent, address(0), caller);
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Slashed);
    }

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
        vm.expectRevert("Invalid proof");
        lightManager.updateAgentStatus(agent, status, proof);
    }

    // ══════════════════════════════════════════ TEST: SUBMIT STATEMENTS ══════════════════════════════════════════════

    function test_submitAttestation(RawAttestation memory ra) public {
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectCall(
            destination,
            abi.encodeWithSelector(
                InterfaceDestination.acceptAttestation.selector, agentIndex[notary], nextSignatureIndex(), attPayload
            )
        );
        lightManager.submitAttestation(attPayload, attSignature);
    }

    function test_submitAttestation_revert_signedByGuard(RawAttestation memory ra) public {
        address guard = domains[0].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(guard, ra);
        vm.expectRevert("Signer is not a Notary");
        lightManager.submitAttestation(attPayload, attSignature);
    }

    function test_submitAttestation_revert_signedByRemoteNotary(RawAttestation memory ra) public {
        address notary = domains[DOMAIN_REMOTE].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        vm.expectRevert("Wrong Notary domain");
        lightManager.submitAttestation(attPayload, attSignature);
    }

    function test_submitAttestation_revert_notaryInDispute(RawAttestation memory ra) public {
        address notary = domains[localDomain()].agent;
        (bytes memory attPayload, bytes memory attSignature) = signAttestation(notary, ra);
        openDispute({guard: domains[0].agent, notary: notary});
        vm.expectRevert("Notary is in dispute");
        lightManager.submitAttestation(attPayload, attSignature);
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
        vm.expectRevert("!destination");
        vm.prank(caller);
        lightManager.remoteWithdrawTips(DOMAIN_SYNAPSE, BONDING_OPTIMISTIC_PERIOD, address(0), 0);
    }

    function test_remoteWithdrawTips_revert_notSynapseChain(uint32 msgOrigin) public {
        vm.assume(msgOrigin != DOMAIN_SYNAPSE);
        skip(BONDING_OPTIMISTIC_PERIOD);
        bytes memory msgPayload = managerMsgPayload(msgOrigin, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert("!synapseDomain");
        managerMsgPrank(msgPayload);
    }

    function test_remoteWithdrawTips_revert_optimisticPeriodNotOver(uint32 proofMaturity) public {
        proofMaturity = proofMaturity % BONDING_OPTIMISTIC_PERIOD;
        skip(proofMaturity);
        bytes memory msgPayload = managerMsgPayload(DOMAIN_SYNAPSE, remoteWithdrawTipsCalldata(address(0), 0));
        vm.expectRevert("!optimisticPeriod");
        managerMsgPrank(msgPayload);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
