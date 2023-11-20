// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {
    AGENT_ROOT_PROPOSAL_TIMEOUT,
    AGENT_TREE_HEIGHT,
    BONDING_OPTIMISTIC_PERIOD,
    FRESH_DATA_TIMEOUT
} from "../libs/Constants.sol";
import {
    AgentRootNotProposed,
    AgentRootTimeoutNotOver,
    IncorrectAgentIndex,
    IncorrectAgentProof,
    IncorrectAgentRoot,
    CallerNotDestination,
    MustBeSynapseDomain,
    NotStuck,
    SynapseDomainForbidden,
    WithdrawTipsOptimisticPeriod
} from "../libs/Errors.sol";
import {MerkleMath} from "../libs/merkle/MerkleMath.sol";
import {AgentFlag, AgentStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager} from "./AgentManager.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

/// @notice LightManager keeps track of all agents on chains other than Synapse Chain.
/// Is uses the Agent Merkle Roots from the Notary-signed attestations to stay in sync with the `BondingManager`.
/// `LightManager` is responsible for the following:
/// - Accepting the Agent Merkle Roots (passing the optimistic period check) from the `Destination` contract.
/// - Using these roots to enable agents to register themselves by proving their status.
/// - Accepting Manager Message from `BondingManager` on Synapse Chain to withdraw tips.
/// - Sending Manager Messages to `BondingManager` on Synapse Chain to slash agents, when their fraud is proven.
contract LightManager is AgentManager, InterfaceLightManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════
    /// @inheritdoc IAgentManager
    bytes32 public agentRoot;

    /// @dev Pending Agent Merkle Root that was proposed by the contract owner.
    bytes32 internal _proposedAgentRoot;

    /// @dev Timestamp when the Agent Merkle Root was proposed by the contract owner.
    uint256 internal _agentRootProposedAt;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private _agentMap;

    // (index => agent)
    mapping(uint256 => address) private _agents;

    // (agent => index)
    mapping(address => uint256) private _agentIndexes;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_) MessagingBase("0.0.3", synapseDomain_) {
        if (localDomain == synapseDomain_) revert SynapseDomainForbidden();
    }

    function initialize(address origin_, address destination_, address inbox_, address owner_) external initializer {
        __AgentManager_init(origin_, destination_, inbox_, owner_);
    }

    // ════════════════════════════════════════════════ OWNER ONLY ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function proposeAgentRootWhenStuck(bytes32 agentRoot_) external onlyOwner onlyWhenStuck {
        if (agentRoot_ == 0) revert IncorrectAgentRoot();
        // Update the proposed agent root, clear the timer if the root is empty
        _proposedAgentRoot = agentRoot_;
        _agentRootProposedAt = block.timestamp;
        emit AgentRootProposed(agentRoot_);
    }

    /// @inheritdoc InterfaceLightManager
    function cancelProposedAgentRoot() external onlyOwner {
        bytes32 cancelledAgentRoot = _proposedAgentRoot;
        if (cancelledAgentRoot == 0) revert AgentRootNotProposed();
        _proposedAgentRoot = 0;
        _agentRootProposedAt = 0;
        emit ProposedAgentRootCancelled(cancelledAgentRoot);
    }

    /// @inheritdoc InterfaceLightManager
    /// @dev Should proceed with the proposed root, even if new Notary data is available.
    /// This is done to prevent rogue Notaries from going offline and then
    /// indefinitely blocking the agent root resolution, thus `onlyWhenStuck` modifier is not used here.
    function resolveProposedAgentRoot() external onlyOwner {
        bytes32 newAgentRoot = _proposedAgentRoot;
        if (newAgentRoot == 0) revert AgentRootNotProposed();
        if (block.timestamp < _agentRootProposedAt + AGENT_ROOT_PROPOSAL_TIMEOUT) revert AgentRootTimeoutNotOver();
        _setAgentRoot(newAgentRoot);
        _proposedAgentRoot = 0;
        _agentRootProposedAt = 0;
        emit ProposedAgentRootResolved(newAgentRoot);
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
        if (msgOrigin != synapseDomain) revert MustBeSynapseDomain();
        // Check that merkle proof is mature enough
        // TODO: separate constant for withdrawing tips optimistic period
        if (proofMaturity < BONDING_OPTIMISTIC_PERIOD) revert WithdrawTipsOptimisticPeriod();
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        // Magic value to return is selector of the called function
        return this.remoteWithdrawTips.selector;
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function proposedAgentRootData() external view returns (bytes32 agentRoot_, uint256 proposedAt_) {
        return (_proposedAgentRoot, _agentRootProposedAt);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _afterAgentSlashed(uint32 domain, address agent, address prover) internal virtual override {
        // Send a manager message to BondingManager on SynChain
        // remoteSlashAgent(msgOrigin, proofMaturity, domain, agent, prover) with the first two security args omitted
        InterfaceOrigin(origin).sendManagerMessage({
            destination: synapseDomain,
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
