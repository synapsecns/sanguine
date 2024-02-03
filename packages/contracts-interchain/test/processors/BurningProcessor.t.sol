// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {BurningProcessor} from "../../src/processors/BurningProcessor.sol";

import {MockGMX} from "../mocks/MockGMX.sol";
import {MockInterchainERC20} from "../mocks/MockInterchainERC20.sol";

import {AbstractProcessorTest} from "./AbstractProcessor.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract BurningProcessorTest is AbstractProcessorTest {
    function deployTokens() internal virtual override {
        token = new MockGMX("Token", address(this));
        icToken = new MockInterchainERC20("IC Token");
    }

    function deployProcessor() internal virtual override {
        processor = new BurningProcessor(address(icToken), address(token));
        // Processor should be able to mint the underlying GMX token
        MockGMX(address(token)).transferOwnership(address(processor));
    }

    // Burn token: token (1) -> icToken (0)
    function test_swap_burnUnderlyingToken() public {
        uint256 amount = 100;
        vm.prank(user);
        processor.swap(1, 0, amount, amount, block.timestamp);
        // Check underlying token balance
        assertEq(token.balanceOf(user), START_BALANCE - amount);
        assertEq(token.balanceOf(address(processor)), 0);
        // Check IC token balance
        assertEq(icToken.balanceOf(user), START_BALANCE + amount);
        assertEq(icToken.balanceOf(address(processor)), 0);
    }

    // Mint token: icToken (0) -> token (1)
    function test_swap_mintUnderlyingToken() public {
        uint256 amount = 100;
        vm.prank(user);
        processor.swap(0, 1, amount, amount, block.timestamp);
        // Check underlying token balance
        assertEq(token.balanceOf(user), START_BALANCE + amount);
        assertEq(token.balanceOf(address(processor)), 0);
        // Check IC token balance
        assertEq(icToken.balanceOf(user), START_BALANCE - amount);
        assertEq(icToken.balanceOf(address(processor)), 0);
    }
}
