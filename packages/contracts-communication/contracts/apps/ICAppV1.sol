// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AbstractICApp} from "./AbstractICApp.sol";

import {IInterchainAppV1} from "../interfaces/IInterchainAppV1.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";

abstract contract ICAppV1 is AbstractICApp, AccessControlEnumerable, IInterchainAppV1 {
    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Stores the address of the latest Interchain Client.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_setLatestClient` instead.
    /// - Should not emit any events: this is done in the calling function.
    function _storeLatestClient(address client) internal override {
        // TODO: implement
    }

    /// @dev Toggle the state of the Interchain Client (allowed/disallowed to send messages to this app).
    /// - The client is checked to be in the opposite state before the change.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_addClient` and `_removeClient` instead.
    /// - Should not emit any events: this is done in the calling functions.
    function _toggleClientState(address client, bool allowed) internal override {
        // TODO: implement
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the configuration of the app for validating the received messages.
    function _getAppConfig() internal view override returns (bytes memory) {
        // TODO: implement
    }

    /// @dev Returns the address of the Execution Service to use for sending messages.
    function _getExecutionService() internal view override returns (address) {
        // TODO: implement
    }

    /// @dev Returns the latest Interchain Client. This is the Client that is used for sending messages.
    function _getLatestClient() internal view override returns (address) {
        // TODO: implement
    }

    /// @dev Returns the list of modules to use for sending messages, as well as validating the received messages.
    function _getModules() internal view override returns (address[] memory) {
        // TODO: implement
    }

    /// @dev Checks if the sender is allowed to send messages to this app.
    function _isAllowedSender(uint256 srcChainId, bytes32 sender) internal view override returns (bool) {
        // TODO: implement
    }

    /// @dev Checks if the caller is an Interchain Client.
    /// Both latest and legacy Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view override returns (bool) {
        // TODO: implement
    }
}
