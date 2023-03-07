// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/libs/Merkle.sol";
import "../../contracts/libs/Snapshot.sol";

import "../utils/SynapseTest.t.sol";
import "../utils/SynapseProofs.t.sol";
import "../utils/libs/Random.t.sol";

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
                guardStates[i][j] = random.nextState({ origin: j + 1, nonce: i + j + 1 });
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
        }
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
        }
    }
}
