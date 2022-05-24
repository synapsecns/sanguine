// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import {IMessageRecipient} from "../../../interfaces/IMessageRecipient.sol";

contract BadRecipient4 is IMessageRecipient {
    function handle(
        uint32,
        uint32,
        bytes32,
        bytes memory
    ) external pure override {
        assembly {
            mstore(0, 0xabcdef)
            return(0, 32)
        }
    }
}
