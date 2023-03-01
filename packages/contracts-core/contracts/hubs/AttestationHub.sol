// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Attestation.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { IAttestationHub } from "../interfaces/IAttestationHub.sol";

/**
 * @notice Hub to accept and save attestations.
 */
abstract contract AttestationHub is IAttestationHub {
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

    /// @inheritdoc IAttestationHub
    function attestationsAmount() external view returns (uint256) {
        return roots.length;
    }

    /// @inheritdoc IAttestationHub
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

    /// @dev Accepts an Attestation signed by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _acceptAttestation(Attestation _att, address _notary) internal {
        bytes32 root = _att.root();
        require(_rootAttestation(root).isEmpty(), "Root already exists");
        rootAttestations[root] = _att.toDestinationAttestation(_notary);
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
