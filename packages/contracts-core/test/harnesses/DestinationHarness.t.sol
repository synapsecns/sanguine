// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Destination } from "../../contracts/Destination.sol";

// TODO: remove/adapt when "go generate" is updated
contract DestinationHarness is Destination {
    constructor(uint32 _domain) Destination(_domain) {}

    /// @notice Adding agents in Go tests
    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _addAgent(_domain, _account);
    }

    /// @notice Removing agents in Go tests
    function removeAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _removeAgent(_domain, _account);
    }
}
