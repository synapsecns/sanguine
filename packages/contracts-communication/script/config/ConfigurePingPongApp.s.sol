// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {PingPongApp} from "../../contracts/apps/PingPongApp.sol";

import {ConfigureOwnableApp, stdJson} from "./ConfigureOwnableApp.s.sol";

contract ConfigurePingPongApp is ConfigureOwnableApp {
    using stdJson for string;

    constructor() ConfigureOwnableApp("PingPongApp") {}

    function afterAppConfigured() internal override {
        printLog("Setting gas limit");
        PingPongApp pingPongApp = PingPongApp(payable(address(app)));
        uint256 gasLimit = config.readUint(".gasLimit");
        if (pingPongApp.gasLimit() != gasLimit) {
            pingPongApp.setGasLimit(gasLimit);
            printSuccessWithIndent(string.concat("Set gas limit to ", vm.toString(gasLimit)));
        } else {
            printSkipWithIndent(string.concat("gas limit already set to ", vm.toString(gasLimit)));
        }
    }
}
