// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Attestation } from "./libs/Attestation.sol";
import { AttestationHub } from "./hubs/AttestationHub.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";
import { GlobalNotaryRegistry } from "./registry/GlobalNotaryRegistry.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AttestationCollector is AttestationHub, GlobalNotaryRegistry, OwnableUpgradeable {
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev All submitted Notary Attestations are stored.
     * As different Notaries might sign attestations with the same nonce,
     * but different root (meaning one of the attestations is fraudulent),
     * we need a system so store all such attestations.
     *
     * `attestedRoots`: attested roots for every (origin, destination, nonce) tuple
     * `signatures`: signatures for every submitted (origin, destination, nonce, root) attestation.
     * We only store the first submitted signature for such attestation.
     *
     * attestationKey = (origin, destination, nonce)
     */
    // [attestationKey => [roots]]
    mapping(uint96 => bytes32[]) internal attestedRoots;
    // [attestationKey => [root => signature]]
    mapping(uint96 => mapping(bytes32 => bytes)) internal signatures;

    /// @dev We are also storing last submitted (nonce, root) attestation for every Notary.
    /// attestationDomains = (origin, destination)
    // [attestationDomains => [notary => latestNonce]]
    mapping(uint64 => mapping(address => uint32)) internal latestNonces;
    // [attestationDomains => [notary => latestRoot]]
    mapping(uint64 => mapping(address => bytes32)) internal latestRoots;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: add/remove notaries upon bonding/unbonding

    function addNotary(uint32 _domain, address _notary) external onlyOwner returns (bool) {
        return _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) external onlyOwner returns (bool) {
        return _removeNotary(_domain, _notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Get i-th attestation for given (origin, destination, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be
     * more than one attestation for given (origin, destination, nonce).
     */
    function getAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        uint256 _index
    ) external view returns (bytes memory) {
        bytes32 root = getRoot(_origin, _destination, _nonce, _index);
        // signature always exists for a stored root
        return _formatAttestation(_origin, _destination, _nonce, root);
    }

    /**
     * @notice Get attestation for (origin, destination, nonce), if exists.
     */
    function getAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) external view returns (bytes memory) {
        require(_signatureExists(_origin, _destination, _nonce, _root), "!signature");
        return _formatAttestation(_origin, _destination, _nonce, _root);
    }

    /**
     * @notice Get latest attestation for the (origin, destination).
     */
    function getLatestAttestation(uint32 _origin, uint32 _destination)
        external
        view
        returns (bytes memory)
    {
        uint256 amount = notariesAmount(_origin);
        require(amount != 0, "!notaries");
        uint32 latestNonce = 0;
        bytes32 latestRoot = bytes32(0);
        uint64 attestationDomains = Attestation.attestationDomains(_origin, _destination);
        for (uint256 i = 0; i < amount; ) {
            address notary = getNotary(_destination, i);
            uint32 nonce = latestNonces[attestationDomains][notary];
            // Check latest Notary's nonce against current latest nonce
            if (nonce > latestNonce) {
                latestRoot = latestRoots[attestationDomains][notary];
                latestNonce = nonce;
            }
            unchecked {
                ++i;
            }
        }
        // Check if we found anything
        require(latestNonce != 0, "No attestations found");
        return _formatAttestation(_origin, _destination, latestNonce, latestRoot);
    }

    /**
     * @notice Get latest nonce for the (origin, destination, notary).
     */
    function getLatestNonce(
        uint32 _origin,
        uint32 _destination,
        address _notary
    ) external view returns (uint32) {
        uint64 attestationDomains = Attestation.attestationDomains(_origin, _destination);
        uint32 latestNonce = latestNonces[attestationDomains][_notary];
        // Check if we found anything
        require(latestNonce != 0, "No nonce found");
        return latestNonce;
    }

    /**
     * @notice Get latest root for the (origin, destination, notary).
     */
    function getLatestRoot(
        uint32 _origin,
        uint32 _destination,
        address _notary
    ) external view returns (bytes32) {
        uint64 attestationDomains = Attestation.attestationDomains(_origin, _destination);
        bytes32 latestRoot = latestRoots[attestationDomains][_notary];
        // Check if we found anything
        require(latestRoot != 0, "No root found");
        return latestRoot;
    }

    /**
     * @notice Get latest attestation for the domain signed by given Notary.
     */
    function getLatestAttestation(
        uint32 _origin,
        uint32 _destination,
        address _notary
    ) external view returns (bytes memory) {
        uint64 attestationDomains = Attestation.attestationDomains(_origin, _destination);
        uint32 nonce = latestNonces[attestationDomains][_notary];
        require(nonce != 0, "No attestations found");
        bytes32 root = latestRoots[attestationDomains][_notary];
        return _formatAttestation(_origin, _destination, nonce, root);
    }

    /**
     * @notice Get amount of attested roots for given (domain, nonce).
     * Assuming no fraud is committed, amount <= 1.
     * If amount > 1, fraud was committed.
     */
    function rootsAmount(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) external view returns (uint256) {
        uint96 attestationKey = Attestation.attestationKey(_origin, _destination, _nonce);
        return attestedRoots[attestationKey].length;
    }

    /**
     * @notice Get i-th root for given (domain, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be more than one root for given (domain, nonce).
     */
    function getRoot(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        uint256 _index
    ) public view returns (bytes32) {
        uint96 attestationKey = Attestation.attestationKey(_origin, _destination, _nonce);
        require(_index < attestedRoots[attestationKey].length, "!index");
        return attestedRoots[attestationKey][_index];
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Both Notary and Guard signatures
     * have been checked at this point (see ReportHub.sol).
     *
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return TRUE if Attestation was stored.
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool) {
        // Get attestation fields
        uint32 origin = _attestationView.attestedOrigin();
        uint32 destination = _attestationView.attestedDestination();
        uint32 nonce = _attestationView.attestedNonce();
        bytes32 root = _attestationView.attestedRoot();
        // Get attestation IDs
        uint96 attestedKey = _attestationView.attestedKey();
        uint64 attestedDomains = _attestationView.attestedDomains();
        // Check if the same Notary have already submitted a more recent attestation
        require(nonce > latestNonces[attestedDomains][_notary], "Outdated attestation");
        // Don't store Attestation, if another Notary
        // have submitted the same (origin, destination, nonce, root) before.
        require(!_signatureExists(origin, destination, nonce, root), "Duplicated attestation");
        // Update Notary's "latest attestation" for (origin, destination)
        latestNonces[attestedDomains][_notary] = nonce;
        latestRoots[attestedDomains][_notary] = root;
        // Save signature and root
        signatures[attestedKey][root] = _attestationView.notarySignature().clone();
        attestedRoots[attestedKey].push(root);
        emit AttestationAccepted(_notary, _attestation);
        return true;
    }

    function _formatAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bytes memory) {
        uint96 attestationKey = Attestation.attestationKey(_origin, _destination, _nonce);
        return
            Attestation.formatAttestation(
                Attestation.formatAttestationData(_origin, _destination, _nonce, _root),
                signatures[attestationKey][_root]
            );
    }

    function _signatureExists(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bool) {
        uint96 attestationKey = Attestation.attestationKey(_origin, _destination, _nonce);
        return signatures[attestationKey][_root].length > 0;
    }
}
