// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";
import { AgentInfo, SystemEntity } from "../../../contracts/libs/Structures.sol";

import { AgentManagerTest } from "./AgentManager.t.sol";

import { LightManager, ISystemContract, SynapseTest } from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
contract LightManagerTest is AgentManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.expectRevert("Can't be deployed on SynChain");
        new LightManager(DOMAIN_SYNAPSE);
    }

    function test_version() public {
        // Check version
        assertEq(lightManager.version(), LATEST_VERSION, "!version");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: UNAUTHORIZED ACCESS (NOT OWNER)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.addAgent(1, address(1));
    }

    function test_removeAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.removeAgent(1, address(1));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address agent) public {
        (bool isActive, ) = lightManager.isActiveAgent(agent);
        // Should not be an already added agent
        vm.assume(!isActive);
        lightManager.addAgent(domain, agent);
        assertTrue(lightManager.isActiveAgent(domain, agent));
    }

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        lightManager.removeAgent(domain, agent);
        assertFalse(lightManager.isActiveAgent(domain, agent));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TEST: REGISTRY SLASH                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_registrySlash_origin(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        vm.expectCall(
            destination,
            abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent)
        );
        vm.prank(origin);
        lightManager.registrySlash(domain, agent);
        assertFalse(lightManager.isActiveAgent(domain, agent));
    }

    function test_registrySlash_revertUnauthorized(address caller) public {
        vm.assume(caller != origin);
        vm.expectRevert("Unauthorized caller");
        vm.prank(caller);
        lightManager.registrySlash(0, address(0));
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
