// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/SnapAttestationHarness.t.sol";

// solhint-disable func-name-mixedcase
contract SnapAttestationLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    struct RawSnapAttestation {
        bytes32 root;
        uint8 height;
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
            rs.height,
            rs.nonce,
            rs.blockNumber,
            rs.timestamp
        );
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(rs.root, rs.height, rs.nonce, rs.blockNumber, rs.timestamp),
            "!formatSnapAttestation"
        );
        checkCastToAttestation({ payload: payload, isAttestation: true });
        // Test getters
        assertEq(libHarness.root(payload), rs.root, "!root");
        assertEq(libHarness.height(payload), rs.height, "!height");
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
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_toDestinationAttestation(
        RawSnapAttestation memory rs,
        address notary,
        uint40 destTimestamp
    ) public {
        vm.warp(destTimestamp);
        bytes memory payload = libHarness.formatSnapAttestation(
            rs.root,
            rs.height,
            rs.nonce,
            rs.blockNumber,
            rs.timestamp
        );
        DestinationAttestation memory destAtt = libHarness.toDestinationAttestation(
            payload,
            notary
        );
        assertEq(destAtt.notary, notary, "!notary");
        assertEq(destAtt.height, rs.height, "!height");
        assertEq(destAtt.nonce, rs.nonce, "!nonce");
        assertEq(destAtt.destTimestamp, destTimestamp, "!destTimestamp");
    }

    function test_isEmpty(DestinationAttestation memory destAtt) public {
        if (destAtt.notary == address(0)) {
            assertTrue(libHarness.isEmpty(destAtt), "!isEmpty: when Empty");
        } else {
            assertFalse(libHarness.isEmpty(destAtt), "!isEmpty: when non-Empty");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formatSummitAttestation(RawSnapAttestation memory rs) public {
        SummitAttestation memory att = SummitAttestation(
            rs.root,
            rs.height,
            rs.blockNumber,
            rs.timestamp
        );
        bytes memory payload = libHarness.formatSummitAttestation(att, rs.nonce);
        assertEq(
            payload,
            libHarness.formatSnapAttestation(
                rs.root,
                rs.height,
                rs.nonce,
                rs.blockNumber,
                rs.timestamp
            ),
            "!formatSummitAttestation"
        );
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
