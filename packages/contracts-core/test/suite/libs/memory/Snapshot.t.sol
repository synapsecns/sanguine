// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IncorrectStatesAmount, UnformattedSnapshot} from "../../../../contracts/libs/Errors.sol";
import {State, StateLib, STATE_LENGTH} from "../../../../contracts/libs/memory/State.sol";
import {SNAPSHOT_TREE_HEIGHT} from "../../../../contracts/libs/Constants.sol";
import {ChainGas, GasData} from "../../../../contracts/libs/stack/GasData.sol";
import {MerkleMath} from "../../../../contracts/libs/merkle/MerkleMath.sol";
import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {SnapshotHarness} from "../../../harnesses/libs/memory/SnapshotHarness.t.sol";

import {Random} from "../../../utils/libs/Random.t.sol";
import {RawState} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract SnapshotLibraryTest is SynapseLibraryTest {
    using StateLib for bytes;
    using MemViewLib for bytes;

    uint256 internal constant MAX_STATES = 32;

    SnapshotHarness internal libHarness;

    function setUp() public {
        libHarness = new SnapshotHarness();
    }

    function test_formatSnapshot(Random memory random, uint256 statesAmount) public {
        // Should be in [1 .. MAX_STATES] range
        statesAmount = bound(statesAmount, 1, MAX_STATES);
        RawState[] memory states = new RawState[](statesAmount);
        bytes[] memory statePayloads = new bytes[](statesAmount);
        bytes32[] memory stateHashes = new bytes32[](statesAmount);
        // Construct fake states having different byte representation
        bytes memory payload = "";
        for (uint256 i = 0; i < statesAmount; ++i) {
            states[i] = random.nextState();
            statePayloads[i] = states[i].formatState();
            payload = bytes.concat(payload, statePayloads[i]);
            // State library is covered in a separate uint test, we assume it is working fine
            (bytes32 leftLeaf, bytes32 rightLeaf) = states[i].castToState().subLeafs();
            // For Snapshot Merkle Tree we use the hash of two sub-leafs as "leaf"
            stateHashes[i] = keccak256(bytes.concat(leftLeaf, rightLeaf));
        }
        // Test formatting of snapshot
        assertEq(libHarness.formatSnapshot(statePayloads), payload, "!formatSnapshot");
        checkCastToSnapshot({payload: payload, isSnapshot: true});
        // Test getters
        assertEq(libHarness.statesAmount(payload), statesAmount, "!statesAmount");
        ChainGas[] memory snapGas = libHarness.snapGas(payload);
        for (uint8 i = 0; i < statesAmount; ++i) {
            assertEq(libHarness.state(payload, i), statePayloads[i], "!state");
            assertEq(snapGas[i].domain(), states[i].origin, "!snapGas.domain");
            assertEq(GasData.unwrap(snapGas[i].gasData()), states[i].gasData.encodeGasData(), "!snapGas.gasData");
        }
        // Test hashing of "valid snapshot"
        bytes32 snapshotSalt = keccak256("SNAPSHOT_VALID_SALT");
        bytes32 hashedSnapshot = keccak256(abi.encodePacked(snapshotSalt, keccak256(payload)));
        assertEq(libHarness.hashValid(payload), hashedSnapshot, "!hashValid");
        // Test root
        // MerkleMath library is covered in a separate uint test, we assume it is working fine
        MerkleMath.calculateRoot(stateHashes, SNAPSHOT_TREE_HEIGHT - 1);
        // Expected merkle root value is stateHashes[0]
        assertEq(libHarness.calculateRoot(payload), stateHashes[0], "!root");
    }

    function test_isSnapshot_length(uint16 length) public {
        bytes memory payload = new bytes(length);
        checkCastToSnapshot({
            payload: payload,
            isSnapshot: length > 0 && length <= MAX_STATES * STATE_LENGTH && length % STATE_LENGTH == 0
        });
    }

    function test_isSnapshot_stateAmounts(uint8 stateAmounts) public {
        bytes memory payload = new bytes(stateAmounts * STATE_LENGTH);
        checkCastToSnapshot({payload: payload, isSnapshot: stateAmounts > 0 && stateAmounts <= MAX_STATES});
    }

    function test_formatSnapshot_tooManyStates(uint256 stateAmounts) public {
        stateAmounts = bound(stateAmounts, MAX_STATES + 1, 255);
        bytes[] memory statePayloads = new bytes[](stateAmounts);
        for (uint256 i = 0; i < stateAmounts; ++i) {
            statePayloads[i] = new bytes(STATE_LENGTH);
        }
        vm.expectRevert(IncorrectStatesAmount.selector);
        libHarness.formatSnapshot(statePayloads);
    }

    function test_formatSnapshot_noStates() public {
        vm.expectRevert(IncorrectStatesAmount.selector);
        libHarness.formatSnapshot(new bytes[](0));
    }

    function checkCastToSnapshot(bytes memory payload, bool isSnapshot) public {
        if (isSnapshot) {
            assertTrue(libHarness.isSnapshot(payload), "!isSnapshot: when valid");
            assertEq(libHarness.castToSnapshot(payload), payload, "!castToState: when valid");
        } else {
            assertFalse(libHarness.isSnapshot(payload), "!isSnapshot: when valid");
            vm.expectRevert(UnformattedSnapshot.selector);
            libHarness.castToSnapshot(payload);
        }
    }
}
