// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Production events ============
import "../../contracts/events/AttestationCollectorEvents.sol";
import "../../contracts/events/DestinationEvents.sol";
import "../../contracts/events/DestinationHubEvents.sol";
import "../../contracts/events/GuardRegistryEvents.sol";
import "../../contracts/events/NotaryManagerEvents.sol";
import "../../contracts/events/NotaryRegistryEvents.sol";
import "../../contracts/events/OriginEvents.sol";
import "../../contracts/events/OriginHubEvents.sol";
// ============ Harness events ============
import "../harnesses/events/AttestationHubHarnessEvents.sol";
import "../harnesses/events/BasicClientHarnessEvents.sol";
import "../harnesses/events/ClientHarnessEvents.sol";
import "../harnesses/events/DestinationHarnessEvents.sol";
import "../harnesses/events/ReportHubHarnessEvents.sol";
import "../harnesses/events/SystemContractHarnessEvents.sol";

// solhint-disable-next-line no-empty-blocks
abstract contract SynapseEvents is
    AttestationCollectorEvents,
    DestinationEvents,
    DestinationHubEvents,
    GuardRegistryEvents,
    NotaryManagerEvents,
    NotaryRegistryEvents,
    OriginEvents,
    OriginHubEvents,
    AttestationHubHarnessEvents,
    BasicClientHarnessEvents,
    ClientHarnessEvents,
    DestinationHarnessEvents,
    ReportHubHarnessEvents,
    SystemContractHarnessEvents
{

}
