// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainAppBaseEvents} from "../events/InterchainAppBaseEvents.sol";
import {IInterchainApp} from "../interfaces/IInterchainApp.sol";

import {AppConfigV1} from "../libs/AppConfig.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract InterchainAppBase is InterchainAppBaseEvents, IInterchainApp {
    using EnumerableSet for EnumerableSet.AddressSet;

    // TODO: naming, visibility
    address public interchain;

    /// @dev Required responses and optimistic period for the module responses.
    AppConfigV1 private _appConfig;
    /// @dev Address of the linked app deployed on the remote chain.
    mapping(uint256 chainId => bytes32 remoteApp) private _linkedApp;
    /// @dev Trusted Interchain modules.
    EnumerableSet.AddressSet private _trustedModules;
    /// @dev Execution Service to use for sending messages.
    address private _executionService;

    error InterchainApp__ModuleAlreadyAdded(address module);
    error InterchainApp__ModuleNotAdded(address module);
    error InterchainApp__SameChainId(uint256 chainId);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the address of the Execution Service to use for sending messages.
    /// @dev Could be overridden in the derived contracts.
    function getExecutionService() public view virtual returns (address) {
        return _executionService;
    }

    /// @notice Returns the list of modules used for sending messages.
    /// @dev Could be overridden in the derived contracts.
    function getSendingModules() public view virtual returns (address[] memory) {
        return _trustedModules.values();
    }

    /// @notice Returns the list of modules used for receiving messages.
    /// @dev Could be overridden in the derived contracts.
    function getReceivingModules() public view virtual returns (address[] memory) {
        return _trustedModules.values();
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Links the remote app to the current app.
    /// Will revert if the chainId is the same as the chainId of the local app.
    /// Note: Should be guarded with permissions check.
    function _linkRemoteApp(uint256 chainId, bytes32 remoteApp) internal {
        if (chainId == block.chainid) revert InterchainApp__SameChainId(chainId);
        _linkedApp[chainId] = remoteApp;
        emit AppLinked(chainId, remoteApp);
    }

    /// @dev Thin wrapper around _linkRemoteApp to accept EVM address as a parameter.
    function _linkRemoteAppEVM(uint256 chainId, address remoteApp) internal {
        _linkRemoteApp(chainId, TypeCasts.addressToBytes32(remoteApp));
    }

    /// @dev Adds the module to the trusted modules set.
    /// Will revert if the module is already added.
    /// Note: Should be guarded with permissions check.
    function _addTrustedModule(address module) internal {
        bool added = _trustedModules.add(module);
        if (!added) revert InterchainApp__ModuleAlreadyAdded(module);
        emit TrustedModuleAdded(module);
    }

    /// @dev Removes the module from the trusted modules set.
    /// Will revert if the module is not added.
    /// Note: Should be guarded with permissions check.
    function _removeTrustedModule(address module) internal {
        bool removed = _trustedModules.remove(module);
        if (!removed) revert InterchainApp__ModuleNotAdded(module);
        emit TrustedModuleRemoved(module);
    }

    /// @dev Sets the app config:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    /// Note: Should be guarded with permissions check.
    function _setAppConfigV1(AppConfigV1 memory appConfig) internal {
        _appConfig = appConfig;
        emit AppConfigV1Set(appConfig.requiredResponses, appConfig.optimisticPeriod);
    }

    /// @dev Sets the execution service address.
    /// Note: Should be guarded with permissions check.
    function _setExecutionService(address executionService) internal {
        _executionService = executionService;
        emit ExecutionServiceSet(executionService);
    }

    /// @dev Sets the interchain client address.
    /// Note: Should be guarded with permissions check.
    function _setInterchainClient(address interchain_) internal {
        interchain = interchain_;
        emit InterchainClientSet(interchain_);
    }

    // ════════════════════════════════════════════ INTERNAL: MESSAGING ════════════════════════════════════════════════
}
