// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IMessageBus {
    error MessageBus__NotEVMReceiver(bytes32 receiver);

    /// @notice Sends a message to a receiving contract address on another chain.
    /// Sender must make sure that the message is unique and not a duplicate message.
    /// @dev Legacy MessageBus only supports V1 of the options format, which specifies only the gas limit.
    /// @param receiver     The bytes32 address of the destination contract to be called
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param message      The arbitrary payload to pass to the destination chain receiver
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
    function sendMessage(
        bytes32 receiver,
        uint256 dstChainId,
        bytes memory message,
        bytes memory options
    )
        external
        payable;

    /// @notice Allows the Interchain Governor to set the message length in bytes to be used for fee estimation.
    /// This does not affect the sendMessage function, but only the fee estimation.
    function setMessageLengthEstimate(uint256 length) external;

    /// @notice Returns the preset message length in bytes used for fee estimation.
    function messageLengthEstimate() external view returns (uint256);

    /// @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the gas limit.
    /// @dev This function is using the preset message length to estimate the gas fee. This should cover most cases,
    /// if the message length is lower than the preset value. For more accurate fee estimation, use estimateFeeExact.
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
    function estimateFee(uint256 dstChainId, bytes memory options) external view returns (uint256);

    /// @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the message length
    /// and the gas limit.
    /// @param dstChainId   The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
    /// @param messageLen   The length of the message to be sent in bytes
    /// @param options      Versioned struct used to instruct relayer on how to proceed with gas limits
    function estimateFeeExact(
        uint256 dstChainId,
        uint256 messageLen,
        bytes memory options
    )
        external
        view
        returns (uint256);
}
