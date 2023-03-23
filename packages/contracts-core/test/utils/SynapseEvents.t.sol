// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ============ Harness events ============
import { BasicClientHarnessEvents } from "../harnesses/events/BasicClientHarnessEvents.sol";
import { ClientHarnessEvents } from "../harnesses/events/ClientHarnessEvents.sol";
import { SystemContractHarnessEvents } from "../harnesses/events/SystemContractHarnessEvents.sol";

// solhint-disable no-empty-blocks
abstract contract SynapseEvents is
    BasicClientHarnessEvents,
    ClientHarnessEvents,
    SystemContractHarnessEvents
{

}
