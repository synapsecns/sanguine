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
        pool = new PoolMock(address(erc20), address(weth));
        vault = new SimpleVaultMock();

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

    function getDepositZapData(address token) public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(vault),
            // Use placeholder zero amount
            payload: abi.encodeCall(vault.deposit, (token, 0, user)),
            // Amount is encoded as the second parameter
            amountPosition: 4 + 32
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
        vm.expectRevert(ZapRouterV1__DeadlineExceeded.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp - 1,
            zapParams: zapParams
        });
    }

    function test_depositERC20_exactAmount_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        vm.expectRevert(ZapRouterV1__AmountInsufficient.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT + 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositERC20_exactAmount_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(AMOUNT);
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: 1,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
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
        vm.expectRevert(ZapRouterV1__DeadlineExceeded.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp - 1,
            zapParams: zapParams
        });
    }

    function test_depositERC20_fullBalance_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        vm.expectRevert(ZapRouterV1__AmountInsufficient.selector);
        userPerformZaps({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT + 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositERC20_fullBalance_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositERC20ZapParams(USE_FULL_BALANCE);
        vm.expectRevert(ZapRouterV1__MsgValueIncorrect.selector);
        userPerformZaps({
            msgValue: 1,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp,
            zapParams: zapParams
        });
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
        vm.expectRevert(ZapRouterV1__DeadlineExceeded.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp - 1,
            zapParams: zapParams
        });
    }

    function test_depositNative_exactAmount_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
        vm.expectRevert(ZapRouterV1__AmountInsufficient.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT + 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositNative_exactAmount_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
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

    function test_depositNative_exactAmount_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(AMOUNT);
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
        vm.expectRevert(ZapRouterV1__DeadlineExceeded.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT,
            deadline: block.timestamp - 1,
            zapParams: zapParams
        });
    }

    function test_depositNative_fullBalance_revert_lastZapAmountInsufficient() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
        vm.expectRevert(ZapRouterV1__AmountInsufficient.selector);
        userPerformZaps({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastZapAmountIn: AMOUNT + 1,
            deadline: block.timestamp,
            zapParams: zapParams
        });
    }

    function test_depositNative_fullBalance_revert_msgValueAboveExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
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

    function test_depositNative_fullBalance_revert_msgValueBelowExpected() public {
        IZapRouterV1.ZapParams[] memory zapParams = getDepositNativeZapParams(USE_FULL_BALANCE);
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
