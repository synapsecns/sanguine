// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IZapRouterV1, IZapRouterV1Errors, ZapRouterV1} from "../../contracts/router/ZapRouterV1.sol";
import {TokenZapV1} from "../../contracts/zaps/TokenZapV1.sol";

import {MockERC20} from "../MockERC20.sol";
import {PoolMock} from "../mocks/PoolMock.sol";
import {SimpleVaultMock} from "../mocks/SimpleVaultMock.sol";
import {WETHMock} from "../mocks/WETHMock.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract ZapRouterV1Test is Test, IZapRouterV1Errors {
    address internal constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 internal constant AMOUNT = 1 ether;
    uint256 internal constant TOKEN_PRICE = 2; // in ETH
    uint256 internal constant EXTRA_FUNDS = 0.1337 ether;
    uint256 internal constant USE_FULL_BALANCE = type(uint256).max;

    ZapRouterV1 internal router;
    TokenZapV1 internal tokenZap;

    MockERC20 internal erc20;
    WETHMock internal weth;
    PoolMock internal pool;
    SimpleVaultMock internal vault;

    address internal user;

    function setUp() public {
        router = new ZapRouterV1();
        tokenZap = new TokenZapV1();

        erc20 = new MockERC20("TKN", 18);
        weth = new WETHMock();
        vault = new SimpleVaultMock();

        pool = new PoolMock(address(weth), address(erc20));
        pool.setRatioWei(TOKEN_PRICE * 1e18);

        user = makeAddr("User");

        // Deal funds to the user
        erc20.mint(user, 10 * AMOUNT);
        deal(user, 20 * AMOUNT);
        vm.prank(user);
        weth.deposit{value: 10 * AMOUNT}();

        // Deal funds to the pool
        erc20.mint(address(pool), 1000 * AMOUNT);
        deal(address(pool), 2000 * AMOUNT);
        vm.prank(address(pool));
        weth.deposit{value: 1000 * AMOUNT}();

        // Approve the router
        vm.prank(user);
        erc20.approve(address(router), type(uint256).max);
        vm.prank(user);
        weth.approve(address(router), type(uint256).max);
    }

    function getWrapZapData() public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(weth),
            payload: abi.encodeCall(weth.deposit, ()),
            // Amount is not encoded
            amountPosition: type(uint256).max
        });
    }

    function getUnwrapZapData() public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(weth),
            payload: abi.encodeCall(weth.withdraw, (AMOUNT)),
            // Amount is encoded as the first parameter
            amountPosition: 4
        });
    }

    function getDepositZapData(address token) public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(vault),
            // Use placeholder zero amount
            payload: abi.encodeCall(vault.deposit, (token, 0, user)),
            // Amount is encoded as the second parameter
            amountPosition: 4 + 32
        });
    }

    function getSwapZapData(address token) public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(pool),
            // Use placeholder zero amount
            payload: abi.encodeCall(pool.swap, (0, token)),
            // Amount is encoded as the first parameter
            amountPosition: 4
        });
    }

    function userPerformZaps(
        uint256 msgValue,
        uint256 amountIn,
        uint256 minLastZapAmountIn,
        uint256 deadline,
        IZapRouterV1.ZapParams[] memory zapParams
    )
        public
        virtual
    {
        vm.prank(user);
        router.performZaps{value: msgValue}({
            zapRecipient: address(tokenZap),
            amountIn: amountIn,
            minLastZapAmountIn: minLastZapAmountIn,
            deadline: deadline,
            zapParams: zapParams
        });
    }

    function checkRevertMsgValueAboveExpectedWithERC20(
        IZapRouterV1.ZapParams[] memory zapParams,
        uint256 minLastZapAmountIn
    )
        public
    {
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: 1,
            amountIn: AMOUNT,
            minLastZapAmountIn: minLastZapAmountIn,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function checkRevertsMsgValueAboveExpectedWithNative(IZapRouterV1.ZapParams[] memory zapParams) public {
        // Just msg.value is too high
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Both msg.value and amountIn are too high
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT + 1,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function checkRevertsMsgValueBelowExpectedWithNative(IZapRouterV1.ZapParams[] memory zapParams) public {
        // Just msg.value is too low
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: AMOUNT - 1,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT - 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Both msg.value and amountIn are too low
        vm.expectRevert(abi.encodeWithSelector(Address.AddressInsufficientBalance.selector, router));
        userPerformZaps({
            msgValue: AMOUNT - 1,
            amountIn: AMOUNT - 1,
            minLastZapAmountIn: AMOUNT - 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function checkRevertDeadlineExceeded(
        uint256 msgValue,
        uint256 lastZapAmountIn,
        IZapRouterV1.ZapParams[] memory zapParams
    )
        public
    {
        vm.expectRevert(ZapRouterV1__DeadlineExceeded.selector);
        userPerformZaps({
            msgValue: msgValue,
            amountIn: AMOUNT,
            minLastZapAmountIn: lastZapAmountIn,
            deadline: block.timestamp - 1,
            zapParams: zapParams
        });
    }

    function checkRevertAmountInsufficient(
        uint256 msgValue,
        uint256 lastZapAmountIn,
        IZapRouterV1.ZapParams[] memory zapParams
    )
        public
    {
        vm.expectRevert(ZapRouterV1__AmountInsufficient.selector);
        userPerformZaps({
            msgValue: msgValue,
            amountIn: AMOUNT,
            minLastZapAmountIn: lastZapAmountIn + 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    // ═══════════════════════════════════════════════ DEPOSIT ERC20 ═══════════════════════════════════════════════════

    function getDepositERC20ZapParams(uint256 amount) public view returns (IZapRouterV1.ZapParams[] memory) {
        return toArray(
            IZapRouterV1.ZapParams({
                token: address(erc20),
                amount: amount,
                msgValue: 0,
                zapData: getDepositZapData(address(erc20))
            })
        );
    }

    function test_depositERC20_exactAmount() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_depositERC20_exactAmount_extraERC20() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        test_depositERC20_exactAmount();
    }

    function test_depositERC20_exactAmount_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositERC20_exactAmount_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositERC20_exactAmount_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, AMOUNT);
    }

    function test_depositERC20_fullBalance() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    /// @notice Extra funds should be used with "full balance" instructions.
    function test_depositERC20_fullBalance_extraERC20() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT + EXTRA_FUNDS);
    }

    function test_depositERC20_fullBalance_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositERC20_fullBalance_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositERC20_fullBalance_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, AMOUNT);
    }

    // ══════════════════════════════════════════════ DEPOSIT NATIVE ═══════════════════════════════════════════════════

    function getDepositNativeZapParams(uint256 amount) public view returns (IZapRouterV1.ZapParams[] memory) {
        return toArray(
            IZapRouterV1.ZapParams({
                token: NATIVE_GAS_TOKEN,
                amount: amount,
                msgValue: AMOUNT,
                zapData: getDepositZapData(NATIVE_GAS_TOKEN)
            })
        );
    }

    function test_depositNative_exactAmount() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_depositNative_exactAmount_extraNative() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        test_depositNative_exactAmount();
    }

    function test_depositNative_exactAmount_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositNative_exactAmount_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositNative_exactAmount_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_depositNative_exactAmount_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    function test_depositNative_fullBalance() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should be used with "full balance" instructions.
    function test_depositNative_fullBalance_extraNative() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT + EXTRA_FUNDS);
    }

    function test_depositNative_fullBalance_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositNative_fullBalance_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_depositNative_fullBalance_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_depositNative_fullBalance_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    // ═══════════════════════════════════════════ SWAP & DEPOSIT ERC20 ════════════════════════════════════════════════

    function getSwapDepositERC20ZapParams(
        uint256 amountSwap,
        uint256 amountDeposit
    )
        public
        view
        returns (IZapRouterV1.ZapParams[] memory)
    {
        return toArray(
            // WETH -> ERC20
            IZapRouterV1.ZapParams({
                token: address(weth),
                amount: amountSwap,
                msgValue: 0,
                zapData: getSwapZapData(address(weth))
            }),
            // deposit ERC20
            IZapRouterV1.ZapParams({
                token: address(erc20),
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(address(erc20))
            })
        );
    }

    function test_swapDepositERC20_exactAmounts() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, amountDeposit);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_swapDepositERC20_exactAmounts_extraFunds() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_swapDepositERC20_exactAmounts();
    }

    function test_swapDepositERC20_exactAmounts_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, amountDeposit);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmounts_revert_lastZapAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, amountDeposit);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmounts_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, amountDeposit);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, amountDeposit);
    }

    function test_swapDepositERC20_exactAmount0() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_exactAmount0_extraFunds() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE + EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_exactAmount0_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmount0_revert_lastZapAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmount0_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, amountDeposit);
    }

    function test_swapDepositERC20_exactAmount1() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_swapDepositERC20_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_swapDepositERC20_exactAmount1();
    }

    function test_swapDepositERC20_exactAmount1_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmount1_revert_lastZapAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_exactAmount1_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, amountDeposit);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, amountDeposit);
    }

    function test_swapDepositERC20_fullBalances() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_fullBalances_extraFunds() public {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = (AMOUNT + EXTRA_FUNDS) * TOKEN_PRICE + EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_fullBalances_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_fullBalances_revert_lastZapAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: amountDeposit, zapParams: zapParams});
    }

    function test_swapDepositERC20_fullBalances_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        IZapRouterV1.ZapParams[] memory zapParams = getSwapDepositERC20ZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20(zapParams, amountDeposit);
    }

    // ════════════════════════════════════════════ WRAP & DEPOSIT WETH ════════════════════════════════════════════════

    function getWrapDepositWETHZapParams(
        uint256 amountWrap,
        uint256 amountDeposit
    )
        public
        view
        returns (IZapRouterV1.ZapParams[] memory)
    {
        return toArray(
            IZapRouterV1.ZapParams({
                token: NATIVE_GAS_TOKEN,
                amount: amountWrap,
                msgValue: AMOUNT,
                zapData: getWrapZapData()
            }),
            IZapRouterV1.ZapParams({
                token: address(weth),
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(address(weth))
            })
        );
    }

    function test_wrapDepositWETH_exactAmounts() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_wrapDepositWETH_exactAmounts_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_wrapDepositWETH_exactAmounts();
    }

    function test_wrapDepositWETH_exactAmounts_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmounts_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmounts_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_exactAmounts_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, AMOUNT);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_exactAmount0() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    function test_wrapDepositWETH_exactAmount0_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT + EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), amountDeposit);
    }

    function test_wrapDepositWETH_exactAmount0_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmount0_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmount0_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_exactAmount0_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_exactAmount1() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, AMOUNT);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_wrapDepositWETH_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_wrapDepositWETH_exactAmount1();
    }

    function test_wrapDepositWETH_exactAmount1_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmount1_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_exactAmount1_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_exactAmount1_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_fullBalances() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    function test_wrapDepositWETH_fullBalances_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT + 2 * EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), amountDeposit);
    }

    function test_wrapDepositWETH_fullBalances_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_fullBalances_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_wrapDepositWETH_fullBalances_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative(zapParams);
    }

    function test_wrapDepositWETH_fullBalances_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getWrapDepositWETHZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative(zapParams);
    }

    // ══════════════════════════════════════════ UNWRAP & DEPOSIT NATIVE ══════════════════════════════════════════════

    function getUnwrapDepositNativeZapParams(
        uint256 amountUnwrap,
        uint256 amountDeposit
    )
        public
        view
        returns (IZapRouterV1.ZapParams[] memory)
    {
        return toArray(
            IZapRouterV1.ZapParams({
                token: address(weth),
                amount: amountUnwrap,
                msgValue: 0,
                zapData: getUnwrapZapData()
            }),
            IZapRouterV1.ZapParams({
                token: NATIVE_GAS_TOKEN,
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(NATIVE_GAS_TOKEN)
            })
        );
    }

    function test_unwrapDepositNative_exactAmounts() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, AMOUNT);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_unwrapDepositNative_exactAmounts_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_unwrapDepositNative_exactAmounts();
    }

    function test_unwrapDepositNative_exactAmounts_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmounts_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, AMOUNT);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmounts_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20({zapParams: zapParams, minLastZapAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_exactAmount0() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    function test_unwrapDepositNative_exactAmount0_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT + EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), amountDeposit);
    }

    function test_unwrapDepositNative_exactAmount0_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmount0_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmount0_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(AMOUNT, USE_FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({zapParams: zapParams, minLastZapAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_exactAmount1() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, AMOUNT);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_unwrapDepositNative_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        test_unwrapDepositNative_exactAmount1();
    }

    function test_unwrapDepositNative_exactAmount1_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmount1_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, AMOUNT);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_exactAmount1_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20({zapParams: zapParams, minLastZapAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_fullBalances() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    function test_unwrapDepositNative_fullBalances_extraFunds() public {
        deal(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        uint256 amountDeposit = AMOUNT + 2 * EXTRA_FUNDS;
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: amountDeposit,
            deadline: block.timestamp,
            zapParams: zapParams
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), amountDeposit);
    }

    function test_unwrapDepositNative_fullBalances_revert_deadlineExceeded() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_fullBalances_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastZapAmountIn: AMOUNT, zapParams: zapParams});
    }

    function test_unwrapDepositNative_fullBalances_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getUnwrapDepositNativeZapParams(USE_FULL_BALANCE, USE_FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({zapParams: zapParams, minLastZapAmountIn: AMOUNT});
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    function toArray(IZapRouterV1.ZapParams memory a) internal pure returns (IZapRouterV1.ZapParams[] memory arr) {
        arr = new IZapRouterV1.ZapParams[](1);
        arr[0] = a;
        return arr;
    }

    function toArray(
        IZapRouterV1.ZapParams memory a,
        IZapRouterV1.ZapParams memory b
    )
        internal
        pure
        returns (IZapRouterV1.ZapParams[] memory arr)
    {
        arr = new IZapRouterV1.ZapParams[](2);
        arr[0] = a;
        arr[1] = b;
        return arr;
    }

    function toArray(
        IZapRouterV1.ZapParams memory a,
        IZapRouterV1.ZapParams memory b,
        IZapRouterV1.ZapParams memory c
    )
        internal
        pure
        returns (IZapRouterV1.ZapParams[] memory arr)
    {
        arr = new IZapRouterV1.ZapParams[](3);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
        return arr;
    }
}
