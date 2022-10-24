// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { SynapseTest } from "../utils/SynapseTest.sol";
import { Bytes29Test } from "../utils/Bytes29Test.sol";
import { Attestation } from "../../contracts/libs/Attestation.sol";
import { Report } from "../../contracts/libs/Report.sol";
import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract ReportTest is SynapseTest, Bytes29Test {
    using Attestation for bytes29;
    using TypedMemView for bytes29;
    using Report for bytes;
    using Report for bytes29;

    uint256 internal constant NOTARY_PK = 1337;
    uint256 internal constant GUARD_PK = 7331;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");
    Report.Flag internal flag = Report.Flag.Fraud;

    bytes internal attestationData;
    bytes internal attestation;
    bytes internal guardSignature;

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();
        _verifyView(_view);
    }

    function test_formatFraudReport() public {
        flag = Report.Flag.Fraud;
        _createTestData();
        bytes memory report = Report.formatFraudReport(attestation, guardSignature);
        _verifyView(report.castToReport());
    }

    function test_formatValidReport() public {
        flag = Report.Flag.Valid;
        _createTestData();
        bytes memory report = Report.formatValidReport(attestation, guardSignature);
        _verifyView(report.castToReport());
    }

    function test_formatReportData() public {
        _createTestData();
        bytes memory reportData = Report.formatReportData(flag, attestation);
        _verifyReportData(reportData);
    }

    function test_formatFraudReportData() public {
        flag = Report.Flag.Fraud;
        _createTestData();
        bytes memory reportData = Report.formatFraudReportData(attestation);
        _verifyReportData(reportData);
    }

    function test_formatValidReportData() public {
        flag = Report.Flag.Valid;
        _createTestData();
        bytes memory reportData = Report.formatValidReportData(attestation);
        _verifyReportData(reportData);
    }

    function test_isReport_tooShort() public {
        _createTestData();
        bytes memory report = abi.encodePacked(
            uint16(attestation.length),
            uint8(flag),
            attestation
        );
        assertFalse(report.castToReport().isReport());
    }

    function test_isReport_noAttestation() public {
        _createTestData();
        bytes memory report = abi.encodePacked(uint16(0), uint8(flag), guardSignature);
        assertFalse(report.castToReport().isReport());
    }

    function test_isReport_invalidAttestation() public {
        _createTestData();
        // Use attestationData instead of full attestation (i.e. no Notary signature)
        bytes memory report = Report.formatReport(flag, attestationData, guardSignature);
        assertFalse(report.castToReport().isReport());
    }

    function test_isReport_emptyPayload() public {
        bytes memory report = bytes("");
        assertFalse(report.castToReport().isReport());
    }

    function test_incorrectType_reportedFraud() public {
        _prepareMistypedTest(SynapseTypes.REPORT).reportedFraud();
    }

    function test_incorrectType_reportedAttestation() public {
        _prepareMistypedTest(SynapseTypes.REPORT).reportedAttestation();
    }

    function test_incorrectType_reportData() public {
        _prepareMistypedTest(SynapseTypes.REPORT).reportData();
    }

    function test_incorrectType_guardSignature() public {
        _prepareMistypedTest(SynapseTypes.REPORT).guardSignature();
    }

    function _createTestData() internal {
        attestationData = Attestation.formatAttestationData(domain, nonce, root);
        bytes memory notarySig = signMessage(NOTARY_PK, attestationData);
        attestation = Attestation.formatAttestation(attestationData, notarySig);
        bytes memory reportData = Report.formatReportData(flag, attestation);
        guardSignature = signMessage(GUARD_PK, reportData);
    }

    function _createTestView() internal override returns (bytes29 _view) {
        _createTestData();
        bytes memory report = Report.formatReport(flag, attestation, guardSignature);
        return report.castToReport();
    }

    function _verifyReportData(bytes memory _reportData) internal {
        bytes memory reportData = abi.encodePacked(uint8(flag), attestationData);
        assertEq(_reportData, reportData, "!reportData");
    }

    function _verifyView(bytes29 _view) internal {
        assertTrue(_view.isReport());

        assertEq(_view.reportedFraud(), flag == Report.Flag.Fraud, "!flag");
        assertEq(_view.reportedAttestation().clone(), attestation, "!attestation");

        _verifyReportData(_view.reportData().clone());
        bytes memory guardSig = signMessage(GUARD_PK, _view.reportData().clone());
        assertEq(_view.guardSignature().clone(), guardSig, "!guardSig");
    }
}
