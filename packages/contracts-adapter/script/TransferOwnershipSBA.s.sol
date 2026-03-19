// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";
import {SynapseScript, stdJson} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract TransferOwnershipSBA is SynapseScript {
    using stdJson for string;

    SynapseBridgeAdapter internal sba;
    string internal multisigConfig;

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testTransferOwnershipSBA() external {}

    function run() external broadcastWithHooks {
        multisigConfig = readGlobalDeployProdConfig("multisig", true);
        string memory configPath = string.concat(".", activeChain);
        if (!multisigConfig.keyExists(configPath)) {
            printFailWithIndent(string.concat("Multisig not set for chain: ", activeChain));
            assert(false);
        }
        address newOwner = multisigConfig.readAddress(configPath);
        if (newOwner == address(0)) {
            printFailWithIndent(string.concat("Multisig is zero address for chain: ", activeChain));
            assert(false);
        }

        address deployment = getDeploymentAddress({contractName: "SynapseBridgeAdapter", revertIfNotFound: true});
        sba = SynapseBridgeAdapter(deployment);

        address currentOwner = sba.owner();
        printInfo(string.concat("Current owner: ", vm.toString(currentOwner)));
        printInfo(string.concat("Target owner:  ", vm.toString(newOwner)));

        if (currentOwner == newOwner) {
            printSkipWithIndent("SynapseBridgeAdapter already owned by target multisig");
            return;
        }
        if (currentOwner != msg.sender) {
            printFailWithIndent(
                string.concat(
                    "Broadcast wallet is not the current owner: ",
                    vm.toString(msg.sender),
                    " != ",
                    vm.toString(currentOwner)
                )
            );
            assert(false);
        }

        printLog("Transferring ownership...");
        sba.transferOwnership(newOwner);
        printSuccessWithIndent(string.concat("SynapseBridgeAdapter owner set to ", vm.toString(newOwner)));
    }
}
