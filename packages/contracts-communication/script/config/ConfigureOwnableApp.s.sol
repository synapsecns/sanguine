// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OwnableApp} from "../../contracts/apps/OwnableApp.sol";
import {AppConfigV1} from "../../contracts/libs/AppConfig.sol";

import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

import {stdJson, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable code-complexity
abstract contract ConfigureOwnableApp is SynapseScript {
    using stdJson for string;

    string public appName;

    string public config;
    OwnableApp public app;

    constructor(string memory appName_) {
        appName = appName_;
    }

    function run(string memory environment) external virtual broadcastWithHooks {
        loadConfig(environment);
        linkRemoteChains();
        syncTrustedModules();
        setAppConfig();
        afterAppConfigured();
    }

    function loadConfig(string memory environment) internal virtual {
        app = OwnableApp(getDeploymentAddress({contractName: appName, revertIfNotFound: true}));
        config = readGlobalDeployConfig({contractName: appName, globalProperty: environment, revertIfNotFound: true});
    }

    function linkRemoteChains() internal virtual {
        printLog("Linking remote chains");
        string[] memory chains = config.readStringArray(".chains");
        for (uint256 i = 0; i < chains.length; i++) {
            string memory chain = chains[i];
            uint256 chainId = chainIds[chain];
            require(chainId != 0, string.concat("Chain not found: ", chain));
            // Skip current chain
            if (chainId == blockChainId()) continue;
            address remoteApp = getDeploymentAddress({chain: chain, contractName: appName, revertIfNotFound: true});
            bytes32 linkedApp = app.getLinkedApp(chainId);
            if (TypeCasts.addressToBytes32(remoteApp) != linkedApp) {
                app.linkRemoteAppEVM(chainId, remoteApp);
                printSuccessWithIndent(string.concat("Linked ", vm.toString(remoteApp), " on ", chain));
            } else {
                printSkipWithIndent(string.concat("already linked to ", vm.toString(remoteApp), " on ", chain));
            }
        }
    }

    function syncTrustedModules() internal virtual {
        printLog("Syncing trusted modules");
        string[] memory moduleNames = config.readStringArray(".trustedModules");
        address[] memory trustedModules = new address[](moduleNames.length);
        for (uint256 i = 0; i < moduleNames.length; i++) {
            trustedModules[i] = getDeploymentAddress({contractName: moduleNames[i], revertIfNotFound: true});
        }
        address[] memory existingModules = app.getReceivingModules();
        // Remove modules that are not in the config
        uint256 removed = 0;
        for (uint256 i = 0; i < existingModules.length; i++) {
            if (!contains(trustedModules, existingModules[i])) {
                app.removeTrustedModule(existingModules[i]);
                printSuccessWithIndent(string.concat("Removed ", vm.toString(existingModules[i])));
                removed++;
            }
        }
        // Add modules that are in the config but not in the app
        uint256 added = 0;
        for (uint256 i = 0; i < trustedModules.length; i++) {
            if (!contains(existingModules, trustedModules[i])) {
                printSuccessWithIndent(
                    string.concat("Added ", vm.toString(trustedModules[i]), " [", moduleNames[i], "]")
                );
                app.addTrustedModule(trustedModules[i]);
                added++;
            }
        }
        if (removed + added == 0) {
            printSkipWithIndent("modules are up to date");
        } else {
            printLog(string.concat("Added ", vm.toString(added), " modules, removed ", vm.toString(removed)));
        }
    }

    function setAppConfig() internal virtual {
        printLog("Setting app config");
        increaseIndent();
        bytes memory rawConfig = config.parseRaw(".appConfig");
        AppConfigV1 memory appConfig = abi.decode(rawConfig, (AppConfigV1));
        printLog(string.concat("Required responses: ", vm.toString(appConfig.requiredResponses)));
        printLog(string.concat("Optimistic period: ", vm.toString(appConfig.optimisticPeriod)));
        AppConfigV1 memory existingConfig = app.getAppConfigV1();
        if (
            appConfig.requiredResponses == existingConfig.requiredResponses
                && appConfig.optimisticPeriod == existingConfig.optimisticPeriod
        ) {
            printSkipWithIndent("config is already set");
        } else {
            app.setAppConfigV1(appConfig);
            printSuccessWithIndent("Config set");
        }
        decreaseIndent();
    }

    function afterAppConfigured() internal virtual {}

    function contains(address[] memory array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) return true;
        }
        return false;
    }
}
