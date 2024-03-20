// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AbstractICApp} from "./AbstractICApp.sol";

import {IInterchainAppV1} from "../interfaces/IInterchainAppV1.sol";
import {AppConfigV1} from "../libs/AppConfig.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract ICAppV1 is AbstractICApp, AccessControlEnumerable, IInterchainAppV1 {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @dev Address of the latest Interchain Client, used for sending messages.
    address private _latestClient;

    /// @dev Required responses and optimistic period for the module responses.
    AppConfigV1 private _appConfigV1;
    /// @dev Address of the linked app deployed on the remote chain.
    mapping(uint256 chainId => bytes32 remoteApp) private _linkedApp;
    /// @dev Interchain Clients allowed to send messages to this app.
    EnumerableSet.AddressSet private _interchainClients;
    /// @dev Trusted Interchain modules.
    EnumerableSet.AddressSet private _trustedModules;
    /// @dev Execution Service to use for sending messages.
    address private _executionService;

    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Stores the address of the latest Interchain Client.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_setLatestClient` instead.
    /// - Should not emit any events: this is done in the calling function.
    function _storeLatestClient(address client) internal override {
        _latestClient = client;
    }

    /// @dev Toggle the state of the Interchain Client (allowed/disallowed to send messages to this app).
    /// - The client is checked to be in the opposite state before the change.
    /// - The exact storage location is up to the implementation.
    /// - Must NOT be called directly: use `_addClient` and `_removeClient` instead.
    /// - Should not emit any events: this is done in the calling functions.
    function _toggleClientState(address client, bool allowed) internal override {
        if (allowed) {
            _interchainClients.add(client);
        } else {
            _interchainClients.remove(client);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the configuration of the app for validating the received messages.
    function _getAppConfig() internal view override returns (bytes memory) {
        return _appConfigV1.encodeAppConfigV1();
    }

    /// @dev Returns the address of the Execution Service to use for sending messages.
    function _getExecutionService() internal view override returns (address) {
        return _executionService;
    }

    /// @dev Returns the latest Interchain Client. This is the Client that is used for sending messages.
    function _getLatestClient() internal view override returns (address) {
        return _latestClient;
    }

    /// @dev Returns the list of modules to use for sending messages, as well as validating the received messages.
    function _getModules() internal view override returns (address[] memory) {
        return _trustedModules.values();
    }

    /// @dev Checks if the sender is allowed to send messages to this app.
    function _isAllowedSender(uint256 srcChainId, bytes32 sender) internal view override returns (bool) {
        return _linkedApp[srcChainId] == sender;
    }

    /// @dev Checks if the caller is an Interchain Client.
    /// Both latest and legacy Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view override returns (bool) {
        return _interchainClients.contains(caller);
    }
}
