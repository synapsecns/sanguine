// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/StateHarness.t.sol";

// solhint-disable func-name-mixedcase
contract StateLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    struct RawState {
        bytes32 root;
        uint32 origin;
        uint32 nonce;
        uint40 blockNumber;
        uint40 timestamp;
    }

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
        assertEq(libHarness.leaf(payload), keccak256(bytes.concat(leftChild, rightChild)), "!leaf");
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
