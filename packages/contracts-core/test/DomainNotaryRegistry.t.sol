// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "forge-std/Test.sol";

import { DomainNotaryRegistryHarness } from "./harnesses/DomainNotaryRegistryHarness.sol";

contract DomainNotaryRegistryTest is Test {
    DomainNotaryRegistryHarness internal registry;

    uint32 internal constant DOMAIN = 1000;

    address internal constant NOTARY_1 = address(1);
    address internal constant NOTARY_2 = address(2);
    address internal constant NOTARY_3 = address(3);
    address internal constant NOTARY_4 = address(4);

    event NotaryAdded(uint32 indexed domain, address notary);
    event NotaryRemoved(uint32 indexed domain, address notary);

    function setUp() public {
        registry = new DomainNotaryRegistryHarness(DOMAIN);
    }

    // solhint-disable func-name-mixedcase
    function test_addNotary_multipleNotaries() public {
        _checkAddNotary(NOTARY_1, true);
        _checkAddNotary(NOTARY_2, true);
        _checkAddNotary(NOTARY_3, true);
        _checkAddNotary(NOTARY_4, true);
    }

    function test_addNotary_twice() public {
        test_addNotary_multipleNotaries();
        _checkAddNotary(NOTARY_1, false);
        _checkAddNotary(NOTARY_2, false);
        _checkAddNotary(NOTARY_3, false);
        _checkAddNotary(NOTARY_4, false);
    }

    function test_removeNotary() public {
        test_addNotary_multipleNotaries();
        _checkRemoveNotary(NOTARY_3, true);
        _checkRemoveNotary(NOTARY_1, true);
        _checkRemoveNotary(NOTARY_4, true);
        _checkRemoveNotary(NOTARY_2, true);
    }

    function test_removeNotary_twice() public {
        test_removeNotary();
        _checkRemoveNotary(NOTARY_3, false);
        _checkRemoveNotary(NOTARY_1, false);
        _checkRemoveNotary(NOTARY_4, false);
        _checkRemoveNotary(NOTARY_2, false);
    }

    function test_notariesAmount() public {
        assertEq(registry.notariesAmount(), 0);
        test_addNotary_multipleNotaries();
        assertEq(registry.notariesAmount(), 4);
        _checkRemoveNotary(NOTARY_3, true);
        assertEq(registry.notariesAmount(), 3);
        _checkRemoveNotary(NOTARY_1, true);
        assertEq(registry.notariesAmount(), 2);
        _checkRemoveNotary(NOTARY_4, true);
        assertEq(registry.notariesAmount(), 1);
        _checkRemoveNotary(NOTARY_2, true);
        assertEq(registry.notariesAmount(), 0);
    }

    function test_getNotary() public {
        test_addNotary_multipleNotaries();
        assertEq(registry.getNotary(0), NOTARY_1);
        assertEq(registry.getNotary(1), NOTARY_2);
        assertEq(registry.getNotary(2), NOTARY_3);
        assertEq(registry.getNotary(3), NOTARY_4);
        _checkRemoveNotary(NOTARY_3, true);
        assertEq(registry.getNotary(0), NOTARY_1);
        assertEq(registry.getNotary(1), NOTARY_2);
        assertEq(registry.getNotary(2), NOTARY_4);
        _checkRemoveNotary(NOTARY_1, true);
        assertEq(registry.getNotary(0), NOTARY_4);
        assertEq(registry.getNotary(1), NOTARY_2);
        _checkRemoveNotary(NOTARY_4, true);
        assertEq(registry.getNotary(0), NOTARY_2);
    }

    function test_allNotaries() public {
        test_addNotary_multipleNotaries();
        {
            address[] memory notaries = registry.allNotaries();
            assertEq(notaries.length, 4);
            assertEq(notaries[0], NOTARY_1);
            assertEq(notaries[1], NOTARY_2);
            assertEq(notaries[2], NOTARY_3);
            assertEq(notaries[3], NOTARY_4);
        }
        _checkRemoveNotary(NOTARY_3, true);
        {
            address[] memory notaries = registry.allNotaries();
            assertEq(notaries.length, 3);
            assertEq(notaries[0], NOTARY_1);
            assertEq(notaries[1], NOTARY_2);
            assertEq(notaries[2], NOTARY_4);
        }
        _checkRemoveNotary(NOTARY_1, true);
        {
            address[] memory notaries = registry.allNotaries();
            assertEq(notaries.length, 2);
            assertEq(notaries[0], NOTARY_4);
            assertEq(notaries[1], NOTARY_2);
        }
        _checkRemoveNotary(NOTARY_4, true);
        {
            address[] memory notaries = registry.allNotaries();
            assertEq(notaries.length, 1);
            assertEq(notaries[0], NOTARY_2);
        }
    }

    function test_isNotary_notNotary() public {
        _checkAddNotary(NOTARY_1, true);
        assertFalse(registry.isNotary(DOMAIN, NOTARY_2));
    }

    function test_isNotary_wrongDomain() public {
        _checkAddNotary(NOTARY_1, true);
        vm.expectRevert("!localDomain");
        registry.isNotary(DOMAIN + 1, NOTARY_1);
    }

    function _checkAddNotary(address _notary, bool _added) internal {
        if (_added) {
            vm.expectEmit(true, true, true, true);
            emit NotaryAdded(DOMAIN, _notary);
        }
        assertEq(registry.addNotary(_notary), _added);
        assertTrue(registry.isNotary(DOMAIN, _notary));
    }

    function _checkRemoveNotary(address _notary, bool _removed) internal {
        if (_removed) {
            vm.expectEmit(true, true, true, true);
            emit NotaryRemoved(DOMAIN, _notary);
        }
        assertEq(registry.removeNotary(_notary), _removed);
        assertFalse(registry.isNotary(DOMAIN, _notary));
    }
}
