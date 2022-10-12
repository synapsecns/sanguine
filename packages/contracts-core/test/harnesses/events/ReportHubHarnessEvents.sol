// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract ReportHubHarnessEvents {
    event LogReport(
        address guard,
        address notary,
        bytes attestation,
        bytes reportView,
        bytes report
    );
}
