// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════
import {AgentFlag, AgentStatus} from "../contracts/libs/Structures.sol";
import {LightManager} from "../contracts/manager/LightManager.sol";
import {LightInbox} from "../contracts/inbox/LightInbox.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DeployMessaging003BaseScript} from "./DeployMessaging003Base.s.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {console, stdJson} from "forge-std/Script.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-console
// solhint-disable ordering
contract DeployMessaging003LightChainScript is DeployMessaging003BaseScript {
    using stdJson for string;
    using Strings for uint256;

    /// @dev Deploys BondingManager or LightManager
    /// Note: requires Origin, Destination and StatementInbox addresses to be set
    function _deployAgentManager() internal override returns (address deployment, bytes memory constructorArgs) {
        // new LightManager(domain)
        constructorArgs = abi.encode(synapseDomain);
        deployment = factoryDeploy(agentManagerName(), type(LightManager).creationCode, constructorArgs);
        require(origin != address(0), "Origin not set");
        require(destination != address(0), "Destination not set");
        require(statementInbox != address(0), "Statement Inbox not set");
    }

    /// @dev Initializes BondingManager or LightManager
    function _initializeAgentManager(address deployment) internal override {
        if (LightManager(deployment).owner() == address(0)) {
            console.log("   %s: initializing", agentManagerName());
            LightManager(deployment).initialize({origin_: origin, destination_: destination, inbox_: statementInbox});
        } else {
            console.log("   %s: already initialized", agentManagerName());
        }
    }

    /// @dev Deploys Inbox or LightInbox
    /// Note: requires AgentManager, Origin and Destination addresses to be set
    function _deployStatementInbox() internal override returns (address deployment, bytes memory constructorArgs) {
        // new LightInbox(synapseDomain)
        constructorArgs = abi.encode(synapseDomain);
        deployment = factoryDeploy(statementInboxName(), type(LightInbox).creationCode, constructorArgs);
        require(agentManager != address(0), "Agent Manager not set");
        require(origin != address(0), "Origin not set");
        require(destination != address(0), "Destination not set");
    }

    /// @dev Initializes Inbox or LightInbox
    function _initializeStatementInbox(address deployment) internal override {
        if (LightInbox(deployment).owner() == address(0)) {
            console.log("   %s: initializing", statementInboxName());
            LightInbox(deployment).initialize({agentManager_: agentManager, origin_: origin, destination_: destination});
        } else {
            console.log("   %s: already initialized", statementInboxName());
        }
    }

    /// @dev Adds agents to BondingManager (no-op for LightManager)
    function _addAgents() internal view override {
        console.log("Adding Agents: skipping for LightManager");
    }

    /// @dev Checks that all agents have been added correctly to BondingManager
    /// or that they could be added to LightManager.
    function _checkAgents() internal override {
        console.log("Adding Agents (simulation)");
        string memory agentRootConfig = loadGlobalDeployConfig("Messaging003AgentRoot");
        uint256[] memory domains = globalConfig.readUintArray(".domains");
        // Agent indexes start from 1
        uint256 expectedIndex = 1;
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = globalConfig.readAddressArray(string.concat(".agents.", domain.toString()));
            string[] memory agentsStr = globalConfig.readStringArray(string.concat(".agents.", domain.toString()));
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                bytes32[] memory proof = agentRootConfig.readBytes32Array(string.concat(".proofs.", agentsStr[j]));
                AgentStatus memory status =
                    AgentStatus({flag: AgentFlag.Active, domain: uint32(domain), index: uint32(expectedIndex)});
                console.log("   [index: %s] [address: %s] [domain: %s]", status.index, agent, status.domain);
                LightManager(agentManager).updateAgentStatus(agent, status, proof);
                ++expectedIndex;
            }
        }
    }
}
