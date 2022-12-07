// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";
import { AttestationHub } from "../../../contracts/hubs/AttestationHub.sol";

import { AttestationHubHarnessEvents } from "../events/AttestationHubHarnessEvents.sol";
import { AgentRegistryExtended } from "../system/AgentRegistryExtended.t.sol";

contract AttestationHubHarness is
    AttestationHubHarnessEvents,
    AttestationHub,
    AgentRegistryExtended
{
    using TypedMemView for bytes29;

    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool) {
        emit LogAttestation(_guards, _notaries, _attestationView.clone(), _attestation);
        return true;
    }
}
