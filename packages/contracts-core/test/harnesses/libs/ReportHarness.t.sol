// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Report } from "../../../contracts/libs/Report.sol";
import { TypedMemView } from "../../../contracts/libs/TypedMemView.sol";

/**
 * @notice Exposes Report methods for testing against golang.
 */
contract ReportHarness {
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               GETTERS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToReport(uint40, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        bytes29 _view = Report.castToReport(_payload);
        return (_view.typeOf(), _view.clone());
    }

    function reportedAttestation(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).reportedAttestation();
        return (_view.typeOf(), _view.clone());
    }

    function reportData(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).reportData();
        return (_view.typeOf(), _view.clone());
    }

    function guardSignature(uint40 _type, bytes memory _payload)
        public
        view
        returns (uint40, bytes memory)
    {
        bytes29 _view = _payload.ref(_type).guardSignature();
        return (_view.typeOf(), _view.clone());
    }

    function reportedFraud(uint40 _type, bytes memory _payload) public pure returns (bool) {
        return _payload.ref(_type).reportedFraud();
    }

    function isReport(bytes memory _payload) public pure returns (bool) {
        return _payload.castToReport().isReport();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              FORMATTERS                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line ordering
    function formatReportData(Report.Flag _flag, bytes memory _attestation)
        public
        view
        returns (bytes memory)
    {
        return Report.formatReportData(_flag, _attestation);
    }

    function formatValidReportData(bytes memory _attestation) public view returns (bytes memory) {
        return Report.formatValidReportData(_attestation);
    }

    function formatFraudReportData(bytes memory _attestation) public view returns (bytes memory) {
        return Report.formatFraudReportData(_attestation);
    }

    function formatReport(
        Report.Flag _flag,
        bytes memory _attestation,
        bytes memory _guardSig
    ) public pure returns (bytes memory) {
        return Report.formatReport(_flag, _attestation, _guardSig);
    }

    function formatValidReport(bytes memory _attestation, bytes memory _guardSig)
        public
        pure
        returns (bytes memory)
    {
        return Report.formatValidReport(_attestation, _guardSig);
    }

    function formatFraudReport(bytes memory _attestation, bytes memory _guardSig)
        public
        pure
        returns (bytes memory)
    {
        return Report.formatFraudReport(_attestation, _guardSig);
    }
}
