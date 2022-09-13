// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Bytes29Test } from "../utils/Bytes29Test.sol";
import { SynapseTypes } from "../../contracts/libs/SynapseTypes.sol";
import { Tips } from "../../contracts/libs/Tips.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

// solhint-disable func-name-mixedcase

contract TipsTest is Bytes29Test {
    using Tips for bytes;
    using TypedMemView for bytes29;
    using Tips for bytes29;

    uint96 internal constant NOTARY_TIP = 1234;
    uint96 internal constant BROADCASTER_TIP = 3456;
    uint96 internal constant PROVER_TIP = 5678;
    uint96 internal constant EXECUTOR_TIP = 7890;
    uint96 internal constant TOTAL_TIPS = NOTARY_TIP + BROADCASTER_TIP + PROVER_TIP + EXECUTOR_TIP;

    function test_formattedCorrectly() public {
        bytes29 _view = _createTestView();

        assertTrue(_view.isTips());

        assertEq(_view.notaryTip(), NOTARY_TIP);
        assertEq(_view.broadcasterTip(), BROADCASTER_TIP);
        assertEq(_view.proverTip(), PROVER_TIP);
        assertEq(_view.executorTip(), EXECUTOR_TIP);

        assertEq(_view.totalTips(), TOTAL_TIPS);
    }

    function test_incorrectType_notaryTip() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_TIPS).notaryTip();
    }

    function test_incorrectType_broadcasterTip() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_TIPS).broadcasterTip();
    }

    function test_incorrectType_proverTip() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_TIPS).proverTip();
    }

    function test_incorrectType_executorTip() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_TIPS).executorTip();
    }

    function test_incorrectType_totalTips() public {
        _prepareMistypedTest(SynapseTypes.MESSAGE_TIPS).totalTips();
    }

    function test_isTips_incorrectVersion() public {
        bytes memory _tips = abi.encodePacked(uint16(0), new bytes(Tips.TIPS_LENGTH - 2));
        assert(_tips.length == Tips.TIPS_LENGTH);
        assertFalse(_tips.castToTips().isTips());
    }

    function test_isTips_emptyPayload() public {
        bytes memory _tips = bytes("");
        assert(_tips.length == 0);
        assertFalse(_tips.castToTips().isTips());
    }

    function test_isTips_tooShort() public {
        bytes memory _tips = abi.encodePacked(uint16(1), new bytes(Tips.TIPS_LENGTH - 3));
        assert(_tips.length < Tips.TIPS_LENGTH);
        assertFalse(_tips.castToTips().isTips());
    }

    function test_isTips_tooLong() public {
        bytes memory _tips = abi.encodePacked(uint16(1), new bytes(Tips.TIPS_LENGTH - 1));
        assert(_tips.length > Tips.TIPS_LENGTH);
        assertFalse(_tips.castToTips().isTips());
    }

    function _createTestView() internal pure override returns (bytes29) {
        bytes memory tips = Tips.formatTips(NOTARY_TIP, BROADCASTER_TIP, PROVER_TIP, EXECUTOR_TIP);
        return tips.castToTips();
    }
}
