// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ICAppV1, OptionsV1} from "../../contracts/apps/ICAppV1.sol";
import {InterchainTxDescriptor} from "../../contracts/libs/InterchainTransaction.sol";

import {IInterchainAppV1Harness} from "../interfaces/IInterchainAppV1Harness.sol";

// solhint-disable func-name-mixedcase
contract ICAppV1Harness is ICAppV1, IInterchainAppV1Harness {
    constructor(address admin) ICAppV1(admin) {}

    function exposed__sendInterchainMessageEVM(
        uint256 dstChainId,
        address receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc)
    {
        return _sendInterchainMessageEVM(dstChainId, receiver, messageFee, options, message);
    }

    function exposed__sendInterchainMessage(
        uint256 dstChainId,
        bytes32 receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc)
    {
        return _sendInterchainMessage(dstChainId, receiver, messageFee, options, message);
    }

    function exposed__sendToLinkedApp(
        uint256 dstChainId,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        external
        returns (InterchainTxDescriptor memory desc)
    {
        return _sendToLinkedApp(dstChainId, messageFee, options, message);
    }

    function exposed__getInterchainFee(
        uint256 dstChainId,
        bytes memory options,
        uint256 messageLen
    )
        external
        view
        returns (uint256)
    {
        return _getInterchainFee(dstChainId, options, messageLen);
    }

    function exposed__getMessageFee(
        uint256 dstChainId,
        OptionsV1 memory options,
        uint256 messageLen
    )
        external
        view
        returns (uint256)
    {
        return _getMessageFee(dstChainId, options, messageLen);
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
