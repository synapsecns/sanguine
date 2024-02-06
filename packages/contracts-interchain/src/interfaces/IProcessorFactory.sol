// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Interface for the Processor Factory.
/// @dev The Processor contract needs the Interchain token address during deployment,
/// and the Interchain token contract needs the Processor address during deployment as well.
/// To remove this circular dependency, the Processor is supposed to be deployed from the factory.
/// This effectively removes the Interchain token address from the Processor constructor arguments,
/// and allows to predict the Processor address in advance and pass it to the Interchain token constructor.
interface IProcessorFactory {
    /// @notice Returns the parameters required for the deployment of the current Processor.
    /// @dev Will revert if no Processor is being deployed.
    function getProcessorDeployParameters() external view returns (address interchainToken, address underlyingToken);
}
