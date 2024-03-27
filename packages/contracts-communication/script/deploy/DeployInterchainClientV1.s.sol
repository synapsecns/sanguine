// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1} from "../../contracts/InterchainClientV1.sol";

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract DeployInterchainClientV1 is SynapseScript {
    address public interchainDB;

    function run() external broadcastWithHooks {
        deployAndSave("InterchainClientV1", cdDeployInterchainClientV1);
    }

    function beforeExecution() internal override {
        super.beforeExecution();
        interchainDB = getDeploymentAddress({contractName: "InterchainDB", revertIfNotFound: true});
    }

    function cdDeployInterchainClientV1() internal returns (address deployedAt, bytes memory constructorArgs) {
        deployedAt = address(new InterchainClientV1(interchainDB, msg.sender));
        constructorArgs = abi.encode(interchainDB, msg.sender);
    }
}
