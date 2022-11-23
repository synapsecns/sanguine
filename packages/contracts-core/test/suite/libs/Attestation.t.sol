// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/AttestationHarness.t.sol";

import "../../../contracts/libs/Auth.sol";

// solhint-disable func-name-mixedcase
contract AttestationLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    AttestationHarness internal libHarness;
    uint256 internal constant NOTARY_PRIV_KEY = 1337;
    // First element is (uint32 origin)
    uint8 internal constant FIRST_ELEMENT_BYTES = 32 / 8;

    function setUp() public override {
        super.setUp();
        libHarness = new AttestationHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root
    ) public {
        // Test formatting of attestation data
        bytes memory attData = libHarness.formatAttestationData(origin, destination, nonce, root);
        assertEq(
            attData,
            abi.encodePacked(origin, destination, nonce, root),
            "!formatAttestationData"
        );
        // Test formatting of attestation
        bytes memory notarySignature = signMessage(NOTARY_PRIV_KEY, attData);
        bytes memory attestation = libHarness.formatAttestation(
            origin,
            destination,
            nonce,
            root,
            notarySignature
        );
        // Test formatter against manually constructed payload
        assertEq(attestation, abi.encodePacked(attData, notarySignature), "!formatAttestation");
        // Both formatters should return the same results
        assertEq(
            attestation,
            libHarness.formatAttestation(attData, notarySignature),
            "!formatAttestation: different"
        );
        // Test "id formatters"
        assertEq(
            libHarness.attestationDomains(origin, destination),
            (uint256(origin) << 32) + destination,
            "!attestationDomains"
        );
        assertEq(
            libHarness.attestationKey(origin, destination, nonce),
            (uint256(origin) << 64) + (uint256(destination) << 32) + nonce,
            "!attestationKey"
        );
        // Test formatting checker
        assertTrue(libHarness.isAttestation(attestation), "!isAttestation");
        // Test getters
        assertEq(
            libHarness.attestedOrigin(SynapseTypes.ATTESTATION, attestation),
            origin,
            "!attestedOrigin"
        );
        assertEq(
            libHarness.attestedDestination(SynapseTypes.ATTESTATION, attestation),
            destination,
            "!attestedOrigin"
        );
        assertEq(
            libHarness.attestedNonce(SynapseTypes.ATTESTATION, attestation),
            nonce,
            "!attestedNonce"
        );
        assertEq(
            libHarness.attestedDomains(SynapseTypes.ATTESTATION, attestation),
            libHarness.attestationDomains(origin, destination),
            "!attestedDomains"
        );
        assertEq(
            libHarness.attestedKey(SynapseTypes.ATTESTATION, attestation),
            libHarness.attestationKey(origin, destination, nonce),
            "!attestedKey"
        );
        assertEq(
            libHarness.attestedRoot(SynapseTypes.ATTESTATION, attestation),
            root,
            "!attestedRoot"
        );
        // Test bytes29 getters
        checkBytes29Getter({
            getter: libHarness.castToAttestation,
            payloadType: SynapseTypes.ATTESTATION,
            payload: attestation,
            expectedType: SynapseTypes.ATTESTATION,
            expectedData: attestation,
            revertMessage: "!castToAttestation"
        });
        checkBytes29Getter({
            getter: libHarness.attestationData,
            payloadType: SynapseTypes.ATTESTATION,
            payload: attestation,
            expectedType: SynapseTypes.ATTESTATION_DATA,
            expectedData: attData,
            revertMessage: "!attestationData"
        });
        checkBytes29Getter({
            getter: libHarness.notarySignature,
            payloadType: SynapseTypes.ATTESTATION,
            payload: attestation,
            expectedType: SynapseTypes.SIGNATURE,
            expectedData: notarySignature,
            revertMessage: "!notarySignature"
        });
    }

    function test_isAttestation_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Attestation's first element (uint32 origin)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isAttestation(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isAttestation: short payload"
        );
    }

    function test_isAttestation_noSignature() public {
        // Use empty payload as signature
        bytes memory signature = "";
        bytes memory payload = libHarness.formatAttestation(
            uint32(0),
            uint32(0),
            uint32(0),
            bytes32(0),
            signature
        );
        assertFalse(libHarness.isAttestation(payload), "!isAttestation: no signature");
    }

    function test_isAttestation_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        assertTrue(libHarness.isAttestation(createTestPayload()), "!isAttestation: test payload");
    }

    function test_isAttestation_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        assertFalse(
            libHarness.isAttestation(cutLastByte(createTestPayload())),
            "!isAttestation: 1 byte shorter"
        );
    }

    function test_isAttestation_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        assertFalse(
            libHarness.isAttestation(addLastByte(createTestPayload())),
            "!isAttestation: 1 byte longer"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_attestedOrigin(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedOrigin(wrongType, payload);
    }

    function test_wrongTypeRevert_attestedDestination(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedDestination(wrongType, payload);
    }

    function test_wrongTypeRevert_attestedNonce(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedNonce(wrongType, payload);
    }

    function test_wrongTypeRevert_attestedRoot(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedRoot(wrongType, payload);
    }

    function test_wrongTypeRevert_attestedDomains(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedDomains(wrongType, payload);
    }

    function test_wrongTypeRevert_attestedKey(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedKey(wrongType, payload);
    }

    function test_wrongTypeRevert_attestationData(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestationData(wrongType, payload);
    }

    function test_wrongTypeRevert_notarySignature(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.notarySignature(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public pure returns (bytes memory) {
        return new bytes(Attestation.ATTESTATION_LENGTH);
    }
}
