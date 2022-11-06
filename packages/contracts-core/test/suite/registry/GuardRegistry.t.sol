// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../tools/registry/GuardRegistryTools.t.sol";
import "../../tools/libs/EnumerableSetTools.t.sol";

// solhint-disable func-name-mixedcase
contract GuardRegistryTest is EnumerableSetTools, GuardRegistryTools {
    function setUp() public override {
        super.setUp();
        createExpectedStates();
        guardRegistry = new GuardRegistryHarness();
        require(GUARDS == ELEMENTS, "!elements");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           TESTS: ADD GUARD                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addGuard_multipleGuards() public {
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            expectGuardAdded({ guardIndex: i });
            guardRegistryAddGuard({ guardIndex: i, returnValue: true });
        }
    }

    function test_addGuard_twice() public {
        test_addGuard_multipleGuards();
        // Should not be possible to add the same Guard twice
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryAddGuard({ guardIndex: i, returnValue: false });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TESTS: REMOVE GUARD                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_removeGuard() public {
        test_addGuard_multipleGuards();
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            expectGuardRemoved({ guardIndex: i });
            guardRegistryRemoveGuard({ guardIndex: i, returnValue: true });
        }
    }

    function test_removeGuard_twice() public {
        test_removeGuard();
        // Should not be possible to remove the same Guard twice
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryRemoveGuard({ guardIndex: i, returnValue: false });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_getGuard_allGuards_getGuard() public {
        test_addGuard_multipleGuards();
        // Remove Guards one by one and check allGuards() and getGuard() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            address[] memory allGuards = guardRegistryAllGuards();
            assertEq(allGuards.length, expectedStates[i].length, "!length");
            for (uint256 j = 0; j < allGuards.length; ++j) {
                address expectedGuard = suiteGuard({ index: expectedStates[i][j] });
                assertEq(allGuards[j], expectedGuard, "!allGuards");
                assertEq(guardRegistryGetGuard(j), expectedGuard, "!getGuard");
            }
            guardRegistryRemoveGuard({ guardIndex: removalOrder[i], returnValue: true });
        }
    }

    // solhint-disable-next-line code-complexity
    function test_isGuard() public {
        for (uint256 j = 0; j < ELEMENTS; ++j) {
            assertFalse(guardRegistryIsGuard(j), "!isGuard: initial state");
        }
        // Add Guards one by one and check isNotary() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryAddGuard({ guardIndex: i, returnValue: true });
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                // Should return true only if Guard was added before
                assertEq(guardRegistryIsGuard(j), j <= i, "!isGuard: adding");
            }
        }
        // Remove Guards one by one and check isNotary() return values
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryRemoveGuard({ guardIndex: i, returnValue: true });
            for (uint256 j = 0; j < ELEMENTS; ++j) {
                // Should return true only if Guard was not removed before
                assertEq(guardRegistryIsGuard(j), j > i, "!isGuard: removing");
            }
        }
    }

    function test_guardsAmount() public {
        assertEq(guardRegistryGuardsAmount(), 0, "!guardsAmount: initial state");
        // Add notaries one by one and check guardsAmount() return value
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryAddGuard({ guardIndex: i, returnValue: true });
            // Added (i + 1) guards so far
            assertEq(guardRegistryGuardsAmount(), i + 1, "!guardsAmount: adding");
        }
        for (uint256 i = 0; i < ELEMENTS; ++i) {
            guardRegistryRemoveGuard({ guardIndex: i, returnValue: true });
            // Removed (i + 1) guards so far
            assertEq(guardRegistryGuardsAmount(), ELEMENTS - (i + 1), "!guardsAmount: removing");
        }
    }
}
