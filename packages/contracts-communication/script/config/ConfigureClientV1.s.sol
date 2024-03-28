// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainClientV1} from "../../contracts/InterchainClientV1.sol";
import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

import {stdJson, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureClientV1 is SynapseScript {
    using stdJson for string;

    string public constant NAME = "InterchainClientV1";

    InterchainClientV1 public client;
    string public config;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        linkRemoteChains();
        setExecutionFees();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        client = InterchainClientV1(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function linkRemoteChains() internal {
        printLog("Linking remote chains");
        string[] memory chains = config.readStringArray(".chains");
        for (uint256 i = 0; i < chains.length; i++) {
            string memory chain = chains[i];
            uint256 chainId = chainIds[chain];
            require(chainId != 0, string.concat("Chain not found: ", chain));
            // Skip current chain
            if (chainId == blockChainId()) continue;
            address remoteClientEVM = getDeploymentAddress({chain: chain, contractName: NAME, revertIfNotFound: true});
            bytes32 remoteClient = TypeCasts.addressToBytes32(remoteClientEVM);
            bytes32 linkedClient = client.getLinkedClient(chainId);
            if (remoteClient != linkedClient) {
                client.setLinkedClient(chainId, remoteClient);
                printSuccessWithIndent(string.concat("Linked ", vm.toString(remoteClientEVM), " on ", chain));
            } else {
                printSkipWithIndent(string.concat("already linked to ", vm.toString(remoteClientEVM), " on ", chain));
            }
        }
    }

    function setExecutionFees() internal {
        printLog("Setting ExecutionFees");
        address executionFees = getDeploymentAddress({contractName: "ExecutionFees", revertIfNotFound: true});
        if (client.executionFees() != executionFees) {
            client.setExecutionFees(executionFees);
            printSuccessWithIndent(string.concat("Set ExecutionFees to ", vm.toString(executionFees)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(executionFees)));
        }
    }
}
