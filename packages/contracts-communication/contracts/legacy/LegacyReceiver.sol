// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyReceiverEvents} from "./events/LegacyReceiverEvents.sol";
import {ILegacyReceiver} from "./interfaces/ILegacyReceiver.sol";
import {IMessageBus} from "./interfaces/IMessageBus.sol";
import {LegacyOptionsLib} from "./libs/LegacyOptions.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

abstract contract LegacyReceiver is Ownable, LegacyReceiverEvents, ILegacyReceiver {
    address public messageBus;
    mapping(uint256 chainId => bytes32 trustedRemote) public trustedRemotes;

    error LegacyReceiver__BalanceTooLow(uint256 actual, uint256 required);
    error LegacyReceiver__NotMessageBus(address caller);
    error LegacyReceiver__NotTrustedRemote(uint256 chainId, bytes32 srcCaller);
    error LegacyReceiver__SameChainId(uint256 chainId);
    error LegacyReceiver__TrustedRemoteNotSet(uint256 chainId);

    constructor(address owner_) Ownable(owner_) {}

    function setMessageBus(address messageBus_) external onlyOwner {
        messageBus = messageBus_;
        emit MessageBusSet(messageBus_);
    }

    function setTrustedRemote(uint256 remoteChainId, bytes32 remoteApp) external onlyOwner {
        if (remoteChainId == block.chainid) {
            revert LegacyReceiver__SameChainId(remoteChainId);
        }
        trustedRemotes[remoteChainId] = remoteApp;
        emit TrustedRemoteSet(remoteChainId, remoteApp);
    }

    function executeMessage(
        bytes32 srcAddress,
        uint256 srcChainId,
        bytes calldata message,
        address executor
    )
        external
    {
        if (msg.sender != messageBus) {
            revert LegacyReceiver__NotMessageBus(msg.sender);
        }
        if (trustedRemotes[srcChainId] != srcAddress) {
            revert LegacyReceiver__NotTrustedRemote(srcChainId, srcAddress);
        }
        _handleMessage(srcAddress, srcChainId, message, executor);
    }

    /// @dev Handle the verified message.
    function _handleMessage(
        bytes32 srcAddress,
        uint256 srcChainId,
        bytes calldata message,
        address executor
    )
        internal
        virtual;

    /// @dev Sends a message to the trusted remote app.
    function _sendMessage(uint256 dstChainId, uint256 messageFee, uint256 gasLimit, bytes memory message) internal {
        bytes32 dstRemote = trustedRemotes[dstChainId];
        if (dstRemote == 0) {
            revert LegacyReceiver__TrustedRemoteNotSet(dstChainId);
        }
        if (address(this).balance < messageFee) {
            revert LegacyReceiver__BalanceTooLow(address(this).balance, messageFee);
        }
        bytes memory options = LegacyOptionsLib.encodeLegacyOptions(gasLimit);
        IMessageBus(messageBus).sendMessage{value: messageFee}({
            receiver: dstRemote,
            dstChainId: dstChainId,
            message: message,
            options: options
        });
    }

    /// @notice Calculates the fee to send a message to the remote app.
    function _getMessageFee(
        uint256 dstChainId,
        uint256 gasLimit,
        uint256 messageLen
    )
        internal
        view
        returns (uint256 fee)
    {
        bytes memory options = LegacyOptionsLib.encodeLegacyOptions(gasLimit);
        fee = IMessageBus(messageBus).estimateFeeExact(dstChainId, options, messageLen);
    }
}
