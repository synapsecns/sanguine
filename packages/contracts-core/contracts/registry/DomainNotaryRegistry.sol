// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AbstractNotaryRegistry } from "./AbstractNotaryRegistry.sol";
import { DomainContext } from "../context/DomainContext.sol";

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice A Registry to keep track of Notaries on a single domain.
 */
abstract contract DomainNotaryRegistry is AbstractNotaryRegistry, DomainContext {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // All active notaries for the tracked chain
    EnumerableSet.AddressSet internal notaries;

    // gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Ensures that there is at least one active Notary.
     */
    modifier haveActiveNotary() {
        require(notariesAmount() != 0, "!notaries");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns addresses of all Notaries.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of notaries is large (see EnumerableSet.values())
     */
    function allNotaries() external view returns (address[] memory) {
        return notaries.values();
    }

    /**
     * @notice Returns i-th Notary. O(1)
     * @dev Will revert if index is out of range
     */
    function getNotary(uint256 _index) public view returns (address) {
        return notaries.at(_index);
    }

    /**
     * @notice Returns amount of active notaries. O(1)
     */
    function notariesAmount() public view returns (uint256) {
        return notaries.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _addNotary(uint32 _domain, address _notary)
        internal
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _addNotary(_notary);
    }

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     */
    function _addNotary(address _notary) internal returns (bool notaryAdded) {
        notaryAdded = notaries.add(_notary);
        if (notaryAdded) {
            emit NotaryAdded(_localDomain(), _notary);
        }
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _removeNotary(uint32 _domain, address _notary)
        internal
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _removeNotary(_notary);
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     */
    function _removeNotary(address _notary) internal returns (bool notaryRemoved) {
        notaryRemoved = notaries.remove(_notary);
        if (notaryRemoved) {
            emit NotaryRemoved(_localDomain(), _notary);
        }
    }

    /**
     * @notice Returns whether given address is a notary for the tracked domain.
     * @dev Reverts if domain doesn't match the tracked domain.
     */
    function _isNotary(uint32 _domain, address _account)
        internal
        view
        override
        onlyLocalDomain(_domain)
        returns (bool)
    {
        return _isNotary(_account);
    }

    /**
     * @notice Returns whether given address is a notary for the tracked domain.
     */
    function _isNotary(address _account) internal view returns (bool) {
        return notaries.contains(_account);
    }
}
