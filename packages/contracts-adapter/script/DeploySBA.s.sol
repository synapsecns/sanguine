// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseBridgeAdapter} from "../src/SynapseBridgeAdapter.sol";

import {SynapseScript} from "@synapsecns/solidity-devops/src/SynapseScript.sol";

// solhint-disable no-empty-blocks
contract DeploySBA is SynapseScript {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testDeploySBA() external {}

    function run() external broadcastWithHooks {
        // TODO: get endpoint address for chain
    }
}
