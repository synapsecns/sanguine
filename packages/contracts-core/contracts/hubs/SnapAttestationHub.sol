// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DestinationAttestation, SnapAttestation } from "../libs/SnapAttestation.sol";

/**
 * @notice Hub to accept and save attestations.
 */
abstract contract SnapAttestationHub {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Tracks all accepted Notary attestations
    // (root => attestation)
    mapping(bytes32 => DestinationAttestation) private rootAttestations;

    /// @dev All snapshot roots from the accepted Notary attestations
    bytes32[] private roots;

    /// @dev gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the total amount of Notaries attestations that have been accepted.
    function attestationsAmount() external view returns (uint256) {
        return roots.length;
    }

    /// @notice Returns an attestation from the list of all accepted Notary attestations.
    /// @dev Index refers to attestation's snapshot root position in `roots` array.
    /// @param _index   Attestation index
    /// @return root    Snapshot root for the attestation
    /// @return destAtt Rest of attestation data that Destination keeps track of
    function getAttestation(uint256 _index)
        external
        view
        returns (bytes32 root, DestinationAttestation memory destAtt)
    {
        require(_index < roots.length, "Index out of range");
        root = roots[_index];
        destAtt = rootAttestations[root];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ACCEPT DATA                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Accepts a SnapAttestation signed by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _acceptAttestation(SnapAttestation _snapAtt, address _notary) internal {
        bytes32 root = _snapAtt.root();
        require(_rootAttestation(root).isEmpty(), "Root already exists");
        rootAttestations[root] = _snapAtt.toDestinationAttestation(_notary);
        roots.push(root);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         CHECK STATEMENT DATA                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the saved attestation for the "Snapshot Merkle Root".
    /// Will return an empty struct, if the root hasn't been submitted in a Notary attestation yet.
    function _rootAttestation(bytes32 _root) internal view returns (DestinationAttestation memory) {
        return rootAttestations[_root];
    }
}
