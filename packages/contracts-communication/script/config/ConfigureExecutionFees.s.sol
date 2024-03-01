// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionFees} from "../../contracts/ExecutionFees.sol";

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureExecutionFees is SynapseScript {
    string public constant NAME = "ExecutionFees";

    ExecutionFees public executionFees;

    function run() external broadcastWithHooks {
        executionFees = ExecutionFees(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
        addClientAsRecorder();
    }

    function addClientAsRecorder() internal {
        printLog("Adding InterchainClientV1 as Recorder");
        address client = getDeploymentAddress({contractName: "InterchainClientV1", revertIfNotFound: true});
        bytes32 role = executionFees.RECORDER_ROLE();
        if (!executionFees.hasRole(role, client)) {
            executionFees.grantRole(role, client);
            printSuccessWithIndent(string.concat("Added ", vm.toString(client), " as Recorder"));
        } else {
            printSkipWithIndent(string.concat(vm.toString(client), " already has Recorder role"));
        }
    }
}
