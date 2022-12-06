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
        attestationHub.addAgent(DOMAIN_REMOTE, suiteNotary(DOMAIN_REMOTE));
    }

    function test_setup() public {
        assertTrue(
            attestationHub.isActiveAgent(DOMAIN_REMOTE, suiteNotary(DOMAIN_REMOTE)),
            "Failed to add notary"
        );
        assertFalse(attestationHub.isActiveAgent(DOMAIN_REMOTE, attacker), "Attacker is Notary");
        assertFalse(
            attestationHub.isActiveAgent(DOMAIN_LOCAL, suiteNotary(DOMAIN_REMOTE)),
            "Added Notary on another domain"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SUBMIT ATTESTATION                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation() public {
        createAttestationMock({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE });
        expectLogAttestation();
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_notNotary() public {
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            signer: attacker
        });
        vm.expectRevert("Signer is not authorized");
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_wrongDomain() public {
        createAttestationMock({ origin: DOMAIN_REMOTE, destination: DOMAIN_LOCAL });
        // notary is not active on REMOTE_DOMAIN
        vm.expectRevert("Signer is not authorized");
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_noNotarySignature() public {
        createAttestationMock({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE });
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
