// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AuthManager } from "./auth/AuthManager.sol";
import { Attestation } from "./libs/Attestation.sol";
import { TypedMemView } from "./libs/TypedMemView.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AttestationCollector is AuthManager, OwnableUpgradeable {
    using Attestation for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event AttestationSubmitted(address indexed updater, bytes attestation);

    event UpdaterAdded(uint32 indexed domain, address updater);

    event UpdaterRemoved(uint32 indexed domain, address updater);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // [homeDomain => [updater => isUpdater]]
    mapping(uint32 => mapping(address => bool)) public isUpdater;

    /**
     * @dev All submitted Notary Attestations are stored.
     * As different Notaries might sign attestations with the same nonce,
     * but different root (meaning one of the attestations is fraudulent),
     * we need a system so store all such attestations.
     *
     * `attestationRoots` stores a list of attested roots for every (domain, nonce) pair
     * `signatures` stores a signature for every submitted (domain, nonce, root) attestation.
     * We only store the first submitted signature for such attestation.
     */
    // [homeDomain => [nonce => [roots]]]
    mapping(uint32 => mapping(uint32 => bytes32[])) internal attestationRoots;
    // [homeDomain => [nonce => [root => signature]]]
    mapping(uint32 => mapping(uint32 => mapping(bytes32 => bytes))) internal signatures;

    /// @dev We are also storing last submitted (nonce, root) attestation for every Notary.
    // [homeDomain => [notary => latestNonce]]
    mapping(uint32 => mapping(address => uint32)) public latestNonce;
    // [homeDomain => [notary => latestRoot]]
    mapping(uint32 => mapping(address => bytes32)) public latestRoot;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[45] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Get i-th attestation for given (domain, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be more than one attestation for given (domain, nonce).
     */
    function getAttestation(
        uint32 _domain,
        uint32 _nonce,
        uint256 _index
    ) external view returns (bytes memory) {
        bytes32 root = getRoot(_domain, _nonce, _index);
        // signature always exists for a stored root
        return _formatAttestation(_domain, _nonce, root);
    }

    /**
     * @notice Get attestation for (domain, nonce, root), if exists.
     */
    function getAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) external view returns (bytes memory) {
        require(_signatureExists(_domain, _nonce, _root), "!signature");
        return _formatAttestation(_domain, _nonce, _root);
    }

    /**
     * @notice Get latest attestation for the domain (WIP).
     */
    function getLatestAttestation(uint32 _domain) external view returns (bytes memory) {
        // TODO: enumerate Notaries to implement this
    }

    /**
     * @notice Get latest attestation for the domain signed by given Notary.
     */
    function getLatestAttestation(uint32 _domain, address _updater)
        external
        view
        returns (bytes memory)
    {
        uint32 nonce = latestNonce[_domain][_updater];
        require(nonce > 0, "No attestations found");
        bytes32 root = latestRoot[_domain][_updater];
        return _formatAttestation(_domain, nonce, root);
    }

    /**
     * @notice Get i-th root for given (domain, nonce), if exists.
     * Assuming no fraud is committed, index = 0 should be used.
     * If fraud was committed, there might be more than one root for given (domain, nonce).
     */
    function getRoot(
        uint32 _domain,
        uint32 _nonce,
        uint256 _index
    ) public view returns (bytes32) {
        require(_index < attestationRoots[_domain][_nonce].length, "!index");
        return attestationRoots[_domain][_nonce][_index];
    }

    /**
     * @notice Get amount of attested roots for given (domain, nonce).
     * Assuming no fraud is committed, amount <= 1.
     * If amount > 1, fraud was committed.
     */
    function rootsAmount(uint32 _domain, uint32 _nonce) external view returns (uint256) {
        return attestationRoots[_domain][_nonce].length;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: add/remove updaters upon bonding/unbonding

    function addUpdater(uint32 _domain, address _updater) external onlyOwner {
        _addUpdater(_domain, _updater);
    }

    function removeUpdater(uint32 _domain, address _updater) external onlyOwner {
        _removeUpdater(_domain, _updater);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function submitAttestation(address _updater, bytes memory _attestation) external {
        bytes29 _view = _checkUpdaterAuth(_updater, _attestation);
        _storeAttestation(_updater, _view);
        emit AttestationSubmitted(_updater, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isUpdater(uint32 _homeDomain, address _updater)
        internal
        view
        override
        returns (bool)
    {
        return isUpdater[_homeDomain][_updater];
    }

    function _isWatchtower(address _watchtower) internal view override returns (bool) {}

    function _addUpdater(uint32 _domain, address _updater) internal {
        if (!isUpdater[_domain][_updater]) {
            isUpdater[_domain][_updater] = true;
            emit UpdaterAdded(_domain, _updater);
        }
    }

    function _removeUpdater(uint32 _domain, address _updater) internal {
        if (isUpdater[_domain][_updater]) {
            isUpdater[_domain][_updater] = false;
            emit UpdaterRemoved(_domain, _updater);
        }
    }

    function _formatAttestation(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bytes memory) {
        return
            Attestation.formatAttestation(
                Attestation.formatAttestationData(_domain, _nonce, _root),
                signatures[_domain][_nonce][_root]
            );
    }

    function _signatureExists(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bool) {
        return signatures[_domain][_nonce][_root].length > 0;
    }

    function _storeAttestation(address _updater, bytes29 _view) internal {
        uint32 domain = _view.attestationDomain();
        uint32 nonce = _view.attestationNonce();
        bytes32 root = _view.attestationRoot();
        require(nonce > latestNonce[domain][_updater], "Outdated attestation");
        latestNonce[domain][_updater] = nonce;
        latestRoot[domain][_updater] = root;
        // Store root & signature only once
        if (!_signatureExists(domain, nonce, root)) {
            signatures[domain][nonce][root] = _view.attestationSignature().clone();
            attestationRoots[domain][nonce].push(root);
        }
    }
}
