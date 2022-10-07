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
import "../events/AttestationHubHarnessEvents.sol";
import "../events/ClientHarnessEvents.sol";
import "../events/DestinationHarnessEvents.sol";
import "../events/ReportHubHarnessEvents.sol";
import "../events/SystemContractHarnessEvents.sol";

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
    ClientHarnessEvents,
    DestinationHarnessEvents,
    ReportHubHarnessEvents,
    SystemContractHarnessEvents
{

}
