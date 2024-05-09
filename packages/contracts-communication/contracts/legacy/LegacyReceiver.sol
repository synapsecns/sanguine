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

    error LegacyReceiver__BalanceBelowMin(uint256 actual, uint256 required);
    error LegacyReceiver__ChainIdNotRemote(uint256 chainId);
    error LegacyReceiver__CallerNotMessageBus(address caller);
    error LegacyReceiver__SrcCallerNotTrusted(uint256 chainId, bytes32 srcCaller);
    error LegacyReceiver__TrustedRemoteZeroAddress(uint256 chainId);

    constructor(address owner_) Ownable(owner_) {}

    function setMessageBus(address messageBus_) external onlyOwner {
        messageBus = messageBus_;
        emit MessageBusSet(messageBus_);
    }

    function setTrustedRemote(uint256 remoteChainId, bytes32 remoteApp) external onlyOwner {
        if (remoteChainId == block.chainid) {
            revert LegacyReceiver__ChainIdNotRemote(remoteChainId);
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
            revert LegacyReceiver__CallerNotMessageBus(msg.sender);
        }
        if (trustedRemotes[srcChainId] != srcAddress) {
            revert LegacyReceiver__SrcCallerNotTrusted(srcChainId, srcAddress);
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
            revert LegacyReceiver__TrustedRemoteZeroAddress(dstChainId);
        }
        if (address(this).balance < messageFee) {
            revert LegacyReceiver__BalanceBelowMin(address(this).balance, messageFee);
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
