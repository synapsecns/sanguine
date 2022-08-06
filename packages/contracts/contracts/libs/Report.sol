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
 *          Report includes an allegedly fraudulent Attestation (which includes the Notary signature),
 *          and Guard signature on such Attestation.
 */
library Report {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // More flag values could be added in the future
    enum FraudFlag {
        UNKNOWN,
        FRAUD
    }

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /**
     * @dev ReportData memory layout
     * [000 .. 001): fraudFlag      uint8    1 bytes
     * [001 .. 041): attData        bytes   40 bytes
     *
     * guardSig is Guard's signature on ReportData
     *
     *      Report memory layout
     * [000 .. 002): attLength      uint16   2 bytes
     * [002 .. 003): fraudFlag      uint8    1 bytes
     * [003 .. AAA]: attestation    bytes   ?? bytes (40 + 64/65 bytes)
     * [AAA .. END): guardSig       bytes   ?? bytes (64/65 bytes)
     *
     *      Closer look into Report (see Attestation.sol)
     * [002 .. 003): fraudFlag      uint8    1 bytes
     * [003 .. 043]: attData        bytes   40 bytes
     * [043 .. AAA): notarySig      bytes   ?? bytes (64/65 bytes)
     *
     * fraudFlag + attData = reportData, so
     *
     *      Report memory layout (sliced alternatively)
     * [000 .. 002): attLength      uint16   2 bytes
     * [002 .. 043): reportData     bytes   41 bytes
     * [043 .. AAA): notarySig      bytes   ?? bytes (64/65 bytes)
     * [AAA .. END): guardSig       bytes   ?? bytes (64/65 bytes)
     */

    uint256 internal constant OFFSET_ATTESTATION_LENGTH = 0;
    uint256 internal constant OFFSET_FRAUD_FLAG = 2;
    uint256 internal constant OFFSET_ATTESTATION = 3;

    uint256 internal constant ATTESTATION_DATA_LENGTH = 40;
    uint256 internal constant REPORT_DATA_LENGTH = 1 + ATTESTATION_DATA_LENGTH;

    /**
     * @notice Returns formatted (packed) report with provided fields
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @param _guardSig     Guard signature on reportData (see formatReportData below)
     * @return Formatted report
     **/
    function formatReport(
        FraudFlag _flag,
        bytes memory _attestation,
        bytes memory _guardSig
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(uint16(_attestation.length), uint8(_flag), _attestation, _guardSig);
    }

    /**
     * @notice Returns formatted (packed) FRAUD report with provided fields
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @param _guardSig     Guard signature on reportData (see formatReportData below)
     * @return Formatted report
     **/
    function formatFraudReport(bytes memory _attestation, bytes memory _guardSig)
        internal
        pure
        returns (bytes memory)
    {
        return formatReport(FraudFlag.FRAUD, _attestation, _guardSig);
    }

    /**
     * @notice Returns formatted (packed) report data with provided fields
     * @param _flag         Flag indicating whether attestation is fraudulent
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatReportData(FraudFlag _flag, bytes memory _attestation)
        internal
        view
        returns (bytes memory)
    {
        bytes memory attestationData = Attestation
            .attestationData(Attestation.attestationView(_attestation))
            .clone();
        return abi.encodePacked(uint8(_flag), attestationData);
    }

    /**
     * @notice Returns formatted (packed) FRAUD report data with provided fields
     * @param _attestation  Formatted attestation (see Attestation.sol)
     * @return Formatted report data
     **/
    function formatFraudReportData(bytes memory _attestation) internal view returns (bytes memory) {
        return formatReportData(FraudFlag.FRAUD, _attestation);
    }

    /**
     * @notice Returns typed memory view over given payload (type = REPORT)
     */
    function reportView(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.REPORT);
    }

    /**
     * @notice Checks that payload is a Report, by checking its length.
     */
    function isReport(bytes29 _view) internal pure returns (bool) {
        uint256 attestationLength = reportAttestationLength(_view);
        // Guard signature needs to exist
        return (_view.len() > OFFSET_ATTESTATION + attestationLength);
    }

    /// @dev No type checks in private functions,
    /// as the type is checked in the function that called this one.
    function reportAttestationLength(bytes29 _view) private pure returns (uint256) {
        return _view.indexUint(OFFSET_ATTESTATION_LENGTH, 2);
    }

    /**
     * @notice Checks Fraud Flag of the Report.
     */
    function reportIsFraud(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bool)
    {
        return _view.indexUint(OFFSET_FRAUD_FLAG, 1) == uint8(FraudFlag.FRAUD);
    }

    /**
     * @notice Returns Report's Data, that is going to be signed by the Guard.
     */
    function reportData(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        return _view.slice(OFFSET_FRAUD_FLAG, REPORT_DATA_LENGTH, SynapseTypes.REPORT_DATA);
    }

    /**
     * @notice Returns Report's Attestation (which is supposed to be signed by the Notary already).
     */
    function reportAttestation(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        return
            _view.slice(
                OFFSET_ATTESTATION,
                reportAttestationLength(_view),
                SynapseTypes.ATTESTATION
            );
    }

    /**
     * @notice Returns Guard's signature on ReportData.
     */
    function reportGuardSignature(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        uint256 offsetSignature = OFFSET_ATTESTATION + reportAttestationLength(_view);
        return _view.slice(offsetSignature, _view.len() - offsetSignature, SynapseTypes.SIGNATURE);
    }
}
