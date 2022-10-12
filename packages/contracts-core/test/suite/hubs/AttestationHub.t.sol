// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/libs/AttestationTools.t.sol";
import { AttestationHubHarness } from "../../harnesses/hubs/AttestationHubHarness.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationHubTest is AttestationTools {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes29;

    AttestationHubHarness internal attestationHub;

    function setUp() public override {
        super.setUp();
        attestationHub = new AttestationHubHarness();
        attestationHub.addNotary(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL));
    }

    function test_setup() public {
        assertTrue(
            attestationHub.isNotary(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL)),
            "Failed to add notary"
        );
        assertFalse(attestationHub.isNotary(DOMAIN_LOCAL, attacker), "Attacker is Notary");
        assertFalse(
            attestationHub.isNotary(DOMAIN_REMOTE, suiteNotary(DOMAIN_LOCAL)),
            "Added Notary on another domain"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SUBMIT ATTESTATION                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation() public {
        createAttestationMock(DOMAIN_LOCAL);
        expectLogAttestation();
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_notNotary() public {
        createAttestationMock({ domain: DOMAIN_LOCAL, signer: attacker });
        vm.expectRevert("Signer is not a notary");
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_wrongDomain() public {
        createAttestationMock(DOMAIN_REMOTE);
        // notary is not active on REMOTE_DOMAIN
        vm.expectRevert("Signer is not a notary");
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_noNotarySignature() public {
        createAttestationMock(DOMAIN_LOCAL);
        // Strip notary signature from attestation payload
        attestationRaw = Attestation.formatAttestation(
            attestationRaw.castToAttestation().attestationData().clone(),
            ""
        );
        vm.expectRevert("Not an attestation");
        attestationHubSubmitAttestation();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationHubSubmitAttestation() public {
        vm.prank(broadcaster);
        attestationHub.submitAttestation(attestationRaw);
    }
}
