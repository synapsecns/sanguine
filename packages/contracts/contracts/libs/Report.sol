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

    modifier onlyType(bytes29 _view, uint40 _type) {
        _view.assertType(_type);
        _;
    }

    /**
     * @dev Report memory layout
     * [000 .. 002): attLength      uint16   2 bytes
     * [002 .. AAA): attestation    bytes   ?? bytes (40 + 64/65 bytes)
     * [AAA .. END): signature      bytes   ?? bytes (64/65 bytes)
     */

    uint256 internal constant OFFSET_ATTESTATION_LENGTH = 0;
    uint256 internal constant OFFSET_ATTESTATION = 2;

    function formatReport(bytes memory _attestation, bytes memory _signature)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(uint16(_attestation.length), _attestation, _signature);
    }

    function castToReport(bytes memory _payload) internal pure returns (bytes29) {
        return _payload.ref(SynapseTypes.REPORT);
    }

    function isReport(bytes29 _view) internal pure returns (bool) {
        uint256 attestationLength = reportAttestationLength(_view);
        // signature needs to exist
        return (_view.len() > OFFSET_ATTESTATION + attestationLength);
    }

    /// @dev No type checks in private functions,
    /// as the type is checked in the function that called this one.
    function reportAttestationLength(bytes29 _view) private pure returns (uint256) {
        return _view.indexUint(OFFSET_ATTESTATION_LENGTH, 2);
    }

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

    function reportSignature(bytes29 _view)
        internal
        pure
        onlyType(_view, SynapseTypes.REPORT)
        returns (bytes29)
    {
        uint256 offsetSignature = OFFSET_ATTESTATION + reportAttestationLength(_view);
        return _view.slice(offsetSignature, _view.len() - offsetSignature, SynapseTypes.SIGNATURE);
    }
}
