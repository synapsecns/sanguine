// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { GuardRegistryHarness } from "./harnesses/GuardRegistryHarness.sol";

contract GuardRegistryTest is Test {
    GuardRegistryHarness internal registry;

    address internal constant GUARD_1 = address(1);
    address internal constant GUARD_2 = address(2);
    address internal constant GUARD_3 = address(3);
    address internal constant GUARD_4 = address(4);

    event GuardAdded(address indexed guard);
    event GuardRemoved(address indexed guard);

    function setUp() public {
        registry = new GuardRegistryHarness();
    }

    // solhint-disable func-name-mixedcase
    function test_addGuard_multipleGuards() public {
        _checkAddGuard(GUARD_1, true);
        _checkAddGuard(GUARD_2, true);
        _checkAddGuard(GUARD_3, true);
        _checkAddGuard(GUARD_4, true);
    }

    function test_addGuard_twice() public {
        test_addGuard_multipleGuards();
        _checkAddGuard(GUARD_1, false);
        _checkAddGuard(GUARD_2, false);
        _checkAddGuard(GUARD_3, false);
        _checkAddGuard(GUARD_4, false);
    }

    function test_removeGuard() public {
        test_addGuard_multipleGuards();
        _checkRemoveGuard(GUARD_3, true);
        _checkRemoveGuard(GUARD_1, true);
        _checkRemoveGuard(GUARD_4, true);
        _checkRemoveGuard(GUARD_2, true);
    }

    function test_removeGuard_twice() public {
        test_removeGuard();
        _checkRemoveGuard(GUARD_3, false);
        _checkRemoveGuard(GUARD_1, false);
        _checkRemoveGuard(GUARD_4, false);
        _checkRemoveGuard(GUARD_2, false);
    }

    function test_guardsAmount() public {
        assertEq(registry.guardsAmount(), 0);
        test_addGuard_multipleGuards();
        assertEq(registry.guardsAmount(), 4);
        _checkRemoveGuard(GUARD_3, true);
        assertEq(registry.guardsAmount(), 3);
        _checkRemoveGuard(GUARD_1, true);
        assertEq(registry.guardsAmount(), 2);
        _checkRemoveGuard(GUARD_4, true);
        assertEq(registry.guardsAmount(), 1);
        _checkRemoveGuard(GUARD_2, true);
        assertEq(registry.guardsAmount(), 0);
    }

    function test_getGuard() public {
        test_addGuard_multipleGuards();
        assertEq(registry.getGuard(0), GUARD_1);
        assertEq(registry.getGuard(1), GUARD_2);
        assertEq(registry.getGuard(2), GUARD_3);
        assertEq(registry.getGuard(3), GUARD_4);
        _checkRemoveGuard(GUARD_3, true);
        assertEq(registry.getGuard(0), GUARD_1);
        assertEq(registry.getGuard(1), GUARD_2);
        assertEq(registry.getGuard(2), GUARD_4);
        _checkRemoveGuard(GUARD_1, true);
        assertEq(registry.getGuard(0), GUARD_4);
        assertEq(registry.getGuard(1), GUARD_2);
        _checkRemoveGuard(GUARD_4, true);
        assertEq(registry.getGuard(0), GUARD_2);
    }

    function test_isGuard_notGuard() public {
        _checkAddGuard(GUARD_1, true);
        assertFalse(registry.isGuard(GUARD_2));
    }

    function test_allGuards() public {
        test_addGuard_multipleGuards();
        {
            address[] memory guards = registry.allGuards();
            assertEq(guards.length, 4);
            assertEq(guards[0], GUARD_1);
            assertEq(guards[1], GUARD_2);
            assertEq(guards[2], GUARD_3);
            assertEq(guards[3], GUARD_4);
        }
        _checkRemoveGuard(GUARD_3, true);
        {
            address[] memory guards = registry.allGuards();
            assertEq(guards.length, 3);
            assertEq(guards[0], GUARD_1);
            assertEq(guards[1], GUARD_2);
            assertEq(guards[2], GUARD_4);
        }
        _checkRemoveGuard(GUARD_1, true);
        {
            address[] memory guards = registry.allGuards();
            assertEq(guards.length, 2);
            assertEq(guards[0], GUARD_4);
            assertEq(guards[1], GUARD_2);
        }
        _checkRemoveGuard(GUARD_4, true);
        {
            address[] memory guards = registry.allGuards();
            assertEq(guards.length, 1);
            assertEq(guards[0], GUARD_2);
        }
    }

    function _checkAddGuard(address _guard, bool _added) internal {
        if (_added) {
            vm.expectEmit(true, true, true, true);
            emit GuardAdded(_guard);
        }
        assertEq(registry.addGuard(_guard), _added);
        assertTrue(registry.isGuard(_guard));
    }

    function _checkRemoveGuard(address _guard, bool _removed) internal {
        if (_removed) {
            vm.expectEmit(true, true, true, true);
            emit GuardRemoved(_guard);
        }
        assertEq(registry.removeGuard(_guard), _removed);
        assertFalse(registry.isGuard(_guard));
    }
}
