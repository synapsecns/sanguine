// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainERC20Test} from "./InterchainERC20.t.sol";
import {RateLimiting} from "../../src/libs/RateLimit.sol";

import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainERC20MintTest is InterchainERC20Test {
    address public bridge;
    address public user;

    uint256 public constant INITIAL_TOTAL_LIMIT = 1000;
    uint256 public constant INITIAL_MINTED = 400;
    uint256 public constant INITIAL_CURRENT_LIMIT = INITIAL_TOTAL_LIMIT - INITIAL_MINTED;

    uint256 public constant SMALL_PERIOD = 1 days / 10;
    uint256 public constant SMALL_PERIOD_CURRENT_LIMIT = INITIAL_CURRENT_LIMIT + 100;
    uint256 public constant LARGE_PERIOD = 1 days * 2;

    function setUp() public override {
        super.setUp();
        bridge = makeAddr("Bridge");
        user = makeAddr("User");
        // Set the total mint limit for the bridge and spend some of it
        authSetTotalMintLimit(bridge, INITIAL_TOTAL_LIMIT);
        authMintToken(address(1337), INITIAL_MINTED);
    }

    function authMintToken(address account, uint256 amount) public {
        vm.prank(bridge);
        token.mint(account, amount);
    }

    function test_mint_revert_unauthorized() public {
        vm.expectRevert(abi.encodeWithSelector(RateLimiting.RateLimiting__LimitExceeded.selector, 100, 0));
        vm.prank(user);
        token.mint(user, 100);
    }

    function test_mint_zeroTimePassed_mintUnderLimit() public {
        uint256 amount = INITIAL_CURRENT_LIMIT / 10;
        authMintToken(user, amount);
        assertEq(token.balanceOf(user), amount);
        assertEq(token.totalSupply(), INITIAL_MINTED + amount);
        assertEq(token.getCurrentMintLimit(bridge), INITIAL_CURRENT_LIMIT - amount);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_zeroTimePassed_mintExactlyLimit() public {
        authMintToken(user, INITIAL_CURRENT_LIMIT);
        assertEq(token.balanceOf(user), INITIAL_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_MINTED + INITIAL_CURRENT_LIMIT);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_zeroTimePassed_revert_mintOverLimit() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_CURRENT_LIMIT + 1, INITIAL_CURRENT_LIMIT
            )
        );
        authMintToken(user, INITIAL_CURRENT_LIMIT + 1);
    }

    function test_mint_timePassed_replenishUnderTotalLimit_mintUnderLimit() public {
        skip(SMALL_PERIOD);
        uint256 amount = SMALL_PERIOD_CURRENT_LIMIT / 10;
        authMintToken(user, amount);
        assertEq(token.balanceOf(user), amount);
        assertEq(token.totalSupply(), INITIAL_MINTED + amount);
        assertEq(token.getCurrentMintLimit(bridge), SMALL_PERIOD_CURRENT_LIMIT - amount);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_timePassed_replenishUnderTotalLimit_mintExactlyLimit() public {
        skip(SMALL_PERIOD);
        authMintToken(user, SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.balanceOf(user), SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.totalSupply(), INITIAL_MINTED + SMALL_PERIOD_CURRENT_LIMIT);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_timePassed_replenishUnderTotalLimit_revert_mintOverLimit() public {
        skip(SMALL_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector,
                SMALL_PERIOD_CURRENT_LIMIT + 1,
                SMALL_PERIOD_CURRENT_LIMIT
            )
        );
        authMintToken(user, SMALL_PERIOD_CURRENT_LIMIT + 1);
    }

    function test_mint_timePassed_replenishOverTotalLimit_mintUnderLimit() public {
        skip(LARGE_PERIOD);
        uint256 amount = INITIAL_TOTAL_LIMIT / 10;
        authMintToken(user, amount);
        assertEq(token.balanceOf(user), amount);
        assertEq(token.totalSupply(), INITIAL_MINTED + amount);
        assertEq(token.getCurrentMintLimit(bridge), INITIAL_TOTAL_LIMIT - amount);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_timePassed_replenishOverTotalLimit_mintExactlyLimit() public {
        skip(LARGE_PERIOD);
        authMintToken(user, INITIAL_TOTAL_LIMIT);
        assertEq(token.balanceOf(user), INITIAL_TOTAL_LIMIT);
        assertEq(token.totalSupply(), INITIAL_MINTED + INITIAL_TOTAL_LIMIT);
        assertEq(token.getCurrentMintLimit(bridge), 0);
        assertEq(token.getTotalMintLimit(bridge), INITIAL_TOTAL_LIMIT);
    }

    function test_mint_timePassed_replenishOverTotalLimit_revert_mintOverLimit() public {
        skip(LARGE_PERIOD);
        vm.expectRevert(
            abi.encodeWithSelector(
                RateLimiting.RateLimiting__LimitExceeded.selector, INITIAL_TOTAL_LIMIT + 1, INITIAL_TOTAL_LIMIT
            )
        );
        authMintToken(user, INITIAL_TOTAL_LIMIT + 1);
    }

    // ════════════════════════════════════════════ TESTS: MINT + PAUSE ════════════════════════════════════════════════

    function test_mint_revert_whenPaused() public {
        authPause();
        vm.expectRevert(Pausable.EnforcedPause.selector);
        authMintToken(user, 100);        
    }

    function test_mint_works_whenPausedAndUnpaused() public {
        authPause();
        authUnpause();
        test_mint_zeroTimePassed_mintUnderLimit();
    }
}
