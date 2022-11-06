// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";
import {
    GlobalNotaryRegistryHarness
} from "../../harnesses/registry/GlobalNotaryRegistryHarness.t.sol";

abstract contract GlobalNotaryRegistryTools is SynapseTestSuite {
    GlobalNotaryRegistryHarness internal globalNotaryRegistry;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectNotaryAdded(uint32 domain, uint256 notaryIndex) public {
        vm.expectEmit(true, true, true, true);
        emit NotaryAdded(domain, suiteNotary(domain, notaryIndex));
    }

    function expectNotaryAdded(uint32 domain, address notary) public {
        vm.expectEmit(true, true, true, true);
        emit NotaryAdded(domain, notary);
    }

    function expectNotaryRemoved(uint32 domain, uint256 notaryIndex) public {
        vm.expectEmit(true, true, true, true);
        emit NotaryRemoved(domain, suiteNotary(domain, notaryIndex));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Try to add a chain's given Notary to a given domain and expect a revert
    function globalNotaryRegistryAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(domain, notaryIndex);
        vm.expectRevert(revertMessage);
        globalNotaryRegistry.addNotary(domain, notary);
    }

    // Try to add a given Notary to a given domain and expect a revert
    function globalNotaryRegistryAddNotary(
        uint32 domain,
        address notary,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        globalNotaryRegistry.addNotary(domain, notary);
    }

    // Try to remove a chain's given Notary from a given domain and expect a revert
    function globalNotaryRegistryRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(domain, notaryIndex);
        vm.expectRevert(revertMessage);
        globalNotaryRegistry.removeNotary(domain, notary);
    }

    // Try to remove a given Notary from a given domain and expect a revert
    function globalNotaryRegistryRemoveNotary(
        uint32 domain,
        address notary,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        globalNotaryRegistry.removeNotary(domain, notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Add a chain's given Notary to a given domain and check the return value
    function globalNotaryRegistryAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            globalNotaryRegistry.addNotary(domain, suiteNotary(domain, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    // Add a given Notary to a given domain and check the return value
    function globalNotaryRegistryAddNotary(
        uint32 domain,
        address notary,
        bool returnValue
    ) public {
        assertEq(globalNotaryRegistry.addNotary(domain, notary), returnValue, "!returnValue");
    }

    // Remove a chain's given Notary to a given domain and check the return value
    function globalNotaryRegistryRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            globalNotaryRegistry.removeNotary(domain, suiteNotary(domain, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    // Remove a given Notary to a given domain and check the return value
    function globalNotaryRegistryRemoveNotary(
        uint32 domain,
        address notary,
        bool returnValue
    ) public {
        assertEq(globalNotaryRegistry.removeNotary(domain, notary), returnValue, "!returnValue");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TRIGGER VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function globalNotaryRegistryAllNotaries(uint32 domain) public view returns (address[] memory) {
        return globalNotaryRegistry.allNotaries(domain);
    }

    function globalNotaryRegistryGetNotary(uint32 domain, uint256 notaryIndex)
        public
        view
        returns (address)
    {
        return globalNotaryRegistry.getNotary(domain, notaryIndex);
    }

    function globalNotaryRegistryIsNotary(uint32 domain, address notary)
        public
        view
        returns (bool)
    {
        return globalNotaryRegistry.isNotary(domain, notary);
    }

    function globalNotaryRegistryIsNotary(uint32 domain, uint256 notaryIndex)
        public
        view
        returns (bool)
    {
        return globalNotaryRegistry.isNotary(domain, suiteNotary(domain, notaryIndex));
    }

    function globalNotaryRegistryNotariesAmount(uint32 domain) public view returns (uint256) {
        return globalNotaryRegistry.notariesAmount(domain);
    }
}
