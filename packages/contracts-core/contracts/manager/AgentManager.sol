// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {FRESH_DATA_TIMEOUT} from "../libs/Constants.sol";
import {
    CallerNotInbox,
    DisputeAlreadyResolved,
    IncorrectAgentDomain,
    IndexOutOfRange,
    GuardInDispute,
    NotaryInDispute,
    NotStuck
} from "../libs/Errors.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "../base/MessagingBase.sol";
import {AgentManagerEvents} from "../events/AgentManagerEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {IStatementInbox} from "../interfaces/IStatementInbox.sol";

/// @notice `AgentManager` is used to keep track of all the bonded agents and their statuses.
/// The exact logic of how the agent statuses are stored and updated is implemented in child contracts,
/// and depends on whether the contract is used on Synapse Chain or on other chains.
/// `AgentManager` is responsible for the following:
/// - Keeping track of all the bonded agents and their statuses.
/// - Keeping track of all the disputes between agents.
/// - Notifying `AgentSecured` contracts about the opened and resolved disputes.
/// - Notifying `AgentSecured` contracts about the slashed agents.
abstract contract AgentManager is MessagingBase, AgentManagerEvents, IAgentManager {
    struct AgentDispute {
        DisputeFlag flag;
        uint88 disputePtr;
        address fraudProver;
    }

    // TODO: do we want to store the dispute timestamp?
    struct OpenedDispute {
        uint32 guardIndex;
        uint32 notaryIndex;
        uint32 slashedIndex;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    address public origin;

    address public destination;

    address public inbox;

    // (agent index => their dispute status)
    mapping(uint256 => AgentDispute) internal _agentDispute;

    // All disputes ever opened
    OpenedDispute[] internal _disputes;

    /// @dev gap for upgrade safety
    uint256[45] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyInbox() {
        if (msg.sender != inbox) revert CallerNotInbox();
        _;
    }

    modifier onlyWhenStuck() {
        // Check if there has been no fresh data from the Notaries for a while.
        (uint40 snapRootTime,,) = InterfaceDestination(destination).destStatus();
        if (block.timestamp < FRESH_DATA_TIMEOUT + snapRootTime) revert NotStuck();
        _;
    }

    // ════════════════════════════════════════════════ INITIALIZER ════════════════════════════════════════════════════

    // solhint-disable-next-line func-name-mixedcase
    function __AgentManager_init(address origin_, address destination_, address inbox_, address owner_)
        internal
        onlyInitializing
    {
        __MessagingBase_init(owner_);
        origin = origin_;
        destination = destination_;
        inbox = inbox_;
    }

    // ════════════════════════════════════════════════ ONLY INBOX ═════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    // solhint-disable-next-line ordering
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external onlyInbox {
        // Check that both agents are not in Dispute yet.
        if (_agentDispute[guardIndex].flag != DisputeFlag.None) revert GuardInDispute();
        if (_agentDispute[notaryIndex].flag != DisputeFlag.None) revert NotaryInDispute();
        _disputes.push(OpenedDispute(guardIndex, notaryIndex, 0));
        // Dispute is stored at length - 1, but we store the index + 1 to distinguish from "not in dispute".
        // TODO: check if we really need to use 88 bits for dispute indexes. Every dispute ends up in one of
        // the agents being slashed, so the number of disputes is limited by the number of agents (currently 2**32).
        // Thus we can do the unsafe cast to uint88.
        uint88 disputePtr = uint88(_disputes.length);
        _agentDispute[guardIndex] = AgentDispute(DisputeFlag.Pending, disputePtr, address(0));
        _agentDispute[notaryIndex] = AgentDispute(DisputeFlag.Pending, disputePtr, address(0));
        // Dispute index is length - 1. Note: report that initiated the dispute has the same index in `Inbox`.
        emit DisputeOpened({disputeIndex: disputePtr - 1, guardIndex: guardIndex, notaryIndex: notaryIndex});
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
        if (_agentDispute[_getIndex(agent)].flag == DisputeFlag.Slashed && status.flag != AgentFlag.Slashed) {
            status.flag = AgentFlag.Fraudulent;
        }
    }

    /// @inheritdoc IAgentManager
    function getDisputesAmount() external view returns (uint256) {
        return _disputes.length;
    }

    /// @inheritdoc IAgentManager
    function getDispute(uint256 index)
        external
        view
        returns (
            address guard,
            address notary,
            address slashedAgent,
            address fraudProver,
            bytes memory reportPayload,
            bytes memory reportSignature
        )
    {
        if (index >= _disputes.length) revert IndexOutOfRange();
        OpenedDispute memory dispute = _disputes[index];
        guard = _getAgent(dispute.guardIndex);
        notary = _getAgent(dispute.notaryIndex);
        if (dispute.slashedIndex > 0) {
            slashedAgent = _getAgent(dispute.slashedIndex);
            fraudProver = _agentDispute[dispute.slashedIndex].fraudProver;
        }
        (reportPayload, reportSignature) = IStatementInbox(inbox).getGuardReport(index);
    }

    /// @inheritdoc IAgentManager
    function disputeStatus(address agent)
        external
        view
        returns (DisputeFlag flag, address rival, address fraudProver, uint256 disputePtr)
    {
        uint256 agentIndex = _getIndex(agent);
        AgentDispute memory agentDispute = _agentDispute[agentIndex];
        flag = agentDispute.flag;
        fraudProver = agentDispute.fraudProver;
        disputePtr = agentDispute.disputePtr;
        if (disputePtr > 0) {
            OpenedDispute memory dispute = _disputes[disputePtr - 1];
            rival = _getAgent(dispute.guardIndex == agentIndex ? dispute.notaryIndex : dispute.guardIndex);
        }
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
        AgentDispute memory agentDispute = _agentDispute[slashedIndex];
        if (agentDispute.flag == DisputeFlag.Slashed) revert DisputeAlreadyResolved();
        agentDispute.flag = DisputeFlag.Slashed;
        agentDispute.fraudProver = prover;
        _agentDispute[slashedIndex] = agentDispute;
        // Check if there was a opened dispute with the slashed agent
        uint32 rivalIndex = 0;
        if (agentDispute.disputePtr != 0) {
            uint256 disputeIndex = agentDispute.disputePtr - 1;
            OpenedDispute memory dispute = _disputes[disputeIndex];
            _disputes[disputeIndex].slashedIndex = slashedIndex;
            // Clear the dispute status for the rival
            rivalIndex = dispute.notaryIndex == slashedIndex ? dispute.guardIndex : dispute.notaryIndex;
            delete _agentDispute[rivalIndex];
            emit DisputeResolved(disputeIndex, slashedIndex, rivalIndex, prover);
        }
        _notifyDisputeResolved(slashedIndex, rivalIndex);
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
