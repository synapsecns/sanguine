// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract DeploySBA is SynapseScript {
    using stdJson for string;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySBA() external {}

    function run() external broadcastWithHooks {
        string memory chainsConfig = readGlobalDeployProdConfig({contractName: "chains", revertIfNotFound: true});
        string memory chainKey = string.concat(".", activeChain);
        if (!chainsConfig.keyExists(chainKey)) {
            printFailWithIndent(string.concat("Chain ", activeChain, " not found in chains config"));
            assert(false);
        }
        address endpointV2 = chainsConfig.readAddress(string.concat(chainKey, ".endpointV2"));
        // constructor(address endpoint_, address owner_)
        bytes memory constructorArgs = abi.encode(endpointV2, msg.sender);
        printInfo(string.concat("EndpointV2: ", vm.toString(endpointV2)));
        printInfo(string.concat("Initial Owner: ", vm.toString(msg.sender)));
        deployAndSave({
            contractName: "SynapseBridgeAdapter",
            constructorArgs: constructorArgs,
            deployCodeFunc: cbDeployCreate2
        });
    }
}
