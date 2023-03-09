// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { AgentInfo, SystemEntity } from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentRegistry } from "./AgentRegistry.sol";
import { ISystemContract, SystemContract } from "./SystemContract.sol";
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";

/**
 * @notice Shared agents registry utilities for Origin, Destination.
 * Agents are added/removed via a system call from a local BondingManager.
 */
abstract contract SystemRegistry is AgentRegistry, SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the off-chain agent needs to be slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _info             Information about agent to slash
     */
    function slashAgent(
        uint256,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: decide if we need to store anything, as the slashing occurred on another chain
        _beforeAgentSlashed(_info);
        _removeAgent(_info.domain, _info.account);
    }

    /// @inheritdoc ISystemContract
    function syncAgent(
        uint256,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        /// @dev Must be called from a local BondingManager. Hence `_rootSubmittedAt` is ignored.
        _updateAgentStatus(_info);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _updateAgentStatus(AgentInfo memory _info) internal {
        if (_info.bonded) {
            _addAgent(_info.domain, _info.account);
        } else {
            _removeAgent(_info.domain, _info.account);
        }
    }

    // solhint-disable no-empty-blocks
    /**
     * @notice Hook that is called before the specified agent was slashed via a system call.
     */
    function _beforeAgentSlashed(AgentInfo memory _info) internal virtual {}
}
