// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFees, ExecutionFeesEvents, IExecutionFees} from "../contracts/ExecutionFees.sol";
import {Test} from "forge-std/Test.sol";

contract ExecutionFeesTest is ExecutionFeesEvents, Test {
    ExecutionFees executionFees;
    address icClient = address(0x123);
    address executor = address(0x456);
    address executorA = address(0x789);
    address public admin = makeAddr("Admin");
    bytes32 transactionId = keccak256("transaction");
    bytes32 transactionIdA = keccak256("transactionA");
    uint256 dstChainId = 1;
    uint256 executionFee = 1 ether;
    uint256 executionFeeA = 2 ether;

    function setUp() public {
        executionFees = new ExecutionFees(admin);
        bytes32 recorderRole = executionFees.RECORDER_ROLE();
        vm.prank(admin);
        executionFees.grantRole(recorderRole, icClient);
    }

    function addExecutionFee(uint256 fee, uint256 chainId, bytes32 txId) internal {
        vm.deal(address(this), fee);
        executionFees.addExecutionFee{value: fee}(chainId, txId);
    }

    function recordExecutor(uint256 chainId, bytes32 txId, address executorAddr) internal {
        vm.prank(address(icClient));
        executionFees.recordExecutor(chainId, txId, executorAddr);
    }

    function test_addExecutionFee_emitsEvent() public {
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee);
        addExecutionFee(executionFee, dstChainId, transactionId);
    }

    function test_addExecutionFee_doesNotAwardFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.accumulatedRewards(executor), 0);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_addExecutionFee_recordsTxFee() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee);
    }

    function test_addExecutionFee_addedTwice_sameTx_emitsEvent() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee + executionFeeA);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
    }

    function test_addExecutionFee_addedTwice_sameTx_recordsTxFee() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee + executionFeeA);
    }

    function test_addExecutionFee_addedTwice_diffTx_emitsEvent() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionIdA, executionFeeA);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
    }

    function test_addExecutionFee_addedTwice_diffTx_recordsTxFee() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee);
        assertEq(executionFees.executionFee(dstChainId, transactionIdA), executionFeeA);
    }

    function test_addExecutionFee_alreadyRecordedNoFee_emitsEvents() public {
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFee);
        addExecutionFee(executionFee, dstChainId, transactionId);
    }

    function test_addExecutionFee_alreadyRecordedNoFee_awardsFees() public {
        recordExecutor(dstChainId, transactionId, executor);
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.accumulatedRewards(executor), executionFee);
        assertEq(executionFees.unclaimedRewards(executor), executionFee);
    }

    function test_addExecutionFee_alreadyRecordedNoFee_recordsTxFee() public {
        recordExecutor(dstChainId, transactionId, executor);
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee);
    }

    function test_addExecutionFee_alreadyRecordedWithFee_emitsEvents() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee + executionFeeA);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFeeA);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
    }

    function test_addExecutionFee_alreadyRecordedWithFee_awardsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), executionFee + executionFeeA);
    }

    function test_addExecutionFee_alreadyRecordedWithFee_recordsTxFee() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee + executionFeeA);
    }

    function test_addExecutionFee_alreadyClaimed_emitsEvents() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee + executionFeeA);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFeeA);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
    }

    function test_addExecutionFee_alreadyClaimed_awardsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), executionFeeA);
    }

    function test_addExecutionFee_alreadyClaimed_recordsTxFee() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionId), executionFee + executionFeeA);
    }

    function test_recordExecutor_emitsEvents() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        vm.expectEmit(address(executionFees));
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFee);
        recordExecutor(dstChainId, transactionId, executor);
    }

    function test_recordExecutor_awardsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee);
        assertEq(executionFees.unclaimedRewards(executor), executionFee);
    }

    function test_recordExecutor_recordsExecutor() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionId), executor);
    }

    function test_recordExecutor_feeNotAdded_emitsEvents() public {
        vm.expectEmit(address(executionFees));
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, 0);
        recordExecutor(dstChainId, transactionId, executor);
    }

    function test_recordExecutor_feeNotAdded_doesNotAwardFees() public {
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.accumulatedRewards(executor), 0);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_recordExecutor_feeNotAdded_recordsExecutor() public {
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionId), executor);
    }

    function test_recordExecutor_addedTwice_sameTx_emitsEvents() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        vm.expectEmit(address(executionFees));
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFee + executionFeeA);
        recordExecutor(dstChainId, transactionId, executor);
    }

    function test_recordExecutor_addedTwice_sameTx_awardsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), executionFee + executionFeeA);
    }

    function test_recordExecutor_addedTwice_sameTx_recordsExecutor() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionId), executor);
    }

    function test_recordExecutor_addedTwice_diffTx_emitsEvents() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        vm.expectEmit(address(executionFees));
        emit ExecutorRecorded(dstChainId, transactionIdA, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFeeA);
        recordExecutor(dstChainId, transactionIdA, executor);
    }

    function test_recordExecutor_addedTwice_diffTx_awardsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee);
        assertEq(executionFees.unclaimedRewards(executor), executionFee);
    }

    function test_recordExecutor_addedTwice_diffTx_recordsExecutor() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionId), executor);
    }

    function setupAddedTwiceClaimedOnce() internal {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
    }

    function test_recordExecutor_addedTwice_diffTxClaimed_emitsEvents() public {
        setupAddedTwiceClaimedOnce();
        vm.expectEmit(address(executionFees));
        emit ExecutorRecorded(dstChainId, transactionIdA, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesAwarded(executor, executionFeeA);
        recordExecutor(dstChainId, transactionIdA, executor);
    }

    function test_recordExecutor_addedTwice_diffTxClaimed_awardsFees() public {
        setupAddedTwiceClaimedOnce();
        recordExecutor(dstChainId, transactionIdA, executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), executionFeeA);
    }

    function test_recordExecutor_addedTwice_diffTxClaimed_recordsExecutor() public {
        setupAddedTwiceClaimedOnce();
        recordExecutor(dstChainId, transactionIdA, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionIdA), executor);
    }

    function test_claimExecutionFees_emitsEvent() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesClaimed(executor, executionFee);
        executionFees.claimExecutionFees(executor);
    }

    function test_claimExecutionFees_transfersEther() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        uint256 beforeBalance = executor.balance;
        executionFees.claimExecutionFees(executor);
        uint256 afterBalance = executor.balance;
        assertEq(afterBalance - beforeBalance, executionFee);
    }

    function test_claimExecutionFees_recordsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_claimExecutionFees_addedTwice_sameTx_emitsEvent() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesClaimed(executor, executionFee + executionFeeA);
        executionFees.claimExecutionFees(executor);
    }

    function test_claimExecutionFees_addedTwice_sameTx_transfersEther() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        uint256 beforeBalance = executor.balance;
        executionFees.claimExecutionFees(executor);
        uint256 afterBalance = executor.balance;
        assertEq(afterBalance - beforeBalance, executionFee + executionFeeA);
    }

    function test_claimExecutionFees_addedTwice_sameTx_recordsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_claimExecutionFees_addedTwice_diffTx_emitsEvent() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesClaimed(executor, executionFee);
        executionFees.claimExecutionFees(executor);
    }

    function test_claimExecutionFees_addedTwice_diffTx_transfersEther() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        uint256 beforeBalance = executor.balance;
        executionFees.claimExecutionFees(executor);
        uint256 afterBalance = executor.balance;
        assertEq(afterBalance - beforeBalance, executionFee);
    }

    function test_claimExecutionFees_addedTwice_diffTx_recordsFees() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        addExecutionFee(executionFeeA, dstChainId, transactionIdA);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_claimExecutionFees_addedTwice_diffTxClaimed_emitsEvent() public {
        setupAddedTwiceClaimedOnce();
        recordExecutor(dstChainId, transactionIdA, executor);
        vm.expectEmit(address(executionFees));
        emit ExecutionFeesClaimed(executor, executionFeeA);
        executionFees.claimExecutionFees(executor);
    }

    function test_claimExecutionFees_addedTwice_diffTxClaimed_transfersEther() public {
        setupAddedTwiceClaimedOnce();
        recordExecutor(dstChainId, transactionIdA, executor);
        uint256 beforeBalance = executor.balance;
        executionFees.claimExecutionFees(executor);
        uint256 afterBalance = executor.balance;
        assertEq(afterBalance - beforeBalance, executionFeeA);
    }

    function test_claimExecutionFees_addedTwice_diffTxClaimed_recordsFees() public {
        setupAddedTwiceClaimedOnce();
        recordExecutor(dstChainId, transactionIdA, executor);
        executionFees.claimExecutionFees(executor);
        assertEq(executionFees.accumulatedRewards(executor), executionFee + executionFeeA);
        assertEq(executionFees.unclaimedRewards(executor), 0);
    }

    function test_executionFee_wrongChainId() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId + 1, transactionId), 0);
    }

    function test_executionFee_wrongTxId() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        assertEq(executionFees.executionFee(dstChainId, transactionIdA), 0);
    }

    function test_executor_wrongChainId() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId + 1, transactionId), address(0));
    }

    function test_executor_wrongTxId() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        assertEq(executionFees.recordedExecutor(dstChainId, transactionIdA), address(0));
    }

    function test_addExecutionFee_revertZeroFee() public {
        vm.expectRevert(IExecutionFees.ExecutionFees__ZeroAmount.selector);
        executionFees.addExecutionFee(dstChainId, transactionId);
    }

    function test_recordExecutor_revertExecutorAlreadyRecorded_same() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectRevert(IExecutionFees.ExecutionFees__AlreadyRecorded.selector);
        recordExecutor(dstChainId, transactionId, executor);
    }

    function test_recordExecutor_revertExecutorAlreadyRecorded_different() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        vm.expectRevert(IExecutionFees.ExecutionFees__AlreadyRecorded.selector);
        recordExecutor(dstChainId, transactionId, executorA);
    }

    function test_recordExecutor_revertExecutorZeroAddress() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        vm.expectRevert(IExecutionFees.ExecutionFees__ZeroAddress.selector);
        recordExecutor(dstChainId, transactionId, address(0));
    }

    function test_claimExecutionFees_revertNoUnclaimedRewards() public {
        addExecutionFee(executionFee, dstChainId, transactionId);
        recordExecutor(dstChainId, transactionId, executor);
        executionFees.claimExecutionFees(executor);
        vm.expectRevert(IExecutionFees.ExecutionFees__ZeroAmount.selector);
        executionFees.claimExecutionFees(executor);
    }
}
