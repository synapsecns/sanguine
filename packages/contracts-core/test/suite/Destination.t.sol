// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SNAPSHOT_MAX_STATES} from "../../contracts/libs/Snapshot.sol";
import {DisputeFlag} from "../../contracts/libs/Structures.sol";
import {ISystemRegistry} from "../../contracts/interfaces/ISystemRegistry.sol";
import {IDisputeHub} from "../../contracts/interfaces/IDisputeHub.sol";

import {InterfaceDestination} from "../../contracts/Destination.sol";
import {Versioned} from "../../contracts/Version.sol";

import {Random} from "../utils/libs/Random.t.sol";
import {RawAttestation, RawState, RawStateIndex} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, SynapseTest} from "../utils/SynapseTest.t.sol";
import {ExecutionHubTest} from "./hubs/ExecutionHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract DestinationTest is ExecutionHubTest {
    // Deploy Production version of Destination and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION) {}

    function test_setupCorrectly() public {
        // Check Agents: currently all Agents are known in LightManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, ISystemRegistry(destination).agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check version
        assertEq(Versioned(destination).version(), LATEST_VERSION, "!version");
        // Check pending Agent Merkle Root
        (bool rootPassed, bool rootPending) = InterfaceDestination(destination).passAgentRoot();
        assertFalse(rootPassed);
        assertFalse(rootPending);
    }

    function test_getSignedAttestation(Random memory random) public {
        uint256 amount = 10;
        bytes[] memory attPayloads = new bytes[](amount);
        bytes[] memory attSignatures = new bytes[](amount);
        for (uint32 index = 0; index < amount; ++index) {
            address notary = domains[localDomain()].agents[random.nextUint256() % DOMAIN_AGENTS];
            RawAttestation memory ra = random.nextAttestation(index + 1);
            (attPayloads[index], attSignatures[index]) = signAttestation(notary, ra);
            InterfaceDestination(destination).submitAttestation(attPayloads[index], attSignatures[index]);
        }
        for (uint32 index = 0; index < amount; ++index) {
            (bytes memory attPayload, bytes memory attSignature) =
                InterfaceDestination(destination).getSignedAttestation(index);
            assertEq(attPayload, attPayloads[index], "!payload");
            assertEq(attSignature, attSignatures[index], "!signature");
        }
    }

    function test_submitAttestation(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        vm.warp(rootSubmittedAt);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        InterfaceDestination(destination).submitAttestation(attPayload, attSig);
        (uint48 snapRootTime,,) = InterfaceDestination(destination).destStatus();
        assertEq(snapRootTime, rootSubmittedAt);
    }

    function test_submitAttestation_updatesAgentRoot(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(ra.agentRoot != agentRootLM);
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        vm.warp(rootSubmittedAt);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        InterfaceDestination(destination).submitAttestation(attPayload, attSig);
        (, uint48 agentRootTime, address statusNotary) = InterfaceDestination(destination).destStatus();
        // Check that values were assigned
        assertEq(InterfaceDestination(destination).nextAgentRoot(), ra.agentRoot);
        assertEq(agentRootTime, rootSubmittedAt);
        assertEq(statusNotary, notary);
    }

    function test_submitAttestation_doesNotOverwritePending(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA,
        uint32 firstRootSubmittedAt,
        uint32 timePassed
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA.agentRoot != agentRootLM);
        vm.assume(firstRA.snapRoot != secondRA.snapRoot);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        timePassed = timePassed % AGENT_ROOT_OPTIMISTIC_PERIOD;
        skip(timePassed);
        // Form a second attestation: Notary 1
        address notaryF = domains[DOMAIN_LOCAL].agents[0];
        address notaryS = domains[DOMAIN_LOCAL].agents[1];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notaryS, secondRA);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notaryS, attPayload, attSig);
        assertTrue(InterfaceDestination(destination).submitAttestation(attPayload, attSig));
        (uint48 snapRootTime, uint48 agentRootTime, address notary) = InterfaceDestination(destination).destStatus();
        assertEq(snapRootTime, block.timestamp);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(notary, notaryF);
    }

    function test_submitAttestation_notAccepted_agentRootUpdated(
        RawAttestation memory firstRA,
        uint32 firstRootSubmittedAt
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA.agentRoot != agentRootLM);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        // Should not accept the attestation before doing any checks,
        // so we could pass empty values here
        assertFalse(InterfaceDestination(destination).submitAttestation("", ""));
        (uint48 snapRootTime, uint48 agentRootTime, address notary) = InterfaceDestination(destination).destStatus();
        assertEq(snapRootTime, firstRootSubmittedAt);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(notary, domains[DOMAIN_LOCAL].agent);
        // Should update the Agent Merkle Root
        assertEq(lightManager.agentRoot(), firstRA.agentRoot);
    }

    function test_submitStateReport_notAccepted_agentRootUpdated(
        RawAttestation memory firstRA,
        uint32 firstRootSubmittedAt
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA.agentRoot != agentRootLM);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        // Should not accept the attestation before doing any checks,
        // so we could pass empty values here
        assertFalse(IDisputeHub(destination).submitStateReport(0, "", "", "", ""));
        (uint48 snapRootTime, uint48 agentRootTime, address notary) = InterfaceDestination(destination).destStatus();
        assertEq(snapRootTime, firstRootSubmittedAt);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(notary, domains[DOMAIN_LOCAL].agent);
        // Should update the Agent Merkle Root
        assertEq(lightManager.agentRoot(), firstRA.agentRoot);
    }

    function test_submitStateReportWithProof_notAccepted_agentRootUpdated(
        RawAttestation memory firstRA,
        uint32 firstRootSubmittedAt
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA.agentRoot != agentRootLM);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        // Should not accept the attestation before doing any checks,
        // so we could pass empty values here
        assertFalse(IDisputeHub(destination).submitStateReportWithProof(0, "", "", new bytes32[](0), "", ""));
        (uint48 snapRootTime, uint48 agentRootTime, address notary) = InterfaceDestination(destination).destStatus();
        assertEq(snapRootTime, firstRootSubmittedAt);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(notary, domains[DOMAIN_LOCAL].agent);
        // Should update the Agent Merkle Root
        assertEq(lightManager.agentRoot(), firstRA.agentRoot);
    }

    function test_passAgentRoot_optimisticPeriodNotOver(
        RawAttestation memory ra,
        uint32 rootSubmittedAt,
        uint32 timePassed
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(ra.agentRoot != agentRootLM);
        // Submit attestation that updates `nextAgentRoot`
        test_submitAttestation_updatesAgentRoot(ra, rootSubmittedAt);
        timePassed = timePassed % AGENT_ROOT_OPTIMISTIC_PERIOD;
        skip(timePassed);
        (bool rootPassed, bool rootPending) = InterfaceDestination(destination).passAgentRoot();
        assertFalse(rootPassed);
        assertTrue(rootPending);
        assertEq(lightManager.agentRoot(), agentRootLM);
    }

    function test_passAgentRoot_optimisticPeriodOver(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        // Submit attestation that updates `nextAgentRoot`
        test_submitAttestation_updatesAgentRoot(ra, rootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        (bool rootPassed, bool rootPending) = InterfaceDestination(destination).passAgentRoot();
        assertTrue(rootPassed);
        assertFalse(rootPending);
        assertEq(lightManager.agentRoot(), ra.agentRoot);
    }

    function test_submitAttestationReport(RawAttestation memory ra) public {
        address prover = makeAddr("Prover");
        // Create Notary signature for the attestation
        address notary = domains[DOMAIN_LOCAL].agent;
        (, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard signature for the report
        address guard = domains[0].agent;
        (bytes memory arPayload, bytes memory arSig) = createSignedAttestationReport(guard, ra);
        vm.expectEmit(true, true, true, true);
        emit Dispute(guard, DOMAIN_LOCAL, notary);
        vm.prank(prover);
        InterfaceDestination(destination).submitAttestationReport(arPayload, arSig, attSig);
        checkDisputeOpened(destination, guard, notary);
    }

    function test_submitStateReport(RawState memory rs, RawStateIndex memory rsi) public boundIndex(rsi) {
        check_submitStateReport(destination, DOMAIN_LOCAL, rs, rsi);
    }

    function test_submitStateReportWithProof(RawState memory rs, RawAttestation memory ra, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        check_submitStateReportWithProof(destination, DOMAIN_LOCAL, rs, ra, rsi);
    }

    // ════════════════════════════════════════════ DISPUTE RESOLUTION ═════════════════════════════════════════════════

    function test_managerSlash(uint256 domainId, uint256 agentId, address prover) public {
        // no counterpart in this test
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        bool isRemoteNotary = !(domain == 0 || domain == DOMAIN_LOCAL);
        if (!isRemoteNotary) {
            vm.expectEmit();
            emit DisputeResolved(address(0), domain, agent);
        }
        vm.expectEmit();
        emit AgentSlashed(domain, agent, prover);
        vm.recordLogs();
        vm.prank(address(lightManager));
        ISystemRegistry(destination).managerSlash(domain, agent, prover);
        if (isRemoteNotary) {
            // Should only emit AgentSlashed for remote Notaries
            assertEq(vm.getRecordedLogs().length, 1);
            assertEq(uint8(IDisputeHub(destination).disputeStatus(agent).flag), uint8(DisputeFlag.None));
        } else {
            assertEq(vm.getRecordedLogs().length, 2);
            checkDisputeResolved({hub: destination, honest: address(0), slashed: agent});
        }
    }

    function test_managerSlash_honestGuard(RawAttestation memory ra) public {
        address guard = domains[0].agent;
        address notary = domains[DOMAIN_LOCAL].agent;
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(ra);
        // Slash the Notary
        vm.prank(address(lightManager));
        ISystemRegistry(destination).managerSlash(DOMAIN_LOCAL, notary, address(0));
        checkDisputeResolved({hub: destination, honest: guard, slashed: notary});
    }

    function test_managerSlash_honestNotary(RawAttestation memory ra) public {
        address guard = domains[0].agent;
        address notary = domains[DOMAIN_LOCAL].agent;
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(ra);
        // Slash the Guard
        vm.prank(address(lightManager));
        ISystemRegistry(destination).managerSlash(0, guard, address(0));
        checkDisputeResolved({hub: destination, honest: notary, slashed: guard});
    }

    // ══════════════════════════════════════════ TESTS: WHILE IN DISPUTE ══════════════════════════════════════════════

    function test_submitAttestation_revert_notaryInDispute(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA
    ) public {
        address notary = domains[DOMAIN_LOCAL].agent;
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(firstRA);
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, secondRA);
        vm.expectRevert("Notary is in dispute");
        InterfaceDestination(destination).submitAttestation(attPayload, attSig);
    }

    function test_submitAttestationReport_revert_guardInDispute(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA
    ) public {
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(firstRA);
        // Try to initiate a dispute between Guard 0 and Notary 1
        address guard = domains[0].agent;
        (bytes memory arPayload, bytes memory arSig) = createSignedAttestationReport(guard, secondRA);
        address notary = domains[DOMAIN_LOCAL].agents[1];
        (, bytes memory attSig) = signAttestation(notary, secondRA);
        vm.expectRevert("Guard already in dispute");
        InterfaceDestination(destination).submitAttestationReport(arPayload, arSig, attSig);
    }

    function test_submitAttestationReport_revert_notaryInDispute(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA
    ) public {
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(firstRA);
        // Try to initiate a dispute between Guard 1 and Notary 0
        address guard = domains[0].agents[1];
        (bytes memory arPayload, bytes memory arSig) = createSignedAttestationReport(guard, secondRA);
        address notary = domains[DOMAIN_LOCAL].agent;
        (, bytes memory attSig) = signAttestation(notary, secondRA);
        vm.expectRevert("Notary already in dispute");
        InterfaceDestination(destination).submitAttestationReport(arPayload, arSig, attSig);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        override
        returns (bytes32 snapRoot, bytes32[] memory snapProof)
    {
        RawAttestation memory ra;
        (ra, snapProof) = createSnapshotProof(sm);
        snapRoot = ra.snapRoot;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(domains[DOMAIN_LOCAL].agent, ra);
        InterfaceDestination(destination).submitAttestation(attPayload, attSig);
    }

    /// @notice Returns local domain for the tested system contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
