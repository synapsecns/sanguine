// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { IMessageRecipient } from "../../../contracts/interfaces/IMessageRecipient.sol";

// solhint-disable no-empty-blocks
contract AppHarness is IMessageRecipient {
    uint32 public optimisticSeconds;

    uint32 public expectedOrigin;
    uint32 public expectedNonce;
    bytes32 public expectedSender;
    bytes32 public expectedContentHash;

    constructor(uint32 optimisticSeconds_) {
        optimisticSeconds = optimisticSeconds_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testAppHarness() external {}

    function prepare(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        bytes memory content
    ) external {
        expectedOrigin = origin;
        expectedNonce = nonce;
        expectedSender = sender;
        expectedContentHash = keccak256(content);
    }

    function handle(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes memory content
    ) external view {
        require(block.timestamp >= rootSubmittedAt + optimisticSeconds, "app: !optimisticSeconds");
        require(origin == expectedOrigin, "app: !origin");
        require(nonce == expectedNonce, "app: !nonce");
        require(sender == expectedSender, "app: !sender");
        require(keccak256(content) == expectedContentHash, "app: !message");
    }
}
