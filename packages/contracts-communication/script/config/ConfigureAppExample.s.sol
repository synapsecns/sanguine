// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ConfigureOwnableApp} from "./ConfigureOwnableApp.s.sol";

contract ConfigureAppExample is ConfigureOwnableApp {
    constructor() ConfigureOwnableApp("InterchainAppExample") {}
}
