// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

contract BadRecipientHandle {
    function handle(uint32, bytes32) external pure {} // solhint-disable-line no-empty-blocks
}
