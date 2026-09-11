// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseOFTAdapterFactory {
    struct ConstructorParams {
        address token;
        address lzEndpoint;
        address owner;
    }

    event AdapterDeployed(address indexed token, address indexed adapter, bytes32 salt);

    /// @notice Deploys an adapter using CREATE2. Only callable by the factory owner.
    /// @param token The SynapseERC20 token to adapt.
    /// @param salt The CREATE2 salt, independent of the token address.
    /// @return adapter The deployed adapter address.
    function deploy(address token, bytes32 salt) external returns (address adapter);

    /// @notice Returns the constructor parameters for the adapter currently being deployed.
    /// @dev Parameters are cleared after deployment.
    function getConstructorParams() external view returns (ConstructorParams memory);
}
