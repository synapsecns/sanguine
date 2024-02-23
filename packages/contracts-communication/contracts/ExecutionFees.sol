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
        require(msg.value > 0, "ExecutionFees: Fee must be greater than 0");
        require(_transactionsByExecutor[transactionId] == address(0), "ExecutionFees: Executor already recorded");
        _executionFees[transactionId] += msg.value;
        emit ExecutionFeeAdded(dstChainId, transactionId, msg.value);
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
        require(_transactionsByExecutor[transactionId] == address(0), "ExecutionFees: Executor already recorded");
        require(_executionFees[transactionId] > 0, "ExecutionFees: No execution fee found");
        _transactionsByExecutor[transactionId] = executor;
        _accumulatedRewards[executor] += _executionFees[transactionId];
        _unclaimedRewards[executor] += _executionFees[transactionId];
        emit ExecutorRecorded(dstChainId, transactionId, executor);
    }

    // @inheritdoc IExecutionFees
    function claimExecutionFees(address executor) external override {
        uint256 amount = _unclaimedRewards[executor];
        require(amount > 0, "ExecutionFees: No unclaimed rewards");
        _unclaimedRewards[executor] = 0;
        payable(executor).sendValue(amount);
        emit ExecutionFeesClaimed(msg.sender, amount);
    }

    // @inheritdoc IExecutionFees
    function getAccumulatedRewards(address executor) external view override returns (uint256 accumulated) {
        return _accumulatedRewards[executor];
    }

    // @inheritdoc IExecutionFees
    function getUnclaimedRewards(address executor) external view override returns (uint256 unclaimed) {
        return _unclaimedRewards[executor];
    }
}
