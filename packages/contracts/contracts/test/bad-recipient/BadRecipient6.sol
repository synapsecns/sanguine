// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import {IMessageRecipient} from "../../../interfaces/IMessageRecipient.sol";

contract BadRecipient6 is IMessageRecipient {
    function handle(
        uint32,
        uint32,
        bytes32,
        bytes memory
    ) external pure override {
        require(false); // solhint-disable-line reason-string
    }
}
