// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { Message } from "../contracts/libs/Message.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { TypedMemView } from "../contracts/libs/TypedMemView.sol";

contract TipsTest is Test {
    using Tips for bytes;
    using TypedMemView for bytes29;
    using Tips for bytes29;

    uint96 internal constant NOTARY_TIP = 1234;
    uint96 internal constant BROADCASTER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant EXECUTOR_TIP = 7890;
    uint96 internal constant TOTAL_TIPS = NOTARY_TIP + BROADCASTER_TIP + PROVER_TIP + EXECUTOR_TIP;

    uint40 internal constant WRONG_TYPE = 1337;

    function test_correctTipsEncoding() public {
        bytes29 tipsView = _createTestData();

        assertEq(tipsView.notaryTip(), NOTARY_TIP);
        assertEq(tipsView.broadcasterTip(), BROADCASTER_TIP);
        assertEq(tipsView.proverTip(), PROVER_TIP);
        assertEq(tipsView.executorTip(), EXECUTOR_TIP);

        assertEq(tipsView.totalTips(), TOTAL_TIPS);
    }

    function test_incorrectType_notaryTip() public {
        _createTestDataMistyped().notaryTip();
    }

    function test_incorrectType_broadcasterTip() public {
        _createTestDataMistyped().broadcasterTip();
    }

    function test_incorrectType_proverTip() public {
        _createTestDataMistyped().proverTip();
    }

    function test_incorrectType_executorTip() public {
        _createTestDataMistyped().executorTip();
    }

    function _createTestData() internal pure returns (bytes29) {
        bytes memory tips = Tips.formatTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
        return tips.tipsView();
    }

    function _createTestDataMistyped() internal returns (bytes29 tipsView) {
        tipsView = _createTestData().castTo(WRONG_TYPE);
        vm.expectRevert(_expectedRevertMessage());
    }

    function _expectedRevertMessage() internal pure returns (bytes memory) {
        (, uint256 g) = TypedMemView.encodeHex(WRONG_TYPE);
        (, uint256 e) = TypedMemView.encodeHex(Message.TIPS_TYPE);
        return
            abi.encodePacked(
                "Type assertion failed. Got 0x",
                uint80(g),
                ". Expected 0x",
                uint80(e)
            );
    }
}
