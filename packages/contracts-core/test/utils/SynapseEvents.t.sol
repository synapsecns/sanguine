// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Production events ============
import { AgentRegistryEvents } from "../../contracts/events/AgentRegistryEvents.sol";
// ============ Harness events ============
import { AgentRegistryHarnessEvents } from "../harnesses/events/AgentRegistryHarnessEvents.sol";
import { BasicClientHarnessEvents } from "../harnesses/events/BasicClientHarnessEvents.sol";
import { ClientHarnessEvents } from "../harnesses/events/ClientHarnessEvents.sol";
import { SystemContractHarnessEvents } from "../harnesses/events/SystemContractHarnessEvents.sol";

// solhint-disable-next-line no-empty-blocks
abstract contract SynapseEvents is
    AgentRegistryEvents,
    // Harnesses events
    AgentRegistryHarnessEvents,
    BasicClientHarnessEvents,
    ClientHarnessEvents,
    SystemContractHarnessEvents
{

}
