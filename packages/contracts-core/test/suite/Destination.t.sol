// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SNAPSHOT_MAX_STATES} from "../../contracts/libs/Snapshot.sol";
import {DisputeFlag} from "../../contracts/libs/Structures.sol";
import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";

import {ChainGas, GasData, InterfaceDestination} from "../../contracts/Destination.sol";
import {Versioned} from "../../contracts/base/Version.sol";

import {fakeSnapshot} from "../utils/libs/FakeIt.t.sol";
import {Random} from "../utils/libs/Random.t.sol";
import {RawAttestation, RawSnapshot, RawState, RawStateIndex} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, AgentStatus, Destination, LightManager, SynapseTest} from "../utils/SynapseTest.t.sol";
import {ExecutionHubTest} from "./hubs/ExecutionHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract DestinationTest is ExecutionHubTest {
    // Deploy Production version of Destination and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION) {}

    // ═══════════════════════════════════════════════ TESTS: SETUP ════════════════════════════════════════════════════

    function test_setupCorrectly() public {
        // Check Agents: currently all Agents are known in LightManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, IAgentSecured(localDestination()).agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check version
        assertEq(Versioned(localDestination()).version(), LATEST_VERSION, "!version");
        // Check pending Agent Merkle Root
        (bool rootPassed, bool rootPending) = InterfaceDestination(localDestination()).passAgentRoot();
        assertFalse(rootPassed);
        assertFalse(rootPending);
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        address caller = random.nextAddress();
        LightManager agentManager = new LightManager(domain);
        bytes32 agentRoot = random.next();
        Destination cleanContract = new Destination(domain, address(agentManager));
        agentManager.initialize(address(0), address(cleanContract));
        vm.prank(caller);
        cleanContract.initialize(agentRoot);
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), address(agentManager), "!agentManager");
        assertEq(cleanContract.nextAgentRoot(), agentRoot, "!nextAgentRoot");
    }

    function initializeLocalContract() public override {
        Destination(localContract()).initialize(0);
    }

    // ════════════════════════════════════════════════ OTHER TESTS ════════════════════════════════════════════════════

    function test_getAttestation(Random memory random) public {
        uint256 amount = 10;
        bytes[] memory attPayloads = new bytes[](amount);
        bytes[] memory attSignatures = new bytes[](amount);
        for (uint32 index = 0; index < amount; ++index) {
            address notary = domains[localDomain()].agents[random.nextUint256() % DOMAIN_AGENTS];
            RawSnapshot memory rs = random.nextSnapshot();
            RawAttestation memory ra = random.nextAttestation(rs, index + 1);
            uint256[] memory snapGas = rs.snapGas();
            (attPayloads[index], attSignatures[index]) = signAttestation(notary, ra);
            lightManager.submitAttestation(attPayloads[index], attSignatures[index], ra._agentRoot, snapGas);
        }
        for (uint32 index = 0; index < amount; ++index) {
            (bytes memory attPayload, bytes memory attSignature) =
                InterfaceDestination(localDestination()).getAttestation(index);
            assertEq(attPayload, attPayloads[index], "!payload");
            assertEq(attSignature, attSignatures[index], "!signature");
        }
    }

    function test_submitAttestation(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        RawSnapshot memory rs = Random(ra.snapRoot).nextSnapshot();
        ra._snapGasHash = rs.snapGasHash();
        ra.setDataHash();
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        vm.warp(rootSubmittedAt);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        (uint48 snapRootTime,,) = InterfaceDestination(localDestination()).destStatus();
        assertEq(snapRootTime, rootSubmittedAt);
    }

    function test_submitAttestation_updatesAgentRoot(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        RawSnapshot memory rs = Random(ra.snapRoot).nextSnapshot();
        ra._snapGasHash = rs.snapGasHash();
        ra.setDataHash();
        uint256[] memory snapGas = rs.snapGas();
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(ra._agentRoot != agentRootLM);
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        vm.warp(rootSubmittedAt);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        (, uint40 agentRootTime, uint32 index) = InterfaceDestination(localDestination()).destStatus();
        // Check that values were assigned
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), ra._agentRoot);
        assertEq(agentRootTime, rootSubmittedAt);
        assertEq(index, agentIndex[notary]);
    }

    function test_submitAttestation_doesNotOverwritePending(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA,
        uint32 firstRootSubmittedAt,
        uint32 timePassed
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA._agentRoot != agentRootLM);
        vm.assume(firstRA.snapRoot != secondRA.snapRoot);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        timePassed = timePassed % AGENT_ROOT_OPTIMISTIC_PERIOD;
        skip(timePassed);
        // Form a second attestation: Notary 1
        RawSnapshot memory rs = Random(secondRA.snapRoot).nextSnapshot();
        secondRA._snapGasHash = rs.snapGasHash();
        secondRA.setDataHash();
        uint256[] memory snapGas = rs.snapGas();
        address notaryF = domains[DOMAIN_LOCAL].agents[0];
        address notaryS = domains[DOMAIN_LOCAL].agents[1];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notaryS, secondRA);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notaryS, attPayload, attSig);
        assertTrue(lightManager.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas));
        (uint40 snapRootTime, uint40 agentRootTime, uint32 index) =
            InterfaceDestination(localDestination()).destStatus();
        assertEq(snapRootTime, block.timestamp);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(index, agentIndex[notaryF]);
    }

    function test_acceptAttestation_revert_notAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert("!agentManager");
        vm.prank(caller);
        InterfaceDestination(localDestination()).acceptAttestation(0, 0, "", 0, new ChainGas[](0));
    }

    function test_acceptAttestation_notAccepted_agentRootUpdated(
        RawAttestation memory firstRA,
        uint32 firstRootSubmittedAt
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA._agentRoot != agentRootLM);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        // Mock a call from lightManager, could as well use the empty values as they won't be checked for validity
        vm.prank(address(lightManager));
        assertFalse(InterfaceDestination(localDestination()).acceptAttestation(0, 0, "", 0, new ChainGas[](0)));
        (uint40 snapRootTime, uint40 agentRootTime, uint32 index) =
            InterfaceDestination(localDestination()).destStatus();
        assertEq(snapRootTime, firstRootSubmittedAt);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(index, agentIndex[domains[DOMAIN_LOCAL].agent]);
        // Should update the Agent Merkle Root
        assertEq(lightManager.agentRoot(), firstRA._agentRoot);
    }

    function test_passAgentRoot_optimisticPeriodNotOver(
        RawAttestation memory ra,
        uint32 rootSubmittedAt,
        uint32 timePassed
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(ra._agentRoot != agentRootLM);
        // Submit attestation that updates `nextAgentRoot`
        test_submitAttestation_updatesAgentRoot(ra, rootSubmittedAt);
        timePassed = timePassed % AGENT_ROOT_OPTIMISTIC_PERIOD;
        skip(timePassed);
        (bool rootPassed, bool rootPending) = InterfaceDestination(localDestination()).passAgentRoot();
        assertFalse(rootPassed);
        assertTrue(rootPending);
        assertEq(lightManager.agentRoot(), agentRootLM);
    }

    function test_passAgentRoot_optimisticPeriodOver(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        // Submit attestation that updates `nextAgentRoot`
        test_submitAttestation_updatesAgentRoot(ra, rootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        (bool rootPassed, bool rootPending) = InterfaceDestination(localDestination()).passAgentRoot();
        assertTrue(rootPassed);
        assertFalse(rootPending);
        assertEq(lightManager.agentRoot(), ra._agentRoot);
    }

    // ═════════════════════════════════════════════════ GAS DATA ══════════════════════════════════════════════════════

    function test_getGasData(Random memory random) public {
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: random.nextUint32()});
        firstSnap.states[1] = random.nextState({origin: DOMAIN_SYNAPSE, nonce: random.nextUint32()});
        RawAttestation memory ra = random.nextAttestation(firstSnap, random.nextUint32());
        // Use current agent root in the attestation
        ra._agentRoot = getAgentRoot();
        ra.setDataHash();
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(firstNotary, ra);
        uint256[] memory firstSnapGas = firstSnap.snapGas();
        // Submit first attestation
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
        uint256 firstSkipTime = random.nextUint32();
        skip(firstSkipTime);
        RawSnapshot memory secondSnap;
        secondSnap.states = new RawState[](1);
        secondSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: random.nextUint32()});
        vm.assume(
            GasData.unwrap(firstSnap.states[0].castToState().gasData())
                != GasData.unwrap(secondSnap.states[0].castToState().gasData())
        );
        RawAttestation memory secondRA = random.nextAttestation(secondSnap, random.nextUint32());
        address secondNotary = domains[DOMAIN_LOCAL].agents[1];
        (attPayload, attSig) = signAttestation(secondNotary, secondRA);
        uint256[] memory secondSnapGas = secondSnap.snapGas();
        // Submit second attestation
        lightManager.submitAttestation(attPayload, attSig, secondRA._agentRoot, secondSnapGas);
        uint256 secondSkipTime = random.nextUint32();
        skip(secondSkipTime);
        // Check getGasData
        GasData firstRemoteGasData = firstSnap.states[0].castToState().gasData();
        GasData firstSynapseGasData = firstSnap.states[1].castToState().gasData();
        GasData secondRemoteGasData = secondSnap.states[0].castToState().gasData();
        emit log_named_uint("Remote gasData: first", GasData.unwrap(firstRemoteGasData));
        emit log_named_uint("Remote gasData: second", GasData.unwrap(secondRemoteGasData));
        emit log_named_uint("Synapse gasData: first", GasData.unwrap(firstSynapseGasData));
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(secondRemoteGasData), "!remoteGasData");
        assertEq(dataMaturity, secondSkipTime, "!remoteDataMaturity");
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(firstSynapseGasData), "!synapseGasData");
        assertEq(dataMaturity, firstSkipTime + secondSkipTime, "!synapseDataMaturity");
    }

    function test_getGasData_localDomain(Random memory random) public {
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: random.nextUint32()});
        firstSnap.states[1] = random.nextState({origin: localDomain(), nonce: random.nextUint32()});
        RawAttestation memory ra = random.nextAttestation(firstSnap, random.nextUint32());
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(firstNotary, ra);
        uint256[] memory firstSnapGas = firstSnap.snapGas();
        // Submit first attestation
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
        uint256 firstSkipTime = random.nextUint32();
        skip(firstSkipTime);
        // Check getGasData
        GasData firstRemoteGasData = firstSnap.states[0].castToState().gasData();
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(firstRemoteGasData), "!remoteGasData");
        assertEq(dataMaturity, firstSkipTime, "!remoteDataMaturity");
        // Should not save data for local domain
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(localDomain());
        assertEq(GasData.unwrap(gasData), 0, "!localGasData");
        assertEq(dataMaturity, 0, "!localDataMaturity");
    }

    function test_getGasData_noDataForDomain(Random memory random, uint32 domain) public {
        vm.assume(domain != DOMAIN_REMOTE && domain != DOMAIN_SYNAPSE);
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: random.nextUint32()});
        firstSnap.states[1] = random.nextState({origin: DOMAIN_SYNAPSE, nonce: random.nextUint32()});
        RawAttestation memory ra = random.nextAttestation(firstSnap, random.nextUint32());
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(firstNotary, ra);
        uint256[] memory firstSnapGas = firstSnap.snapGas();
        // Submit first attestation
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
        skip(random.nextUint32());
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(domain);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function test_getGasData_notaryInDispute(Random memory random) public {
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: random.nextUint32()});
        firstSnap.states[1] = random.nextState({origin: DOMAIN_SYNAPSE, nonce: random.nextUint32()});
        RawAttestation memory ra = random.nextAttestation(firstSnap, random.nextUint32());
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(firstNotary, ra);
        uint256[] memory firstSnapGas = firstSnap.snapGas();
        // Submit first attestation
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
        skip(random.nextUint32());
        // Open dispute
        openDispute({guard: domains[0].agent, notary: firstNotary});
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    // TODO: move to AgentManager test
    /*
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
        lightManager.submitAttestationReport(arPayload, arSig, attSig);
        checkDisputeOpened(destination, guard, notary);
    }

    function test_submitStateReport(RawState memory rs, RawStateIndex memory rsi) public boundIndex(rsi) {
        check_submitStateReportWithSnapshot(destination, DOMAIN_LOCAL, rs, rsi);
    }

    function test_submitStateReportWithProof(RawState memory rs, RawAttestation memory ra, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        check_submitStateReportWithSnapshotProof(destination, DOMAIN_LOCAL, rs, ra, rsi);
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
        IAgentSecured(destination).managerSlash(domain, agent, prover);
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
        IAgentSecured(destination).managerSlash(DOMAIN_LOCAL, notary, address(0));
        checkDisputeResolved({hub: destination, honest: guard, slashed: notary});
    }

    function test_managerSlash_honestNotary(RawAttestation memory ra) public {
        address guard = domains[0].agent;
        address notary = domains[DOMAIN_LOCAL].agent;
        // Put Notary 0 and Guard 0 in dispute
        test_submitAttestationReport(ra);
        // Slash the Guard
        vm.prank(address(lightManager));
        IAgentSecured(destination).managerSlash(0, guard, address(0));
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
        lightManager.submitAttestation(attPayload, attSig);
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
        lightManager.submitAttestationReport(arPayload, arSig, attSig);
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
        lightManager.submitAttestationReport(arPayload, arSig, attSig);
    }
    */

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        override
        returns (bytes32 snapRoot, bytes32[] memory snapProof)
    {
        RawAttestation memory ra;
        (ra, snapProof) = createSnapshotProof(sm);
        RawSnapshot memory rs = fakeSnapshot(sm.rs, sm.rsi);
        uint256[] memory snapGas = rs.snapGas();
        snapRoot = ra.snapRoot;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(domains[DOMAIN_LOCAL].agent, ra);
        lightManager.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
