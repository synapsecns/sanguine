// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {UlnConfig} from "@layerzerolabs/lz-evm-messagelib-v2/contracts/uln/UlnBase.sol";
import {ILayerZeroEndpointV2 as ILZEV2} from
    "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/ILayerZeroEndpointV2.sol";
import {IMessageLib} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLib.sol";
import {
    IMessageLibManager,
    SetConfigParam
} from "@layerzerolabs/lz-evm-protocol-v2/contracts/interfaces/IMessageLibManager.sol";
import {AddressCast} from "@layerzerolabs/lz-evm-protocol-v2/contracts/libs/AddressCast.sol";
import {IOAppCore} from "@layerzerolabs/oapp-evm/contracts/oapp/interfaces/IOAppCore.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {StringUtils, SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

interface ILayerZeroEndpointV2 is ILZEV2 {
    function delegates(address _oapp) external view returns (address);
}

abstract contract LayerZeroWiring is SynapseScript {
    using AddressCast for address;
    using stdJson for string;
    using StringUtils for string;

    uint32 internal constant CONFIG_TYPE_ULN = 2;

    IOAppCore internal app;
    string internal deploymentName;

    string internal chainsConfig;
    string internal chainsConfigRoot;
    uint256 internal longestChainLength;
    uint256 internal longestConfirmationsLength;

    ILayerZeroEndpointV2 internal endpoint;
    address internal receiveLibrary;
    address internal sendLibrary;
    mapping(address => SetConfigParam[]) internal setConfigParams;
    bool internal printMultisigTxs;
    bool internal printPeerMultisigTxs;

    string[] internal allChains;
    mapping(string => uint32) internal eidByChainName;
    mapping(uint32 => bool) internal eidSkipped;

    string internal dvnsConfig;

    string internal securityConfig;
    string internal securityConfigRoot;
    address[] internal requiredDVNs;

    function loadConfigs() internal virtual {
        dvnsConfig = readGlobalDeployProdConfig("dvns", true);

        securityConfig = readGlobalDeployProdConfig("security", true);
        string[] memory dvnNames = securityConfig.readStringArray(".DVNs");
        for (uint256 i = 0; i < dvnNames.length; ++i) {
            requiredDVNs.push(dvnsConfig.readAddress(string.concat(".", activeChain, ".", dvnNames[i])));
        }
        requiredDVNs = sortAddresses(requiredDVNs);

        chainsConfig = readGlobalDeployProdConfig("chains", true);
        if (!chainsConfig.keyExists(chainConfigPath(activeChain, ""))) {
            printFailWithIndent(string.concat("Chain ", activeChain, " not found in chains config"));
            assert(false);
        }
        allChains = vm.parseJsonKeys(chainsConfig, ".");
        loadChainsConfig(chainsConfig.readAddress(chainConfigPath(activeChain, ".endpointV2")));
    }

    function loadChainsConfig(address endpointAddress) internal {
        if (!chainsConfig.keyExists(chainConfigPath(activeChain, ""))) {
            printFailWithIndent(string.concat("Chain ", activeChain, " not found in chains config"));
            assert(false);
        }
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            uint256 chainLength = chain.length();
            if (chainLength > longestChainLength) {
                longestChainLength = chainLength;
            }
            uint256 confirmationsLength = vm.toString(getChainConfirmations(chain)).length();
            if (confirmationsLength > longestConfirmationsLength) {
                longestConfirmationsLength = confirmationsLength;
            }
            eidByChainName[chain] = uint32(chainsConfig.readUint(chainConfigPath(chain, ".eid")));
        }
        endpoint = ILayerZeroEndpointV2(endpointAddress);
        receiveLibrary = chainsConfig.readAddress(chainConfigPath(activeChain, ".receiveUln302"));
        sendLibrary = chainsConfig.readAddress(chainConfigPath(activeChain, ".sendUln302"));
    }

    function wireApp(address appAddress, string memory name) internal {
        app = IOAppCore(appAddress);
        deploymentName = name;
        address appEndpoint = address(app.endpoint());
        if (appEndpoint != address(endpoint)) {
            printFailWithIndent(
                string.concat(
                    "Endpoint mismatch: config has ",
                    vm.toString(address(endpoint)),
                    " but app has ",
                    vm.toString(appEndpoint)
                )
            );
            assert(false);
        }
        loadSkippedEids();
        setPrintMultisigTxs();
        setPeers();
        setSendLibrary();
        setReceiveLibrary();
        setSendConfig();
        setReceiveConfig();
    }

    function loadSkippedEids() internal {
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            uint32 eid = eidByChainName[chain];
            if (!IMessageLib(sendLibrary).isSupportedEid(eid) || !IMessageLib(receiveLibrary).isSupportedEid(eid)) {
                eidSkipped[eid] = true;
                printLogWithIndent(
                    string.concat(unicode"⚠️ ", chain, " is not supported (eid: ", vm.toString(eid), ")")
                );
            }
        }
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

    function setPrintMultisigTxs() internal {
        address owner = Ownable(address(app)).owner();
        if (msg.sender != owner) {
            printPeerMultisigTxs = true;
            printInfo(
                "Broadcast wallet is not the app owner, printing peer multisig calldata instead of submitting txs"
            );
            printAuthorityInfo(owner, "Owner");
        }

        address delegate = endpoint.delegates(address(app));
        if (msg.sender != delegate) {
            printMultisigTxs = true;
            printInfo("Wallet is not the app delegate; printing endpoint multisig calldata instead of submitting txs");
            printAuthorityInfo(delegate, "Delegate");
        }
    }

    function setPeers() internal {
        printLog("Setting peers...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            address remoteApp =
                getDeploymentAddress({chain: chain, contractName: deploymentName, revertIfNotFound: false});
            if (remoteApp == address(0)) {
                printSkipWithIndent(string.concat(chain, " doesn't have ", deploymentName, " deployed"));
                continue;
            }
            bytes32 peer = remoteApp.toBytes32();
            uint32 eid = eidByChainName[chain];
            if (eidSkipped[eid]) {
                printSkipWithIndent(string.concat(chain, " is not supported"));
                continue;
            }
            if (app.peers(eid) == peer) {
                printSkipWithIndent(string.concat(chain, " already has peer set"));
                continue;
            }
            if (printPeerMultisigTxs) {
                printMultisigTx(
                    string.concat(formatChainName(chain), "set peer"),
                    address(app),
                    abi.encodeCall(IOAppCore.setPeer, (eid, peer))
                );
                continue;
            }
            app.setPeer(eid, peer);
            printSuccessWithIndent(string.concat(formatChainName(chain), vm.toString(remoteApp)));
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
            if (eidSkipped[eid]) {
                printSkipWithIndent(string.concat(chain, " is not supported"));
                continue;
            }
            address curSendLibrary = endpoint.getSendLibrary(address(app), eid);
            bool isDefault = endpoint.isDefaultSendLibrary(address(app), eid);
            if (curSendLibrary == sendLibrary && !isDefault) {
                printSkipWithIndent(string.concat(chain, " already has send library set"));
                continue;
            }
            if (printMultisigTxs) {
                printMultisigTx(
                    string.concat(formatChainName(chain), "set send library"),
                    address(endpoint),
                    abi.encodeCall(IMessageLibManager.setSendLibrary, (address(app), eid, sendLibrary))
                );
                continue;
            }
            endpoint.setSendLibrary(address(app), eid, sendLibrary);
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
            if (eidSkipped[eid]) {
                printSkipWithIndent(string.concat(chain, " is not supported"));
                continue;
            }
            (address curReceiveLibrary, bool isDefault) = endpoint.getReceiveLibrary(address(app), eid);
            if (curReceiveLibrary == receiveLibrary && !isDefault) {
                printSkipWithIndent(string.concat(chain, " already has receive library set"));
                continue;
            }
            if (printMultisigTxs) {
                printMultisigTx(
                    string.concat(formatChainName(chain), "set receive library"),
                    address(endpoint),
                    abi.encodeCall(IMessageLibManager.setReceiveLibrary, (address(app), eid, receiveLibrary, 0))
                );
                continue;
            }
            endpoint.setReceiveLibrary(address(app), eid, receiveLibrary, 0);
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
        if (eidSkipped[eid]) {
            printSkipWithIndent(string.concat(chain, " is not supported"));
            return;
        }
        bytes memory curConfig = getCurrentUlnConfig(lib, eid);
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
            SetConfigParam[] memory params = setConfigParams[sendLibrary];
            if (printMultisigTxs) {
                printMultisigTx(
                    "set send config",
                    address(endpoint),
                    abi.encodeCall(IMessageLibManager.setConfig, (address(app), sendLibrary, params))
                );
                return;
            }
            endpoint.setConfig({_oapp: address(app), _lib: sendLibrary, _params: params});
        }
    }

    function setReceiveConfig() internal {
        printLog("Setting receive config...");
        for (uint256 i = 0; i < allChains.length; ++i) {
            string memory chain = allChains[i];
            if (chain.equals(activeChain)) {
                continue;
            }
            uint64 confirmations = getChainConfirmations(chain);
            UlnConfig memory ulnConfig = getUlnConfig(confirmations);
            prepareUlnConfig(chain, "receive", ulnConfig);
        }
        if (setConfigParams[receiveLibrary].length > 0) {
            SetConfigParam[] memory params = setConfigParams[receiveLibrary];
            if (printMultisigTxs) {
                printMultisigTx(
                    "set receive config",
                    address(endpoint),
                    abi.encodeCall(IMessageLibManager.setConfig, (address(app), receiveLibrary, params))
                );
                return;
            }
            endpoint.setConfig({_oapp: address(app), _lib: receiveLibrary, _params: params});
        }
    }

    function chainConfigPath(string memory chain, string memory suffix) internal view returns (string memory) {
        return string.concat(chainsConfigRoot, ".", chain, suffix);
    }

    function formatChainName(string memory chain) internal view returns (string memory) {
        return string.concat(chain, ": ", string(" ").duplicate(longestChainLength - chain.length()));
    }

    function formatConfirmations(uint64 confirmations) internal view returns (string memory) {
        string memory str = vm.toString(confirmations);
        return string.concat(string(" ").duplicate(longestConfirmationsLength - str.length()), str);
    }

    function getChainConfirmations(string memory chain) internal view returns (uint64 confirmations) {
        string memory configPath = string.concat(securityConfigRoot, ".blockConfirmations.", chain);
        if (!securityConfig.keyExists(configPath)) {
            printFailWithIndent(string.concat("Block confirmations not set for chain: ", chain));
            assert(false);
        }
        confirmations = uint64(securityConfig.readUint(configPath));
    }

    function getUlnConfig(uint64 confirmations) internal view virtual returns (UlnConfig memory) {
        return UlnConfig({
            confirmations: confirmations,
            requiredDVNCount: uint8(requiredDVNs.length),
            optionalDVNCount: 0,
            optionalDVNThreshold: 0,
            requiredDVNs: requiredDVNs,
            optionalDVNs: new address[](0)
        });
    }

    function getCurrentUlnConfig(address lib, uint32 eid) internal view virtual returns (bytes memory) {
        return endpoint.getConfig({_oapp: address(app), _lib: lib, _eid: eid, _configType: CONFIG_TYPE_ULN});
    }

    function printAuthorityInfo(address authority, string memory role) internal view {
        printLogWithIndent(string.concat("App:         ", vm.toString(address(app))));
        printLogWithIndent(string.concat(role, ":", string(" ").duplicate(12 - role.length()), vm.toString(authority)));
        printLogWithIndent(string.concat("Broadcaster: ", vm.toString(msg.sender)));
    }

    function printMultisigTx(string memory action, address to, bytes memory data) internal view {
        printLogWithIndent(action);
        printLogWithIndent(string.concat("to:   ", vm.toString(to)));
        printLogWithIndent(string.concat("data: ", vm.toString(data)));
    }
}
