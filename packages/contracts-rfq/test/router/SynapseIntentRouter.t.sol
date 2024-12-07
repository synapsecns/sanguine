// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {
    ISynapseIntentRouter,
    ISynapseIntentRouterErrors,
    SynapseIntentRouter
} from "../../contracts/router/SynapseIntentRouter.sol";
import {TokenZapV1} from "../../contracts/zaps/TokenZapV1.sol";

import {MockERC20} from "../mocks/MockERC20.sol";
import {PoolMock} from "../mocks/PoolMock.sol";
import {SimpleVaultMock} from "../mocks/SimpleVaultMock.sol";
import {WETHMock} from "../mocks/WETHMock.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseIntentRouterTest is Test, ISynapseIntentRouterErrors {
    address internal constant NATIVE_GAS_TOKEN = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 internal constant AMOUNT = 1 ether;
    uint256 internal constant EXTRA_FUNDS = 0.1337 ether;
    uint256 internal constant TOKEN_PRICE = 2; // in ETH
    uint256 internal constant FULL_BALANCE = type(uint256).max;

    SynapseIntentRouter internal router;
    TokenZapV1 internal tokenZap;

    MockERC20 internal erc20;
    WETHMock internal weth;
    PoolMock internal pool;
    SimpleVaultMock internal vault;

    address internal user;

    modifier withExtraFunds() {
        erc20.mint(address(tokenZap), EXTRA_FUNDS);
        weth.mint(address(tokenZap), EXTRA_FUNDS);
        deal(address(tokenZap), EXTRA_FUNDS);
        _;
    }

    function setUp() public {
        router = new SynapseIntentRouter();
        tokenZap = new TokenZapV1();

        erc20 = new MockERC20("TKN", 18);
        weth = new WETHMock();
        vault = new SimpleVaultMock();

        pool = new PoolMock(address(weth), address(erc20));
        pool.setRatioWei(TOKEN_PRICE * 1e18);

        user = makeAddr("User");

        // Deal funds to the user
        erc20.mint(user, 10 * AMOUNT);
        weth.mint(user, 10 * AMOUNT);
        deal(user, 10 * AMOUNT);

        // Deal funds to the pool
        erc20.mint(address(pool), 1000 * AMOUNT);
        weth.mint(address(pool), 1000 * AMOUNT);
        deal(address(pool), 1000 * AMOUNT);

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
            amountPosition: type(uint256).max,
            finalToken: address(weth),
            forwardTo: address(0)
        });
    }

    function getUnwrapZapData(address forwardTo) public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(weth),
            payload: abi.encodeCall(weth.withdraw, (AMOUNT)),
            // Amount is encoded as the first parameter
            amountPosition: 4,
            finalToken: NATIVE_GAS_TOKEN,
            forwardTo: forwardTo
        });
    }

    function getSwapZapData(address token, address forwardTo) public view returns (bytes memory) {
        address otherToken = token == address(weth) ? address(erc20) : address(weth);
        return tokenZap.encodeZapData({
            target: address(pool),
            // Use placeholder zero amount
            payload: abi.encodeCall(pool.swap, (0, token)),
            // Amount is encoded as the first parameter
            amountPosition: 4,
            finalToken: otherToken,
            forwardTo: forwardTo
        });
    }

    function getDepositZapData(address token) public view returns (bytes memory) {
        return tokenZap.encodeZapData({
            target: address(vault),
            // Use placeholder zero amount
            payload: abi.encodeCall(vault.deposit, (token, 0, user)),
            // Amount is encoded as the second parameter
            amountPosition: 4 + 32,
            finalToken: address(0),
            forwardTo: address(0)
        });
    }

    function completeUserIntent(
        uint256 msgValue,
        uint256 amountIn,
        uint256 minLastStepAmountIn,
        uint256 deadline,
        ISynapseIntentRouter.StepParams[] memory steps
    )
        public
        virtual
    {
        vm.prank(user);
        router.completeIntent{value: msgValue}({
            zapRecipient: address(tokenZap),
            amountIn: amountIn,
            minLastStepAmountIn: minLastStepAmountIn,
            deadline: deadline,
            steps: steps
        });
    }

    function checkRevertMsgValueAboveExpectedWithERC20(
        ISynapseIntentRouter.StepParams[] memory steps,
        uint256 lastStepAmountIn
    )
        public
    {
        vm.expectRevert(SIR__MsgValueIncorrect.selector);
        completeUserIntent({
            msgValue: 1,
            amountIn: AMOUNT,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp,
            steps: steps
        });
    }

    function checkRevertsMsgValueAboveExpectedWithNative(
        ISynapseIntentRouter.StepParams[] memory steps,
        uint256 lastStepAmountIn
    )
        public
    {
        // Just msg.value is too high
        vm.expectRevert(SIR__MsgValueIncorrect.selector);
        completeUserIntent({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp,
            steps: steps
        });
        // Both msg.value and amountIn are too high
        vm.expectRevert(SIR__MsgValueIncorrect.selector);
        completeUserIntent({
            msgValue: AMOUNT + 1,
            amountIn: AMOUNT + 1,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp,
            steps: steps
        });
    }

    function checkRevertsMsgValueBelowExpectedWithNative(
        ISynapseIntentRouter.StepParams[] memory steps,
        uint256 lastStepAmountIn
    )
        public
    {
        // Just msg.value is too low
        vm.expectRevert(SIR__MsgValueIncorrect.selector);
        completeUserIntent({
            msgValue: AMOUNT - 1,
            amountIn: AMOUNT,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp,
            steps: steps
        });
        // Both msg.value and amountIn are too low
        vm.expectRevert(abi.encodeWithSelector(Address.AddressInsufficientBalance.selector, router));
        completeUserIntent({
            msgValue: AMOUNT - 1,
            amountIn: AMOUNT - 1,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp,
            steps: steps
        });
    }

    function checkRevertDeadlineExceeded(
        uint256 msgValue,
        uint256 lastStepAmountIn,
        ISynapseIntentRouter.StepParams[] memory steps
    )
        public
    {
        vm.expectRevert(SIR__DeadlineExceeded.selector);
        completeUserIntent({
            msgValue: msgValue,
            amountIn: AMOUNT,
            minLastStepAmountIn: lastStepAmountIn,
            deadline: block.timestamp - 1,
            steps: steps
        });
    }

    function checkRevertAmountInsufficient(
        uint256 msgValue,
        uint256 lastStepAmountIn,
        ISynapseIntentRouter.StepParams[] memory steps
    )
        public
    {
        vm.expectRevert(SIR__AmountInsufficient.selector);
        completeUserIntent({
            msgValue: msgValue,
            amountIn: AMOUNT,
            minLastStepAmountIn: lastStepAmountIn + 1,
            deadline: block.timestamp,
            steps: steps
        });
    }

    // ═══════════════════════════════════════════════ DEPOSIT ERC20 ═══════════════════════════════════════════════════

    function getDepositERC20Steps(uint256 amount) public view returns (ISynapseIntentRouter.StepParams[] memory) {
        return toArray(
            ISynapseIntentRouter.StepParams({
                token: address(erc20),
                amount: amount,
                msgValue: 0,
                zapData: getDepositZapData(address(erc20))
            })
        );
    }

    function test_depositERC20_exactAmount() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(AMOUNT);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_depositERC20_exactAmount_extraFunds() public withExtraFunds {
        test_depositERC20_exactAmount();
    }

    function test_depositERC20_exactAmount_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(AMOUNT);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_depositERC20_exactAmount_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(AMOUNT);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_depositERC20_exactAmount_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_depositERC20_fullBalance() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT);
    }

    /// @notice Extra funds should be used with "full balance" instructions.
    function test_depositERC20_fullBalance_extraFunds() public withExtraFunds {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(erc20)), AMOUNT + EXTRA_FUNDS);
    }

    function test_depositERC20_fullBalance_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_depositERC20_fullBalance_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_depositERC20_fullBalance_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositERC20Steps(FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    // ══════════════════════════════════════════════ DEPOSIT NATIVE ═══════════════════════════════════════════════════

    function getDepositNativeSteps(uint256 amount) public view returns (ISynapseIntentRouter.StepParams[] memory) {
        return toArray(
            ISynapseIntentRouter.StepParams({
                token: NATIVE_GAS_TOKEN,
                amount: amount,
                msgValue: AMOUNT,
                zapData: getDepositZapData(NATIVE_GAS_TOKEN)
            })
        );
    }

    function test_depositNative_exactAmount() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_depositNative_exactAmount_extraFunds() public withExtraFunds {
        test_depositNative_exactAmount();
    }

    function test_depositNative_exactAmount_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_depositNative_exactAmount_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_depositNative_exactAmount_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_depositNative_exactAmount_revert_msgValueBelowExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(AMOUNT);
        checkRevertsMsgValueBelowExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT - 1});
    }

    function test_depositNative_fullBalance() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should be used with "full balance" instructions.
    function test_depositNative_fullBalance_extraFunds() public withExtraFunds {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT + EXTRA_FUNDS);
    }

    function test_depositNative_fullBalance_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_depositNative_fullBalance_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_depositNative_fullBalance_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_depositNative_fullBalance_revert_msgValueBelowExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getDepositNativeSteps(FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT - 1});
    }

    // ═══════════════════════════════════════════ SWAP & DEPOSIT ERC20 ════════════════════════════════════════════════

    function getSwapDepositERC20Steps(
        uint256 amountSwap,
        uint256 amountDeposit
    )
        public
        view
        returns (ISynapseIntentRouter.StepParams[] memory)
    {
        return toArray(
            // WETH -> ERC20
            ISynapseIntentRouter.StepParams({
                token: address(weth),
                amount: amountSwap,
                msgValue: 0,
                zapData: getSwapZapData(address(weth), address(0))
            }),
            // deposit ERC20
            ISynapseIntentRouter.StepParams({
                token: address(erc20),
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(address(erc20))
            })
        );
    }

    function test_swapDepositERC20_exactAmounts() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, amountDeposit);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_swapDepositERC20_exactAmounts_extraFunds() public withExtraFunds {
        test_swapDepositERC20_exactAmounts();
    }

    function test_swapDepositERC20_exactAmounts_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, amountDeposit);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: amountDeposit, steps: steps});
    }

    function test_swapDepositERC20_exactAmounts_revert_lastStepAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, amountDeposit);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: amountDeposit + 1, steps: steps});
    }

    function test_swapDepositERC20_exactAmounts_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, amountDeposit);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: amountDeposit});
    }

    function test_swapDepositERC20_exactAmount0() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Extra funds should be used with the final "full balance" instruction.
    function test_swapDepositERC20_exactAmount0_extraFunds() public withExtraFunds {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE + EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_exactAmount0_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: amountDeposit, steps: steps});
    }

    function test_swapDepositERC20_exactAmount0_revert_lastStepAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: amountDeposit + 1, steps: steps});
    }

    function test_swapDepositERC20_exactAmount0_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(AMOUNT, FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: amountDeposit});
    }

    function test_swapDepositERC20_exactAmount1() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_swapDepositERC20_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual withExtraFunds {
        test_swapDepositERC20_exactAmount1();
    }

    function test_swapDepositERC20_exactAmount1_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: amountDeposit, steps: steps});
    }

    function test_swapDepositERC20_exactAmount1_revert_lastStepAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: amountDeposit + 1, steps: steps});
    }

    function test_swapDepositERC20_exactAmount1_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, amountDeposit);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: amountDeposit});
    }

    function test_swapDepositERC20_fullBalances() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    /// @notice Extra funds should be used with both "full balance" instructions.
    function test_swapDepositERC20_fullBalances_extraFunds() public withExtraFunds {
        uint256 amountDeposit = (AMOUNT + EXTRA_FUNDS) * TOKEN_PRICE + EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(erc20)), amountDeposit);
    }

    function test_swapDepositERC20_fullBalances_revert_deadlineExceeded() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: 0, lastStepAmountIn: amountDeposit, steps: steps});
    }

    function test_swapDepositERC20_fullBalances_revert_lastStepAmountInsufficient() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: 0, lastStepAmountIn: amountDeposit + 1, steps: steps});
    }

    function test_swapDepositERC20_fullBalances_revert_msgValueAboveExpected() public {
        uint256 amountDeposit = AMOUNT * TOKEN_PRICE;
        ISynapseIntentRouter.StepParams[] memory steps = getSwapDepositERC20Steps(FULL_BALANCE, FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: amountDeposit});
    }

    // ════════════════════════════════════════════ WRAP & DEPOSIT WETH ════════════════════════════════════════════════

    function getWrapDepositWETHSteps(
        uint256 amountWrap,
        uint256 amountDeposit
    )
        public
        view
        returns (ISynapseIntentRouter.StepParams[] memory)
    {
        return toArray(
            // ETH -> WETH
            ISynapseIntentRouter.StepParams({
                token: NATIVE_GAS_TOKEN,
                amount: amountWrap,
                msgValue: AMOUNT,
                zapData: getWrapZapData()
            }),
            // deposit WETH
            ISynapseIntentRouter.StepParams({
                token: address(weth),
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(address(weth))
            })
        );
    }

    function test_wrapDepositWETH_exactAmounts() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, AMOUNT);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_wrapDepositWETH_exactAmounts_extraFunds() public withExtraFunds {
        test_wrapDepositWETH_exactAmounts();
    }

    function test_wrapDepositWETH_exactAmounts_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_wrapDepositWETH_exactAmounts_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_wrapDepositWETH_exactAmounts_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, AMOUNT);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_wrapDepositWETH_exactAmount0() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Extra funds should be used with the final "full balance" instruction.
    function test_wrapDepositWETH_exactAmount0_extraFunds() public withExtraFunds {
        uint256 amountDeposit = AMOUNT + EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(weth)), amountDeposit);
    }

    function test_wrapDepositWETH_exactAmount0_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_wrapDepositWETH_exactAmount0_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_wrapDepositWETH_exactAmount0_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_wrapDepositWETH_exactAmount0_revert_msgValueBelowExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(AMOUNT, FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT - 1});
    }

    function test_wrapDepositWETH_exactAmount1() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_wrapDepositWETH_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual withExtraFunds {
        test_wrapDepositWETH_exactAmount1();
    }

    function test_wrapDepositWETH_exactAmount1_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_wrapDepositWETH_exactAmount1_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_wrapDepositWETH_exactAmount1_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_wrapDepositWETH_exactAmount1_revert_msgValueBelowExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, AMOUNT);
        checkRevertsMsgValueBelowExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT - 1});
    }

    function test_wrapDepositWETH_fullBalances() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, address(weth)), AMOUNT);
    }

    /// @notice Extra funds should be used with both "full balance" instructions.
    function test_wrapDepositWETH_fullBalances_extraFunds() public withExtraFunds {
        uint256 amountDeposit = AMOUNT + 2 * EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: AMOUNT,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, address(weth)), amountDeposit);
    }

    function test_wrapDepositWETH_fullBalances_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_wrapDepositWETH_fullBalances_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_wrapDepositWETH_fullBalances_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertsMsgValueAboveExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_wrapDepositWETH_fullBalances_revert_msgValueBelowExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getWrapDepositWETHSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertsMsgValueBelowExpectedWithNative({steps: steps, lastStepAmountIn: AMOUNT - 1});
    }

    // ══════════════════════════════════════════ UNWRAP & DEPOSIT NATIVE ══════════════════════════════════════════════

    function getUnwrapDepositNativeSteps(
        uint256 amountUnwrap,
        uint256 amountDeposit
    )
        public
        view
        returns (ISynapseIntentRouter.StepParams[] memory)
    {
        return toArray(
            // WETH -> ETH
            ISynapseIntentRouter.StepParams({
                token: address(weth),
                amount: amountUnwrap,
                msgValue: 0,
                zapData: getUnwrapZapData(address(0))
            }),
            // Deposit ETH
            ISynapseIntentRouter.StepParams({
                token: NATIVE_GAS_TOKEN,
                amount: amountDeposit,
                msgValue: 0,
                zapData: getDepositZapData(NATIVE_GAS_TOKEN)
            })
        );
    }

    function test_unwrapDepositNative_exactAmounts() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, AMOUNT);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should have no effect on "exact amount" instructions.
    function test_unwrapDepositNative_exactAmounts_extraFunds() public withExtraFunds {
        test_unwrapDepositNative_exactAmounts();
    }

    function test_unwrapDepositNative_exactAmounts_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_unwrapDepositNative_exactAmounts_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_unwrapDepositNative_exactAmounts_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_exactAmount0() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should be used with the final "full balance" instruction.
    function test_unwrapDepositNative_exactAmount0_extraFunds() public withExtraFunds {
        uint256 amountDeposit = AMOUNT + EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), amountDeposit);
    }

    function test_unwrapDepositNative_exactAmount0_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_unwrapDepositNative_exactAmount0_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_unwrapDepositNative_exactAmount0_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(AMOUNT, FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_exactAmount1() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, AMOUNT);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Should succeed with extra funds if no balance checks are performed.
    /// Last action is "use exact amount", so extra funds have no effect.
    function test_unwrapDepositNative_exactAmount1_extraFunds_revertWithBalanceChecks() public virtual withExtraFunds {
        test_unwrapDepositNative_exactAmount1();
    }

    function test_unwrapDepositNative_exactAmount1_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, AMOUNT);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_unwrapDepositNative_exactAmount1_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, AMOUNT);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_unwrapDepositNative_exactAmount1_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, AMOUNT);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    function test_unwrapDepositNative_fullBalances() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: AMOUNT,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), AMOUNT);
    }

    /// @notice Extra funds should be used with both "full balance" instructions.
    function test_unwrapDepositNative_fullBalances_extraFunds() public withExtraFunds {
        uint256 amountDeposit = AMOUNT + 2 * EXTRA_FUNDS;
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, FULL_BALANCE);
        completeUserIntent({
            msgValue: 0,
            amountIn: AMOUNT,
            minLastStepAmountIn: amountDeposit,
            deadline: block.timestamp,
            steps: steps
        });
        // Check that the vault registered the deposit with the extra funds
        assertEq(vault.balanceOf(user, NATIVE_GAS_TOKEN), amountDeposit);
    }

    function test_unwrapDepositNative_fullBalances_revert_deadlineExceeded() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertDeadlineExceeded({msgValue: AMOUNT, lastStepAmountIn: AMOUNT, steps: steps});
    }

    function test_unwrapDepositNative_fullBalances_revert_lastStepAmountInsufficient() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertAmountInsufficient({msgValue: AMOUNT, lastStepAmountIn: AMOUNT + 1, steps: steps});
    }

    function test_unwrapDepositNative_fullBalances_revert_msgValueAboveExpected() public {
        ISynapseIntentRouter.StepParams[] memory steps = getUnwrapDepositNativeSteps(FULL_BALANCE, FULL_BALANCE);
        checkRevertMsgValueAboveExpectedWithERC20({steps: steps, lastStepAmountIn: AMOUNT});
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    function toArray(ISynapseIntentRouter.StepParams memory a)
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory arr)
    {
        arr = new ISynapseIntentRouter.StepParams[](1);
        arr[0] = a;
        return arr;
    }

    function toArray(
        ISynapseIntentRouter.StepParams memory a,
        ISynapseIntentRouter.StepParams memory b
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory arr)
    {
        arr = new ISynapseIntentRouter.StepParams[](2);
        arr[0] = a;
        arr[1] = b;
        return arr;
    }

    function toArray(
        ISynapseIntentRouter.StepParams memory a,
        ISynapseIntentRouter.StepParams memory b,
        ISynapseIntentRouter.StepParams memory c
    )
        internal
        pure
        returns (ISynapseIntentRouter.StepParams[] memory arr)
    {
        arr = new ISynapseIntentRouter.StepParams[](3);
        arr[0] = a;
        arr[1] = b;
        arr[2] = c;
        return arr;
    }
}
