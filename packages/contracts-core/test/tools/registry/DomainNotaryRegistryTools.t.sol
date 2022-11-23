// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../utils/SynapseTestSuite.t.sol";
import {
    DomainNotaryRegistryHarness
} from "../../harnesses/registry/DomainNotaryRegistryHarness.t.sol";

abstract contract DomainNotaryRegistryTools is SynapseTestSuite {
    DomainNotaryRegistryHarness internal domainNotaryRegistry;

    /**
     * @dev For tests we will be using Notaries from DOMAIN_LOCAL.
     * We want to test both "domain-specific" functions and those who
     * let you specify the domain. The latter should reject domains other than DOMAIN_LOCAL.
     * For these reasons you will see code like:
     *      domainNotaryRegistry.addNotary(domain, suiteNotary(DOMAIN_LOCAL, notaryIndex));
     * Which will use arbitrary domain, but will try to add a specific Notary.
     */

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectNotaryAdded(uint256 notaryIndex) public {
        expectNotaryAdded(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL, notaryIndex));
    }

    function expectNotaryAdded(uint32 domain, address notary) public {
        vm.expectEmit(true, true, true, true);
        emit BeforeNotaryAdded(domain, notary);
        vm.expectEmit(true, true, true, true);
        emit NotaryAdded(domain, notary);
        vm.expectEmit(true, true, true, true);
        emit AfterNotaryAdded(domain, notary);
    }

    function expectNotaryRemoved(uint256 notaryIndex) public {
        expectNotaryRemoved(DOMAIN_LOCAL, suiteNotary(DOMAIN_LOCAL, notaryIndex));
    }

    function expectNotaryRemoved(uint32 domain, address notary) public {
        vm.expectEmit(true, true, true, true);
        emit BeforeNotaryRemoved(domain, notary);
        vm.expectEmit(true, true, true, true);
        emit NotaryRemoved(domain, notary);
        vm.expectEmit(true, true, true, true);
        emit AfterNotaryRemoved(domain, notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Try to add given Notary and expect a revert
    function domainNotaryRegistryAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(DOMAIN_LOCAL, notaryIndex);
        vm.expectRevert(revertMessage);
        domainNotaryRegistry.addNotary(domain, notary);
    }

    // Try to remove given Notary and expect a revert
    function domainNotaryRegistryRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(DOMAIN_LOCAL, notaryIndex);
        vm.expectRevert(revertMessage);
        domainNotaryRegistry.removeNotary(domain, notary);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Add a given notary via addNotary(notary) and check the return value
    function domainNotaryRegistryAddNotary(uint256 notaryIndex, bool returnValue) public {
        assertEq(
            domainNotaryRegistry.addNotary(suiteNotary(DOMAIN_LOCAL, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    // Add a given notary via addNotary(domain, notary) and check the return value
    function domainNotaryRegistryAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            domainNotaryRegistry.addNotary(domain, suiteNotary(DOMAIN_LOCAL, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    // Remove a given notary via removeNotary(notary) and check the return value
    function domainNotaryRegistryRemoveNotary(uint256 notaryIndex, bool returnValue) public {
        assertEq(
            domainNotaryRegistry.removeNotary(suiteNotary(DOMAIN_LOCAL, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    // Remove a given notary via removeNotary(domain, notary) and check the return value
    function domainNotaryRegistryRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            domainNotaryRegistry.removeNotary(domain, suiteNotary(DOMAIN_LOCAL, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TRIGGER VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function domainNotaryRegistryAllNotaries() public view returns (address[] memory) {
        return domainNotaryRegistry.allNotaries();
    }

    function domainNotaryRegistryGetNotary(uint256 notaryIndex) public view returns (address) {
        return domainNotaryRegistry.getNotary(notaryIndex);
    }

    // Trigger isNotary(notary)
    function domainNotaryRegistryIsNotary(uint256 notaryIndex) public view returns (bool) {
        return domainNotaryRegistry.isNotary(suiteNotary(DOMAIN_LOCAL, notaryIndex));
    }

    // Trigger isNotary(domain, notary)
    function domainNotaryRegistryIsNotary(uint32 domain, uint256 notaryIndex)
        public
        view
        returns (bool)
    {
        return domainNotaryRegistry.isNotary(domain, suiteNotary(DOMAIN_LOCAL, notaryIndex));
    }

    function domainNotaryRegistryNotariesAmount() public view returns (uint256) {
        return domainNotaryRegistry.notariesAmount();
    }
}
