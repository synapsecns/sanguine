// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/registry/DomainNotaryRegistryTools.t.sol";
import "../../tools/libs/EnumerableSetTools.t.sol";

// solhint-disable func-name-mixedcase
contract DomainNotaryRegistryTest is EnumerableSetTools, DomainNotaryRegistryTools {
    function setUp() public override {
        super.setUp();
        createExpectedStates();
        domainNotaryRegistry = new DomainNotaryRegistryHarness(DOMAIN_LOCAL);
        require(NOTARIES_PER_CHAIN == ELEMENTS, "!elements");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TESTS: ADD NOTARY                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addNotary_multipleNotaries() public {
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            expectNotaryAdded({ notaryIndex: i });
            domainNotaryRegistryAddNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
        }
    }

    function test_addNotary_twice() public {
        test_addNotary_multipleNotaries();
        // Should not be possible to add the same Notary twice
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryAddNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: false
            });
        }
    }

    function test_addNotary_revert_wrongDomain() public {
        // Specifying wrong domain leads to reverting
        domainNotaryRegistryAddNotary({
            domain: DOMAIN_REMOTE,
            notaryIndex: 0,
            revertMessage: "!localDomain"
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: REMOVE NOTARY                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_removeNotary() public {
        test_addNotary_multipleNotaries();
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            expectNotaryRemoved({ notaryIndex: i });
            domainNotaryRegistryRemoveNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
        }
    }

    function test_removeNotary_twice() public {
        test_removeNotary();
        // Should not be possible to remove the same Notary twice
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryRemoveNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: false
            });
        }
    }

    function test_removeNotary_revert_wrongDomain() public {
        domainNotaryRegistryRemoveNotary({
            domain: DOMAIN_REMOTE,
            notaryIndex: 0,
            revertMessage: "!localDomain"
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_getNotary_allNotaries_getNotary() public {
        test_addNotary_multipleNotaries();
        // Remove Notaries one by one and check allNotaries() and getNotary() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            address[] memory allNotaries = domainNotaryRegistryAllNotaries();
            assertEq(allNotaries.length, expectedStates[i].length, "!length");
            for (uint256 j = 0; j < allNotaries.length; ++j) {
                address expectedNotary = suiteNotary({
                    domain: DOMAIN_LOCAL,
                    index: expectedStates[i][j]
                });
                assertEq(allNotaries[j], expectedNotary, "!allNotaries");
                assertEq(domainNotaryRegistryGetNotary(j), expectedNotary, "!getNotary");
            }
            domainNotaryRegistryRemoveNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: removalOrder[i],
                returnValue: true
            });
        }
    }

    function test_isNotary_afterAdding() public {
        for (uint256 j = 0; j < ELEMENTS; ++j) {
            assertFalse(
                domainNotaryRegistryIsNotary({ domain: DOMAIN_LOCAL, notaryIndex: j }),
                "!isNotary: initial state"
            );
        }
        // Add Notaries one by one and check isNotary() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryAddNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                // Should return true only if Notary was added before
                assertEq(
                    domainNotaryRegistryIsNotary({ domain: DOMAIN_LOCAL, notaryIndex: j }),
                    j <= i,
                    "!isNotary: adding"
                );
            }
        }
    }

    function test_isNotary_afterRemoving() public {
        test_addNotary_multipleNotaries();
        for (uint256 j = 0; j < ELEMENTS; ++j) {
            assertTrue(
                domainNotaryRegistryIsNotary({ domain: DOMAIN_LOCAL, notaryIndex: j }),
                "!isNotary: initial state"
            );
        }
        // Remove Notaries one by one and check isNotary() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryRemoveNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                // Should return true only if Notary was not removed before
                assertEq(
                    domainNotaryRegistryIsNotary({ domain: DOMAIN_LOCAL, notaryIndex: j }),
                    j > i,
                    "!isNotary: removing"
                );
            }
        }
    }

    function test_isNotary_revert_wrongDomain() public {
        vm.expectRevert("!localDomain");
        // Specifying wrong domain leads to reverting in the view function as well
        domainNotaryRegistryIsNotary({ domain: DOMAIN_REMOTE, notaryIndex: 0 });
    }

    function test_notariesAmount() public {
        assertEq(domainNotaryRegistryNotariesAmount(), 0, "!notariesAmount: initial state");
        // Add notaries one by one and check notariesAmount() return value
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryAddNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
            // Added (i + 1) notaries so far
            assertEq(domainNotaryRegistryNotariesAmount(), i + 1, "!notariesAmount: adding");
        }
        // Remove notaries one by one and check notariesAmount() return value
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            domainNotaryRegistryRemoveNotary({
                domain: DOMAIN_LOCAL,
                notaryIndex: i,
                returnValue: true
            });
            // Removed (i + 1) notaries so far
            assertEq(
                domainNotaryRegistryNotariesAmount(),
                ELEMENTS - (i + 1),
                "!notariesAmount: removing"
            );
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   TEST: DOMAIN SPECIFIC FUNCTIONS                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_domainSpecific() public {
        assertEq(
            domainNotaryRegistryIsNotary({ notaryIndex: 0 }),
            false,
            "!isNotary: initial state"
        );

        expectNotaryAdded({ notaryIndex: 0 });
        domainNotaryRegistryAddNotary({ notaryIndex: 0, returnValue: true });
        assertEq(domainNotaryRegistryIsNotary({ notaryIndex: 0 }), true, "!isNotary: adding");
        // Should not add the same notary twice
        domainNotaryRegistryAddNotary({ notaryIndex: 0, returnValue: false });

        expectNotaryRemoved({ notaryIndex: 0 });
        domainNotaryRegistryRemoveNotary({ notaryIndex: 0, returnValue: true });
        assertEq(domainNotaryRegistryIsNotary({ notaryIndex: 0 }), false, "!isNotary: removing");
        // Should not delete the same notary twice
        domainNotaryRegistryRemoveNotary({ notaryIndex: 0, returnValue: false });
    }
}
