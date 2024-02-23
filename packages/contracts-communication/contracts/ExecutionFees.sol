// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionFees} from "./interfaces/IExecutionFees.sol";
import {ExecutionFeesEvents} from "./events/ExecutionFeesEvents.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract ExecutionFees is AccessControl, ExecutionFeesEvents, IExecutionFees {
    using Address for address payable;

    bytes32 public constant RECORDER_ROLE = keccak256("RECORDER_ROLE");

    // Interchain Transaction IDs => Total Execution Fees
    mapping(bytes32 => uint256) private _executionFees;
    // Executor Addresses => Total Accumulated Fees
    mapping(address => uint256) private _accumulatedRewards;
    // Executor Addresses => Currently unclaimed fees
    mapping(address => uint256) private _unclaimedRewards;
    // Interchain Transaction IDs => Executor Addresses
    mapping(bytes32 => address) private _transactionsByExecutor;

    constructor(address initialAdmin) {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin);
    }

    // @inheritdoc IExecutionFees
    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable override {
        if (msg.value == 0) revert ExecutionFees__ZeroAmount();
        _executionFees[transactionId] += msg.value;
        emit ExecutionFeeAdded(dstChainId, transactionId, _executionFees[transactionId]);
        address executor = _transactionsByExecutor[transactionId];
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
        override
        onlyRole(RECORDER_ROLE)
    {
        if (executor == address(0)) revert ExecutionFees__ZeroAddress();
        if (_transactionsByExecutor[transactionId] != address(0)) revert ExecutionFees__AlreadyRecorded();
        uint256 fee = _executionFees[transactionId];
        _transactionsByExecutor[transactionId] = executor;
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        _awardFee(executor, fee);
    }

    // @inheritdoc IExecutionFees
    function claimExecutionFees(address executor) external override {
        uint256 amount = _unclaimedRewards[executor];
        if (amount == 0) revert ExecutionFees__ZeroAmount();
        _unclaimedRewards[executor] = 0;
        payable(executor).sendValue(amount);
        emit ExecutionFeesClaimed(executor, amount);
    }

    // @inheritdoc IExecutionFees
    function getAccumulatedRewards(address executor) external view override returns (uint256 accumulated) {
        return _accumulatedRewards[executor];
    }

    // @inheritdoc IExecutionFees
    function getUnclaimedRewards(address executor) external view override returns (uint256 unclaimed) {
        return _unclaimedRewards[executor];
    }

    /// @dev Award the executor with the fee for completing the transaction.
    function _awardFee(address executor, uint256 fee) internal {
        _accumulatedRewards[executor] += fee;
        _unclaimedRewards[executor] += fee;
        emit ExecutionFeesAwarded(executor, fee);
    }
}
