// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { DestinationHarness } from "../harnesses/DestinationHarness.t.sol";
import { OriginHarness } from "../harnesses/OriginHarness.t.sol";
import { SummitHarness } from "../harnesses/SummitHarness.t.sol";

import { Test } from "forge-std/Test.sol";

interface Initializable {
    function initialize() external;
}

// disable-solhint func-name-mixedcase
contract HarnessTest is Test {
    function test_destinationHarness(
        uint32 domain,
        address notary,
        address guard
    ) public {
        vm.assume(notary != guard);
        domain = uint32(bound(domain, 1, type(uint32).max));
        DestinationHarness destination = new DestinationHarness(domain);
        Initializable(address(destination)).initialize();
        // LOCAL Notary
        destination.addAgent(domain, notary);
        assertTrue(destination.isActiveAgent(domain, notary));
        destination.removeAgent(domain, notary);
        assertFalse(destination.isActiveAgent(domain, notary));
        // Guard
        destination.addAgent(0, guard);
        assertTrue(destination.isActiveAgent(0, guard));
        destination.removeAgent(0, guard);
        assertFalse(destination.isActiveAgent(0, guard));
    }

    function test_originHarness(
        uint32 domain,
        address notary,
        address guard
    ) public {
        vm.assume(notary != guard);
        domain = uint32(bound(domain, 1, type(uint32).max));
        OriginHarness origin = new OriginHarness(1);
        Initializable(address(origin)).initialize();
        // Any Notary
        origin.addAgent(domain, notary);
        assertTrue(origin.isActiveAgent(domain, notary));
        origin.removeAgent(domain, notary);
        assertFalse(origin.isActiveAgent(domain, notary));
        // Guard
        origin.addAgent(0, guard);
        assertTrue(origin.isActiveAgent(0, guard));
        origin.removeAgent(0, guard);
        assertFalse(origin.isActiveAgent(0, guard));
    }

    function test_summitHarness(
        uint32 domain,
        address notary,
        address guard
    ) public {
        vm.assume(notary != guard);
        domain = uint32(bound(domain, 1, type(uint32).max));
        SummitHarness summit = new SummitHarness();
        Initializable(address(summit)).initialize();
        // Any Notary
        summit.addAgent(domain, notary);
        assertTrue(summit.isActiveAgent(domain, notary));
        summit.removeAgent(domain, notary);
        assertFalse(summit.isActiveAgent(domain, notary));
        // Guard
        summit.addAgent(0, guard);
        assertTrue(summit.isActiveAgent(0, guard));
        summit.removeAgent(0, guard);
        assertFalse(summit.isActiveAgent(0, guard));
    }
}
