// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SynapseTest } from "./utils/SynapseTest.sol";

import { MessageHarness } from "./harnesses/MessageHarness.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract MessageTest is SynapseTest {
    MessageHarness messageHarness;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    uint32 originDomain;
    bytes32 sender;
    uint32 nonce;
    uint32 destinationDomain;
    uint32 optimisticSeconds;
    bytes32 recipient;
    bytes tips;
    bytes messageBody;

    function setUp() public override {
        super.setUp();
        messageHarness = new MessageHarness();
        originDomain = 1000;
        sender = bytes32("AAAA THE SENDOOOOOR");
        nonce = 42;
        destinationDomain = 2000;
        optimisticSeconds = 4;
        recipient = bytes32("AAAA THE RECEIVOOOR");
        tips = getDefaultTips();
        messageBody = bytes("Messagoooor");
    }

    function test_formatMessage() public {
        bytes memory message = messageHarness.formatMessage(
            originDomain,
            sender,
            nonce,
            destinationDomain,
            recipient,
            optimisticSeconds,
            tips,
            messageBody
        );

        assertEq(messageHarness.origin(message), originDomain);
        assertEq(messageHarness.sender(message), sender);
        assertEq(messageHarness.nonce(message), nonce);
        assertEq(messageHarness.destination(message), destinationDomain);
        assertEq(messageHarness.recipient(message), recipient);
        assertEq(messageHarness.optimisticSeconds(message), optimisticSeconds);
        assertEq(messageHarness.tips(message), tips);
        assertEq(messageHarness.body(message), messageBody);
        assertEq(messageHarness.leaf(message), keccak256(message));
    }

    function test_messageHash() public {
        bytes memory message = messageHarness.formatMessage(
            originDomain,
            sender,
            nonce,
            destinationDomain,
            recipient,
            optimisticSeconds,
            tips,
            messageBody
        );

        bytes32 messageHash = messageHarness.messageHash(
            originDomain,
            sender,
            nonce,
            destinationDomain,
            recipient,
            optimisticSeconds,
            tips,
            messageBody
        );

        assertEq(messageHash, keccak256(message));
    }
}
