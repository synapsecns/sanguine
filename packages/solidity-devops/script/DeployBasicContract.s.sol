// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {BasicContract} from "./BasicContract.sol";
import {SynapseScript, StringUtils} from "../src/SynapseScript.sol";

/// Note: it is recommended to handle the contract configuration in a separate script.
/// - `deployAndSave` will skip the deployment if the contract artifact is already saved in the deployments directory.
/// - `deployAndSave` will return the deployment address regardless of whether the contract
/// has been deployed previously or just now.
/// - `deployAndSaveAs` behaves similarly to `deployAndSave`, but allows for custom aliasing.
contract DeployBasicContract is SynapseScript {
    using StringUtils for *;

    string public contractName = "BasicContract";

    /// @notice Showcases different deployment methods.

    function run() external broadcastWithHooks {
        // Deploy using the contract-specific callback
        deployAndSave({contractName: contractName, deployContractFunc: cbDeployBasicContract});
        // Or, deploy using a generic callback for vanilla deployment
        // Need to craft the constructor arguments manually
        deployAndSave({contractName: contractName, constructorArgs: abi.encode(42), deployCodeFunc: cbDeploy});
        // Or, use the generic callback for CREATE2 deployment
        // Need to craft the constructor arguments manually
        deployAndSave({contractName: contractName, constructorArgs: abi.encode(42), deployCodeFunc: cbDeployCreate2});
    }

    /// @notice Showcases how to deploy a contract with a custom alias.
    /// Note: it is recommended to handle the contract configuration in a separate script.
    function run(string memory aliasSuffix) external broadcastWithHooks {
        // Typical alias derivation scheme
        string memory contractAlias = contractName.concat(".", aliasSuffix);
        // Deploy using the contract-specific callback
        deployAndSaveAs({
            contractName: contractName,
            contractAlias: contractAlias,
            deployContractFunc: cbDeployBasicContract
        });
        // Or, deploy using a generic callback for vanilla deployment
        // Need to craft the constructor arguments manually
        deployAndSaveAs({
            contractName: contractName,
            contractAlias: contractAlias,
            constructorArgs: abi.encode(42),
            deployCodeFunc: cbDeploy
        });
        // Or, use the generic callback for CREATE2 deployment
        // Need to craft the constructor arguments manually
        deployAndSaveAs({
            contractName: contractName,
            contractAlias: contractAlias,
            constructorArgs: abi.encode(42),
            deployCodeFunc: cbDeployCreate2
        });
    }

    /// @notice Contract-specific deployment callback.
    /// - MUST follow the interface
    ///     - function deployContractFunc() internal returns (address deployedAt, bytes memory constructorArgs)
    /// - MUST return the correct deployment address and constructor arguments
    /// - SHOULD only deploy contract without any additional configuration
    function cbDeployBasicContract() internal returns (address deployedAt, bytes memory constructorArgs) {
        constructorArgs = abi.encode(42);
        deployedAt = address(new BasicContract(42));
    }
}
