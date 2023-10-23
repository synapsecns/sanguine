// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    BaseClientOptimisticPeriod,
    CallerNotDestination,
    IncorrectNonce,
    IncorrectSender,
    IncorrectRecipient,
    ZeroProofMaturity
} from "../../../contracts/libs/Errors.sol";
import {MessageRecipient} from "../../../contracts/client/MessageRecipient.sol";

import {BaseClientHarness} from "../../harnesses/client/BaseClientHarness.t.sol";
import {SynapseTest} from "../../utils/SynapseTest.t.sol";
import {InterfaceOrigin} from "../../mocks/OriginMock.t.sol";

import {RawBaseMessage, RawHeader, RawRequest} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
contract BaseClientTest is SynapseTest {
    BaseClientHarness public client;

    // Deploy mocks instead of the production contracts
    constructor() SynapseTest(0) {}

    function setUp() public override {
        super.setUp();
        client = new BaseClientHarness(origin, destination);
    }

    function test_sendBaseMessage(
        address user,
        uint32 destination_,
        uint256 tipsValue,
        RawRequest memory rr,
        uint32 nonce,
        bytes32 msgHash,
        bytes memory content
    ) public {
        vm.assume(destination_ != 0 && destination_ != DOMAIN_LOCAL);
        vm.label(user, "User");
        // Set some sensible limit for fuzzed tips values
        tipsValue = tipsValue % (2 ** 32);
        uint192 encodedRequest = rr.encodeRequest();
        vm.deal(user, tipsValue);
        // Get expected values for sending a message
        bytes32 recipient = client.trustedSender(destination_);
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes memory expectedCall = abi.encodeWithSelector(
            InterfaceOrigin.sendBaseMessage.selector, destination_, recipient, optimisticPeriod, encodedRequest, content
        );
        MessageRecipient.MessageRequest memory request =
            MessageRecipient.MessageRequest({gasDrop: rr.gasDrop, gasLimit: rr.gasLimit, version: rr.version});
        // Mock returned values for sendBaseMessage call
        vm.mockCall(
            origin, abi.encodeWithSelector(InterfaceOrigin.sendBaseMessage.selector), abi.encode(nonce, msgHash)
        );
        vm.expectCall(origin, tipsValue, expectedCall);
        vm.prank(user);
        (uint32 nonce_, bytes32 msgHash_) =
            client.sendBaseMessage{value: tipsValue}(destination_, tipsValue, request, content);
        assertEq(nonce_, nonce);
        assertEq(msgHash_, msgHash);
    }

    function test_sendBaseMessage_revert_recipientNotSet(
        address user,
        uint256 tipsValue,
        MessageRecipient.MessageRequest memory request
    ) public {
        // There is no trustedSender for this domain => will revert in BaseClient
        uint32 destination_ = 0;
        vm.label(user, "User");
        // Set some sensible limit for fuzzed tips values
        tipsValue = tipsValue % (2 ** 32);
        vm.deal(user, tipsValue);
        vm.expectRevert(IncorrectRecipient.selector);
        vm.prank(user);
        client.sendBaseMessage{value: tipsValue}(destination_, tipsValue, request, "");
    }

    function test_sendBaseMessage_tipsValueNotMsgValue() public {
        address user = makeAddr("User");
        uint32 destination_ = 1;
        uint256 tipsValue = 1337;
        uint256 msgValue = 6969;
        // Use empty request for this test
        MessageRecipient.MessageRequest memory request;
        vm.deal(user, msgValue);
        bytes memory expectedCall = abi.encodeWithSelector(InterfaceOrigin.sendBaseMessage.selector);
        // Should call origin with provided tipsValue as msg.value
        vm.expectCall(origin, tipsValue, expectedCall);
        vm.prank(user);
        client.sendBaseMessage{value: msgValue}(destination_, tipsValue, request, "");
    }

    function test_getMinimumTipsValue(
        uint32 destination_,
        RawRequest memory rr,
        uint256 contentLength,
        uint256 expectedResult
    ) public {
        uint192 encodedRequest = rr.encodeRequest();
        MessageRecipient.MessageRequest memory request =
            MessageRecipient.MessageRequest({gasDrop: rr.gasDrop, gasLimit: rr.gasLimit, version: rr.version});
        // (destination, paddedRequest, contentLength)
        bytes memory expectedCall = abi.encodeWithSelector(
            InterfaceOrigin.getMinimumTipsValue.selector, destination_, encodedRequest, contentLength
        );
        vm.mockCall(origin, expectedCall, abi.encode(expectedResult));
        vm.expectCall(origin, expectedCall);
        uint256 result = client.getMinimumTipsValue(destination_, request, contentLength);
        assertEq(result, expectedResult);
    }

    function test_receiveBaseMessage(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        uint256 secondsPassed,
        uint32 version,
        bytes memory content
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL && rh.nonce != 0);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        secondsPassed = bound(secondsPassed, optimisticPeriod, 2 * optimisticPeriod);
        vm.warp(rootSubmittedAt + secondsPassed);
        vm.expectEmit();
        // msg.value should be zero
        emit BaseMessageReceived(0, rh.origin, rh.nonce, version, content);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, secondsPassed, version, content);
    }

    function test_receiveBaseMessage_revert_notDestination(RawHeader memory rh, uint256 rootSubmittedAt, address caller)
        public
    {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL && rh.nonce != 0);
        vm.assume(caller != destination);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt + optimisticPeriod);
        vm.expectRevert(CallerNotDestination.selector);
        vm.prank(caller);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, 0, "");
    }

    function test_receiveBaseMessage_revert_notTrustedSender(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        bytes32 sender
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL && rh.nonce != 0);
        vm.assume(sender != client.trustedSender(rh.origin));
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt + optimisticPeriod);
        vm.expectRevert(IncorrectSender.selector);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, 0, "");
    }

    function test_receiveBaseMessage_revert_zeroNonce(RawHeader memory rh, uint256 rootSubmittedAt) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL);
        rh.nonce = 0;
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt + optimisticPeriod);
        vm.expectRevert(IncorrectNonce.selector);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, 0, "");
    }

    function test_receiveBaseMessage_revert_optimisticPeriodNotOver(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        uint256 secondsPassed
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL && rh.nonce != 0);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        secondsPassed = bound(secondsPassed, 1, optimisticPeriod - 1);
        vm.warp(rootSubmittedAt + secondsPassed);
        vm.expectRevert(BaseClientOptimisticPeriod.selector);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, secondsPassed, 0, "");
    }

    function test_receiveBaseMessage_revert_optimisticPeriodMinus1Second() public {
        uint32 timePassed = client.optimisticPeriod() - 1;
        bytes32 sender = client.trustedSender(DOMAIN_REMOTE);
        skip(timePassed);
        vm.expectRevert(BaseClientOptimisticPeriod.selector);
        vm.prank(destination);
        client.receiveBaseMessage({
            origin_: DOMAIN_REMOTE,
            nonce: 1,
            sender: sender,
            proofMaturity: timePassed,
            version: 0,
            content: ""
        });
    }

    function test_receiveBaseMessage_revert_zeroProofMaturity(RawHeader memory rh, uint256 rootSubmittedAt) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL && rh.nonce != 0);
        uint32 optimisticPeriod = 0;
        // Get expected values for receiving a message
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt);
        vm.expectRevert(ZeroProofMaturity.selector);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, 0, "");
    }
}
