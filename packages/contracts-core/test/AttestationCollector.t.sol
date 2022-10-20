// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "./utils/SynapseTest.sol";

import { AttestationCollectorHarness } from "./harnesses/AttestationCollectorHarness.sol";

// solhint-disable func-name-mixedcase
contract AttestationCollectorTest is SynapseTest {
    AttestationCollectorHarness internal collector;

    uint32 internal nonce = 420;
    bytes32 internal root = "root";

    uint256[] internal notariesPK;
    address[] internal notaries;

    uint32[][] internal attestedNonces;
    bytes32[][] internal attestedRoots;
    mapping(uint32 => uint256) internal rootsAmount;

    uint256 internal constant NOTARIES_AMOUNT = 4;

    event AttestationSubmitted(address indexed notary, bytes attestation);
    event NotaryAdded(uint32 indexed domain, address indexed notary);
    event NotaryRemoved(uint32 indexed domain, address indexed notary);

    function setUp() public override {
        super.setUp();
        collector = new AttestationCollectorHarness();
        collector.initialize();

        notariesPK = new uint256[](NOTARIES_AMOUNT);
        notaries = new address[](NOTARIES_AMOUNT);
        for (uint256 i = 0; i < NOTARIES_AMOUNT; ++i) {
            notariesPK[i] = 42 + i * 69;
            notaries[i] = vm.addr(notariesPK[i]);
        }

        attestedNonces = new uint32[][](NOTARIES_AMOUNT);
        attestedRoots = new bytes32[][](NOTARIES_AMOUNT);
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

    function test_addNotaries() public {
        for (uint256 i = 0; i < NOTARIES_AMOUNT; ++i) {
            vm.expectEmit(true, true, true, true);
            emit NotaryAdded(localDomain, notaries[i]);
            collector.addNotary(localDomain, notaries[i]);
            assertTrue(collector.isNotary(localDomain, notaries[i]));
        }
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
        collector.removeNotary(localDomain, fakeNotary);
    }

    function test_submitAttestation() public {
        test_addNotary();
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        vm.expectEmit(true, true, true, true);
        emit AttestationSubmitted(notary, attestation);
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_notNotary() public {
        test_addNotary();
        (bytes memory attestation, ) = signOriginAttestation(fakeNotaryPK, nonce, root);
        vm.expectRevert("Signer is not a notary");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_wrongDomain() public {
        test_addNotary();
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, nonce, root);
        // Signer is not set as notary for the `remoteDomain`
        vm.expectRevert("Signer is not a notary");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_zeroNonce() public {
        test_addNotary();
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, 0, root);
        vm.expectRevert("Outdated attestation");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_outdated() public {
        test_submitAttestation();
        (bytes memory attestation, ) = signOriginAttestation(notaryPK, nonce, root);
        vm.expectRevert("Outdated attestation");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestation_duplicate() public {
        test_submitAttestation();
        test_addNotaries();
        (bytes memory attestation, ) = signOriginAttestation(notariesPK[0], nonce, root);
        // duplicate attestation should not be stored
        vm.expectRevert("Duplicated attestation");
        collector.submitAttestation(attestation);
    }

    function test_submitAttestations() public {
        test_addNotaries();
        // First Notary submits attestations with nonces: [1, 2, 5]
        _submitTestAttestation(1, 0, true);
        _submitTestAttestation(2, 0, true);
        _submitTestAttestation(5, 0, true);
        // Second Notary submits attestations with nonces: [1, 3, 6]
        // duplicate attestation is not stored
        _submitTestAttestation(1, 1, 0, false);
        _submitTestAttestation(3, 1, true);
        _submitTestAttestation(6, 1, true);
        // Third Notary submits nonces [1, 6, 7]
        // first two are conflicting attestations, they are stored
        _submitTestAttestation(1, 2, true);
        _submitTestAttestation(6, 2, true);
        _submitTestAttestation(7, 2, true);
        // Fourth Notary submits all duplicate attestations for nonces [1, 6]
        // duplicates are not stored
        _submitTestAttestation(1, 3, 0, false);
        _submitTestAttestation(1, 3, 2, false);
        _submitTestAttestation(6, 3, 1, false);
        _submitTestAttestation(6, 3, 2, false);

        // Submit a few fresh attestations
        _submitTestAttestation(9, 2, true);
        _submitTestAttestation(10, 0, true);
        _submitTestAttestation(12, 3, true);
        _submitTestAttestation(8, 1, true);
    }

    function test_getAttestation() public {
        test_submitAttestations();
        _checkGetAttestation(1, 0, 0);
        _checkGetAttestation(2, 0, 0);
        _checkGetAttestation(5, 0, 0);

        _checkGetAttestation(3, 1, 0);
        _checkGetAttestation(6, 1, 0);

        // conflicting attestations are stored with index > 0
        _checkGetAttestation(1, 2, 1);
        _checkGetAttestation(6, 2, 1);
        _checkGetAttestation(7, 2, 0);

        _checkGetAttestation(9, 2, 0);
        _checkGetAttestation(10, 0, 0);
        _checkGetAttestation(12, 3, 0);
        _checkGetAttestation(8, 1, 0);
    }

    function test_getAttestation_noSignature() public {
        test_submitAttestations();
        vm.expectRevert("!signature");
        // Nonce 6 was submitted only by Notaries 1 and 2
        collector.getAttestation(localDomain, 6, _generateTestRoot(6, 3));
    }

    function test_getRoot_noAttestations() public {
        test_submitAttestations();
        vm.expectRevert("!index");
        collector.getRoot(localDomain, 4, 0);
    }

    function test_getLatestAttestation_noNotaryAttestations() public {
        test_submitAttestation();
        test_addNotaries();
        vm.expectRevert("No attestations found");
        collector.getLatestAttestation(localDomain, notaries[0]);
    }

    function test_getLatestAttestation_noAttestations() public {
        test_addNotaries();
        vm.expectRevert("No attestations found");
        collector.getLatestAttestation(localDomain);
    }

    function test_getLatestAttestation_noNotaries() public {
        vm.expectRevert("!notaries");
        collector.getLatestAttestation(localDomain);
    }

    function test_latestNonce() public {
        test_submitAttestations();
        for (uint256 i = 0; i < NOTARIES_AMOUNT; ++i) {
            uint256 notaryNonces = attestedNonces[i].length;
            if (notaryNonces == 0) {
                assertEq(collector.latestNonce(localDomain, notaries[i]), 0);
            } else {
                assertEq(
                    collector.latestNonce(localDomain, notaries[i]),
                    attestedNonces[i][notaryNonces - 1]
                );
            }
        }
    }

    function test_rootsAmount() public {
        test_submitAttestations();
        for (uint32 _nonce = 0; _nonce < 16; ++_nonce) {
            assertEq(collector.rootsAmount(localDomain, _nonce), rootsAmount[_nonce]);
        }
    }

    function _submitTestAttestation(
        uint32 _nonce,
        uint256 _notaryIndex,
        bool _stored
    ) internal {
        // Make Notary sign "their own" attestation
        _submitTestAttestation(_nonce, _notaryIndex, _notaryIndex, _stored);
    }

    function _submitTestAttestation(
        uint32 _nonce,
        uint256 _notaryIndex,
        uint256 _notaryGenerationIndex,
        bool _stored
    ) internal {
        // Create attestation based by index specified for attestation generation
        bytes32 _root = _generateTestRoot(_nonce, _notaryGenerationIndex);
        bytes memory attestation = _generateTestAttestation(
            _nonce,
            _notaryIndex,
            _notaryGenerationIndex
        );
        if (_stored) {
            vm.expectEmit(true, true, true, true);
            emit AttestationSubmitted(notaries[_notaryIndex], attestation);
            // Store testing info for later checking
            attestedNonces[_notaryIndex].push(_nonce);
            attestedRoots[_notaryIndex].push(_root);
            ++rootsAmount[_nonce];
        } else {
            vm.expectRevert("Duplicated attestation");
        }
        // Use potentially another notary index for signing
        assertEq(collector.submitAttestation(attestation), _stored);
        // Check both getLatestAttestation() functions
        _checkLatestAttestations();
    }

    // This will generate unique root for every Notary, even with the same nonce
    function _generateTestRoot(uint32 _nonce, uint256 _notaryIndex)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode("root", _nonce, _notaryIndex));
    }

    // solhint-disable-next-line ordering
    function _generateTestAttestation(
        uint32 _nonce,
        uint256 _notaryIndex,
        uint256 _notaryGenerationIndex
    ) internal returns (bytes memory attestation) {
        (attestation, ) = signOriginAttestation(
            notariesPK[_notaryIndex],
            _nonce,
            _generateTestRoot(_nonce, _notaryGenerationIndex)
        );
    }

    function _checkGetAttestation(
        uint32 _nonce,
        uint256 _notaryIndex,
        uint256 _attestationIndex
    ) internal {
        _checkGetAttestation(_nonce, _notaryIndex, _notaryIndex, _attestationIndex);
    }

    function _checkGetAttestation(
        uint32 _nonce,
        uint256 _notaryIndex,
        uint256 _notaryGenerationIndex,
        uint256 _attestationIndex
    ) internal {
        bytes32 _root = _generateTestRoot(_nonce, _notaryGenerationIndex);
        bytes memory attestation = _generateTestAttestation(
            _nonce,
            _notaryIndex,
            _notaryGenerationIndex
        );
        assertEq(collector.getRoot(localDomain, _nonce, _attestationIndex), _root);
        assertEq(collector.getAttestation(localDomain, _nonce, _attestationIndex), attestation);
        assertEq(collector.getAttestation(localDomain, _nonce, _root), attestation);
    }

    // solhint-disable-next-line code-complexity
    function _checkLatestAttestations() internal {
        uint32 latestNonce = 0;
        bytes memory latestAttestation;
        for (uint256 i = 0; i < NOTARIES_AMOUNT; ++i) {
            if (attestedNonces[i].length == 0) continue;
            uint256 indexLast = attestedNonces[i].length - 1;
            if (attestedNonces[i][indexLast] == 0) continue;
            assert(attestedRoots[i][indexLast] != bytes32(0));
            (bytes memory attestation, ) = signOriginAttestation(
                notariesPK[i],
                attestedNonces[i][indexLast],
                attestedRoots[i][indexLast]
            );
            assertEq(collector.getLatestAttestation(localDomain, notaries[i]), attestation);

            if (attestedNonces[i][indexLast] > latestNonce) {
                latestNonce = attestedNonces[i][indexLast];
                latestAttestation = attestation;
            }
        }
        if (latestNonce != 0) {
            assertEq(collector.getLatestAttestation(localDomain), latestAttestation);
        }
    }
}
