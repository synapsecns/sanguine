// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";
import { AttestationHub } from "../../../contracts/hubs/AttestationHub.sol";

import { AttestationHubHarnessEvents } from "../events/AttestationHubHarnessEvents.sol";
import { GlobalNotaryRegistryHarness } from "../registry/GlobalNotaryRegistryHarness.t.sol";

contract AttestationHubHarness is
    AttestationHubHarnessEvents,
    AttestationHub,
    GlobalNotaryRegistryHarness
{
    using TypedMemView for bytes29;

    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool) {
        emit LogAttestation(_notary, _attestationView.clone(), _attestation);
        return true;
    }
}
