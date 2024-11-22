// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TokenZapV1IntegrationTest, IFastBridge, IFastBridgeV2} from "./TokenZapV1.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2TokenZapV1SrcTest is TokenZapV1IntegrationTest {
    event BridgeRequested(
        bytes32 indexed transactionId,
        address indexed sender,
        bytes request,
        uint32 destChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        bool sendChainGas
    );

    function setUp() public virtual override {
        vm.chainId(SRC_CHAIN_ID);
        super.setUp();
    }

    function mintTokens() public virtual override {
        deal(user, SRC_AMOUNT);
        srcToken.mint(user, SRC_AMOUNT);
        vm.prank(user);
        srcToken.approve(address(fastBridge), type(uint256).max);
    }

    function bridge(
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2,
        bool isToken
    )
        public
    {
        vm.prank({msgSender: user, txOrigin: user});
        fastBridge.bridge{value: isToken ? 0 : SRC_AMOUNT}(params, paramsV2);
    }

    function expectEventBridgeRequested(
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2,
        bool isToken
    )
        public
    {
        bytes memory encodedBridgeTx = encodeBridgeTx(params, paramsV2);
        bytes32 txId = keccak256(encodedBridgeTx);
        vm.expectEmit(address(fastBridge));
        emit BridgeRequested({
            transactionId: txId,
            sender: user,
            request: encodedBridgeTx,
            destChainId: DST_CHAIN_ID,
            originToken: isToken ? address(srcToken) : NATIVE_GAS_TOKEN,
            destToken: isToken ? address(dstToken) : NATIVE_GAS_TOKEN,
            originAmount: SRC_AMOUNT,
            destAmount: DST_AMOUNT,
            sendChainGas: paramsV2.zapNative > 0
        });
    }

    function checkBalances(bool isToken) public view {
        if (isToken) {
            assertEq(srcToken.balanceOf(user), 0);
            assertEq(srcToken.balanceOf(address(fastBridge)), SRC_AMOUNT);
        } else {
            assertEq(address(user).balance, 0);
            assertEq(address(fastBridge).balance, SRC_AMOUNT);
        }
    }

    function test_bridge_depositTokenParams() public {
        expectEventBridgeRequested({params: tokenParams, paramsV2: depositTokenParams, isToken: true});
        bridge({params: tokenParams, paramsV2: depositTokenParams, isToken: true});
        checkBalances({isToken: true});
    }

    function test_bridge_depositTokenWithZapNativeParams() public {
        expectEventBridgeRequested({params: tokenParams, paramsV2: depositTokenWithZapNativeParams, isToken: true});
        bridge({params: tokenParams, paramsV2: depositTokenWithZapNativeParams, isToken: true});
        checkBalances({isToken: true});
    }

    function test_bridge_depositTokenRevertParams() public {
        expectEventBridgeRequested({params: tokenParams, paramsV2: depositTokenRevertParams, isToken: true});
        bridge({params: tokenParams, paramsV2: depositTokenRevertParams, isToken: true});
        checkBalances({isToken: true});
    }

    function test_bridge_depositNativeParams() public {
        expectEventBridgeRequested({params: nativeParams, paramsV2: depositNativeParams, isToken: false});
        bridge({params: nativeParams, paramsV2: depositNativeParams, isToken: false});
        checkBalances({isToken: false});
    }

    function test_bridge_depositNativeNoAmountParams() public {
        expectEventBridgeRequested({params: nativeParams, paramsV2: depositNativeNoAmountParams, isToken: false});
        bridge({params: nativeParams, paramsV2: depositNativeNoAmountParams, isToken: false});
        checkBalances({isToken: false});
    }

    function test_bridge_depositNativeRevertParams() public {
        expectEventBridgeRequested({params: nativeParams, paramsV2: depositNativeRevertParams, isToken: false});
        bridge({params: nativeParams, paramsV2: depositNativeRevertParams, isToken: false});
        checkBalances({isToken: false});
    }
}
