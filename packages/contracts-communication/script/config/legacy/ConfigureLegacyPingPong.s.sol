// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyPingPong} from "../../../contracts/legacy/LegacyPingPong.sol";
import {TypeCasts} from "../../../contracts/libs/TypeCasts.sol";

import {stdJson, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureLegacyPingPong is SynapseScript {
    using stdJson for string;

    string public constant NAME = "LegacyPingPong";

    string public config;
    LegacyPingPong public legacyPingPong;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        setMessageBus();
        setTrustedRemotes();
        setGasLimit();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        address deployment = getDeploymentAddress({contractName: NAME, revertIfNotFound: true});
        legacyPingPong = LegacyPingPong(payable(deployment));
    }

    function setMessageBus() internal {
        printLog("Setting MessageBus");
        address messageBus = getDeploymentAddress({contractName: "MessageBus", revertIfNotFound: true});
        if (legacyPingPong.messageBus() != messageBus) {
            legacyPingPong.setMessageBus(messageBus);
            printSuccessWithIndent(string.concat("Set MessageBus to ", vm.toString(messageBus)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(messageBus)));
        }
    }

    function setTrustedRemotes() internal {
        printLog("Setting trusted remotes");
        string[] memory chains = config.readStringArray(".chains");
        for (uint256 i = 0; i < chains.length; i++) {
            string memory chain = chains[i];
            uint256 chainId = chainIds[chain];
            require(chainId != 0, string.concat("Chain not found: ", chain));
            // Skip current chain
            if (chainId == blockChainId()) continue;
            address remotePingPong = getDeploymentAddress({chain: chain, contractName: NAME, revertIfNotFound: true});
            bytes32 remotePingPongBytes32 = TypeCasts.addressToBytes32(remotePingPong);
            bytes32 trustedRemote = legacyPingPong.trustedRemotes(chainId);
            if (remotePingPongBytes32 != trustedRemote) {
                legacyPingPong.setTrustedRemote(chainId, remotePingPongBytes32);
                printSuccessWithIndent(
                    string.concat("Set trusted remote to ", vm.toString(remotePingPong), " on ", chain)
                );
            } else {
                printSkipWithIndent(string.concat("already set to ", vm.toString(remotePingPong), " on ", chain));
            }
        }
    }

    function setGasLimit() internal {
        printLog("Setting gas limit");
        uint256 gasLimit = config.readUint(".gasLimit");
        if (legacyPingPong.gasLimit() != gasLimit) {
            legacyPingPong.setGasLimit(gasLimit);
            printSuccessWithIndent(string.concat("Set gas limit to ", vm.toString(gasLimit)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(gasLimit)));
        }
    }
}
