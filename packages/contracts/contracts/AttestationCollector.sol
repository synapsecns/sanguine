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

    // stores the latest nonce
    mapping(uint32 => uint32) private domainNonces;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[49] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
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
        _storeAttestation(_view);
        emit AttestationSubmitted(_updater, _attestation);
    }

    function latestNonce(uint32 domain) external view returns (uint32) {
        return domainNonces[domain];
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

    function _storeAttestation(bytes29 _view) internal {
        // TODO: implement storing logic for easy retrieval
        uint32 domain = Attestation.attestationDomain(_view);

        uint32 newNonce = Attestation.attestationNonce(_view);

        if (newNonce > domainNonces[domain]) {
            domainNonces[domain] = newNonce;
        }
    }
}
