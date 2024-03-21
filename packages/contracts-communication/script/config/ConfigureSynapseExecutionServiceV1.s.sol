// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceV1} from "../../contracts/execution/SynapseExecutionServiceV1.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureSynapseExecutionServiceV1 is SynapseScript {
    using stdJson for string;

    string public constant NAME = "ExecutionService";

    SynapseExecutionServiceV1 public service;
    string public config;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        setGovernor();
        setExecutorEOA();
        setGasOracle();
        setInterchainClient();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        service = SynapseExecutionServiceV1(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function setGovernor() internal {
        printLog("Setting Governor");
        bytes32 governorRole = service.GOVERNOR_ROLE();
        if (!service.hasRole(governorRole, msg.sender)) {
            service.grantRole(governorRole, msg.sender);
            printSuccessWithIndent(string.concat("Granted Governor role to ", vm.toString(msg.sender)));
        } else {
            printSkipWithIndent(string.concat("governor role already granted to ", vm.toString(msg.sender)));
        }
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

    function setInterchainClient() internal {
        // TODO: should support multiple clients
        address interchainClientV1 = getDeploymentAddress({contractName: "InterchainClientV1", revertIfNotFound: true});
        printLog("Setting InterchainClient on ExecutionService");
        if (!service.hasRole(service.IC_CLIENT_ROLE(), interchainClientV1)) {
            service.grantRole(service.IC_CLIENT_ROLE(), interchainClientV1);
            printSuccessWithIndent(string.concat("Granted IC_CLIENT_ROLE to ", vm.toString(interchainClientV1)));
        } else {
            printSkipWithIndent(string.concat("IC_CLIENT_ROLE already granted to ", vm.toString(interchainClientV1)));
        }
    }
}
