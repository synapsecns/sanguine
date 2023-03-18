// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { MerkleLib } from "../../../contracts/libs/Merkle.sol";
import {
    Snapshot,
    SnapshotLib,
    SNAPSHOT_MAX_STATES,
    State,
    StateLib
} from "../../../contracts/libs/Snapshot.sol";
import { fakeStates, SummitState } from "../libs/FakeIt.t.sol";
import { AttestationProofGenerator } from "./AttestationProofGenerator.t.sol";

import { Test } from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract AttestationProofGeneratorTest is Test {
    using SnapshotLib for bytes;
    using StateLib for bytes;

    AttestationProofGenerator internal proofGen;

    function setUp() public {
        proofGen = new AttestationProofGenerator();
    }

    function test_attestationProof(
        SummitState memory state,
        uint256 statesAmount,
        uint256 stateIndex
    ) public {
        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);
        (bytes[] memory states, State[] memory ptrs) = fakeStates(state, statesAmount, stateIndex);
        Snapshot snapshot = SnapshotLib.formatSnapshot(ptrs).castToSnapshot();

        proofGen.acceptSnapshot(states);
        bytes32 snapshotRoot = proofGen.root();
        assertEq(snapshotRoot, snapshot.root(), "!snapshotRoot");

        (bytes32 item, ) = state.formatSummitState().castToState().subLeafs();
        bytes32[] memory proof = proofGen.generateProof(stateIndex);
        assertEq(MerkleLib.branchRoot(item, proof, stateIndex << 1), snapshotRoot, "!proof");
    }
}
