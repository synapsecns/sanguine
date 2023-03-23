// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";
import { AgentManagerTest } from "./AgentManager.t.sol";

import { ISystemContract, Summit, SynapseTest } from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
contract BondingManagerTest is AgentManagerTest {
    // Deploy Production version of Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_SUMMIT) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: UNAUTHORIZED ACCESS (NOT OWNER)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.addAgent(1, address(1));
    }

    function test_removeAgent_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        bondingManager.removeAgent(1, address(1));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent(uint32 domain, address agent) public {
        (bool isActive, ) = bondingManager.isActiveAgent(agent);
        // Should not be an already added agent
        vm.assume(!isActive);
        bondingManager.addAgent(domain, agent);
        assertTrue(bondingManager.isActiveAgent(domain, agent));
    }

    function test_removeAgent(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        bondingManager.removeAgent(domain, agent);
        assertFalse(bondingManager.isActiveAgent(domain, agent));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TEST: REGISTRY SLASH                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_registrySlash_origin(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        vm.expectCall(
            summit,
            abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent)
        );
        vm.prank(originSynapse);
        bondingManager.registrySlash(domain, agent);
        assertFalse(bondingManager.isActiveAgent(domain, agent));
    }

    function test_registrySlash_summit(uint32 domain, address agent) public {
        test_addAgent(domain, agent);
        vm.expectCall(
            originSynapse,
            abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent)
        );
        vm.prank(summit);
        bondingManager.registrySlash(domain, agent);
        assertFalse(bondingManager.isActiveAgent(domain, agent));
    }

    function test_registrySlash_revertUnauthorized(address caller) public {
        vm.assume(caller != originSynapse && caller != summit);
        vm.expectRevert("Unauthorized caller");
        vm.prank(caller);
        bondingManager.registrySlash(0, address(0));
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
