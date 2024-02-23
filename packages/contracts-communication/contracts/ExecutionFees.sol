// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFeesEvents} from "./events/ExecutionFeesEvents.sol";
import {IExecutionFees} from "./interfaces/IExecutionFees.sol";

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract ExecutionFees is AccessControl, ExecutionFeesEvents, IExecutionFees {
    bytes32 public constant RECORDER_ROLE = keccak256("RECORDER_ROLE");

    /// @inheritdoc IExecutionFees
    mapping(uint256 chainId => mapping(bytes32 transactionId => uint256 fee)) public executionFee;
    /// @inheritdoc IExecutionFees
    mapping(address executor => uint256 totalAccumulated) public accumulatedRewards;
    /// @inheritdoc IExecutionFees
    mapping(address executor => uint256 totalClaimed) public unclaimedRewards;
    /// @inheritdoc IExecutionFees
    mapping(uint256 chainId => mapping(bytes32 transactionId => address executor)) public recordedExecutor;

    constructor(address initialAdmin) {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin);
    }

    // @inheritdoc IExecutionFees
    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable {
        if (msg.value == 0) revert ExecutionFees__ZeroAmount();
        executionFee[dstChainId][transactionId] += msg.value;
        // Use the new total fee as the event parameter.
        emit ExecutionFeeAdded(dstChainId, transactionId, executionFee[dstChainId][transactionId]);
        address executor = recordedExecutor[dstChainId][transactionId];
        // If the executor is recorded, the previous fee has been awarded already. Award the new fee.
        if (executor != address(0)) {
            _awardFee(executor, msg.value);
        }
    }

    // @inheritdoc IExecutionFees
    function recordExecutor(
        uint256 dstChainId,
        bytes32 transactionId,
        address executor
    )
        external
        onlyRole(RECORDER_ROLE)
    {
        if (executor == address(0)) revert ExecutionFees__ZeroAddress();
        if (recordedExecutor[dstChainId][transactionId] != address(0)) revert ExecutionFees__AlreadyRecorded();
        uint256 fee = executionFee[dstChainId][transactionId];
        recordedExecutor[dstChainId][transactionId] = executor;
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        _awardFee(executor, fee);
    }

    // @inheritdoc IExecutionFees
    function claimExecutionFees(address executor) external {
        uint256 amount = unclaimedRewards[executor];
        if (amount == 0) revert ExecutionFees__ZeroAmount();
        unclaimedRewards[executor] = 0;
        Address.sendValue(payable(executor), amount);
        emit ExecutionFeesClaimed(executor, amount);
    }

    /// @dev Award the executor with the fee for completing the transaction.
    function _awardFee(address executor, uint256 fee) internal {
        accumulatedRewards[executor] += fee;
        unclaimedRewards[executor] += fee;
        emit ExecutionFeesAwarded(executor, fee);
    }
}
