// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract DeploySynapseModule is SynapseScript {
    address public interchainDB;

    function run() external broadcastWithHooks {
        deployAndSave("SynapseModule", cdDeploySynapseModule);
    }

    function beforeExecution() internal override {
        super.beforeExecution();
        interchainDB = getDeploymentAddress({contractName: "InterchainDB", revertIfNotFound: true});
    }

    function cdDeploySynapseModule() internal returns (address deployedAt, bytes memory constructorArgs) {
        deployedAt = address(new SynapseModule(interchainDB, msg.sender));
        constructorArgs = abi.encode(interchainDB, msg.sender);
    }
}
