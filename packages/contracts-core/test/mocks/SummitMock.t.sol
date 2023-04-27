// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentStatus, InterfaceSummit} from "../../contracts/interfaces/InterfaceSummit.sol";
import {ExecutionHubMock} from "./hubs/ExecutionHubMock.t.sol";
import {SnapshotHubMock} from "./hubs/SnapshotHubMock.t.sol";
import {AgentSecuredMock} from "./base/AgentSecuredMock.t.sol";

// solhint-disable no-empty-blocks
contract SummitMock is ExecutionHubMock, SnapshotHubMock, AgentSecuredMock, InterfaceSummit {
    /// @notice Prevents this contract from being included in the coverage report
    function testSummitMock() external {}

    function acceptReceipt(
        AgentStatus memory rcptNotaryStatus,
        AgentStatus memory attNotaryStatus,
        uint256 sigIndex,
        bytes memory rcptPayload,
        uint32 attNonce
    ) external returns (bool wasAccepted) {}

    function acceptSnapshot(AgentStatus memory status, uint256 sigIndex, bytes memory snapPayload)
        external
        returns (bytes memory attPayload)
    {}

    function distributeTips() external returns (bool queuePopped) {}

    function withdrawTips(uint32 origin, uint256 amount) external {}

    function actorTips(address actor, uint32 origin) external view returns (uint128 earned, uint128 claimed) {}

    function receiptQueueLength() external view returns (uint256) {}

    function getLatestState(uint32 origin) external view returns (bytes memory statePayload) {}
}
