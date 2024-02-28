// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {BasicContract} from "./BasicContract.sol";
import {SynapseScript} from "../src/SynapseScript.sol";

contract Script is SynapseScript {
    function run() external broadcastWithHooks {
        payable(msg.sender).transfer(0);
        deployAndSave("BasicContract", cbDeployBasicContract);
    }

    function cbDeployBasicContract() internal returns (address deployedAt, bytes memory constructorArgs) {
        constructorArgs = abi.encode(42);
        deployedAt = address(new BasicContract(42));
    }
}
