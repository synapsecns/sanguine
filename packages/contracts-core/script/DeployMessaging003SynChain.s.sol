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
    function _deployAgentManager() internal override returns (address) {
        // new BondingManager(domain)
        bytes memory constructorArgs = abi.encode(localDomain);
        return factoryDeploy(agentManagerName(), type(BondingManager).creationCode, constructorArgs);
    }

    /// @dev Initializes BondingManager or LightManager
    function _initializeAgentManager() internal override {
        BondingManager(agentManager).initialize({
            origin_: origin,
            destination_: destination,
            inbox_: statementInbox,
            summit_: summit
        });
    }

    /// @dev Deploys Inbox or LightInbox
    function _deployStatementInbox() internal override returns (address) {
        // new Inbox(domain)
        bytes memory constructorArgs = abi.encode(localDomain);
        return factoryDeploy(statementInboxName(), type(Inbox).creationCode, constructorArgs);
    }

    /// @dev Initializes Inbox or LightInbox
    function _initializeStatementInbox() internal override {
        Inbox(statementInbox).initialize({
            agentManager_: agentManager,
            origin_: origin,
            destination_: destination,
            summit_: summit
        });
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
                // Get a proof of non-inclusion
                bytes32[] memory proof = BondingManager(agentManager).getProof(agent);
                BondingManager(agentManager).addAgent(uint32(domain), agent, proof);
                console.log("   %s on domain [%s]", agent, domain);
            }
        }
        // Save resulting agent root for deployments on other chains
        bytes32 agentRoot = BondingManager(agentManager).agentRoot();
        string memory agentRootConfig = "agentRoot";
        agentRootConfig = agentRootConfig.serialize("initialAgentRoot", agentRoot);
        agentRootConfig.write(globalDeployConfigPath("Messaging003AgentRoot"));
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
