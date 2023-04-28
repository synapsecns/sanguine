// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {ISnapshotHub} from "../../contracts/interfaces/ISnapshotHub.sol";
import {SNAPSHOT_TREE_HEIGHT} from "../../contracts/libs/Constants.sol";
import {MerkleMath} from "../../contracts/libs/MerkleMath.sol";

import {InterfaceSummit} from "../../contracts/Summit.sol";
import {Versioned} from "../../contracts/base/Version.sol";

import {AgentFlag, AgentStatus, SynapseTest} from "../utils/SynapseTest.t.sol";
import {State, RawAttestation, RawState, RawStateIndex} from "../utils/libs/SynapseStructs.t.sol";
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
        notaryAttestations[0] = RawAttestation({
            snapRoot: bytes32(0),
            agentRoot: bytes32(0),
            nonce: 0,
            blockNumber: uint40(block.number),
            timestamp: uint40(block.timestamp)
        });
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
        assertEq(ISnapshotHub(summit).getAttestation(0), notaryAttestations[0].formatAttestation(), "!getAttestation");
    }

    function test_acceptGuardSnapshot_revert_notAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert("!agentManager");
        vm.prank(caller);
        InterfaceSummit(summit).acceptGuardSnapshot(0, 0, "");
    }

    function test_acceptNotarySnapshot_revert_notAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert("!agentManager");
        vm.prank(caller);
        InterfaceSummit(summit).acceptNotarySnapshot(0, 0, 0, "");
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
            expectDisputeResolved(notary, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(bondingManager.verifyAttestation(attPayload, attSig), isValid, "!returnValue");
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
        (bytes memory arPayload, bytes memory arSig) = createSignedAttestationReport(guard, ra);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationReport(arPayload, arSig);
            // TODO: check that anyone could make the call
            expectStatusUpdated(AgentFlag.Fraudulent, 0, guard);
            expectDisputeResolved(guard, address(0), address(this));
        }
        vm.recordLogs();
        assertEq(bondingManager.verifyAttestationReport(arPayload, arSig), isValid, "!returnValue");
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
            bytes memory attPayload =
                bondingManager.submitSnapshot(guardSnapshots[i].snapshot, guardSnapshots[i].signature);
            assertEq(attPayload, "", "Guard: non-empty attestation");
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

            bytes[] memory rawStates = new bytes[](STATES);
            State[] memory states = new State[](STATES);
            for (uint256 j = 0; j < STATES; ++j) {
                // Pick a random Guard to choose their state for domain (J+1)
                // To ensure that all Notary snapshots are different pick Guard
                // with the same index as Notary for the first state
                uint256 guardIndex = j == 0 ? i : random.nextUint256() % DOMAIN_AGENTS;
                rawStates[j] = guardStates[guardIndex][j].formatState();
                states[j] = guardStates[guardIndex][j].castToState();
            }

            // Calculate root and height using AttestationProofGenerator
            acceptSnapshot(rawStates);
            ra.snapRoot = getSnapshotRoot();
            ra.agentRoot = getAgentRoot();
            // This is i-th submitted attestation so far, but attestation nonce starts from 1
            ra.nonce = i + 1;
            notaryAttestations[ra.nonce] = ra;
            bytes memory attestation = ra.formatAttestation();

            address notary = domains[DOMAIN_LOCAL].agents[i];
            (snapPayloads[i], snapSignatures[i]) = signSnapshot(notary, states);
            // Nothing should be saved before Notary submitted their first snapshot
            assertEq(ISnapshotHub(summit).getLatestNotaryAttestation(notary), "");

            vm.expectEmit(true, true, true, true);
            emit AttestationSaved(attestation);
            vm.expectEmit(true, true, true, true);
            emit SnapshotAccepted(DOMAIN_LOCAL, notary, snapPayloads[i], snapSignatures[i]);
            // Should pass the resulting attestation to Destination: acceptAttestation(status, sigIndex, attestation)
            vm.expectCall(
                destinationSynapse,
                abi.encodeWithSelector(
                    InterfaceDestination.acceptAttestation.selector, agentIndex[notary], type(uint256).max, attestation
                )
            );

            bytes memory attPayload = bondingManager.submitSnapshot(snapPayloads[i], snapSignatures[i]);
            assertEq(attPayload, attestation, "Notary: incorrect attestation");
            // Check attestation getter
            assertEq(ISnapshotHub(summit).getAttestation(ra.nonce), attestation, "!getAttestation");
            assertEq(
                ISnapshotHub(summit).getLatestNotaryAttestation(notary), attestation, "!getLatestNotaryAttestation"
            );

            // Check proofs for every State in the Notary snapshot
            for (uint256 j = 0; j < STATES; ++j) {
                bytes32[] memory snapProof = ISnapshotHub(summit).getSnapshotProof(ra.nonce, j);
                // Item to prove is State's "left sub-leaf"
                (bytes32 item,) = states[j].subLeafs();
                // Item index is twice the state index (since it's a left child)
                assertEq(
                    MerkleMath.proofRoot(2 * j, item, snapProof, SNAPSHOT_TREE_HEIGHT), ra.snapRoot, "!getSnapshotProof"
                );
            }

            // Check latest Notary States
            for (uint32 j = 0; j < STATES; ++j) {
                assertEq(ISnapshotHub(summit).getLatestAgentState(j + 1, notary), rawStates[j], "!latestState: notary");
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
            assertEq(ISnapshotHub(summit).getLatestNotaryAttestation(guard), "");
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

    // ══════════════════════════════════════════════ DISPUTE OPENING ══════════════════════════════════════════════════
    // TODO: move to AgentManager test
    /*
    function test_submitStateReport(uint256 domainId, RawState memory rs, RawStateIndex memory rsi)
        public
        boundIndex(rsi)
    {
        // Restrict to non-zero existing domains
        domainId = bound(domainId, 1, allDomains.length - 1);
        check_submitStateReportWithSnapshot(summit, allDomains[domainId], rs, rsi);
    }

    function test_submitStateReportWithProof(
        uint256 domainId,
        RawState memory rs,
        RawAttestation memory ra,
        RawStateIndex memory rsi
    ) public boundIndex(rsi) {
        // Restrict to non-zero existing domains
        domainId = bound(domainId, 1, allDomains.length - 1);
        check_submitStateReportWithSnapshotProof(summit, allDomains[domainId], rs, ra, rsi);
    }

    // ════════════════════════════════════════════ DISPUTE RESOLUTION ═════════════════════════════════════════════════

    function test_managerSlash(uint256 domainId, uint256 agentId, address prover) public {
        // no counterpart in this test
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        vm.expectEmit();
        emit DisputeResolved(address(0), domain, agent);
        vm.expectEmit();
        emit AgentSlashed(domain, agent, prover);
        vm.recordLogs();
        vm.prank(address(bondingManager));
        IAgentSecured(summit).managerSlash(domain, agent, prover);
        assertEq(vm.getRecordedLogs().length, 2);
        checkDisputeResolved({hub: summit, honest: address(0), slashed: agent});
    }

    function test_managerSlash_honestGuard(RawState memory rs) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 1));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, rs, rsi);
        (address guard, address notary) = (domains[0].agents[0], domains[domain].agents[0]);
        // Slash the Notary
        vm.prank(address(bondingManager));
        IAgentSecured(summit).managerSlash(domain, notary, address(0));
        checkDisputeResolved({hub: summit, honest: guard, slashed: notary});
    }

    function test_managerSlash_honestNotary(RawState memory rs) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 1));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, rs, rsi);
        (address guard, address notary) = (domains[0].agents[0], domains[domain].agents[0]);
        // Slash the Guard
        vm.prank(address(bondingManager));
        IAgentSecured(summit).managerSlash(0, guard, address(0));
        checkDisputeResolved({hub: summit, honest: notary, slashed: guard});
    }

    // ══════════════════════════════════════════ TESTS: WHILE IN DISPUTE ══════════════════════════════════════════════

    function test_submitStateReport_revert_notaryInDispute(RawState memory firstRS, RawState memory secondRS) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Create Notary 0 snapshot
        (address guard, address notary) = (domains[0].agents[1], domains[domain].agents[0]);
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, secondRS, rsi);
        // Create report by Guard 1
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, secondRS);
        vm.expectRevert("Notary already in dispute");
        bondingManager.submitStateReportWithSnapshot(0, srPayload, srSig, snapPayload, snapSig);
    }

    function test_submitStateReport_revert_guardInDispute(RawState memory firstRS, RawState memory secondRS) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Create Notary 1 snapshot
        (address guard, address notary) = (domains[0].agents[0], domains[domain].agents[1]);
        (bytes memory snapPayload, bytes memory snapSig) = createSignedSnapshot(notary, secondRS, rsi);
        // Create report by Guard 0
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, secondRS);
        vm.expectRevert("Guard already in dispute");
        bondingManager.submitStateReportWithSnapshot(0, srPayload, srSig, snapPayload, snapSig);
    }

    function test_submitStateReportWithProof_revert_notaryInDispute(
        RawState memory firstRS,
        RawState memory secondRS,
        RawAttestation memory ra
    ) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Create Notary 0 attestation
        (address guard, address notary) = (domains[0].agents[1], domains[domain].agents[0]);
        ra = createAttestation(secondRS, ra, rsi);
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard 1 signature for the report
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, secondRS);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        vm.expectRevert("Notary already in dispute");
        bondingManager.submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
        );
    }

    function test_submitStateReportWithProof_revert_guardInDispute(
        RawState memory firstRS,
        RawState memory secondRS,
        RawAttestation memory ra
    ) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Create Notary 1 attestation
        (address guard, address notary) = (domains[0].agents[0], domains[domain].agents[1]);
        ra = createAttestation(secondRS, ra, rsi);
        (bytes memory attPayload, bytes memory attSig) = signAttestation(notary, ra);
        // Create Guard 0 signature for the report
        (bytes memory srPayload, bytes memory srSig) = createSignedStateReport(guard, secondRS);
        // Generate Snapshot Proof
        bytes32[] memory snapProof = genSnapshotProof(rsi.stateIndex);
        vm.expectRevert("Guard already in dispute");
        bondingManager.submitStateReportWithSnapshotProof(
            rsi.stateIndex, srPayload, srSig, snapProof, attPayload, attSig
        );
    }

    function test_submitSnapshot_revert_notaryInDispute(RawState memory firstRS, RawState memory secondRS) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        uint32 domain = allDomains[domainId];
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Make sure state nonce is non-zero
        if (secondRS.nonce == 0) secondRS.nonce = 1;
        // Create Guard 1 snapshot
        (address guard, address notary) = (domains[0].agents[1], domains[domain].agents[0]);
        (bytes memory snapPayload, bytes memory guardSig) = createSignedSnapshot(guard, secondRS, rsi);
        // Guard 1 submits snapshot
        bondingManager.submitSnapshot(snapPayload, guardSig);
        // Notary 0 signs the same snapshot
        bytes memory notarySig = signSnapshot(notary, snapPayload);
        vm.expectRevert("Notary is in dispute");
        bondingManager.submitSnapshot(snapPayload, notarySig);
    }

    function test_submitSnapshot_success_guardInDispute(RawState memory firstRS, RawState memory secondRS) public {
        (uint256 domainId, RawStateIndex memory rsi) = (1, RawStateIndex(0, 0));
        // Put Notary 0 and Guard 0 in dispute; rsi.statesAmount is set to 1
        test_submitStateReport(domainId, firstRS, rsi);
        // Make sure state nonce is non-zero
        if (secondRS.nonce == 0) secondRS.nonce = 1;
        // Create Guard 1 snapshot
        address guard = domains[0].agents[0];
        (bytes memory snapPayload, bytes memory guardSig) = createSignedSnapshot(guard, secondRS, rsi);
        // Guard 0 submits snapshot - being in dispute does not interfere with future snapshots
        vm.expectEmit();
        emit SnapshotAccepted(0, guard, snapPayload, guardSig);
        bondingManager.submitSnapshot(snapPayload, guardSig);
    }
    */
    // ═════════════════════════════════════════════════ OVERRIDES ═════════════════════════════════════════════════════

    function systemContract() public view override returns (address) {
        return summit;
    }

    /// @notice Returns local domain for the tested system contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
