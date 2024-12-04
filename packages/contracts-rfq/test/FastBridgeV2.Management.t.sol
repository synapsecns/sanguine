// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IAdmin} from "../contracts/interfaces/IAdmin.sol";
import {IAdminV2Errors} from "../contracts/interfaces/IAdminV2Errors.sol";

import {FastBridgeV2, FastBridgeV2Test} from "./FastBridgeV2.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2ManagementTest is FastBridgeV2Test, IAdminV2Errors {
    uint256 public constant FEE_RATE_MAX = 1e4; // 1%
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");

    uint256 public constant MIN_CANCEL_DELAY = 1 hours;
    uint256 public constant DEFAULT_CANCEL_DELAY = 1 days;

    address public admin = makeAddr("Admin");
    address public governorA = makeAddr("Governor A");

    address public proverA = makeAddr("Prover A");
    address public proverB = makeAddr("Prover B");

    event ProverAdded(address prover);
    event ProverRemoved(address prover);

    event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay);
    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

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

    function addProver(address caller, address prover) public {
        vm.prank(caller);
        fastBridge.addProver(prover);
    }

    function removeProver(address caller, address prover) public {
        vm.prank(caller);
        fastBridge.removeProver(prover);
    }

    function setGovernor(address caller, address newGovernor) public {
        vm.prank(caller);
        fastBridge.grantRole(GOVERNOR_ROLE, newGovernor);
    }

    function setCancelDelay(address caller, uint256 newCancelDelay) public {
        vm.prank(caller);
        fastBridge.setCancelDelay(newCancelDelay);
    }

    function setProtocolFeeRate(address caller, uint256 newFeeRate) public {
        vm.prank(caller);
        fastBridge.setProtocolFeeRate(newFeeRate);
    }

    function sweepProtocolFees(address caller, address token, address recipient) public {
        vm.prank(caller);
        fastBridge.sweepProtocolFees(token, recipient);
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

    function test_defaultCancelDelay() public view {
        assertEq(fastBridge.cancelDelay(), DEFAULT_CANCEL_DELAY);
    }

    // ════════════════════════════════════════════════ ADD PROVER ═════════════════════════════════════════════════════

    function test_addProver() public {
        vm.expectEmit(address(fastBridge));
        emit ProverAdded(proverA);
        addProver(governor, proverA);
        assertEq(fastBridge.getActiveProverID(proverA), 1);
        assertEq(fastBridge.getActiveProverID(proverB), 0);
        address[] memory provers = fastBridge.getProvers();
        assertEq(provers.length, 1);
        assertEq(provers[0], proverA);
    }

    function test_addProver_twice() public {
        test_addProver();
        vm.expectEmit(address(fastBridge));
        emit ProverAdded(proverB);
        addProver(governor, proverB);
        assertEq(fastBridge.getActiveProverID(proverA), 1);
        assertEq(fastBridge.getActiveProverID(proverB), 2);
        address[] memory provers = fastBridge.getProvers();
        assertEq(provers.length, 2);
        assertEq(provers[0], proverA);
        assertEq(provers[1], proverB);
    }

    function test_addProver_twice_afterRemoval() public {
        test_removeProver_twice();
        // Add B back
        vm.expectEmit(address(fastBridge));
        emit ProverAdded(proverB);
        addProver(governor, proverB);
        assertEq(fastBridge.getActiveProverID(proverA), 0);
        assertEq(fastBridge.getActiveProverID(proverB), 2);
        address[] memory provers = fastBridge.getProvers();
        assertEq(provers.length, 1);
        assertEq(provers[0], proverB);
        // Add A back
        vm.expectEmit(address(fastBridge));
        emit ProverAdded(proverA);
        addProver(governor, proverA);
        assertEq(fastBridge.getActiveProverID(proverA), 1);
        assertEq(fastBridge.getActiveProverID(proverB), 2);
        provers = fastBridge.getProvers();
        assertEq(provers.length, 2);
        assertEq(provers[0], proverA);
        assertEq(provers[1], proverB);
    }

    function test_addProver_revertNotGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        addProver(caller, proverA);
    }

    function test_addProver_revertAlreadyActive() public {
        test_addProver();
        vm.expectRevert(ProverAlreadyActive.selector);
        addProver(governor, proverA);
    }

    function test_addProver_revertTooManyProvers() public {
        for (uint256 i = 0; i < type(uint16).max; i++) {
            addProver(governor, address(uint160(i)));
        }
        vm.expectRevert(ProverCapacityExceeded.selector);
        addProver(governor, proverA);
    }

    // ═══════════════════════════════════════════════ REMOVE PROVER ═══════════════════════════════════════════════════

    function test_removeProver() public {
        test_addProver_twice();
        vm.expectEmit(address(fastBridge));
        emit ProverRemoved(proverA);
        removeProver(governor, proverA);
        assertEq(fastBridge.getActiveProverID(proverA), 0);
        assertEq(fastBridge.getActiveProverID(proverB), 2);
        address[] memory provers = fastBridge.getProvers();
        assertEq(provers.length, 1);
        assertEq(provers[0], proverB);
    }

    function test_removeProver_twice() public {
        test_removeProver();
        vm.expectEmit(address(fastBridge));
        emit ProverRemoved(proverB);
        removeProver(governor, proverB);
        assertEq(fastBridge.getActiveProverID(proverA), 0);
        assertEq(fastBridge.getActiveProverID(proverB), 0);
        address[] memory provers = fastBridge.getProvers();
        assertEq(provers.length, 0);
    }

    function test_removeProver_revertNotGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        removeProver(caller, proverA);
    }

    function test_removeProver_revertNeverBeenActive() public {
        vm.expectRevert(ProverNotActive.selector);
        removeProver(governor, proverA);
    }

    function test_removeProver_revertNotActive() public {
        test_removeProver();
        vm.expectRevert(ProverNotActive.selector);
        removeProver(governor, proverA);
    }

    // ═════════════════════════════════════════════ SET CANCEL DELAY ══════════════════════════════════════════════════

    function test_setCancelDelay() public {
        vm.expectEmit(address(fastBridge));
        emit CancelDelayUpdated(DEFAULT_CANCEL_DELAY, 4 days);
        setCancelDelay(governor, 4 days);
        assertEq(fastBridge.cancelDelay(), 4 days);
    }

    function test_setCancelDelay_twice() public {
        test_setCancelDelay();
        vm.expectEmit(address(fastBridge));
        emit CancelDelayUpdated(4 days, 8 days);
        setCancelDelay(governor, 8 days);
        assertEq(fastBridge.cancelDelay(), 8 days);
    }

    function test_setCancelDelay_revertBelowMin() public {
        vm.expectRevert(CancelDelayBelowMin.selector);
        setCancelDelay(governor, MIN_CANCEL_DELAY - 1);
    }

    function test_setCancelDelay_revertNotGovernor(address caller) public {
        vm.assume(caller != governor);
        expectUnauthorized(caller, fastBridge.GOVERNOR_ROLE());
        setCancelDelay(caller, 4 days);
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
        vm.expectRevert(FeeRateAboveMax.selector);
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

    function test_chainGasAmountZero() public view {
        assertEq(fastBridge.chainGasAmount(), 0);
    }

    function test_setChainGasAmount_revert() public {
        // Generic revert: this function should not be in the V2 interface
        vm.expectRevert();
        vm.prank(governor);
        IAdmin(address(fastBridge)).setChainGasAmount(123);
    }
}
