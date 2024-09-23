// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {FastBridgeV2, FastBridgeV2Test} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2ManagementTest is FastBridgeV2Test {
    uint256 public constant FEE_RATE_MAX = 1e4; // 1%
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    address public admin = makeAddr("Admin");
    address public governorA = makeAddr("Governor A");

    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);
    event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount);

    function deployFastBridge() public override returns (FastBridgeV2) {
        return new FastBridgeV2(admin);
    }

    function configureFastBridge() public override {
        setGovernor(admin, governor);
    }

    function mintTokens() public override {
        srcToken.mint(address(fastBridge), 100);
        deal(address(fastBridge), 200);
        cheatCollectedProtocolFees(address(srcToken), 100);
        cheatCollectedProtocolFees(ETH_ADDRESS, 200);
    }

    function setGovernor(address caller, address newGovernor) public {
        vm.prank(caller);
        fastBridge.grantRole(GOVERNOR_ROLE, newGovernor);
    }

    function setProtocolFeeRate(address caller, uint256 newFeeRate) public {
        vm.prank(caller);
        fastBridge.setProtocolFeeRate(newFeeRate);
    }

    function sweepProtocolFees(address caller, address token, address recipient) public {
        vm.prank(caller);
        fastBridge.sweepProtocolFees(token, recipient);
    }

    function setChainGasAmount(address caller, uint256 newChainGasAmount) public {
        vm.prank(caller);
        fastBridge.setChainGasAmount(newChainGasAmount);
    }

    function test_grantGovernorRole() public {
        assertFalse(fastBridge.hasRole(GOVERNOR_ROLE, governorA));
        setGovernor(admin, governorA);
        assertTrue(fastBridge.hasRole(GOVERNOR_ROLE, governorA));
    }

    function test_grantGovernorRole_revertNotAdmin(address caller) public {
        vm.assume(caller != admin);
        expectUnauthorized(caller, fastBridge.DEFAULT_ADMIN_ROLE());
        setGovernor(caller, governorA);
    }

    // ═══════════════════════════════════════════ SET PROTOCOL FEE RATE ═══════════════════════════════════════════════

    function test_setProtocolFeeRate() public {
        vm.expectEmit(address(fastBridge));
        emit FeeRateUpdated(0, 123);
        setProtocolFeeRate(governor, 123);
        assertEq(fastBridge.protocolFeeRate(), 123);
    }

    function test_setProtocolFeeRate_twice() public {
        test_setProtocolFeeRate();
        vm.expectEmit(address(fastBridge));
        emit FeeRateUpdated(123, FEE_RATE_MAX);
        setProtocolFeeRate(governor, FEE_RATE_MAX);
        assertEq(fastBridge.protocolFeeRate(), FEE_RATE_MAX);
    }

    function test_setProtocolFeeRate_revert_tooHigh() public {
        vm.expectRevert("newFeeRate > max");
        setProtocolFeeRate(governor, FEE_RATE_MAX + 1);
    }

    function test_setProtocolFeeRate_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        setProtocolFeeRate(caller, 123);
    }

    // ════════════════════════════════════════════ SWEEP PROTOCOL FEES ════════════════════════════════════════════════

    function test_sweepProtocolFees_erc20() public {
        vm.expectEmit(address(fastBridge));
        emit FeesSwept(address(srcToken), governorA, 100);
        sweepProtocolFees(governor, address(srcToken), governorA);
        assertEq(srcToken.balanceOf(address(fastBridge)), 0);
        assertEq(srcToken.balanceOf(governorA), 100);
        assertEq(fastBridge.protocolFees(address(srcToken)), 0);
    }

    function test_sweepProtocolFees_eth() public {
        vm.expectEmit(address(fastBridge));
        emit FeesSwept(ETH_ADDRESS, governorA, 200);
        sweepProtocolFees(governor, ETH_ADDRESS, governorA);
        assertEq(address(fastBridge).balance, 0);
        assertEq(governorA.balance, 200);
        assertEq(fastBridge.protocolFees(ETH_ADDRESS), 0);
    }

    function test_sweepProtocolFees_revertNotGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        sweepProtocolFees(caller, address(srcToken), governorA);
    }

    // ═══════════════════════════════════════════ SET CHAIN GAS AMOUNT ════════════════════════════════════════════════

    function test_setChainGasAmount() public {
        vm.expectEmit(address(fastBridge));
        emit ChainGasAmountUpdated(0, 123);
        setChainGasAmount(governor, 123);
        assertEq(fastBridge.chainGasAmount(), 123);
    }

    function test_setChainGasAmount_twice() public {
        test_setChainGasAmount();
        vm.expectEmit(address(fastBridge));
        emit ChainGasAmountUpdated(123, 456);
        setChainGasAmount(governor, 456);
        assertEq(fastBridge.chainGasAmount(), 456);
    }

    function test_setChainGasAmount_revertNotGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        setChainGasAmount(caller, 123);
    }
}
