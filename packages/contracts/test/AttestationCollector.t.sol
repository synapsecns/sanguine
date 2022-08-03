// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "./utils/SynapseTest.sol";

import { AttestationCollectorHarness } from "./harnesses/AttestationCollectorHarness.sol";

contract AttestationCollectorTest is SynapseTest {
    AttestationCollectorHarness internal collector;

    event AttestationSubmitted(address indexed notary, bytes attestation);

    event NotaryAdded(uint32 indexed domain, address notary);

    event NotaryRemoved(uint32 indexed domain, address notary);

    uint32 internal nonce = 420;
    bytes32 internal root = "root";

    function setUp() public override {
        super.setUp();
        collector = new AttestationCollectorHarness();
        collector.initialize();
    }

    function test_cannotInitializeTwice() public {
        vm.expectRevert("Initializable: contract is already initialized");
        collector.initialize();
    }

    function test_addNotary() public {
        vm.expectEmit(true, true, true, true);
        emit NotaryAdded(localDomain, notary);
        collector.addNotary(localDomain, notary);
    }

    function test_addNotary_notOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1337));
        collector.addNotary(localDomain, fakeNotary);
    }

    function test_removeNotary() public {
        test_addNotary();
        emit NotaryRemoved(localDomain, notary);
        collector.removeNotary(localDomain, notary);
    }

    function test_removeNotary_notOwner() public {
        test_addNotary();
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1337));
        collector.addNotary(localDomain, fakeNotary);
    }

    function test_submitAttestation() public {
        test_addNotary();
        (bytes memory attestation, ) = signHomeAttestation(notaryPK, nonce, root);
        vm.expectEmit(true, true, true, true);
        emit AttestationSubmitted(notary, attestation);
        collector.submitAttestation(notary, attestation);
    }

    function test_submitAttestation_invalidSignature() public {
        test_addNotary();
        (bytes memory attestation, ) = signHomeAttestation(fakeNotaryPK, nonce, root);
        vm.expectRevert("Invalid signature");
        collector.submitAttestation(notary, attestation);
    }

    function test_submitAttestation_notNotary() public {
        test_addNotary();
        (bytes memory attestation, ) = signHomeAttestation(fakeNotaryPK, nonce, root);
        vm.expectRevert("Signer is not an notary");
        collector.submitAttestation(fakeNotary, attestation);
    }

    function test_submitAttestation_wrongDomain() public {
        test_addNotary();
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, root);
        // Signer is not set as notary for the `remoteDomain`
        vm.expectRevert("Signer is not an notary");
        collector.submitAttestation(notary, attestation);
    }
}
