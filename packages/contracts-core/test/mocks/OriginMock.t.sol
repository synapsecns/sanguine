// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./system/SystemContractMock.t.sol";

// solhint-disable no-empty-blocks
contract OriginMock is SystemContractMock {
    uint32 public messagesDispatched;
    bytes[] public messageBodies;

    constructor(uint32 _domain) SystemContractMock(_domain) {}

    function dispatch(
        uint32,
        bytes32,
        uint32,
        bytes memory,
        bytes memory _messageBody
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        messageNonce = ++messagesDispatched;
        messageBodies.push(_messageBody);
        messageHash;
    }
}
