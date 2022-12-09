// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/AttestationHarness.t.sol";

import "../../../contracts/libs/Auth.sol";

// solhint-disable func-name-mixedcase
contract AttestationLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    AttestationHarness internal libHarness;
    uint256 internal constant MAX_SIGNERS = 10;
    uint256 internal constant GUARD_PRIV_KEY = 7331;
    uint256 internal constant NOTARY_PRIV_KEY = 1337;
    // Anything shorter than this is never a formatted Attestation
    uint8 internal incompletePayloadLength = uint8(Attestation.OFFSET_FIRST_SIGNATURE);

    function setUp() public override {
        super.setUp();
        libHarness = new AttestationHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line code-complexity
    function test_formattedCorrectly(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root,
        uint256 guardSigs,
        uint256 notarySigs
    ) public {
        // Limit amount of signers
        guardSigs = guardSigs % MAX_SIGNERS;
        notarySigs = notarySigs % MAX_SIGNERS;
        uint256 agentSigs = guardSigs + notarySigs;
        // Should be at least one signer
        vm.assume(agentSigs != 0);
        // Test formatting of attestation data
        bytes memory attData = libHarness.formatAttestationData(origin, destination, nonce, root);
        assertEq(
            attData,
            abi.encodePacked(origin, destination, nonce, root),
            "!formatAttestationData"
        );
        // Test formatting of attestation
        bytes[] memory guardSignatures = new bytes[](guardSigs);
        for (uint256 i = 0; i < guardSigs; ++i) {
            // Use unique PK for each agent
            guardSignatures[i] = signMessage(GUARD_PRIV_KEY + i, attData);
        }
        bytes[] memory notarySignatures = new bytes[](notarySigs);
        for (uint256 i = 0; i < notarySigs; ++i) {
            // Use unique PK for each agent
            notarySignatures[i] = signMessage(NOTARY_PRIV_KEY + i, attData);
        }
        bytes memory attestation = libHarness.formatAttestation(
            attData,
            guardSignatures,
            notarySignatures
        );
        {
            bytes memory allSigs = "";
            // Test formatting of attestation
            for (uint256 i = 0; i < guardSigs; ++i) {
                allSigs = bytes.concat(allSigs, guardSignatures[i]);
            }

            for (uint256 i = 0; i < notarySigs; ++i) {
                allSigs = bytes.concat(allSigs, notarySignatures[i]);
            }
            // Sanity check
            assert(allSigs.length == 65 * agentSigs);
            // Test formatter against manually constructed payload
            assertEq(
                attestation,
                abi.encodePacked(attData, uint8(guardSigs), uint8(notarySigs), allSigs),
                "!formatAttestation"
            );
        }
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
        (uint8 _guardSigs, uint8 _notarySigs) = libHarness.agentSignatures(
            SynapseTypes.ATTESTATION,
            attestation
        );
        assertEq(_guardSigs, guardSigs, "!agentSignatures: guardSigs");
        assertEq(_notarySigs, notarySigs, "!agentSignatures: notarySigs");
        assertEq(
            libHarness.guardSignatures(SynapseTypes.ATTESTATION, attestation),
            guardSigs,
            "!guardSignatures"
        );
        assertEq(
            libHarness.notarySignatures(SynapseTypes.ATTESTATION, attestation),
            notarySigs,
            "!notarySignatures"
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
        for (uint256 i = 0; i < guardSigs; ++i) {
            libHarness.setIndex(i);
            checkBytes29Getter({
                getter: libHarness.guardSignature,
                payloadType: SynapseTypes.ATTESTATION,
                payload: attestation,
                expectedType: SynapseTypes.SIGNATURE,
                expectedData: guardSignatures[i],
                revertMessage: "!guardSignature"
            });
        }
        for (uint256 i = 0; i < notarySigs; ++i) {
            libHarness.setIndex(i);
            checkBytes29Getter({
                getter: libHarness.notarySignature,
                payloadType: SynapseTypes.ATTESTATION,
                payload: attestation,
                expectedType: SynapseTypes.SIGNATURE,
                expectedData: notarySignatures[i],
                revertMessage: "!notarySignature"
            });
        }
    }

    function test_isAttestation_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Attestation's first few elements (data + agentSigs)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, incompletePayloadLength, data);
        assertFalse(libHarness.isAttestation(payload), "!isAttestation: short payload");
    }

    function test_isAttestation_noSignatures() public {
        bytes memory attData = libHarness.formatAttestationData(
            uint32(0),
            uint32(0),
            uint32(0),
            bytes32(0)
        );
        bytes memory payload = libHarness.formatAttestation(
            attData,
            new bytes[](0),
            new bytes[](0)
        );
        assertFalse(libHarness.isAttestation(payload), "!isAttestation: no signatures");
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
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedOrigin(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestedDestination(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedDestination(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestedNonce(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedNonce(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestedRoot(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedRoot(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestedDomains(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedDomains(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestedKey(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestedKey(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_attestationData(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.attestationData(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_guardSignature(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.guardSignature(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_notarySignature(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.notarySignature(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_agentSignatures(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.agentSignatures(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_guardSignatures(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.guardSignatures(wrongType, createTestPayload());
    }

    function test_wrongTypeRevert_notarySignatures(uint40 wrongType) public {
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.ATTESTATION });
        libHarness.notarySignatures(wrongType, createTestPayload());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public pure returns (bytes memory) {
        bytes memory attData = new bytes(Attestation.ATTESTATION_DATA_LENGTH);
        bytes memory mockSig = new bytes(ByteString.SIGNATURE_LENGTH);
        return abi.encodePacked(attData, uint8(1), uint8(1), mockSig, mockSig);
    }
}
