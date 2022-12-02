// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";

abstract contract BondingManager is SystemContract {
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
        ISystemRouter.SystemEntity _caller,
        AgentInfo memory _info
    ) external override onlySystemRouter {
        bool forwardUpdate;
        if (_callOrigin == _localDomain()) {
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
            // Forward information about slashed agent to remote chains
            forwardUpdate = true;
        } else {
            // Validate security params for cross-chain slashing
            _assertCrossChainSlashing(_rootSubmittedAt, _callOrigin, _caller);
            // Forward information about slashed agent to remote chains
            // only if BondingManager is deployed on Synapse Chain
            forwardUpdate = _onSynapseChain();
        }
        // Forward information about the slashed agent to local Registries
        // Forward information about slashed agent to remote chains if needed
        _updateLocalRegistries({
            _data: _dataSlashAgent(_info),
            _forwardUpdate: forwardUpdate,
            _callOrigin: _callOrigin
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _updateLocalRegistries(
        bytes memory _data,
        bool _forwardUpdate,
        uint32 _callOrigin
    ) internal {
        // Pass data to all System Registries. This could lead to duplicated data, meaning that
        // every Registry is responsible for ignoring the data it already has. This makes Registries
        // a bit more complex, but greatly reduces the complexity of BondingManager.
        systemRouter.systemMultiCall({
            _destination: _localDomain(),
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
     * @notice Forward data with an agent status update (due to
     * a system call from `_callOrigin`).
     * @dev If BondingManager is deployed on Synapse Chain, all other chains should be notified.
     * Otherwise, only Synapse Chain should be notified.
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
        ISystemRouter.SystemEntity _caller
    ) internal view virtual;

    /**
     * @notice Returns a list of local System Registries: system contracts, keeping track
     * of active Notaries and Guards.
     */
    function _localSystemRegistries()
        internal
        pure
        returns (ISystemRouter.SystemEntity[] memory recipients)
    {
        recipients = new ISystemRouter.SystemEntity[](2);
        recipients[0] = ISystemRouter.SystemEntity.Origin;
        recipients[1] = ISystemRouter.SystemEntity.Destination;
    }
}
