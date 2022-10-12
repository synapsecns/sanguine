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
    ▏*║                          TESTS: ADD NOTARY                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary_multipleNotaries() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < ELEMENTS; ++i) {
                expectNotaryAdded({ domain: domain, notaryIndex: i });
                globalNotaryRegistryAddNotary({
                    domain: domain,
                    notaryIndex: i,
                    returnValue: true
                });
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
}
