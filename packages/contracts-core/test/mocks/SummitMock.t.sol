// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, InterfaceSummit} from "../../contracts/interfaces/InterfaceSummit.sol";
import {SnapshotHubMock} from "./hubs/SnapshotHubMock.t.sol";
import {SystemRegistryMock} from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract SummitMock is SnapshotHubMock, SystemRegistryMock, InterfaceSummit {
    /// @notice Prevents this contract from being included in the coverage report
    function testSummitMock() external {}

    function acceptReceipt(
        address notary,
        AgentStatus memory status,
        bytes memory rcptPayload,
        bytes memory rcptSignature
    ) external returns (bool wasAccepted) {}

    function acceptSnapshot(
        address agent,
        AgentStatus memory status,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bytes memory attPayload) {}

    function distributeTips() external returns (bool queuePopped) {}

    function withdrawTips(uint32 origin, uint256 amount) external {}

    function actorTips(address actor, uint32 origin) external view returns (uint128 earned, uint128 claimed) {}

    function receiptQueueLength() external view returns (uint256) {}

    function getLatestState(uint32 origin) external view returns (bytes memory statePayload) {}

    function getSignedSnapshot(uint256 nonce)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {}
}
