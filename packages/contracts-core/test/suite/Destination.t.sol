// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DISPUTE_TIMEOUT_NOTARY} from "../../contracts/libs/Constants.sol";
import {CallerNotInbox, DisputeTimeoutNotOver, NotaryInDispute, OutdatedNonce} from "../../contracts/libs/Errors.sol";
import {SNAPSHOT_MAX_STATES} from "../../contracts/libs/memory/Snapshot.sol";
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
        Destination dst = Destination(localDestination());
        // Check Agents: all Agents are known in LightManager post-setup
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, dst.agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check AgentManager and Inbox
        assertEq(dst.agentManager(), localAgentManager(), "!agentManager");
        assertEq(dst.inbox(), localInbox(), "!inbox");
        // Check version
        assertEq(dst.version(), LATEST_VERSION, "!version");
        // Check pending Agent Merkle Root
        (bool rootPending) = dst.passAgentRoot();
        assertFalse(rootPending);
    }

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new Destination({synapseDomain_: 1, agentManager_: address(2), inbox_: address(3)});
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.assume(domain != DOMAIN_SYNAPSE);
        address caller = random.nextAddress();
        vm.chainId(uint256(domain));
        LightManager agentManager = new LightManager(DOMAIN_SYNAPSE);
        address inbox_ = random.nextAddress();
        bytes32 agentRoot = random.next();
        Destination cleanContract = new Destination(DOMAIN_SYNAPSE, address(agentManager), inbox_);
        agentManager.initialize(address(0), address(cleanContract), inbox_);
        vm.prank(caller);
        cleanContract.initialize(agentRoot);
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), address(agentManager), "!agentManager");
        assertEq(cleanContract.inbox(), inbox_, "!inbox");
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
            lightInbox.submitAttestation(attPayloads[index], attSignatures[index], ra._agentRoot, snapGas);
        }
        for (uint32 index = 0; index < amount; ++index) {
            (bytes memory attPayload, bytes memory attSignature) =
                InterfaceDestination(localDestination()).getAttestation(index);
            assertEq(attPayload, attPayloads[index], "!payload");
            assertEq(attSignature, attSignatures[index], "!signature");
        }
    }

    function test_submitAttestation(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        vm.assume(ra.nonce != 0);
        RawSnapshot memory rs = Random(ra.snapRoot).nextSnapshot();
        ra._snapGasHash = rs.snapGasHash();
        ra.setDataHash();
        uint256[] memory snapGas = rs.snapGas();
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        vm.warp(rootSubmittedAt);
        vm.expectEmit();
        emit AttestationAccepted(DOMAIN_LOCAL, notary, attPayload, attSig);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        (uint48 snapRootTime,,) = InterfaceDestination(localDestination()).destStatus();
        assertEq(snapRootTime, rootSubmittedAt);
        assertEq(InterfaceDestination(localDestination()).lastAttestationNonce(agentIndex[notary]), ra.nonce);
    }

    function test_submitAttestation_updatesAgentRoot(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        vm.assume(ra.nonce != 0);
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
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
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
        vm.assume(firstRA.nonce != 0 && secondRA.nonce != 0);
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
        assertTrue(lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas));
        (uint40 snapRootTime, uint40 agentRootTime, uint32 index) =
            InterfaceDestination(localDestination()).destStatus();
        assertEq(snapRootTime, block.timestamp);
        assertEq(agentRootTime, firstRootSubmittedAt);
        assertEq(index, agentIndex[notaryF]);
    }

    function test_acceptAttestation_revert_notInbox(address caller) public {
        vm.assume(caller != localInbox());
        vm.expectRevert(CallerNotInbox.selector);
        vm.prank(caller);
        InterfaceDestination(localDestination()).acceptAttestation(0, 0, "", 0, new ChainGas[](0));
    }

    function test_acceptAttestation_revert_blockTimestampOverflow() public {
        address notary = domains[DOMAIN_LOCAL].agent;

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        uint256[] memory snapGas = rawSnap.snapGas();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);

        vm.warp(2 ** 40);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
    }

    function test_acceptAttestation_revert_notaryInDispute(uint256 notaryId) public {
        address notary = domains[DOMAIN_LOCAL].agents[notaryId % DOMAIN_AGENTS];
        openDispute({guard: domains[0].agent, notary: notary});
        vm.prank(address(lightInbox));
        vm.expectRevert(NotaryInDispute.selector);
        InterfaceDestination(localDestination()).acceptAttestation(agentIndex[notary], 0, "", 0, new ChainGas[](0));
    }

    function test_acceptAttestation_revert_notaryWonDisputeTimeout() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        uint256[] memory snapGas = rawSnap.snapGas();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);

        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
        skip(7 days);
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        vm.expectRevert(DisputeTimeoutNotOver.selector);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
    }

    function test_acceptAttestation_afterNotaryDisputeTimeout() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        bytes32 snapRoot = rawSnap.castToSnapshot().calculateRoot();
        // Sanity check
        assert(testedEH().getAttestationNonce(snapRoot) == 0);
        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        uint256[] memory snapGas = rawSnap.snapGas();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);

        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
        skip(7 days);
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
        skip(DISPUTE_TIMEOUT_NOTARY);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        // Snapshot root should be known to ExecutionHub (Destination)
        assertGt(testedEH().getAttestationNonce(snapRoot), 0);
    }

    function test_acceptAttestation_revert_lowerNonce() public {
        Random memory random = Random("salt");
        RawAttestation memory firstRA = random.nextAttestation({nonce: 2});
        test_submitAttestation({ra: firstRA, rootSubmittedAt: 1000});
        skip(100);
        RawSnapshot memory rawSnap = random.nextSnapshot();
        uint256[] memory snapGas = rawSnap.snapGas();
        RawAttestation memory secondRA = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, secondRA);
        vm.expectRevert(OutdatedNonce.selector);
        lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas);
    }

    function test_acceptAttestation_revert_sameNonce() public {
        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        uint256[] memory snapGas = rawSnap.snapGas();
        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        skip(100);
        vm.expectRevert(OutdatedNonce.selector);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
    }

    function test_acceptAttestation_passesAgentRoot(
        RawAttestation memory firstRA,
        RawAttestation memory secondRA,
        uint32 firstRootSubmittedAt
    ) public {
        bytes32 agentRootLM = lightManager.agentRoot();
        vm.assume(firstRA._agentRoot != agentRootLM);
        vm.assume(firstRA.snapRoot != secondRA.snapRoot);
        vm.assume(secondRA.nonce > firstRA.nonce);
        test_submitAttestation(firstRA, firstRootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        // Form a second attestation: Notary 1
        RawSnapshot memory rs = Random(secondRA.snapRoot).nextSnapshot();
        secondRA._snapGasHash = rs.snapGasHash();
        secondRA.setDataHash();
        uint256[] memory snapGas = rs.snapGas();
        address notaryS = domains[DOMAIN_LOCAL].agents[1];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notaryS, secondRA);
        uint256 newRootTimestamp = block.timestamp;
        assertTrue(lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas));
        (uint40 snapRootTime, uint40 agentRootTime, uint32 index) =
            InterfaceDestination(localDestination()).destStatus();
        // Dest status should point to the new root
        assertEq(snapRootTime, newRootTimestamp);
        assertEq(agentRootTime, newRootTimestamp);
        assertEq(index, agentIndex[notaryS]);
        // New agent root should be pending in Destination
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), secondRA._agentRoot);
        // Should update the Agent Merkle Root
        assertEq(lightManager.agentRoot(), firstRA._agentRoot);
    }

    function test_acceptAttestation_success_diffNotary_lowerNonce() public {
        Random memory random = Random("salt");
        RawAttestation memory firstRA = random.nextAttestation({nonce: 2});
        test_submitAttestation({ra: firstRA, rootSubmittedAt: 1000});
        skip(100);
        RawSnapshot memory rawSnap = random.nextSnapshot();
        uint256[] memory snapGas = rawSnap.snapGas();
        RawAttestation memory secondRA = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        address notary = domains[DOMAIN_LOCAL].agents[1];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, secondRA);
        bool success = lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas);
        assertTrue(success);
    }

    function test_acceptAttestation_success_diffNotary_diffAtt_sameNonce() public {
        Random memory random = Random("salt");
        RawAttestation memory firstRA = random.nextAttestation({nonce: 2});
        test_submitAttestation({ra: firstRA, rootSubmittedAt: 1000});
        skip(100);
        RawSnapshot memory rawSnap = random.nextSnapshot();
        uint256[] memory snapGas = rawSnap.snapGas();
        RawAttestation memory secondRA = random.nextAttestation({rawSnap: rawSnap, nonce: 2});
        address notary = domains[DOMAIN_LOCAL].agents[1];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, secondRA);
        bool success = lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, snapGas);
        assertTrue(success);
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
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        assertTrue(rootPending);
        assertEq(lightManager.agentRoot(), agentRootLM);
    }

    function test_passAgentRoot_optimisticPeriodOver(RawAttestation memory ra, uint32 rootSubmittedAt) public {
        // Submit attestation that updates `nextAgentRoot`
        test_submitAttestation_updatesAgentRoot(ra, rootSubmittedAt);
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        assertFalse(rootPending);
        assertEq(lightManager.agentRoot(), ra._agentRoot);
    }

    /// @dev Submits a mock attestation and opens a dispute for the Notary that signed it.
    function prepareAgentRootDisputeTest() internal returns (bytes32 newAgentRoot) {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        newAgentRoot = ra._agentRoot;
        uint256[] memory snapGas = rawSnap.snapGas();

        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        // Sanity check
        assert(InterfaceDestination(localDestination()).nextAgentRoot() == newAgentRoot);
        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
    }

    /// @dev Resolves test dispute above in favor of the Notary.
    function prepareNotaryWonDisputeTest() internal {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
    }

    function test_passAgentRoot_notaryInDispute_optimisticPeriodNotOver() public {
        bytes32 oldRoot = lightManager.agentRoot();
        prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD - 1);
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should not pass the root
        assertEq(lightManager.agentRoot(), oldRoot);
        // Should clear pending
        assertFalse(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), oldRoot);
    }

    function test_passAgentRoot_notaryInDispute_optimisticPeriodOver() public {
        bytes32 oldRoot = lightManager.agentRoot();
        prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD);
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should not pass the root
        assertEq(lightManager.agentRoot(), oldRoot);
        // Should clear pending
        assertFalse(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), oldRoot);
    }

    function test_passAgentRoot_notaryWonDisputeTimeout_optimisticPeriodNotOver() public {
        bytes32 oldRoot = lightManager.agentRoot();
        bytes32 newRoot = prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD - DISPUTE_TIMEOUT_NOTARY);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        // Time since attestation was submitted: AGENT_ROOT_OPTIMISTIC_PERIOD - 1
        // Time sinceNotary won the dispute: DISPUTE_TIMEOUT_NOTARY - 1
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should not pass the root
        assertEq(lightManager.agentRoot(), oldRoot);
        // Should not clear pending
        assertTrue(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), newRoot);
    }

    function test_passAgentRoot_notaryWonDisputeTimeout_optimisticPeriodOver() public {
        bytes32 oldRoot = lightManager.agentRoot();
        bytes32 newRoot = prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD - (DISPUTE_TIMEOUT_NOTARY - 1));
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        // Time since attestation was submitted: AGENT_ROOT_OPTIMISTIC_PERIOD
        // Time sinceNotary won the dispute: DISPUTE_TIMEOUT_NOTARY - 1
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should not pass the root
        assertEq(lightManager.agentRoot(), oldRoot);
        // Should not clear pending
        assertTrue(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), newRoot);
    }

    function test_passAgentRoot_afterNotaryDisputeTimeout_optimisticPeriodNotOver() public {
        bytes32 oldRoot = lightManager.agentRoot();
        bytes32 newRoot = prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD - 1 - DISPUTE_TIMEOUT_NOTARY);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        // Time since attestation was submitted: AGENT_ROOT_OPTIMISTIC_PERIOD - 1
        // Time sinceNotary won the dispute: DISPUTE_TIMEOUT_NOTARY
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should not pass the root
        assertEq(lightManager.agentRoot(), oldRoot);
        // Should not clear pending
        assertTrue(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), newRoot);
    }

    function test_passAgentRoot_afterNotaryDisputeTimeout_optimisticPeriodOver() public {
        bytes32 newRoot = prepareAgentRootDisputeTest();
        skip(AGENT_ROOT_OPTIMISTIC_PERIOD - DISPUTE_TIMEOUT_NOTARY);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        // Time since attestation was submitted: AGENT_ROOT_OPTIMISTIC_PERIOD
        // Time sinceNotary won the dispute: DISPUTE_TIMEOUT_NOTARY
        bool rootPending = InterfaceDestination(localDestination()).passAgentRoot();
        // Should pass the root
        assertEq(lightManager.agentRoot(), newRoot);
        // Should clear pending
        assertFalse(rootPending);
        assertEq(InterfaceDestination(localDestination()).nextAgentRoot(), newRoot);
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
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
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
        lightInbox.submitAttestation(attPayload, attSig, secondRA._agentRoot, secondSnapGas);
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
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
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
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, firstSnapGas);
        skip(random.nextUint32());
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(domain);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function prepareGasDataDisputeTest() internal returns (GasData remoteGasData, GasData synapseGasData) {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = RawSnapshot(new RawState[](2));
        rawSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 1});
        rawSnap.states[1] = random.nextState({origin: DOMAIN_SYNAPSE, nonce: 2});
        remoteGasData = rawSnap.states[0].castToState().gasData();
        synapseGasData = rawSnap.states[1].castToState().gasData();

        RawAttestation memory ra = random.nextAttestation({rawSnap: rawSnap, nonce: 1});
        uint256[] memory snapGas = rawSnap.snapGas();
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
        // Sanity checks
        {
            (GasData gasData,) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
            assert(GasData.unwrap(gasData) == GasData.unwrap(remoteGasData));
        }
        {
            (GasData gasData,) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
            assert(GasData.unwrap(gasData) == GasData.unwrap(synapseGasData));
        }
        openTestDispute({guardIndex: agentIndex[guard], notaryIndex: agentIndex[notary]});
    }

    function test_getGasData_notaryInDispute() public {
        prepareGasDataDisputeTest();
        skip(7 days);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function test_getGasData_notaryWonDisputeTimeout() public {
        prepareGasDataDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function test_getGasData_afterNotaryDisputeTimeout() public {
        (GasData remoteGasData, GasData synapseGasData) = prepareGasDataDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(remoteGasData));
        assertEq(dataMaturity, 7 days + DISPUTE_TIMEOUT_NOTARY);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_SYNAPSE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(synapseGasData));
        assertEq(dataMaturity, 7 days + DISPUTE_TIMEOUT_NOTARY);
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
        RawSnapshot memory rs = fakeSnapshot(sm.rs, sm.rsi);
        uint256[] memory snapGas = rs.snapGas();
        snapRoot = ra.snapRoot;
        (bytes memory attPayload, bytes memory attSig) = signAttestation(domains[DOMAIN_LOCAL].agent, ra);
        lightInbox.submitAttestation(attPayload, attSig, ra._agentRoot, snapGas);
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
