// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {SynapseOFTAdapterFactory} from "../src/SynapseOFTAdapterFactory.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Deploys one token's OFT adapter through the saved, owner-controlled factory.
contract DeploySynapseOFTAdapter is SynapseScript {
    using stdJson for string;

    SynapseOFTAdapterFactory internal factory;
    address internal token;
    bytes32 internal salt;

    /// @notice Excludes this script from coverage reports.
    function testDeploySynapseOFTAdapter() external {}

    /// @param tokenId The token identifier used in the config and deployment alias, e.g. SYN.
    function run(string memory tokenId) external broadcastWithHooks {
        string memory environment = vm.envOr("DEPLOY_ENVIRONMENT", string("production"));
        if (keccak256(bytes(environment)) == keccak256("production")) environment = ENVIRONMENT_PROD;
        require(bytes(tokenId).length > 0, "Token identifier is empty");
        string memory aliasName = string.concat("SynapseOFTAdapter.", tokenId);
        string memory config = readGlobalDeployConfig("SynapseOFTAdapter", environment, true);
        string memory tokenPath = string.concat(".tokens.", tokenId);
        token = config.readAddress(string.concat(tokenPath, ".addresses.", activeChain));
        salt = config.readBytes32(string.concat(tokenPath, ".salt"));
        factory = SynapseOFTAdapterFactory(getDeploymentAddress("SynapseOFTAdapterFactory", true));
        require(factory.owner() == msg.sender, "Wallet does not own factory");
        require(factory.endpoint() != address(0), "Factory is not initialized");
        require(token.code.length > 0, "Token is not a contract");

        address predicted = predictAddress(address(factory), type(SynapseOFTAdapter).creationCode, salt);
        printInfo(string.concat("Token: ", vm.toString(token)));
        printInfo(string.concat("Salt: ", vm.toString(salt)));
        printInfo(string.concat("Predicted adapter: ", vm.toString(predicted)));
        SynapseOFTAdapter adapter = SynapseOFTAdapter(
            deployAndSaveAs({
                contractName: "SynapseOFTAdapter",
                contractAlias: aliasName,
                deployContractFunc: deployAdapter
            })
        );
        require(address(adapter) == predicted, "Adapter address mismatch");
        require(adapter.token() == token, "Adapter token mismatch");
        require(address(adapter.endpoint()) == factory.endpoint(), "Adapter endpoint mismatch");
        require(adapter.owner() == msg.sender, "Adapter owner mismatch");
    }

    function deployAdapter() internal returns (address deployedAt, bytes memory constructorArgs) {
        deployedAt = factory.deploy(token, salt);
        constructorArgs = "";
    }
}
