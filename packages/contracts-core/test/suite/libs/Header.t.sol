// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SynapseLibraryTest, TypedMemView } from "../../utils/SynapseLibraryTest.t.sol";
import { HeaderHarness } from "../../harnesses/libs/HeaderHarness.t.sol";

import { HeaderLib } from "../../../contracts/libs/Header.sol";

// solhint-disable func-name-mixedcase
contract HeaderLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    HeaderHarness internal libHarness;
    // First element is (uint16 version)
    uint8 internal constant FIRST_ELEMENT_BYTES = 16 / 8;

    function setUp() public override {
        super.setUp();
        libHarness = new HeaderHarness();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: FORMATTING                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_formatHeader(
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
                HeaderLib.HEADER_VERSION,
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
        checkCastToHeader({ payload: payload, isHeader: true });
        // Test getters
        assertEq(libHarness.version(payload), HeaderLib.HEADER_VERSION, "!headerVersion");
        assertEq(libHarness.origin(payload), origin, "!origin");
        assertEq(libHarness.sender(payload), sender, "!sender");
        assertEq(libHarness.nonce(payload), nonce, "!nonce");
        assertEq(libHarness.destination(payload), destination, "!destination");
        assertEq(libHarness.recipient(payload), recipient, "!recipient");
        assertEq(libHarness.optimisticSeconds(payload), optimisticSeconds, "!optimisticSeconds");
        assertEq(
            libHarness.recipientAddress(payload),
            address(uint160(uint256(recipient))),
            "!recipientAddress"
        );
    }

    function test_isHeader_firstElementIncomplete(uint8 payloadLength, bytes32 data) public {
        // Payload having less bytes than Header's first element (uint16 version)
        // should be correctly treated as unformatted (i.e. with no reverts)
        bytes memory payload = createShortPayload(payloadLength, FIRST_ELEMENT_BYTES, data);
        checkCastToHeader({ payload: payload, isHeader: false });
    }

    function test_isHeader_testPayload() public {
        // Check that manually constructed test payload is considered formatted
        bytes memory payload = createTestPayload();
        checkCastToHeader({ payload: payload, isHeader: true });
    }

    function test_isHeader_shorterLength() public {
        // Check that manually constructed test payload without the last byte
        // is not considered formatted
        bytes memory payload = cutLastByte(createTestPayload());
        checkCastToHeader({ payload: payload, isHeader: false });
    }

    function test_isHeader_longerLength() public {
        // Check that manually constructed test payload with an extra last byte
        // is not considered formatted
        bytes memory payload = addLastByte(createTestPayload());
        checkCastToHeader({ payload: payload, isHeader: false });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               HELPERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function checkCastToHeader(bytes memory payload, bool isHeader) public {
        if (isHeader) {
            assertTrue(libHarness.isHeader(payload), "!isHeader: when valid");
            assertEq(libHarness.castToHeader(payload), payload, "!castToHeader: when valid");
        } else {
            assertFalse(libHarness.isHeader(payload), "!isHeader: when valid");
            vm.expectRevert("Not a header payload");
            libHarness.castToHeader(payload);
        }
    }

    function createTestPayload() public view returns (bytes memory) {
        return libHarness.formatHeader(0, bytes32(0), 0, 0, bytes32(0), 0);
    }
}
