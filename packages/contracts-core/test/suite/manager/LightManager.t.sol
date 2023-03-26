// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";

import { AgentManagerTest } from "./AgentManager.t.sol";

import {
    LightManager,
    ISystemContract,
    ISystemRegistry,
    SynapseTest
} from "../../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
contract LightManagerTest is AgentManagerTest {
    // Deploy mocks for every messaging contract
    constructor() SynapseTest(0) {}

    function test_initializer(
        address caller,
        address _origin,
        address _destination
    ) public {
        lightManager = new LightManager(DOMAIN_LOCAL);
        vm.prank(caller);
        lightManager.initialize(ISystemRegistry(_origin), ISystemRegistry(_destination));
        assertEq(lightManager.owner(), caller);
        assertEq(address(lightManager.origin()), _origin);
        assertEq(address(lightManager.destination()), _destination);
    }

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

    // TODO: this should be only called by Destination
    function test_setAgentRoot_revert_notOwner(address caller) public {
        vm.assume(caller != address(this));
        expectRevertNotOwner();
        vm.prank(caller);
        lightManager.setAgentRoot(bytes32(uint256(1)));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TESTS: ADD/REMOVE AGENTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_new(
        address caller,
        uint32 domain,
        address agent
    ) public {
        (bool isActive, ) = lightManager.isActiveAgent(agent);
        // Should not be an already added agent
        vm.assume(!isActive);
        bytes32 root = addNewAgent(domain, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        // Anyone could add agents in Light Manager
        vm.prank(caller);
        lightManager.addAgent(domain, agent, proof, agentIndex[agent]);
        checkActive(lightManager, domain, agent);
    }

    function test_setAgentRoot(bytes32 root) public {
        lightManager.setAgentRoot(root);
        assertEq(lightManager.agentRoot(), root, "!agentRoot");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    TEST: UPDATE AGENTS (REVERTS)                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_invalidProof(uint256 domainId, uint256 agentId) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        bytes32[] memory proof = getAgentProof(agent);
        // This succeeds, but doesn't do anything, as agent was already added
        lightManager.addAgent(domain, agent, proof, agentIndex[agent]);
        // Change agent root, so old proofs are no longer valid
        test_setAgentRoot(bytes32(0));
        checkInactive(lightManager, domain, agent);
        vm.expectRevert("Invalid proof");
        lightManager.addAgent(domain, agent, proof, agentIndex[agent]);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TEST: REGISTRY SLASH                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_registrySlash_origin(uint32 domain, address agent) public {
        test_addAgent_new(address(this), domain, agent);
        vm.expectCall(
            destination,
            abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent)
        );
        vm.prank(origin);
        lightManager.registrySlash(domain, agent);
        // assertFalse(lightManager.isActiveAgent(domain, agent));
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
