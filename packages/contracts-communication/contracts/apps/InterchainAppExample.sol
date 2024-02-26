// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainAppBase, AppConfigV1, OptionsV1} from "./InterchainAppBase.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract InterchainAppExample is InterchainAppBase, Ownable {
    event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes message);
    event MessageSent(uint256 dstChainId, uint256 dbNonce, bytes32 transactionId);

    constructor(address owner_) Ownable(owner_) {}

    /// @notice Allows the owner to link the remote app for the given chain ID.
    /// - This address will be used as the receiver for the messages sent from this chain.
    /// - This address will be the only trusted sender for the messages sent to this chain.
    function linkRemoteApp(uint256 chainId, bytes32 remoteApp) external onlyOwner {
        _linkRemoteApp(chainId, remoteApp);
    }

    /// @notice This wrapper for `linkRemoteApp` to accept EVM address as a parameter.
    function linkRemoteAppEVM(uint256 chainId, address remoteApp) external onlyOwner {
        _linkRemoteAppEVM(chainId, remoteApp);
    }

    /// @notice Allows the owner to add the module to the trusted modules set.
    /// - This set of modules will be used to verify message sent from this chain.
    /// - This set of modules will be used to verify message sent to this chain.
    function addTrustedModule(address module) external onlyOwner {
        _addTrustedModule(module);
    }

    /// @notice Allows the owner to remove the module from the trusted modules set.
    function removeTrustedModule(address module) external onlyOwner {
        _removeTrustedModule(module);
    }

    /// @notice Allows the owner to set the app config for the current app. App config includes:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    function setAppConfigV1(AppConfigV1 memory appConfig) external onlyOwner {
        _setAppConfigV1(appConfig);
    }

    /// @notice Allows the owner to set the address of the Execution Service.
    /// This address will be used to request execution of the messages sent from this chain,
    /// by supplying the Service's execution fee.
    function setExecutionService(address executionService) external onlyOwner {
        _setExecutionService(executionService);
    }

    /// @notice Allows the owner to set the address of the InterchainClient contract.
    function setInterchainClient(address interchain_) external onlyOwner {
        _setInterchainClient(interchain_);
    }

    /// @notice Sends a basic message to the destination chain.
    function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes calldata message) external payable {
        (bytes32 transactionId, uint256 dbNonce) = _sendInterchainMessage({
            dstChainId: dstChainId,
            messageFee: msg.value,
            options: OptionsV1({gasLimit: gasLimit, gasAirdrop: 0}),
            message: message
        });
        emit MessageSent(dstChainId, dbNonce, transactionId);
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        bytes calldata message
    )
        internal
        override
    {
        emit MessageReceived(srcChainId, sender, dbNonce, message);
    }
}
