// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBus} from "../../../contracts/legacy/MessageBus.sol";
import {TypeCasts} from "../../../contracts/libs/TypeCasts.sol";

import {stdJson, ConfigureAppV1} from "../ConfigureAppV1.s.sol";

contract ConfigureMessageBus is ConfigureAppV1 {
    using stdJson for string;

    constructor() ConfigureAppV1("MessageBus") {}

    function beforeAppConfigured() internal override {
        setInterchainGovernor();
    }

    function afterAppConfigured() internal override {
        setMessageLengthEstimate();
    }

    function setInterchainGovernor() internal {
        printLog("Setting Interchain Governor");
        bytes32 icGovernorRole = app.IC_GOVERNOR_ROLE();
        if (!app.hasRole(icGovernorRole, msg.sender)) {
            app.grantRole(icGovernorRole, msg.sender);
            printSuccessWithIndent("Granted IC_GOVERNOR_ROLE");
        } else {
            printSkipWithIndent("IC_GOVERNOR_ROLE already granted");
        }
    }

    function setMessageLengthEstimate() internal {
        printLog("Setting message length estimate");
        MessageBus messageBus = MessageBus(payable(address(app)));
        uint256 messageLengthEstimate = config.readUint(".messageLengthEstimate");
        if (messageBus.messageLengthEstimate() != messageLengthEstimate) {
            messageBus.setMessageLengthEstimate(messageLengthEstimate);
            printSuccessWithIndent(string.concat("Set to ", vm.toString(messageLengthEstimate)));
        } else {
            printSkipWithIndent(string.concat("already set to ", vm.toString(messageLengthEstimate)));
        }
    }
}
