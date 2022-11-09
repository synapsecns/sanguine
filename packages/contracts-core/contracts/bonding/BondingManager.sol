// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "../system/SystemContract.sol";

abstract contract BondingManager is SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating that an active Notary was slashed.
     * @param _domain           Domain where the slashed Notary was active
     * @param _notary           Active Notary that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function slashNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external override onlySystemRouter {
        bool forwardUpdate;
        if (_callOrigin == _localDomain()) {
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
            // Forward information about slashed Notary to remote chains
            forwardUpdate = true;
        } else {
            // Validate security params for cross-chain slashing
            _assertCrossChainSlashing(_callOrigin, _caller, _rootSubmittedAt);
            // Forward information about slashed Notary to remote chains
            // only if BondingManager is deployed on Synapse Chain
            forwardUpdate = _onSynapseChain();
        }
        // Forward information about slashed Notary to local system registries
        // Forward information about slashed Notary to remote chains if needed
        _localUpdateNotary({
            _selector: SystemContract.slashNotary.selector,
            _domain: _domain,
            _notary: _notary,
            _callOrigin: _callOrigin,
            _forwardUpdate: forwardUpdate
        });
    }

    /**
     * @notice Receive a system call indicating that an active Guard was slashed.
     * @param _guard            Active Guard that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function slashGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external override onlySystemRouter {
        bool forwardUpdate;
        if (_callOrigin == _localDomain()) {
            // Only Origin can slash agents on local domain
            _assertEntityAllowed(ORIGIN, _caller);
            // Forward information about slashed Guard to remote chains
            forwardUpdate = true;
        } else {
            // Validate security params for cross-chain slashing
            _assertCrossChainSlashing(_callOrigin, _caller, _rootSubmittedAt);
            // Forward information about slashed Guard to remote chains
            // only if BondingManager is deployed on Synapse Chain
            forwardUpdate = _onSynapseChain();
        }
        // Forward information about slashed Guard to local system registries
        // Forward information about slashed Guard to remote chains if needed
        _localUpdateGuard({
            _selector: SystemContract.slashNotary.selector,
            _guard: _guard,
            _callOrigin: _callOrigin,
            _forwardUpdate: forwardUpdate
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║          INTERNAL HELPERS: UPDATE AGENT (BOND/UNBOND/SLASH)          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Update Notary status on local system registries.
     * @dev Use selectors for bondNotary, unbondNotary or slashNotary based on the use case.
     */
    function _localUpdateNotary(
        bytes4 _selector,
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        bool _forwardUpdate
    ) internal {
        uint32 local = _localDomain();
        bytes memory data = abi.encodeWithSelector(_selector, _domain, _notary);
        // Notary is updated on Origin, only if domain matches the local domain
        // Notary is updated on Destination, only if domain doesn't match the local domain
        if (_domain == local) {
            // If a system call originated on local domain, it means Origin was the caller.
            // Meaning there's no need to call Origin again.
            if (_callOrigin != local) {
                systemRouter.systemCall({
                    _destination: local,
                    _optimisticSeconds: 0,
                    _recipient: ISystemRouter.SystemEntity.Origin,
                    _data: data
                });
            }
        } else {
            // Call Destination
            systemRouter.systemCall({
                _destination: local,
                _optimisticSeconds: 0,
                _recipient: ISystemRouter.SystemEntity.Destination,
                _data: data
            });
        }
        if (_forwardUpdate) {
            _forwardUpdateData(data, _callOrigin);
        }
    }

    /**
     * @notice Update Guard status on local system registries.
     * @dev Use selectors for bondGuard, unbondGuard or slashGuard based on the use case.
     */
    function _localUpdateGuard(
        bytes4 _selector,
        address _guard,
        uint32 _callOrigin,
        bool _forwardUpdate
    ) internal {
        bytes memory data = abi.encodeWithSelector(_selector, _guard);
        if (_callOrigin == _localDomain()) {
            // If a system call originated on local domain, it means Origin was the caller.
            // Meaning there's no need to call Origin again.
            systemRouter.systemCall({
                _destination: _localDomain(),
                _optimisticSeconds: 0,
                _recipient: ISystemRouter.SystemEntity.Destination,
                _data: data
            });
        } else {
            bytes[] memory dataArray = new bytes[](2);
            // TODO: use wrapper for system multicall instead
            dataArray[0] = dataArray[1] = data;
            systemRouter.systemMultiCall({
                _destination: _localDomain(),
                _optimisticSeconds: 0,
                _recipients: _localSystemRegistries(),
                _dataArray: dataArray
            });
        }
        if (_forwardUpdate) {
            _forwardUpdateData(data, _callOrigin);
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
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
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
