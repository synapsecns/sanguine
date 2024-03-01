// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Deploys a contract that takes no arguments in its constructor.
contract DeployNoArgs is SynapseScript {
    function run(string memory contractName) external broadcastWithHooks {
        bytes memory constructorArgs = "";
        deployAndSave(contractName, constructorArgs, cbDeploy);
    }
}
