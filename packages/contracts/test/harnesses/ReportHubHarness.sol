// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";
import { ReportHub } from "../../contracts/hubs/ReportHub.sol";

import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { GlobalNotaryRegistryHarness } from "./GlobalNotaryRegistryHarness.sol";

contract ReportHubHarness is ReportHub, GuardRegistryHarness, GlobalNotaryRegistryHarness {
    using TypedMemView for bytes29;

    event LogFraudFlag(bool flag);
    event LogReport(address guard, address notary, bytes attestation, bytes report);

    function _checkFraudFlag(bool _flag) internal override {
        emit LogFraudFlag(_flag);
    }

    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes memory _report
    ) internal override returns (bool) {
        emit LogReport(_guard, _notary, _attestationView.clone(), _report);
        return true;
    }
}
