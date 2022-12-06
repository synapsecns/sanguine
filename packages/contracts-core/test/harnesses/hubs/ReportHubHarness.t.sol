// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";
import { ReportHub } from "../../../contracts/hubs/ReportHub.sol";

import { AttestationHubHarness } from "./AttestationHubHarness.t.sol";
import { ReportHubHarnessEvents } from "../events/ReportHubHarnessEvents.sol";

contract ReportHubHarness is ReportHubHarnessEvents, ReportHub, AttestationHubHarness {
    using TypedMemView for bytes29;

    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool) {
        /* TODO(Chi): enable reports once co-signed Attestation is implemented
        emit LogReport(_guard, _notary, _attestationView.clone(), _reportView.clone(), _report);
        return true;
        */
    }
}
