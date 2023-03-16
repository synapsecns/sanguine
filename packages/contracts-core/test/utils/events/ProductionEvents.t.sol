// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AgentRegistryEvents } from "../../../contracts/events/AgentRegistryEvents.sol";
import { DestinationEvents } from "../../../contracts/events/DestinationEvents.sol";
import { DisputeHubEvents } from "../../../contracts/events/DisputeHubEvents.sol";
import { OriginEvents } from "../../../contracts/events/OriginEvents.sol";
import { SnapshotHubEvents } from "../../../contracts/events/SnapshotHubEvents.sol";
import { SummitEvents } from "../../../contracts/events/SummitEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    AgentRegistryEvents,
    DestinationEvents,
    DisputeHubEvents,
    OriginEvents,
    SnapshotHubEvents,
    SummitEvents
{

}
