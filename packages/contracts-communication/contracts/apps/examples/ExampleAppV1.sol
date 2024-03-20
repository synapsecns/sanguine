// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ICAppV1} from "../ICAppV1.sol";

import {InterchainTxDescriptor} from "../../libs/InterchainTransaction.sol";
import {OptionsV1} from "../../libs/Options.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

contract ExampleAppV1 is ICAppV1 {
    event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message);
    event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId);

    constructor(address admin) ICAppV1(admin) {
        _grantRole(IC_GOVERNOR_ROLE, admin);
    }

    /// @notice Allows the Admin to withdraw the native asset from the contract.
    function withdraw() external onlyRole(DEFAULT_ADMIN_ROLE) {
        Address.sendValue(payable(msg.sender), address(this).balance);
    }

    /// @notice Sends a basic message to the destination chain.
    function sendMessage(
        uint256 dstChainId,
        uint256 gasLimit,
        uint256 gasAirdrop,
        bytes calldata message
    )
        external
        payable
    {
        InterchainTxDescriptor memory desc = _sendToLinkedApp({
            dstChainId: dstChainId,
            messageFee: msg.value,
            options: OptionsV1({gasLimit: gasLimit, gasAirdrop: gasAirdrop}),
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
