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
     * `attestedRoots` stores a list of attested roots for every (origin, destination, nonce) tuple
     * `signatures` stores a signature for every submitted (origin, destination, nonce, root) attestation.
     * We only store the first submitted signature for such attestation.
     */
    // [attestion_key(origin,destination,nonce) => [roots]]
    mapping(uint96 => bytes32[]) internal attestedRoots;
    // [attestion_key(origin,destination,nonce) => [root => signature]]
    mapping(uint96 => mapping(bytes32 => bytes)) internal signatures;

    /// @dev We are also storing last submitted (nonce, root) attestation for every Notary.
    // [attested_domain(origin,destination) => [notary => latestNonce]]
    mapping(uint64 => mapping(address => uint32)) public latestNonces;
    // [attested_domain(origin,destination) => [notary => latestRoot]]
    mapping(uint64 => mapping(address => bytes32)) public latestRoots;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event AttestationSubmitted(address indexed notary, bytes attestation);

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
     * If fraud was committed, there might be more than one attestation for given (origin_domain, destination_domain, nonce).
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
        uint256 amount = notariesAmount(_destination);
        require(amount != 0, "!notaries");
        uint32 latestNonce = 0;
        bytes32 latestRoot = bytes32(0);
        uint64 attestedDomains = Attestation.attestedDomains(_origin, _destination);
        for (uint256 i = 0; i < amount; ) {
            address notary = getNotary(_destination, i);
            uint32 nonce = latestNonces[attestedDomains][notary];
            // Check latest Notary's nonce against current latest nonce
            if (nonce > latestNonce) {
                latestRoot = latestRoots[attestedDomains][notary];
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
     * @notice Get latest attestation for the domain signed by given Notary.
     */
    function getLatestAttestation(
        uint32 _origin,
        uint32 _destination,
        address _notary
    ) external view returns (bytes memory) {
        uint64 attestedDomains = Attestation.attestedDomains(_origin, _destination);
        uint32 nonce = latestNonces[attestedDomains][_notary];
        require(nonce != 0, "No attestations found");
        bytes32 root = latestRoots[attestedDomains][_notary];
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
        uint96 attestionKey = Attestation.attestionKey(_origin, _destination, _nonce);
        return attestedRoots[attestionKey].length;
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
        uint96 attestionKey = Attestation.attestionKey(_origin, _destination, _nonce);
        require(_index < attestedRoots[attestionKey].length, "!index");
        return attestedRoots[attestionKey][_index];
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
        uint32 origin = _attestationView.attestedOrigin();
        uint32 destination = _attestationView.attestedDestination();
        uint32 nonce = _attestationView.attestedNonce();
        bytes32 root = _attestationView.attestedRoot();

        uint96 attestionKey = _attestationView.attestionKey();
        uint64 attestedDomains = _attestationView.attestedDomains();

        require(nonce > latestNonces[attestedDomains][_notary], "Outdated attestation");
        // Don't store Attestation, if another Notary
        // have submitted the same (domain, nonce, root) before.
        require(!_signatureExists(origin, destination, nonce, root), "Duplicated attestation");
        latestNonces[attestedDomains][_notary] = nonce;
        latestRoots[attestedDomains][_notary] = root;
        signatures[attestionKey][root] = _attestationView.notarySignature().clone();
        attestedRoots[attestionKey].push(root);
        emit AttestationSubmitted(_notary, _attestation);
        return true;
    }

    function _formatAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bytes memory) {
        uint96 attestionKey = Attestation.attestionKey(_origin, _destination, _nonce);
        return
            Attestation.formatAttestation(
                Attestation.formatAttestationData(_origin, _destination, _nonce, _root),
                signatures[attestionKey][_root]
            );
    }

    function _signatureExists(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root
    ) internal view returns (bool) {
        uint96 attestionKey = Attestation.attestionKey(_origin, _destination, _nonce);
        return signatures[attestionKey][_root].length > 0;
    }
}
