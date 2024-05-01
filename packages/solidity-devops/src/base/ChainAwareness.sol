// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {PathFinder} from "./PathFinder.sol";
import {StringUtils} from "../libs/StringUtils.sol";

import {VmSafe} from "forge-std/Vm.sol";

abstract contract ChainAwareness is PathFinder {
    using StringUtils for *;

    /// @notice Name of the active chain, should match the name of the directory in deployments
    string public activeChain;

    /// @dev Whether `chainIds` and `chainNames` have been loaded
    bool private chainIdsLoaded;
    /// @notice Map from chain name to chainId
    mapping(string => uint256) public chainIds;
    /// @notice Map from chainId to chain name
    mapping(uint256 => string) public chainNames;

    // ══════════════════════════════════════ ACTIVE CHAIN: FILE PATH GETTERS ══════════════════════════════════════════

    /// @notice Returns the path to the FRESH contract deployment JSON for a contract on the active chain.
    /// Example: ".deployments/mainnet/SynapseRouter.json"
    function getFreshDeploymentFN(string memory contractName) internal view returns (string memory) {
        return getFreshDeploymentFN(activeChain, contractName);
    }

    /// @notice Returns the path to the contract deployment JSON for a contract on the active chain.
    /// Example: "deployments/mainnet/SynapseRouter.json"
    function getDeploymentFN(string memory contractName) internal view returns (string memory) {
        return getDeploymentFN(activeChain, contractName);
    }

    /// @notice Returns the path to the generic file on the active chain.
    /// @dev Useful for the files that are not specific to a contract, but are specific to a chain.
    function getChainGenericFN(string memory fileName) internal view returns (string memory) {
        return getChainGenericFN(activeChain, fileName);
    }

    /// @notice Returns the path to the contract deployment configuration JSON for a contract on the active chain.
    /// Example: "script/configs/mainnet/SynapseRouter.json"
    function getDeployConfigFN(string memory contractName) internal view returns (string memory) {
        return getDeployConfigFN(activeChain, contractName);
    }

    // ══════════════════════════════════════════════ CHAIN ID UTILS ═══════════════════════════════════════════════════

    /// @notice Returns the chain ID for a given chain by reading the chain ID file in the deployments directory.
    /// @dev Returns 0 or reverts if the chain ID is not found based on the flag.
    function getChainId(string memory chain, bool revertIfNotFound) internal returns (uint256) {
        string memory chainIdFile = getDeploymentsPath().concat(chain, "/.chainId");
        // Read only the first line of the file
        try vm.readLine(chainIdFile) returns (string memory chainId) {
            vm.closeFile(chainIdFile);
            return chainId.toUint();
        } catch {
            if (revertIfNotFound) revert("ChainAwareness: chain ID not found");
            return 0;
        }
    }

    /// @notice Sets active chain to the one matching block.chainid value.
    /// Reverts if the chain is not supported.
    function loadActiveChain() internal {
        loadChainIds();
        uint256 chainId = blockChainId();
        activeChain = chainNames[chainId];
        require(
            activeChain.length() > 0, StringUtils.concat("ChainAwareness: unsupported chain ID: ", chainId.fromUint())
        );
    }

    /// @notice Reads all chain:chainId pairs from the deployments directory,
    /// and saves them in `chainIds` and `chainNames` mappings.
    /// @dev Will do this lazily, only once per the script run.
    function loadChainIds() internal {
        if (chainIdsLoaded) return;
        // Read all entries in the deployments directory
        VmSafe.DirEntry[] memory entries = vm.readDir(getDeploymentsPath());
        // Iterate over all entries that are directories
        for (uint256 i = 0; i < entries.length; i++) {
            if (!entries[i].isDir) {
                continue;
            }
            // Extract chain name from the path: everything that comes after the last slash
            uint256 lastSlash = entries[i].path.lastIndexOf("/");
            require(
                lastSlash != StringUtils.NOT_FOUND,
                StringUtils.concat("ChainAwareness: invalid path: ", entries[i].path)
            );

            // Chain name is everything that comes after the last slash
            string memory chainName = entries[i].path.suffix(lastSlash + 1);
            // Only save the chain if it has a valid chainId
            uint256 chainId = getChainId({chain: chainName, revertIfNotFound: false});
            if (chainId == 0) continue;
            // Sanity check that no duplicates are saved
            require(chainIds[chainName] == 0, StringUtils.concat("ChainAwareness: duplicate chain name: ", chainName));
            require(
                chainNames[chainId].length() == 0,
                StringUtils.concat("ChainAwareness: duplicate chain ID: ", chainId.fromUint())
            );
            chainIds[chainName] = chainId;
            chainNames[chainId] = chainName;
        }
        chainIdsLoaded = true;
    }

    /// @notice Wrapper for block.chainid, which only exists in Solidity 0.8+
    function blockChainId() internal view returns (uint256 chainId) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            chainId := chainid()
        }
    }
}
