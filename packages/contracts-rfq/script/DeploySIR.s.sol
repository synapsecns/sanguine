// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract DeploySIR is SynapseScript {
    string public constant LATEST_SIR = "SynapseIntentRouter";
    string public constant LATEST_ZAP = "TokenZapV1";

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySIR() external {}

    function run() external broadcastWithHooks {
        // TODO: create2 salts
        deployAndSave({contractName: LATEST_SIR, constructorArgs: "", deployCodeFunc: cbDeployCreate2});
        deployAndSave({contractName: LATEST_ZAP, constructorArgs: "", deployCodeFunc: cbDeployCreate2});
    }
}
