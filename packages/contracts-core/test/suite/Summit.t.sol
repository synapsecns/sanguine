// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IAgentRegistry } from "../../contracts/interfaces/IAgentRegistry.sol";
import { ISnapshotHub } from "../../contracts/interfaces/ISnapshotHub.sol";
import { MerkleLib } from "../../contracts/libs/Merkle.sol";
import { SnapshotLib, SummitAttestation } from "../../contracts/libs/Snapshot.sol";
import { State, StateLib, SummitState } from "../../contracts/libs/State.sol";

import { InterfaceSummit } from "../../contracts/Summit.sol";

import { SynapseTest } from "../utils/SynapseTest.t.sol";
import { SynapseProofs } from "../utils/SynapseProofs.t.sol";
import { Random } from "../utils/libs/Random.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract SummitTest is SynapseTest, SynapseProofs {
    using StateLib for bytes;

    struct SignedSnapshot {
        bytes snapshot;
        bytes signature;
    }

    uint256 internal constant STATES = 10;

    mapping(uint256 => mapping(uint256 => SummitState)) internal guardStates;
    mapping(uint256 => SignedSnapshot) internal guardSnapshots;

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
    }

    function test_guardSnapshots(Random memory random) public {
        // Every Guard submits a snapshot with a random state for domains in [1 .. DOMAINS] range
        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            State[] memory states = new State[](STATES);
            for (uint32 j = 0; j < STATES; ++j) {
                // Use random non-zero nonce for every state
                uint32 nonce = uint32(bound(random.nextUint32(), 1, type(uint32).max));
                guardStates[i][j] = random.nextState({ origin: j + 1, nonce: nonce });
                states[j] = guardStates[i][j].formatSummitState().castToState();
            }
            bytes memory snapshot = SnapshotLib.formatSnapshot(states);
            bytes memory signature = signMessage(domains[0].agents[i], keccak256(snapshot));
            guardSnapshots[i] = SignedSnapshot(snapshot, signature);
        }

        for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
            // Check that every State is saved
            for (uint256 j = 0; j < STATES; ++j) {
                vm.expectEmit(true, true, true, true);
                emit StateSaved(guardStates[i][j].formatSummitState());
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
                    guardStates[i][j].formatSummitState(),
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
            SummitAttestation memory sa;
            sa.blockNumber = random.nextUint40();
            sa.timestamp = random.nextUint40();
            vm.roll(sa.blockNumber);
            vm.warp(sa.timestamp);

            bytes[] memory rawStates = new bytes[](STATES);
            State[] memory states = new State[](STATES);
            for (uint256 j = 0; j < STATES; ++j) {
                // Pick a random Guard to choose their state for domain (J+1)
                uint256 guardIndex = random.nextUint256() % DOMAIN_AGENTS;
                rawStates[j] = guardStates[guardIndex][j].formatSummitState();
                states[j] = rawStates[j].castToState();
            }

            // Calculate root and height using AttestationProofGenerator
            acceptSnapshot(rawStates);
            sa.root = getSnapshotRoot();
            sa.height = getSnapshotHeight();
            // This is i-th submitted attestation so far
            bytes memory attestation = sa.formatSummitAttestation({ _nonce: i });

            address notary = domains[DOMAIN_LOCAL].agents[i];
            bytes memory snapshot = SnapshotLib.formatSnapshot(states);
            bytes memory signature = signMessage(notary, keccak256(snapshot));

            vm.expectEmit(true, true, true, true);
            emit AttestationSaved(attestation);
            vm.expectEmit(true, true, true, true);
            emit SnapshotAccepted(DOMAIN_LOCAL, notary, snapshot, signature);
            InterfaceSummit(summit).submitSnapshot(snapshot, signature);

            // Check proofs for every State in the Notary snapshot
            for (uint256 j = 0; j < STATES; ++j) {
                bytes32[] memory snapProof = ISnapshotHub(summit).getSnapshotProof(i, j);
                // Item to prove is State's "left sub-leaf"
                (bytes32 item, ) = states[j].subLeafs();
                // Item index is twice the state index (since it's a left child)
                assertEq(
                    MerkleLib.branchRoot(item, snapProof, 2 * j),
                    sa.root,
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
            SummitState memory latestState;
            for (uint32 i = 0; i < DOMAIN_AGENTS; ++i) {
                if (guardStates[i][j].nonce > latestState.nonce) {
                    latestState = guardStates[i][j];
                }
            }
            assertEq(
                InterfaceSummit(summit).getLatestState(j + 1),
                latestState.formatSummitState(),
                "!getLatestState"
            );
        }
    }
}
