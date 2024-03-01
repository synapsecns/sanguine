// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionService} from "../../contracts/ExecutionService.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureExecutionFees is SynapseScript {
    using stdJson for string;

    string public constant NAME = "ExecutionService";

    ExecutionService public service;
    string public config;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        setExecutorEOA();
        setGasOracle();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        service = ExecutionService(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function setExecutorEOA() internal {
        printLog("Setting ExecutorEOA");
        address executorEOA = config.readAddress(".executorEOA");
        if (service.executorEOA() != executorEOA) {
            service.setExecutorEOA(executorEOA);
            printSuccessWithIndent(string.concat("Set ExecutorEOA to ", vm.toString(executorEOA)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(executorEOA)));
        }
    }

    function setGasOracle() internal {
        string memory gasOracleName = config.readString(".gasOracleName");
        printLog(string.concat("Setting GasOracle to ", gasOracleName));
        address gasOracle = getDeploymentAddress({contractName: gasOracleName, revertIfNotFound: true});
        if (address(service.gasOracle()) != gasOracle) {
            service.setGasOracle(gasOracle);
            printSuccessWithIndent(string.concat("Set GasOracle to ", vm.toString(gasOracle)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(gasOracle)));
        }
    }
}
