// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentManagerEvents} from "../../../contracts/events/AgentManagerEvents.sol";
import {DestinationEvents} from "../../../contracts/events/DestinationEvents.sol";
import {ExecutionHubEvents} from "../../../contracts/events/ExecutionHubEvents.sol";
import {GasOracleEvents} from "../../../contracts/events/GasOracleEvents.sol";
import {InboxEvents} from "../../../contracts/events/InboxEvents.sol";
import {OriginEvents} from "../../../contracts/events/OriginEvents.sol";
import {SnapshotHubEvents} from "../../../contracts/events/SnapshotHubEvents.sol";
import {StatementInboxEvents} from "../../../contracts/events/StatementInboxEvents.sol";
import {SummitEvents} from "../../../contracts/events/SummitEvents.sol";

// solhint-disable no-empty-blocks
abstract contract ProductionEvents is
    AgentManagerEvents,
    DestinationEvents,
    ExecutionHubEvents,
    GasOracleEvents,
    InboxEvents,
    OriginEvents,
    SnapshotHubEvents,
    StatementInboxEvents,
    SummitEvents
{}
