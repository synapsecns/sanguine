// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {DISPUTE_TIMEOUT_NOTARY} from "../libs/Constants.sol";
import {ChainContext} from "../libs/ChainContext.sol";
import {CallerNotAgentManager, CallerNotInbox} from "../libs/Errors.sol";
import {AgentStatus, DisputeFlag, DisputeStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {IAgentSecured} from "../interfaces/IAgentSecured.sol";
import {MessagingBase} from "./MessagingBase.sol";

/**
 * @notice Base contract for messaging contracts that are secured by the agent manager.
 * `AgentSecured` relies on `AgentManager` to provide the following functionality:
 * - Keep track of agents and their statuses.
 * - Pass agent-signed statements that were verified by the agent manager.
 * - These statements are considered valid indefinitely, unless the agent is disputed.
 * - Disputes are opened and resolved by the agent manager.
 * > `AgentSecured` implementation should never use statements signed by agents that are disputed.
 */
abstract contract AgentSecured is MessagingBase, IAgentSecured {
    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    address public immutable agentManager;

    /// @inheritdoc IAgentSecured
    address public immutable inbox;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (agent index => their dispute status: flag, openedAt, resolvedAt)
    mapping(uint32 => DisputeStatus) internal _disputes;

    /// @dev gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyAgentManager() {
        if (msg.sender != agentManager) revert CallerNotAgentManager();
        _;
    }

    modifier onlyInbox() {
        if (msg.sender != inbox) revert CallerNotInbox();
        _;
    }

    constructor(string memory version_, uint32 synapseDomain_, address agentManager_, address inbox_)
        MessagingBase(version_, synapseDomain_)
    {
        agentManager = agentManager_;
        inbox = inbox_;
    }

    // ════════════════════════════════════════════ ONLY AGENT MANAGER ═════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external onlyAgentManager {
        uint40 openedAt = ChainContext.blockTimestamp();
        DisputeStatus memory status = DisputeStatus({flag: DisputeFlag.Pending, openedAt: openedAt, resolvedAt: 0});
        _disputes[guardIndex] = status;
        _disputes[notaryIndex] = status;
    }

    /// @inheritdoc IAgentSecured
    function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) external onlyAgentManager {
        // Update the dispute status of the slashed agent first.
        uint40 resolvedAt = ChainContext.blockTimestamp();
        _disputes[slashedIndex].flag = DisputeFlag.Slashed;
        _disputes[slashedIndex].resolvedAt = resolvedAt;
        // Mark the rival agent as not disputed, if there was an ongoing dispute.
        if (rivalIndex != 0) {
            _disputes[rivalIndex].flag = DisputeFlag.None;
            _disputes[rivalIndex].resolvedAt = resolvedAt;
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function agentStatus(address agent) external view returns (AgentStatus memory) {
        return _agentStatus(agent);
    }

    /// @inheritdoc IAgentSecured
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {
        return _getAgent(index);
    }

    /// @inheritdoc IAgentSecured
    function latestDisputeStatus(uint32 agentIndex) external view returns (DisputeStatus memory) {
        return _disputes[agentIndex];
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address agent) internal view returns (AgentStatus memory) {
        return IAgentManager(agentManager).agentStatus(agent);
    }

    /// @dev Returns agent and their status for a given agent index. Returns zero values for non existing indexes.
    function _getAgent(uint256 index) internal view returns (address agent, AgentStatus memory status) {
        return IAgentManager(agentManager).getAgent(index);
    }

    /// @dev Checks if a Dispute exists for the given Notary. This function returns true, if
    /// Notary is in ongoing Dispute, or if Dispute was resolved not in Notary's favor.
    /// In both cases we can't trust Notary's data.
    /// Note: Agent-Secured contracts can trust Notary data only if both `_notaryDisputeExists` and
    /// `_notaryDisputeTimeout` return false.
    function _notaryDisputeExists(uint32 notaryIndex) internal view returns (bool) {
        return _disputes[notaryIndex].flag != DisputeFlag.None;
    }

    /// @dev Checks if a Notary recently won a Dispute and is still in the "post-dispute" timeout period.
    /// In this period we still can't trust Notary's data, though we can optimistically assume that
    /// that the data will be correct after the timeout (assuming no new Disputes are opened).
    /// Note: Agent-Secured contracts can trust Notary data only if both `_notaryDisputeExists` and
    /// `_notaryDisputeTimeout` return false.
    function _notaryDisputeTimeout(uint32 notaryIndex) internal view returns (bool) {
        DisputeStatus memory status = _disputes[notaryIndex];
        // Exit early if Notary is in ongoing Dispute / slashed.
        if (status.flag != DisputeFlag.None) return false;
        // Check if Notary has been in any Dispute at all.
        if (status.openedAt == 0) return false;
        // Otherwise check if the Dispute timeout is still active.
        return block.timestamp < status.resolvedAt + DISPUTE_TIMEOUT_NOTARY;
    }
}
