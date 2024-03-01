// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OwnableApp} from "./OwnableApp.sol";
import {OptionsV1} from "../libs/Options.sol";

contract InterchainAppExample is OwnableApp {
    event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes message);
    event MessageSent(uint256 dstChainId, uint256 dbNonce, bytes32 transactionId);

    constructor(address owner_) OwnableApp(owner_) {}

    /// @notice Sends a basic message to the destination chain.
    function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes calldata message) external payable {
        (bytes32 transactionId, uint256 dbNonce) = _sendInterchainMessage({
            dstChainId: dstChainId,
            messageFee: msg.value,
            options: OptionsV1({gasLimit: gasLimit, gasAirdrop: 0}),
            message: message
        });
        emit MessageSent(dstChainId, dbNonce, transactionId);
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        bytes calldata message
    )
        internal
        override
    {
        emit MessageReceived(srcChainId, sender, dbNonce, message);
    }
}
