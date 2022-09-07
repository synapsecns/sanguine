// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AbstractNotaryRegistry } from "./AbstractNotaryRegistry.sol";

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice A Registry to keep track of all Notaries on a single domain.
 */
contract DomainNotaryRegistry is AbstractNotaryRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint32 private immutable trackedDomain;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    EnumerableSet.AddressSet internal notaries;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line var-name-mixedcase
    uint256[49] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event DomainNotaryAdded(address notary);

    event DomainNotaryRemoved(address notary);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _trackedDomain) {
        trackedDomain = _trackedDomain;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function allNotaries() external view returns (address[] memory) {
        return notaries.values();
    }

    function getNotary(uint256 _index) external view returns (address) {
        return notaries.at(_index);
    }

    function notariesAmount() external view returns (uint256) {
        return notaries.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addNotary(address _notary) internal returns (bool notaryAdded) {
        notaryAdded = notaries.add(_notary);
        if (notaryAdded) {
            emit DomainNotaryAdded(_notary);
        }
    }

    function _removeNotary(address _notary) internal returns (bool notaryRemoved) {
        notaryRemoved = notaries.remove(_notary);
        if (notaryRemoved) {
            emit DomainNotaryRemoved(_notary);
        }
    }

    function _isNotary(uint32 _domain, address _notary) internal view override returns (bool) {
        require(_domain == trackedDomain, "Wrong domain");
        return notaries.contains(_notary);
    }
}
