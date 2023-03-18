// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";

import { SystemRouterMock } from "../mocks/system/SystemRouterMock.t.sol";

/// @notice Harness for standalone Go tests.
/// Do not use for tests requiring interactions between messaging contracts.
contract DestinationHarness is Destination {
    constructor(uint32 _domain) Destination(_domain) {
        // Add Mock for SystemRouter for standalone tests
        systemRouter = new SystemRouterMock();
    }

    /// @notice Adding agents in Go tests
    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _addAgent(_domain, _account);
    }

    /// @notice Removing agents in Go tests
    function removeAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _removeAgent(_domain, _account);
    }
}
