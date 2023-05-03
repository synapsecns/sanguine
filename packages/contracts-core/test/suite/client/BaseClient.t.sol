// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

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
        bytes memory content
    ) public {
        vm.assume(destination_ != 0 && destination_ != DOMAIN_LOCAL);
        vm.label(user, "User");
        // Set some sensible limit for fuzzed tips values
        tipsValue = tipsValue % (2 ** 32);
        uint160 encodedRequest = rr.encodeRequest();
        vm.deal(user, tipsValue);
        // Get expected values for sending a message
        bytes32 recipient = client.trustedSender(destination_);
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes memory expectedCall = abi.encodeWithSelector(
            InterfaceOrigin.sendBaseMessage.selector, destination_, recipient, optimisticPeriod, encodedRequest, content
        );
        vm.expectCall(origin, tipsValue, expectedCall);
        vm.prank(user);
        client.sendBaseMessage{value: tipsValue}(destination_, encodedRequest, content);
    }

    function test_sendBaseMessage_revert_recipientNotSet(address user, uint256 tipsValue, RawRequest memory rr)
        public
    {
        // There is no trustedSender for this domain => will revert in BaseClient
        uint32 destination_ = 0;
        vm.label(user, "User");
        // Set some sensible limit for fuzzed tips values
        tipsValue = tipsValue % (2 ** 32);
        uint160 encodedRequest = rr.encodeRequest();
        vm.deal(user, tipsValue);
        vm.expectRevert("BaseClient: !recipient");
        vm.prank(user);
        client.sendBaseMessage{value: tipsValue}(destination_, encodedRequest, "");
    }

    function test_receiveBaseMessage(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        uint256 secondsPassed,
        bytes memory content
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        secondsPassed = bound(secondsPassed, optimisticPeriod, 2 * optimisticPeriod);
        vm.warp(rootSubmittedAt + secondsPassed);
        vm.expectEmit();
        // msg.value should be zero
        emit BaseMessageReceived(0, rh.origin, rh.nonce, content);
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, secondsPassed, content);
    }

    function test_receiveBaseMessage_revert_notDestination(RawHeader memory rh, uint256 rootSubmittedAt, address caller)
        public
    {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL);
        vm.assume(caller != destination);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt + optimisticPeriod);
        vm.expectRevert("BaseClient: !destination");
        vm.prank(caller);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, "");
    }

    function test_receiveBaseMessage_revert_notTrustedSender(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        bytes32 sender
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL);
        vm.assume(sender != client.trustedSender(rh.origin));
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        vm.warp(rootSubmittedAt + optimisticPeriod);
        vm.expectRevert("BaseClient: !trustedSender");
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, optimisticPeriod, "");
    }

    function test_receiveBaseMessage_revert_optimisticPeriodNotOver(
        RawHeader memory rh,
        uint256 rootSubmittedAt,
        uint256 secondsPassed
    ) public {
        vm.assume(rh.origin != 0 && rh.origin != DOMAIN_LOCAL);
        // Get expected values for receiving a message
        uint32 optimisticPeriod = client.optimisticPeriod();
        bytes32 sender = client.trustedSender(rh.origin);
        // Set some sensible restrictions for timestamps
        rootSubmittedAt = bound(rootSubmittedAt, 1, 1e10);
        secondsPassed = bound(secondsPassed, 0, optimisticPeriod - 1);
        vm.warp(rootSubmittedAt + secondsPassed);
        vm.expectRevert("BaseClient: !optimisticPeriod");
        vm.prank(destination);
        client.receiveBaseMessage(rh.origin, rh.nonce, sender, secondsPassed, "");
    }
}
