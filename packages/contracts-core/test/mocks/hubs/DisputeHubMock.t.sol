// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DisputeStatus, IDisputeHub} from "../../../contracts/interfaces/IDisputeHub.sol";

// solhint-disable no-empty-blocks
contract DisputeHubMock is IDisputeHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testDisputeHubMock() external {}

    function openDispute(address guard, uint32 domain, address notary) external {}

    function disputeStatus(address agent) external view returns (DisputeStatus memory status) {}
}
