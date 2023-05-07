// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AGENT_TREE_HEIGHT, BONDING_OPTIMISTIC_PERIOD, SYNAPSE_DOMAIN} from "../libs/Constants.sol";
import {
    IncorrectAgentIndex,
    IncorrectAgentProof,
    CallerNotDestination,
    MustBeSynapseDomain,
    SynapseDomainForbidden,
    WithdrawTipsOptimisticPeriod
} from "../libs/Errors.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
import {AgentFlag, AgentStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager} from "./AgentManager.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is AgentManager, InterfaceLightManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════
    /// @inheritdoc IAgentManager
    bytes32 public agentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private _agentMap;

    // (index => agent)
    mapping(uint256 => address) private _agents;

    // (agent => index)
    mapping(address => uint256) private _agentIndexes;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) MessagingBase("0.0.3", domain) {
        if (domain == SYNAPSE_DOMAIN) revert SynapseDomainForbidden();
    }

    function initialize(address origin_, address destination_, address inbox_) external initializer {
        __AgentManager_init(origin_, destination_, inbox_);
        __Ownable_init();
    }

    // ═══════════════════════════════════════════════ AGENTS LOGIC ════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external {
        address storedAgent = _agents[status.index];
        if (storedAgent != address(0) && storedAgent != agent) revert IncorrectAgentIndex();
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(status.flag, status.domain, agent);
        bytes32 root = agentRoot;
        // Check that proof matches the latest merkle root
        if (MerkleMath.proofRoot(status.index, leaf, proof, AGENT_TREE_HEIGHT) != root) revert IncorrectAgentProof();
        // Save index => agent in the map
        if (storedAgent == address(0)) {
            _agents[status.index] = agent;
            _agentIndexes[agent] = status.index;
        }
        // Update the agent status against this root
        _agentMap[root][agent] = status;
        emit StatusUpdated(status.flag, status.domain, agent);
        // Notify local AgentSecured contracts, if agent flag is Slashed
        if (status.flag == AgentFlag.Slashed) {
            // This will revert if the agent has been slashed earlier
            _resolveDispute(status.index, msg.sender);
        }
    }

    /// @inheritdoc InterfaceLightManager
    function setAgentRoot(bytes32 agentRoot_) external {
        // Only destination can pass AgentRoot to be set
        if (msg.sender != destination) revert CallerNotDestination();
        _setAgentRoot(agentRoot_);
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        if (msg.sender != destination) revert CallerNotDestination();
        // Only AgentManager on Synapse Chain can give instructions to withdraw tips
        if (msgOrigin != SYNAPSE_DOMAIN) revert MustBeSynapseDomain();
        // Check that merkle proof is mature enough
        // TODO: separate constant for withdrawing tips optimistic period
        if (proofMaturity < BONDING_OPTIMISTIC_PERIOD) revert WithdrawTipsOptimisticPeriod();
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        // Magic value to return is selector of the called function
        return this.remoteWithdrawTips.selector;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _afterAgentSlashed(uint32 domain, address agent, address prover) internal virtual override {
        // Send a manager message to BondingManager on SynChain
        // remoteSlashAgent(msgOrigin, proofMaturity, domain, agent, prover) with the first two security args omitted
        InterfaceOrigin(origin).sendManagerMessage({
            destination: SYNAPSE_DOMAIN,
            optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
            payload: abi.encodeWithSelector(InterfaceBondingManager.remoteSlashAgent.selector, domain, agent, prover)
        });
    }

    /// @dev Notify local AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal override {
        // Origin contract doesn't need to know about the dispute
        IAgentSecured(destination).openDispute(guardIndex, notaryIndex);
    }

    /// @dev Notify local AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal override {
        // Origin contract doesn't need to know about the dispute
        IAgentSecured(destination).resolveDispute(slashedIndex, rivalIndex);
    }

    /// @dev Updates the Agent Merkle Root that Light Manager is tracking.
    function _setAgentRoot(bytes32 _agentRoot) internal {
        if (agentRoot != _agentRoot) {
            agentRoot = _agentRoot;
            emit RootUpdated(_agentRoot);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the stored status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _storedAgentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agentRoot][agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes, or for indexes
    /// of the agents that have not been added to Light Manager yet.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        return _agents[index];
    }

    /// @dev Returns the index of the agent in the Agent Merkle Tree. Returns zero for non existing agents, or
    /// for agents that have not been added to Light Manager yet.
    function _getIndex(address agent) internal view override returns (uint256 index) {
        return _agentIndexes[agent];
    }
}
