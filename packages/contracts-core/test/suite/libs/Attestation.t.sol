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
    uint8 internal incompletePayloadLength = uint8(AttestationLib.OFFSET_FIRST_SIGNATURE);

    function setUp() public override {
        super.setUp();
        libHarness = new AttestationHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line code-complexity
    function test_formattedCorrectly_attestation(
        RawAttestation memory ra,
        uint256 guardsAmount,
        uint256 notariesAmount
    ) public {
        // Limit amount of signers
        guardsAmount = guardsAmount % MAX_SIGNERS;
        notariesAmount = notariesAmount % MAX_SIGNERS;
        uint256 agentSigs = guardsAmount + notariesAmount;
        // Should be at least one signer
        vm.assume(agentSigs != 0);
        // Test formatting of attestation data
        bytes memory attData = libHarness.formatAttestationData(
            ra.origin,
            ra.destination,
            ra.nonce,
            ra.root,
            ra.blockNumber,
            ra.timestamp
        );
        // Test formatting of attestation
        bytes[] memory guardSignatures = new bytes[](guardsAmount);
        bytes[] memory notarySignatures = new bytes[](notariesAmount);
        bytes memory attestation;
        {
            uint256[] memory guardPrivKeys = new uint256[](guardsAmount);
            for (uint256 i = 0; i < guardsAmount; ++i) {
                // Use unique PK for each agent
                guardPrivKeys[i] = GUARD_PRIV_KEY + i;
                guardSignatures[i] = signMessage(guardPrivKeys[i], attData);
            }
            bytes memory guardSignaturesPayload = signMessage(guardPrivKeys, attData);
            uint256[] memory notaryPrivKeys = new uint256[](notariesAmount);
            for (uint256 i = 0; i < notariesAmount; ++i) {
                // Use unique PK for each agent
                notaryPrivKeys[i] = NOTARY_PRIV_KEY + i;
                notarySignatures[i] = signMessage(notaryPrivKeys[i], attData);
            }
            bytes memory notarySignaturesPayload = signMessage(notaryPrivKeys, attData);
            attestation = libHarness.formatAttestation(
                attData,
                guardSignaturesPayload,
                notarySignaturesPayload
            );
            // Both formatters should return the same payload
            assertEq(
                libHarness.formatAttestationFromViews(
                    attData,
                    guardSignaturesPayload,
                    notarySignaturesPayload
                ),
                attestation,
                "!formatAttestationFromViews"
            );
        }
        {
            bytes memory allSigs = "";
            for (uint256 i = 0; i < guardsAmount; ++i) {
                allSigs = bytes.concat(allSigs, guardSignatures[i]);
            }
            for (uint256 i = 0; i < notariesAmount; ++i) {
                allSigs = bytes.concat(allSigs, notarySignatures[i]);
            }
            // Sanity check
            assert(allSigs.length == 65 * agentSigs);
            // Test formatter against manually constructed payload
            assertEq(
                attestation,
                abi.encodePacked(attData, uint8(guardsAmount), uint8(notariesAmount), allSigs),
                "!formatAttestation"
            );
        }
        checkCastToAttestation({ payload: attestation, isAttestation: true });
        // Test getters
        assertEq(libHarness.data(attestation), attData, "!data");
        (uint8 _guardsAmount, uint8 _notariesAmount) = libHarness.agentsAmount(attestation);
        assertEq(_guardsAmount, guardsAmount, "!agentsAmount: guardsAmount");
        assertEq(_notariesAmount, notariesAmount, "!agentsAmount: notariesAmount");
        assertEq(libHarness.guardsAmount(attestation), guardsAmount, "!guardsAmount");
        assertEq(libHarness.notariesAmount(attestation), notariesAmount, "!notariesAmount");
        // Check signature getters
        for (uint256 i = 0; i < guardsAmount; ++i) {
            assertEq(
                libHarness.guardSignature(attestation, i),
                guardSignatures[i],
                "!guardSignature"
            );
        }
        for (uint256 i = 0; i < notariesAmount; ++i) {
            assertEq(
                libHarness.notarySignature(attestation, i),
                notarySignatures[i],
                "!notarySignature"
            );
        }
    }

    function test_formatAttestation_revert_incorrectSigPayloadLength() public {
        bytes memory attData = new bytes(AttestationLib.ATTESTATION_DATA_LENGTH);
        bytes memory singleSig = new bytes(ByteString.SIGNATURE_LENGTH);
        bytes memory shorterSig = new bytes(singleSig.length - 1);
        bytes memory longerSig = new bytes(singleSig.length + 1);
        vm.expectRevert("!signaturesLength");
        libHarness.formatAttestation(attData, shorterSig, singleSig);
        vm.expectRevert("!signaturesLength");
        libHarness.formatAttestation(attData, singleSig, longerSig);
    }

    function test_formatAttestation_revert_tooManySignatures() public {
        bytes memory attData = new bytes(AttestationLib.ATTESTATION_DATA_LENGTH);
        bytes memory singleSig = new bytes(ByteString.SIGNATURE_LENGTH);
        bytes memory manySigs = new bytes(256 * ByteString.SIGNATURE_LENGTH);
        vm.expectRevert("Too many signatures");
        libHarness.formatAttestation(attData, manySigs, singleSig);
        vm.expectRevert("Too many signatures");
        libHarness.formatAttestation(attData, singleSig, manySigs);
    }

    function test_isAttestation_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Attestation's first few elements (data + agentSigs)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, incompletePayloadLength, data);
        checkCastToAttestation({ payload: payload, isAttestation: false });
    }

    function test_isAttestation_noSignatures() public {
        bytes memory attData = new bytes(AttestationLib.ATTESTATION_DATA_LENGTH);
        bytes memory payload = libHarness.formatAttestation(attData, new bytes(0), new bytes(0));
        checkCastToAttestation({ payload: payload, isAttestation: false });
    }

    function test_isAttestation_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        bytes memory payload = createTestPayload();
        checkCastToAttestation({ payload: payload, isAttestation: true });
    }

    function test_isAttestation_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        bytes memory payload = cutLastByte(createTestPayload());
        checkCastToAttestation({ payload: payload, isAttestation: false });
    }

    function test_isAttestation_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        bytes memory payload = addLastByte(createTestPayload());
        checkCastToAttestation({ payload: payload, isAttestation: false });
    }

    function test_guardSignature_revert_outOfRange(
        uint8 guardsAmount,
        uint8 notariesAmount,
        uint8 guardIndex
    ) public {
        vm.assume(guardsAmount != 0 || notariesAmount != 0);
        vm.assume(guardIndex >= guardsAmount);
        bytes memory attestation = createTestPayload(guardsAmount, notariesAmount);
        vm.expectRevert("Out of range");
        libHarness.guardSignature(attestation, guardIndex);
    }

    function test_notarySignature_revert_outOfRange(
        uint8 guardsAmount,
        uint8 notariesAmount,
        uint8 notaryIndex
    ) public {
        vm.assume(guardsAmount != 0 || notariesAmount != 0);
        vm.assume(notaryIndex >= notariesAmount);
        bytes memory attestation = createTestPayload(guardsAmount, notariesAmount);
        vm.expectRevert("Out of range");
        libHarness.notarySignature(attestation, notaryIndex);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ATTESTATION DATA                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly_attestationData(RawAttestation memory ra) public {
        // Test formatting of attestation data
        bytes memory attData = libHarness.formatAttestationData(
            ra.origin,
            ra.destination,
            ra.nonce,
            ra.root,
            ra.blockNumber,
            ra.timestamp
        );
        assertEq(
            attData,
            abi.encodePacked(
                ra.origin,
                ra.destination,
                ra.nonce,
                ra.root,
                ra.blockNumber,
                ra.timestamp
            ),
            "!formatAttestationData"
        );
        checkCastToAttestationData({ payload: attData, isAttestationData: true });
        // Test pack/unpack of fields
        uint64 _domains = (uint64(ra.origin) << 32) + ra.destination;
        uint96 _key = (uint96(ra.origin) << 64) + (uint96(ra.destination) << 32) + ra.nonce;
        {
            assertEq(libHarness.packDomains(ra.origin, ra.destination), _domains, "!packDomains");
            assertEq(libHarness.packKey(ra.origin, ra.destination, ra.nonce), _key, "!packKey");
            (uint32 unpackedOrigin, uint32 unpackedDestination) = libHarness.unpackDomains(
                _domains
            );
            assertEq(unpackedOrigin, ra.origin, "!unpackDomains: origin");
            assertEq(unpackedDestination, ra.destination, "!unpackDomains: destination");
            (uint64 unpackedDomains, uint32 unpackedNonce) = libHarness.unpackKey(_key);
            assertEq(unpackedDomains, _domains, "!unpackKey: domains");
            assertEq(unpackedNonce, ra.nonce, "!unpackKey: nonce");
        }
        // Test getters
        assertEq(libHarness.origin(attData), ra.origin, "!origin");
        assertEq(libHarness.destination(attData), ra.destination, "!destination");
        assertEq(libHarness.nonce(attData), ra.nonce, "!nonce");
        assertEq(libHarness.domains(attData), _domains, "!domains");
        assertEq(libHarness.key(attData), _key, "!key");
        assertEq(libHarness.root(attData), ra.root, "!root");
        assertEq(libHarness.blockNumber(attData), ra.blockNumber, "!blockNumber");
        assertEq(libHarness.timestamp(attData), ra.timestamp, "!timestamp");
    }

    function test_isAttestationData(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToAttestationData({
            payload: payload,
            isAttestationData: length == AttestationLib.ATTESTATION_DATA_LENGTH
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToAttestation(bytes memory payload, bool isAttestation) public {
        if (isAttestation) {
            assertTrue(libHarness.isAttestation(payload), "!isAttestation: when valid");
            assertEq(
                libHarness.castToAttestation(payload),
                payload,
                "!castToAttestation: when valid"
            );
        } else {
            assertFalse(libHarness.isAttestation(payload), "!isAttestation: when valid");
            vm.expectRevert("Not an attestation");
            libHarness.castToAttestation(payload);
        }
    }

    function checkCastToAttestationData(bytes memory payload, bool isAttestationData) public {
        if (isAttestationData) {
            assertTrue(libHarness.isAttestationData(payload), "!isAttestationData: when valid");
            assertEq(
                libHarness.castToAttestationData(payload),
                payload,
                "!castToAttestation: when valid"
            );
        } else {
            assertFalse(libHarness.isAttestationData(payload), "!isAttestationData: when valid");
            vm.expectRevert("Not an attestation data");
            libHarness.castToAttestationData(payload);
        }
    }

    function createTestPayload() public pure returns (bytes memory) {
        return createTestPayload({ guardsAmount: 1, notariesAmount: 1 });
    }

    function createTestPayload(uint8 guardsAmount, uint8 notariesAmount)
        public
        pure
        returns (bytes memory)
    {
        bytes memory mockData = new bytes(AttestationLib.ATTESTATION_DATA_LENGTH);
        bytes memory mockGuardSigs = new bytes(ByteString.SIGNATURE_LENGTH * guardsAmount);
        bytes memory mockNotarySigs = new bytes(ByteString.SIGNATURE_LENGTH * notariesAmount);
        return
            abi.encodePacked(mockData, guardsAmount, notariesAmount, mockGuardSigs, mockNotarySigs);
    }
}
