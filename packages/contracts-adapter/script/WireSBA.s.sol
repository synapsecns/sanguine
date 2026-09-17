// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {LayerZeroWiring} from "./helpers/LayerZeroWiring.sol";

// solhint-disable no-empty-blocks
contract WireSBA is LayerZeroWiring {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testWireSBA() external {}

    function run() external broadcastWithHooks {
        loadConfigs();
        address deployment = getDeploymentAddress({contractName: "SynapseBridgeAdapter", revertIfNotFound: true});
        wireApp(deployment, "SynapseBridgeAdapter");
    }
}
