// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1, ISynapseGasOracleV1} from "../../contracts/oracles/SynapseGasOracleV1.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {stdJson, StringUtils, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable custom-errors
contract ConfigureSynapseGasOracleV1 is SynapseScript {
    using stdJson for string;
    using StringUtils for *;

    string public constant NAME = "SynapseGasOracleV1";

    string public config;
    SynapseGasOracleV1 public gasOracle;

    function run(string memory environment) external broadcastWithHooks {
        loadConfig(environment);
        setLocalNativePrice();
        setAllRemoteGasData();
    }

    function loadConfig(string memory environment) internal {
        config = readGlobalDeployConfig({contractName: NAME, globalProperty: environment, revertIfNotFound: true});
        gasOracle = SynapseGasOracleV1(getDeploymentAddress({contractName: NAME, revertIfNotFound: true}));
    }

    function setLocalNativePrice() internal {
        printLog("Setting local native price");
        uint256 nativePrice = config.readUint(string.concat(".", activeChain, ".nativePrice"));
        string memory desc = string.concat(nativePrice.fromWei(), " ETH");
        if (nativePrice != gasOracle.getLocalNativePrice()) {
            gasOracle.setLocalNativePrice(nativePrice);
            printSuccessWithIndent(string.concat("Set local native price to ", desc));
        } else {
            printSkipWithIndent(string.concat("already set to ", desc));
        }
    }

    function setAllRemoteGasData() internal {
        printLog("Setting remote gas data");
        string[] memory chains = vm.parseJsonKeys(config, ".");
        for (uint256 i = 0; i < chains.length; i++) {
            string memory chain = chains[i];
            uint64 chainId = SafeCast.toUint64(chainIds[chain]);
            require(chainId != 0, string.concat("Chain not found: ", chain));
            // Skip current chain
            if (chainId == blockChainId()) continue;
            setRemoteGasData(chainId, chain);
        }
    }

    function setRemoteGasData(uint64 chainId, string memory chain) internal withIndent {
        printLog(chain);
        bytes memory chainConfigRaw = config.parseRaw(string.concat(".", chain));
        ISynapseGasOracleV1.RemoteGasData memory chainConfig =
            abi.decode(chainConfigRaw, (ISynapseGasOracleV1.RemoteGasData));
        ISynapseGasOracleV1.RemoteGasData memory current = gasOracle.getRemoteGasData(chainId);
        if (!equals(chainConfig, current)) {
            gasOracle.setRemoteGasData(chainId, chainConfig);
        }
        string memory desc = string.concat(chainConfig.calldataPrice.fromFloat(9), " gwei");
        if (current.calldataPrice != chainConfig.calldataPrice) {
            printSuccessWithIndent(string.concat("Set calldataPrice to ", desc));
        } else {
            printSkipWithIndent(string.concat("calldataPrice already set to ", desc));
        }
        desc = string.concat(chainConfig.gasPrice.fromFloat(9), " gwei");
        if (current.gasPrice != chainConfig.gasPrice) {
            printSuccessWithIndent(string.concat("Set gasPrice to ", desc));
        } else {
            printSkipWithIndent(string.concat("gasPrice already set to ", desc));
        }
        desc = string.concat(chainConfig.nativePrice.fromWei(), " ETH");
        if (current.nativePrice != chainConfig.nativePrice) {
            printSuccessWithIndent(string.concat("Set nativePrice to ", desc));
        } else {
            printSkipWithIndent(string.concat("nativePrice already set to ", desc));
        }
    }

    function equals(
        ISynapseGasOracleV1.RemoteGasData memory a,
        ISynapseGasOracleV1.RemoteGasData memory b
    )
        internal
        pure
        returns (bool)
    {
        return a.calldataPrice == b.calldataPrice && a.gasPrice == b.gasPrice && a.nativePrice == b.nativePrice;
    }
}
