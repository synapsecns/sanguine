// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService} from "./interfaces/IExecutionService.sol";
import {ExecutionServiceEvents} from "./events/ExecutionServiceEvents.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {IGasOracle} from "./interfaces/IGasOracle.sol";

import {OptionsLib, OptionsV1} from "./libs/Options.sol";

contract ExecutionService is ExecutionServiceEvents, Ownable, IExecutionService {
    using OptionsLib for bytes;

    address public interchainClient;
    address public executorEOA;
    IGasOracle public gasOracle;

    constructor() Ownable(msg.sender) {}

    function setInterchainClient(address _interchainClient) external onlyOwner {
        interchainClient = _interchainClient;
        emit InterchainClientUpdated(_interchainClient);
    }

    function setExecutorEOA(address _executorEOA) external onlyOwner {
        executorEOA = _executorEOA;
        emit ExecutorEOAUpdated(executorEOA);
    }

    function setGasOracle(address _gasOracle) external onlyOwner {
        gasOracle = IGasOracle(_gasOracle);
        emit GasOracleUpdated(_gasOracle);
    }

    modifier onlyInterchainClient() {
        require(msg.sender == interchainClient, "ExecutionService: caller is not the InterchainClient");
        _;
    }

    // @inheritdoc
    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
        override onlyInterchainClient
    {
        require(executionFee > getExecutionFee(dstChainId, txPayloadSize, options), "ExecutionService: execution fee is not high enough");
        emit ExecutionRequested(dstChainId, txPayloadSize, transactionId, executionFee, options);
    }

    // @inheritdoc
    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        external
        view
        override
        returns (uint256)
    {
        (uint8 version, bytes memory data) = options.decodeVersionedOptions();
        if (version == OptionsLib.OPTIONS_V1) {
            OptionsV1 memory optionsV1 = data.decodeOptionsV1();
            uint256 baseCost = gasOracle.estimateTxCostInLocalUnits(dstChainId, optionsV1.gasLimit, txPayloadSize);
            if (optionsV1.gasAirdrop > 0) {
                baseCost += gasOracle.convertRemoteValueToLocalUnits(dstChainId, optionsV1.gasAirdrop);
            }
            return baseCost;
        } else {
            revert("Unsupported options version: version must be OPTIONS_V1");
        }
    }
}
