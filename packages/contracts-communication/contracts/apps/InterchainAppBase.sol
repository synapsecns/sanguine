// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {InterchainAppBaseEvents} from "../events/InterchainAppBaseEvents.sol";
import {IInterchainApp} from "../interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "../interfaces/IInterchainClientV1.sol";

import {AppConfigV1} from "../libs/AppConfig.sol";
import {OptionsV1} from "../libs/Options.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract InterchainAppBase is InterchainAppBaseEvents, IInterchainApp {
    using EnumerableSet for EnumerableSet.AddressSet;

    // TODO: naming, visibility
    address public interchain;

    /// @dev Required responses and optimistic period for the module responses.
    AppConfigV1 private _appConfig;
    /// @dev Address of the linked app deployed on the remote chain.
    mapping(uint256 chainId => bytes32 remoteApp) private _linkedApp;
    /// @dev Trusted Interchain modules.
    EnumerableSet.AddressSet private _trustedModules;
    /// @dev Execution Service to use for sending messages.
    address private _executionService;

    error InterchainApp__BalanceTooLow(uint256 actual, uint256 expected);
    error InterchainApp__CallerNotInterchainClient(address caller);
    error InterchainApp__InterchainClientNotSet();
    error InterchainApp__ModuleAlreadyAdded(address module);
    error InterchainApp__ModuleNotAdded(address module);
    error InterchainApp__ReceiverNotSet(uint256 chainId);
    error InterchainApp__SameChainId(uint256 chainId);
    error InterchainApp__SenderNotAllowed(uint256 srcChainId, bytes32 sender);

    /// @inheritdoc IInterchainApp
    function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes calldata message) external payable {
        if (msg.sender != interchain) revert InterchainApp__CallerNotInterchainClient(msg.sender);
        if (srcChainId == block.chainid) revert InterchainApp__SameChainId(srcChainId);
        if (!isAllowedSender(srcChainId, sender)) revert InterchainApp__SenderNotAllowed(srcChainId, sender);
        _receiveMessage(srcChainId, sender, nonce, message);
        // Note: application may elect to emit an event in `_receiveMessage`, so we don't emit a generic event here.
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainApp
    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {
        // Note: the getters for app config and modules could be overridden in the derived contracts.
        appConfig = getAppConfig().encodeAppConfigV1();
        modules = getReceivingModules();
    }

    /// @notice Returns the app config for receiving messages.
    /// @dev Could be overridden in the derived contracts.
    function getAppConfig() public view virtual returns (AppConfigV1 memory) {
        return _appConfig;
    }

    /// @notice Returns the address of the Execution Service to use for sending messages.
    /// @dev Could be overridden in the derived contracts.
    function getExecutionService() public view virtual returns (address) {
        return _executionService;
    }

    /// @notice Returns the linked app address (as bytes32) for the given chain ID.
    /// @dev Could be overridden in the derived contracts.
    function getLinkedApp(uint256 chainId) public view virtual returns (bytes32) {
        return _linkedApp[chainId];
    }

    /// @notice Returns the list of modules used for sending messages.
    /// @dev Could be overridden in the derived contracts.
    function getSendingModules() public view virtual returns (address[] memory) {
        return _trustedModules.values();
    }

    /// @notice Returns the list of modules used for receiving messages.
    /// @dev Could be overridden in the derived contracts.
    function getReceivingModules() public view virtual returns (address[] memory) {
        return _trustedModules.values();
    }

    /// @notice Checks whether the sender is allowed to send messages to the current app.
    /// @dev Could be overridden in the derived contracts.
    function isAllowedSender(uint256 srcChainId, bytes32 sender) public view virtual returns (bool) {
        return _linkedApp[srcChainId] == sender;
    }

    // ═══════════════════════════════════════════ INTERNAL: MANAGEMENT ════════════════════════════════════════════════

    /// @dev Links the remote app to the current app.
    /// Will revert if the chainId is the same as the chainId of the local app.
    /// Note: Should be guarded with permissions check.
    function _linkRemoteApp(uint256 chainId, bytes32 remoteApp) internal {
        if (chainId == block.chainid) revert InterchainApp__SameChainId(chainId);
        _linkedApp[chainId] = remoteApp;
        emit AppLinked(chainId, remoteApp);
    }

    /// @dev Thin wrapper around _linkRemoteApp to accept EVM address as a parameter.
    function _linkRemoteAppEVM(uint256 chainId, address remoteApp) internal {
        _linkRemoteApp(chainId, TypeCasts.addressToBytes32(remoteApp));
    }

    /// @dev Adds the module to the trusted modules set.
    /// Will revert if the module is already added.
    /// Note: Should be guarded with permissions check.
    function _addTrustedModule(address module) internal {
        bool added = _trustedModules.add(module);
        if (!added) revert InterchainApp__ModuleAlreadyAdded(module);
        emit TrustedModuleAdded(module);
    }

    /// @dev Removes the module from the trusted modules set.
    /// Will revert if the module is not added.
    /// Note: Should be guarded with permissions check.
    function _removeTrustedModule(address module) internal {
        bool removed = _trustedModules.remove(module);
        if (!removed) revert InterchainApp__ModuleNotAdded(module);
        emit TrustedModuleRemoved(module);
    }

    /// @dev Sets the app config:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    /// Note: Should be guarded with permissions check.
    function _setAppConfigV1(AppConfigV1 memory appConfig) internal {
        _appConfig = appConfig;
        emit AppConfigV1Set(appConfig.requiredResponses, appConfig.optimisticPeriod);
    }

    /// @dev Sets the execution service address.
    /// Note: Should be guarded with permissions check.
    function _setExecutionService(address executionService) internal {
        _executionService = executionService;
        emit ExecutionServiceSet(executionService);
    }

    /// @dev Sets the interchain client address.
    /// Note: Should be guarded with permissions check.
    function _setInterchainClient(address interchain_) internal {
        interchain = interchain_;
        emit InterchainClientSet(interchain_);
    }

    // ════════════════════════════════════════════ INTERNAL: MESSAGING ════════════════════════════════════════════════

    /// @dev Thin wrapper around _sendInterchainMessage to send the message to the linked app.
    function _sendInterchainMessage(
        uint256 dstChainId,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        internal
    {
        _sendInterchainMessage(dstChainId, getLinkedApp(dstChainId), messageFee, options, message);
    }

    /// @dev Thin wrapper around _sendInterchainMessage to accept EVM address as a parameter.
    function _sendInterchainMessageEVM(
        uint256 dstChainId,
        address receiver,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        internal
    {
        _sendInterchainMessage(dstChainId, TypeCasts.addressToBytes32(receiver), messageFee, options, message);
    }

    /// @dev Performs necessary checks and sends an interchain message.
    function _sendInterchainMessage(
        uint256 dstChainId,
        bytes32 receiver,
        uint256 messageFee,
        OptionsV1 memory options,
        bytes memory message
    )
        internal
    {
        address cachedInterchain = interchain;
        if (cachedInterchain == address(0)) revert InterchainApp__InterchainClientNotSet();
        if (dstChainId == block.chainid) revert InterchainApp__SameChainId(dstChainId);
        if (receiver == 0) revert InterchainApp__ReceiverNotSet(dstChainId);
        if (address(this).balance < messageFee) revert InterchainApp__BalanceTooLow(address(this).balance, messageFee);
        IInterchainClientV1(cachedInterchain).interchainSend{value: messageFee}(
            dstChainId, receiver, getExecutionService(), getSendingModules(), options.encodeOptionsV1(), message
        );
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint64 nonce,
        bytes calldata message
    )
        internal
        virtual;
}
