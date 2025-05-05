// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseIntentPreviewer} from "../contracts/router/SynapseIntentPreviewer.sol";
import {SynapseIntentRouter} from "../contracts/router/SynapseIntentRouter.sol";
import {TokenZapV1} from "../contracts/zaps/TokenZapV1.sol";

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract DeploySIR is SynapseScript {
    string public constant LATEST_SIR = "SynapseIntentRouter";
    string public constant LATEST_SIP = "SynapseIntentPreviewer";
    string public constant LATEST_ZAP = "TokenZapV1";

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySIR() external {}

    function run() external broadcastWithHooks {
        // TODO: create2 salts
        address sir = deployAndSave({contractName: LATEST_SIR, constructorArgs: "", deployCodeFunc: cbDeployCreate2});
        address sip = deployAndSave({contractName: LATEST_SIP, constructorArgs: "", deployCodeFunc: cbDeployCreate2});
        address zap = deployAndSave({contractName: LATEST_ZAP, constructorArgs: "", deployCodeFunc: cbDeployCreate2});
        printLog("Checking deployments");
        checkDeployment(LATEST_SIR, sir, type(SynapseIntentRouter).runtimeCode);
        checkDeployment(LATEST_SIP, sip, type(SynapseIntentPreviewer).runtimeCode);
        checkDeployment(LATEST_ZAP, zap, type(TokenZapV1).runtimeCode);
    }

    function checkDeployment(string memory contractName, address deployedAt, bytes memory contractCode) internal view {
        if (keccak256(deployedAt.code) != keccak256(contractCode)) {
            printFailWithIndent(
                string.concat(contractName, " deployment at ", vm.toString(deployedAt), " is not the latest version")
            );
            assert(false);
        }
        printSuccessWithIndent(string.concat(contractName, " latest version deployed at ", vm.toString(deployedAt)));
    }
}
