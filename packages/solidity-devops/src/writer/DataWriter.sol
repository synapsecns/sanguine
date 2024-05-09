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
    /// Note: writes the JSON file to the FRESH deployments directory. The written file needs to be moved
    /// to the correct location outside of the deployment script.
    /// Note: will not include the ABI in the output JSON.
    function writeDeploymentArtifact(
        string memory chain,
        string memory contractAlias,
        string memory artifact
    )
        internal
        returns (string memory path)
    {
        // Use contract alias to determine the deployment path
        path = getFreshDeploymentFN(chain, contractAlias);
        // First write the deployment JSON without the ABI
        writeJson(StringUtils.concat("Saving deployment for ", contractAlias, " on ", chain), path, artifact);
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
