// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {LegacyReceiver} from "./LegacyReceiver.sol";

import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @notice A simple app that sends a message to the remote PingPongApp, which will respond with a message back.
/// This app can be loaded with a native asset, which will be used to pay for the messages sent.
/// Note: we deal with uint256 chain IDs in this contract as the legacy MessageBus contract is using uint256 chain IDs.
contract LegacyPingPong is LegacyReceiver {
    uint256 internal constant MESSAGE_LENGTH = 32;
    uint256 internal constant DEFAULT_GAS_LIMIT = 500_000;
    uint256 public gasLimit;

    event GasLimitSet(uint256 gasLimit);
    event PingDisrupted(uint256 counter);
    event PingReceived(uint256 counter);
    event PingSent(uint256 counter);

    constructor(address owner_) LegacyReceiver(owner_) {
        _setGasLimit(DEFAULT_GAS_LIMIT);
    }

    /// @notice Enables the contract to accept native asset.
    receive() external payable {}

    /// @notice Allows the Owner to set the gas limit for the interchain messages.
    function setGasLimit(uint256 gasLimit_) external onlyOwner {
        _setGasLimit(gasLimit_);
    }

    /// @notice Allows the Owner to withdraw the native asset from the contract.
    function withdraw() external onlyOwner {
        Address.sendValue(payable(msg.sender), address(this).balance);
    }

    /// @notice Starts the ping-pong message exchange with the remote PingPongApp.
    function startPingPong(uint256 dstChainId, uint256 counter) external {
        _sendPingPongMessage({dstChainId: dstChainId, counter: counter, lowBalanceRevert: true});
    }

    /// @notice Returns the fee to send a single ping message to the remote PingPongApp.
    function getPingFee(uint256 dstChainId) public view returns (uint256) {
        return _getMessageFee({dstChainId: dstChainId, gasLimit: gasLimit, messageLen: MESSAGE_LENGTH});
    }

    /// @dev Handle the verified message.
    function _handleMessage(
        bytes32, // srcAddress
        uint256 srcChainId,
        bytes calldata message,
        address // executor
    )
        internal
        override
    {
        uint256 counter = abi.decode(message, (uint256));
        emit PingReceived(counter);
        if (counter > 0) {
            _sendPingPongMessage({dstChainId: srcChainId, counter: counter - 1, lowBalanceRevert: false});
        }
    }

    /// @dev Sends a message to the PingPongApp on the remote chain.
    /// If `counter > 0`, the remote app will respond with a message to this app, decrementing the counter.
    /// Once the counter reaches 0, the remote app will not respond.
    function _sendPingPongMessage(uint256 dstChainId, uint256 counter, bool lowBalanceRevert) internal {
        uint256 pingFee = getPingFee(dstChainId);
        if (address(this).balance < pingFee && !lowBalanceRevert) {
            emit PingDisrupted(counter);
            return;
        }
        bytes memory message = abi.encode(counter);
        _sendMessage({dstChainId: dstChainId, messageFee: pingFee, gasLimit: gasLimit, message: message});
        emit PingSent(counter);
    }

    /// @dev Sets the gas limit for the interchain messages.
    function _setGasLimit(uint256 gasLimit_) internal {
        gasLimit = gasLimit_;
        emit GasLimitSet(gasLimit_);
    }
}
