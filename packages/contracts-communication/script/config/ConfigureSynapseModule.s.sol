// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseModule} from "../../contracts/modules/SynapseModule.sol";

import {stdJson, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

contract ConfigureSynapseModule is SynapseScript {
    using stdJson for string;

    string public constant NAME = "SynapseModule";

    string public config;
    SynapseModule public module;

    address[] public removedVerifiers;
    address[] public addedVerifiers;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        setClaimFeeFraction();
        setFeeCollector();
        setGasOracle();
        setThreshold();
        syncVerifiers();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        module = SynapseModule(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function setClaimFeeFraction() internal {
        printLog("Setting claimFeeFraction");
        uint256 claimFeeBPS = config.readUint(".claimFeeBPS");
        // Convert from basis points to wei
        uint256 claimFeeFraction = claimFeeBPS * 1e18 / 10_000;
        string memory desc = string.concat(vm.toString(claimFeeFraction), " [", vm.toString(claimFeeBPS), " BPS]");
        if (module.getClaimFeeFraction() != claimFeeFraction) {
            module.setClaimFeeFraction(claimFeeFraction);
            printSuccessWithIndent(string.concat("Set claimFeeFraction to ", desc));
        } else {
            printSkipWithIndent(string.concat("already set to ", desc));
        }
    }

    function setFeeCollector() internal {
        printLog("Setting feeCollector");
        address feeCollector = config.readAddress(".feeCollector");
        if (module.feeCollector() != feeCollector) {
            module.setFeeCollector(feeCollector);
            printSuccessWithIndent(string.concat("Set feeCollector to ", vm.toString(feeCollector)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(feeCollector)));
        }
    }

    function setGasOracle() internal {
        string memory gasOracleName = config.readString(".gasOracleName");
        printLog(string.concat("Setting GasOracle to ", gasOracleName));
        address gasOracle = getDeploymentAddress({contractName: gasOracleName, revertIfNotFound: true});
        if (module.gasOracle() != gasOracle) {
            module.setGasOracle(gasOracle);
            printSuccessWithIndent(string.concat("Set gasOracle to ", vm.toString(gasOracle)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(gasOracle)));
        }
    }

    function setThreshold() internal {
        printLog("Setting threshold");
        uint256 threshold = config.readUint(".threshold");
        if (module.getThreshold() != threshold) {
            module.setThreshold(threshold);
            printSuccessWithIndent(string.concat("Set threshold to ", vm.toString(threshold)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(threshold)));
        }
    }

    function syncVerifiers() internal {
        printLog("Syncing verifiers");
        address[] memory verifiers = config.readAddressArray(".verifiers");
        address[] memory existingVerifiers = module.getVerifiers();
        // Remove verifiers that are not in the config
        for (uint256 i = 0; i < existingVerifiers.length; i++) {
            if (!contains(verifiers, existingVerifiers[i])) {
                removedVerifiers.push(existingVerifiers[i]);
                printSuccessWithIndent(string.concat("Removing ", vm.toString(existingVerifiers[i])));
            }
        }
        if (removedVerifiers.length > 0) {
            module.removeVerifiers(removedVerifiers);
        }
        // Add verifiers that are in the config but not in the module
        for (uint256 i = 0; i < verifiers.length; i++) {
            if (!contains(existingVerifiers, verifiers[i])) {
                addedVerifiers.push(verifiers[i]);
                printSuccessWithIndent(string.concat("Adding ", vm.toString(verifiers[i])));
            }
        }
        if (addedVerifiers.length > 0) {
            module.addVerifiers(addedVerifiers);
        }
        if (removedVerifiers.length + addedVerifiers.length == 0) {
            printSkipWithIndent("verifiers are up to date");
        } else {
            printLog(
                string.concat(
                    "Added ",
                    vm.toString(addedVerifiers.length),
                    " verifiers, removed ",
                    vm.toString(removedVerifiers.length)
                )
            );
        }
    }

    function contains(address[] memory array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) return true;
        }
        return false;
    }
}
