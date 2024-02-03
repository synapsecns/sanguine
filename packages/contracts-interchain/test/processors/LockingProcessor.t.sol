// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {LockingProcessor} from "../../src/processors/LockingProcessor.sol";

import {MockERC20} from "../mocks/MockERC20.sol";
import {MockInterchainERC20} from "../mocks/MockInterchainERC20.sol";

import {AbstractProcessorTest} from "./AbstractProcessor.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract LockingProcessorTest is AbstractProcessorTest {
    function deployTokens() internal virtual override {
        token = new MockERC20("Token");
        icToken = new MockInterchainERC20("IC Token");
    }

    function deployProcessor() internal virtual override {
        processor = new LockingProcessor(address(icToken), address(token));
        // Mint the underlying token to the processor (backing the user IC token balance)
        token.mintPublic(address(processor), START_BALANCE);
    }

    // Lock token: token (1) -> icToken (0)
    function test_swap_lockUnderlyingToken() public {
        uint256 amount = 100;
        vm.prank(user);
        processor.swap(1, 0, amount, 0, type(uint256).max);
        // Check underlying token balances
        assertEq(token.balanceOf(user), START_BALANCE - amount);
        assertEq(token.balanceOf(address(processor)), START_BALANCE + amount);
        // Check IC token balance
        assertEq(icToken.balanceOf(user), START_BALANCE + amount);
        assertEq(icToken.balanceOf(address(processor)), 0);
    }

    // Unlock token: icToken (0) -> token (1)
    function test_swap_unlockUnderlyingToken() public {
        uint256 amount = 100;
        vm.prank(user);
        processor.swap(0, 1, amount, 0, type(uint256).max);
        // Check underlying token balance
        assertEq(token.balanceOf(user), START_BALANCE + amount);
        assertEq(token.balanceOf(address(processor)), START_BALANCE - amount);
        // Check IC token balance
        assertEq(icToken.balanceOf(user), START_BALANCE - amount);
        assertEq(icToken.balanceOf(address(processor)), 0);
    }
}
