// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceV1} from "../../contracts/execution/SynapseExecutionServiceV1.sol";

import {DeployProxy} from "./DeployProxy.s.sol";

contract DeploySynapseExecutionServiceV1 is DeployProxy {
    function run() external broadcastWithHooks {
        bytes memory initData = abi.encodeCall(SynapseExecutionServiceV1.initialize, (msg.sender));
        address implementation = deployAndSaveAs({
            contractName: "SynapseExecutionServiceV1",
            contractAlias: "SynapseExecutionServiceV1.Implementation",
            constructorArgs: "",
            deployCodeFunc: cbDeploy
        });
        deployAndSaveProxy({
            contractName: "SynapseExecutionServiceV1",
            implementation: implementation,
            proxyAdminOwner: msg.sender,
            initData: initData
        });
    }
}
