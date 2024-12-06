// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseIntentRouter} from "../../contracts/interfaces/ISynapseIntentRouter.sol";
import {IDefaultExtendedPool} from "../../contracts/legacy/router/interfaces/IDefaultExtendedPool.sol";
import {Action, DefaultParams} from "../../contracts/legacy/router/libs/Structs.sol";
import {SynapseIntentPreviewer} from "../../contracts/router/SynapseIntentPreviewer.sol";

import {ZapDataV1Harness} from "../harnesses/ZapDataV1Harness.sol";

import {DefaultPoolMock} from "../mocks/DefaultPoolMock.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {LimitedToken, SwapQuery, SwapQuoterMock} from "../mocks/SwapQuoterMock.sol";
import {WETHMock} from "../mocks/WETHMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseIntentPreviewerTest is Test {
    address internal constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 internal constant AMOUNT_IN = 1.337 ether;
    uint256 internal constant SWAP_AMOUNT_OUT = 4.2 ether;
    uint256 internal constant ALL_ACTIONS_MASK = type(uint256).max;
    uint256 internal constant FULL_AMOUNT = type(uint256).max;

    uint8 internal constant TOKEN_IN_INDEX = 2;
    uint8 internal constant TOKEN_OUT_INDEX = 1;
    uint8 internal constant TOKENS = 3;
    uint8 internal constant LP_TOKEN_INDEX = type(uint8).max;

    ZapDataV1Harness internal zapDataLib;

    SynapseIntentPreviewer internal sip;
    address internal defaultPoolMock;
    address internal swapQuoterMock;

    address internal weth;
    address internal tokenA;
    address internal tokenB;
    address internal lpToken;

    address internal routerAdapterMock = makeAddr("Router Adapter Mock");

    function setUp() public {
        sip = new SynapseIntentPreviewer();

        defaultPoolMock = address(new DefaultPoolMock());
        swapQuoterMock = address(new SwapQuoterMock());

        weth = address(new WETHMock());
        tokenA = address(new MockERC20("A", 18));
        tokenB = address(new MockERC20("B", 18));
        lpToken = address(new MockERC20("LP", 18));

        zapDataLib = new ZapDataV1Harness();

        vm.label(defaultPoolMock, "DefaultPoolMock");
        vm.label(swapQuoterMock, "SwapQuoterMock");
        vm.label(weth, "WETHMock");
        vm.label(tokenA, "TokenA");
        vm.label(tokenB, "TokenB");
        vm.label(lpToken, "LPToken");
        vm.label(address(zapDataLib), "ZapDataV1Harness");

        vm.mockCall({
            callee: defaultPoolMock,
            data: abi.encodeCall(DefaultPoolMock.swapStorage, ()),
            returnData: abi.encode(0, 0, 0, 0, 0, 0, lpToken)
        });
    }

    function mockGetAmountOut(address tokenIn, address tokenOut, uint256 amountIn, SwapQuery memory mockQuery) public {
        LimitedToken memory token = LimitedToken({actionMask: ALL_ACTIONS_MASK, token: tokenIn});
        vm.mockCall({
            callee: swapQuoterMock,
            data: abi.encodeCall(SwapQuoterMock.getAmountOut, (token, tokenOut, amountIn)),
            returnData: abi.encode(mockQuery)
        });
    }

    function mockGetToken(uint8 tokenIndex, address token) public {
        vm.mockCall({
            callee: defaultPoolMock,
            data: abi.encodeCall(DefaultPoolMock.getToken, (tokenIndex)),
            returnData: abi.encode(token)
        });
    }

    function getSwapQuery(address tokenOut) public view returns (SwapQuery memory) {
        return SwapQuery({
            routerAdapter: routerAdapterMock,
            tokenOut: tokenOut,
            minAmountOut: SWAP_AMOUNT_OUT,
            deadline: type(uint256).max,
            rawParams: abi.encode(
                DefaultParams({
                    action: Action.Swap,
                    pool: defaultPoolMock,
                    tokenIndexFrom: TOKEN_IN_INDEX,
                    tokenIndexTo: TOKEN_OUT_INDEX
                })
            )
        });
    }

    function getSwapZapData() public view returns (bytes memory) {
        return getSwapZapData(TOKEN_IN_INDEX, TOKEN_OUT_INDEX);
    }

    function getSwapZapData(uint8 indexIn, uint8 indexOut) public view returns (bytes memory) {
        return zapDataLib.encodeV1({
            target_: defaultPoolMock,
            // swap(tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
            payload_: abi.encodeCall(DefaultPoolMock.swap, (indexIn, indexOut, 0, 0, type(uint256).max)),
            // Amount (dx) is encoded as the third parameter
            amountPosition_: 4 + 32 * 2
        });
    }

    function test_getSwapZapData() public view {
        for (uint8 i = 0; i < TOKENS; i++) {
            for (uint8 j = 0; j < TOKENS; j++) {
                bytes memory zapData = getSwapZapData(i, j);
                bytes memory payload = zapDataLib.payload(zapData, AMOUNT_IN);
                // swap(tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
                assertEq(payload, abi.encodeCall(DefaultPoolMock.swap, (i, j, AMOUNT_IN, 0, type(uint256).max)));
            }
        }
    }

    function getAddLiquidityQuery(address tokenOut) public view returns (SwapQuery memory) {
        return SwapQuery({
            routerAdapter: routerAdapterMock,
            tokenOut: tokenOut,
            minAmountOut: SWAP_AMOUNT_OUT,
            deadline: type(uint256).max,
            rawParams: abi.encode(
                DefaultParams({
                    action: Action.AddLiquidity,
                    pool: defaultPoolMock,
                    tokenIndexFrom: TOKEN_IN_INDEX,
                    tokenIndexTo: LP_TOKEN_INDEX
                })
            )
        });
    }

    function getAddLiquidityZapData() public view returns (bytes memory) {
        return getAddLiquidityZapData(TOKEN_IN_INDEX);
    }

    function getAddLiquidityZapData(uint8 indexIn) public view returns (bytes memory) {
        uint256[] memory amounts = new uint256[](TOKENS);
        return zapDataLib.encodeV1({
            target_: defaultPoolMock,
            // addLiquidity(amounts, minToMint, deadline)
            payload_: abi.encodeCall(IDefaultExtendedPool.addLiquidity, (amounts, 0, type(uint256).max)),
            // Amount is encoded within `amounts` at `TOKEN_IN_INDEX`, `amounts` is encoded after
            // (amounts.offset, minToMint, deadline, amounts.length)
            amountPosition_: 4 + 32 * (4 + indexIn)
        });
    }

    function test_getAddLiquidityZapData() public view {
        for (uint8 i = 0; i < TOKENS; i++) {
            bytes memory zapData = getAddLiquidityZapData(i);
            bytes memory payload = zapDataLib.payload(zapData, AMOUNT_IN);
            uint256[] memory amounts = new uint256[](TOKENS);
            amounts[i] = AMOUNT_IN;
            // addLiquidity(amounts, minToMint, deadline)
            assertEq(payload, abi.encodeCall(IDefaultExtendedPool.addLiquidity, (amounts, 0, type(uint256).max)));
        }
    }

    function getRemoveLiquidityQuery(address tokenOut) public view returns (SwapQuery memory) {
        return SwapQuery({
            routerAdapter: routerAdapterMock,
            tokenOut: tokenOut,
            minAmountOut: SWAP_AMOUNT_OUT,
            deadline: type(uint256).max,
            rawParams: abi.encode(
                DefaultParams({
                    action: Action.RemoveLiquidity,
                    pool: defaultPoolMock,
                    tokenIndexFrom: LP_TOKEN_INDEX,
                    tokenIndexTo: TOKEN_OUT_INDEX
                })
            )
        });
    }

    function getRemoveLiquidityZapData() public view returns (bytes memory) {
        return getRemoveLiquidityZapData(TOKEN_OUT_INDEX);
    }

    function getRemoveLiquidityZapData(uint8 indexOut) public view returns (bytes memory) {
        return zapDataLib.encodeV1({
            target_: defaultPoolMock,
            // removeLiquidityOneToken(tokenAmount, tokenIndex, minAmount, deadline)
            payload_: abi.encodeCall(IDefaultExtendedPool.removeLiquidityOneToken, (0, indexOut, 0, type(uint256).max)),
            // Amount (tokenAmount) is encoded as the first parameter
            amountPosition_: 4
        });
    }

    function test_getRemoveLiquidityZapData() public view {
        for (uint8 i = 0; i < TOKENS; i++) {
            bytes memory zapData = getRemoveLiquidityZapData(i);
            bytes memory payload = zapDataLib.payload(zapData, AMOUNT_IN);
            // removeLiquidityOneToken(tokenAmount, tokenIndex, minAmount, deadline)
            assertEq(
                payload,
                abi.encodeCall(IDefaultExtendedPool.removeLiquidityOneToken, (AMOUNT_IN, i, 0, type(uint256).max))
            );
        }
    }

    function getWrapETHQuery(address tokenOut) public view returns (SwapQuery memory) {
        return SwapQuery({
            routerAdapter: routerAdapterMock,
            tokenOut: tokenOut,
            minAmountOut: AMOUNT_IN,
            deadline: type(uint256).max,
            rawParams: abi.encode(
                DefaultParams({
                    action: Action.HandleEth,
                    pool: address(0),
                    tokenIndexFrom: LP_TOKEN_INDEX,
                    tokenIndexTo: LP_TOKEN_INDEX
                })
            )
        });
    }

    function getWrapETHZapData() public view returns (bytes memory) {
        return zapDataLib.encodeV1({
            target_: weth,
            // deposit()
            payload_: abi.encodeCall(WETHMock.deposit, ()),
            // Amount is not encoded
            amountPosition_: zapDataLib.AMOUNT_NOT_PRESENT()
        });
    }

    function test_getWrapETHZapData() public view {
        bytes memory zapData = getWrapETHZapData();
        bytes memory payload = zapDataLib.payload(zapData, AMOUNT_IN);
        // deposit()
        assertEq(payload, abi.encodeCall(WETHMock.deposit, ()));
    }

    function getUnwrapWETHQuery(address tokenOut) public view returns (SwapQuery memory) {
        return SwapQuery({
            routerAdapter: routerAdapterMock,
            tokenOut: tokenOut,
            minAmountOut: AMOUNT_IN,
            deadline: type(uint256).max,
            rawParams: abi.encode(
                DefaultParams({
                    action: Action.HandleEth,
                    pool: address(0),
                    tokenIndexFrom: LP_TOKEN_INDEX,
                    tokenIndexTo: LP_TOKEN_INDEX
                })
            )
        });
    }

    function getUnwrapWETHZapData() public view returns (bytes memory) {
        return zapDataLib.encodeV1({
            target_: weth,
            // withdraw(amount)
            payload_: abi.encodeCall(WETHMock.withdraw, (0)),
            // Amount is encoded as the first parameter
            amountPosition_: 4
        });
    }

    function test_getUnwrapWETHZapData() public view {
        bytes memory zapData = getUnwrapWETHZapData();
        bytes memory payload = zapDataLib.payload(zapData, AMOUNT_IN);
        // withdraw(amount)
        assertEq(payload, abi.encodeCall(WETHMock.withdraw, (AMOUNT_IN)));
    }

    function assertEq(ISynapseIntentRouter.StepParams memory a, ISynapseIntentRouter.StepParams memory b) public pure {
        assertEq(a.token, b.token);
        assertEq(a.amount, b.amount);
        assertEq(a.msgValue, b.msgValue);
        assertEq(a.zapData, b.zapData);
    }

    // ════════════════════════════════════════════════ ZERO STEPS ═════════════════════════════════════════════════════

    function test_previewIntent_noOp_token() public view {
        (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps) =
            sip.previewIntent({swapQuoter: swapQuoterMock, tokenIn: tokenA, tokenOut: tokenA, amountIn: AMOUNT_IN});
        // Checks
        assertEq(amountOut, AMOUNT_IN);
        assertEq(steps.length, 0);
    }

    function test_previewIntent_noOp_native() public view {
        (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps) = sip.previewIntent({
            swapQuoter: swapQuoterMock,
            tokenIn: NATIVE_GAS_TOKEN,
            tokenOut: NATIVE_GAS_TOKEN,
            amountIn: AMOUNT_IN
        });
        // Checks
        assertEq(amountOut, AMOUNT_IN);
        assertEq(steps.length, 0);
    }

    function test_previewIntent_zeroAmountOut() public {
        // tokenOut is always populated
        SwapQuery memory emptyQuery;
        emptyQuery.tokenOut = tokenB;
        mockGetAmountOut({tokenIn: tokenA, tokenOut: tokenB, amountIn: AMOUNT_IN, mockQuery: emptyQuery});
        (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps) =
            sip.previewIntent({swapQuoter: swapQuoterMock, tokenIn: tokenA, tokenOut: tokenB, amountIn: AMOUNT_IN});
        // Checks
        assertEq(amountOut, 0);
        assertEq(steps.length, 0);
    }

    // ════════════════════════════════════════════════ SINGLE STEP ════════════════════════════════════════════════════

    function checkSingleStepIntent(
        address tokenIn,
        address tokenOut,
        uint256 expectedAmountOut,
        ISynapseIntentRouter.StepParams memory expectedStep
    )
        public
        view
    {
        // Preview intent
        (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps) =
            sip.previewIntent({swapQuoter: swapQuoterMock, tokenIn: tokenIn, tokenOut: tokenOut, amountIn: AMOUNT_IN});
        // Checks
        assertEq(amountOut, expectedAmountOut);
        assertEq(steps.length, 1);
        assertEq(steps[0], expectedStep);
    }

    function test_previewIntent_swap() public {
        SwapQuery memory mockQuery = getSwapQuery(tokenB);
        mockGetToken(TOKEN_IN_INDEX, tokenA);
        mockGetToken(TOKEN_OUT_INDEX, tokenB);
        mockGetAmountOut({tokenIn: tokenA, tokenOut: tokenB, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        ISynapseIntentRouter.StepParams memory expectedStep = ISynapseIntentRouter.StepParams({
            token: tokenA,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getSwapZapData()
        });
        checkSingleStepIntent(tokenA, tokenB, SWAP_AMOUNT_OUT, expectedStep);
    }

    function test_previewIntent_addLiquidity() public {
        SwapQuery memory mockQuery = getAddLiquidityQuery(lpToken);
        mockGetToken(TOKEN_IN_INDEX, tokenA);
        mockGetAmountOut({tokenIn: tokenA, tokenOut: lpToken, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        ISynapseIntentRouter.StepParams memory expectedStep = ISynapseIntentRouter.StepParams({
            token: tokenA,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getAddLiquidityZapData()
        });
        checkSingleStepIntent(tokenA, lpToken, SWAP_AMOUNT_OUT, expectedStep);
    }

    function test_previewIntent_removeLiquidity() public {
        SwapQuery memory mockQuery = getRemoveLiquidityQuery(tokenB);
        mockGetToken(TOKEN_OUT_INDEX, tokenB);
        mockGetAmountOut({tokenIn: lpToken, tokenOut: tokenB, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        ISynapseIntentRouter.StepParams memory expectedStep = ISynapseIntentRouter.StepParams({
            token: lpToken,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getRemoveLiquidityZapData()
        });
        checkSingleStepIntent(lpToken, tokenB, SWAP_AMOUNT_OUT, expectedStep);
    }

    function test_previewIntent_wrapETH() public {
        SwapQuery memory mockQuery = getWrapETHQuery(weth);
        mockGetAmountOut({tokenIn: NATIVE_GAS_TOKEN, tokenOut: weth, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        ISynapseIntentRouter.StepParams memory expectedStep = ISynapseIntentRouter.StepParams({
            token: NATIVE_GAS_TOKEN,
            amount: FULL_AMOUNT,
            msgValue: AMOUNT_IN,
            zapData: getWrapETHZapData()
        });
        checkSingleStepIntent(NATIVE_GAS_TOKEN, weth, AMOUNT_IN, expectedStep);
    }

    function test_previewIntent_unwrapETH() public {
        SwapQuery memory mockQuery = getUnwrapWETHQuery(NATIVE_GAS_TOKEN);
        mockGetAmountOut({tokenIn: weth, tokenOut: NATIVE_GAS_TOKEN, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        ISynapseIntentRouter.StepParams memory expectedStep = ISynapseIntentRouter.StepParams({
            token: weth,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getUnwrapWETHZapData()
        });
        checkSingleStepIntent(weth, NATIVE_GAS_TOKEN, AMOUNT_IN, expectedStep);
    }

    // ════════════════════════════════════════════════ DOUBLE STEP ════════════════════════════════════════════════════

    function checkDoubleStepIntent(
        address tokenIn,
        address tokenOut,
        uint256 expectedAmountOut,
        ISynapseIntentRouter.StepParams memory expectedStep0,
        ISynapseIntentRouter.StepParams memory expectedStep1
    )
        public
        view
    {
        // Preview intent
        (uint256 amountOut, ISynapseIntentRouter.StepParams[] memory steps) =
            sip.previewIntent({swapQuoter: swapQuoterMock, tokenIn: tokenIn, tokenOut: tokenOut, amountIn: AMOUNT_IN});
        // Checks
        assertEq(amountOut, expectedAmountOut);
        assertEq(steps.length, 2);
        assertEq(steps[0], expectedStep0);
        assertEq(steps[1], expectedStep1);
    }

    function test_previewIntent_swap_unwrapWETH() public {
        SwapQuery memory mockQuery = getSwapQuery(weth);
        mockGetToken(TOKEN_IN_INDEX, tokenA);
        mockGetToken(TOKEN_OUT_INDEX, weth);
        mockGetAmountOut({tokenIn: tokenA, tokenOut: NATIVE_GAS_TOKEN, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        // step0: tokenA -> weth
        ISynapseIntentRouter.StepParams memory expectedStep0 = ISynapseIntentRouter.StepParams({
            token: tokenA,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getSwapZapData()
        });
        // step1: weth -> NATIVE_GAS_TOKEN
        ISynapseIntentRouter.StepParams memory expectedStep1 = ISynapseIntentRouter.StepParams({
            token: weth,
            amount: FULL_AMOUNT,
            msgValue: 0,
            zapData: getUnwrapWETHZapData()
        });
        checkDoubleStepIntent(tokenA, NATIVE_GAS_TOKEN, SWAP_AMOUNT_OUT, expectedStep0, expectedStep1);
    }

    function test_previewIntent_wrapETH_swap() public {
        SwapQuery memory mockQuery = getSwapQuery(tokenB);
        mockGetToken(TOKEN_IN_INDEX, weth);
        mockGetToken(TOKEN_OUT_INDEX, tokenB);
        mockGetAmountOut({tokenIn: NATIVE_GAS_TOKEN, tokenOut: tokenB, amountIn: AMOUNT_IN, mockQuery: mockQuery});
        // step0: NATIVE_GAS_TOKEN -> weth
        ISynapseIntentRouter.StepParams memory expectedStep0 = ISynapseIntentRouter.StepParams({
            token: NATIVE_GAS_TOKEN,
            amount: FULL_AMOUNT,
            msgValue: AMOUNT_IN,
            zapData: getWrapETHZapData()
        });
        // step1: weth -> tokenB
        ISynapseIntentRouter.StepParams memory expectedStep1 =
            ISynapseIntentRouter.StepParams({token: weth, amount: FULL_AMOUNT, msgValue: 0, zapData: getSwapZapData()});
        checkDoubleStepIntent(NATIVE_GAS_TOKEN, tokenB, SWAP_AMOUNT_OUT, expectedStep0, expectedStep1);
    }
}
