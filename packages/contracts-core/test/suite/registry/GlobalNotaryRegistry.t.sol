// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/registry/GlobalNotaryRegistryTools.t.sol";
import "../../tools/libs/EnumerableSetTools.t.sol";

// solhint-disable func-name-mixedcase
contract GlobalNotaryRegistryTest is EnumerableSetTools, GlobalNotaryRegistryTools {
    function setUp() public override {
        super.setUp();
        createExpectedStates();
        globalNotaryRegistry = new GlobalNotaryRegistryHarness();
        require(NOTARIES_PER_CHAIN == ELEMENTS, "!elements");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    TESTS: REVERT WHEN DOMAIN == 0                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary_revert_domainZero(address notary) public {
        vm.expectRevert("!domain: zero");
        globalNotaryRegistry.addNotary(0, notary);
    }

    function test_removeNotary_revert_domainZero(address notary) public {
        vm.expectRevert("!domain: zero");
        globalNotaryRegistry.removeNotary(0, notary);
    }

    function test_isNotary_revert_domainZero(address notary) public {
        vm.expectRevert("!domain: zero");
        globalNotaryRegistry.isNotary(0, notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: ADD NOTARY                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary_multipleNotaries() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            // Check amount of domains before adding
            assertEq(globalNotaryRegistryDomainsAmount(), d, "!domainsAmount: before adding");
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                expectNotaryAdded({ domain: domain, notaryIndex: i });
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: true
                });
                // Check amount of domains after adding
                assertEq(
                    globalNotaryRegistryDomainsAmount(),
                    d + 1,
                    "!domainsAmount: after adding"
                );
            }
        }
    }

    function test_addNotary_twice() public {
        test_addNotary_multipleNotaries();
        // Should not be possible to add the same Notary twice for the same domain
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: false
                });
                // Amount of domains should not change
                assertEq(
                    globalNotaryRegistryDomainsAmount(),
                    DOMAINS,
                    "!domainsAmount: after readding"
                );
            }
        }
    }

    function test_addNotary_sameNotaryTwoDomains() public {
        test_addNotary_multipleNotaries();
        // Should not be possible to add the same Notary twice for a different domain
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                // Fetch notary address, that is active on another domain
                address notary = suiteForeignNotary({ domain: domain, index: i });
                // Should not be possible to add this Notary for given domain
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notary: notary,
                    returnValue: false
                });
                // Amount of domains should not change
                assertEq(
                    globalNotaryRegistryDomainsAmount(),
                    DOMAINS,
                    "!domainsAmount: after reading"
                );
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: REMOVE NOTARY                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_removeNotary() public {
        test_addNotary_multipleNotaries();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                expectNotaryRemoved({ domain: domain, notaryIndex: i });
                globalNotaryRegistryRemoveNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: true
                });
                // Amount of domains should not change, as long as there's an active notary left
                if (i != ELEMENTS - 1) {
                    assertEq(
                        globalNotaryRegistryDomainsAmount(),
                        DOMAINS - d,
                        "!domainsAmount: deleting, active notary left"
                    );
                } else {
                    assertEq(
                        globalNotaryRegistryDomainsAmount(),
                        DOMAINS - d - 1,
                        "!domainsAmount: deleting, all domain notaries removed"
                    );
                }
            }
        }
    }

    function test_removeNotary_twice() public {
        test_removeNotary();
        // Should not be possible to remove the same Notary twice
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryRemoveNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: false
                });
                // Amount of domains should not change
                assertEq(
                    globalNotaryRegistryDomainsAmount(),
                    0,
                    "!domainsAmount: after redeleting"
                );
            }
        }
    }

    function test_removeNotary_addSameDomain() public {
        test_removeNotary();
        // Should be possible to add Notary back on the same domain once deleted
        test_addNotary_multipleNotaries();
    }

    function test_removeNotary_addAnotherDomain() public {
        test_removeNotary();
        // Should be possible to add Notary back on another domain once deleted
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            // Check amount of domains before adding
            assertEq(globalNotaryRegistryDomainsAmount(), d, "!domainsAmount: before adding");
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                // Fetch notary address, that used to be active on another domain
                address notary = suiteForeignNotary({ domain: domain, index: i });
                // Since notary is inactive, we should be able to add it on given domain
                expectNotaryAdded(domain, notary);
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notary: notary,
                    returnValue: true
                });
                // Check amount of domains after adding
                assertEq(
                    globalNotaryRegistryDomainsAmount(),
                    d + 1,
                    "!domainsAmount: after adding"
                );
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_getNotary_allNotaries_getNotary() public {
        test_addNotary_multipleNotaries();
        // Remove Notaries one by one and check allNotaries() and getNotary() return values
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                address[] memory allNotaries = globalNotaryRegistryAllNotaries(domain);
                assertEq(allNotaries.length, expectedStates[i].length, "!length");
                for (uint256 j = 0; j < allNotaries.length; ++j) {
                    address expectedNotary = suiteNotary({
                        domain: domain,
                        index: expectedStates[i][j]
                    });
                    assertEq(allNotaries[j], expectedNotary, "!allNotaries");
                    assertEq(
                        globalNotaryRegistryGetNotary({ domain: domain, notaryIndex: j }),
                        expectedNotary,
                        "!getNotary"
                    );
                }
                globalNotaryRegistryRemoveNotary({
                    domain: domain,
                    notaryIndex: removalOrder[i],
                    returnValue: true
                });
            }
        }
    }

    function test_isNotary_afterAdding() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            // Check initial state
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                assertFalse(
                    globalNotaryRegistryIsNotary({ domain: domain, notaryIndex: j }),
                    "!isNotary: initial state"
                );
            }
            // Add Notaries one by one and check isNotary() return values
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: true
                });
                for (uint256 j = 0; j < ELEMENTS; ++j) {
                    // Should return true only if Notary was added before
                    assertEq(
                        globalNotaryRegistryIsNotary({ domain: domain, notaryIndex: j }),
                        j <= i,
                        "!isNotary: adding"
                    );
                }
            }
        }
    }

    function test_isNotary_afterRemoving() public {
        test_addNotary_multipleNotaries();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            // Check initial state
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                assertTrue(
                    globalNotaryRegistryIsNotary({ domain: domain, notaryIndex: j }),
                    "!isNotary: initial state"
                );
            }
            // Remove Notaries one by one and check isNotary() return values
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryRemoveNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: true
                });
                for (uint256 j = 0; j < ELEMENTS; ++j) {
                    // Should return true only if Notary was not removed before
                    assertEq(
                        globalNotaryRegistryIsNotary({ domain: domain, notaryIndex: j }),
                        j > i,
                        "!isNotary: removing"
                    );
                }
            }
        }
    }

    function test_notariesAmount_afterAdding() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            // Add notaries one by one and check notariesAmount() return value
            uint32 domain = domains[d];
            assertEq(
                globalNotaryRegistryNotariesAmount(domain),
                0,
                "!notariesAmount: initial state"
            );
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryAddNotary(domain, i, true);
                // Added (i + 1) notaries so far
                assertEq(
                    globalNotaryRegistryNotariesAmount(domain),
                    i + 1,
                    "!notariesAmount: adding"
                );
            }
        }
    }

    function test_notariesAmount_afterRemoving() public {
        test_addNotary_multipleNotaries();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            // Remove notaries one by one and check notariesAmount() return value
            uint32 domain = domains[d];
            assertEq(
                globalNotaryRegistryNotariesAmount(domain),
                ELEMENTS,
                "!notariesAmount: initial state"
            );
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                globalNotaryRegistryRemoveNotary(domain, i, true);
                // Removed (i + 1) notaries so far
                assertEq(
                    globalNotaryRegistryNotariesAmount(domain),
                    ELEMENTS - (i + 1),
                    "!notariesAmount: removing"
                );
            }
        }
    }

    function test_allDomains_getDomain_domainsAmount() public {
        // Add first Notary
        globalNotaryRegistryAddNotary({ domain: DOMAIN_LOCAL, notaryIndex: 0, returnValue: true });
        // Checks # 1
        assertEq(globalNotaryRegistryDomainsAmount(), 1, "!domainsAmount (#1)");
        assertEq(globalNotaryRegistryGetDomain(0), DOMAIN_LOCAL, "!getDomain(0) (#1)");
        uint32[] memory allDomains = globalNotaryRegistryAllDomains();
        assertEq(allDomains.length, 1, "!length (#1)");
        assertEq(allDomains[0], DOMAIN_LOCAL, "!domains[0] (#1)");
        // Add Notary on another domain
        globalNotaryRegistryAddNotary({ domain: DOMAIN_REMOTE, notaryIndex: 0, returnValue: true });
        // Checks # 2
        assertEq(globalNotaryRegistryDomainsAmount(), 2, "!domainsAmount (#2)");
        assertEq(globalNotaryRegistryGetDomain(0), DOMAIN_LOCAL, "!getDomain(0) (#2)");
        assertEq(globalNotaryRegistryGetDomain(1), DOMAIN_REMOTE, "!getDomain(1) (#2)");
        allDomains = globalNotaryRegistryAllDomains();
        assertEq(allDomains.length, 2, "!length (#2)");
        assertEq(allDomains[0], DOMAIN_LOCAL, "!domains[0] (#2)");
        assertEq(allDomains[1], DOMAIN_REMOTE, "!domains[1] (#2)");
        // Remove Notary on first domain
        globalNotaryRegistryRemoveNotary({
            domain: DOMAIN_LOCAL,
            notaryIndex: 0,
            returnValue: true
        });
        // Checks #3
        assertEq(globalNotaryRegistryDomainsAmount(), 1, "!domainsAmount (#3)");
        assertEq(globalNotaryRegistryGetDomain(0), DOMAIN_REMOTE, "!getDomain (#3)");
        allDomains = globalNotaryRegistryAllDomains();
        assertEq(allDomains.length, 1, "!length (#3)");
        assertEq(allDomains[0], DOMAIN_REMOTE, "!domains[0] (#3)");
    }
}
