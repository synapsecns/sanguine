// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {
    CallerNotInbox,
    IncorrectAgentDomain,
    DisputeAlreadyResolved,
    GuardInDispute,
    NotaryInDispute
} from "../libs/Errors.sol";
import {AgentFlag, AgentStatus, Dispute, DisputeFlag} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "../base/MessagingBase.sol";
import {AgentManagerEvents} from "../events/AgentManagerEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";

abstract contract AgentManager is MessagingBase, AgentManagerEvents, IAgentManager {
    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    address public origin;

    address public destination;

    address public inbox;

    // (agent index => their dispute status)
    mapping(uint256 => Dispute) internal _disputes;

    /// @dev gap for upgrade safety
    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyInbox() {
        if (msg.sender != inbox) revert CallerNotInbox();
        _;
    }

    // ════════════════════════════════════════════════ INITIALIZER ════════════════════════════════════════════════════

    // solhint-disable-next-line func-name-mixedcase
    function __AgentManager_init(address origin_, address destination_, address inbox_) internal onlyInitializing {
        origin = origin_;
        destination = destination_;
        inbox = inbox_;
    }

    // ════════════════════════════════════════════════ ONLY INBOX ═════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    // solhint-disable-next-line ordering
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external onlyInbox {
        // Check that both agents are not in Dispute yet
        if (_disputes[guardIndex].flag != DisputeFlag.None) revert GuardInDispute();
        if (_disputes[notaryIndex].flag != DisputeFlag.None) revert NotaryInDispute();
        _updateDispute(guardIndex, Dispute(DisputeFlag.Pending, notaryIndex, address(0)));
        _updateDispute(notaryIndex, Dispute(DisputeFlag.Pending, guardIndex, address(0)));
        _notifyDisputeOpened(guardIndex, notaryIndex);
    }

    /// @inheritdoc IAgentManager
    function slashAgent(uint32 domain, address agent, address prover) external onlyInbox {
        _slashAgent(domain, agent, prover);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {
        agent = _getAgent(index);
        if (agent != address(0)) status = agentStatus(agent);
    }

    /// @inheritdoc IAgentManager
    function agentStatus(address agent) public view returns (AgentStatus memory status) {
        status = _storedAgentStatus(agent);
        // If agent was proven to commit fraud, but their slashing wasn't completed, return the Fraudulent flag.
        if (_disputes[_getIndex(agent)].flag == DisputeFlag.Slashed && status.flag != AgentFlag.Slashed) {
            status.flag = AgentFlag.Fraudulent;
        }
    }

    /// @inheritdoc IAgentManager
    function disputeStatus(address agent) external view returns (Dispute memory) {
        return _disputes[_getIndex(agent)];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Hook that is called after agent was slashed in AgentManager and AgentSecured contracts were notified.
    // solhint-disable-next-line no-empty-blocks
    function _afterAgentSlashed(uint32 domain, address agent, address prover) internal virtual {}

    /// @dev Child contract should implement the logic for notifying AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal virtual;

    /// @dev Child contract should implement the logic for notifying AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal virtual;

    /// @dev Slashes the Agent and notifies the local Destination and Origin contracts about the slashed agent.
    /// Should be called when the agent fraud was confirmed.
    function _slashAgent(uint32 domain, address agent, address prover) internal {
        // Check that agent is Active/Unstaking and that the domains match
        AgentStatus memory status = _storedAgentStatus(agent);
        status.verifyActiveUnstaking();
        if (status.domain != domain) revert IncorrectAgentDomain();
        // The "stored" agent status is not updated yet, however agentStatus() will return AgentFlag.Fraudulent
        emit StatusUpdated(AgentFlag.Fraudulent, domain, agent);
        // This will revert if the agent has been slashed earlier
        _resolveDispute(status.index, prover);
        // Call "after slash" hook - this allows Bonding/Light Manager to add custom "after slash" logic
        _afterAgentSlashed(domain, agent, prover);
    }

    /// @dev Resolves a Dispute between a slashed Agent and their Rival (if there was one).
    function _resolveDispute(uint32 slashedIndex, address prover) internal {
        Dispute memory dispute = _disputes[slashedIndex];
        if (dispute.flag == DisputeFlag.Slashed) revert DisputeAlreadyResolved();
        (dispute.flag, dispute.fraudProver) = (DisputeFlag.Slashed, prover);
        _updateDispute(slashedIndex, dispute);
        // Clear Dispute status for the Rival
        if (dispute.rivalIndex != 0) {
            _updateDispute(dispute.rivalIndex, Dispute(DisputeFlag.None, 0, address(0)));
        }
        _notifyDisputeResolved(slashedIndex, dispute.rivalIndex);
    }

    /// @dev Updates a dispute status for the agent and emits an event.
    function _updateDispute(uint256 agentIndex, Dispute memory dispute) internal {
        _disputes[agentIndex] = dispute;
        emit DisputeUpdated(_getAgent(agentIndex), dispute);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Generates leaf to be saved in the Agent Merkle Tree
    function _agentLeaf(AgentFlag flag, uint32 domain, address agent) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(flag, domain, agent));
    }

    /// @dev Returns the last known status for the agent from the Agent Merkle Tree.
    /// Note: the actual agent status (returned by `agentStatus()`) may differ, if agent fraud was proven.
    function _storedAgentStatus(address agent) internal view virtual returns (AgentStatus memory);

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view virtual returns (address);

    /// @dev Returns the index of the agent in the Agent Merkle Tree. Returns zero for non existing agents.
    function _getIndex(address agent) internal view virtual returns (uint256);
}
