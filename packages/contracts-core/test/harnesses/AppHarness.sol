// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { IMessageRecipient } from "../../contracts/interfaces/IMessageRecipient.sol";

contract AppHarness is IMessageRecipient {
    uint32 public optimisticSeconds;

    uint32 public expectedOrigin;
    uint32 public expectedNonce;
    bytes32 public expectedSender;
    bytes32 public expectedMessageBodyHash;

    constructor(uint32 _optimisticSeconds) {
        optimisticSeconds = _optimisticSeconds;
    }

    function prepare(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        bytes memory _message
    ) external {
        expectedOrigin = _origin;
        expectedNonce = _nonce;
        expectedSender = _sender;
        expectedMessageBodyHash = keccak256(_message);
    }

    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootTimestamp,
        bytes memory _message
    ) external view {
        require(block.timestamp >= _rootTimestamp + optimisticSeconds, "app: !optimisticSeconds");
        require(_origin == expectedOrigin, "app: !origin");
        require(_nonce == expectedNonce, "app: !nonce");
        require(_sender == expectedSender, "app: !sender");
        require(keccak256(_message) == expectedMessageBodyHash, "app: !message");
    }
}
