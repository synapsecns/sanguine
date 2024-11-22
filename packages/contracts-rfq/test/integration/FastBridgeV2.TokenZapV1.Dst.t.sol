// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {TokenZapV1IntegrationTest, VaultManyArguments, IFastBridge, IFastBridgeV2} from "./TokenZapV1.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2TokenZapV1DstTest is TokenZapV1IntegrationTest {
    event BridgeRelayed(
        bytes32 indexed transactionId,
        address indexed relayer,
        address indexed to,
        uint32 originChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        uint256 chainGasAmount
    );

    function setUp() public virtual override {
        vm.chainId(DST_CHAIN_ID);
        super.setUp();
    }

    function mintTokens() public virtual override {
        deal(relayer, DST_AMOUNT);
        dstToken.mint(relayer, DST_AMOUNT);
        vm.prank(relayer);
        dstToken.approve(address(fastBridge), type(uint256).max);
    }

    function relay(
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2,
        bool isToken
    )
        public
    {
        bytes memory encodedBridgeTx = encodeBridgeTx(params, paramsV2);
        vm.prank({msgSender: relayer, txOrigin: relayer});
        fastBridge.relay{value: isToken ? paramsV2.zapNative : DST_AMOUNT}(encodedBridgeTx);
    }

    function expectEventBridgeRelayed(
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2,
        bool isToken
    )
        public
    {
        bytes32 txId = keccak256(encodeBridgeTx(params, paramsV2));
        vm.expectEmit(address(fastBridge));
        emit BridgeRelayed({
            transactionId: txId,
            relayer: relayer,
            to: address(dstZap),
            originChainId: SRC_CHAIN_ID,
            originToken: isToken ? address(srcToken) : NATIVE_GAS_TOKEN,
            destToken: isToken ? address(dstToken) : NATIVE_GAS_TOKEN,
            originAmount: SRC_AMOUNT,
            destAmount: DST_AMOUNT,
            chainGasAmount: paramsV2.zapNative
        });
    }

    function checkBalances(bool isToken) public view {
        if (isToken) {
            assertEq(dstToken.balanceOf(user), 0);
            assertEq(dstToken.balanceOf(relayer), 0);
            assertEq(dstToken.balanceOf(address(fastBridge)), 0);
            assertEq(dstToken.balanceOf(address(dstZap)), 0);
            assertEq(dstToken.balanceOf(address(dstVault)), DST_AMOUNT);
            assertEq(dstVault.balanceOf(user, address(dstToken)), DST_AMOUNT);
        } else {
            assertEq(address(user).balance, 0);
            assertEq(address(relayer).balance, 0);
            assertEq(address(fastBridge).balance, 0);
            assertEq(address(dstZap).balance, 0);
            assertEq(address(dstVault).balance, DST_AMOUNT);
            assertEq(dstVault.balanceOf(user, NATIVE_GAS_TOKEN), DST_AMOUNT);
        }
    }

    function test_relay_depositTokenParams() public {
        expectEventBridgeRelayed({params: tokenParams, paramsV2: depositTokenParams, isToken: true});
        relay({params: tokenParams, paramsV2: depositTokenParams, isToken: true});
        checkBalances({isToken: true});
    }

    function test_relay_depositTokenWithZapNativeParams() public {
        expectEventBridgeRelayed({params: tokenParams, paramsV2: depositTokenWithZapNativeParams, isToken: true});
        relay({params: tokenParams, paramsV2: depositTokenWithZapNativeParams, isToken: true});
        checkBalances({isToken: true});
        // Extra ETH will be also custodied by the Vault
        assertEq(address(dstVault).balance, ZAP_NATIVE);
    }

    function test_relay_depositTokenRevertParams_revert() public {
        vm.expectRevert(VaultManyArguments.VaultManyArguments__SomeError.selector);
        relay({params: tokenParams, paramsV2: depositTokenRevertParams, isToken: true});
    }

    function test_relay_depositNativeParams() public {
        expectEventBridgeRelayed({params: nativeParams, paramsV2: depositNativeParams, isToken: false});
        relay({params: nativeParams, paramsV2: depositNativeParams, isToken: false});
        checkBalances({isToken: false});
    }

    function test_relay_depositNativeNoAmountParams() public {
        expectEventBridgeRelayed({params: nativeParams, paramsV2: depositNativeNoAmountParams, isToken: false});
        relay({params: nativeParams, paramsV2: depositNativeNoAmountParams, isToken: false});
        checkBalances({isToken: false});
    }

    function test_relay_depositNativeRevertParams_revert() public {
        vm.expectRevert(VaultManyArguments.VaultManyArguments__SomeError.selector);
        relay({params: nativeParams, paramsV2: depositNativeRevertParams, isToken: false});
    }
}
