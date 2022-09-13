// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { SynapseTest } from "./utils/SynapseTest.sol";
import { AttestationHubHarness } from "./harnesses/AttestationHubHarness.sol";

import { Attestation } from "../contracts/libs/Attestation.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase
contract AttestationHubTest is SynapseTest {
    using Attestation for bytes;
    using Attestation for bytes29;
    using TypedMemView for bytes29;

    AttestationHubHarness internal attestationHub;

    uint32 internal domain = 1234;
    uint32 internal nonce = 4321;
    bytes32 internal root = keccak256("root");

    bytes internal attestation;
    bytes internal attestationData;

    event LogAttestation(address notary, bytes attestationView, bytes attestation);

    function setUp() public override {
        super.setUp();
        attestationHub = new AttestationHubHarness();
        attestationHub.addNotary(domain, notary);
    }

    function test_setUp() public {
        assertTrue(attestationHub.isNotary(domain, notary), "!notary");
        assertFalse(attestationHub.isNotary(domain, fakeNotary), "!fakeNotary");
    }

    function test_submitAttestation() public {
        _createTestAttestation(notaryPK);
        vm.expectEmit(true, true, true, true);
        emit LogAttestation(notary, attestation, attestation);
        attestationHub.submitAttestation(attestation);
    }

    function test_submitAttestation_notAttestation() public {
        _createTestAttestation(notaryPK);
        // Exclude Notary signature from the attestation
        attestation = Attestation.formatAttestation(attestationData, bytes(""));
        // Attestation data alone is not an attestation
        vm.expectRevert("Not an attestation");
        attestationHub.submitAttestation(attestation);
    }

    function test_submitAttestation_notNotary() public {
        _createTestAttestation(fakeNotaryPK);
        vm.expectRevert("Signer is not a notary");
        attestationHub.submitAttestation(attestation);
    }

    function test_submitAttestation_wrongDomain() public {
        // Change domain
        ++domain;
        // Sanity check: signer is not Notary for another domain
        assert(!attestationHub.isNotary(domain, notary));
        _createTestAttestation(notaryPK);
        // Signer is not a notary for this domain
        vm.expectRevert("Signer is not a notary");
        attestationHub.submitAttestation(attestation);
    }

    function _createTestAttestation(uint256 _notaryPK) internal {
        attestationData = Attestation.formatAttestationData(domain, nonce, root);
        bytes memory notarySig = signMessage(_notaryPK, attestationData);
        attestation = Attestation.formatAttestation(attestationData, notarySig);
    }
}
