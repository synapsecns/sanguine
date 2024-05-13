// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusBaseTest, LegacyMessage, LegacyReceiverMock} from "./MessageBus.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract MessageBusDstTest is MessageBusBaseTest {
    uint64 public constant MOCK_DB_NONCE = 1_234_567;

    LegacyMessage public legacyMsg;
    bytes public encodedLegacyMsg;

    function setUp() public override {
        super.setUp();
        legacyMsg =
            LegacyMessage({srcSender: srcSender, dstReceiver: dstReceiver, srcNonce: MOCK_NONCE, message: MESSAGE});
        encodedLegacyMsg = encodeLegacyMessage(legacyMsg);
    }

    function messageBusAppReceive() internal {
        vm.prank(icClient);
        messageBus.appReceive({
            srcChainId: REMOTE_CHAIN_ID,
            sender: remoteMessageBusBytes32,
            dbNonce: MOCK_DB_NONCE,
            message: encodedLegacyMsg
        });
    }

    function test_appReceive_callsReceiver() public {
        bytes memory expectedCalldata =
            abi.encodeCall(LegacyReceiverMock.executeMessage, (srcSenderBytes32, REMOTE_CHAIN_ID, MESSAGE, icClient));
        vm.expectCall({callee: dstReceiver, msgValue: 0, data: expectedCalldata, count: 1});
        messageBusAppReceive();
    }

    function test_appReceive_emitsEvent() public {
        expectEventExecuted({srcChainId: REMOTE_CHAIN_ID, legacyMsg: legacyMsg});
        messageBusAppReceive();
    }
}
