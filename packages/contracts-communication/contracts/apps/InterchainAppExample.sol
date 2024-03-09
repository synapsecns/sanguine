// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OwnableApp} from "./OwnableApp.sol";
import {InterchainTxDescriptor} from "../libs/InterchainTransaction.sol";
import {OptionsV1} from "../libs/Options.sol";

contract InterchainAppExample is OwnableApp {
    event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message);
    event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId);

    constructor(address owner_) OwnableApp(owner_) {}

    /// @notice Sends a basic message to the destination chain.
    function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes calldata message) external payable {
        InterchainTxDescriptor memory desc = _sendInterchainMessage({
            dstChainId: dstChainId,
            messageFee: msg.value,
            options: OptionsV1({gasLimit: gasLimit, gasAirdrop: 0}),
            message: message
        });
        emit MessageSent(dstChainId, desc.dbNonce, desc.entryIndex, desc.transactionId);
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        internal
        override
    {
        emit MessageReceived(srcChainId, sender, dbNonce, entryIndex, message);
    }
}
