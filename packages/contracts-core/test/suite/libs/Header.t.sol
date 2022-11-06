// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseLibraryTest.t.sol";
import "../../harnesses/libs/HeaderHarness.t.sol";

import "../../../contracts/libs/Header.sol";

// solhint-disable func-name-mixedcase
contract HeaderLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    HeaderHarness internal libHarness;
    // First element is (uint32 origin)
    uint8 internal constant FIRST_ELEMENT_BYTES = 32 / 8;

    function setUp() public override {
        super.setUp();
        libHarness = new HeaderHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formattedCorrectly(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticSeconds
    ) public {
        // Test formatting
        bytes memory payload = libHarness.formatHeader(
            origin,
            sender,
            nonce,
            destination,
            recipient,
            optimisticSeconds
        );
        assertEq(
            payload,
            abi.encodePacked(
                Header.HEADER_VERSION,
                origin,
                sender,
                nonce,
                destination,
                recipient,
                optimisticSeconds
            ),
            "!formatHeader"
        );
        // Test formatting checker
        assertTrue(libHarness.isHeader(payload), "!isHeader");
        // Test getters
        assertEq(libHarness.origin(SynapseTypes.MESSAGE_HEADER, payload), origin, "!origin");
        assertEq(libHarness.sender(SynapseTypes.MESSAGE_HEADER, payload), sender, "!sender");
        assertEq(libHarness.nonce(SynapseTypes.MESSAGE_HEADER, payload), nonce, "!nonce");
        assertEq(
            libHarness.destination(SynapseTypes.MESSAGE_HEADER, payload),
            destination,
            "!destination"
        );
        assertEq(
            libHarness.recipient(SynapseTypes.MESSAGE_HEADER, payload),
            recipient,
            "!recipient"
        );
        assertEq(
            libHarness.optimisticSeconds(SynapseTypes.MESSAGE_HEADER, payload),
            optimisticSeconds,
            "!optimisticSeconds"
        );
        assertEq(
            libHarness.recipientAddress(SynapseTypes.MESSAGE_HEADER, payload),
            address(uint160(uint256(recipient))),
            "!origin"
        );
    }

    function test_isHeader_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Header's first element (uint32 origin)
        // should be correctly treated as unformatted (i.e. with no reverts)
        assertFalse(
            libHarness.isHeader(createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data)),
            "!isHeader: short payload"
        );
    }

    function test_isHeader_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        assertTrue(libHarness.isHeader(createTestPayload()), "!isHeader: test payload");
    }

    function test_isHeader_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        assertFalse(
            libHarness.isHeader(cutLastByte(createTestPayload())),
            "!isHeader: 1 byte shorter"
        );
    }

    function test_isHeader_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        assertFalse(
            libHarness.isHeader(addLastByte(createTestPayload())),
            "!isHeader: 1 byte longer"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: WRONG TYPE                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_wrongTypeRevert_origin(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.origin(wrongType, payload);
    }

    function test_wrongTypeRevert_sender(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.sender(wrongType, payload);
    }

    function test_wrongTypeRevert_nonce(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.nonce(wrongType, payload);
    }

    function test_wrongTypeRevert_destination(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.destination(wrongType, payload);
    }

    function test_wrongTypeRevert_recipient(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.recipient(wrongType, payload);
    }

    function test_wrongTypeRevert_optimisticSeconds(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.optimisticSeconds(wrongType, payload);
    }

    function test_wrongTypeRevert_recipientAddress(uint40 wrongType) public {
        bytes memory payload = createTestPayload();
        expectRevertWrongType({ wrongType: wrongType, correctType: SynapseTypes.MESSAGE_HEADER });
        libHarness.recipientAddress(wrongType, payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatHeader(0, bytes32(0), 0, 0, bytes32(0), 0);
    }
}
