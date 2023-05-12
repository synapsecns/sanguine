// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {CallerNotAgentManager, CallerNotInbox} from "../libs/Errors.sol";
import {AgentStatus, DisputeFlag} from "../libs/Structures.sol";
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

    // (agent index => their dispute flag: None/Pending/Slashed)
    mapping(uint32 => DisputeFlag) internal _disputes;

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

    constructor(string memory version_, uint32 localDomain_, address agentManager_, address inbox_)
        MessagingBase(version_, localDomain_)
    {
        agentManager = agentManager_;
        inbox = inbox_;
    }

    // ════════════════════════════════════════════ ONLY AGENT MANAGER ═════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external onlyAgentManager {
        _disputes[guardIndex] = DisputeFlag.Pending;
        _disputes[notaryIndex] = DisputeFlag.Pending;
    }

    /// @inheritdoc IAgentSecured
    function resolveDispute(uint32 slashedIndex, uint32 honestIndex) external onlyAgentManager {
        _disputes[slashedIndex] = DisputeFlag.Slashed;
        if (honestIndex != 0) delete _disputes[honestIndex];
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

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address agent) internal view returns (AgentStatus memory) {
        return IAgentManager(agentManager).agentStatus(agent);
    }

    /// @dev Returns agent and their status for a given agent index. Returns zero values for non existing indexes.
    function _getAgent(uint256 index) internal view returns (address agent, AgentStatus memory status) {
        return IAgentManager(agentManager).getAgent(index);
    }

    /// @dev Checks if the agent with the given index is in a dispute.
    function _isInDispute(uint32 agentIndex) internal view returns (bool) {
        return _disputes[agentIndex] != DisputeFlag.None;
    }
}
