// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OwnableApp} from "./OwnableApp.sol";

import {OptionsV1} from "../libs/Options.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @notice A simple app that sends a message to the remote PingPongApp, which will respond with a message back.
/// This app can be loaded with a native asset, which will be used to pay for the messages sent.
contract PingPongApp is OwnableApp {
    uint256 internal constant DEFAULT_GAS_LIMIT = 500_000;
    uint256 public gasLimit;

    event GasLimitSet(uint256 gasLimit);
    event PingDisrupted(uint256 counter);
    event PingReceived(uint256 counter);
    event PingSent(uint256 counter);

    error PingPongApp__LowBalance(uint256 required);

    constructor(address owner_) OwnableApp(owner_) {
        _setGasLimit(DEFAULT_GAS_LIMIT);
    }

    /// @notice Enables the contract to accept native asset.
    receive() external payable {}

    /// @notice Allows the owner to set the gas limit for the interchain messages.
    function setGasLimit(uint256 gasLimit_) external onlyOwner {
        _setGasLimit(gasLimit_);
    }

    /// @notice Allows the owner to withdraw the native asset from the contract.
    function withdraw() external onlyOwner {
        Address.sendValue(payable(msg.sender), address(this).balance);
    }

    /// @notice Starts the ping-pong message exchange with the remote PingPongApp.
    function startPingPong(uint256 dstChainId, uint256 counter) external {
        // Revert if the balance is lower than the message fee.
        _sendPingPongMessage(dstChainId, counter, true);
    }

    /// @notice Returns the fee to send a single ping message to the remote PingPongApp.
    function getPingFee(uint256 dstChainId) external view returns (uint256) {
        OptionsV1 memory options = OptionsV1({gasLimit: gasLimit, gasAirdrop: 0});
        bytes memory message = abi.encode(uint256(0));
        return _getInterchainFee(dstChainId, options, message);
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32, // sender
        uint256, // dbNonce
        bytes calldata message
    )
        internal
        override
    {
        uint256 counter = abi.decode(message, (uint256));
        emit PingReceived(counter);
        if (counter > 0) {
            // Don't revert if the balance is low, just stop sending messages.
            _sendPingPongMessage({dstChainId: srcChainId, counter: counter - 1, lowBalanceRevert: false});
        }
    }

    /// @dev Sends a message to the PingPongApp on the remote chain.
    /// If `counter > 0`, the remote app will respond with a message to this app, decrementing the counter.
    /// Once the counter reaches 0, the remote app will not respond.
    function _sendPingPongMessage(uint256 dstChainId, uint256 counter, bool lowBalanceRevert) internal {
        OptionsV1 memory options = OptionsV1({gasLimit: gasLimit, gasAirdrop: 0});
        bytes memory message = abi.encode(counter);
        uint256 messageFee = _getInterchainFee(dstChainId, options, message);
        if (address(this).balance < messageFee) {
            if (lowBalanceRevert) {
                revert PingPongApp__LowBalance({required: messageFee});
            } else {
                emit PingDisrupted(counter);
                return;
            }
        }
        _sendInterchainMessage(dstChainId, messageFee, options, message);
        emit PingSent(counter);
    }

    /// @dev Sets the gas limit for the interchain messages.
    function _setGasLimit(uint256 gasLimit_) internal {
        gasLimit = gasLimit_;
        emit GasLimitSet(gasLimit_);
    }
}
