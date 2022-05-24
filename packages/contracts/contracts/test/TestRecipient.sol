// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import {IMessageRecipient} from "../../interfaces/IMessageRecipient.sol";

contract TestRecipient is IMessageRecipient {
    bool public processed = false;

    // solhint-disable-next-line payable-fallback
    fallback() external {
        revert("Fallback");
    }

    function handle(
        uint32,
        uint32,
        bytes32,
        bytes memory
    ) external pure override {} // solhint-disable-line no-empty-blocks

    function receiveString(string calldata _str)
        public
        pure
        returns (string memory)
    {
        return _str;
    }

    function processCall(bool callProcessed) public {
        processed = callProcessed;
    }

    function message() public pure returns (string memory) {
        return "message received";
    }
}
