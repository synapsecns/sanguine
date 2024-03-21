// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {StringUtils} from "../libs/StringUtils.sol";
import {ChainAwareReader} from "../reader/ChainAwareReader.sol";
import {ChainAwareWriter} from "../writer/ChainAwareWriter.sol";

import {console2, stdJson} from "forge-std/Script.sol";

abstract contract DeploymentSaver is ChainAwareReader, ChainAwareWriter {
    using stdJson for string;
    using StringUtils for *;

    function checkBeforeDeploying(string memory contractName) internal view returns (address deployedAt) {
        return checkBeforeDeploying(contractName, contractName);
    }

    function checkBeforeDeploying(
        string memory contractName,
        string memory contractAlias
    )
        internal
        view
        returns (address deployedAt)
    {
        printLog(
            contractName.equals(contractAlias)
                ? StringUtils.concat("Deploying: ", contractName)
                : StringUtils.concat("Deploying: ", contractName, " as ", contractAlias)
        );
        // Check if the contract is already deployed
        deployedAt = getDeploymentAddress({contractName: contractAlias, revertIfNotFound: false});
        if (deployedAt != address(0)) {
            printSkipWithIndent(StringUtils.concat("already deployed at ", vm.toString(deployedAt)));
        }
    }

    /// @notice Deploys a contract and saves the deployment JSON, if the contract hasn't been deployed yet.
    /// See `deployAndSaveAs` below for more details.
    function deployAndSave(
        string memory contractName,
        bytes memory constructorArgs,
        function(string memory, bytes memory) internal returns (address) deployCodeFunc
    )
        internal
        returns (address deployedAt)
    {
        // Use contractName as contractAlias by default
        return deployAndSaveAs(contractName, contractName, constructorArgs, deployCodeFunc);
    }

    /// @notice Deploys a contract and saves the deployment JSON, if the contract hasn't been deployed yet.
    /// Needs to be passed a generic function that deploys the contract, that follows this signature:
    /// deployCodeFunc(string memory contractName, bytes memory constructorArgs) internal returns (address deployedAt);
    /// - Example: cbDeploy() and cbDeployCreate2() could be used here, as they follow the signature. Or anything
    /// else that could deploy raw bytecode of any contract, most likely using assembly/factory approach.
    /// - Note: contract should be configured outside of `deployCodeFunc`.
    function deployAndSaveAs(
        string memory contractName,
        string memory contractAlias,
        bytes memory constructorArgs,
        function(string memory, bytes memory) internal returns (address) deployCodeFunc
    )
        internal
        returns (address deployedAt)
    {
        deployedAt = checkBeforeDeploying(contractName, contractAlias);
        if (deployedAt != address(0)) {
            return deployedAt;
        }
        // Trigger callback to deploy the contract
        deployedAt = deployCodeFunc(contractName, constructorArgs);
        printSuccessWithIndent(StringUtils.concat("Deployed at ", vm.toString(deployedAt)));
        // Save the deployment JSON
        saveDeployment(contractName, contractAlias, deployedAt, constructorArgs);
    }

    /// @notice Deploys a contract and saves the deployment JSON, if the contract hasn't been deployed yet.
    /// See `deployAndSaveAs` below for more details.
    function deployAndSave(
        string memory contractName,
        function() internal returns (address, bytes memory) deployContractFunc
    )
        internal
        returns (address deployedAt)
    {
        // Use contractName as contractAlias by default
        return deployAndSaveAs(contractName, contractName, deployContractFunc);
    }

    /// @notice Deploys a contract and saves the deployment JSON, if the contract hasn't been deployed yet.
    /// Needs to be passed a contract-specific function that deploys the contract, that follows this signature:
    /// deployContractFunc() internal returns (address deployedAt, bytes memory constructorArgs);
    /// - Example: use anything that deploys a specific contract, most likely using `new Contract(...);` approach.
    /// - Note: contract should be configured outside of `deployContractFunc`.
    function deployAndSaveAs(
        string memory contractName,
        string memory contractAlias,
        function() internal returns (address, bytes memory) deployContractFunc
    )
        internal
        returns (address deployedAt)
    {
        deployedAt = checkBeforeDeploying(contractName, contractAlias);
        if (deployedAt != address(0)) {
            return deployedAt;
        }
        // Trigger callback to deploy the specific contract
        bytes memory constructorArgs;
        (deployedAt, constructorArgs) = deployContractFunc();
        printSuccessWithIndent(StringUtils.concat("Deployed at ", vm.toString(deployedAt)));
        // Save the deployment JSON
        saveDeployment(contractName, contractAlias, deployedAt, constructorArgs);
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    /// @notice Produces a JSON string that can be used to save a contract deployment.
    /// Note: contract ABI is not included in the output.
    function serializeDeploymentData(
        address deployedAt,
        bytes memory constructorArgs
    )
        internal
        returns (string memory data)
    {
        data = "deployment";
        data.serialize("address", deployedAt);
        if (constructorArgs.length == 0) {
            return data.serialize("constructorArgs", string("0x"));
        } else {
            return data.serialize("constructorArgs", constructorArgs);
        }
    }

    /// @notice Saves the deployment JSON for a contract on a given chain under the specified alias.
    /// Example: contractName = "LinkedPool", contractAlias = "LinkedPool.USDC"
    /// Note: writes to the FRESH deployment path, which is moved to the correct location after the contract is deployed.
    /// Note: requires ffi to be turned on, and jq to be installed.
    function saveDeployment(
        string memory contractName,
        string memory contractAlias,
        address deployedAt,
        bytes memory constructorArgs
    )
        internal
        withIndent
    {
        writeDeploymentArtifact({
            contractName: contractName,
            contractAlias: contractAlias,
            artifactWithoutABI: serializeDeploymentData(deployedAt, constructorArgs)
        });
    }

    /// @notice Saves the deployment JSON for a contract on a given chain under the specified alias.
    /// Example: contractName = "LinkedPool", contractAlias = "LinkedPool.USDC"
    /// Note: writes to the FRESH deployment path, which is moved to the correct location after the contract is deployed.
    /// Note: will not include the ABI in the output JSON. Unlike `saveDeployment`, has no dependencies.
    function saveDeploymentWithoutABI(
        string memory contractAlias,
        address deployedAt,
        bytes memory constructorArgs
    )
        internal
        withIndent
    {
        writeDeploymentArtifactWithoutABI({
            contractAlias: contractAlias,
            artifactWithoutABI: serializeDeploymentData(deployedAt, constructorArgs)
        });
    }
}
