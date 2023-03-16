// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentRegistry } from "../../contracts/interfaces/IAgentRegistry.sol";
import { ISnapshotHub } from "../../contracts/interfaces/ISnapshotHub.sol";
import { MerkleLib } from "../../contracts/libs/Merkle.sol";
import { AgentInfo, SystemEntity } from "../../contracts/libs/Structures.sol";

import { InterfaceSummit } from "../../contracts/Summit.sol";
import { Versioned } from "../../contracts/Version.sol";

import { ISystemContract, SynapseTest } from "../utils/SynapseTest.t.sol";
import { SynapseProofs } from "../utils/SynapseProofs.t.sol";
import {
    AttestationFlag,
    State,
    RawAttestation,
    RawAttestationReport,
    RawState
} from "../utils/libs/SynapseStructs.t.sol";
import { Random } from "../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract SummitTest is SynapseTest, SynapseProofs {
    struct SignedSnapshot {
        bytes snapshot;
        bytes signature;
    }

    struct AttestationMask {
        bool diffRoot;
        bool diffHeight;
        bool diffBlockNumber;
        bool diffTimestamp;
    }

    uint256 internal constant STATES = 10;

    mapping(uint256 => mapping(uint256 => RawState)) internal guardStates;
    mapping(uint256 => SignedSnapshot) internal guardSnapshots;
    mapping(uint256 => RawAttestation) internal notaryAttestations;

    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    function test_setupCorrectly() public {
        // Check Agents
        // Summit should know about agents from all domains, including Guards
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                assertTrue(IAgentRegistry(summit).isActiveAgent(domain, agent), "!agent");
            }
        }
        // Check version
        assertEq(Versioned(summit).version(), LATEST_VERSION, "!version");
    }

    function test_verifyAttestation_existingNonce(Random memory random, AttestationMask memory mask)
        public
    {
        test_notarySnapshots(random);
        // Restrict nonce to existing ones
        uint32 nonce = uint32(bound(random.nextUint32(), 0, DOMAIN_AGENTS - 1));
        // Attestation is valid if and only if all four fields match
        bool isValid = !(mask.diffRoot ||
            mask.diffHeight ||
            mask.diffBlockNumber ||
            mask.diffTimestamp);
        RawAttestation memory ra = notaryAttestations[nonce];
        if (mask.diffRoot) ra.root = ra.root ^ bytes32(uint256(1));
        if (mask.diffHeight) ra.height = ra.height ^ 1;
        if (mask.diffBlockNumber) ra.blockNumber = ra.blockNumber ^ 1;
        if (mask.diffTimestamp) ra.timestamp = ra.timestamp ^ 1;
        verifyAttestation(random, ra, isValid);
    }

    function test_verifyAttestation_unknownNonce(Random memory random, RawAttestation memory ra)
        public
    {
        test_notarySnapshots(random);
        // Restrict nonce to existing ones
        ra.nonce = uint32(bound(ra.nonce, DOMAIN_AGENTS, type(uint32).max));
        verifyAttestation(random, ra, false);
    }

    function verifyAttestation(
        Random memory random,
        RawAttestation memory ra,
        bool isValid
    ) public {
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
            expectAgentSlashed(domain, notary);
        }
        vm.recordLogs();
        assertEq(
            InterfaceSummit(summit).verifyAttestation(attPayload, attSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
        // Verify report on constructed attestation
        verifyAttestationReport(random, ra, isValid);
    }

    function verifyAttestationReport(
        Random memory random,
        RawAttestation memory ra,
        bool isAttestationValid
    ) public {
        // Report is considered invalid, if reported attestation is valid
        bool isValid = !isAttestationValid;
        // Pick random Guard
        uint256 guardIndex = bound(random.nextUint256(), 0, DOMAIN_AGENTS - 1);
        address guard = domains[0].agents[guardIndex];
        RawAttestationReport memory rawAR = RawAttestationReport(
            uint8(AttestationFlag.Invalid),
            ra
        );
        (bytes memory arPayload, bytes memory arSig) = signAttestationReport(guard, rawAR);
        if (!isValid) {
            // Expect Events to be emitted
            vm.expectEmit(true, true, true, true);
            emit InvalidAttestationReport(arPayload, arSig);
            expectAgentSlashed(0, guard);
        }
        vm.recordLogs();
        assertEq(
            InterfaceSummit(summit).verifyAttestationReport(arPayload, arSig),
            isValid,
            "!returnValue"
        );
        if (isValid) {
            assertEq(vm.getRecordedLogs().length, 0, "Emitted logs when shouldn't");
        }
    }

    function expectAgentSlashed(uint32 domain, address agent) public {
        vm.expectEmit(true, true, true, true);
        emit AgentRemoved(domain, agent);
        vm.expectEmit(true, true, true, true);
        emit AgentSlashed(domain, agent);
        // Should slash Agents on Synapse Chain registries
        bytes memory expectedCall = _expectedSlashCall(domain, agent);
        vm.expectCall(originSynapse, expectedCall);
        vm.expectCall(destinationSynapse, expectedCall);
        // Should forward Slash system calls
        bytes memory data = _dataSlashAgentCall(domain, agent);
        _expectRemoteCallBondingManager(DOMAIN_LOCAL, data);
        _expectRemoteCallBondingManager(DOMAIN_REMOTE, data);
    }

    function test_guardSnapshots(Random memory random) public {
        // Every Guard submits a snapshot with a random state for domains in [1 .. DOMAINS] range
        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            State[] memory states = new State[](STATES);
            for (uint32 j = 0; j < STATES; ++j) {
                // Use random non-zero nonce for every state
                uint32 nonce = uint32(bound(random.nextUint32(), 1, type(uint32).max));
                guardStates[i][j] = random.nextState({ origin: j + 1, nonce: nonce });
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
            emit SnapshotAccepted(
                0,
                domains[0].agents[i],
                guardSnapshots[i].snapshot,
                guardSnapshots[i].signature
            );
            InterfaceSummit(summit).submitSnapshot(
                guardSnapshots[i].snapshot,
                guardSnapshots[i].signature
            );
            // Check latest Guard States
            for (uint32 j = 0; j < STATES; ++j) {
                assertEq(
                    ISnapshotHub(summit).getLatestAgentState(j + 1, domains[0].agents[i]),
                    guardStates[i][j].formatState(),
                    "!latestState: guard"
                );
            }
        }

        // Check global latest state
        checkLatestState();
    }

    function test_notarySnapshots(Random memory random) public {
        // Every Guard submits a snapshot with a random state for domains in [1 .. DOMAINS] range
        test_guardSnapshots(random);

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
                uint256 guardIndex = random.nextUint256() % DOMAIN_AGENTS;
                rawStates[j] = guardStates[guardIndex][j].formatState();
                states[j] = guardStates[guardIndex][j].castToState();
            }

            // Calculate root and height using AttestationProofGenerator
            acceptSnapshot(rawStates);
            ra.root = getSnapshotRoot();
            ra.height = getSnapshotHeight();
            // This is i-th submitted attestation so far
            ra.nonce = i;
            notaryAttestations[i] = ra;
            bytes memory attestation = ra.formatAttestation();

            address notary = domains[DOMAIN_LOCAL].agents[i];
            (bytes memory snapPayload, bytes memory snapSig) = signSnapshot(notary, states);

            vm.expectEmit(true, true, true, true);
            emit AttestationSaved(attestation);
            vm.expectEmit(true, true, true, true);
            emit SnapshotAccepted(DOMAIN_LOCAL, notary, snapPayload, snapSig);
            InterfaceSummit(summit).submitSnapshot(snapPayload, snapSig);

            // Check proofs for every State in the Notary snapshot
            for (uint256 j = 0; j < STATES; ++j) {
                bytes32[] memory snapProof = ISnapshotHub(summit).getSnapshotProof(i, j);
                // Item to prove is State's "left sub-leaf"
                (bytes32 item, ) = states[j].subLeafs();
                // Item index is twice the state index (since it's a left child)
                assertEq(
                    MerkleLib.branchRoot(item, snapProof, 2 * j),
                    ra.root,
                    "!getSnapshotProof"
                );
            }

            // Check latest Notary States
            for (uint32 j = 0; j < STATES; ++j) {
                assertEq(
                    ISnapshotHub(summit).getLatestAgentState(j + 1, notary),
                    rawStates[j],
                    "!latestState: notary"
                );
            }
        }

        // Check global latest state
        checkLatestState();
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
            assertEq(
                InterfaceSummit(summit).getLatestState(j + 1),
                latestState.formatState(),
                "!getLatestState"
            );
        }
    }

    function _expectRemoteCallBondingManager(uint32 domain, bytes memory data) internal {
        vm.expectCall(
            address(systemRouterSynapse),
            abi.encodeWithSelector(
                systemRouterSynapse.systemCall.selector,
                domain, // destination
                BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                SystemEntity.BondingManager, //recipient
                data
            )
        );
    }

    function _dataSlashAgentCall(uint32 domain, address notary)
        internal
        pure
        returns (bytes memory)
    {
        return
            abi.encodeWithSelector(
                ISystemContract.slashAgent.selector,
                0, // rootSubmittedAt
                0, // callOrigin
                0, // systemCaller
                AgentInfo(domain, notary, false)
            );
    }

    function _expectedSlashCall(uint32 domain, address notary)
        internal
        view
        returns (bytes memory)
    {
        return
            abi.encodeWithSelector(
                ISystemContract.slashAgent.selector,
                block.timestamp,
                DOMAIN_SYNAPSE,
                SystemEntity.BondingManager,
                AgentInfo(domain, notary, false)
            );
    }
}
