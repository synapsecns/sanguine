// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridgeV2, IFastBridgeV2} from "../contracts/FastBridgeV2.sol";

import {FastBridge, FastBridgeTest} from "./FastBridge.t.sol";

contract FastBridgeV2ParityTest is FastBridgeTest {
    function deployFastBridge() internal override {
        address fastBridgeV2 = address(new FastBridgeV2(owner));
        fastBridge = FastBridge(fastBridgeV2);
    }

    // ════════════════════════════════════════════ OVERRIDES FOR TESTS ════════════════════════════════════════════════

    /// @dev The same test, but expecting a different error
    function test_failedClaimNoProof() public virtual override {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, ) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);

        vm.warp(block.timestamp + 31 minutes);

        vm.expectRevert(senderIncorrectSelector());
        fastBridge.claim(request, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    // ════════════════════════════════════════ OVERRIDES FOR CUSTOM ERRORS ════════════════════════════════════════════

    // FastBridgeV2 uses a different naming convention for errors, so we need to override the error selectors

    function amountIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__AmountIncorrect.selector;
    }

    function chainIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__ChainIncorrect.selector;
    }

    function deadlineExceededSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineExceeded.selector;
    }

    function deadlineNotExceededSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineNotExceeded.selector;
    }

    function deadlineTooShortSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineTooShort.selector;
    }

    function disputePeriodNotPassedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DisputePeriodNotPassed.selector;
    }

    function disputePeriodPassedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DisputePeriodPassed.selector;
    }

    function msgValueIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__MsgValueIncorrect.selector;
    }

    function senderIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__SenderIncorrect.selector;
    }

    function statusIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__StatusIncorrect.selector;
    }

    function transactionRelayedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__TransactionRelayed.selector;
    }

    function zeroAddressSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__ZeroAddress.selector;
    }
}
