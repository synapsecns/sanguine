// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AbstractICAppEvents} from "../events/AbstractICAppEvents.sol";

import {IInterchainApp} from "../interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "../interfaces/IInterchainClientV1.sol";

import {InterchainTxDescriptor} from "../libs/InterchainTransaction.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

abstract contract AbstractICApp is AbstractICAppEvents, IInterchainApp {
    using TypeCasts for address;

    error InterchainApp__AlreadyLatestClient(address client);
    error InterchainApp__BalanceBelowMin(uint256 balance, uint256 minRequired);
    error InterchainApp__ChainIdNotRemote(uint64 chainId);
    error InterchainApp__ClientAlreadyAdded(address client);
    error InterchainApp__InterchainClientZeroAddress();
    error InterchainApp__NotInterchainClient(address account);
    error InterchainApp__ReceiverZeroAddress(uint64 chainId);
    error InterchainApp__SenderNotAllowed(uint64 srcChainId, bytes32 sender);

    /// @inheritdoc IInterchainApp
    function appReceive(
        uint64 srcChainId,
        bytes32 sender,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable
    {
        if (!_isInterchainClient(msg.sender)) {
            revert InterchainApp__NotInterchainClient(msg.sender);
        }
        if (srcChainId == block.chainid) {
            revert InterchainApp__ChainIdNotRemote(srcChainId);
        }
        if (!_isAllowedSender(srcChainId, sender)) {
            revert InterchainApp__SenderNotAllowed(srcChainId, sender);
        }
        _receiveMessage(srcChainId, sender, dbNonce, entryIndex, message);
    }

    /// @inheritdoc IInterchainApp
    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {
        appConfig = _getAppConfig();
        modules = _getModules();
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Performs necessary checks and adds an Interchain Client.
    /// Optionally sets the latest client to this one.
    /// Note: should be guarded with permission checks in the derived contracts.
    function _addClient(address client, bool updateLatest) internal {
        if (client == address(0)) {
            revert InterchainApp__InterchainClientZeroAddress();
        }
        if (_isInterchainClient(client)) {
            revert InterchainApp__ClientAlreadyAdded(client);
        }
        _toggleClientState(client, true);
        emit InterchainClientAdded(client);
        if (updateLatest) {
            _setLatestClient(client);
        }
    }

    /// @dev Performs necessary checks and removes an Interchain Client. If this client is the latest one,
    /// the latest client is set to zero address (effectively pausing the app ability to send messages).
    /// Note: should be guarded with permission checks in the derived contracts.
    function _removeClient(address client) internal {
        if (!_isInterchainClient(client)) {
            revert InterchainApp__NotInterchainClient(client);
        }
        _toggleClientState(client, false);
        emit InterchainClientRemoved(client);
        if (client == _getLatestClient()) {
            _setLatestClient(address(0));
        }
    }

    /// @dev Sets the latest Interchain Client to one of the allowed clients. Setting the client to zero address
    /// is allowed and effectively pauses the app ability to send messages (but still allows to receive them).
    /// Note: should be guarded with permission checks in the derived contracts.
    function _setLatestClient(address client) internal {
        // New latest client must be an allowed client or zero address.
        if (!_isInterchainClient(client) && client != address(0)) {
            revert InterchainApp__NotInterchainClient(client);
        }
        if (client == _getLatestClient()) {
            revert InterchainApp__AlreadyLatestClient(client);
        }
        _storeLatestClient(client);
        emit LatestClientSet(client);
    }

    /// @dev Stores the address of the latest Interchain Client.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_setLatestClient` instead.
    /// - Should not emit any events: this is done in the calling function.
    function _storeLatestClient(address client) internal virtual;

    /// @dev Toggle the state of the Interchain Client (allowed/disallowed to send messages to this app).
    /// - The client is checked to be in the opposite state before the change.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_addClient` and `_removeClient` instead.
    /// - Should not emit any events: this is done in the calling functions.
    function _toggleClientState(address client, bool allowed) internal virtual;

    // ════════════════════════════════════════════ INTERNAL: MESSAGING ════════════════════════════════════════════════

    /// @dev Thin wrapper around _sendInterchainMessage to accept EVM address as a parameter.
    function _sendInterchainMessageEVM(
        uint64 dstChainId,
        address receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        internal
        returns (InterchainTxDescriptor memory desc)
    {
        return _sendInterchainMessage(dstChainId, receiver.addressToBytes32(), messageFee, options, message);
    }

    /// @dev Performs necessary checks and sends an interchain message.
    function _sendInterchainMessage(
        uint64 dstChainId,
        bytes32 receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        internal
        returns (InterchainTxDescriptor memory desc)
    {
        address client = _getLatestClient();
        if (client == address(0)) {
            revert InterchainApp__InterchainClientZeroAddress();
        }
        if (dstChainId == block.chainid) {
            revert InterchainApp__ChainIdNotRemote(dstChainId);
        }
        if (receiver == 0) {
            revert InterchainApp__ReceiverZeroAddress(dstChainId);
        }
        if (address(this).balance < messageFee) {
            revert InterchainApp__BalanceBelowMin({balance: address(this).balance, minRequired: messageFee});
        }
        return IInterchainClientV1(client).interchainSend{value: messageFee}(
            dstChainId, receiver, _getExecutionService(), _getModules(), options, message
        );
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint64 srcChainId,
        bytes32 sender,
        uint64 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        internal
        virtual;

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the fee for sending an Interchain message.
    function _getInterchainFee(
        uint64 dstChainId,
        bytes memory options,
        uint256 messageLen
    )
        internal
        view
        returns (uint256)
    {
        address client = _getLatestClient();
        if (client == address(0)) {
            revert InterchainApp__InterchainClientZeroAddress();
        }
        return IInterchainClientV1(client).getInterchainFee(
            dstChainId, _getExecutionService(), _getModules(), options, messageLen
        );
    }

    /// @dev Returns the configuration of the app for validating the received messages.
    function _getAppConfig() internal view virtual returns (bytes memory);

    /// @dev Returns the address of the Execution Service to use for sending messages.
    function _getExecutionService() internal view virtual returns (address);

    /// @dev Returns the latest Interchain Client. This is the Client that is used for sending messages.
    function _getLatestClient() internal view virtual returns (address);

    /// @dev Returns the list of modules to use for sending messages, as well as validating the received messages.
    function _getModules() internal view virtual returns (address[] memory);

    /// @dev Checks if the sender is allowed to send messages to this app.
    function _isAllowedSender(uint64 srcChainId, bytes32 sender) internal view virtual returns (bool);

    /// @dev Checks if the caller is an Interchain Client.
    /// Both latest and legacy Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view virtual returns (bool);
}
