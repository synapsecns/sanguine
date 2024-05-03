// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ClaimableFeesEvents} from "../../contracts/events/ClaimableFeesEvents.sol";
import {IClaimableFees} from "../../contracts/interfaces/IClaimableFees.sol";

import {ClaimableFeesHarness} from "../harnesses/ClaimableFeesHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract ClaimableFeesTest is ClaimableFeesEvents, Test {
    uint256 public constant INITIAL_BALANCE = 10 ether;
    uint256 public constant CLAIMABLE_AMOUNT = 5 ether;

    uint256 public constant ONE_PERCENT = 0.01e18;

    ClaimableFeesHarness public harness;

    address public claimer = makeAddr("Claimer");
    address public recipient = makeAddr("Recipient");

    event BeforeFeesClaimed(uint256 amount, uint256 reward);

    function setUp() public {
        harness = new ClaimableFeesHarness();
        deal(address(harness), INITIAL_BALANCE);
    }

    function expectEvents(uint256 claimerReward) public {
        vm.expectEmit(address(harness));
        emit BeforeFeesClaimed(CLAIMABLE_AMOUNT, claimerReward);
        vm.expectEmit(address(harness));
        emit FeesClaimed(recipient, CLAIMABLE_AMOUNT - claimerReward, claimer, claimerReward);
    }

    function test_claimFees_zeroFraction_events() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: 0});
        expectEvents({claimerReward: 0});
        vm.prank(claimer);
        harness.claimFees();
    }

    function test_claimFees_zeroFraction_balances() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: 0});
        vm.prank(claimer);
        harness.claimFees();
        assertEq(address(harness).balance, INITIAL_BALANCE - CLAIMABLE_AMOUNT);
        assertEq(claimer.balance, 0);
        assertEq(recipient.balance, CLAIMABLE_AMOUNT);
    }

    function test_claimFees_zeroFraction_state() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: 0});
        vm.prank(claimer);
        harness.claimFees();
        assertEq(harness.getClaimableAmount(), 0);
        assertEq(harness.getClaimerReward(), 0);
        assertEq(harness.getClaimerFraction(), 0);
        assertEq(harness.getFeeRecipient(), recipient);
    }

    function test_claimFees_nonZeroFraction_events() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: ONE_PERCENT});
        expectEvents({claimerReward: 5e16});
        vm.prank(claimer);
        harness.claimFees();
    }

    function test_claimFees_nonZeroFraction_balances() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: ONE_PERCENT});
        vm.prank(claimer);
        harness.claimFees();
        assertEq(address(harness).balance, INITIAL_BALANCE - CLAIMABLE_AMOUNT);
        assertEq(claimer.balance, 5e16);
        assertEq(recipient.balance, CLAIMABLE_AMOUNT - 5e16);
    }

    function test_claimFees_nonZeroFraction_state() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: ONE_PERCENT});
        vm.prank(claimer);
        harness.claimFees();
        assertEq(harness.getClaimableAmount(), 0);
        assertEq(harness.getClaimerReward(), 0);
        assertEq(harness.getClaimerFraction(), ONE_PERCENT);
        assertEq(harness.getFeeRecipient(), recipient);
    }

    function test_claimFees_revert_zeroAmount() public {
        harness.setup({claimableAmount: 0, feeRecipient: recipient, claimerFraction: ONE_PERCENT});
        vm.expectRevert(IClaimableFees.ClaimableFees__FeeAmountZero.selector);
        vm.prank(claimer);
        harness.claimFees();
    }

    function test_claimFees_revert_zeroFeeRecipient() public {
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: address(0), claimerFraction: ONE_PERCENT});
        vm.expectRevert(IClaimableFees.ClaimableFees__FeeRecipientZeroAddress.selector);
        vm.prank(claimer);
        harness.claimFees();
    }

    function test_claimFees_revert_claimerFractionExceedsOnePercent() public {
        uint256 tooBigFraction = ONE_PERCENT + 1;
        harness.setup({claimableAmount: CLAIMABLE_AMOUNT, feeRecipient: recipient, claimerFraction: tooBigFraction});
        vm.expectRevert(
            abi.encodeWithSelector(
                IClaimableFees.ClaimableFees__ClaimerFractionAboveMax.selector, tooBigFraction, ONE_PERCENT
            )
        );
        vm.prank(claimer);
        harness.claimFees();
    }
}
