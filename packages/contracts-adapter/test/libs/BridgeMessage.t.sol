// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {BridgeMessage, BridgeMessageHarness} from "../harnesses/BridgeMessageHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract BridgeMessageTest is Test {
    BridgeMessageHarness internal harness;

    function setUp() public {
        harness = new BridgeMessageHarness();
    }

    function test_roundTrip(address recipient, address srcToken, uint256 amount) public view {
        bytes memory payload = harness.encodeBridgeMessage(recipient, srcToken, amount);
        (address decodedRecipient, address decodedSrcToken, uint256 decodedAmount) =
            harness.decodeBridgeMessage(payload);
        assertEq(decodedRecipient, recipient);
        assertEq(decodedSrcToken, srcToken);
        assertEq(decodedAmount, amount);
    }

    function test_roundTrip_randomData() public {
        address recipient = makeAddr("Random Address");
        address srcToken = makeAddr("Random Token");
        uint256 amount = uint256(keccak256("Random Amount"));
        bytes memory payload = harness.encodeBridgeMessage(recipient, srcToken, amount);
        (address decodedRecipient, address decodedSrcToken, uint256 decodedAmount) =
            harness.decodeBridgeMessage(payload);
        assertEq(decodedRecipient, recipient);
        assertEq(decodedSrcToken, srcToken);
        assertEq(decodedAmount, amount);
    }

    function test_decodeBridgeMessage_revert_invalidPayloadLength(uint16 length) public {
        vm.assume(length != 32 * 3);
        bytes memory payload = new bytes(length);
        vm.expectRevert(BridgeMessage.BridgeMessage__InvalidPayload.selector);
        harness.decodeBridgeMessage(payload);
    }
}
