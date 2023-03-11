// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    AttestationFlag,
    AttestationReport,
    AttestationReportLib,
    TypedMemView
} from "../../../contracts/libs/AttestationReport.sol";

/// @notice Exposes Report methods for testing against golang.
contract ReportHarness {
    using AttestationReportLib for bytes;
    using AttestationReportLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      ATTESTATION REPORT GETTERS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToAttestationReport(bytes memory _payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        AttestationReport _attestation = AttestationReportLib.castToAttestationReport(_payload);
        return _attestation.unwrap().clone();
    }

    function arHash(bytes memory _payload) public pure returns (bytes32) {
        return _payload.castToAttestationReport().hash();
    }

    function arFlag(bytes memory _payload) public pure returns (AttestationFlag) {
        return _payload.castToAttestationReport().flag();
    }

    function arAttestation(bytes memory _payload) public view returns (bytes memory) {
        return _payload.castToAttestationReport().attestation().unwrap().clone();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    ATTESTATION REPORT FORMATTERS                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function formatAttestationReport(AttestationFlag _flag, bytes memory _attPayload)
        public
        pure
        returns (bytes memory)
    {
        return AttestationReportLib.formatAttestationReport(_flag, _attPayload);
    }

    function isAttestationReport(bytes memory _payload) public pure returns (bool) {
        return _payload.ref(0).isAttestationReport();
    }
}
