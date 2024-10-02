// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridgeRecipient {
    function fastBridgeTransferReceived(
        address token,
        uint256 amount,
        bytes memory callParams
    )
        external
        payable
        returns (bytes4);
}
