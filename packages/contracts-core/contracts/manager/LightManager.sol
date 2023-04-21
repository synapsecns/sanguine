// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AGENT_TREE_HEIGHT} from "../libs/Constants.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
import {AgentFlag, AgentStatus, SlashStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager} from "./AgentManager.sol";
import {DomainContext} from "../context/DomainContext.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
import {Versioned} from "../Version.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is Versioned, AgentManager, InterfaceLightManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════
    /// @inheritdoc IAgentManager
    bytes32 public agentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private _agentMap;

    // (index => agent)
    mapping(uint256 => address) private _agents;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) DomainContext(domain) Versioned("0.0.3") {
        require(!_onSynapseChain(), "Can't be deployed on SynChain");
    }

    function initialize(address origin_, address destination_) external initializer {
        __AgentManager_init(origin_, destination_);
        __Ownable_init();
    }

    // ═══════════════════════════════════════════════ AGENTS LOGIC ════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external {
        address storedAgent = _agents[status.index];
        require(storedAgent == address(0) || storedAgent == agent, "Invalid agent index");
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(status.flag, status.domain, agent);
        bytes32 root = agentRoot;
        // Check that proof matches the latest merkle root
        require(MerkleMath.proofRoot(status.index, leaf, proof, AGENT_TREE_HEIGHT) == root, "Invalid proof");
        // Save index => agent in the map
        if (storedAgent == address(0)) _agents[status.index] = agent;
        // Update the agent status against this root
        _agentMap[root][agent] = status;
        emit StatusUpdated(status.flag, status.domain, agent);
        // Notify local Registries, if agent flag is Slashed
        if (status.flag == AgentFlag.Slashed) {
            // Prover is msg.sender
            _notifySlashing(DESTINATION | ORIGIN, status.domain, agent, msg.sender);
        }
    }

    /// @inheritdoc InterfaceLightManager
    function setAgentRoot(bytes32 agentRoot_) external {
        require(msg.sender == destination, "Only Destination sets agent root");
        _setAgentRoot(agentRoot_);
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        require(msg.sender == destination, "!destination");
        // Only AgentManager on Synapse Chain can give instructions to withdraw tips
        require(msgOrigin == SYNAPSE_DOMAIN, "!synapseDomain");
        // Check that merkle proof is mature enough
        require(proofMaturity >= BONDING_OPTIMISTIC_PERIOD, "!optimisticPeriod");
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        // Magic value to return is selector of the called function
        return this.remoteWithdrawTips.selector;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _afterRegistrySlash(uint32 domain, address agent, address prover) internal override {
        // Send a manager message to BondingManager on SynChain
        // remoteRegistrySlash(msgOrigin, proofMaturity, domain, agent, prover) with the first two security args omitted
        InterfaceOrigin(origin).sendManagerMessage({
            destination: SYNAPSE_DOMAIN,
            optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
            payload: abi.encodeWithSelector(InterfaceBondingManager.remoteRegistrySlash.selector, domain, agent, prover)
        });
    }

    /// @dev Updates the Agent Merkle Root that Light Manager is tracking.
    function _setAgentRoot(bytes32 _agentRoot) internal {
        if (agentRoot != _agentRoot) {
            agentRoot = _agentRoot;
            emit RootUpdated(_agentRoot);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _agentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agentRoot][agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        return _agents[index];
    }
}
