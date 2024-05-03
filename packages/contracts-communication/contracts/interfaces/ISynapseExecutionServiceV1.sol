// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IExecutionService} from "./IExecutionService.sol";

interface ISynapseExecutionServiceV1 is IExecutionService {
    error SynapseExecutionService__ExecutorZeroAddress();
    error SynapseExecutionService__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);
    error SynapseExecutionService__GasOracleNotContract(address gasOracle);
    error SynapseExecutionService__GasOracleZeroAddress();
    error SynapseExecutionService__OptionsVersionNotSupported(uint16 version);

    function setClaimerFraction(uint256 claimerFraction) external;
    function setExecutorEOA(address executorEOA_) external;
    function setGasOracle(address gasOracle_) external;
    function setGlobalMarkup(uint256 globalMarkup_) external;

    function gasOracle() external view returns (address);
    function globalMarkup() external view returns (uint256);
}
