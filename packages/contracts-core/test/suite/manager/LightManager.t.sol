// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { SystemEntity } from "../../../contracts/libs/Structures.sol";
import { ISystemRegistry } from "../../../contracts/interfaces/ISystemRegistry.sol";

import { AgentManagerTest } from "./AgentManager.t.sol";

import {
    AgentFlag,
    AgentStatus,
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

    function test_setAgentRoot_revert_notDestination(address caller) public {
        vm.assume(caller != destination);
        vm.expectRevert("Only Destination sets agent root");
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
        // Should not be an already added agent
        vm.assume(lightManager.agentStatus(agent).flag == AgentFlag.Unknown);
        bytes32 root = addNewAgent(domain, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        vm.expectEmit();
        emit StatusUpdated(AgentFlag.Active, domain, agent);
        // Anyone could add agents in Light Manager
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Active);
    }

    function test_updateAgentStatus_slashed(
        address caller,
        uint256 domainId,
        uint256 agentId
    ) public {
        (uint32 domain, address agent) = getAgent(domainId, agentId);
        // Set flag to Slashed in the Merkle Tree
        bytes32 root = updateAgent(AgentFlag.Slashed, agent);
        test_setAgentRoot(root);
        bytes32[] memory proof = getAgentProof(agent);
        vm.expectEmit();
        emit StatusUpdated(AgentFlag.Slashed, domain, agent);
        bytes memory expectedCall = abi.encodeWithSelector(
            ISystemRegistry.managerSlash.selector,
            domain,
            agent
        );
        vm.expectCall(destination, expectedCall);
        vm.expectCall(origin, expectedCall);
        // Anyone could add agents in Light Manager
        vm.prank(caller);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
        checkAgentStatus(agent, lightManager.agentStatus(agent), AgentFlag.Slashed);
    }

    function test_setAgentRoot(bytes32 root) public {
        vm.expectEmit();
        emit RootUpdated(root);
        vm.prank(destination);
        lightManager.setAgentRoot(root);
        assertEq(lightManager.agentRoot(), root, "!agentRoot");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    TEST: UPDATE AGENTS (REVERTS)                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_addAgent_revert_invalidProof(uint256 domainId, uint256 agentId) public {
        (, address agent) = getAgent(domainId, agentId);
        bytes32[] memory proof = getAgentProof(agent);
        AgentStatus memory status = getAgentStatus(agent);
        // This succeeds, but doesn't do anything, as agent was already added
        lightManager.updateAgentStatus(agent, status, proof);
        // Change agent root, so old proofs are no longer valid
        test_setAgentRoot(bytes32(0));
        assertEq(uint8(lightManager.agentStatus(agent).flag), uint8(AgentFlag.Unknown));
        vm.expectRevert("Invalid proof");
        lightManager.updateAgentStatus(agent, status, proof);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         TEST: REGISTRY SLASH                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_registrySlash_origin(
        uint32 domain,
        address agent,
        address prover
    ) public {
        test_addAgent_new(address(this), domain, agent);
        bytes memory data = _remoteSlashData(domain, agent, prover);
        vm.expectEmit();
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        vm.expectCall(
            destination,
            abi.encodeWithSelector(ISystemRegistry.managerSlash.selector, domain, agent)
        );
        // (_destination, _optimisticSeconds, _recipient, _data)
        vm.expectCall(
            address(systemRouter),
            abi.encodeWithSelector(
                systemRouter.systemCall.selector,
                DOMAIN_SYNAPSE,
                BONDING_OPTIMISTIC_PERIOD,
                SystemEntity.AgentManager,
                data
            )
        );
        vm.prank(origin);
        lightManager.registrySlash(domain, agent, prover);
        assertEq(uint8(lightManager.agentStatus(agent).flag), uint8(AgentFlag.Fraudulent));
        (bool isSlashed, address _prover) = lightManager.slashStatus(agent);
        assertTrue(isSlashed);
        assertEq(_prover, prover);
    }

    function test_registrySlash_revertUnauthorized(address caller) public {
        vm.assume(caller != origin);
        vm.expectRevert("Unauthorized caller");
        vm.prank(caller);
        // Try to slash an existing agent
        lightManager.registrySlash(0, domains[0].agent, address(0));
    }

    function _localDomain() internal pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
