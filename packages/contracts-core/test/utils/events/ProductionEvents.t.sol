// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DestinationEvents } from "../../../contracts/events/DestinationEvents.sol";
import { OriginEvents } from "../../../contracts/events/OriginEvents.sol";
import { SnapshotHubEvents } from "../../../contracts/events/SnapshotHubEvents.sol";
import { SummitEvents } from "../../../contracts/events/SummitEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    DestinationEvents,
    OriginEvents,
    SnapshotHubEvents,
    SummitEvents
{

}
