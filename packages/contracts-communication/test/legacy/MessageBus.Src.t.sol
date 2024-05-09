// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusBaseTest, InterchainClientV1Mock, LegacyMessage, LegacyOptionsLib} from "./MessageBus.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract MessageBusSrcTest is MessageBusBaseTest {
    LegacyMessage public legacyMsg;
    bytes public encodedLegacyMsg;

    function setUp() public override {
        super.setUp();
        for (uint256 i = 0; i < MOCK_NONCE; ++i) {
            vm.prank(srcSender);
            messageBus.sendMessage({
                receiver: dstReceiverBytes32,
                dstChainId: REMOTE_CHAIN_ID,
                message: "",
                options: legacyOptions
            });
        }
        legacyMsg =
            LegacyMessage({srcSender: srcSender, dstReceiver: dstReceiver, srcNonce: MOCK_NONCE, message: MESSAGE});
        encodedLegacyMsg = encodeLegacyMessage(legacyMsg);
        deal(srcSender, MOCK_FEE);
    }

    function messageBusSendMessage(bytes32 receiver, bytes memory options) internal {
        vm.prank(srcSender);
        messageBus.sendMessage{value: MOCK_FEE}({
            receiver: receiver,
            dstChainId: REMOTE_CHAIN_ID,
            message: MESSAGE,
            options: options
        });
    }

    function test_sendMessage_callsClient() public {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainClientV1Mock.interchainSend,
            (REMOTE_CHAIN_ID, remoteMessageBusBytes32, execService, icModules, icOptions, encodedLegacyMsg)
        );
        vm.expectCall({callee: icClient, msgValue: MOCK_FEE, data: expectedCalldata, count: 1});
        messageBusSendMessage(dstReceiverBytes32, legacyOptions);
    }

    function test_sendMessage_emitsEvent() public {
        expectEventMessageSent({srcChainId: LOCAL_CHAIN_ID, dstChainId: REMOTE_CHAIN_ID, legacyMsg: legacyMsg});
        messageBusSendMessage(dstReceiverBytes32, legacyOptions);
    }

    function test_sendMessage_increasesNonce() public {
        messageBusSendMessage(dstReceiverBytes32, legacyOptions);
        assertEq(messageBus.nonce(), MOCK_NONCE + 1);
    }

    function test_sendMessage_revert_ReceiverNotEVM() public {
        bytes32 receiver = keccak256("GM");
        expectRevertReceiverNotEVM(receiver);
        messageBusSendMessage(receiver, legacyOptions);
    }

    function test_sendMessage_revert_incorrectOptionsVersion(uint16 version) public {
        vm.assume(version != LegacyOptionsLib.LEGACY_OPTIONS_VERSION);
        bytes memory invalidOpts = abi.encodePacked(version, uint256(1));
        expectRevertPayloadInvalid(invalidOpts);
        messageBusSendMessage(dstReceiverBytes32, invalidOpts);
    }

    function test_sendMessage_revert_incorrectOptionsLength(bytes memory invalidOpts) public {
        vm.assume(invalidOpts.length != legacyOptions.length);
        expectRevertPayloadInvalid(invalidOpts);
        messageBusSendMessage(dstReceiverBytes32, invalidOpts);
    }

    function test_estimateFee_revert_incorrectOptionsVersion(uint16 version) public {
        vm.assume(version != LegacyOptionsLib.LEGACY_OPTIONS_VERSION);
        bytes memory invalidOpts = abi.encodePacked(version, uint256(1));
        expectRevertPayloadInvalid(invalidOpts);
        messageBus.estimateFee(REMOTE_CHAIN_ID, invalidOpts);
    }

    function test_estimateFee_revert_incorrectOptionsLength(bytes memory invalidOpts) public {
        vm.assume(invalidOpts.length != legacyOptions.length);
        expectRevertPayloadInvalid(invalidOpts);
        messageBus.estimateFee(REMOTE_CHAIN_ID, invalidOpts);
    }

    function test_estimateFeeExact_revert_incorrectOptionsVersion(uint16 version) public {
        vm.assume(version != LegacyOptionsLib.LEGACY_OPTIONS_VERSION);
        bytes memory invalidOpts = abi.encodePacked(version, uint256(1));
        expectRevertPayloadInvalid(invalidOpts);
        messageBus.estimateFeeExact(REMOTE_CHAIN_ID, invalidOpts, MESSAGE.length);
    }

    function test_estimateFeeExact_revert_incorrectOptionsLength(bytes memory invalidOpts) public {
        vm.assume(invalidOpts.length != legacyOptions.length);
        expectRevertPayloadInvalid(invalidOpts);
        messageBus.estimateFeeExact(REMOTE_CHAIN_ID, invalidOpts, MESSAGE.length);
    }
}
