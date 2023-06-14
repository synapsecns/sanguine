// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, InterfaceSummit} from "../../contracts/interfaces/InterfaceSummit.sol";
import {SnapshotHubMock} from "./hubs/SnapshotHubMock.t.sol";
import {AgentSecuredMock} from "./base/AgentSecuredMock.t.sol";

// solhint-disable no-empty-blocks
contract SummitMock is SnapshotHubMock, AgentSecuredMock, InterfaceSummit {
    /// @notice Prevents this contract from being included in the coverage report
    function testSummitMock() external {}

    function acceptReceipt(
        uint32 rcptNotaryIndex,
        uint32 attNotaryIndex,
        uint256 sigIndex,
        uint32 attNonce,
        uint256 paddedTips,
        bytes memory rcptBodyPayload
    ) external returns (bool wasAccepted) {}

    function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes memory snapPayload) external {}

    function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes memory snapPayload)
        external
        returns (bytes memory attPayload)
    {}

    function distributeTips() external returns (bool queuePopped) {}

    function withdrawTips(uint32 origin, uint256 amount) external {}

    function actorTips(address actor, uint32 origin) external view returns (uint128 earned, uint128 claimed) {}

    function receiptQueueLength() external view returns (uint256) {}

    function getLatestState(uint32 origin) external view returns (bytes memory statePayload) {}
}
