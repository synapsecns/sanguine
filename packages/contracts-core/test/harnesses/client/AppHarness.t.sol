// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";

contract AppHarness is IMessageRecipient {
    uint32 public optimisticSeconds;

    uint32 public expectedOrigin;
    uint32 public expectedNonce;
    bytes32 public expectedSender;
    bytes32 public expectedMessageBodyHash;

    constructor(uint32 optimisticSeconds_) {
        optimisticSeconds = optimisticSeconds_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testAppHarness() external {}

    function prepare(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        bytes memory message
    ) external {
        expectedOrigin = origin;
        expectedNonce = nonce;
        expectedSender = sender;
        expectedMessageBodyHash = keccak256(message);
    }

    function handle(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes memory message
    ) external view {
        require(block.timestamp >= rootSubmittedAt + optimisticSeconds, "app: !optimisticSeconds");
        require(origin == expectedOrigin, "app: !origin");
        require(nonce == expectedNonce, "app: !nonce");
        require(sender == expectedSender, "app: !sender");
        require(keccak256(message) == expectedMessageBodyHash, "app: !message");
    }
}
