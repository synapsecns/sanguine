// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseComposer} from "../src/SynapseComposer.sol";
import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";

import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Deploys a token's HyperCore composer using its saved OFT adapter and global configuration.
contract DeploySynapseComposer is SynapseScript {
    using stdJson for string;

    string public constant NAME = "SynapseComposer";

    /// @notice Excludes this script from coverage reports.
    // solhint-disable-next-line no-empty-blocks
    function testDeploySynapseComposer() external {}

    /// @param tokenId The token identifier used in the config and deployment alias, e.g. SYN.
    function run(string memory tokenId) external broadcastWithHooks {
        require(block.chainid == 999 || block.chainid == 998, "Composer requires HyperEVM");
        require(bytes(tokenId).length > 0, "Token identifier is empty");
        string memory environment = vm.envOr("DEPLOY_ENVIRONMENT", string("production"));
        if (keccak256(bytes(environment)) == keccak256("production")) environment = ENVIRONMENT_PROD;
        string memory config = readGlobalDeployConfig("SynapseOFTAdapter", environment, true);
        string memory tokenPath = string.concat(".tokens.", tokenId);
        string memory composerPath = string.concat(tokenPath, ".composer");
        SynapseOFTAdapter adapter =
            SynapseOFTAdapter(getDeploymentAddress(string.concat("SynapseOFTAdapter.", tokenId), true));
        require(address(adapter).code.length > 0, "Adapter is not a contract");
        require(
            adapter.token() == config.readAddress(string.concat(tokenPath, ".addresses.", activeChain)),
            "Adapter token mismatch"
        );
        require(
            address(adapter.endpoint()) == config.readAddress(string.concat(".endpoints.", activeChain)),
            "Adapter endpoint mismatch"
        );

        uint64 coreIndexId = SafeCast.toUint64(config.readUint(string.concat(composerPath, ".coreIndexId")));
        int8 assetDecimalDiff = SafeCast.toInt8(config.readInt(string.concat(composerPath, ".assetDecimalDiff")));
        address recoveryAddress = config.readAddress(string.concat(composerPath, ".recoveryAddress"));
        require(assetDecimalDiff >= -2 && assetDecimalDiff <= 18, "Invalid asset decimal difference");
        require(recoveryAddress != address(0), "Recovery address is zero");

        bytes memory constructorArgs = abi.encode(address(adapter), coreIndexId, assetDecimalDiff, recoveryAddress);
        printInfo(string.concat("OFT adapter: ", vm.toString(address(adapter))));
        printInfo(string.concat("Recovery address: ", vm.toString(recoveryAddress)));
        SynapseComposer composer = SynapseComposer(
            payable(
                deployAndSaveAs({
                    contractName: NAME,
                    contractAlias: string.concat(NAME, ".", tokenId),
                    constructorArgs: constructorArgs,
                    deployCodeFunc: cbDeploy
                })
            )
        );
        require(composer.OFT() == address(adapter), "Composer adapter mismatch");
        require(composer.ERC20_CORE_INDEX_ID() == coreIndexId, "Composer Core index mismatch");
        require(composer.ERC20_DECIMAL_DIFF() == assetDecimalDiff, "Composer decimals mismatch");
        require(composer.RECOVERY_ADDRESS() == recoveryAddress, "Composer recovery mismatch");
    }
}
