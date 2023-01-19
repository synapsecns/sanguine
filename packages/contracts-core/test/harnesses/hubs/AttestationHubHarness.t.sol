// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/Attestation.sol";
import { AttestationHub } from "../../../contracts/hubs/AttestationHub.sol";

import { AttestationHubHarnessEvents } from "../events/AttestationHubHarnessEvents.sol";
import { AgentRegistryExtended } from "../system/AgentRegistryExtended.t.sol";

contract AttestationHubHarness is
    AttestationHubHarnessEvents,
    AttestationHub,
    AgentRegistryExtended
{
    using AttestationLib for Attestation;
    using TypedMemView for bytes29;

    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        Attestation _att,
        bytes memory _attPayload
    ) internal override returns (bool) {
        emit LogAttestation(_guards, _notaries, _att.unwrap().clone(), _attPayload);
        return true;
    }
}
