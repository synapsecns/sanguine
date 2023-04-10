// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import {
    AttestationFlag,
    AttestationReport,
    AttestationReportLib,
    TypedMemView
} from "../../../contracts/libs/AttestationReport.sol";

// solhint-disable ordering
/// @notice Exposes Report methods for testing against golang.
contract AttestationReportHarness {
    using AttestationReportLib for bytes;
    using AttestationReportLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToAttestationReport(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        AttestationReport attReport = AttestationReportLib.castToAttestationReport(payload);
        return attReport.unwrap().clone();
    }

    function hash(bytes memory payload) public pure returns (bytes32) {
        return payload.castToAttestationReport().hash();
    }

    function flag(bytes memory payload) public pure returns (AttestationFlag) {
        return payload.castToAttestationReport().flag();
    }

    function attestation(bytes memory payload) public view returns (bytes memory) {
        return payload.castToAttestationReport().attestation().unwrap().clone();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatAttestationReport(AttestationFlag flag_, bytes memory attPayload)
        public
        pure
        returns (bytes memory)
    {
        return AttestationReportLib.formatAttestationReport(flag_, attPayload);
    }

    function isAttestationReport(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isAttestationReport();
    }
}
