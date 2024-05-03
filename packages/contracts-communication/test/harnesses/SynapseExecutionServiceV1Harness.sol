// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseExecutionServiceV1} from "../../contracts/execution/SynapseExecutionServiceV1.sol";

// solhint-disable no-empty-blocks
/// @notice This harness is supposed to be used IN TESTS ONLY.
// DO NOT use this contract in production.
contract SynapseExecutionServiceV1Harness is SynapseExecutionServiceV1 {
    constructor() {
        // Grant the deployer all roles to simplify testing
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(GOVERNOR_ROLE, msg.sender);
    }

    function _disableInitializers() internal override {
        // No-op so that we can use the implementation w/o the proxy in Go tests
    }
}
