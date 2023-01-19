// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/ReportHarness.t.sol";

import "../../../contracts/libs/Auth.sol";
import "../../../contracts/libs/Attestation.sol";

// solhint-disable func-name-mixedcase
contract ReportLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    ReportHarness internal libHarness;
    uint256 internal constant GUARD_PRIV_KEY = 7331;
    uint256 internal constant NOTARY_PRIV_KEY = 1337;
    // First element is (uint8 flag)
    uint8 internal constant FIRST_ELEMENT_BYTES = 8 / 8;

    function setUp() public override {
        super.setUp();
        libHarness = new ReportHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    /* TODO (Chi): enable once Reports are reimplemented
    function test_formattedCorrectly(
        bool isFraud,
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root
    ) public {
        // Prepare attestation - this has been tested in a dedicated unit test
        bytes memory attestationData = Attestation.formatAttestationData(
            origin,
            destination,
            nonce,
            root
        );
        bytes memory notarySig = signMessage(NOTARY_PRIV_KEY, attestationData);
        bytes memory attestation = Attestation.formatAttestation(attestationData, notarySig);
        // Prepare report data
        Report.Flag flag = isFraud ? Report.Flag.Fraud : Report.Flag.Valid;
        bytes memory reportData = libHarness.formatReportData(flag, attestation);
        // Test formatter against manually constructed payload
        assertEq(reportData, abi.encodePacked(flag, attestationData), "!formatReportData");
        // Both formatters should return the same results
        assertEq(
            reportData,
            (isFraud ? libHarness.formatFraudReportData : libHarness.formatValidReportData)(
                attestation
            ),
            string.concat("!format", isFraud ? "Fraud" : "Valid", "ReportData")
        );
        // Prepare report
        bytes memory guardSig = signMessage(GUARD_PRIV_KEY, reportData);
        bytes memory report = libHarness.formatReport(flag, attestation, guardSig);
        // Test formatter against manually constructed payload
        assertEq(
            report,
            abi.encodePacked(flag, attestationData, notarySig, guardSig),
            "!formatReport"
        );
        // Both formatters should return the same results
        assertEq(
            report,
            (isFraud ? libHarness.formatFraudReport : libHarness.formatValidReport)(
                attestation,
                guardSig
            ),
            string.concat("!format", isFraud ? "Fraud" : "Valid", "Report")
        );
        // Test formatting checker
        assertTrue(libHarness.isReport(report), "!isReport");
        // Test getters
        assertEq(libHarness.reportedFraud(SynapseTypes.REPORT, report), isFraud, "!reportedFraud");
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.reportedAttestation,
            payloadType: SynapseTypes.REPORT,
            payload: report,
            expectedType: SynapseTypes.ATTESTATION,
            expectedData: attestation,
            revertMessage: "!reportedAttestation"
        });
        checkBytes29Getter({
            getter: libHarness.reportData,
            payloadType: SynapseTypes.REPORT,
            payload: report,
            expectedType: SynapseTypes.REPORT_DATA,
            expectedData: reportData,
            revertMessage: "!reportData"
        });
        checkBytes29Getter({
            getter: libHarness.guardSignature,
            payloadType: SynapseTypes.REPORT,
            payload: report,
            expectedType: SynapseTypes.SIGNATURE,
            expectedData: guardSig,
            revertMessage: "!guardSignature"
        });
    }

    function test_isReport_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Report's first element (uint8 flag)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isReport(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isReport: short payload"
        );
    }

    function test_isReport_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        assertTrue(libHarness.isReport(createTestPayload()), "!isReport: test payload");
    }

    function test_isReport_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        assertFalse(
            libHarness.isReport(cutLastByte(createTestPayload())),
            "!isReport: 1 byte shorter"
        );
    }

    function test_isReport_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        assertFalse(
            libHarness.isReport(addLastByte(createTestPayload())),
            "!isReport: 1 byte longer"
        );
    }

    function test_isReport_unknownFlag(uint8 flag) public {
        // Wrong flag value means payload is not a formatted Report
        vm.assume(flag != uint8(Report.Flag.Valid) && flag != uint8(Report.Flag.Fraud));
        bytes memory report = abi.encodePacked(flag, new bytes(Report.REPORT_LENGTH - 1));
        // Sanity check
        assert(report.length == Report.REPORT_LENGTH);
        assertFalse(libHarness.isReport(report), "!isReport: unknown flag");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    /* TODO (Chi): enable once Reports are reimplemented
    function test_wrongTypeRevert_reportedAttestation(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.REPORT });
        libHarness.reportedAttestation(wrongType, payload);
    }

    function test_wrongTypeRevert_reportData(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.REPORT });
        libHarness.reportData(wrongType, payload);
    }

    function test_wrongTypeRevert_guardSignature(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.REPORT });
        libHarness.guardSignature(wrongType, payload);
    }

    function test_wrongTypeRevert_reportedFraud(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.REPORT });
        libHarness.reportedFraud(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public pure returns (bytes memory) {
        return new bytes(ReportLib.REPORT_LENGTH);
    }
}
