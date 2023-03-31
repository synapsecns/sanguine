// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IMessageRecipient {
    function handle(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 rootSubmittedAt,
        bytes memory content
    ) external;
}
