// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AGENT_TREE_HEIGHT } from "../libs/Constants.sol";
import { MerkleLib } from "../libs/Merkle.sol";
import { AgentFlag, AgentStatus } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentManager, IAgentManager, ISystemRegistry } from "./AgentManager.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { ILightManager } from "../interfaces/ILightManager.sol";
import { Versioned } from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager, ILightManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Latest known Agent Merkle Root
    bytes32 private latestAgentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) public agentStatus;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      CONSTRUCTOR & INITIALIZER                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(uint32 _domain) DomainContext(_domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize(ISystemRegistry _origin, ISystemRegistry _destination)
        external
        initializer
    {
        __AgentManager_init(_origin, _destination);
        __Ownable_init();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             AGENTS LOGIC                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc ILightManager
    function addAgent(
        uint32 _domain,
        address _agent,
        bytes32[] memory _proof,
        uint256 _index
    ) external {
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(AgentFlag.Active, _domain, _agent);
        bytes32 root = latestAgentRoot;
        // Check that proof matches the latest merkle root
        require(
            MerkleLib.proofRoot(_index, leaf, _proof, AGENT_TREE_HEIGHT) == root,
            "Invalid proof"
        );
        // Mark agent as registered against this root
        agentStatus[root][_agent] = AgentStatus(AgentFlag.Active, _domain, uint32(_index));
    }

    /// @inheritdoc ILightManager
    function setAgentRoot(bytes32 _agentRoot) external onlyOwner {
        // TODO: only destination should be able to call this
        if (latestAgentRoot != _agentRoot) {
            latestAgentRoot = _agentRoot;
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SLASHING LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function registrySlash(uint32 _domain, address _agent) external {
        // On chains other than Synapse Chain only Origin could slash Agents
        // TODO: add "marked for external slashing" logic
        if (msg.sender == address(origin)) {
            destination.managerSlash(_domain, _agent);
            // TODO: issue a system call to BondingManager on SynChain
        } else {
            revert("Unauthorized caller");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc IAgentManager
    function agentRoot() public view override returns (bytes32) {
        return latestAgentRoot;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _agentStatus(address _agent) internal view override returns (AgentStatus memory) {
        return agentStatus[latestAgentRoot][_agent];
    }
}
