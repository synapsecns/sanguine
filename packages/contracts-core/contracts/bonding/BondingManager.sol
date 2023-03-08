// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";
import { ISystemContract, SystemContract } from "../system/SystemContract.sol";
import "../Version.sol";

abstract contract BondingManager is SystemContract, Version0_0_2 {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating the off-chain agent needs to be slashed.
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _info             Information about agent to slash
     */
    function slashAgent(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller,
        AgentInfo memory _info
    ) external onlySystemRouter {
        bool forwardUpdate;
        if (_callOrigin == localDomain) {
            // Forward information about slashed agent to remote chains
            forwardUpdate = true;
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
        } else {
            // Forward information about slashed agent to remote chains
            // only if BondingManager is deployed on Synapse Chain
            forwardUpdate = _onSynapseChain();
            // Validate security params for cross-chain slashing
            _assertCrossChainSlashing(_rootSubmittedAt, _callOrigin, _caller);
        }
        // Forward information about the slashed agent to local Registries
        // Forward information about slashed agent to remote chains if needed
        _slashAgentLocalRegistries(_info, forwardUpdate, _callOrigin);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: generalize this further when Agent Merkle Tree is implemented

    /// @dev Passes an "update status" message to local Registries:
    /// that an Agent has been slashed
    function _slashAgentLocalRegistries(
        AgentInfo memory _info,
        bool _forwardUpdate,
        uint32 _callOrigin
    ) internal {
        _updateLocalRegistries(_dataSlashAgent(_info), _forwardUpdate, _callOrigin);
    }

    /// @dev Passes an "update status" message to local Registries:
    /// that an Agent has been added / removed
    function _syncAgentLocalRegistries(AgentInfo memory _info, uint32 _callOrigin) internal {
        // Forward information about added/removed agent to remote chains
        // only if BondingManager is deployed on Synapse Chain
        bool forwardUpdate = _onSynapseChain();
        _updateLocalRegistries(_dataSyncAgent(_info), forwardUpdate, _callOrigin);
    }

    /// @dev Passes an "update status" message to local Registries:
    /// that an Agent has been added / removed / slashed
    function _updateLocalRegistries(
        bytes memory _data,
        bool _forwardUpdate,
        uint32 _callOrigin
    ) internal {
        // Pass data to all System Registries. This could lead to duplicated data, meaning that
        // every Registry is responsible for ignoring the data it already has. This makes Registries
        // a bit more complex, but greatly reduces the complexity of BondingManager.
        systemRouter.systemMultiCall({
            _destination: localDomain,
            _optimisticSeconds: 0,
            _recipients: _localSystemRegistries(),
            _data: _data
        });
        // Forward data cross-chain, if requested
        if (_forwardUpdate) {
            _forwardUpdateData(_data, _callOrigin);
        }
    }

    /**
     * @notice Forward data with an agent status update (due to a system call from `_callOrigin`).
     * @dev If BondingManager is deployed on Synapse Chain, all chains should be notified,
     * excluding `_callOrigin` and Synapse Chain.
     * If BondingManager ois not deployed on Synapse CHain, only Synapse Chain should be notified.
     */
    function _forwardUpdateData(bytes memory _data, uint32 _callOrigin) internal virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Perform all required security checks for a cross-chain
     * system call for slashing an agent.
     */
    function _assertCrossChainSlashing(
        uint256 _rootSubmittedAt,
        uint32 _callOrigin,
        SystemEntity _caller
    ) internal view virtual;

    /**
     * @notice Returns a list of local System Registries: system contracts, keeping track
     * of active Notaries and Guards.
     */
    function _localSystemRegistries() internal pure returns (SystemEntity[] memory recipients) {
        recipients = new SystemEntity[](2);
        recipients[0] = SystemEntity.Origin;
        recipients[1] = SystemEntity.Destination;
    }
}
