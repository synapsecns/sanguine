// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Production events ============
import "../../contracts/events/AgentRegistryEvents.sol";
import "../../contracts/events/AttestationHubEvents.sol";
import "../../contracts/events/AttestationCollectorEvents.sol";
import "../../contracts/events/DestinationEvents.sol";
import "../../contracts/events/GuardRegistryEvents.sol";
import "../../contracts/events/NotaryRegistryEvents.sol";
import "../../contracts/events/OriginEvents.sol";
import "../../contracts/events/OriginHubEvents.sol";
// ============ Harness events ============
import "../harnesses/events/AgentRegistryHarnessEvents.sol";
import "../harnesses/events/AttestationHubHarnessEvents.sol";
import "../harnesses/events/BasicClientHarnessEvents.sol";
import "../harnesses/events/ClientHarnessEvents.sol";
import "../harnesses/events/DestinationHarnessEvents.sol";
import "../harnesses/events/NotaryRegistryHarnessEvents.sol";
import "../harnesses/events/ReportHubHarnessEvents.sol";
import "../harnesses/events/SystemContractHarnessEvents.sol";
// ============ Mocks events ============
import "../mocks/events/SystemContractMockEvents.sol";

// solhint-disable-next-line no-empty-blocks
abstract contract SynapseEvents is
    AgentRegistryEvents,
    AttestationCollectorEvents,
    AttestationHubEvents,
    DestinationEvents,
    GuardRegistryEvents,
    NotaryRegistryEvents,
    OriginEvents,
    OriginHubEvents,
    // Harnesses events
    AgentRegistryHarnessEvents,
    AttestationHubHarnessEvents,
    BasicClientHarnessEvents,
    ClientHarnessEvents,
    DestinationHarnessEvents,
    NotaryRegistryHarnessEvents,
    ReportHubHarnessEvents,
    SystemContractHarnessEvents,
    // Mocks events
    SystemContractMockEvents
{

}
