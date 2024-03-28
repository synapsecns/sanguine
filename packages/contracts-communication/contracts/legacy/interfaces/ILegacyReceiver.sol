// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ILegacyReceiver {
    function executeMessage(bytes32 srcAddress, uint256 srcChainId, bytes memory message, address executor) external;
}
