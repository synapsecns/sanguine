// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {HyperCoreMessage, HyperCoreMessageHarness} from "../harnesses/HyperCoreMessageHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering, max-line-length
contract HyperCoreMessageTest is Test {
    HyperCoreMessageHarness internal harness;

    function setUp() public {
        harness = new HyperCoreMessageHarness();
    }

    function test_roundTrip(address recipient, address srcToken, uint256 amount) public view {
        bytes memory payload = harness.encodeHyperCoreMessage(recipient, srcToken, amount);
        (address decodedRecipient, address decodedSrcToken, uint256 decodedAmount) =
            harness.decodeHyperCoreMessage(payload);
        assertEq(decodedRecipient, recipient);
        assertEq(decodedSrcToken, srcToken);
        assertEq(decodedAmount, amount);
    }

    function test_decodeHyperCoreMessage_revert_invalidPayloadLength(uint16 length) public {
        vm.assume(length != 32 * 4);
        bytes memory payload = new bytes(length);
        vm.expectRevert(HyperCoreMessage.HyperCoreMessage__InvalidPayload.selector);
        harness.decodeHyperCoreMessage(payload);
    }

    function test_decodeHyperCoreMessage_revert_unsupportedVersion(uint8 version) public {
        vm.assume(version != 1);
        bytes memory payload = abi.encode(version, address(1), address(2), 3);
        vm.expectRevert(abi.encodeWithSelector(HyperCoreMessage.HyperCoreMessage__UnsupportedVersion.selector, version));
        harness.decodeHyperCoreMessage(payload);
    }
}
