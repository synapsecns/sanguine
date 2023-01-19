// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./Attestation.sol";

/// @dev Report is a memory view over a formatted report payload.
type Report is bytes29;
/// @dev ReportData is a memory view over a formatted report data.
type ReportData is bytes29;

/**
 * @notice Library for formatting the Guard Reports.
 * Reports are submitted to Origin contracts in order to slash a fraudulent Notary.
 * Reports are submitted to Destination contracts in order to blacklist
 * an allegedly fraudulent Notary.
 * Just like an Attestation, a Report could be checked on Origin contract
 * on the chain the reported Notary is attesting.
 * Report includes:
 * - Flag, indicating whether the reported attestation is fraudulent.
 * - Reported Attestation (Attestation data and Notary signature on that data).
 * - Guard signature on Report data.
 */
library ReportLib {
    using AttestationLib for bytes;
    using AttestationLib for Attestation;
    using AttestationLib for AttestationData;
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev More flag values could be added in the future,
     *      e.g. flag indicating "type" of fraud.
     *      Going forward, Flag.Valid is guaranteed to be
     *      the only Flag specifying a valid attestation.
     *
     *      Flag.Valid indicates a reported valid Attestation.
     *      Flag.Fraud indicates a reported fraud Attestation.
     */
    enum Flag {
        Valid,
        Fraud
    }

    // TODO: revisit once Report structure is finalized
    /**
     * @dev ReportData memory layout
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 045): attData        bytes   44 bytes
     *
     * guardSig is Guard's signature on ReportData
     *
     *      Report memory layout
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 110): attestation    bytes   109 bytes (44 + 65 bytes)
     * [110 .. 175): guardSig       bytes   65 bytes
     *
     *      Unpack attestation field (see Attestation.sol)
     * [000 .. 001): flag           uint8    1 bytes
     * [001 .. 045): attData        bytes   44 bytes
     * [045 .. 110): notarySig      bytes   65 bytes
     * [110 .. 175): guardSig       bytes   65 bytes
     *
     * notarySig is Notary's signature on AttestationData
     *
     * flag + attData = reportData (see above), so
     *
     *      Report memory layout (sliced alternatively)
     * [000 .. 045): reportData     bytes   45 bytes
     * [045 .. 110): notarySig      bytes   65 bytes
     * [110 .. 171): guardSig       bytes   61 bytes
     */

    uint256 internal constant OFFSET_FLAG = 0;
    uint256 internal constant OFFSET_ATTESTATION = 1;

    uint256 internal constant ATTESTATION_DATA_LENGTH = 44;
    uint256 internal constant REPORT_DATA_LENGTH = 1 + ATTESTATION_DATA_LENGTH;
    uint256 internal constant REPORT_LENGTH = REPORT_DATA_LENGTH + 2 * 65;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             REPORT DATA                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns formatted report data with provided fields
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attPayload   Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatReportData(Flag _flag, bytes memory _attPayload)
        internal
        view
        returns (bytes memory)
    {
        // Extract attestation data from payload
        AttestationData attData = _attPayload.castToAttestation().data();
        // Construct report data
        return abi.encodePacked(uint8(_flag), attData.unwrap().clone());
    }

    /**
     * @notice Returns formatted report data on valid attestation with provided fields
     * @param _validAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatValidReportData(bytes memory _validAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Valid, _validAttestation);
    }

    /**
     * @notice Returns formatted report data on fraud attestation with provided fields
     * @param _fraudAttestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatFraudReportData(bytes memory _fraudAttestation)
        internal
        view
        returns (bytes memory)
    {
        return formatReportData(Flag.Fraud, _fraudAttestation);
    }

    /**
     * @notice Returns a ReportData view over the given payload.
     * @dev Will revert if the payload is not a report data.
     */
    function castToReportData(bytes memory _payload) internal pure returns (ReportData) {
        return castToReportData(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a ReportData view.
     * @dev Will revert if the memory view is not over a report data.
     */
    function castToReportData(bytes29 _view) internal pure returns (ReportData) {
        require(isReportData(_view), "Not a report data");
        return ReportData.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted Attestation Data.
    function isReportData(bytes29 _view) internal pure returns (bool) {
        return _view.len() == REPORT_DATA_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(ReportData _data) internal pure returns (bytes29) {
        return ReportData.unwrap(_data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                REPORT                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns formatted report payload with provided fields.
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attPayload   Formatted attestation (see Attestation.sol)
     * @param _guardSig     Guard signature on reportData (see formatReportData below)
     * @return Formatted report
     **/
    function formatReport(
        Flag _flag,
        bytes memory _attPayload,
        bytes memory _guardSig
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(uint8(_flag), _attPayload, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a valid attestation with provided fields.
     * @param _validAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatValidReport(bytes memory _validAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Valid, _validAttestation, _guardSig);
    }

    /**
     * @notice Returns formatted report payload on a fraud attestation with provided fields.
     * @param _fraudAttestation Formatted attestation (see Attestation.sol)
     * @param _guardSig         Guard signature on reportData (see ReportData section above)
     * @return Formatted report
     **/
    function formatFraudReport(bytes memory _fraudAttestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(Flag.Fraud, _fraudAttestation, _guardSig);
    }

    /**
     * @notice Returns a Report view over the given payload.
     * @dev Will revert if the payload is not a report.
     */
    function castToReport(bytes memory _payload) internal pure returns (Report) {
        return castToReport(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Report view.
     * @dev Will revert if the memory view is not over a report.
     */
    function castToReport(bytes29 _view) internal pure returns (Report) {
        require(isReport(_view), "Not a report data");
        return Report.wrap(_view);
    }

    /**
     * @notice Checks that a payload is a formatted Report payload.
     */
    function isReport(bytes29 _view) internal pure returns (bool) {
        // TODO: revisit once Report structure is finalized
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Report _data) internal pure returns (bytes29) {
        return Report.unwrap(_data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            REPORT SLICING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: revisit once Report structure is finalized

    /**
     * @notice Returns Report's Data, that is going to be signed by the Guard.
     */
    function reportData(bytes29 _view) internal pure returns (ReportData) {}

    /**
     * @notice Returns Guard's signature on ReportData.
     */
    function guardSignature(bytes29 _view) internal pure returns (Signature) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          PRIVATE FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Returns int value of Report flag.
     *      Needed to prevent overflow when casting to Flag.
     */
    function _flagIntValue(bytes29 _view) private pure returns (uint8 flagIntValue) {
        flagIntValue = uint8(_view.indexUint({ _index: OFFSET_FLAG, _bytes: 1 }));
    }
}
