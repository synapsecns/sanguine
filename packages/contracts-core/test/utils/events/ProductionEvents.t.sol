// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManagerEvents} from "../../../contracts/events/AgentManagerEvents.sol";
import {BondingManagerEvents} from "../../../contracts/events/BondingManagerEvents.sol";
import {DestinationEvents} from "../../../contracts/events/DestinationEvents.sol";
import {ExecutionHubEvents} from "../../../contracts/events/ExecutionHubEvents.sol";
import {OriginEvents} from "../../../contracts/events/OriginEvents.sol";
import {SnapshotHubEvents} from "../../../contracts/events/SnapshotHubEvents.sol";
import {SummitEvents} from "../../../contracts/events/SummitEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    AgentManagerEvents,
    BondingManagerEvents,
    DestinationEvents,
    ExecutionHubEvents,
    OriginEvents,
    SnapshotHubEvents,
    SummitEvents
{}
