// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/SnapAttestationHarness.t.sol";

// solhint-disable func-name-mixedcase
contract SnapAttestationLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    struct RawSnapAttestation {
        bytes32 root;
        uint8 depth;
        uint32 nonce;
        uint40 blockNumber;
        uint40 timestamp;
    }

    SnapAttestationHarness internal libHarness;

    function setUp() public override {
        libHarness = new SnapAttestationHarness();
    }

    function test_formatSnapAttestation(RawSnapAttestation memory rs) public {
        bytes memory payload = libHarness.formatSnapAttestation(
            rs.root,
            rs.depth,
            rs.nonce,
            rs.blockNumber,
            rs.timestamp
        );
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(rs.root, rs.depth, rs.nonce, rs.blockNumber, rs.timestamp),
            "!formatSnapAttestation"
        );
        checkCastToAttestation({ payload: payload, isAttestation: true });
        // Test getters
        assertEq(libHarness.root(payload), rs.root, "!root");
        assertEq(libHarness.depth(payload), rs.depth, "!depth");
        assertEq(libHarness.nonce(payload), rs.nonce, "!nonce");
        assertEq(libHarness.blockNumber(payload), rs.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(payload), rs.timestamp, "!timestamp");
    }

    function test_isAttestation(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToAttestation({
            payload: payload,
            isAttestation: length == SNAP_ATTESTATION_LENGTH
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToAttestation(bytes memory payload, bool isAttestation) public {
        if (isAttestation) {
            assertTrue(libHarness.isAttestation(payload), "!isAttestation: when valid");
            assertEq(
                libHarness.castToSnapAttestation(payload),
                payload,
                "!castToSnapAttestation: when valid"
            );
        } else {
            assertFalse(libHarness.isAttestation(payload), "!isAttestation: when valid");
            vm.expectRevert("Not a snapshot attestation");
            libHarness.castToSnapAttestation(payload);
        }
    }
}
