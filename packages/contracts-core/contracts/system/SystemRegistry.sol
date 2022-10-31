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
     * @notice Adds a new notary, as they staked their bond on SynChain.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain       Domain, where new Notary is active
     * @param _notary       New Notary to add
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
     * @notice Removes an active notary, as they requested the unstaking of their bond on SynChain.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain       Domain, where new Notary was active
     * @param _notary       Active Notary to remove
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
     * @notice Removes an active notary, as they were slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _domain       Domain, where new Notary was active
     * @param _notary       Active Notary to remove
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
     * @notice Adds a new guard, as they staked their bond on SynChain.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard        New Guard to add
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
     * @notice Removes an active guard, as they requested the unstaking of their bond on SynChain.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard        Active Guard to remove
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
     * @notice Removes an active guard, as they were slashed.
     * @dev Must be called from a local BondingManager. Therefore
     * `uint256 _rootSubmittedAt` is ignored.
     * @param _guard        Active Guard to remove
     * @param _callOrigin   Domain, where system call originated
     * @param _caller       Entity which performed a system call
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
