// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { State, StateLib, STATE_LENGTH } from "../../../contracts/libs/State.sol";
import { SNAPSHOT_TREE_HEIGHT } from "../../../contracts/libs/Constants.sol";
import { MerkleList } from "../../../contracts/libs/MerkleList.sol";
import { SynapseLibraryTest, TypedMemView } from "../../utils/SynapseLibraryTest.t.sol";
import { SnapshotHarness } from "../../harnesses/libs/SnapshotHarness.t.sol";

// solhint-disable func-name-mixedcase
contract SnapshotLibraryTest is SynapseLibraryTest {
    using StateLib for bytes;
    using TypedMemView for bytes;

    uint256 internal constant MAX_STATES = 32;

    SnapshotHarness internal libHarness;

    function setUp() public override {
        libHarness = new SnapshotHarness();
    }

    function test_formatSnapshot(uint256 statesAmount) public {
        // Should be in [1 .. MAX_STATES] range
        statesAmount = bound(statesAmount, 1, MAX_STATES);
        bytes memory payload = new bytes(statesAmount * STATE_LENGTH);
        bytes[] memory statePayloads = new bytes[](statesAmount);
        bytes32[] memory stateHashes = new bytes32[](statesAmount);
        // Construct fake states having different byte representation
        for (uint256 i = 0; i < statesAmount; ++i) {
            statePayloads[i] = new bytes(STATE_LENGTH);
            for (uint256 j = 0; j < STATE_LENGTH; ++j) {
                bytes1 b = bytes1(uint8(i + 1));
                statePayloads[i][j] = b;
                payload[i * STATE_LENGTH + j] = b;
            }
            // State library is covered in a separate uint test, we assume it is working fine
            (bytes32 leftLeaf, bytes32 rightLeaf) = statePayloads[i].castToState().subLeafs();
            // For Snapshot Merkle Tree we use the hash of two sub-leafs as "leaf"
            stateHashes[i] = keccak256(bytes.concat(leftLeaf, rightLeaf));
        }
        bytes32 snapshotSalt = keccak256("SNAPSHOT_SALT");
        bytes32 hashedSnapshot = keccak256(abi.encodePacked(snapshotSalt, keccak256(payload)));
        // Test formatting of snapshot
        assertEq(libHarness.formatSnapshot(statePayloads), payload, "!formatSnapshot");
        checkCastToSnapshot({ payload: payload, isSnapshot: true });
        // Test getters
        assertEq(libHarness.statesAmount(payload), statesAmount, "!statesAmount");
        for (uint256 i = 0; i < statesAmount; ++i) {
            assertEq(libHarness.state(payload, i), statePayloads[i], "!state");
        }
        // Test hashing
        assertEq(libHarness.hash(payload), hashedSnapshot, "!hash");
        // Test root
        // MerkleList library is covered in a separate uint test, we assume it is working fine
        MerkleList.calculateRoot(stateHashes, SNAPSHOT_TREE_HEIGHT - 1);
        // Expected merkle root value is stateHashes[0]
        assertEq(libHarness.root(payload), stateHashes[0], "!root");
    }

    function test_isSnapshot_length(uint16 length) public {
        bytes memory payload = new bytes(length);
        checkCastToSnapshot({
            payload: payload,
            isSnapshot: length > 0 &&
                length <= MAX_STATES * STATE_LENGTH &&
                length % STATE_LENGTH == 0
        });
    }

    function test_isSnapshot_stateAmounts(uint8 stateAmounts) public {
        bytes memory payload = new bytes(stateAmounts * STATE_LENGTH);
        checkCastToSnapshot({
            payload: payload,
            isSnapshot: stateAmounts > 0 && stateAmounts <= MAX_STATES
        });
    }

    function test_formatSnapshot_tooManyStates(uint256 stateAmounts) public {
        stateAmounts = bound(stateAmounts, MAX_STATES + 1, 255);
        bytes[] memory statePayloads = new bytes[](stateAmounts);
        for (uint256 i = 0; i < stateAmounts; ++i) {
            statePayloads[i] = new bytes(STATE_LENGTH);
        }
        vm.expectRevert("Invalid states amount");
        libHarness.formatSnapshot(statePayloads);
    }

    function test_formatSnapshot_noStates() public {
        vm.expectRevert("Invalid states amount");
        libHarness.formatSnapshot(new bytes[](0));
    }

    function checkCastToSnapshot(bytes memory payload, bool isSnapshot) public {
        if (isSnapshot) {
            assertTrue(libHarness.isSnapshot(payload), "!isSnapshot: when valid");
            assertEq(libHarness.castToSnapshot(payload), payload, "!castToState: when valid");
        } else {
            assertFalse(libHarness.isSnapshot(payload), "!isSnapshot: when valid");
            vm.expectRevert("Not a snapshot");
            libHarness.castToSnapshot(payload);
        }
    }
}
