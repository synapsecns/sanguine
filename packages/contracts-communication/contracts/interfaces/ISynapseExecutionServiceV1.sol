// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IExecutionService} from "./IExecutionService.sol";

interface ISynapseExecutionServiceV1 is IExecutionService {
    error SynapseExecutionService__GasOracleNotSet();
    error SynapseExecutionService__FeeAmountTooLow(uint256 actual, uint256 required);
    error SynapseExecutionService__OptionsVersionNotSupported(uint16 version);
    error SynapseExecutionService__ZeroAddress();

    /// @notice Allows the contract governor to set the address of the EOA account that will be used
    /// to execute transactions on the remote chains.
    function setExecutorEOA(address executorEOA_) external;

    /// @notice Allows the contract governor to set the address of the gas oracle.
    function setGasOracle(address gasOracle_) external;

    /// @notice Allows the contract governor to set the global markup that the Execution Service charges
    /// on top of the GasOracle's gas cost estimates.
    function setGlobalMarkup(uint256 globalMarkup_) external;

    /// @notice Address of the gas oracle used for estimating the gas cost of the transactions.
    function gasOracle() external view returns (address);

    /// @notice The markup that the Execution Service charges on top of the GasOracle's gas cost estimates.
    /// Zero markup means that the Execution Service charges the exact gas cost estimated by the GasOracle.
    /// The markup is denominated in Wei, 1e18 being 100%.
    function globalMarkup() external view returns (uint256);
}
