// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AbstractICApp, InterchainTxDescriptor} from "./AbstractICApp.sol";

import {InterchainAppV1Events} from "../events/InterchainAppV1Events.sol";
import {IInterchainAppV1} from "../interfaces/IInterchainAppV1.sol";
import {AppConfigV1} from "../libs/AppConfig.sol";
import {OptionsV1} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract ICAppV1 is AbstractICApp, AccessControlEnumerable, InterchainAppV1Events, IInterchainAppV1 {
    using EnumerableSet for EnumerableSet.AddressSet;
    using TypeCasts for address;
    using TypeCasts for bytes32;

    /// @notice Role to manage the Interchain setup of the app.
    bytes32 public constant IC_GOVERNOR_ROLE = keccak256("IC_GOVERNOR_ROLE");

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

    /// @inheritdoc IInterchainAppV1
    function addInterchainClient(address client, bool updateLatest) external onlyRole(IC_GOVERNOR_ROLE) {
        _addClient(client, updateLatest);
    }

    /// @inheritdoc IInterchainAppV1
    function removeInterchainClient(address client) external onlyRole(IC_GOVERNOR_ROLE) {
        _removeClient(client);
    }

    /// @inheritdoc IInterchainAppV1
    function setLatestInterchainClient(address client) external onlyRole(IC_GOVERNOR_ROLE) {
        _setLatestClient(client);
    }

    /// @inheritdoc IInterchainAppV1
    function linkRemoteApp(uint256 chainId, bytes32 remoteApp) external onlyRole(IC_GOVERNOR_ROLE) {
        _linkRemoteApp(chainId, remoteApp);
    }

    /// @inheritdoc IInterchainAppV1
    function linkRemoteAppEVM(uint256 chainId, address remoteApp) external onlyRole(IC_GOVERNOR_ROLE) {
        _linkRemoteApp(chainId, remoteApp.addressToBytes32());
    }

    /// @inheritdoc IInterchainAppV1
    function addTrustedModule(address module) external onlyRole(IC_GOVERNOR_ROLE) {
        if (module == address(0)) {
            revert InterchainApp__ModuleZeroAddress();
        }
        bool added = _trustedModules.add(module);
        if (!added) {
            revert InterchainApp__ModuleAlreadyAdded(module);
        }
        emit TrustedModuleAdded(module);
    }

    /// @inheritdoc IInterchainAppV1
    function removeTrustedModule(address module) external onlyRole(IC_GOVERNOR_ROLE) {
        bool removed = _trustedModules.remove(module);
        if (!removed) {
            revert InterchainApp__ModuleNotAdded(module);
        }
        emit TrustedModuleRemoved(module);
    }

    /// @inheritdoc IInterchainAppV1
    function setAppConfigV1(AppConfigV1 memory appConfig) external onlyRole(IC_GOVERNOR_ROLE) {
        if (appConfig.requiredResponses == 0 || appConfig.optimisticPeriod == 0) {
            revert InterchainApp__InvalidAppConfig(appConfig.requiredResponses, appConfig.optimisticPeriod);
        }
        _appConfigV1 = appConfig;
        emit AppConfigV1Set(appConfig.requiredResponses, appConfig.optimisticPeriod);
    }

    /// @inheritdoc IInterchainAppV1
    function setExecutionService(address executionService) external onlyRole(IC_GOVERNOR_ROLE) {
        _executionService = executionService;
        emit ExecutionServiceSet(executionService);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainAppV1
    function getAppConfigV1() external view returns (AppConfigV1 memory) {
        return _appConfigV1;
    }

    /// @inheritdoc IInterchainAppV1
    function getExecutionService() external view returns (address) {
        return _executionService;
    }

    /// @inheritdoc IInterchainAppV1
    function getInterchainClients() external view returns (address[] memory) {
        return _interchainClients.values();
    }

    /// @inheritdoc IInterchainAppV1
    function getLatestInterchainClient() external view returns (address) {
        return _latestClient;
    }

    /// @inheritdoc IInterchainAppV1
    function getLinkedApp(uint256 chainId) external view returns (bytes32) {
        return _linkedApp[chainId];
    }

    /// @inheritdoc IInterchainAppV1
    function getLinkedAppEVM(uint256 chainId) external view returns (address linkedAppEVM) {
        bytes32 linkedApp = _linkedApp[chainId];
        linkedAppEVM = linkedApp.bytes32ToAddress();
        if (linkedAppEVM.addressToBytes32() != linkedApp) {
            revert InterchainApp__NotEVMLinkedApp(linkedApp);
        }
    }

    /// @inheritdoc IInterchainAppV1
    function getModules() external view returns (address[] memory) {
        return _trustedModules.values();
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Links the remote app to the current app.
    /// Will revert if the chainId is the same as the chainId of the local app.
    /// Note: Should be guarded with permissions check.
    function _linkRemoteApp(uint256 chainId, bytes32 remoteApp) internal {
        if (chainId == block.chainid) {
            revert InterchainApp__SameChainId(chainId);
        }
        if (remoteApp == 0) {
            revert InterchainApp__AppZeroAddress();
        }
        _linkedApp[chainId] = remoteApp;
        emit AppLinked(chainId, remoteApp);
    }

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

    // ════════════════════════════════════════════ INTERNAL: MESSAGING ════════════════════════════════════════════════

    /// @dev Thin wrapper around _sendInterchainMessage to send the message to the linked app.
    function _sendToLinkedApp(
        uint256 dstChainId,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        internal
        returns (InterchainTxDescriptor memory)
    {
        bytes memory encodedOptions = options.encodeOptionsV1();
        return _sendInterchainMessage(dstChainId, _linkedApp[dstChainId], messageFee, encodedOptions, message);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the fee to send a message to the linked app on the remote chain.
    function _getMessageFee(
        uint256 dstChainId,
        OptionsV1 memory options,
        bytes memory message
    )
        internal
        view
        returns (uint256)
    {
        bytes memory encodedOptions = options.encodeOptionsV1();
        return _getInterchainFee(dstChainId, encodedOptions, message);
    }

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
