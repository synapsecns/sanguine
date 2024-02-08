// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Thin wrapper around the Create2 EVM opcode.
library Create2 {
    error Create2__AlreadyDeployed(address deployedAt);
    error Create2__DeploymentFailed();

    /// @notice Deploys a contract using the CREATE2 opcode, and returns the deployed contract's address.
    /// @dev Will revert if the contract is already deployed at the predicted address, or if the deployment fails.
    /// @param value            The amount of wei to send with the deployment.
    /// @param salt             A salt to make the CREATE2 address unique.
    /// @param creationCode     The contract's creation bytecode.
    /// @return deployedAt      The address of the deployed contract.
    function deploy(uint256 value, bytes32 salt, bytes memory creationCode) internal returns (address deployedAt) {
        deployedAt = predictDeployment(address(this), salt, creationCode);
        if (deployedAt.code.length > 0) revert Create2__AlreadyDeployed(deployedAt);
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // create2(value, offset, size, salt)
            deployedAt := create2(value, add(creationCode, 0x20), mload(creationCode), salt)
        }
        if (deployedAt == address(0)) revert Create2__DeploymentFailed();
    }

    /// @notice Predicts the address of a contract that would be deployed using the CREATE2 opcode.
    /// @param deployer         The address of the deploying contract.
    /// @param salt             A salt to make the CREATE2 address unique.
    /// @param creationCode     The contract's creation bytecode.
    /// @return deployedAt      The address of the deployed contract.
    function predictDeployment(
        address deployer,
        bytes32 salt,
        bytes memory creationCode
    )
        internal
        pure
        returns (address deployedAt)
    {
        deployedAt = address(
            uint160(uint256(keccak256(abi.encodePacked(bytes1(0xff), deployer, salt, keccak256(creationCode)))))
        );
    }
}
