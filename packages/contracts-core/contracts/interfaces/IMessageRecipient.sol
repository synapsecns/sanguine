// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IMessageRecipient {
    function handle(
        uint32 _origin,
        uint32 _nonce,
        bytes32 _sender,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) external;
}
