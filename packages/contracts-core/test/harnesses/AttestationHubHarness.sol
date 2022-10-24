// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { AttestationHub } from "../../contracts/hubs/AttestationHub.sol";

import { GlobalNotaryRegistryHarness } from "./GlobalNotaryRegistryHarness.sol";

contract AttestationHubHarness is AttestationHub, GlobalNotaryRegistryHarness {
    using TypedMemView for bytes29;

    event LogAttestation(address notary, bytes attestationView, bytes attestation);

    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool) {
        emit LogAttestation(_notary, _attestationView.clone(), _attestation);
        return true;
    }
}
