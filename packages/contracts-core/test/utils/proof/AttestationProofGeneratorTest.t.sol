// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MerkleLib} from "../../../contracts/libs/Merkle.sol";
import {Snapshot, SNAPSHOT_MAX_STATES, SNAPSHOT_TREE_HEIGHT, State} from "../../../contracts/libs/Snapshot.sol";
import {fakeSnapshot, RawState, RawSnapshot} from "../libs/FakeIt.t.sol";
import {AttestationProofGenerator} from "./AttestationProofGenerator.t.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract AttestationProofGeneratorTest is Test {
    AttestationProofGenerator internal proofGen;

    function setUp() public {
        proofGen = new AttestationProofGenerator();
    }

    function test_attestationProof(RawState memory rawState, uint256 statesAmount, uint256 stateIndex) public {
        statesAmount = bound(statesAmount, 1, SNAPSHOT_MAX_STATES);
        stateIndex = bound(stateIndex, 0, statesAmount - 1);
        RawSnapshot memory rawSnap = fakeSnapshot(rawState, statesAmount, stateIndex);
        Snapshot snapshot = rawSnap.castToSnapshot();
        State state = rawState.castToState();
        bytes[] memory states = rawSnap.formatStates();
        proofGen.acceptSnapshot(states);
        bytes32 snapshotRoot = proofGen.root();
        assertEq(snapshotRoot, snapshot.root(), "!snapshotRoot");

        (bytes32 item,) = state.subLeafs();
        uint256 itemIndex = stateIndex << 1;
        bytes32[] memory proof = proofGen.generateProof(stateIndex);
        assertEq(MerkleLib.proofRoot(itemIndex, item, proof, SNAPSHOT_TREE_HEIGHT), snapshotRoot, "!proof");
    }
}
