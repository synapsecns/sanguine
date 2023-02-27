// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IOrigin {
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) external returns (uint32 messageNonce, bytes32 messageHash);
}
