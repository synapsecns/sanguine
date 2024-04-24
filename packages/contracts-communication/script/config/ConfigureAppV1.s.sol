// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ICAppV1} from "../../contracts/apps/ICAppV1.sol";
import {AppConfigV1} from "../../contracts/libs/AppConfig.sol";

import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {stdJson, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

abstract contract ConfigureAppV1 is SynapseScript {
    using stdJson for string;

    string public appName;

    string public config;
    ICAppV1 public app;

    constructor(string memory appName_) {
        appName = appName_;
    }

    function run(string memory environment) external virtual broadcastWithHooks {
        loadConfig(environment);
        beforeAppConfigured();
        linkRemoteChains();
        syncTrustedModules();
        setAppConfig();
        setExecutionService();
        setLatestInterchainClient();
        syncInterchainClients();
        afterAppConfigured();
    }

    function loadConfig(string memory environment) internal virtual {
        app = ICAppV1(getDeploymentAddress({contractName: appName, revertIfNotFound: true}));
        config = readGlobalDeployConfig({contractName: appName, globalProperty: environment, revertIfNotFound: true});
    }

    function linkRemoteChains() internal virtual {
        printLog("Linking remote chains");
        string[] memory chains = config.readStringArray(".chains");
        for (uint256 i = 0; i < chains.length; i++) {
            string memory chain = chains[i];
            uint64 chainId = SafeCast.toUint64(chainIds[chain]);
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
        address[] memory existingModules = app.getModules();
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
        uint256 requiredResponses = config.readUint(".appConfig.requiredResponses");
        uint256 optimisticPeriod = config.readUint(".appConfig.optimisticPeriod");
        printLog(string.concat("Required responses: ", vm.toString(requiredResponses)));
        printLog(string.concat("Optimistic period: ", vm.toString(optimisticPeriod)));
        AppConfigV1 memory existingConfig = app.getAppConfigV1();
        if (
            requiredResponses == existingConfig.requiredResponses && optimisticPeriod == existingConfig.optimisticPeriod
        ) {
            printSkipWithIndent("config is already set");
        } else {
            app.setAppConfigV1(requiredResponses, optimisticPeriod);
            printSuccessWithIndent("Config set");
        }
        decreaseIndent();
    }

    function setExecutionService() internal virtual {
        printLog("Setting execution service");
        address executionService =
            getDeploymentAddress({contractName: "SynapseExecutionServiceV1", revertIfNotFound: true});
        if (app.getExecutionService() != executionService) {
            app.setExecutionService(executionService);
            printSuccessWithIndent(string.concat("Execution service set to ", vm.toString(executionService)));
        } else {
            printSkipWithIndent(string.concat("execution service is already set to ", vm.toString(executionService)));
        }
    }

    function setLatestInterchainClient() internal virtual {
        printLog("Setting latest interchain client");
        string memory clientName = config.readString(".latestInterchainClient");
        address client = getDeploymentAddress({contractName: clientName, revertIfNotFound: true});
        string memory desc = string.concat("set to ", clientName, " [", vm.toString(client), "]");
        if (app.getLatestInterchainClient() != client) {
            if (contains(app.getInterchainClients(), client)) {
                app.setLatestInterchainClient(client);
                printSuccessWithIndent(string.concat("Latest client ", desc));
            } else {
                app.addInterchainClient({client: client, updateLatest: true});
                printSuccessWithIndent(string.concat("Latest client added and ", desc));
            }
        } else {
            printSkipWithIndent(string.concat("client is already set to ", vm.toString(client)));
        }
    }

    function syncInterchainClients() internal virtual {
        printLog("Syncing interchain clients");
        string[] memory clientNames = config.readStringArray(".interchainClients");
        address[] memory clients = new address[](clientNames.length);
        for (uint256 i = 0; i < clientNames.length; i++) {
            clients[i] = getDeploymentAddress({contractName: clientNames[i], revertIfNotFound: true});
        }
        address[] memory existingClients = app.getInterchainClients();
        // Remove clients that are not in the config
        uint256 removed = 0;
        for (uint256 i = 0; i < existingClients.length; i++) {
            if (!contains(clients, existingClients[i])) {
                app.removeInterchainClient(existingClients[i]);
                printSuccessWithIndent(string.concat("Removed ", vm.toString(existingClients[i])));
                removed++;
            }
        }
        // Add clients that are in the config but not in the app
        uint256 added = 0;
        for (uint256 i = 0; i < clients.length; i++) {
            if (!contains(existingClients, clients[i])) {
                printSuccessWithIndent(string.concat("Added ", vm.toString(clients[i]), " [", clientNames[i], "]"));
                app.addInterchainClient({client: clients[i], updateLatest: false});
                added++;
            }
        }
        if (removed + added == 0) {
            printSkipWithIndent("clients are up to date");
        } else {
            printLog(string.concat("Added ", vm.toString(added), " clients, removed ", vm.toString(removed)));
        }
    }

    // solhint-disable-next-line no-empty-blocks
    function beforeAppConfigured() internal virtual {}

    // solhint-disable-next-line no-empty-blocks
    function afterAppConfigured() internal virtual {}

    function contains(address[] memory array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) return true;
        }
        return false;
    }
}
