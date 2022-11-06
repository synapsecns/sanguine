// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/libs/ReportTools.t.sol";
import { ReportHubHarness } from "../../harnesses/hubs/ReportHubHarness.t.sol";

// solhint-disable func-name-mixedcase
contract ReportHubTest is ReportTools {
    using Attestation for bytes;
    using Attestation for bytes29;
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes29;

    ReportHubHarness internal reportHub;

    function setUp() public override {
        super.setUp();
        reportHub = new ReportHubHarness();
        reportHub.addNotary(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL));
        reportHub.addGuard(suiteGuard());
    }

    function test_setUp() public {
        assertTrue(
            reportHub.isNotary(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL)),
            "Failed to add notary"
        );
        assertFalse(reportHub.isNotary(DOMAIN_LOCAL, attacker), "Attacker is Notary");
        assertFalse(
            reportHub.isNotary(DOMAIN_REMOTE, suiteNotary(DOMAIN_LOCAL)),
            "Added Notary on another domain"
        );

        assertTrue(reportHub.isGuard(suiteGuard()), "Failed to add Guard");
        assertFalse(reportHub.isGuard(attacker), "Attacker is Guard");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: SUBMIT REPORT                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitReport_fraud() public {
        createAttestationMock(DOMAIN_LOCAL);
        createReport(Report.Flag.Fraud);
        expectLogReport();
        reportHubSubmitReport();
    }

    function test_submitReport_valid() public {
        createAttestationMock(DOMAIN_LOCAL);
        createReport(Report.Flag.Valid);
        expectLogReport();
        reportHubSubmitReport();
    }

    function test_submitReport_revert_notNotary() public {
        createAttestationMock({ domain: DOMAIN_LOCAL, signer: attacker });
        createReport(Report.Flag.Fraud);
        vm.expectRevert("Signer is not a notary");
        reportHubSubmitReport();
    }

    function test_submitReport_revert_wrongDomain() public {
        createAttestationMock({ domain: DOMAIN_REMOTE, signer: attacker });
        createReport(Report.Flag.Fraud);
        // notary is not active on REMOTE_DOMAIN
        vm.expectRevert("Signer is not a notary");
        reportHubSubmitReport();
    }

    function test_submitReport_revert_noNotarySignature() public {
        createAttestationMock(DOMAIN_LOCAL);
        // Strip notary signature from attestation payload
        attestationRaw = Attestation.formatAttestation(
            attestationRaw.castToAttestation().attestationData().clone(),
            ""
        );
        createReport(Report.Flag.Fraud);
        // Report on something that is not an Attestation is not a Report
        vm.expectRevert("Not a report");
        reportHubSubmitReport();
    }

    function test_submitReport_revert_noGuardSignature() public {
        createAttestationMock(DOMAIN_LOCAL);
        // Strip guard signature from report payload
        reportRaw = Report.formatFraudReport(attestationRaw, "");
        vm.expectRevert("Not a report");
        reportHubSubmitReport();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function reportHubSubmitReport() public {
        vm.prank(broadcaster);
        reportHub.submitReport(reportRaw);
    }
}
