// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IInterchainApp} from "../interfaces/IInterchainApp.sol";

abstract contract AbstractICApp is IInterchainApp {
    error InterchainApp__NotInterchainClient(address account);
    error InterchainApp__SameChainId(uint256 chainId);
    error InterchainApp__SenderNotAllowed(uint256 srcChainId, bytes32 sender);

    /// @inheritdoc IInterchainApp
    function appReceive(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
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
            revert InterchainApp__SameChainId(srcChainId);
        }
        if (!_isAllowedSender(srcChainId, sender)) {
            revert InterchainApp__SenderNotAllowed(srcChainId, sender);
        }
        _receiveMessage(srcChainId, sender, dbNonce, entryIndex, message);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        internal
        virtual;

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if the sender is allowed to send messages to this app.
    function _isAllowedSender(uint256 srcChainId, bytes32 sender) internal view virtual returns (bool);

    /// @dev Checks if the caller is an Interchain Client. Only Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view virtual returns (bool);
}
