// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {StringUtils, SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

// solhint-disable custom-errors
contract TopUpGasScript is SynapseScript {
    using Address for address payable;
    using StringUtils for *;

    uint256 public constant MINIMUM_BALANCE = 1 ether;

    function run(string memory contractAlias) external payable broadcastWithHooks {
        address deployment = getDeploymentAddress({contractName: contractAlias, revertIfNotFound: true});
        string memory desc = MINIMUM_BALANCE.fromWei().concat(" ETH");
        if (deployment.balance >= MINIMUM_BALANCE) {
            printSkipWithIndent(contractAlias.concat(" already has at least ", desc));
            return;
        }
        payable(deployment).sendValue(MINIMUM_BALANCE);
        printSuccessWithIndent(contractAlias.concat(" topped up with ", desc));
    }
}
