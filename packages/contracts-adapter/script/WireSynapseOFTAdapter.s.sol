// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {LayerZeroWiring} from "./helpers/LayerZeroWiring.sol";

import {UlnBase, UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {
    EnforcedOptionParam,
    IOAppOptionsType3
} from "@layerzerolabs/oapp-evm/contracts/oapp/interfaces/IOAppOptionsType3.sol";
import {OptionsBuilder} from "@layerzerolabs/oapp-evm/contracts/oapp/libs/OptionsBuilder.sol";
import {stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Wires a token's OFT peers, message libraries, ULN security configuration, and enforced options.
contract WireSynapseOFTAdapter is LayerZeroWiring {
    using OptionsBuilder for bytes;
    using stdJson for string;

    /// @notice Excludes this script from coverage reports.
    // solhint-disable-next-line no-empty-blocks
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
        setEnforcedOptions();
    }

    function setEnforcedOptions() internal {
        printLog("Setting enforced options...");
        SynapseOFTAdapter adapter = SynapseOFTAdapter(address(app));
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            uint32 eid = eidByChainName[chain];
            if (eid == eidByChainName[activeChain] || eidSkipped[eid]) continue;
            string memory path = string.concat(".wiring.enforcedOptions.", chain);
            uint256 receiveGas = chainsConfig.readUint(string.concat(path, ".lzReceiveGas"));
            uint256 composeGas = chainsConfig.readUint(string.concat(path, ".lzComposeGas"));
            require(receiveGas > 0 && receiveGas <= type(uint128).max, "Invalid receive gas");
            require(composeGas <= type(uint128).max, "Invalid compose gas");
            bytes memory options = OptionsBuilder.newOptions().addExecutorLzReceiveOption(uint128(receiveGas), 0);
            bytes memory composeOptions = options;
            if (composeGas > 0) {
                composeOptions = composeOptions.addExecutorLzComposeOption(0, uint128(composeGas), 0);
            }
            setEnforcedOption(adapter, chain, adapter.SEND(), options);
            setEnforcedOption(adapter, chain, adapter.SEND_AND_CALL(), composeOptions);
        }
    }

    function setEnforcedOption(
        SynapseOFTAdapter adapter,
        string memory chain,
        uint16 msgType,
        bytes memory options
    )
        internal
    {
        uint32 eid = eidByChainName[chain];
        string memory action = string.concat(formatChainName(chain), "enforced options type ", vm.toString(msgType));
        if (keccak256(adapter.enforcedOptions(eid, msgType)) == keccak256(options)) {
            printSkipWithIndent(string.concat(action, " already set"));
            return;
        }
        EnforcedOptionParam[] memory params = new EnforcedOptionParam[](1);
        params[0] = EnforcedOptionParam({eid: eid, msgType: msgType, options: options});
        if (printPeerMultisigTxs) {
            printMultisigTx(action, address(adapter), abi.encodeCall(IOAppOptionsType3.setEnforcedOptions, (params)));
            return;
        }
        adapter.setEnforcedOptions(params);
        printSuccessWithIndent(action);
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
