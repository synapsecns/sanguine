// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyMessageLibHarness, LegacyMessageLib} from "../harnesses/LegacyMessageLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LegacyMessageLibTest is Test {
    struct LegacyMessage {
        address srcSender;
        address dstReceiver;
        uint64 srcNonce;
        bytes message;
    }

    LegacyMessageLibHarness public libHarness;

    function setUp() public {
        libHarness = new LegacyMessageLibHarness();
    }

    function test_encodeLegacyMessage() public {
        address srcSender = makeAddr("Sender");
        address dstReceiver = makeAddr("Receiver");
        uint64 srcNonce = 1337;
        bytes memory message = new bytes(512);
        bytes memory encoded = libHarness.encodeLegacyMessage(srcSender, dstReceiver, srcNonce, message);
        (address newSrcSender, address newDstReceiver, uint64 newSrcNonce, bytes memory newMessage) =
            libHarness.decodeLegacyMessage(encoded);
        assertEq(newSrcSender, srcSender);
        assertEq(newDstReceiver, dstReceiver);
        assertEq(newSrcNonce, srcNonce);
        assertEq(newMessage, message);
    }

    function test_encodeLegacyMessage_short() public {
        address srcSender = makeAddr("Sender");
        address dstReceiver = makeAddr("Receiver");
        uint64 srcNonce = 1337;
        bytes memory message = "Hello, World!";
        bytes memory encoded = libHarness.encodeLegacyMessage(srcSender, dstReceiver, srcNonce, message);
        (address newSrcSender, address newDstReceiver, uint64 newSrcNonce, bytes memory newMessage) =
            libHarness.decodeLegacyMessage(encoded);
        assertEq(newSrcSender, srcSender);
        assertEq(newDstReceiver, dstReceiver);
        assertEq(newSrcNonce, srcNonce);
        assertEq(newMessage, message);
    }

    function test_encodeLegacyMessageRoundtrip(LegacyMessage memory legacyMsg) public {
        bytes memory encoded = libHarness.encodeLegacyMessage(
            legacyMsg.srcSender, legacyMsg.dstReceiver, legacyMsg.srcNonce, legacyMsg.message
        );
        (address newSrcSender, address newDstReceiver, uint64 newSrcNonce, bytes memory newMessage) =
            libHarness.decodeLegacyMessage(encoded);
        assertEq(newSrcSender, legacyMsg.srcSender);
        assertEq(newDstReceiver, legacyMsg.dstReceiver);
        assertEq(newSrcNonce, legacyMsg.srcNonce);
        assertEq(newMessage, legacyMsg.message);
    }
}
