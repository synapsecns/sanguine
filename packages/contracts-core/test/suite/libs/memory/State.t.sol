// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedState} from "../../../../contracts/libs/Errors.sol";
import {STATE_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {GasData, StateHarness} from "../../../harnesses/libs/memory/StateHarness.t.sol";

import {RawState} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract StateLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    StateHarness internal libHarness;

    function setUp() public {
        libHarness = new StateHarness();
    }

    function test_formatState(RawState memory rs) public {
        GasData gasData = rs.gasData.castToGasData();
        bytes memory payload =
            libHarness.formatState(rs.root, rs.origin, rs.nonce, rs.blockNumber, rs.timestamp, gasData);
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(rs.root, rs.origin, rs.nonce, rs.blockNumber, rs.timestamp, rs.gasData.encodeGasData()),
            "!formatState"
        );
        checkCastToState({payload: payload, isState: true});
        // Test getters
        assertEq(libHarness.root(payload), rs.root, "!root");
        assertEq(libHarness.origin(payload), rs.origin, "!origin");
        assertEq(libHarness.nonce(payload), rs.nonce, "!nonce");
        assertEq(libHarness.blockNumber(payload), rs.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(payload), rs.timestamp, "!timestamp");
        assertEq(GasData.unwrap(libHarness.gasData(payload)), rs.gasData.encodeGasData(), "!gasData");
        // Test creating a leaf
        bytes32 leftChild = keccak256(abi.encodePacked(rs.root, rs.origin));
        bytes32 rightChild =
            keccak256(abi.encodePacked(rs.nonce, rs.blockNumber, rs.timestamp, rs.gasData.encodeGasData()));
        (bytes32 leftLeaf, bytes32 rightLeaf) = libHarness.subLeafs(payload);
        assertEq(libHarness.leftLeaf(rs.root, rs.origin), leftChild, "!leftLeaf");
        assertEq(leftLeaf, leftChild, "!subLeafs: left");
        assertEq(libHarness.rightLeaf(rs.nonce, rs.blockNumber, rs.timestamp, gasData), rightChild, "!rightLeaf");
        assertEq(rightLeaf, rightChild, "!subLeafs: right");
        assertEq(libHarness.leaf(payload), keccak256(abi.encodePacked(leftChild, rightChild)), "!leaf");
        // Test hashing of "invalid state"
        bytes32 stateInvalidSalt = keccak256("STATE_INVALID_SALT");
        bytes32 hashedState = keccak256(abi.encodePacked(stateInvalidSalt, keccak256(payload)));
        assertEq(libHarness.hashInvalid(payload), hashedState, "!hashInvalid");
    }

    function test_equals_identical(RawState memory a) public {
        bytes memory aa = a.formatState();
        assertTrue(libHarness.equals(aa, aa), "!equals: identical");
    }

    function test_equals_different(RawState memory a, uint256 mask) public {
        RawState memory b = a.modifyState(mask);
        bytes memory aa = a.formatState();
        bytes memory bb = b.formatState();
        assertFalse(libHarness.equals(aa, bb), "!equals: different");
    }

    function test_isState(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToState({payload: payload, isState: length == STATE_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToState(bytes memory payload, bool isState) public {
        if (isState) {
            assertTrue(libHarness.isState(payload), "!isState: when valid");
            assertEq(libHarness.castToState(payload), payload, "!castToState: when valid");
        } else {
            assertFalse(libHarness.isState(payload), "!isState: when valid");
            vm.expectRevert(UnformattedState.selector);
            libHarness.castToState(payload);
        }
    }
}
