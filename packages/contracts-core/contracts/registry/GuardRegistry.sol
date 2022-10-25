// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AbstractGuardRegistry } from "./AbstractGuardRegistry.sol";

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice A Registry to keep track of Guards on all domains.
 * @dev It is assumed that the Guard signature is valid on all chains.
 */
contract GuardRegistry is AbstractGuardRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // All active guards
    EnumerableSet.AddressSet internal guards;

    // gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns addresses of all Guards.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of notaries is large (see EnumerableSet.values())
     */
    function allGuards() external view returns (address[] memory) {
        return guards.values();
    }

    /**
     * @notice Returns i-th Guard. O(1)
     * @dev Will revert if index is out of range
     */
    function getGuard(uint256 _index) external view returns (address) {
        return guards.at(_index);
    }

    /**
     * @notice Returns amount of active guards. O(1)
     */
    function guardsAmount() external view returns (uint256) {
        return guards.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new guard, emits an event only if guard was added.
     */
    function _addGuard(address _guard) internal override returns (bool guardAdded) {
        guardAdded = guards.add(_guard);
        if (guardAdded) {
            emit GuardAdded(_guard);
        }
    }

    /**
     * @notice Tries to remove a guard, emits an event only if guard was removed.
     */
    function _removeGuard(address _guard) internal override returns (bool guardRemoved) {
        guardRemoved = guards.remove(_guard);
        if (guardRemoved) {
            emit GuardRemoved(_guard);
        }
    }

    /**
     * @notice Returns whether given address is a guard.
     */
    function _isGuard(address _account) internal view override returns (bool) {
        return guards.contains(_account);
    }
}
