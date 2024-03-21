// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService, ISynapseExecutionServiceV1} from "../interfaces/ISynapseExecutionServiceV1.sol";
import {SynapseExecutionServiceEvents} from "../events/SynapseExecutionServiceEvents.sol";

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract SynapseExecutionServiceV1 is
    AccessControlUpgradeable,
    SynapseExecutionServiceEvents,
    ISynapseExecutionServiceV1
{
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

    // TODO: Use ERC-7201 to manage storage
    /// @inheritdoc IExecutionService
    address public executorEOA;
    /// @inheritdoc ISynapseExecutionServiceV1
    address public gasOracle;

    constructor() {
        // Ensure that the implementation contract could not be initialized
        _disableInitializers();
    }

    function initialize(address admin) external initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setExecutorEOA(address executorEOA_) external onlyRole(GOVERNOR_ROLE) {
        if (executorEOA_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        executorEOA = executorEOA_;
        emit ExecutorEOASet(executorEOA_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setGasOracle(address gasOracle_) external onlyRole(GOVERNOR_ROLE) {
        if (gasOracle_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        gasOracle = gasOracle_;
        emit GasOracleSet(gasOracle_);
    }

    /// @inheritdoc IExecutionService
    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
    {
        // TODO: implement
    }

    /// @inheritdoc IExecutionService
    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        returns (uint256)
    {
        // TODO: implement
    }
}
