// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {StringUtils, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

interface IWithdrawable {
    function withdraw() external;
}

// solhint-disable custom-errors
contract WithdrawGasScript is SynapseScript {
    using StringUtils for *;

    function run(string memory contractAlias) external broadcastWithHooks {
        address deployment = getDeploymentAddress({contractName: contractAlias, revertIfNotFound: true});
        uint256 balanceBefore = msg.sender.balance;
        IWithdrawable(deployment).withdraw();
        uint256 delta = msg.sender.balance - balanceBefore;
        if (delta == 0) {
            revert("No ETH was withdrawn");
        }
        printSuccessWithIndent(string.concat("Withdrawn ", delta.fromWei(), " ETH"));
    }
}
