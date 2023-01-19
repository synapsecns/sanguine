// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../../../contracts/libs/Report.sol";

/**
 * @notice Exposes ReportLib methods for testing against golang.
 */
contract ReportHarness {
    using ReportLib for bytes;
    using ReportLib for bytes29;
    using ReportLib for Report;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToReport(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        return ReportLib.castToReport(_payload).unwrap().clone();
    }

    function isReport(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isReport();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line ordering
    function formatReportData(ReportLib.Flag _flag, bytes memory _attestation)
        public
        view
        returns (bytes memory)
    {
        return ReportLib.formatReportData(_flag, _attestation);
    }

    function formatValidReportData(bytes memory _attestation) public view returns (bytes memory) {
        return ReportLib.formatValidReportData(_attestation);
    }

    function formatFraudReportData(bytes memory _attestation) public view returns (bytes memory) {
        return ReportLib.formatFraudReportData(_attestation);
    }

    function formatReport(
        ReportLib.Flag _flag,
        bytes memory _attestation,
        bytes memory _guardSig
    ) public pure returns (bytes memory) {
        return ReportLib.formatReport(_flag, _attestation, _guardSig);
    }

    function formatValidReport(bytes memory _attestation, bytes memory _guardSig)
        public
        pure
        returns (bytes memory)
    {
        return ReportLib.formatValidReport(_attestation, _guardSig);
    }

    function formatFraudReport(bytes memory _attestation, bytes memory _guardSig)
        public
        pure
        returns (bytes memory)
    {
        return ReportLib.formatFraudReport(_attestation, _guardSig);
    }
}
