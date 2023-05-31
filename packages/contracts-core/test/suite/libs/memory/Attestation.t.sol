// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {UnformattedAttestation} from "../../../../contracts/libs/Errors.sol";
import {ATTESTATION_LENGTH} from "../../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../../utils/SynapseLibraryTest.t.sol";
import {AttestationHarness, MemViewLib} from "../../../harnesses/libs/memory/AttestationHarness.t.sol";

import {RawAttestation} from "../../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    AttestationHarness internal libHarness;

    function setUp() public {
        libHarness = new AttestationHarness();
    }

    function test_formatAttestation(RawAttestation memory ra) public {
        bytes memory payload =
            libHarness.formatAttestation(ra.snapRoot, ra.dataHash, ra.nonce, ra.blockNumber, ra.timestamp);
        // Test formatting of state
        assertEq(
            payload,
            abi.encodePacked(ra.snapRoot, ra.dataHash, ra.nonce, ra.blockNumber, ra.timestamp),
            "!formatAttestation"
        );
        checkCastToAttestation({payload: payload, isAttestation: true});
        // Test getters
        assertEq(libHarness.snapRoot(payload), ra.snapRoot, "!snapRoot");
        assertEq(libHarness.dataHash(payload), ra.dataHash, "!dataHash");
        assertEq(libHarness.nonce(payload), ra.nonce, "!nonce");
        assertEq(libHarness.blockNumber(payload), ra.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(payload), ra.timestamp, "!timestamp");
        // Test hashing of "valid attestation"
        bytes32 attestationSalt = keccak256("ATTESTATION_VALID_SALT");
        bytes32 hashedAttestation = keccak256(abi.encodePacked(attestationSalt, keccak256(payload)));
        assertEq(libHarness.hashValid(payload), hashedAttestation, "!hashValid");
        // Test hashing of "invalid attestation"
        bytes32 attestationInvalidSalt = keccak256("ATTESTATION_INVALID_SALT");
        hashedAttestation = keccak256(abi.encodePacked(attestationInvalidSalt, keccak256(payload)));
        assertEq(libHarness.hashInvalid(payload), hashedAttestation, "!hashInvalid");
    }

    function test_dataHash(bytes32 agentRoot, bytes32 snapGasHash) public {
        bytes32 dataHash = libHarness.dataHash(agentRoot, snapGasHash);
        bytes32 expected = keccak256(abi.encodePacked(agentRoot, snapGasHash));
        assertEq(dataHash, expected, "!dataHash");
    }

    function test_isAttestation(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToAttestation({payload: payload, isAttestation: length == ATTESTATION_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToAttestation(bytes memory payload, bool isAttestation) public {
        if (isAttestation) {
            assertTrue(libHarness.isAttestation(payload), "!isAttestation: when valid");
            assertEq(libHarness.castToAttestation(payload), payload, "!castToAttestation: when valid");
        } else {
            assertFalse(libHarness.isAttestation(payload), "!isAttestation: when valid");
            vm.expectRevert(UnformattedAttestation.selector);
            libHarness.castToAttestation(payload);
        }
    }
}
