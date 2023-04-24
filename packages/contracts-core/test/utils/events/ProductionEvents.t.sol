// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManagerEvents} from "../../../contracts/events/AgentManagerEvents.sol";
import {AgentSecuredEvents} from "../../../contracts/events/AgentSecuredEvents.sol";
import {BondingManagerEvents} from "../../../contracts/events/BondingManagerEvents.sol";
import {DestinationEvents} from "../../../contracts/events/DestinationEvents.sol";
import {DisputeHubEvents} from "../../../contracts/events/DisputeHubEvents.sol";
import {ExecutionHubEvents} from "../../../contracts/events/ExecutionHubEvents.sol";
import {OriginEvents} from "../../../contracts/events/OriginEvents.sol";
import {SnapshotHubEvents} from "../../../contracts/events/SnapshotHubEvents.sol";
import {SummitEvents} from "../../../contracts/events/SummitEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    AgentManagerEvents,
    AgentSecuredEvents,
    BondingManagerEvents,
    DestinationEvents,
    DisputeHubEvents,
    ExecutionHubEvents,
    OriginEvents,
    SnapshotHubEvents,
    SummitEvents
{}
