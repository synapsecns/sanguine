// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DestinationAttestation } from "../libs/SnapAttestation.sol";

interface ISnapAttestationHub {
    /**
     * @notice Returns the total amount of Notaries attestations that have been accepted.
     */
    function attestationsAmount() external view returns (uint256);

    /**
     * @notice Returns an attestation from the list of all accepted Notary attestations.
     * @dev Index refers to attestation's snapshot root position in `roots` array.
     * @param _index   Attestation index
     * @return root    Snapshot root for the attestation
     * @return destAtt Rest of attestation data that Destination keeps track of
     */
    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, DestinationAttestation memory destAtt);
}
