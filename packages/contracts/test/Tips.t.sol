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

    uint96 internal constant notary_TIP = 1234;
    uint96 internal constant RELAYER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant PROCESSOR_TIP = 7890;
    uint96 internal constant TOTAL_TIPS = notary_TIP + RELAYER_TIP + PROVER_TIP + PROCESSOR_TIP;

    uint40 internal constant WRONG_TYPE = 1337;

    function test_correctTipsEncoding() public {
        bytes29 tipsView = _createTestData();

        assertEq(tipsView.notaryTip(), notary_TIP);
        assertEq(tipsView.relayerTip(), RELAYER_TIP);
        assertEq(tipsView.proverTip(), PROVER_TIP);
        assertEq(tipsView.processorTip(), PROCESSOR_TIP);

        assertEq(tipsView.totalTips(), TOTAL_TIPS);
    }

    function test_incorrectType_notaryTip() public {
        _createTestDataMistyped().notaryTip();
    }

    function test_incorrectType_relayerTip() public {
        _createTestDataMistyped().relayerTip();
    }

    function test_incorrectType_proverTip() public {
        _createTestDataMistyped().proverTip();
    }

    function test_incorrectType_processorTip() public {
        _createTestDataMistyped().processorTip();
    }

    function _createTestData() internal pure returns (bytes29) {
        bytes memory tips = Tips.formatTips(notary_TIP, RELAYER_TIP, PROVER_TIP, PROCESSOR_TIP);
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
