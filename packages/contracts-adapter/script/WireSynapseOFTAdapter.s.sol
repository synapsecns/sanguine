// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {LayerZeroWiring} from "./helpers/LayerZeroWiring.sol";

import {UlnBase, UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Wires a token's OFT peers, message libraries, and ULN security configuration.
contract WireSynapseOFTAdapter is LayerZeroWiring {
    using stdJson for string;

    /// @notice Excludes this script from coverage reports.
    function testWireSynapseOFTAdapter() external {}

    /// @param tokenId The token identifier used in the global config and deployment alias, e.g. SYN.
    function run(string memory tokenId) external broadcastWithHooks {
        string memory environment = vm.envOr("DEPLOY_ENVIRONMENT", string("production"));
        if (keccak256(bytes(environment)) == keccak256("production")) environment = ENVIRONMENT_PROD;
        string memory config = readGlobalDeployConfig("SynapseOFTAdapter", environment, true);
        string memory tokenPath = string.concat(".tokens.", tokenId, ".addresses");
        address token = config.readAddress(string.concat(tokenPath, ".", activeChain));
        // Configure routes before remote tokens and adapters are deployed; missing peers are skipped.
        allChains = vm.parseJsonKeys(config, ".wiring.chains");
        chainsConfig = config;
        securityConfig = config;
        chainsConfigRoot = ".wiring.chains";
        securityConfigRoot = ".wiring";
        requiredDVNs = config.readAddressArray(string.concat(".wiring.requiredDVNs.", activeChain));
        require(requiredDVNs.length > 0 && requiredDVNs.length <= 127, "Invalid required DVN count");
        requiredDVNs = sortAddresses(requiredDVNs);
        address previousDvn;
        for (uint256 i = 0; i < requiredDVNs.length; ++i) {
            require(requiredDVNs[i] > previousDvn, "Invalid or duplicate DVN");
            previousDvn = requiredDVNs[i];
        }
        loadChainsConfig(config.readAddress(string.concat(".endpoints.", activeChain)));

        string memory aliasName = string.concat("SynapseOFTAdapter.", tokenId);
        address deployment = getDeploymentAddress(aliasName, true);
        require(SynapseOFTAdapter(deployment).token() == token, "Adapter token mismatch");
        wireApp(deployment, aliasName);
    }

    function getUlnConfig(uint64 confirmations) internal view override returns (UlnConfig memory config) {
        config = super.getUlnConfig(confirmations);
        // Explicitly disable optional DVNs; zero would inherit the library default.
        config.optionalDVNCount = type(uint8).max;
    }

    function getCurrentUlnConfig(address lib, uint32 eid) internal view override returns (bytes memory) {
        // Compare raw app settings so inherited defaults are replaced with explicit configuration.
        return abi.encode(UlnBase(lib).getAppUlnConfig(address(app), eid));
    }
}
