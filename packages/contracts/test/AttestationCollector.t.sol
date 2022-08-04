// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "./utils/SynapseTest.sol";

import { AttestationCollectorHarness } from "./harnesses/AttestationCollectorHarness.sol";

contract AttestationCollectorTest is SynapseTest {
    AttestationCollectorHarness internal collector;

    event AttestationSubmitted(address indexed updater, bytes attestation);

    event UpdaterAdded(uint32 indexed domain, address updater);

    event UpdaterRemoved(uint32 indexed domain, address updater);

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

    function test_addUpdater() public {
        vm.expectEmit(true, true, true, true);
        emit UpdaterAdded(localDomain, updater);
        collector.addUpdater(localDomain, updater);
    }

    function test_addUpdater_notOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1337));
        collector.addUpdater(localDomain, fakeUpdater);
    }

    function test_removeUpdater() public {
        test_addUpdater();
        emit UpdaterRemoved(localDomain, updater);
        collector.removeUpdater(localDomain, updater);
    }

    function test_removeUpdater_notOwner() public {
        test_addUpdater();
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(address(1337));
        collector.addUpdater(localDomain, fakeUpdater);
    }

    function test_submitAttestation() public {
        test_addUpdater();
        (bytes memory attestation, ) = signHomeAttestation(updaterPK, nonce, root);
        vm.expectEmit(true, true, true, true);
        emit AttestationSubmitted(updater, attestation);
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_notUpdater() public {
        test_addUpdater();
        (bytes memory attestation, ) = signHomeAttestation(fakeUpdaterPK, nonce, root);
        vm.expectRevert("Signer is not an updater");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_wrongDomain() public {
        test_addUpdater();
        (bytes memory attestation, ) = signRemoteAttestation(updaterPK, nonce, root);
        // Signer is not set as updater for the `remoteDomain`
        vm.expectRevert("Signer is not an updater");
        collector.submitAttestation(attestation);
    }
}
