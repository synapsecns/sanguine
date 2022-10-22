// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseTest } from "./utils/SynapseTest.sol";
import { ReportHubHarness } from "./harnesses/ReportHubHarness.sol";

import { Attestation } from "../contracts/libs/Attestation.sol";
import { Report } from "../contracts/libs/Report.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract ReportHubTest is SynapseTest {
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes29;

    ReportHubHarness internal reportHub;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");
    Report.Flag internal flag = Report.Flag.Fraud;

    bytes internal attestation;
    bytes internal attestationData;
    bytes internal report;

    event LogReport(
        address guard,
        address notary,
        bytes attestation,
        bytes reportView,
        bytes report
    );

    function setUp() public override {
        super.setUp();
        reportHub = new ReportHubHarness();
        reportHub.addNotary(domain, notary);
        reportHub.addGuard(guard);
    }

    function test_setUp() public {
        assertTrue(reportHub.isNotary(domain, notary), "!notary");
        assertTrue(reportHub.isGuard(guard), "!guard");
        assertFalse(reportHub.isNotary(domain, fakeNotary), "!fakeNotary");
        assertFalse(reportHub.isGuard(fakeGuard), "!fakeGuard");
    }

    function test_submitReport() public {
        _createTestReport(notaryPK, guardPK);
        vm.expectEmit(true, true, true, true);
        emit LogReport(guard, notary, attestation, report, report);
        reportHub.submitReport(report);
    }

    function test_submitReport_validFlag() public {
        flag = Report.Flag.Valid;
        test_submitReport();
    }

    function test_submitReport_notReport() public {
        _createTestReport(notaryPK, guardPK);
        // Exclude Guard signature from the report
        report = Report.formatReport(flag, attestation, bytes(""));
        // payload without guard signature is not a Report
        vm.expectRevert("Not a report");
        reportHub.submitReport(report);
    }

    function test_submitReport_notGuard() public {
        _createTestReport(notaryPK, fakeGuardPK);
        vm.expectRevert("Signer is not a guard");
        reportHub.submitReport(report);
    }

    function test_submitReport_notAttestation() public {
        _createTestReport(notaryPK, guardPK);
        bytes memory guardSig = report.castToReport().guardSignature().clone();
        // Exclude Notary signature from the attestation
        report = Report.formatReport(flag, attestationData, guardSig);
        // Report on something that is not an Attestation is not a Report
        vm.expectRevert("Not a report");
        reportHub.submitReport(report);
    }

    function test_submitReport_notNotary() public {
        _createTestReport(fakeNotaryPK, guardPK);
        vm.expectRevert("Signer is not a notary");
        reportHub.submitReport(report);
    }

    function _createTestReport(uint256 _notaryPK, uint256 _guardPK) internal {
        attestationData = Attestation.formatAttestationData(domain, nonce, root);
        bytes memory notarySig = signMessage(_notaryPK, attestationData);
        attestation = Attestation.formatAttestation(attestationData, notarySig);
        bytes memory reportData = Report.formatReportData(flag, attestation);
        bytes memory guardSig = signMessage(_guardPK, reportData);
        report = Report.formatReport(flag, attestation, guardSig);
    }
}
