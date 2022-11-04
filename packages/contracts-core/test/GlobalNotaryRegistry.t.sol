// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Test.sol";

import { GlobalNotaryRegistryHarness } from "./harnesses/GlobalNotaryRegistryHarness.sol";

// solhint-disable func-name-mixedcase
contract GlobalNotaryRegistryTest is Test {
    address internal constant NOTARY_1 = address(1);
    address internal constant NOTARY_2 = address(2);
    address internal constant NOTARY_3 = address(3);
    address internal constant NOTARY_4 = address(4);

    uint32 internal constant DOMAIN_1 = 1234;
    uint32 internal constant DOMAIN_2 = 4321;

    GlobalNotaryRegistryHarness internal registry;

    event NotaryAdded(uint32 indexed domain, address indexed notary);
    event NotaryRemoved(uint32 indexed domain, address indexed notary);

    function setUp() public {
        registry = new GlobalNotaryRegistryHarness();
    }

    function test_getDomain() public {
        _checkAddNotary(DOMAIN_1, NOTARY_1, true, 1);
        assertEq(registry.getDomain(0), DOMAIN_1, "!getDomain (0)");
        _checkAddNotary(DOMAIN_2, NOTARY_2, true, 2);
        assertEq(registry.getDomain(1), DOMAIN_2, "!getDomain (1)");

        _checkRemoveNotary(DOMAIN_1, NOTARY_1, true, 1);
        assertEq(registry.getDomain(0), DOMAIN_2, "!getDomain (2)");
        _checkRemoveNotary(DOMAIN_2, NOTARY_2, true, 0);
    }

    function test_allDomains() public {
        test_addNotary_multipleDomains();
        uint32[] memory domains = registry.allDomains();
        assertEq(domains.length, 2, "!length (0)");
        assertEq(domains[0], DOMAIN_1, "!domains[0] (0)");
        assertEq(domains[1], DOMAIN_2, "!domains[1] (0)");

        _checkRemoveNotary(DOMAIN_1, NOTARY_1, true, 1);
        domains = registry.allDomains();
        assertEq(domains.length, 1, "!length (1)");
        assertEq(domains[0], DOMAIN_2, "!domains[0] (1)");

        _checkRemoveNotary(DOMAIN_2, NOTARY_2, true, 0);
        domains = registry.allDomains();
        assertEq(domains.length, 0, "!length (0)");
    }

    function test_addNotary_multipleDomains() public {
        _checkAddNotary(DOMAIN_1, NOTARY_1, true, 1);
        _checkAddNotary(DOMAIN_2, NOTARY_1, false, 1);

        _checkAddNotary(DOMAIN_2, NOTARY_2, true, 2);
        _checkAddNotary(DOMAIN_1, NOTARY_2, false, 2);
    }

    function test_addNotary_afterDeleting() public {
        test_addNotary_multipleDomains();
        _checkRemoveNotary(DOMAIN_1, NOTARY_1, true, 1);
        _checkRemoveNotary(DOMAIN_2, NOTARY_2, true, 0);

        _checkAddNotary(DOMAIN_2, NOTARY_1, true, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_1, false, 1);

        _checkAddNotary(DOMAIN_1, NOTARY_2, true, 2);
        _checkAddNotary(DOMAIN_2, NOTARY_1, false, 2);
    }

    function test_addNotary_multipleNotaries() public {
        _checkAddNotary(DOMAIN_1, NOTARY_1, true, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_2, true, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_3, true, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_4, true, 1);
    }

    function test_addNotary_twice() public {
        test_addNotary_multipleNotaries();
        _checkAddNotary(DOMAIN_1, NOTARY_1, false, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_2, false, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_3, false, 1);
        _checkAddNotary(DOMAIN_1, NOTARY_4, false, 1);
    }

    function test_removeNotary() public {
        test_addNotary_multipleNotaries();
        _checkRemoveNotary(DOMAIN_1, NOTARY_3, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_1, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_2, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_4, true, 0);
    }

    function test_removeNotary_twice() public {
        test_addNotary_multipleNotaries();
        _checkRemoveNotary(DOMAIN_1, NOTARY_3, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_3, false, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_1, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_1, false, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_2, true, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_2, false, 1);
        _checkRemoveNotary(DOMAIN_1, NOTARY_4, true, 0);
        _checkRemoveNotary(DOMAIN_1, NOTARY_4, false, 0);
    }

    function _checkAddNotary(
        uint32 _domain,
        address _notary,
        bool _added,
        uint256 _activeDomains
    ) internal {
        if (_added) {
            vm.expectEmit(true, true, true, true);
            emit NotaryAdded(_domain, _notary);
        }
        assertEq(registry.addNotary(_domain, _notary), _added);
        if (_added) {
            assertTrue(registry.isNotary(_domain, _notary));
        }
        assertEq(registry.domainsAmount(), _activeDomains, "!domainsAmount");
    }

    function _checkRemoveNotary(
        uint32 _domain,
        address _notary,
        bool _removed,
        uint256 _activeDomains
    ) internal {
        if (_removed) {
            vm.expectEmit(true, true, true, true);
            emit NotaryRemoved(_domain, _notary);
        }
        assertEq(registry.removeNotary(_domain, _notary), _removed);
        assertFalse(registry.isNotary(_domain, _notary));
        assertEq(registry.domainsAmount(), _activeDomains, "!domainsAmount");
    }
}
