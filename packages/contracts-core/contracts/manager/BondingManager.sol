// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AgentFlag, AgentStatus, SlashStatus} from "../libs/Structures.sol";
import {DynamicTree, MerkleTree} from "../libs/MerkleTree.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager} from "./AgentManager.sol";
import {DomainContext} from "../context/DomainContext.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
import {Versioned} from "../Version.sol";

/// @notice BondingManager keeps track of all existing _agents.
/// Used on the Synapse Chain, serves as the "source of truth" for LightManagers on remote chains.
contract BondingManager is Versioned, AgentManager, InterfaceBondingManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (agent => their status)
    mapping(address => AgentStatus) private _agentMap;

    // A list of all agent accounts. First entry is address(0) to make agent indexes start from 1.
    address[] private _agents;

    // Merkle Tree for Agents.
    // leafs[0] = 0
    // leafs[index > 0] = keccak(agentFlag, domain, _agents[index])
    DynamicTree private _agentTree;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) DomainContext(domain) Versioned("0.0.3") {
        require(_onSynapseChain(), "Only deployed on SynChain");
    }

    function initialize(address origin_, address destination_) external initializer {
        __AgentManager_init(origin_, destination_);
        __Ownable_init();
        // Insert a zero address to make indexes for Agents start from 1.
        // Zeroed index is supposed to be used as a sentinel value meaning "no agent".
        _agents.push(address(0));
    }

    // ════════════════════════════════════════════ AGENTS LOGIC (MVP) ═════════════════════════════════════════════════

    // TODO: remove these MVP functions once token staking is implemented

    /// @inheritdoc InterfaceBondingManager
    function addAgent(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check current status of the added agent
        AgentStatus memory status = _agentStatus(agent);
        // Agent index in `_agents`
        uint32 index;
        // Leaf representing currently saved agent information in the tree
        bytes32 oldValue;
        if (status.flag == AgentFlag.Unknown) {
            // Unknown address could be added to any domain
            // New agent will need to be added to `_agents` list
            require(_agents.length < type(uint32).max, "Agents list if full");
            index = uint32(_agents.length);
            // Current leaf for index is bytes32(0), which is already assigned to `leaf`
            _agents.push(agent);
        } else if (status.flag == AgentFlag.Resting && status.domain == domain) {
            // Resting agent could be only added back to the same domain
            // Agent is already in `_agents`, fetch the saved index
            index = status.index;
            // Generate the current leaf for the agent
            // oldValue includes the domain information, so we didn't had to check it above.
            // However, we are still doing this check to have a more appropriate revert string,
            // if a resting agent is requesting to be added to another domain.
            oldValue = _agentLeaf(AgentFlag.Resting, domain, agent);
        } else {
            // Any other flag indicates that agent could not be added
            revert("Agent could not be added");
        }
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Active, domain, index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function initiateUnstaking(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check current status of the unstaking agent
        AgentStatus memory status = _agentStatus(agent);
        // Could only initiate the unstaking for the active agent for the domain
        require(status.flag == AgentFlag.Active && status.domain == domain, "Unstaking could not be initiated");
        // Leaf representing currently saved agent information in the tree.
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if an agent is initiating the unstaking, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(AgentFlag.Active, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Unstaking, domain, status.index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function completeUnstaking(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check current status of the unstaking agent
        AgentStatus memory status = _agentStatus(agent);
        // Could only complete the unstaking, if it was previously initiated
        // TODO: add more checks (time-based, possibly collecting info from other chains)
        require(status.flag == AgentFlag.Unstaking && status.domain == domain, "Unstaking could not be completed");
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if an agent is completing the unstaking, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(AgentFlag.Unstaking, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Resting, domain, status.index), agent);
    }

    // ══════════════════════════════════════════════ SLASHING LOGIC ═══════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function completeSlashing(uint32 domain, address agent, bytes32[] memory proof) external {
        // Check that slashing was initiated by one of the System Registries
        require(slashStatus[agent].isSlashed, "Slashing not initiated");
        // Check that agent is Active/Unstaking and that the domains match
        AgentStatus memory status = _agentStatus(agent);
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking) && status.domain == domain,
            "Slashing could not be completed"
        );
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if anyone is completing the slashing, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(status.flag, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Slashed, domain, status.index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function remoteRegistrySlash(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        require(msg.sender == destination, "!destination");
        // Check that merkle proof is mature enough
        require(proofMaturity >= BONDING_OPTIMISTIC_PERIOD, "!optimisticPeriod");
        // TODO: do we need to save this?
        msgOrigin;
        // Check that Agent hasn't been already slashed and initiate the slashing
        _initiateSlashing(domain, agent, prover);
        // Notify local registries about the slashing
        _notifySlashing(DESTINATION | ORIGIN, domain, agent, prover);
        // Magic value to return is selector of the called function
        return this.remoteRegistrySlash.selector;
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function withdrawTips(address recipient, uint32 origin_, uint256 amount) external {
        require(msg.sender == destination, "Only Summit withdraws tips");
        if (origin_ == localDomain) {
            // Call local Origin to withdraw tips
            InterfaceOrigin(address(origin)).withdrawTips(recipient, amount);
        } else {
            // For remote chains: send a manager message to remote LightManager to handle the withdrawal
            // remoteWithdrawTips(msgOrigin, proofMaturity, recipient, amount) with the first two security args omitted
            InterfaceOrigin(origin).sendManagerMessage({
                destination: origin_,
                optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
                payload: abi.encodeWithSelector(InterfaceLightManager.remoteWithdrawTips.selector, recipient, amount)
            });
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function agentRoot() external view override returns (bytes32) {
        return _agentTree.root;
    }

    /// @inheritdoc InterfaceBondingManager
    function agentLeaf(address agent) external view returns (bytes32 leaf) {
        return _getLeaf(agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function leafsAmount() external view returns (uint256 amount) {
        return _agents.length;
    }

    /// @inheritdoc InterfaceBondingManager
    function getProof(address agent) external view returns (bytes32[] memory proof) {
        bytes32[] memory leafs = allLeafs();
        AgentStatus memory status = _agentStatus(agent);
        // Use next available index for unknown agents
        uint256 index = status.flag == AgentFlag.Unknown ? _agents.length : status.index;
        return MerkleMath.calculateProof(leafs, index);
    }

    /// @inheritdoc InterfaceBondingManager
    function allLeafs() public view returns (bytes32[] memory leafs) {
        return getLeafs(0, _agents.length);
    }

    /// @inheritdoc InterfaceBondingManager
    function getLeafs(uint256 indexFrom, uint256 amount) public view returns (bytes32[] memory leafs) {
        uint256 amountTotal = _agents.length;
        require(indexFrom < amountTotal, "Out of range");
        if (indexFrom + amount > amountTotal) {
            amount = amountTotal - indexFrom;
        }
        leafs = new bytes32[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            leafs[i] = _getLeaf(indexFrom + i);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Updates value in the Agent Merkle Tree to reflect the `newStatus`.
    /// Will revert, if supplied proof for the old value is incorrect.
    function _updateLeaf(bytes32 oldValue, bytes32[] memory proof, AgentStatus memory newStatus, address agent)
        internal
    {
        // New leaf value for the agent in the Agent Merkle Tree
        bytes32 newValue = _agentLeaf(newStatus.flag, newStatus.domain, agent);
        // This will revert if the proof for the old value is incorrect
        bytes32 newRoot = _agentTree.update(newStatus.index, oldValue, proof, newValue);
        _agentMap[agent] = newStatus;
        emit StatusUpdated(newStatus.flag, newStatus.domain, agent);
        emit RootUpdated(newRoot);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the status of the agent.
    function _agentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        if (index < _agents.length) {
            agent = _agents[index];
        }
    }

    /// @dev Returns the current leaf representing agent in the Agent Merkle Tree.
    function _getLeaf(address agent) internal view returns (bytes32 leaf) {
        AgentStatus memory status = _agentStatus(agent);
        if (status.flag != AgentFlag.Unknown) {
            return _agentLeaf(status.flag, status.domain, agent);
        }
        // Return empty leaf for unknown _agents
    }

    /// @dev Returns a leaf from the Agent Merkle Tree with a given index.
    function _getLeaf(uint256 index) internal view returns (bytes32 leaf) {
        if (index != 0) {
            return _getLeaf(_agents[index]);
        }
        // Return empty leaf for a zero index
    }
}
