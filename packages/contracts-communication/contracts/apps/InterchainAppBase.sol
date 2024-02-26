// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainAppBaseEvents} from "../events/InterchainAppBaseEvents.sol";
import {IInterchainApp} from "../interfaces/IInterchainApp.sol";

import {AppConfigV1} from "../libs/AppConfig.sol";

abstract contract InterchainAppBase is InterchainAppBaseEvents, IInterchainApp {
    // TODO: naming, visibility
    address public interchain;

    /// @dev Required responses and optimistic period for the module responses.
    AppConfigV1 private _appConfig;

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
