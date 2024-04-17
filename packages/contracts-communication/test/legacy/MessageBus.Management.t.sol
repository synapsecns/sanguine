// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyMessageLib} from "../../contracts/legacy/libs/LegacyMessage.sol";

import {MessageBusBaseTest, InterchainClientV1Mock} from "./MessageBus.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract MessageBusManagementTest is MessageBusBaseTest {
    uint256 public constant LENGTH_ESTIMATE = 100;

    function mockInterchainFees(uint256 length) internal {
        uint256 legacyMessageLen = LegacyMessageLib.payloadSize(length);
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainClientV1Mock.getInterchainFee,
            (REMOTE_CHAIN_ID, execService, icModules, icOptions, legacyMessageLen)
        );
        vm.mockCall(icClient, expectedCalldata, abi.encode(MOCK_FEE));
    }

    function test_setMessageLengthEstimate_emitsEvent() public {
        expectEventMessageLengthEstimateSet(LENGTH_ESTIMATE);
        vm.prank(governor);
        messageBus.setMessageLengthEstimate(LENGTH_ESTIMATE);
    }

    function test_setMessageLengthEstimate_setsLength() public {
        vm.prank(governor);
        messageBus.setMessageLengthEstimate(LENGTH_ESTIMATE);
        assertEq(messageBus.messageLengthEstimate(), LENGTH_ESTIMATE);
    }

    function test_setMessageLengthEstimate_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        messageBus.setMessageLengthEstimate(LENGTH_ESTIMATE);
    }

    function test_estimateFee_usesLengthEstimate() public {
        mockInterchainFees(LENGTH_ESTIMATE);
        vm.prank(governor);
        messageBus.setMessageLengthEstimate(LENGTH_ESTIMATE);
        uint256 fee = messageBus.estimateFee(REMOTE_CHAIN_ID, legacyOptions);
        assertEq(fee, MOCK_FEE);
    }

    function test_estimateFeeExact() public {
        mockInterchainFees(2 * LENGTH_ESTIMATE);
        vm.prank(governor);
        messageBus.setMessageLengthEstimate(LENGTH_ESTIMATE);
        uint256 fee = messageBus.estimateFeeExact(REMOTE_CHAIN_ID, legacyOptions, 2 * LENGTH_ESTIMATE);
        assertEq(fee, MOCK_FEE);
    }
}
