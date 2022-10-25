// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { ReportHub } from "../../contracts/hubs/ReportHub.sol";

import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { GlobalNotaryRegistryHarness } from "./GlobalNotaryRegistryHarness.sol";

contract ReportHubHarness is ReportHub, GuardRegistryHarness, GlobalNotaryRegistryHarness {
    using TypedMemView for bytes29;

    event LogReport(
        address guard,
        address notary,
        bytes attestation,
        bytes reportView,
        bytes report
    );

    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool) {
        emit LogReport(_guard, _notary, _attestationView.clone(), _reportView.clone(), _report);
        return true;
    }
}
