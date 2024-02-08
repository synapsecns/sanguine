// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Interface for the Interchain Factory.
/// @dev The Processor contract needs the Interchain token address during deployment,
/// and the Interchain token contract needs the Processor address during deployment as well.
/// To remove this circular dependency, both contracts are supposed to be deployed from the factory.
/// This effectively removes their address dependency during deployment, allowing to predict their addresses in advance.
/// The predicted addresses are stored in the Interchain Factory during the deployment, and are used
/// in the constructor of the Processor and Interchain token contracts using the exposed callbacks.
/// Note: assuming InterchainFactory is deployed on the same address on different chains,
/// the address of the interchain token will depend on the following values:
/// - Deployer address (if enforced by the factory)
/// - Interchain Token name, symbol, and decimals (constructor values for the Interchain Token)
interface IInterchainFactory {
    /// @notice Returns the parameters required for the deployment of the current Interchain Token.
    /// @dev Will revert if no Interchain Token is being deployed.
    function getInterchainTokenDeployParameters() external view returns (address initialAdmin, address processor);

    /// @notice Returns the parameters required for the deployment of the current Processor.
    /// @dev Will revert if no Processor is being deployed.
    function getProcessorDeployParameters() external view returns (address interchainToken, address underlyingToken);
}
