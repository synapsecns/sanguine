// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainAppBaseEvents} from "../events/InterchainAppBaseEvents.sol";
import {IInterchainApp} from "../interfaces/IInterchainApp.sol";

import {AppConfigV1} from "../libs/AppConfig.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

abstract contract InterchainAppBase is InterchainAppBaseEvents, IInterchainApp {
    // TODO: naming, visibility
    address public interchain;

    /// @dev Required responses and optimistic period for the module responses.
    AppConfigV1 private _appConfig;

    /// @dev Address of the linked app deployed on the remote chain.
    mapping(uint256 chainId => bytes32 remoteApp) private _linkedApp;

    error InterchainApp__SameChainId(uint256 chainId);

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

    /// @dev Sets the app config:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    /// Note: Should be guarded with permissions check.
    function _setAppConfigV1(AppConfigV1 memory appConfig) internal {
        _appConfig = appConfig;
        emit AppConfigV1Set(appConfig.requiredResponses, appConfig.optimisticPeriod);
    }

    /// @dev Sets the interchain client address.
    /// Note: Should be guarded with permissions check.
    function _setInterchainClient(address interchain_) internal {
        interchain = interchain_;
        emit InterchainClientSet(interchain_);
    }
}
