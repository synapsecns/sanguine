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
    function _deployAgentManager() internal override returns (address) {
        // new LightManager(domain)
        bytes memory constructorArgs = abi.encode(localDomain);
        return factoryDeploy(agentManagerName(), type(LightManager).creationCode, constructorArgs);
    }

    /// @dev Initializes BondingManager or LightManager
    function _initializeAgentManager() internal override {
        LightManager(agentManager).initialize({origin_: origin, destination_: destination, inbox_: statementInbox});
    }

    /// @dev Deploys Inbox or LightInbox
    function _deployStatementInbox() internal override returns (address) {
        // new LightInbox(domain)
        bytes memory constructorArgs = abi.encode(localDomain);
        return factoryDeploy(statementInboxName(), type(LightInbox).creationCode, constructorArgs);
    }

    /// @dev Initializes Inbox or LightInbox
    function _initializeStatementInbox() internal override {
        LightInbox(statementInbox).initialize({agentManager_: agentManager, origin_: origin, destination_: destination});
    }

    /// @dev Adds agents to BondingManager (no-op for LightManager)
    function _addAgents() internal view override {
        console.log("Adding Agents: skipping for LightManager");
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
                ++expectedIndex;
            }
        }
    }
}
