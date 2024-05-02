// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyMessageLibHarness} from "../harnesses/LegacyMessageLibHarness.sol";

import {Test} from "forge-std/Test.sol";

struct LegacyMessage {
    address srcSender;
    address dstReceiver;
    uint64 srcNonce;
    bytes message;
}

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LegacyMessageLibTest is Test {
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

    function test_encodeLegacyMessageRoundtrip(LegacyMessage memory legacyMsg) public view {
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

    function test_payloadSize(LegacyMessage memory legacyMsg) public view {
        uint256 size = libHarness.payloadSize(legacyMsg.message.length);
        uint256 expectedSize = libHarness.encodeLegacyMessage(
            legacyMsg.srcSender, legacyMsg.dstReceiver, legacyMsg.srcNonce, legacyMsg.message
        ).length;
        assertEq(size, expectedSize);
    }

    function test_payloadSize_fuzzBytesOnly(bytes memory message) public view {
        LegacyMessage memory legacyMsg;
        legacyMsg.message = message;
        test_payloadSize(legacyMsg);
    }
}
