// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";
import { GuardRegistryHarness } from "../../harnesses/registry/GuardRegistryHarness.t.sol";

abstract contract GuardRegistryTools is SynapseTestSuite {
    GuardRegistryHarness internal guardRegistry;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectGuardAdded(uint256 guardIndex) public {
        vm.expectEmit(true, true, true, true);
        emit GuardAdded(suiteGuard(guardIndex));
    }

    function expectGuardRemoved(uint256 guardIndex) public {
        vm.expectEmit(true, true, true, true);
        emit GuardRemoved(suiteGuard(guardIndex));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Try to add a given Guard and expect a revert
    function guardRegistryAddGuard(uint256 guardIndex, bytes memory revertMessage) public {
        address guard = suiteGuard(guardIndex);
        vm.expectRevert(revertMessage);
        guardRegistry.addGuard(guard);
    }

    // Try to remove a given Guard and expect a revert
    function guardRegistryRemoveGuard(uint256 guardIndex, bytes memory revertMessage) public {
        address guard = suiteGuard(guardIndex);
        vm.expectRevert(revertMessage);
        guardRegistry.removeGuard(guard);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Add a given Guard and check the return value
    function guardRegistryAddGuard(uint256 guardIndex, bool returnValue) public {
        assertEq(guardRegistry.addGuard(suiteGuard(guardIndex)), returnValue, "!returnValue");
    }

    // Remove a given Guard and check the return value
    function guardRegistryRemoveGuard(uint256 guardIndex, bool returnValue) public {
        assertEq(guardRegistry.removeGuard(suiteGuard(guardIndex)), returnValue, "!returnValue");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TRIGGER VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function guardRegistryAllGuards() public view returns (address[] memory) {
        return guardRegistry.allGuards();
    }

    function guardRegistryGetGuard(uint256 guardIndex) public view returns (address) {
        return guardRegistry.getGuard(guardIndex);
    }

    function guardRegistryIsGuard(uint256 guardIndex) public view returns (bool) {
        return guardRegistry.isGuard(suiteGuard(guardIndex));
    }

    function guardRegistryGuardsAmount() public view returns (uint256) {
        return guardRegistry.guardsAmount();
    }
}
