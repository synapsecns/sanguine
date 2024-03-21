// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService, ISynapseExecutionServiceV1} from "../interfaces/ISynapseExecutionServiceV1.sol";
import {SynapseExecutionServiceEvents} from "../events/SynapseExecutionServiceEvents.sol";
import {IGasOracle} from "../interfaces/IGasOracle.sol";
import {OptionsLib, OptionsV1} from "../libs/Options.sol";

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract SynapseExecutionServiceV1 is
    AccessControlUpgradeable,
    SynapseExecutionServiceEvents,
    ISynapseExecutionServiceV1
{
    /// @custom:storage-location erc7201:Synapse.ExecutionService.V1
    struct SynapseExecutionServiceV1Storage {
        address executorEOA;
        address gasOracle;
    }

    // keccak256(abi.encode(uint256(keccak256("Synapse.ExecutionService.V1")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant SYNAPSE_EXECUTION_SERVICE_V1_STORAGE_LOCATION =
        0xabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100;

    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

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
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.executorEOA = executorEOA_;
        emit ExecutorEOASet(executorEOA_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setGasOracle(address gasOracle_) external onlyRole(GOVERNOR_ROLE) {
        if (gasOracle_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        $.gasOracle = gasOracle_;
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
        onlyRole(IC_CLIENT_ROLE)
    {
        uint256 requiredFee = getExecutionFee(dstChainId, txPayloadSize, options);
        if (executionFee < requiredFee) {
            revert SynapseExecutionService__FeeAmountTooLow({actual: executionFee, required: requiredFee});
        }
        emit ExecutionRequested({transactionId: transactionId, client: msg.sender});
    }

    /// @inheritdoc IExecutionService
    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        public
        view
        returns (uint256 executionFee)
    {
        address cachedGasOracle = gasOracle();
        if (cachedGasOracle == address(0)) {
            revert SynapseExecutionService__GasOracleNotSet();
        }
        // TODO: the "exact version" check should be generalized
        (uint8 version,) = OptionsLib.decodeVersionedOptions(options);
        if (version > OptionsLib.OPTIONS_V1) {
            revert SynapseExecutionService__OptionsVersionNotSupported(version);
        }
        OptionsV1 memory optionsV1 = OptionsLib.decodeOptionsV1(options);
        // TODO: there should be a markup applied to the execution fee
        executionFee = IGasOracle(cachedGasOracle).estimateTxCostInLocalUnits({
            remoteChainId: dstChainId,
            gasLimit: optionsV1.gasLimit,
            calldataSize: txPayloadSize
        });
        if (optionsV1.gasAirdrop > 0) {
            executionFee += IGasOracle(cachedGasOracle).convertRemoteValueToLocalUnits({
                remoteChainId: dstChainId,
                value: optionsV1.gasAirdrop
            });
        }
    }

    /// @inheritdoc IExecutionService
    function executorEOA() public view returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.executorEOA;
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function gasOracle() public view returns (address) {
        SynapseExecutionServiceV1Storage storage $ = _getSynapseExecutionServiceV1Storage();
        return $.gasOracle;
    }

    /// @dev ERC-7201 slot accessor
    function _getSynapseExecutionServiceV1Storage() private pure returns (SynapseExecutionServiceV1Storage storage $) {
        assembly {
            $.slot := SYNAPSE_EXECUTION_SERVICE_V1_STORAGE_LOCATION
        }
    }
}
