// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {SystemMessageHarness} from "../../harnesses/libs/SystemMessageHarness.t.sol";

import {Random} from "../../utils/libs/Random.t.sol";
import {SystemEntity, RawSystemMessage} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract SystemMessageLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    uint256 internal constant OFFSET_ARGUMENTS = 6;

    SystemMessageHarness internal libHarness;

    function setUp() public {
        libHarness = new SystemMessageHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formattedCorrectly(RawSystemMessage memory rsm, uint256 words) public {
        // Make sure sender/recipient fit into SystemEntity
        rsm.sender = uint8(bound(rsm.sender, 0, uint8(type(SystemEntity).max)));
        rsm.recipient = uint8(bound(rsm.recipient, 0, uint8(type(SystemEntity).max)));
        // Set a sensible limit for the total payload length
        words = words % MAX_SYSTEM_CALL_WORDS;
        rsm.callData.args = Random("args").nextBytesWords(words);
        bytes memory callData = rsm.callData.formatCallData();
        // Format the system message
        bytes memory payload =
            libHarness.formatSystemMessage(SystemEntity(rsm.sender), SystemEntity(rsm.recipient), callData);
        // Test formatter against manually constructed payload
        assertEq(payload, abi.encodePacked(rsm.sender, rsm.recipient, callData), "!formatSystemMessage");
        checkCastToSystemMessage({payload: payload, isSystemMessage: true});
        // Test getters
        assertEq(uint8(libHarness.sender(payload)), rsm.sender, "!sender");
        assertEq(uint8(libHarness.recipient(payload)), rsm.recipient, "!recipient");
        assertEq(libHarness.callData(payload), callData, "!callData");
    }

    function test_isSystemMessage(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToSystemMessage({payload: payload, isSystemMessage: length % 32 == OFFSET_ARGUMENTS});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToSystemMessage(bytes memory payload, bool isSystemMessage) public {
        if (isSystemMessage) {
            assertTrue(libHarness.isSystemMessage(payload), "!isSystemMessage: when valid");
            assertEq(libHarness.castToSystemMessage(payload), payload, "!castToSystemMessage: when valid");
        } else {
            assertFalse(libHarness.isSystemMessage(payload), "!isSystemMessage: when valid");
            vm.expectRevert("Not a system message");
            libHarness.castToSystemMessage(payload);
        }
    }
}
