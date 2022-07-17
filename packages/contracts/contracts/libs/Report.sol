// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";
import { Attestation } from "./Attestation.sol";
import { SynapseTypes } from "./SynapseTypes.sol";

/**
 * @notice  Reports are submitted on Home contracts in order to slash a fraudulent Notary.
 *          Reports are submitted on ReplicaManager contracts in order to blacklist
 *          an allegedly fraudulent Notary.
 *          Just like an Attestation, a Report could be checked on Home contract
 *          back on the chain the Notary in question is attesting.
 */
library Report {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /**
     * @dev ReportData memory layout (to be signed by the Guard)
     * [000 .. 020): notary         address 20 bytes
     * [020 .. END): attestation    bytes   ?? bytes (40 + 64/65 bytes)
     *
     *      Report memory layout
     * [000 .. 002): dataLength     uint16   2 bytes
     * [002 .. AAA): data           bytes   ?? bytes (24 + 40 + 64/65 bytes)
     * [AAA .. END): signature      bytes   ?? bytes (64/65 bytes)
     */

    uint256 internal constant OFFSET_NOTARY = 0;
    uint256 internal constant OFFSET_ATTESTATION = 20;

    uint256 internal constant OFFSET_DATA_LENGTH = 0;
    uint256 internal constant OFFSET_DATA = 2;

    function formatReport(bytes memory _data, bytes memory _signature)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(uint16(_data.length), _data, _signature);
    }

    function formatReportData(address _notary, bytes memory _attestation)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(_notary, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                REPORT                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function castToReport(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.REPORT);
    }

    function isReport(bytes29 _view) internal pure returns (bool) {
        uint256 dataLength = reportDataLength(_view);
        // signature needs to exist
        if (_view.len() <= OFFSET_DATA + dataLength) return false;
        return isReportData(reportData(_view));
    }

    /// @dev No type checks in private functions,
    /// as the type is checked in the function that called this one.
    function reportDataLength(bytes29 _view) private pure returns (uint256) {
        return _view.indexUint(OFFSET_DATA_LENGTH, 2);
    }

    function reportData(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        return _view.slice(OFFSET_DATA, reportDataLength(_view), SynapseTypes.REPORT_DATA);
    }

    function reportSignature(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        uint256 offsetSignature = OFFSET_DATA + reportDataLength(_view);
        return _view.slice(offsetSignature, _view.len() - offsetSignature, SynapseTypes.SIGNATURE);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             REPORT DATA                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function isReportData(bytes29 _view) internal pure returns (bool) {
        // Attestation needs to exist
        if (_view.len() <= OFFSET_ATTESTATION) return false;
        return Attestation.isAttestation(reportAttestation(_view));
    }

    // TODO: rename to Notary
    function reportUpdater(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT_DATA)
        returns (address)
    {
        return _view.indexAddress(OFFSET_NOTARY);
    }

    function reportAttestation(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT_DATA)
        returns (bytes29)
    {
        return
            _view.slice(
                OFFSET_ATTESTATION,
                _view.len() - OFFSET_ATTESTATION,
                SynapseTypes.ATTESTATION
            );
    }
}
