// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ATTESTATION_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, MemViewLib} from "../../utils/SynapseLibraryTest.t.sol";
import {RawAttestation, RawAttestationReport} from "../../utils/libs/SynapseStructs.t.sol";

import {AttestationFlag, AttestationReportHarness} from "../../harnesses/libs/AttestationReportHarness.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationReportLibraryTest is SynapseLibraryTest {
    using MemViewLib for bytes;

    AttestationReportHarness internal libHarness;

    function setUp() public {
        libHarness = new AttestationReportHarness();
    }

    function test_formatAttestationReport(RawAttestationReport memory rra) public {
        // Make sure flag fits into enum
        rra.flag = uint8(bound(rra.flag, 0, uint8(type(AttestationFlag).max)));
        // This is tested in AttestationLibraryTest, we assume it's working here
        bytes memory attestation = rra.attestation.formatAttestation();
        bytes memory payload = libHarness.formatAttestationReport(AttestationFlag(rra.flag), attestation);
        assertEq(payload, abi.encodePacked(rra.flag, attestation), "!formatAttestationReport");
        checkCastToAttestationReport({payload: payload, isAttestationReport: true});
        // Check getters
        assertEq(uint8(libHarness.flag(payload)), rra.flag, "!flag");
        assertEq(libHarness.attestation(payload), attestation, "!attestation");
        // Test hashing
        bytes32 attestationReportSalt = keccak256("ATTESTATION_REPORT_SALT");
        bytes32 hashedAttestationReport = keccak256(abi.encodePacked(attestationReportSalt, keccak256(payload)));
        assertEq(libHarness.hash(payload), hashedAttestationReport, "!hash");
    }

    function test_isAttestation(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToAttestationReport({payload: payload, isAttestationReport: length == 1 + ATTESTATION_LENGTH});
    }

    function checkCastToAttestationReport(bytes memory payload, bool isAttestationReport) public {
        if (isAttestationReport) {
            assertTrue(libHarness.isAttestationReport(payload), "!isAttestationReport: when valid");
            assertEq(libHarness.castToAttestationReport(payload), payload, "!castToAttestation: when valid");
        } else {
            assertFalse(libHarness.isAttestationReport(payload), "!isAttestationReport: when valid");
            vm.expectRevert("Not an attestation report");
            libHarness.castToAttestationReport(payload);
        }
    }
}
