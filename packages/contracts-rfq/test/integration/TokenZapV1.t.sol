// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {FastBridgeV2, IFastBridge, IFastBridgeV2} from "../../contracts/FastBridgeV2.sol";
import {BridgeTransactionV2Lib} from "../../contracts/libs/BridgeTransactionV2.sol";
import {ZapDataV1} from "../../contracts/libs/ZapDataV1.sol";
import {TokenZapV1} from "../../contracts/zaps/TokenZapV1.sol";

import {MockERC20} from "../MockERC20.sol";
import {VaultManyArguments} from "../mocks/VaultManyArguments.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable ordering
abstract contract TokenZapV1IntegrationTest is Test {
    address internal constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    uint32 internal constant SRC_CHAIN_ID = 1337;
    uint32 internal constant DST_CHAIN_ID = 7331;

    uint256 internal constant SRC_AMOUNT = 1 ether;
    uint256 internal constant DST_AMOUNT = 0.9999 ether;
    uint256 internal constant ZAP_NATIVE = 123_456;

    FastBridgeV2 internal fastBridge;
    TokenZapV1 internal dstZap;

    address internal user = makeAddr("User");
    address internal relayer = makeAddr("Relayer");

    MockERC20 internal srcToken;
    MockERC20 internal dstToken;

    VaultManyArguments internal dstVault;

    IFastBridge.BridgeParams internal tokenParams;
    IFastBridge.BridgeParams internal nativeParams;

    IFastBridgeV2.BridgeParamsV2 internal depositTokenParams;
    IFastBridgeV2.BridgeParamsV2 internal depositTokenWithZapNativeParams;
    IFastBridgeV2.BridgeParamsV2 internal depositTokenRevertParams;
    IFastBridgeV2.BridgeParamsV2 internal depositNativeParams;
    IFastBridgeV2.BridgeParamsV2 internal depositNativeNoAmountParams;
    IFastBridgeV2.BridgeParamsV2 internal depositNativeRevertParams;

    function setUp() public virtual {
        fastBridge = new FastBridgeV2(address(this));
        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), address(this));
        fastBridge.addProver(relayer);

        srcToken = new MockERC20("SRC", 18);
        dstToken = new MockERC20("DST", 18);

        dstZap = new TokenZapV1();
        dstVault = new VaultManyArguments();

        createFixtures();
        mintTokens();
    }

    function createFixtures() public virtual {
        tokenParams = IFastBridge.BridgeParams({
            dstChainId: DST_CHAIN_ID,
            sender: user,
            to: address(dstZap),
            originToken: address(srcToken),
            destToken: address(dstToken),
            originAmount: SRC_AMOUNT,
            destAmount: DST_AMOUNT,
            sendChainGas: false,
            deadline: block.timestamp + 1 days
        });
        nativeParams = IFastBridge.BridgeParams({
            dstChainId: DST_CHAIN_ID,
            sender: user,
            to: address(dstZap),
            originToken: NATIVE_GAS_TOKEN,
            destToken: NATIVE_GAS_TOKEN,
            originAmount: SRC_AMOUNT,
            destAmount: DST_AMOUNT,
            sendChainGas: false,
            deadline: block.timestamp + 1 days
        });
        // Deposit token
        bytes memory zapData = dstZap.encodeZapData({
            target: address(dstVault),
            payload: getDepositPayload(address(dstToken)),
            amountPosition: 4 + 32 * 2
        });
        depositTokenParams.zapData = zapData;
        depositTokenWithZapNativeParams.zapData = zapData;
        depositTokenWithZapNativeParams.zapNative = ZAP_NATIVE;
        // Deposit native
        depositNativeParams.zapData = dstZap.encodeZapData({
            target: address(dstVault),
            payload: getDepositPayload(NATIVE_GAS_TOKEN),
            amountPosition: 4 + 32 * 2
        });
        // Deposit no amount
        depositNativeNoAmountParams.zapData = dstZap.encodeZapData({
            target: address(dstVault),
            payload: getDepositNoAmountPayload(),
            amountPosition: ZapDataV1.AMOUNT_NOT_PRESENT
        });
        // Deposit revert
        depositTokenRevertParams.zapData = dstZap.encodeZapData({
            target: address(dstVault),
            payload: getDepositRevertPayload(),
            amountPosition: ZapDataV1.AMOUNT_NOT_PRESENT
        });
        depositNativeRevertParams.zapData = dstZap.encodeZapData({
            target: address(dstVault),
            payload: getDepositRevertPayload(),
            amountPosition: ZapDataV1.AMOUNT_NOT_PRESENT
        });
    }

    function mintTokens() public virtual;

    function encodeBridgeTx(
        IFastBridge.BridgeParams memory params,
        IFastBridgeV2.BridgeParamsV2 memory paramsV2
    )
        public
        pure
        returns (bytes memory)
    {
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTx = IFastBridgeV2.BridgeTransactionV2({
            originChainId: SRC_CHAIN_ID,
            destChainId: params.dstChainId,
            originSender: params.sender,
            destRecipient: params.to,
            originToken: params.originToken,
            destToken: params.destToken,
            originAmount: params.originAmount,
            destAmount: params.destAmount,
            // No protocol fees for the test
            originFeeAmount: 0,
            deadline: params.deadline,
            // Single tx is sent, so nonce is 0
            nonce: 0,
            exclusivityRelayer: address(0),
            exclusivityEndTime: 0,
            zapNative: paramsV2.zapNative,
            zapData: paramsV2.zapData
        });
        return BridgeTransactionV2Lib.encodeV2(bridgeTx);
    }

    function getDepositPayload(address token) public view returns (bytes memory) {
        return abi.encodeCall(dstVault.deposit, (token, abi.encode(token), DST_AMOUNT, user, abi.encode(user)));
    }

    function getDepositNoAmountPayload() public view returns (bytes memory) {
        return abi.encodeCall(dstVault.depositNoAmount, (user));
    }

    function getDepositRevertPayload() public view returns (bytes memory) {
        return abi.encodeCall(dstVault.depositWithRevert, ());
    }
}
