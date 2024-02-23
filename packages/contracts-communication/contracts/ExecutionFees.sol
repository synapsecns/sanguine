// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFeesEvents} from "./events/ExecutionFeesEvents.sol";
import {IExecutionFees} from "./interfaces/IExecutionFees.sol";

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract ExecutionFees is AccessControl, ExecutionFeesEvents, IExecutionFees {
    bytes32 public constant RECORDER_ROLE = keccak256("RECORDER_ROLE");

    mapping(uint256 chainId => mapping(bytes32 transactionId => uint256 fee)) internal _executionFees;
    mapping(address executor => uint256 totalAccumulated) internal _accumulatedRewards;
    mapping(address executor => uint256 totalClaimed) internal _unclaimedRewards;
    mapping(uint256 chainId => mapping(bytes32 transactionId => address executor)) internal _recordedExecutor;

    constructor(address initialAdmin) {
        _grantRole(DEFAULT_ADMIN_ROLE, initialAdmin);
    }

    // @inheritdoc IExecutionFees
    function addExecutionFee(uint256 dstChainId, bytes32 transactionId) external payable {
        if (msg.value == 0) revert ExecutionFees__ZeroAmount();
        _executionFees[dstChainId][transactionId] += msg.value;
        // Use the new total fee as the event parameter.
        emit ExecutionFeeAdded(dstChainId, transactionId, _executionFees[dstChainId][transactionId]);
        address executor = _recordedExecutor[dstChainId][transactionId];
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
        if (_recordedExecutor[dstChainId][transactionId] != address(0)) revert ExecutionFees__AlreadyRecorded();
        uint256 fee = _executionFees[dstChainId][transactionId];
        _recordedExecutor[dstChainId][transactionId] = executor;
        emit ExecutorRecorded(dstChainId, transactionId, executor);
        _awardFee(executor, fee);
    }

    // @inheritdoc IExecutionFees
    function claimExecutionFees(address executor) external {
        uint256 amount = _unclaimedRewards[executor];
        if (amount == 0) revert ExecutionFees__ZeroAmount();
        _unclaimedRewards[executor] = 0;
        Address.sendValue(payable(executor), amount);
        emit ExecutionFeesClaimed(executor, amount);
    }

    // @inheritdoc IExecutionFees
    function getAccumulatedRewards(address executor) external view returns (uint256 accumulated) {
        return _accumulatedRewards[executor];
    }

    // @inheritdoc IExecutionFees
    function getUnclaimedRewards(address executor) external view returns (uint256 unclaimed) {
        return _unclaimedRewards[executor];
    }

    // @inheritdoc IExecutionFees
    function getExecutionFee(uint256 dstChainId, bytes32 transactionId) external view returns (uint256 fee) {
        return _executionFees[dstChainId][transactionId];
    }

    // @inheritdoc IExecutionFees
    function getRecordedExecutor(uint256 dstChainId, bytes32 transactionId) external view returns (address executor) {
        return _recordedExecutor[dstChainId][transactionId];
    }

    /// @dev Award the executor with the fee for completing the transaction.
    function _awardFee(address executor, uint256 fee) internal {
        _accumulatedRewards[executor] += fee;
        _unclaimedRewards[executor] += fee;
        emit ExecutionFeesAwarded(executor, fee);
    }
}
