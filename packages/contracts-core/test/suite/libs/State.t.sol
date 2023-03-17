// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { STATE_LENGTH } from "../../../contracts/libs/Constants.sol";

import { SynapseLibraryTest, TypedMemView } from "../../utils/SynapseLibraryTest.t.sol";
import { OriginState, StateHarness, SummitState } from "../../harnesses/libs/StateHarness.t.sol";

struct RawState {
    bytes32 root;
    uint32 origin;
    uint32 nonce;
    uint40 blockNumber;
    uint40 timestamp;
}

// solhint-disable func-name-mixedcase
contract StateLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    StateHarness internal libHarness;

    function setUp() public override {
        libHarness = new StateHarness();
    }

    function test_formatState(RawState memory rs) public {
        bytes memory payload = libHarness.formatState(
            rs.root,
            rs.origin,
            rs.nonce,
            rs.blockNumber,
            rs.timestamp
        );
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(rs.root, rs.origin, rs.nonce, rs.blockNumber, rs.timestamp),
            "!formatState"
        );
        checkCastToState({ payload: payload, isState: true });
        // Test getters
        assertEq(libHarness.root(payload), rs.root, "!root");
        assertEq(libHarness.origin(payload), rs.origin, "!origin");
        assertEq(libHarness.nonce(payload), rs.nonce, "!nonce");
        assertEq(libHarness.blockNumber(payload), rs.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(payload), rs.timestamp, "!timestamp");
        // Test creating a leaf
        bytes32 leftChild = keccak256(abi.encodePacked(rs.root, rs.origin));
        bytes32 rightChild = keccak256(abi.encodePacked(rs.nonce, rs.blockNumber, rs.timestamp));
        (bytes32 leftLeaf, bytes32 rightLeaf) = libHarness.subLeafs(payload);
        assertEq(libHarness.leftLeaf(rs.root, rs.origin), leftChild, "!leftLeaf");
        assertEq(leftLeaf, leftChild, "!subLeafs: left");
        assertEq(
            libHarness.rightLeaf(rs.nonce, rs.blockNumber, rs.timestamp),
            rightChild,
            "!rightLeaf"
        );
        assertEq(rightLeaf, rightChild, "!subLeafs: right");
        assertEq(
            libHarness.leaf(payload),
            keccak256(abi.encodePacked(leftChild, rightChild)),
            "!leaf"
        );
    }

    function test_equals(RawState memory a, uint256 mask) public {
        RawState memory b;
        b.root = a.root ^ bytes32(mask & 1);
        b.origin = a.origin ^ uint32(mask & 2);
        b.nonce = a.nonce ^ uint32(mask & 4);
        b.blockNumber = a.blockNumber ^ uint40(mask & 8);
        b.timestamp = a.timestamp ^ uint40(mask & 16);
        bytes memory aa = abi.encodePacked(a.root, a.origin, a.nonce, a.blockNumber, a.timestamp);
        bytes memory bb = abi.encodePacked(b.root, b.origin, b.nonce, b.blockNumber, b.timestamp);
        assertEq(libHarness.equals(aa, bb), keccak256(aa) == keccak256(bb));
    }

    function test_originState_parity(RawState memory rs) public {
        vm.roll(rs.blockNumber);
        vm.warp(rs.timestamp);
        OriginState memory originState = libHarness.originState();
        assertEq(originState.blockNumber, rs.blockNumber, "!blockNumber");
        assertEq(originState.timestamp, rs.timestamp, "!timestamp");
        bytes memory payload = libHarness.formatOriginState(
            originState,
            rs.root,
            rs.origin,
            rs.nonce
        );
        assertEq(
            payload,
            libHarness.formatState(rs.root, rs.origin, rs.nonce, rs.blockNumber, rs.timestamp),
            "!formatState(originState)"
        );
        assertTrue(libHarness.equalToOrigin(payload, originState), "!equalToOrigin");
    }

    function test_equalToOrigin(RawState memory a, uint256 mask) public {
        // OriginState is equal if and only if both fields are unchanged
        bool isEqual = mask & 3 == 0;
        RawState memory b;
        // Set some of the OriginState fields to different values depending on the mask
        b.blockNumber = a.blockNumber ^ uint40(mask & 1);
        b.timestamp = a.timestamp ^ uint40(mask & 2);
        assertEq(
            libHarness.equalToOrigin(
                libHarness.formatState(a.root, a.origin, a.nonce, a.blockNumber, a.timestamp),
                OriginState(b.blockNumber, b.timestamp)
            ),
            isEqual
        );
    }

    function test_summitState_parity(RawState memory rs) public {
        // State -> SummitState -> State trip test
        vm.roll(rs.blockNumber);
        vm.warp(rs.timestamp);
        bytes memory payload = libHarness.formatState(
            rs.root,
            rs.origin,
            rs.nonce,
            rs.blockNumber,
            rs.timestamp
        );
        SummitState memory state = libHarness.toSummitState(payload);
        assertEq(libHarness.formatSummitState(state), payload, "!summitState");
    }

    function test_isState(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToState({ payload: payload, isState: length == STATE_LENGTH });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToState(bytes memory payload, bool isState) public {
        if (isState) {
            assertTrue(libHarness.isState(payload), "!isState: when valid");
            assertEq(libHarness.castToState(payload), payload, "!castToState: when valid");
        } else {
            assertFalse(libHarness.isState(payload), "!isState: when valid");
            vm.expectRevert("Not a state");
            libHarness.castToState(payload);
        }
    }

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatState(bytes32(0), 0, 0, 0, 0);
    }
}
