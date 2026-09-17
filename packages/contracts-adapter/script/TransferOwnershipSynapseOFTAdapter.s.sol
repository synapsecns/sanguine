// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";

import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

interface IEndpointDelegate {
    function delegates(address app) external view returns (address);
}

/// @notice Transfers an OFT adapter's endpoint delegate and ownership to the configured multisig.
contract TransferOwnershipSynapseOFTAdapter is SynapseScript {
    using stdJson for string;

    /// @notice Excludes this script from coverage reports.
    // solhint-disable-next-line no-empty-blocks
    function testTransferOwnershipSynapseOFTAdapter() external {}

    /// @param tokenId The token identifier used in the config and deployment alias, e.g. SYN.
    function run(string memory tokenId) external broadcastWithHooks {
        require(bytes(tokenId).length > 0, "Token identifier is empty");
        string memory environment = vm.envOr("DEPLOY_ENVIRONMENT", string("production"));
        if (keccak256(bytes(environment)) == keccak256("production")) environment = ENVIRONMENT_PROD;
        string memory config = readGlobalDeployConfig("SynapseOFTAdapter", environment, true);
        address multisig = config.readAddress(string.concat(".multisig.", activeChain));
        require(multisig != address(0), "Multisig is zero address");
        require(multisig.code.length > 0, "Multisig is not a contract");

        SynapseOFTAdapter adapter =
            SynapseOFTAdapter(getDeploymentAddress(string.concat("SynapseOFTAdapter.", tokenId), true));
        require(
            adapter.token() == config.readAddress(string.concat(".tokens.", tokenId, ".addresses.", activeChain)),
            "Adapter token mismatch"
        );
        address endpoint = address(adapter.endpoint());
        require(endpoint == config.readAddress(string.concat(".endpoints.", activeChain)), "Adapter endpoint mismatch");
        address owner = adapter.owner();
        address delegate = IEndpointDelegate(endpoint).delegates(address(adapter));
        printInfo(string.concat("Adapter: ", vm.toString(address(adapter))));
        printInfo(string.concat("Current owner: ", vm.toString(owner)));
        printInfo(string.concat("Current delegate: ", vm.toString(delegate)));
        printInfo(string.concat("Target multisig: ", vm.toString(multisig)));
        if (owner == multisig && delegate == multisig) {
            printSkipWithIndent("Ownership and delegate already transferred");
            return;
        }
        require(owner == msg.sender, "Wallet does not own adapter");

        // setDelegate is owner-only, so update it before transferring adapter ownership.
        if (delegate != multisig) {
            adapter.setDelegate(multisig);
            printSuccessWithIndent("Endpoint delegate transferred");
        }
        require(IEndpointDelegate(endpoint).delegates(address(adapter)) == multisig, "Delegate transfer failed");
        if (owner != multisig) {
            adapter.transferOwnership(multisig);
            printSuccessWithIndent("Adapter ownership transferred");
        }
        require(adapter.owner() == multisig, "Ownership transfer failed");
    }
}
