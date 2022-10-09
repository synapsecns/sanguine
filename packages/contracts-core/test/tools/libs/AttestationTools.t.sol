// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";

abstract contract AttestationTools is SynapseTestSuite {
    // Values specifying the need to mock the data instead of using the provided data
    uint32 internal constant MOCK_ATTESTATION_NONCE = type(uint32).max - 1337;
    bytes32 internal constant MOCK_ATTESTATION_ROOT = "Mock the attestation root";

    // Mock nonce
    uint32 private mockNonce = 42;
    // Saved attestation data
    address internal attestationNotary;
    uint32 internal attestationDomain;
    uint32 internal attestationNonce;
    bytes32 internal attestationRoot;
    // Saved attestation payloads
    bytes internal attestationRaw;
    bytes internal signatureNotary;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   CREATE ATTESTATION (GIVEN ROOT)                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Chain's default notary attestation with given nonce and given root
    function createAttestation(
        uint32 domain,
        uint32 nonce,
        bytes32 root
    ) public {
        // Use first Notary by default
        createAttestation({ domain: domain, nonce: nonce, root: root, notaryIndex: 0 });
    }

    // Chain's given notary attestation with given nonce and given root
    function createAttestation(
        uint32 domain,
        uint32 nonce,
        bytes32 root,
        uint256 notaryIndex
    ) public {
        createAttestation({
            domain: domain,
            nonce: nonce,
            root: root,
            signer: suiteNotary(domain, notaryIndex),
            salt: notaryIndex
        });
    }

    // Signer's attestation with given nonce and given root
    function createAttestation(
        uint32 domain,
        uint32 nonce,
        bytes32 root,
        address signer,
        uint256 salt
    ) public {
        saveAttestationData(domain, signer);
        saveMockableAttestationData(nonce, root, salt);
        createAttestation();
    }

    // Create attestation using all the saved data
    function createAttestation() public {
        (attestationRaw, signatureNotary) = signAttestation(
            attestationDomain,
            attestationNonce,
            attestationRoot,
            attestationNotary
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   CREATE ATTESTATION (MOCKED ROOT)                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Chain's default notary attestation with mocked nonce and root
    function createAttestationMock(uint32 domain) public {
        createAttestationMock(domain, MOCK_ATTESTATION_NONCE);
    }

    // Chain's default notary attestation with given nonce and mocked root
    function createAttestationMock(uint32 domain, uint32 nonce) public {
        createAttestationMock({ domain: domain, nonce: nonce, notaryIndex: 0, salt: 0 });
    }

    // Chain's given notary attestation with given nonce and mocked root
    function createAttestationMock(
        uint32 domain,
        uint32 nonce,
        uint256 notaryIndex,
        uint256 salt
    ) public {
        createAttestation({
            domain: domain,
            nonce: nonce,
            root: MOCK_ATTESTATION_ROOT,
            signer: suiteNotary(domain, notaryIndex),
            salt: salt
        });
    }

    // Signer's attestation with mocked nonce and root
    function createAttestationMock(uint32 domain, address signer) public {
        createAttestation({
            domain: domain,
            nonce: MOCK_ATTESTATION_NONCE,
            root: MOCK_ATTESTATION_ROOT,
            signer: signer,
            salt: 0
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SAVE TEST DATA                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function saveAttestationData(uint32 domain, address signer) public {
        attestationDomain = domain;
        attestationNotary = signer;
    }

    function saveMockableAttestationData(
        uint32 nonce,
        bytes32 root,
        uint256 salt
    ) public {
        // For mocking: use the mock value and increase it for later use
        attestationNonce = nonce == MOCK_ATTESTATION_NONCE ? mockNonce++ : nonce;
        // For mocking: use saved nonce and salt to create a mocked root
        attestationRoot = root == MOCK_ATTESTATION_ROOT
            ? _createMockRoot(attestationNonce, salt)
            : root;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectLogAttestation() public {
        vm.expectEmit(true, true, true, true);
        emit LogAttestation(attestationNotary, attestationRaw, attestationRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Generate a unique mock nonce
    function _createMockNonce() internal returns (uint32) {
        // Return old value and increase the mocked nonce
        return mockNonce++;
    }

    // Generate a unique mock root for given nonce and salt.
    // Using notary index as salt will lead to conflicting attestations:
    // different roots for the same nonce. Useful for fraud testing
    function _createMockRoot(uint32 nonce, uint256 salt) internal pure returns (bytes32) {
        return keccak256(abi.encode("test root", nonce, salt));
    }
}
