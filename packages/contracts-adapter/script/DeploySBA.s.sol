// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter, SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";

import {StringUtils, SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

interface IBridge {
    function bridgeVersion() external view returns (uint256);
}

// solhint-disable code-complexity, no-empty-blocks
/// @notice This script deploys a SynapseBridgeAdapter contract and performs its initial configuration.
/// The LayerZero-specific configuration is done in a separate script, as it requires SBA to be deployed on all chains.
contract DeploySBA is SynapseScript {
    using stdJson for string;
    using StringUtils for string;

    bytes32 internal constant SALT = 0xcabb90224154e99298f486ccf31bda85ac9727771f7f19dbd8e19644c6491fc5;

    SynapseBridgeAdapter internal sba;

    string internal chainsConfig;
    uint256 internal longestChainName;

    string internal tokensConfig;
    string[] internal unsupportedTokenIDs;

    mapping(string => ISynapseBridgeAdapter.RemoteToken[]) internal remoteTokens;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySBA() external {}

    function run() external broadcastWithHooks {
        loadConfigs();
        // Every action below will be skipped if already done
        sba = SynapseBridgeAdapter(deployAdapter());
        setBridge();
        // addTokens();
    }

    function loadConfigs() internal {
        // GMX will have its own independent migration process
        unsupportedTokenIDs.push("GMX");

        chainsConfig = readGlobalDeployProdConfig("chains", true);
        tokensConfig = readGlobalDeployProdConfig("tokens", true);

        if (!chainsConfig.keyExists(string.concat(".", activeChain))) {
            printFailWithIndent(string.concat("Chain ", activeChain, " not found in chains config"));
            assert(false);
        }
        string[] memory chainNames = vm.parseJsonKeys(chainsConfig, ".");
        for (uint256 i = 0; i < chainNames.length; ++i) {
            if (chainNames[i].length() > longestChainName) {
                longestChainName = chainNames[i].length();
            }
        }
    }

    function deployAdapter() internal returns (address deployment) {
        address endpointV2 = chainsConfig.readAddress(string.concat(".", activeChain, ".endpointV2"));
        // constructor(address endpoint_, address owner_)
        bytes memory constructorArgs = abi.encode(endpointV2, msg.sender);
        printInfo(string.concat("EndpointV2:    ", vm.toString(endpointV2)));
        printInfo(string.concat("Initial Owner: ", vm.toString(msg.sender)));
        // Note: this will skip the deployment if it already exists
        setNextDeploymentSalt(SALT);
        deployment = deployAndSave({
            contractName: "SynapseBridgeAdapter",
            constructorArgs: constructorArgs,
            deployCodeFunc: cbDeployCreate2
        });
    }

    function setBridge() internal {
        printLog("Setting bridge...");
        address bridge = chainsConfig.readAddress(string.concat(".", activeChain, ".synapseBridge"));
        printInfo(bridge, "Bridge");
        address implementation = address(
            uint160(uint256(vm.load(bridge, 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc)))
        );
        printInfo(implementation, "Implementation");
        address curBridge = sba.bridge();
        if (curBridge == address(0)) {
            sba.setBridge(bridge);
            printSuccessWithIndent(string.concat("Bridge set to ", vm.toString(bridge)));
            return;
        }
        if (curBridge == bridge) {
            printSkipWithIndent(string.concat("bridge already set to ", vm.toString(bridge)));
            return;
        }
        printFailWithIndent(
            string.concat(
                "Bridge address mismatch: expected ", vm.toString(bridge), " but got ", vm.toString(curBridge)
            )
        );
        assert(false);
    }

    function printInfo(address addr, string memory name) internal {
        increaseIndent();
        printLog(string.concat(name, ": ", vm.toString(addr)));
        if (addr.code.length == 0) {
            printFailWithIndent("Not a contract");
            assert(false);
        }
        printLogWithIndent(string.concat("Version: ", vm.toString(IBridge(addr).bridgeVersion())));
        printLogWithIndent(string.concat("Code length: ", vm.toString(addr.code.length)));
        printLogWithIndent(string.concat("Code hash: ", vm.toString(keccak256(addr.code))));
        decreaseIndent();
    }

    function addTokens() internal {
        printLog("Adding tokens...");
        increaseIndent();
        string[] memory tokenIDs = vm.parseJsonKeys(tokensConfig, ".");
        for (uint256 i = 0; i < tokenIDs.length; ++i) {
            addToken(tokenIDs[i]);
        }
        decreaseIndent();
    }

    function addToken(string memory tokenID) internal {
        printLog(tokenID);
        if (remoteTokens[tokenID].length > 0) {
            printFailWithIndent("Duplicate token ID");
            assert(false);
        }
        string[] memory chains = vm.parseJsonKeys(tokensConfig, string.concat(".", tokenID));
        if (!contains(chains, activeChain)) {
            printSkipWithIndent(string.concat("not deployed on ", activeChain));
            return;
        }
        if (chains.length < 2) {
            printSkipWithIndent("not deployed on multiple chains");
            return;
        }
        if (contains(unsupportedTokenIDs, tokenID)) {
            printSkipWithIndent("will not be supported by SBA");
            return;
        }
        address localToken = tokensConfig.readAddress(string.concat(".", tokenID, ".", activeChain, ".tokenAddress"));
        bool isUnderlying = tokensConfig.readBool(string.concat(".", tokenID, ".", activeChain, ".isUnderlying"));
        ISynapseBridgeAdapter.TokenType tokenType =
            isUnderlying ? ISynapseBridgeAdapter.TokenType.WithdrawDeposit : ISynapseBridgeAdapter.TokenType.MintBurn;
        printInfo(
            string.concat(
                activeChain, ": ", vm.toString(localToken), " ", isUnderlying ? "WithdrawDeposit" : "MintBurn"
            )
        );
        for (uint256 i = 0; i < chains.length; ++i) {
            string memory remoteChain = chains[i];
            if (remoteChain.equals(activeChain)) continue;
            uint32 remoteEid = uint32(chainsConfig.readUint(string.concat(".", remoteChain, ".eid")));
            address remoteToken =
                tokensConfig.readAddress(string.concat(".", tokenID, ".", remoteChain, ".tokenAddress"));
            address curRemoteToken = sba.getRemoteAddress(remoteEid, localToken);
            if (curRemoteToken == address(0)) {
                remoteTokens[tokenID].push(ISynapseBridgeAdapter.RemoteToken(remoteEid, remoteToken));
                printSuccessWithIndent(
                    string.concat(
                        remoteChain,
                        ": ",
                        string(" ").duplicate(longestChainName - remoteChain.length()),
                        vm.toString(remoteToken)
                    )
                );
            } else if (curRemoteToken != remoteToken) {
                printFailWithIndent(string.concat("remote token mismatch for ", remoteChain));
                assert(false);
            }
        }
        if (remoteTokens[tokenID].length == 0) {
            printSkipWithIndent("no remote tokens to add");
            return;
        }
        sba.addToken(localToken, tokenType, remoteTokens[tokenID]);
    }

    function contains(string[] memory arr, string memory str) internal pure returns (bool) {
        for (uint256 i = 0; i < arr.length; ++i) {
            if (arr[i].equals(str)) {
                return true;
            }
        }
        return false;
    }
}
