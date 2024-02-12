// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {Deployer} from "./deploy/Deployer.sol";
import {DeploymentSaver} from "./deploy/DeploymentSaver.sol";
import {StringUtils} from "./libs/StringUtils.sol";

import {Script} from "forge-std/Script.sol";

abstract contract SynapseBaseScript is Script, Deployer, DeploymentSaver {
    uint256 private initialDeployerNonce;

    /// @notice Common pattern for running a script.
    modifier broadcastWithHooks() {
        beforeExecution();
        vm.startBroadcast();
        _;
        vm.stopBroadcast();
        afterExecution();
    }

    /// @notice Common pattern for running a script off-chain.
    modifier offChainWithHooks() {
        beforeExecution();
        _;
        afterExecution();
    }

    /// @notice Hook that is called before the script is executed.
    /// @dev Could be overridden to load custom data before the script is executed.
    /// Make sure to call `super.beforeExecution()` in the overridden function.
    function beforeExecution() internal virtual {
        loadDevopsConfig();
        loadActiveChain();
        loadEnvCreate2Factory();
        initialDeployerNonce = vm.getNonce(msg.sender);
    }

    /// @notice Hook that is called after the script is executed.
    /// @dev Could be overridden to perform various checks after the script is executed.
    /// Make sure to call `super.afterExecution()` in the overridden function.
    function afterExecution() internal virtual {
        // Check if the deployer address has the same nonce as before the script was executed
        uint256 finalDeployerNonce = vm.getNonce(msg.sender);
        printInfo(
            StringUtils.concat(
                "Total amount of broadcasted txs: ", vm.toString(finalDeployerNonce - initialDeployerNonce)
            )
        );
    }
}
