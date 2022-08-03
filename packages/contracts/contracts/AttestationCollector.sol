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

    event AttestationSubmitted(address indexed notary, bytes attestation);

    event NotaryAdded(uint32 indexed domain, address notary);

    event NotaryRemoved(uint32 indexed domain, address notary);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // [homeDomain => [notary => isNotary]]
    mapping(uint32 => mapping(address => bool)) public isNotary;

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

    // TODO: add/remove notarys upon bonding/unbonding

    function addNotary(uint32 _domain, address _notary) external onlyOwner {
        _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) external onlyOwner {
        _removeNotary(_domain, _notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function submitAttestation(address _notary, bytes memory _attestation) external {
        bytes29 _view = _checkNotaryAuth(_notary, _attestation);
        _storeAttestation(_view);
        emit AttestationSubmitted(_notary, _attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isNotary(uint32 _homeDomain, address _notary) internal view override returns (bool) {
        return isNotary[_homeDomain][_notary];
    }

    function _isGuard(address _guard) internal view override returns (bool) {}

    function _addNotary(uint32 _domain, address _notary) internal {
        if (!isNotary[_domain][_notary]) {
            isNotary[_domain][_notary] = true;
            emit NotaryAdded(_domain, _notary);
        }
    }

    function _removeNotary(uint32 _domain, address _notary) internal {
        if (isNotary[_domain][_notary]) {
            isNotary[_domain][_notary] = false;
            emit NotaryRemoved(_domain, _notary);
        }
    }

    function _storeAttestation(bytes29 _view) internal {
        // TODO: implement storing logic for easy retrieval
    }
}
