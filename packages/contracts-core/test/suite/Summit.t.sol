// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CallerNotInbox, DisputeTimeoutNotOver, NotaryInDispute} from "../../contracts/libs/Errors.sol";
import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {ISnapshotHub} from "../../contracts/interfaces/ISnapshotHub.sol";
import {DISPUTE_TIMEOUT_NOTARY, SNAPSHOT_TREE_HEIGHT} from "../../contracts/libs/Constants.sol";
import {MerkleMath} from "../../contracts/libs/merkle/MerkleMath.sol";

import {InterfaceSummit} from "../../contracts/Summit.sol";
import {Versioned} from "../../contracts/base/Version.sol";

import {AgentFlag, AgentStatus, Summit, SynapseTest} from "../utils/SynapseTest.t.sol";
import {State, RawAttestation, RawSnapshot, RawState, RawStateIndex} from "../utils/libs/SynapseStructs.t.sol";
import {Random} from "../utils/libs/Random.t.sol";
import {AgentSecuredTest} from "./hubs/ExecutionHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable code-complexity
contract SummitTest is AgentSecuredTest {
    struct SignedSnapshot {
        bytes snapshot;
        bytes signature;
    }

    uint256 internal constant STATES = 10;

    mapping(uint256 => mapping(uint256 => RawState)) internal guardStates;
    mapping(uint256 => SignedSnapshot) internal guardSnapshots;
    mapping(uint256 => RawAttestation) internal notaryAttestations;

    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    function setUp() public virtual override {
        RawAttestation memory empty = RawAttestation({
            snapRoot: 0,
            dataHash: 0,
            _agentRoot: 0,
            _snapGasHash: 0,
            nonce: 0,
            blockNumber: uint40(block.number),
            timestamp: uint40(block.timestamp)
        });
        empty.setDataHash();
        notaryAttestations[0] = empty;
        super.setUp();
    }

    function test_setupCorrectly() public {
        // Check Agents
        // Summit should know about agents from all domains, including Guards
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, IAgentSecured(summit).agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check version
        assertEq(Versioned(summit).version(), LATEST_VERSION, "!version");
        // Check attestation getter for zero nonce
        (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas) = ISnapshotHub(summit).getAttestation(0);
        assertEq(attPayload, notaryAttestations[0].formatAttestation(), "!attPayload");
        assertEq(agentRoot, 0, "!agentRoot");
        assertEq(snapGas.length, 0, "!snapGas");
    }

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new Summit({synapseDomain_: 1, agentManager_: address(2), inbox_: address(3)});
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        vm.chainId(domain);
        address agentManager = random.nextAddress();
        address inbox_ = random.nextAddress();
        address caller = random.nextAddress();
        Summit cleanContract = new Summit(domain, agentManager, inbox_);
        vm.prank(caller);
        cleanContract.initialize();
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.agentManager(), agentManager, "!agentManager");
        assertEq(cleanContract.inbox(), inbox_, "!inbox");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
    }

    function initializeLocalContract() public override {
        Summit(localContract()).initialize();
    }

    function test_acceptGuardSnapshot_revert_notInbox(address caller) public {
        vm.assume(caller != localInbox());
        vm.expectRevert(CallerNotInbox.selector);
        vm.prank(caller);
        InterfaceSummit(summit).acceptGuardSnapshot(0, 0, "");
    }

    function test_acceptNotarySnapshot_revert_notInbox(address caller) public {
        vm.assume(caller != localInbox());
        vm.expectRevert(CallerNotInbox.selector);
        vm.prank(caller);
        InterfaceSummit(summit).acceptNotarySnapshot(0, 0, 0, "");
    }

    function test_acceptNotarySnapshot_revert_blockTimestampOverflow() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        // Another Guard signs the snapshot
        (bytes memory snapPayload, bytes memory guardSignature) = signSnapshot(guard, rawSnap);
        bytes memory notarySig = signSnapshot(notary, snapPayload);
        inbox.submitSnapshot(snapPayload, guardSignature);
        vm.warp(2 ** 40);
        vm.expectRevert("SafeCast: value doesn't fit in 40 bits");
        inbox.submitSnapshot(snapPayload, notarySig);
    }

    function test_verifyAttestation_existingNonce(Random memory random, uint256 mask) public {
        test_notarySnapshots(random);
        // Restrict nonce to existing ones
        uint32 nonce = uint32(bound(random.nextUint32(), 0, DOMAIN_AGENTS));
        // Attestation is valid if and only if all four fields match
        (bool isValid, RawAttestation memory ra) = notaryAttestations[nonce].modifyAttestation(mask);
        verifyAttestation(random, ra, isValid);
    }

    function test_verifyAttestation_unknownNonce(Random memory random, RawAttestation memory ra) public {
        test_notarySnapshots(random);
        // Restrict nonce to non-existing ones
        ra.nonce = uint32(bound(ra.nonce, DOMAIN_AGENTS + 1, type(uint32).max));
        verifyAttestation(random, ra, false);
    }

    function verifyAttestation(Random memory random, RawAttestation memory ra, bool isValid) public {
        // Pick random domain expect for 0
        uint256 domainIndex = bound(random.nextUint256(), 1, allDomains.length - 1);
        uint32 domain = allDomains[domainIndex];
        // Pick random Notary
        uint256 notaryIndex = bound(random.nextUint256(), 0, DOMAIN_AGENTS - 1);
        address notary = domains[domain].agents[notaryIndex];
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestation(attPayload, attSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, domain, notary);
            expectDisputeResolved(0, notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(inbox.verifyAttestation(attPayload, attSig), isValid, "!returnValue");
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
        // Verify report on constructed attestation
        verifyAttestationReport(random, ra, isValid);
    }

    function verifyAttestationReport(Random memory random, RawAttestation memory ra, bool isAttestationValid) public {
        // Report is considered invalid, if reported attestation is valid
        bool isValid = !isAttestationValid;
        // Pick random Guard
        uint256 guardIndex = bound(random.nextUint256(), 0, DOMAIN_AGENTS - 1);
        address guard = domains[0].agents[guardIndex];
        (bytes memory attPayload, bytes memory arSig) = signAttestationReport(guard, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationReport(attPayload, arSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, 0, guard);
            expectDisputeResolved(0, guard, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(inbox.verifyAttestationReport(attPayload, arSig), isValid, "!returnValue");
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function test_guardSnapshots(Random memory random) public {
        // Every Guard submits a snapshot with a random state for domains in [1 .. DOMAINS] range
        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            State[] memory states = new State[](STATES);
            for (uint32 j = 0; j < STATES; ++j) {
                // Use random non-zero nonce for every state
                uint32 nonce = uint32(bound(random.nextUint32(), 1, type(uint32).max));
                guardStates[i][j] = random.nextState({origin: j + 1, nonce: nonce});
                states[j] = guardStates[i][j].castToState();
            }
            address guard = domains[0].agents[i];
            (bytes memory snapPayload, bytes memory snapSig) = signSnapshot(guard, states);
            guardSnapshots[i] = SignedSnapshot(snapPayload, snapSig);
        }

        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            // Check that every State is saved
            for (uint256 j = 0; j < STATES; ++j) {
                vm.expectEmit(true, true, true, true);
                emit StateSaved(guardStates[i][j].formatState());
            }
            vm.expectEmit(true, true, true, true);
            emit SnapshotAccepted(0, domains[0].agents[i], guardSnapshots[i].snapshot, guardSnapshots[i].signature);
            (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas) =
                inbox.submitSnapshot(guardSnapshots[i].snapshot, guardSnapshots[i].signature);
            assertEq(attPayload, "", "Guard: non-empty attestation");
            assertEq(agentRoot, bytes32(0), "Guard: non-empty agent root");
            assertEq(snapGas.length, 0, "Guard: non-empty snap gas data");
            // Check latest Guard States
            for (uint32 j = 0; j < STATES; ++j) {
                assertEq(
                    ISnapshotHub(summit).getLatestAgentState(j + 1, domains[0].agents[i]),
                    guardStates[i][j].formatState(),
                    "!latestState: guard"
                );
            }
        }

        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            (bytes memory snapPayload, bytes memory snapSignature) = ISnapshotHub(summit).getGuardSnapshot(i);
            assertEq(snapPayload, guardSnapshots[i].snapshot, "!snapshot");
            assertEq(snapSignature, guardSnapshots[i].signature, "!signature");
        }

        // Check global latest state
        checkLatestState();
    }

    function test_notarySnapshots(Random memory random) public {
        // Every Guard submits a snapshot with a random state for domains in [1 .. DOMAINS] range
        test_guardSnapshots(random);

        bytes[] memory snapPayloads = new bytes[](DOMAIN_AGENTS);
        bytes[] memory snapSignatures = new bytes[](DOMAIN_AGENTS);

        // Every Notary submits a snapshot with a random Guard state for all domains
        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            // Set random timestamp and block height
            RawAttestation memory ra;
            ra.blockNumber = random.nextUint40();
            ra.timestamp = random.nextUint40();
            vm.roll(ra.blockNumber);
            vm.warp(ra.timestamp);

            RawSnapshot memory rs;
            rs.states = new RawState[](STATES);
            for (uint256 j = 0; j < STATES; ++j) {
                // Pick a random Guard to choose their state for domain (J+1)
                // To ensure that all Notary snapshots are different pick Guard
                // with the same index as Notary for the first state
                uint256 guardIndex = j == 0 ? i : random.nextUint256() % DOMAIN_AGENTS;
                rs.states[j] = guardStates[guardIndex][j];
            }

            // Calculate root and height using AttestationProofGenerator
            acceptSnapshot(rs);
            ra.snapRoot = getSnapshotRoot();
            ra._agentRoot = getAgentRoot();
            ra._snapGasHash = rs.snapGasHash();
            ra.setDataHash();
            // This is i-th submitted attestation so far, but attestation nonce starts from 1
            ra.nonce = i + 1;
            notaryAttestations[ra.nonce] = ra;
            bytes memory attestation = ra.formatAttestation();

            address notary = domains[DOMAIN_LOCAL].agents[i];
            (snapPayloads[i], snapSignatures[i]) = signSnapshot(notary, rs);
            // Nothing should be saved before Notary submitted their first snapshot
            (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas) =
                ISnapshotHub(summit).getLatestNotaryAttestation(notary);
            assertEq(attPayload, "");
            assertEq(agentRoot, bytes32(0));
            assertEq(snapGas.length, 0);

            vm.expectEmit(true, true, true, true);
            emit AttestationSaved(attestation);
            vm.expectEmit(true, true, true, true);
            emit SnapshotAccepted(DOMAIN_LOCAL, notary, snapPayloads[i], snapSignatures[i]);
            // Should pass to Destination: acceptAttestation(status, sigIndex, attestation, agentRoot, snapGas)
            vm.expectCall(
                destinationSynapse,
                abi.encodeWithSelector(
                    InterfaceDestination.acceptAttestation.selector,
                    agentIndex[notary],
                    type(uint256).max,
                    attestation,
                    ra._agentRoot,
                    rs.snapGas()
                )
            );

            (attPayload, agentRoot, snapGas) = inbox.submitSnapshot(snapPayloads[i], snapSignatures[i]);
            assertEq(attPayload, attestation, "Notary: incorrect attestation");
            assertEq(agentRoot, ra._agentRoot, "Notary: incorrect agent root");
            assertEq(keccak256(abi.encodePacked(snapGas)), ra._snapGasHash, "Notary: incorrect snap gas hash");
            // Check attestation getter
            (attPayload, agentRoot, snapGas) = ISnapshotHub(summit).getAttestation(ra.nonce);
            assertEq(attPayload, attestation, "!getAttestation");
            assertEq(agentRoot, ra._agentRoot, "!getAttestation: agent root");
            assertEq(keccak256(abi.encodePacked(snapGas)), ra._snapGasHash, "!getAttestation: gas hash");
            (attPayload, agentRoot, snapGas) = ISnapshotHub(summit).getLatestNotaryAttestation(notary);
            assertEq(attPayload, attestation, "!latestAttestation");
            assertEq(agentRoot, ra._agentRoot, "!latestAttestation: agent root");
            assertEq(keccak256(abi.encodePacked(snapGas)), ra._snapGasHash, "!latestAttestation: gas hash");

            // Check proofs for every State in the Notary snapshot
            for (uint8 j = 0; j < STATES; ++j) {
                bytes32[] memory snapProof = ISnapshotHub(summit).getSnapshotProof(ra.nonce, j);
                // Item to prove is State's "left sub-leaf"
                (bytes32 item,) = rs.states[j].castToState().subLeafs();
                // Item index is twice the state index (since it's a left child)
                assertEq(
                    MerkleMath.proofRoot(2 * j, item, snapProof, SNAPSHOT_TREE_HEIGHT), ra.snapRoot, "!getSnapshotProof"
                );
            }

            // Check latest Notary States
            for (uint32 j = 0; j < STATES; ++j) {
                assertEq(
                    ISnapshotHub(summit).getLatestAgentState(j + 1, notary),
                    rs.states[j].formatState(),
                    "!latestState: notary"
                );
            }
        }

        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            (bytes memory snapPayload, bytes memory snapSignature) = ISnapshotHub(summit).getNotarySnapshot(i);
            assertEq(snapPayload, snapPayloads[i], "!payload");
            assertEq(snapSignature, snapSignatures[i], "!signature");
        }

        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            address guard = domains[0].agents[i];
            // No Attestations should be saved for Guards
            (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas) =
                ISnapshotHub(summit).getLatestNotaryAttestation(guard);
            assertEq(attPayload, "");
            assertEq(agentRoot, bytes32(0));
            assertEq(snapGas.length, 0);
        }

        // Check global latest state
        checkLatestState();
    }

    function test_getLatestState_empty(uint32 domain) public {
        assertEq(InterfaceSummit(summit).getLatestState(domain), "");
    }

    function checkLatestState() public {
        // Check global latest state
        for (uint32 j = 0; j < STATES; ++j) {
            RawState memory latestState;
            for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
                if (guardStates[i][j].nonce > latestState.nonce) {
                    latestState = guardStates[i][j];
                }
            }
            assertEq(InterfaceSummit(summit).getLatestState(j + 1), latestState.formatState(), "!getLatestState");
        }
    }

    // ══════════════════════════════════════════ TESTS: WHILE IN DISPUTE ══════════════════════════════════════════════

    function test_submitSnapshot_revert_notaryInDispute(RawState memory rs) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 1));
        uint32 domain = allDomains[domainId];
        (address guard0, address guard1, address notary) =
            (domains[0].agents[0], domains[0].agents[1], domains[domain].agents[0]);
        // Put Notary 0 and Guard 0 in dispute
        openDispute({guard: guard0, notary: notary});
        // Make sure state nonce is non-zero
        if (rs.nonce == 0) rs.nonce = 1;
        // Create Guard 1 snapshot
        (bytes memory snapPayload, bytes memory guardSig) = createSignedSnapshot(guard1, rs, rsi);
        // Guard 1 submits snapshot
        inbox.submitSnapshot(snapPayload, guardSig);
        // Notary 0 signs the same snapshot
        bytes memory notarySig = signSnapshot(notary, snapPayload);
        vm.expectRevert(NotaryInDispute.selector);
        inbox.submitSnapshot(snapPayload, notarySig);
    }

    function test_submitSnapshot_success_guardInDispute(RawState memory rs) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 1));
        uint32 domain = allDomains[domainId];
        (address guard, address notary) = (domains[0].agents[0], domains[domain].agents[0]);
        // Put Notary 0 and Guard 0 in dispute
        openDispute({guard: guard, notary: notary});
        // Make sure state nonce is non-zero
        if (rs.nonce == 0) rs.nonce = 1;
        (bytes memory snapPayload, bytes memory guardSig) = createSignedSnapshot(guard, rs, rsi);
        // Guard 0 submits snapshot - being in dispute does not interfere with future snapshots
        vm.expectEmit();
        emit SnapshotAccepted(0, guard, snapPayload, guardSig);
        inbox.submitSnapshot(snapPayload, guardSig);
    }

    // ═════════════════════════════════════════ TESTS: NOTARY WON DISPUTE ═════════════════════════════════════════════

    function prepareSubmitSnapshotDisputeTest() internal returns (bytes memory snapPayload, bytes memory notarySig) {
        address notary = domains[DOMAIN_LOCAL].agent;
        address reportGuard = domains[0].agent;
        address snapshotGuard = domains[0].agents[1];

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = random.nextSnapshot();
        // Another Guard signs the snapshot
        bytes memory guardSignature;
        (snapPayload, guardSignature) = signSnapshot(snapshotGuard, rawSnap);
        notarySig = signSnapshot(notary, snapPayload);
        inbox.submitSnapshot(snapPayload, guardSignature);
        openTestDispute({guardIndex: agentIndex[reportGuard], notaryIndex: agentIndex[notary]});
    }

    /// @dev Resolves test dispute above in favor of the Notary.
    function prepareNotaryWonDisputeTest() internal {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
    }

    function test_submitSnapshot_revert_notaryWonDisputeTimeout() public {
        (bytes memory snapPayload, bytes memory notarySig) = prepareSubmitSnapshotDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        vm.expectRevert(DisputeTimeoutNotOver.selector);
        inbox.submitSnapshot(snapPayload, notarySig);
    }

    function test_submitSnapshot_afterNotaryDisputeTimeout() public {
        address notary = domains[DOMAIN_LOCAL].agent;
        (bytes memory snapPayload, bytes memory notarySig) = prepareSubmitSnapshotDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        inbox.submitSnapshot(snapPayload, notarySig);
        (bytes memory snapPayload_,) = ISnapshotHub(summit).getNotarySnapshot(0);
        assertEq(snapPayload_, snapPayload);
    }

    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    function localContract() public view override returns (address) {
        return summit;
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
