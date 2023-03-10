// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ATTESTATION_LENGTH } from "../../../contracts/libs/Constants.sol";

import { SynapseLibraryTest, TypedMemView } from "../../utils/SynapseLibraryTest.t.sol";
import {
    AttestationHarness,
    DestinationAttestation,
    SummitAttestation,
    TypedMemView
} from "../../harnesses/libs/AttestationHarness.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    struct RawAttestation {
        bytes32 root;
        uint8 height;
        uint32 nonce;
        uint40 blockNumber;
        uint40 timestamp;
    }

    AttestationHarness internal libHarness;

    function setUp() public override {
        libHarness = new AttestationHarness();
    }

    function test_formatAttestation(RawAttestation memory ra) public {
        bytes memory payload = libHarness.formatAttestation(
            ra.root,
            ra.height,
            ra.nonce,
            ra.blockNumber,
            ra.timestamp
        );
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(ra.root, ra.height, ra.nonce, ra.blockNumber, ra.timestamp),
            "!formatAttestation"
        );
        checkCastToAttestation({ payload: payload, isAttestation: true });
        // Test getters
        assertEq(libHarness.root(payload), ra.root, "!root");
        assertEq(libHarness.height(payload), ra.height, "!height");
        assertEq(libHarness.nonce(payload), ra.nonce, "!nonce");
        assertEq(libHarness.blockNumber(payload), ra.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(payload), ra.timestamp, "!timestamp");
        // Test hashing
        bytes32 attestationSalt = keccak256("ATTESTATION_SALT");
        bytes32 hashedAttestation = keccak256(
            abi.encodePacked(attestationSalt, keccak256(payload))
        );
        assertEq(libHarness.hash(payload), hashedAttestation, "!hash");
    }

    function test_isAttestation(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToAttestation({ payload: payload, isAttestation: length == ATTESTATION_LENGTH });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_toDestinationAttestation(
        RawAttestation memory ra,
        address notary,
        uint40 destTimestamp
    ) public {
        vm.warp(destTimestamp);
        bytes memory payload = libHarness.formatAttestation(
            ra.root,
            ra.height,
            ra.nonce,
            ra.blockNumber,
            ra.timestamp
        );
        DestinationAttestation memory destAtt = libHarness.toDestinationAttestation(
            payload,
            notary
        );
        assertEq(destAtt.notary, notary, "!notary");
        assertEq(destAtt.height, ra.height, "!height");
        assertEq(destAtt.nonce, ra.nonce, "!nonce");
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

    function test_formatSummitAttestation(RawAttestation memory ra) public {
        SummitAttestation memory att = SummitAttestation(
            ra.root,
            ra.height,
            ra.blockNumber,
            ra.timestamp
        );
        bytes memory payload = libHarness.formatSummitAttestation(att, ra.nonce);
        assertEq(
            payload,
            libHarness.formatAttestation(
                ra.root,
                ra.height,
                ra.nonce,
                ra.blockNumber,
                ra.timestamp
            ),
            "!formatSummitAttestation"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPEra                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToAttestation(bytes memory payload, bool isAttestation) public {
        if (isAttestation) {
            assertTrue(libHarness.isAttestation(payload), "!isAttestation: when valid");
            assertEq(
                libHarness.castToAttestation(payload),
                payload,
                "!castToAttestation: when valid"
            );
        } else {
            assertFalse(libHarness.isAttestation(payload), "!isAttestation: when valid");
            vm.expectRevert("Not an attestation");
            libHarness.castToAttestation(payload);
        }
    }
}
