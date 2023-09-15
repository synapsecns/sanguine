// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════
import {AgentFlag, AgentStatus} from "../contracts/libs/Structures.sol";
import {BondingManager} from "../contracts/manager/BondingManager.sol";
import {Inbox} from "../contracts/inbox/Inbox.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DeployMessaging003BaseScript} from "./DeployMessaging003Base.s.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {console, stdJson} from "forge-std/Script.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-console
// solhint-disable ordering
contract DeployMessaging003SynChainScript is DeployMessaging003BaseScript {
    using stdJson for string;
    using Strings for uint256;

    /// @dev Deploys BondingManager or LightManager
    /// Note: requires Origin, Destination, StatementInbox and Summit addresses to be set
    function _deployAgentManager() internal override returns (address deployment, bytes memory constructorArgs) {
        // new BondingManager(domain)
        constructorArgs = abi.encode(synapseDomain);
        deployment = factoryDeploy(agentManagerName(), type(BondingManager).creationCode, constructorArgs);
        require(origin != address(0), "Origin not set");
        require(destination != address(0), "Destination not set");
        require(statementInbox != address(0), "Statement Inbox not set");
        require(summit != address(0), "Summit not set");
    }

    /// @dev Initializes BondingManager or LightManager
    function _initializeAgentManager(address deployment) internal override {
        if (BondingManager(deployment).owner() == address(0)) {
            console.log("   %s: initializing", agentManagerName());
            BondingManager(deployment).initialize({
                origin_: origin,
                destination_: destination,
                inbox_: statementInbox,
                summit_: summit
            });
        } else {
            console.log("   %s: already initialized", agentManagerName());
        }
    }

    /// @dev Deploys Inbox or LightInbox
    /// Note: requires AgentManager, Origin, Destination and Summit addresses to be set
    function _deployStatementInbox() internal override returns (address deployment, bytes memory constructorArgs) {
        // new Inbox(domain)
        constructorArgs = abi.encode(synapseDomain);
        deployment = factoryDeploy(statementInboxName(), type(Inbox).creationCode, constructorArgs);
        require(agentManager != address(0), "Agent Manager not set");
        require(origin != address(0), "Origin not set");
        require(destination != address(0), "Destination not set");
        require(summit != address(0), "Summit not set");
    }

    /// @dev Initializes Inbox or LightInbox
    function _initializeStatementInbox(address deployment) internal override {
        if (Inbox(deployment).owner() == address(0)) {
            console.log("   %s: initializing", statementInboxName());
            Inbox(deployment).initialize({
                agentManager_: agentManager,
                origin_: origin,
                destination_: destination,
                summit_: summit
            });
        } else {
            console.log("   %s: already initialized", statementInboxName());
        }
    }

    /// @dev Adds agents to BondingManager (no-op for LightManager)
    function _addAgents() internal override {
        console.log("Adding Agents");
        uint256[] memory domains = globalConfig.readUintArray(".domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = globalConfig.readAddressArray(string.concat(".agents.", domain.toString()));
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                if (BondingManager(agentManager).agentStatus(agent).flag != AgentFlag.Unknown) {
                    console.log("   [address: %s] [domain: %s] skipped (already added)", agent, domain);
                    continue;
                }
                // Get a proof of non-inclusion
                bytes32[] memory proof = BondingManager(agentManager).getProof(agent);
                BondingManager(agentManager).addAgent(uint32(domain), agent, proof);
                console.log("   [address: %s] [domain: %s] added", agent, domain);
            }
        }
        string memory proofsKey = "proofs";
        string memory proofsJson = "";
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = globalConfig.readAddressArray(string.concat(".agents.", domain.toString()));
            string[] memory agentsStr = globalConfig.readStringArray(string.concat(".agents.", domain.toString()));
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                // Get a proof of inclusion
                bytes32[] memory proof = BondingManager(agentManager).getProof(agent);
                proofsJson = proofsKey.serialize(agentsStr[j], proof);
            }
        }
        // Save resulting agent root for deployments on other chains
        bytes32 agentRoot = BondingManager(agentManager).agentRoot();
        string memory agentRootConfig = "agentRoot";
        agentRootConfig.serialize("initialAgentRoot", agentRoot);
        agentRootConfig = agentRootConfig.serialize("proofs", proofsJson);
        string memory path = globalDeployConfigPath("Messaging003AgentRoot");
        agentRootConfig.write(path);
    }

    /// @dev Checks that all agents have been added correctly to BondingManager
    /// or that they could be added to LightManager.
    function _checkAgents() internal override {
        console.log("Checking Agents");
        uint256[] memory domains = globalConfig.readUintArray(".domains");
        // Agent indexes start from 1
        uint256 expectedIndex = 1;
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = globalConfig.readAddressArray(string.concat(".agents.", domain.toString()));
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                AgentStatus memory status = BondingManager(agentManager).agentStatus(agent);
                console.log("   [index: %s] [address: %s] [domain: %s]", status.index, agent, status.domain);
                require(status.flag == AgentFlag.Active, "Agent is not active");
                require(status.domain == uint32(domain), "Agent is not on correct domain");
                require(status.index == expectedIndex, "Agent has incorrect index");
                ++expectedIndex;
            }
        }
    }
}
