// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2SrcExclusivityTest} from "./FastBridgeV2.Src.Exclusivity.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcArbitraryCallTest is FastBridgeV2SrcExclusivityTest {
    bytes public constant CALL_PARAMS = abi.encode("Hello, World!");
    uint256 public constant CALL_VALUE = 1_337_420;

    function createFixturesV2() public virtual override {
        super.createFixturesV2();
        setTokenTestCallParams(CALL_PARAMS);
        setEthTestCallParams(CALL_PARAMS);
    }

    // Contract should accept callParams with length up to 2^16 - 1,
    // so that the callParams.length is encoded in 2 bytes.

    function test_bridge_token_callParamsLengthMax() public {
        bytes memory callParams = new bytes(2 ** 16 - 1);
        setTokenTestCallParams(callParams);
        test_bridge_token();
    }

    function test_bridge_eth_callParamsLengthMax() public {
        bytes memory callParams = new bytes(2 ** 16 - 1);
        setEthTestCallParams(callParams);
        test_bridge_eth();
    }

    function test_bridge_token_revert_callParamsLengthAboveMax() public {
        bytes memory callParams = new bytes(2 ** 16);
        setTokenTestCallParams(callParams);
        vm.expectRevert(CallParamsLengthAboveMax.selector);
        bridge({caller: userA, msgValue: 0, params: tokenParams, paramsV2: tokenParamsV2});
    }

    function test_bridge_eth_revert_callParamsLengthAboveMax() public {
        bytes memory callParams = new bytes(2 ** 16);
        setEthTestCallParams(callParams);
        vm.expectRevert(CallParamsLengthAboveMax.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    // ══════════════════════════════════════ WITH CALL VALUE, NO CALL PARAMS ══════════════════════════════════════════

    function test_bridge_token_withCallValue_noCallParams() public {
        setTokenTestCallParams("");
        setTokenTestCallValue(CALL_VALUE);
        test_bridge_token();
    }

    function test_bridge_token_diffSender_withCallValue_noCallParams() public {
        setTokenTestCallParams("");
        setTokenTestCallValue(CALL_VALUE);
        test_bridge_token_diffSender();
    }

    function test_bridge_eth_withCallValue_noCallParams_revert() public {
        setEthTestCallParams("");
        setEthTestCallValue(CALL_VALUE);
        vm.expectRevert(NativeTokenCallValueNotSupported.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    function test_bridge_eth_diffSender_withCallValue_noCallParams_revert() public {
        setEthTestCallParams("");
        setEthTestCallValue(CALL_VALUE);
        vm.expectRevert(NativeTokenCallValueNotSupported.selector);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    // ═══════════════════════════════════════ WITH CALL VALUE & CALL PARAMS ═══════════════════════════════════════════

    function test_bridge_token_withCallValue_withCallParams() public {
        setTokenTestCallValue(CALL_VALUE);
        test_bridge_token();
    }

    function test_bridge_token_diffSender_withCallValue_withCallParams() public {
        setTokenTestCallValue(CALL_VALUE);
        test_bridge_token_diffSender();
    }

    function test_bridge_eth_withCallValue_withCallParams_revert() public {
        setEthTestCallValue(CALL_VALUE);
        vm.expectRevert(NativeTokenCallValueNotSupported.selector);
        bridge({caller: userA, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }

    function test_bridge_eth_diffSender_withCallValue_withCallParams_revert() public {
        setEthTestCallValue(CALL_VALUE);
        vm.expectRevert(NativeTokenCallValueNotSupported.selector);
        bridge({caller: userB, msgValue: ethParams.originAmount, params: ethParams, paramsV2: ethParamsV2});
    }
}
