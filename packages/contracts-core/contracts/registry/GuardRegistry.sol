// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AbstractGuardRegistry } from "./AbstractGuardRegistry.sol";

import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice A Registry to keep track of all Guards.
 */
contract GuardRegistry is AbstractGuardRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    EnumerableSet.AddressSet internal guards;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line var-name-mixedcase
    uint256[49] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event GuardAdded(address guard);

    event GuardRemoved(address guard);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function allGuards() external view returns (address[] memory) {
        return guards.values();
    }

    function getGuard(uint256 _index) external view returns (address) {
        return guards.at(_index);
    }

    function guardsAmount() external view returns (uint256) {
        return guards.length();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addGuard(address _guard) internal returns (bool guardAdded) {
        guardAdded = guards.add(_guard);
        if (guardAdded) {
            emit GuardAdded(_guard);
        }
    }

    function _removeGuard(address _guard) internal returns (bool guardRemoved) {
        guardRemoved = guards.remove(_guard);
        if (guardRemoved) {
            emit GuardRemoved(_guard);
        }
    }

    function _isGuard(address _guard) internal view override returns (bool) {
        return guards.contains(_guard);
    }
}
