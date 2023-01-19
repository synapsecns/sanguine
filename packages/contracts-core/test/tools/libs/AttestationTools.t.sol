// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";

abstract contract AttestationTools is SynapseTestSuite {
    // Values specifying the need to mock the data instead of using the provided data
    uint32 internal constant MOCK_ATTESTATION_NONCE = type(uint32).max - 1337;
    bytes32 internal constant MOCK_ATTESTATION_ROOT = "Mock the attestation root";

    // Mock nonce
    uint32 internal mockNonce = 42;
    // Saved attestation data
    address[] internal attestationGuards;
    address[] internal attestationNotaries;
    RawAttestation internal ra;
    // Saved attestation ids
    uint64 internal attestationDomains;
    uint96 internal attestationKey;
    // Saved attestation payloads
    bytes internal attestationRaw;
    bytes internal signaturesGuard;
    bytes internal signaturesNotary;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   CREATE ATTESTATION (GIVEN ROOT)                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Chain's default notary attestation with given nonce and given root
    function createAttestation(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root
    ) public {
        // Use first Guard and Notary by default
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: root,
            guardIndex: 0,
            notaryIndex: 0
        });
    }

    // Chain's given notary attestation with given nonce and given root
    function createAttestation(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root,
        uint256 guardIndex,
        uint256 notaryIndex
    ) public {
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: root,
            guardSigner: suiteGuard(guardIndex),
            notarySigner: suiteNotary(destination, notaryIndex),
            salt: notaryIndex
        });
    }

    // Signer's attestation with given nonce and given root
    function createAttestation(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root,
        address guardSigner,
        address notarySigner,
        uint256 salt
    ) public {
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: root,
            guardSigners: castToArray(guardSigner),
            notarySigners: castToArray(notarySigner),
            salt: salt
        });
    }

    // Signer's attestation with given nonce and given root
    function createAttestation(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root,
        address[] memory guardSigners,
        address[] memory notarySigners,
        uint256 salt
    ) public {
        saveAttestationData(origin, destination, guardSigners, notarySigners);
        saveMockableAttestationData(nonce, root, salt);
        saveAttestationIDs();
        createAttestation();
    }

    // Create attestation using all the saved data
    function createAttestation() public {
        // This will also save the attestation data into ra.data
        (ra.data, attestationRaw, signaturesGuard, signaturesNotary) = signAttestation(
            ra,
            attestationGuards,
            attestationNotaries
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    CREATE ATTESTATION (SAME DATA)                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createSameAttestation() public returns (bytes memory) {
        createAttestation();
        return attestationRaw;
    }

    function createSameAttestation(bool isGuard, address agent) public returns (bytes memory) {
        if (isGuard) {
            attestationGuards = castToArray(agent);
            delete attestationNotaries;
        } else {
            delete attestationGuards;
            attestationNotaries = castToArray(agent);
        }
        return createSameAttestation();
    }

    // Create attestation with the same data, having exactly one signature
    function createSameAttestation(uint256 guardSigs, uint256 notarySigs)
        public
        returns (bytes memory)
    {
        (attestationGuards, attestationNotaries) = _createSigners(
            ra.destination,
            guardSigs,
            notarySigs
        );
        return createSameAttestation();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   CREATE ATTESTATION (MOCKED ROOT)                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Chain's default notary attestation with mocked nonce and root
    function createAttestationMock(uint32 origin, uint32 destination) public {
        createAttestationMock(origin, destination, MOCK_ATTESTATION_NONCE);
    }

    // Chain's default notary attestation with given nonce and mocked root
    function createAttestationMock(
        uint32 origin,
        uint32 destination,
        uint32 nonce
    ) public {
        createAttestationMock({
            origin: origin,
            destination: destination,
            nonce: nonce,
            guardIndex: 0,
            notaryIndex: 0,
            salt: 0
        });
    }

    // Chain's given notary attestation with given nonce and mocked root
    function createAttestationMock(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        uint256 guardIndex,
        uint256 notaryIndex,
        uint256 salt
    ) public {
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: nonce,
            root: MOCK_ATTESTATION_ROOT,
            guardSigner: suiteGuard(guardIndex),
            notarySigner: suiteNotary(destination, notaryIndex),
            salt: salt
        });
    }

    // Signer's attestation with mocked nonce and root
    function createAttestationMock(
        uint32 origin,
        uint32 destination,
        address[] memory guardSigners,
        address[] memory notarySigners
    ) public {
        createAttestation({
            origin: origin,
            destination: destination,
            nonce: MOCK_ATTESTATION_NONCE,
            root: MOCK_ATTESTATION_ROOT,
            guardSigners: guardSigners,
            notarySigners: notarySigners,
            salt: 0
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SAVE TEST DATA                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function saveAttestationData(
        uint32 origin,
        uint32 destination,
        address[] memory guardSigners,
        address[] memory notarySigners
    ) public {
        ra.origin = origin;
        ra.destination = destination;
        attestationGuards = guardSigners;
        attestationNotaries = notarySigners;
    }

    function saveMockableAttestationData(
        uint32 nonce,
        bytes32 root,
        uint256 salt
    ) public {
        // For mocking: use the mock value and increase it for later use
        ra.nonce = nonce == MOCK_ATTESTATION_NONCE ? _createMockNonce() : nonce;
        // For mocking: use saved nonce and salt to create a mocked root
        ra.root = root == MOCK_ATTESTATION_ROOT ? _createMockRoot(ra.nonce, salt) : root;
    }

    function saveAttestationMetadata() public {
        saveAttestationMetadata(block.number, block.timestamp);
    }

    function saveAttestationMetadata(uint256 blockNumber, uint256 timestamp) public {
        ra.blockNumber = uint40(blockNumber);
        ra.timestamp = uint40(timestamp);
    }

    function saveAttestationIDs() public {
        attestationDomains = AttestationLib.packDomains(ra.origin, ra.destination);
        attestationKey = AttestationLib.packKey(ra.origin, ra.destination, ra.nonce);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAttestationAccepted() public {
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(attestationGuards, attestationNotaries, attestationRaw);
    }

    function expectLogAttestation() public {
        vm.expectEmit(true, true, true, true);
        emit LogAttestation(attestationGuards, attestationNotaries, attestationRaw, attestationRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Generate a unique mock nonce
    function _createMockNonce() internal returns (uint32) {
        // Return old value and increase the mocked nonce
        return mockNonce++;
    }

    function _createSigners(
        uint32 destination,
        uint256 guardSigs,
        uint256 notarySigs
    ) internal view returns (address[] memory guardSigners, address[] memory notarySigners) {
        return _createSigners(destination, guardSigs, notarySigs, type(uint256).max);
    }

    function _createSigners(
        uint32 destination,
        uint256 guardSigs,
        uint256 notarySigs,
        uint256 attackerIndex
    ) internal view returns (address[] memory guardSigners, address[] memory notarySigners) {
        guardSigners = new address[](guardSigs);
        for (uint256 i = 0; i < guardSigs; ++i) {
            guardSigners[i] = (i == attackerIndex) ? attacker : suiteGuard(i);
        }
        notarySigners = new address[](notarySigs);
        for (uint256 i = 0; i < notarySigs; ++i) {
            notarySigners[i] = (guardSigs + i == attackerIndex)
                ? attacker
                : suiteNotary(destination, i);
        }
    }

    // Generate a unique mock root for given nonce and salt.
    // Using notary index as salt will lead to conflicting attestations:
    // different roots for the same nonce. Useful for fraud testing
    function _createMockRoot(uint32 nonce, uint256 salt) internal pure returns (bytes32) {
        return keccak256(abi.encode("test root", nonce, salt));
    }
}
