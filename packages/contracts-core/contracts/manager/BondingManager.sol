// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {BONDING_OPTIMISTIC_PERIOD} from "../libs/Constants.sol";
import {
    AgentCantBeAdded,
    CallerNotDestination,
    CallerNotSummit,
    DisputeAlreadyResolved,
    DisputeNotOpened,
    IncorrectAgentDomain,
    IncorrectOriginDomain,
    IndexOutOfRange,
    MustBeSynapseDomain,
    SlashAgentOptimisticPeriod,
    SynapseDomainForbidden
} from "../libs/Errors.sol";
import {DynamicTree, MerkleMath} from "../libs/merkle/MerkleTree.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager} from "./AgentManager.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

/// @notice BondingManager keeps track of all existing agents on the Synapse Chain.
/// It utilizes a dynamic Merkle Tree to store the agent information. This enables passing only the
/// latest merkle root of this tree (referenced as the Agent Merkle Root) to the remote chains,
/// so that the agents could "register" themselves by proving their current status against this root.
/// `BondingManager` is responsible for the following:
/// - Keeping track of all existing agents, as well as their statuses. In the MVP version there is no token staking,
///   which will be added in the future. Nonetheless, the agent statuses are still stored in the Merkle Tree, and
///   the agent slashing is still possible, though with no reward/penalty for the reporter/reported.
/// - Marking agents as "ready to be slashed" once their fraud is proven on the local or remote chain. Anyone could
///   complete the slashing by providing the proof of the current agent status against the current Agent Merkle Root.
/// - Sending Manager Message to remote `LightManager` to withdraw collected tips from the remote chain.
/// - Accepting Manager Message from remote `LightManager` to slash agents on the Synapse Chain, when their fraud
///   is proven on the remote chain.
contract BondingManager is AgentManager, InterfaceBondingManager {
    using SafeCast for uint256;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // The address of the Summit contract.
    address public summit;

    // (agent => their status)
    mapping(address => AgentStatus) private _agentMap;

    // (domain => past and current agents for domain)
    mapping(uint32 => address[]) private _domainAgents;

    // A list of all agent accounts. First entry is address(0) to make agent indexes start from 1.
    address[] private _agents;

    // Merkle Tree for Agents.
    // leafs[0] = 0
    // leafs[index > 0] = keccak(agentFlag, domain, _agents[index])
    DynamicTree private _agentTree;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_) MessagingBase("0.0.3", synapseDomain_) {
        if (localDomain != synapseDomain) revert MustBeSynapseDomain();
    }

    function initialize(address origin_, address destination_, address inbox_, address summit_, address owner_)
        external
        initializer
    {
        __AgentManager_init(origin_, destination_, inbox_, owner_);
        summit = summit_;
        // Insert a zero address to make indexes for Agents start from 1.
        // Zeroed index is supposed to be used as a sentinel value meaning "no agent".
        _agents.push(address(0));
    }

    // ════════════════════════════════════════════ AGENTS LOGIC (MVP) ═════════════════════════════════════════════════

    // TODO: remove these MVP functions once token staking is implemented

    /// @inheritdoc InterfaceBondingManager
    function addAgent(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        if (domain == synapseDomain) revert SynapseDomainForbidden();
        // Check the STORED status of the added agent in the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
        // Agent index in `_agents`
        uint32 index;
        // Leaf representing currently saved agent information in the tree
        bytes32 oldValue;
        if (status.flag == AgentFlag.Unknown) {
            // Unknown address could be added to any domain
            // New agent will need to be added to `_agents` list: could not have more than 2**32 agents
            // TODO: consider using more than 32 bits for agent indexes
            index = _agents.length.toUint32();
            // Current leaf for index is bytes32(0), which is already assigned to `leaf`
            _agents.push(agent);
            _domainAgents[domain].push(agent);
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
            revert AgentCantBeAdded();
        }
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Active, domain, index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function initiateUnstaking(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check the CURRENT status of the unstaking agent
        AgentStatus memory status = agentStatus(agent);
        // Could only initiate the unstaking for the active agent for the domain
        status.verifyActive();
        if (status.domain != domain) revert IncorrectAgentDomain();
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
        // Check the CURRENT status of the unstaking agent
        AgentStatus memory status = agentStatus(agent);
        // Could only complete the unstaking, if it was previously initiated
        // TODO: add more checks (time-based, possibly collecting info from other chains)
        status.verifyUnstaking();
        if (status.domain != domain) revert IncorrectAgentDomain();
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if an agent is completing the unstaking, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(AgentFlag.Unstaking, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Resting, domain, status.index), agent);
    }

    // ════════════════════════════════════════════════ ONLY OWNER ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function resolveDisputeWhenStuck(uint32 domain, address slashedAgent) external onlyOwner onlyWhenStuck {
        AgentDispute memory slashedDispute = _agentDispute[_getIndex(slashedAgent)];
        if (slashedDispute.flag == DisputeFlag.None) revert DisputeNotOpened();
        if (slashedDispute.flag == DisputeFlag.Slashed) revert DisputeAlreadyResolved();
        // This will revert if domain doesn't match the agent's domain.
        _slashAgent({domain: domain, agent: slashedAgent, prover: address(0)});
    }

    // ══════════════════════════════════════════════ SLASHING LOGIC ═══════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function completeSlashing(uint32 domain, address agent, bytes32[] memory proof) external {
        // Check the CURRENT status of the unstaking agent
        AgentStatus memory status = agentStatus(agent);
        // Could only complete the slashing, if it was previously initiated
        status.verifyFraudulent();
        if (status.domain != domain) revert IncorrectAgentDomain();
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if anyone is completing the slashing, but specifies incorrect domain.
        bytes32 oldValue = _getLeaf(agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Slashed, domain, status.index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        if (msg.sender != destination) revert CallerNotDestination();
        // Check that merkle proof is mature enough
        // TODO: separate constant for slashing optimistic period
        if (proofMaturity < BONDING_OPTIMISTIC_PERIOD) revert SlashAgentOptimisticPeriod();
        // TODO: do we need to save domain where the agent was slashed?
        // Message needs to be sent from the remote chain
        if (msgOrigin == localDomain) revert IncorrectOriginDomain();
        // Slash agent and notify local AgentSecured contracts
        _slashAgent(domain, agent, prover);
        // Magic value to return is selector of the called function
        return this.remoteSlashAgent.selector;
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function withdrawTips(address recipient, uint32 origin_, uint256 amount) external {
        // Only Summit can withdraw tips
        if (msg.sender != summit) revert CallerNotSummit();
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
    function getActiveAgents(uint32 domain) external view returns (address[] memory agents) {
        uint256 amount = _domainAgents[domain].length;
        agents = new address[](amount);
        uint256 activeAgents = 0;
        for (uint256 i = 0; i < amount; ++i) {
            address agent = _domainAgents[domain][i];
            if (agentStatus(agent).flag == AgentFlag.Active) {
                agents[activeAgents++] = agent;
            }
        }
        if (activeAgents != amount) {
            // Shrink the returned array by storing the required length in memory
            // solhint-disable-next-line no-inline-assembly
            assembly {
                mstore(agents, activeAgents)
            }
        }
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
        // Use the STORED agent status from the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
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
        if (indexFrom >= amountTotal) revert IndexOutOfRange();
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

    /// @dev Notify local AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal override {
        IAgentSecured(destination).openDispute(guardIndex, notaryIndex);
        IAgentSecured(summit).openDispute(guardIndex, notaryIndex);
    }

    /// @dev Notify local AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal override {
        IAgentSecured(destination).resolveDispute(slashedIndex, rivalIndex);
        IAgentSecured(summit).resolveDispute(slashedIndex, rivalIndex);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the status of the agent.
    function _storedAgentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        if (index < _agents.length) {
            agent = _agents[index];
        }
    }

    /// @dev Returns the index of the agent in the Agent Merkle Tree. Returns zero for non existing agents.
    function _getIndex(address agent) internal view override returns (uint256 index) {
        return _agentMap[agent].index;
    }

    /// @dev Returns the current leaf representing agent in the Agent Merkle Tree.
    function _getLeaf(address agent) internal view returns (bytes32 leaf) {
        // Get the agent status STORED in the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
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
