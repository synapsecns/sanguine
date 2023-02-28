// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Production events ============
import "../../contracts/events/AgentRegistryEvents.sol";
import "../../contracts/events/GuardRegistryEvents.sol";
import "../../contracts/events/NotaryRegistryEvents.sol";
// ============ Harness events ============
import "../harnesses/events/AgentRegistryHarnessEvents.sol";
import "../harnesses/events/BasicClientHarnessEvents.sol";
import "../harnesses/events/ClientHarnessEvents.sol";
import "../harnesses/events/NotaryRegistryHarnessEvents.sol";
import "../harnesses/events/SystemContractHarnessEvents.sol";
// ============ Mocks events ============
import "../mocks/events/SystemContractMockEvents.sol";

// solhint-disable-next-line no-empty-blocks
abstract contract SynapseEvents is
    AgentRegistryEvents,
    GuardRegistryEvents,
    NotaryRegistryEvents,
    // Harnesses events
    AgentRegistryHarnessEvents,
    BasicClientHarnessEvents,
    ClientHarnessEvents,
    NotaryRegistryHarnessEvents,
    SystemContractHarnessEvents,
    // Mocks events
    SystemContractMockEvents
{

}
