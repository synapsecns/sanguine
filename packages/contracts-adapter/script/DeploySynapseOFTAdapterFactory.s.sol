// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapterFactory} from "../src/SynapseOFTAdapterFactory.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Deploys the OFT adapter factory using CREATE2 and initializes its chain-specific endpoint.
contract DeploySynapseOFTAdapterFactory is SynapseScript {
    using stdJson for string;

    string public constant NAME = "SynapseOFTAdapterFactory";

    /// @notice Excludes this script from coverage reports.
    // solhint-disable-next-line no-empty-blocks
    function testDeploySynapseOFTAdapterFactory() external {}

    function run() external broadcastWithHooks {
        string memory environment = vm.envOr("DEPLOY_ENVIRONMENT", string("production"));
        if (keccak256(bytes(environment)) == keccak256("production")) environment = ENVIRONMENT_PROD;
        string memory config = readGlobalDeployConfig("SynapseOFTAdapter", environment, true);
        address endpoint = config.readAddress(string.concat(".endpoints.", activeChain));
        require(endpoint.code.length > 0, "Endpoint is not a contract");

        bytes memory constructorArgs = abi.encode(msg.sender);
        address predicted = predictAddress(getInitCode(NAME, constructorArgs), bytes32(0));
        printInfo(string.concat("Initial owner: ", vm.toString(msg.sender)));
        printInfo(string.concat("Endpoint: ", vm.toString(endpoint)));
        printInfo(string.concat("Predicted factory: ", vm.toString(predicted)));
        setNextDeploymentSalt(bytes32(0));
        SynapseOFTAdapterFactory factory = SynapseOFTAdapterFactory(
            deployAndSave({contractName: NAME, constructorArgs: constructorArgs, deployCodeFunc: cbDeployCreate2})
        );
        require(address(factory) == predicted, "Factory address mismatch");
        require(factory.owner() == msg.sender, "Factory owner mismatch");
        if (factory.endpoint() == address(0)) {
            factory.initialize(endpoint);
            printSuccessWithIndent("Factory endpoint initialized");
        } else {
            require(factory.endpoint() == endpoint, "Factory endpoint mismatch");
            printSkipWithIndent("Factory endpoint already initialized");
        }
    }
}
