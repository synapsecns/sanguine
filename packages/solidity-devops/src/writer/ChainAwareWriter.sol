// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {ChainAwareness} from "../base/ChainAwareness.sol";
import {StringUtils} from "../libs/StringUtils.sol";
import {DataWriter} from "../writer/DataWriter.sol";

abstract contract ChainAwareWriter is ChainAwareness, DataWriter {
    using StringUtils for *;

    /// @notice Writes the deployment JSON for a contract on the active chain under the specified alias.
    /// Example: contractName = "LinkedPool", contractAlias = "LinkedPool.USDC"
    /// Note: writes to the FRESH deployment path, which is moved to the correct location after the contract is deployed.
    /// Note: will not include the ABI in the output JSON.
    function writeDeploymentArtifact(string memory contractAlias, string memory artifact) internal {
        writeDeploymentArtifact(activeChain, contractAlias, artifact);
    }

    /// @notice Writes the deploy config for a contract on the active chain.
    function writeDeployConfig(string memory contractName, string memory configData) internal {
        writeDeployConfig(activeChain, contractName, configData);
    }
}
