// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { BondingManagerEvents } from "../../../contracts/events/BondingManagerEvents.sol";
import { DestinationEvents } from "../../../contracts/events/DestinationEvents.sol";
import { DisputeHubEvents } from "../../../contracts/events/DisputeHubEvents.sol";
import { ExecutionHubEvents } from "../../../contracts/events/ExecutionHubEvents.sol";
import { OriginEvents } from "../../../contracts/events/OriginEvents.sol";
import { SnapshotHubEvents } from "../../../contracts/events/SnapshotHubEvents.sol";
import { SummitEvents } from "../../../contracts/events/SummitEvents.sol";
import { SystemRegistryEvents } from "../../../contracts/events/SystemRegistryEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    BondingManagerEvents,
    DestinationEvents,
    DisputeHubEvents,
    ExecutionHubEvents,
    OriginEvents,
    SnapshotHubEvents,
    SummitEvents,
    SystemRegistryEvents
{

}
