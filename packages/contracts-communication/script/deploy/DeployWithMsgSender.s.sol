// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

/// @notice Deploys a contract that takes a single address argument in its constructor.
/// The address is the sender of the transaction that deploys the contract.
contract DeployWithMsgSender is SynapseScript {
    function run(string memory contractName) external broadcastWithHooks {
        bytes memory constructorArgs = abi.encode(msg.sender);
        deployAndSave(contractName, constructorArgs, cbDeploy);
    }
}
