// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainAppBaseEvents} from "../events/InterchainAppBaseEvents.sol";
import {IInterchainApp} from "../interfaces/IInterchainApp.sol";

abstract contract InterchainAppBase is InterchainAppBaseEvents, IInterchainApp {
    // TODO: naming
    address public interchain;

    /// @dev Sets the interchain client address.
    /// Note: Should be guarded with permissions check.
    function _setInterchainClient(address interchain_) internal {
        interchain = interchain_;
        emit InterchainClientSet(interchain_);
    }
}
