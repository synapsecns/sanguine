// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AbstractICApp, InterchainTxDescriptor} from "./AbstractICApp.sol";

import {InterchainAppV1Events} from "../events/InterchainAppV1Events.sol";
import {IInterchainAppV1} from "../interfaces/IInterchainAppV1.sol";
import {AppConfigV1, APP_CONFIG_GUARD_DISABLED} from "../libs/AppConfig.sol";
import {OptionsV1} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

abstract contract ICAppV1 is AbstractICApp, AccessControlEnumerable, InterchainAppV1Events, IInterchainAppV1 {
    using EnumerableSet for EnumerableSet.AddressSet;
    using TypeCasts for address;
    using TypeCasts for bytes32;

    /// @notice Role to manage the Interchain setup of the app.
    bytes32 public constant IC_GOVERNOR_ROLE = keccak256("IC_GOVERNOR_ROLE");

    /// @dev Address of the latest Interchain Client, used for sending messages.
    /// Note: packed in a single storage slot with the `_requiredResponses` and `_optimisticPeriod`.
    address private _latestClient;
    /// @dev Required responses and optimistic period for the module responses.
    uint16 private _requiredResponses;
    uint48 private _optimisticPeriod;

    /// @dev Address of the linked app deployed on the remote chain.
    mapping(uint64 chainId => bytes32 remoteApp) private _linkedApp;
    /// @dev Interchain Clients allowed to send messages to this app.
    EnumerableSet.AddressSet private _interchainClients;
    /// @dev Trusted Interchain modules.
    EnumerableSet.AddressSet private _trustedModules;
    /// @dev Execution Service to use for sending messages.
    address private _executionService;

    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @notice Allows the governor to add the interchain client to the allowed clients set,
    /// and optionally set the latest client to this one.
    /// Note: only the allowed clients can send messages to this app.
    /// Note: the latest client is used for sending messages from this app.
    /// @param client       The address of the interchain client to add.
    /// @param updateLatest Whether to set the latest client to this one.
    function addInterchainClient(address client, bool updateLatest) external onlyRole(IC_GOVERNOR_ROLE) {
        _addClient(client, updateLatest);
    }

    /// @notice Allows the governor to remove the interchain client from the allowed clients set.
    /// If the client is the latest client, the latest client is set to the zero address.
    /// @param client       The address of the interchain client to remove.
    function removeInterchainClient(address client) external onlyRole(IC_GOVERNOR_ROLE) {
        _removeClient(client);
    }

    /// @notice Allows the governor to set the address of the latest interchain client.
    /// @dev The new latest client must be an allowed client or zero address.
    /// Setting the client to zero address effectively pauses the app ability to send messages,
    /// while allowing to receive them.
    /// @param client       The address of the latest interchain client.
    function setLatestInterchainClient(address client) external onlyRole(IC_GOVERNOR_ROLE) {
        _setLatestClient(client);
    }

    /// @notice Allows the governor to link the remote app for the given chain ID.
    /// - This address will be used as the receiver for the messages sent from this chain.
    /// - This address will be the only trusted sender for the messages sent to this chain.
    /// @param chainId      The remote chain ID.
    /// @param remoteApp    The address of the remote app to link.
    function linkRemoteApp(uint64 chainId, bytes32 remoteApp) external onlyRole(IC_GOVERNOR_ROLE) {
        _linkRemoteApp(chainId, remoteApp);
    }

    /// @notice Thin wrapper for `linkRemoteApp` to accept EVM address as a parameter.
    function linkRemoteAppEVM(uint64 chainId, address remoteApp) external onlyRole(IC_GOVERNOR_ROLE) {
        _linkRemoteApp(chainId, remoteApp.addressToBytes32());
    }

    /// @notice Allows the governor to add the module to the trusted modules set.
    /// - This set of modules will be used to verify both sent and received messages.
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

    /// @notice Allows the governor to remove the module from the trusted modules set.
    function removeTrustedModule(address module) external onlyRole(IC_GOVERNOR_ROLE) {
        bool removed = _trustedModules.remove(module);
        if (!removed) {
            revert InterchainApp__ModuleNotAdded(module);
        }
        emit TrustedModuleRemoved(module);
    }

    /// @notice Allows the governor to set the app config for the current app. App config includes:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    function setAppConfigV1(uint256 requiredResponses, uint256 optimisticPeriod) external onlyRole(IC_GOVERNOR_ROLE) {
        if (requiredResponses == 0) {
            revert InterchainApp__AppConfigInvalid(requiredResponses, optimisticPeriod);
        }
        _requiredResponses = SafeCast.toUint16(requiredResponses);
        _optimisticPeriod = SafeCast.toUint48(optimisticPeriod);
        emit AppConfigV1Set(requiredResponses, optimisticPeriod);
    }

    /// @notice Allows the governor to set the address of the Execution Service.
    /// This address will be used to request execution of the messages sent from this chain,
    /// by supplying the Service's execution fee.
    function setExecutionService(address executionService) external onlyRole(IC_GOVERNOR_ROLE) {
        _executionService = executionService;
        emit ExecutionServiceSet(executionService);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @notice Returns the app config for the current app:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    /// - guardFlag: the flag indicating the guard type (0 - none, 1 - client's default, 2 - custom)
    /// - guard: the address of the guard contract (if the guardFlag is set to 2)
    function getAppConfigV1() public view returns (AppConfigV1 memory) {
        (uint8 guardFlag, address guard) = _getGuardConfig();
        return AppConfigV1({
            requiredResponses: _requiredResponses,
            optimisticPeriod: _optimisticPeriod,
            guardFlag: guardFlag,
            guard: guard
        });
    }

    /// @notice Returns the address of the Execution Service used by this app for sending messages.
    // solhint-disable-next-line ordering
    function getExecutionService() external view returns (address) {
        return _executionService;
    }

    /// @notice Returns the list of Interchain Clients allowed to send messages to this app.
    function getInterchainClients() external view returns (address[] memory) {
        return _interchainClients.values();
    }

    /// @notice Returns the address of the latest interchain client.
    /// This address is used for sending messages from this app.
    function getLatestInterchainClient() external view returns (address) {
        return _latestClient;
    }

    /// @notice Returns the linked app address (as bytes32) for the given chain ID.
    function getLinkedApp(uint64 chainId) external view returns (bytes32) {
        return _linkedApp[chainId];
    }

    /// @notice Thin wrapper for `getLinkedApp` to return the linked app address as EVM address.
    /// @dev Will revert if the linked app address is not an EVM address.
    function getLinkedAppEVM(uint64 chainId) external view returns (address linkedAppEVM) {
        bytes32 linkedApp = _linkedApp[chainId];
        linkedAppEVM = linkedApp.bytes32ToAddress();
        if (linkedAppEVM.addressToBytes32() != linkedApp) {
            revert InterchainApp__LinkedAppNotEVM(linkedApp);
        }
    }

    /// @notice Returns the list of Interchain Modules trusted by this app.
    /// This set of modules will be used to verify both sent and received messages.
    function getModules() external view returns (address[] memory) {
        return _trustedModules.values();
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Links the remote app to the current app.
    /// Will revert if the chainId is the same as the chainId of the local app.
    /// Note: Should be guarded with permissions check.
    function _linkRemoteApp(uint64 chainId, bytes32 remoteApp) internal {
        if (chainId == block.chainid) {
            revert InterchainApp__ChainIdNotRemote(chainId);
        }
        if (remoteApp == 0) {
            revert InterchainApp__RemoteAppZeroAddress();
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
        uint64 dstChainId,
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
        uint64 dstChainId,
        OptionsV1 memory options,
        uint256 messageLen
    )
        internal
        view
        returns (uint256)
    {
        bytes memory encodedOptions = options.encodeOptionsV1();
        return _getInterchainFee(dstChainId, encodedOptions, messageLen);
    }

    /// @dev Returns the configuration of the app for validating the received messages.
    function _getAppConfig() internal view override returns (bytes memory) {
        return getAppConfigV1().encodeAppConfigV1();
    }

    /// @dev Returns the guard flag and address in the app config.
    /// By default, the ICApp does not opt in for any guard, but it can be overridden in the derived contracts.
    function _getGuardConfig() internal view virtual returns (uint8 guardFlag, address guard) {
        return (APP_CONFIG_GUARD_DISABLED, address(0));
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
    function _isAllowedSender(uint64 srcChainId, bytes32 sender) internal view override returns (bool) {
        return _linkedApp[srcChainId] == sender;
    }

    /// @dev Checks if the caller is an Interchain Client.
    /// Both latest and legacy Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view override returns (bool) {
        return _interchainClients.contains(caller);
    }
}
