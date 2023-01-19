// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/libs/AttestationTools.t.sol";
import { AttestationHubHarness } from "../../harnesses/hubs/AttestationHubHarness.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationHubTest is AttestationTools {
    AttestationHubHarness internal attestationHub;

    function setUp() public override {
        super.setUp();
        attestationHub = new AttestationHubHarness();
        // Add suite Guards
        for (uint256 i = 0; i < GUARDS; ++i) {
            attestationHub.addAgent(0, suiteGuard(i));
        }
        // Add suite Notaries
        for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
            attestationHub.addAgent(DOMAIN_REMOTE, suiteNotary(DOMAIN_REMOTE, i));
        }
    }

    function test_setup() public {
        for (uint256 i = 0; i < GUARDS; ++i) {
            assertTrue(attestationHub.isActiveAgent(0, suiteGuard(i)), "Failed to add guard");
        }
        assertFalse(attestationHub.isActiveAgent(0, attacker), "Attacker is Guard");

        for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
            assertTrue(
                attestationHub.isActiveAgent(DOMAIN_REMOTE, suiteNotary(DOMAIN_REMOTE, i)),
                "Failed to add notary"
            );
            assertFalse(
                attestationHub.isActiveAgent(DOMAIN_LOCAL, suiteNotary(DOMAIN_REMOTE, i)),
                "Added Notary on another domain"
            );
        }
        assertFalse(attestationHub.isActiveAgent(DOMAIN_REMOTE, attacker), "Attacker is Notary");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SUBMIT ATTESTATION                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation(uint256 guardSigs, uint256 notarySigs) public {
        guardSigs = guardSigs % GUARDS;
        notarySigs = notarySigs % NOTARIES_PER_CHAIN;
        // Should be at least one signature
        vm.assume(guardSigs + notarySigs != 0);
        (address[] memory guardSigners, address[] memory notarySigners) = _createSigners({
            destination: DOMAIN_REMOTE,
            guardSigs: guardSigs,
            notarySigs: notarySigs
        });
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: guardSigners,
            notarySigners: notarySigners
        });
        expectLogAttestation();
        attestationHubSubmitAttestation();
    }

    function test_submitAttestation_revert_notAgent(
        uint256 guardSigs,
        uint256 notarySigs,
        uint256 attackerIndex
    ) public {
        guardSigs = guardSigs % GUARDS;
        notarySigs = notarySigs % NOTARIES_PER_CHAIN;
        // Should be at least one signature
        vm.assume(guardSigs + notarySigs != 0);
        // Pick one agent to substitute to attacker: this could be a Guard or a Notary
        attackerIndex = attackerIndex % (guardSigs + notarySigs);
        (address[] memory guardSigners, address[] memory notarySigners) = _createSigners({
            destination: DOMAIN_REMOTE,
            guardSigs: guardSigs,
            notarySigs: notarySigs,
            attackerIndex: attackerIndex
        });
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: guardSigners,
            notarySigners: notarySigners
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

    function test_submitAttestation_revert_noSignatures() public {
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: new address[](0),
            notarySigners: new address[](0)
        });
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
