// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { Header } from "../contracts/libs/Header.sol";
import { Message } from "../contracts/libs/Message.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract HeaderTest is Test {
    using Header for bytes;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using Header for bytes29;
    using Message for bytes29;

    uint32 internal constant ORIGIN = 1234;
    bytes32 internal constant SENDER = bytes32("sender");
    uint32 internal constant NONCE = 3456;
    uint32 internal constant DESTINATION = 5678;
    bytes32 internal constant RECIPIENT = bytes32("recipient");
    uint32 internal constant OPTIMISTIC_SECONDS = 7890;

    uint40 internal constant WRONG_TYPE = 1337;

    function test_encodedCorrectly() public {
        bytes29 headerView = _createTestData();

        assertEq(headerView.headerVersion(), Header.HEADER_VERSION);
        assertEq(headerView.origin(), ORIGIN);
        assertEq(headerView.sender(), SENDER);
        assertEq(headerView.nonce(), NONCE);
        assertEq(headerView.destination(), DESTINATION);
        assertEq(headerView.recipient(), RECIPIENT);
        assertEq(headerView.optimisticSeconds(), OPTIMISTIC_SECONDS);
    }

    function test_incorrectType_headerVersion() public {
        _createTestDataMistyped().headerVersion();
    }

    function test_incorrectType_origin() public {
        _createTestDataMistyped().origin();
    }

    function test_incorrectType_sender() public {
        _createTestDataMistyped().sender();
    }

    function test_incorrectType_nonce() public {
        _createTestDataMistyped().nonce();
    }

    function test_incorrectType_destination() public {
        _createTestDataMistyped().destination();
    }

    function test_incorrectType_recipient() public {
        _createTestDataMistyped().recipient();
    }

    function test_incorrectType_recipientAddress() public {
        _createTestDataMistyped().recipientAddress();
    }

    function test_incorrectType_optimisticSeconds() public {
        _createTestDataMistyped().optimisticSeconds();
    }

    function _createTestData() internal pure returns (bytes29) {
        bytes memory _header = Header.formatHeader(
            ORIGIN,
            SENDER,
            NONCE,
            DESTINATION,
            RECIPIENT,
            OPTIMISTIC_SECONDS
        );
        return _header.headerView();
    }

    function _createTestDataMistyped() internal returns (bytes29 headerView) {
        headerView = _createTestData().castTo(WRONG_TYPE);
        vm.expectRevert(_expectedRevertMessage());
    }

    function _expectedRevertMessage() internal pure returns (bytes memory) {
        (, uint256 g) = TypedMemView.encodeHex(WRONG_TYPE);
        (, uint256 e) = TypedMemView.encodeHex(Message.HEADER_TYPE);
        return
            abi.encodePacked(
                "Type assertion failed. Got 0x",
                uint80(g),
                ". Expected 0x",
                uint80(e)
            );
    }
}
