// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcExclusivityTest} from "./FastBridgeV2.Src.Exclusivity.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcZapTest is FastBridgeV2SrcExclusivityTest {
    bytes public constant ZAP_DATA = abi.encode("Hello, World!");
    uint256 public constant ZAP_NATIVE = 1_337_420;

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        setTokenTestZapData(ZAP_DATA);
        setEthTestZapData(ZAP_DATA);
    }

    // Contract should accept zapData with length up to 2^16 - 1,
    // so that the zapData.length is encoded in 2 bytes.

    function test_bridge_token_zapDataLengthMax() public {
        bytes memory zapData = new bytes(2 ** 16 - 1);
        setTokenTestZapData(zapData);
        test_bridge_token();
    }

    function test_bridge_eth_zapDataLengthMax() public {
        bytes memory zapData = new bytes(2 ** 16 - 1);
        setEthTestZapData(zapData);
        test_bridge_eth();
    }

    function test_bridge_token_revert_zapDataLengthAboveMax() public {
        bytes memory zapData = new bytes(2 ** 16);
        setTokenTestZapData(zapData);
        vm.expectRevert(ZapDataLengthAboveMax.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
    }

    function test_bridge_eth_revert_zapDataLengthAboveMax() public {
        bytes memory zapData = new bytes(2 ** 16);
        setEthTestZapData(zapData);
        vm.expectRevert(ZapDataLengthAboveMax.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    // ══════════════════════════════════════ WITH CALL VALUE, NO CALL PARAMS ══════════════════════════════════════════

    function test_bridge_token_withZapNative_noZapData() public {
        setTokenTestZapData("");
        setTokenTestZapNative(ZAP_NATIVE);
        test_bridge_token();
    }

    function test_bridge_token_diffSender_withZapNative_noZapData() public {
        setTokenTestZapData("");
        setTokenTestZapNative(ZAP_NATIVE);
        test_bridge_token_diffSender();
    }

    function test_bridge_eth_withZapNative_noZapData_revert() public {
        setEthTestZapData("");
        setEthTestZapNative(ZAP_NATIVE);
        vm.expectRevert(ZapNativeNotSupported.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    function test_bridge_eth_diffSender_withZapNative_noZapData_revert() public {
        setEthTestZapData("");
        setEthTestZapNative(ZAP_NATIVE);
        vm.expectRevert(ZapNativeNotSupported.selector);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    // ═══════════════════════════════════════ WITH CALL VALUE & CALL PARAMS ═══════════════════════════════════════════

    function test_bridge_token_withZapNative_withZapData() public {
        setTokenTestZapNative(ZAP_NATIVE);
        test_bridge_token();
    }

    function test_bridge_token_diffSender_withZapNative_withZapData() public {
        setTokenTestZapNative(ZAP_NATIVE);
        test_bridge_token_diffSender();
    }

    function test_bridge_eth_withZapNative_withZapData_revert() public {
        setEthTestZapNative(ZAP_NATIVE);
        vm.expectRevert(ZapNativeNotSupported.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    function test_bridge_eth_diffSender_withZapNative_withZapData_revert() public {
        setEthTestZapNative(ZAP_NATIVE);
        vm.expectRevert(ZapNativeNotSupported.selector);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }
}
