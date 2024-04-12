// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBus} from "../../contracts/legacy/MessageBus.sol";

/// @notice Harness for TESTS ONLY
contract MessageBusHarness is MessageBus {
    constructor(address admin) MessageBus(admin) {}

    function setNonce(uint64 nonce_) external {
        nonce = nonce_;
    }
}
