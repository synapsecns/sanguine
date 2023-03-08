// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Production events ============
import "../../contracts/events/AgentRegistryEvents.sol";
// ============ Harness events ============
import "../harnesses/events/AgentRegistryHarnessEvents.sol";
import "../harnesses/events/BasicClientHarnessEvents.sol";
import "../harnesses/events/ClientHarnessEvents.sol";
import "../harnesses/events/SystemContractHarnessEvents.sol";

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
