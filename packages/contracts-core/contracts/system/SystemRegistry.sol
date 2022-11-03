// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { SystemContract } from "./SystemContract.sol";
import { AbstractGuardRegistry } from "../registry/AbstractGuardRegistry.sol";
import { AbstractNotaryRegistry } from "../registry/AbstractNotaryRegistry.sol";

/**
 * @notice Shared agents registry utilities for Origin, Destination.
 * Agents are added/removed via a system call from a local BondingManager.
 */
abstract contract SystemRegistry is AbstractGuardRegistry, AbstractNotaryRegistry, SystemContract {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating that a new Notary staked a bond.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain           Domain where the new Notary will be active
     * @param _notary           New Notary that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function bondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: (applied below as well) _addNotary() can return false,
        // if Notary is already active. Determine if we need to revert in this case.
        _addNotary(_domain, _notary);
    }

    /**
     * @notice Receive a system call indicating that an active Notary unstaked their bond.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain           Domain where the Notary was active
     * @param _notary           Active Notary that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function unbondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        _removeNotary(_domain, _notary);
    }

    /**
     * @notice Receive a system call indicating that an active Notary was slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain           Domain where the slashed Notary was active
     * @param _notary           Active Notary that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function slashNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: decide if we need to store anything, as the slashing occurred on another chain
        _removeNotary(_domain, _notary);
    }

    /**
     * @notice Receive a system call indicating that a new Guard staked a bond.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard            New Guard that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function bondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        _addGuard(_guard);
    }

    /**
     * @notice Receive a system call indicating that an active Guard unstaked their bond.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard            Active Guard that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function unbondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        _removeGuard(_guard);
    }

    /**
     * @notice Receive a system call indicating that an active Guard was slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard            Active Guard that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     */
    function slashGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256
    ) external override onlySystemRouter onlyLocalBondingManager(_callOrigin, _caller) {
        // TODO: decide if we need to store anything, as the slashing occurred on another chain
        _removeGuard(_guard);
    }
}
