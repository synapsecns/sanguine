// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AuthManager } from "./auth/AuthManager.sol";
import { Attestation } from "./libs/Attestation.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AttestationCollector is AuthManager, OwnableUpgradeable {
    using Attestation for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event AttestationSubmitted(address indexed updater, bytes attestation, bytes signature);

    event UpdaterAdded(uint32 indexed domain, address updater);

    event UpdaterRemoved(uint32 indexed domain, address updater);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // [homeDomain => [updater => isUpdater]]
    mapping(uint32 => mapping(address => bool)) public isUpdater;

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

    function submitAttestation(
        address _updater,
        bytes memory _attestation,
        bytes memory _signature
    ) external {
        bytes29 _view = _checkUpdaterAuth(_updater, _attestation, _signature);
        _storeAttestation(_view);
        emit AttestationSubmitted(_updater, _attestation, _signature);
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
    }
}
