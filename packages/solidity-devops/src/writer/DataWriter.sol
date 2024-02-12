// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {Logger} from "../base/Logger.sol";
import {PathFinder} from "../base/PathFinder.sol";
import {StringUtils} from "../libs/StringUtils.sol";

import {stdJson} from "forge-std/StdJson.sol";

abstract contract DataWriter is PathFinder, Logger {
    using stdJson for string;
    using StringUtils for *;

    /// @notice Writes a JSON data to a file, and prints a log message.
    function writeJson(string memory descriptionLog, string memory path, string memory data) internal {
        printLogWithIndent(descriptionLog);
        createDirIfRequired(path);
        data.write(path);
    }

    /// @notice Writes the deployment JSON for a contract on a given chain under the specified alias.
    /// Example: contractName = "LinkedPool", contractAlias = "LinkedPool.USDC"
    /// Note: writes to the FRESH deployment path, which is moved to the correct location after the contract is deployed.
    /// Note: requires ffi to be turned on, and jq to be installed.
    function writeDeploymentArtifact(
        string memory chain,
        string memory contractName,
        string memory contractAlias,
        string memory artifactWithoutABI
    )
        internal
        returns (string memory path)
    {
        path = writeDeploymentArtifactWithoutABI(chain, contractAlias, artifactWithoutABI);
        // Then, append the ABI to the deployment JSON. This will put the "abi" key after the "address" key,
        // improving readability of the JSON file.
        // Use contract name to determine the artifact path
        string memory fullJson = addJsonKey({pathInput: getArtifactFN(contractName), pathOutput: path, key: ".abi"});
        // Finally, write the full deployment JSON
        fullJson.write(path);
    }

    /// @notice Writes the deployment JSON for a contract on a given chain under the specified alias.
    /// Example: contractName = "LinkedPool", contractAlias = "LinkedPool.USDC"
    /// Note: writes to the FRESH deployment path, which is moved to the correct location after the contract is deployed.
    /// Note: will not include the ABI in the output JSON. Unlike `writeDeploymentArtifact`, has no dependencies.
    function writeDeploymentArtifactWithoutABI(
        string memory chain,
        string memory contractAlias,
        string memory artifactWithoutABI
    )
        internal
        returns (string memory path)
    {
        // Use contract alias to determine the deployment path
        path = getFreshDeploymentFN(chain, contractAlias);
        // First write the deployment JSON without the ABI
        writeJson(StringUtils.concat("Saving deployment for ", contractAlias, " on ", chain), path, artifactWithoutABI);
    }

    /// @notice Writes the deploy config for a contract on a given chain.
    function writeDeployConfig(string memory chain, string memory contractName, string memory configData) internal {
        writeJson(
            StringUtils.concat("Saving deploy config for ", contractName, " on ", chain),
            getDeployConfigFN(chain, contractName),
            configData
        );
    }

    /// @notice Writes the global deploy config that is shared across all chains for a contract.
    function writeGlobalDeployConfig(
        string memory contractName,
        string memory globalProperty,
        string memory configData
    )
        internal
    {
        writeJson(
            StringUtils.concat("Saving global config for ", contractName, ": ", globalProperty),
            getGlobalDeployConfigFN(contractName, globalProperty),
            configData
        );
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    /// @notice Reads value associated with a key from the input JSON file, and then writes it to the output JSON file.
    /// Will overwrite the value in the output JSON file if it already exists, otherwise will append it.
    /// Note: requires ffi to be turned on, and jq to be installed.
    function addJsonKey(
        string memory pathInput,
        string memory pathOutput,
        string memory key
    )
        internal
        returns (string memory fullInputData)
    {
        assertFileExists(pathInput);
        assertFileExists(pathOutput);
        // Example: jq .abi=$data.abi --argfile data path/to/input.json path/to/output.json
        string[] memory inputs = new string[](6);
        inputs[0] = "jq";
        inputs[1] = key.concat(" = $data", key);
        inputs[2] = "--argfile";
        inputs[3] = "data";
        inputs[4] = pathInput;
        inputs[5] = pathOutput;
        return string(vm.ffi(inputs));
    }

    /// @notice Creates a directory where the file will be saved if it doesn't exist yet.
    function createDirIfRequired(string memory filePath) internal {
        // Path to the directory is everything that comes before the last slash
        uint256 lastSlash = filePath.lastIndexOf("/");
        if (lastSlash == StringUtils.NOT_FOUND) {
            return;
        }
        string memory dirPath = filePath.prefix(lastSlash);
        if (!vm.exists(dirPath)) {
            printLog(StringUtils.concat("Creating directory: ", dirPath));
            vm.createDir({path: dirPath, recursive: true});
        }
    }
}
