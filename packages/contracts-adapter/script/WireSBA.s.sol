// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";

import {UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {ILayerZeroEndpointV2} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {SetConfigParam} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLibManager.sol";
import {AddressCast} from "@layerzerolabs/lz-evm-protocol-v2/contracts/libs/AddressCast.sol";
import {StringUtils, SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks, ordering
contract WireSBA is SynapseScript {
    using AddressCast for address;
    using stdJson for string;
    using StringUtils for string;

    uint32 internal constant CONFIG_TYPE_ULN = 2;

    SynapseBridgeAdapter internal sba;

    string internal chainsConfig;
    uint256 internal longestChainLength;
    uint256 internal longestConfirmationsLength;

    ILayerZeroEndpointV2 internal endpoint;
    address internal receiveLibrary;
    address internal sendLibrary;
    mapping(address => SetConfigParam[]) internal setConfigParams;

    string[] internal allChains;
    mapping(string => uint32) internal eidByChainName;

    string internal dvnsConfig;

    string internal securityConfig;
    uint256 internal confirmationTimeMs;
    address[] internal requiredDVNs;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testWireSBA() external {}

    function run() external broadcastWithHooks {
        loadConfigs();
        address deployment = getDeploymentAddress({contractName: "SynapseBridgeAdapter", revertIfNotFound: true});
        sba = SynapseBridgeAdapter(deployment);
        // Every action below will be skipped if already done
        setPeers();
        setSendLibrary();
        setReceiveLibrary();
        setSendConfig();
        setReceiveConfig();
    }

    function formatChainName(string memory chain) internal view returns (string memory) {
        return string.concat(chain, ": ", string(" ").duplicate(longestChainLength - chain.length()));
    }

    function formatConfirmations(uint64 confirmations) internal view returns (string memory) {
        string memory str = vm.toString(confirmations);
        return string.concat(string(" ").duplicate(longestConfirmationsLength - str.length()), str);
    }

    function getChainConfirmations(string memory chain) internal view returns (uint64 confirmations) {
        if (confirmationTimeMs == 0) {
            printFailWithIndent("Confirmation time not set");
            assert(false);
        }
        uint256 blockTimeMs = chainsConfig.readUint(string.concat(".", chain, ".blockTime"));
        confirmations = uint64(confirmationTimeMs / blockTimeMs);
        // Round up to be a multiple of 10
        confirmations = (confirmations + 9) / 10 * 10;
    }

    function getUlnConfig(uint64 confirmations) internal view returns (UlnConfig memory) {
        return UlnConfig({
            confirmations: confirmations,
            requiredDVNCount: uint8(requiredDVNs.length),
            optionalDVNCount: 0,
            optionalDVNThreshold: 0,
            requiredDVNs: requiredDVNs,
            optionalDVNs: new address[](0)
        });
    }

    function loadConfigs() internal {
        dvnsConfig = readGlobalDeployProdConfig("dvns", true);

        securityConfig = readGlobalDeployProdConfig("security", true);
        confirmationTimeMs = readConfirmationTimeMs();
        string[] memory dvnNames = securityConfig.readStringArray(".DVNs");
        for (uint256 i = 0; i < dvnNames.length; ++i) {
            requiredDVNs.push(dvnsConfig.readAddress(string.concat(".", activeChain, ".", dvnNames[i])));
        }
        requiredDVNs = sortAddresses(requiredDVNs);

        chainsConfig = readGlobalDeployProdConfig("chains", true);
        if (!chainsConfig.keyExists(string.concat(".", activeChain))) {
            printFailWithIndent(string.concat("Chain ", activeChain, " not found in chains config"));
            assert(false);
        }
        allChains = vm.parseJsonKeys(chainsConfig, ".");
        for (uint256 i = 0; i < allChains.length; ++i) {
            uint256 chainLength = allChains[i].length();
            if (chainLength > longestChainLength) {
                longestChainLength = chainLength;
            }
            uint256 confirmationsLength = vm.toString(getChainConfirmations(allChains[i])).length();
            if (confirmationsLength > longestConfirmationsLength) {
                longestConfirmationsLength = confirmationsLength;
            }
            eidByChainName[allChains[i]] = uint32(chainsConfig.readUint(string.concat(".", allChains[i], ".eid")));
        }
        endpoint = ILayerZeroEndpointV2(chainsConfig.readAddress(string.concat(".", activeChain, ".endpointV2")));
        receiveLibrary = chainsConfig.readAddress(string.concat(".", activeChain, ".receiveUln302"));
        sendLibrary = chainsConfig.readAddress(string.concat(".", activeChain, ".sendUln302"));
    }

    function readConfirmationTimeMs() internal view returns (uint256) {
        uint256 confirmationTimeSeconds = securityConfig.readUint(".confirmationTimeSeconds");
        logTime(confirmationTimeSeconds, 1 days, "days") || logTime(confirmationTimeSeconds, 1 hours, "hours")
            || logTime(confirmationTimeSeconds, 1 minutes, "minutes")
            || logTime(confirmationTimeSeconds, 1 seconds, "seconds");
        return confirmationTimeSeconds * 1000;
    }

    function logTime(uint256 time, uint256 period, string memory unit) internal view returns (bool logged) {
        if (time % period == 0) {
            printInfo(string.concat("Confirmation time: ", vm.toString(time / period), " ", unit));
            return true;
        }
        return false;
    }

    function sortAddresses(address[] memory addresses) internal returns (address[] memory sorted) {
        uint256[] memory tmp = new uint256[](addresses.length);
        for (uint256 i = 0; i < addresses.length; ++i) {
            tmp[i] = uint256(uint160(addresses[i]));
        }
        tmp = vm.sort(tmp);
        sorted = new address[](addresses.length);
        for (uint256 i = 0; i < addresses.length; ++i) {
            sorted[i] = address(uint160(uint256(tmp[i])));
        }
    }

    function setPeers() internal {
        printLog("Setting peers...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            address remoteSBA =
                getDeploymentAddress({chain: chain, contractName: "SynapseBridgeAdapter", revertIfNotFound: false});
            if (remoteSBA == address(0)) {
                printSkipWithIndent(string.concat(chain, " doesn't have SBA deployed"));
                continue;
            }
            bytes32 peer = remoteSBA.toBytes32();
            uint32 eid = eidByChainName[chain];
            if (sba.peers(eid) == peer) {
                printSkipWithIndent(string.concat(chain, " already has peer set"));
                continue;
            }
            sba.setPeer(eid, peer);
            printSuccessWithIndent(string.concat(formatChainName(chain), vm.toString(remoteSBA)));
        }
    }

    function setSendLibrary() internal {
        printLog("Setting send library...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            uint32 eid = eidByChainName[chain];
            address curSendLibrary = endpoint.getSendLibrary(address(sba), eid);
            bool isDefault = endpoint.isDefaultSendLibrary(address(sba), eid);
            if (curSendLibrary == sendLibrary && !isDefault) {
                printSkipWithIndent(string.concat(chain, " already has send library set"));
                continue;
            }
            endpoint.setSendLibrary(address(sba), eid, sendLibrary);
            printSuccessWithIndent(string.concat(formatChainName(chain), vm.toString(sendLibrary)));
        }
    }

    function setReceiveLibrary() internal {
        printLog("Setting receive library...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            uint32 eid = eidByChainName[chain];
            (address curReceiveLibrary, bool isDefault) = endpoint.getReceiveLibrary(address(sba), eid);
            if (curReceiveLibrary == receiveLibrary && !isDefault) {
                printSkipWithIndent(string.concat(chain, " already has receive library set"));
                continue;
            }
            // Note: we don't need to use the grace period here
            endpoint.setReceiveLibrary(address(sba), eid, receiveLibrary, 0);
            printSuccessWithIndent(string.concat(formatChainName(chain), vm.toString(receiveLibrary)));
        }
    }

    function prepareUlnConfig(string memory chain, string memory libName, UlnConfig memory ulnConfig) internal {
        address lib = libName.equals("send") ? sendLibrary : libName.equals("receive") ? receiveLibrary : address(0);
        if (lib == address(0)) {
            printFailWithIndent(string.concat("Invalid library name: ", libName));
            assert(false);
        }
        uint32 eid = eidByChainName[chain];
        bytes memory curConfig =
            endpoint.getConfig({_oapp: address(sba), _lib: lib, _eid: eid, _configType: CONFIG_TYPE_ULN});
        if (keccak256(curConfig) == keccak256(abi.encode(ulnConfig))) {
            printSkipWithIndent(string.concat(formatChainName(chain), " already has ", libName, " config set"));
            return;
        }
        setConfigParams[lib].push(
            SetConfigParam({eid: eid, configType: CONFIG_TYPE_ULN, config: abi.encode(ulnConfig)})
        );
        printSuccessWithIndent(
            string.concat(
                formatChainName(chain),
                vm.toString(ulnConfig.requiredDVNCount),
                " DVNs, ",
                formatConfirmations(ulnConfig.confirmations),
                " confirmations"
            )
        );
    }

    function setSendConfig() internal {
        printLog("Setting send config...");
        // For send config the active chain confirmations are used
        uint64 confirmations = getChainConfirmations(activeChain);
        UlnConfig memory ulnConfig = getUlnConfig(confirmations);
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            prepareUlnConfig(chain, "send", ulnConfig);
        }
        if (setConfigParams[sendLibrary].length > 0) {
            endpoint.setConfig({_oapp: address(sba), _lib: sendLibrary, _params: setConfigParams[sendLibrary]});
        }
    }

    function setReceiveConfig() internal {
        printLog("Setting receive config...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            // For receive config the remote chain confirmations are used
            uint64 confirmations = getChainConfirmations(chain);
            UlnConfig memory ulnConfig = getUlnConfig(confirmations);
            prepareUlnConfig(chain, "receive", ulnConfig);
        }
        if (setConfigParams[receiveLibrary].length > 0) {
            endpoint.setConfig({_oapp: address(sba), _lib: receiveLibrary, _params: setConfigParams[receiveLibrary]});
        }
    }
}
