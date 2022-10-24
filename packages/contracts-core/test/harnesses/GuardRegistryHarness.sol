// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { GuardRegistry } from "../../contracts/registry/GuardRegistry.sol";

contract GuardRegistryHarness is GuardRegistry {
    function addGuard(address _guard) public returns (bool) {
        return _addGuard(_guard);
    }

    function removeGuard(address _guard) public returns (bool) {
        return _removeGuard(_guard);
    }

    function isGuard(address _guard) public view returns (bool) {
        return _isGuard(_guard);
    }
}
