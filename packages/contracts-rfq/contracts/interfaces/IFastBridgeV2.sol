// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
    /// @dev This is an extension of BridgeStatus enum, which adds new statuses for FastBridgeV2.
    /// The values from 0 to 4 are the same as BridgeStatus.
    /// Disputed: the transaction has been disputed by the Guard, but not yet resolved.
    /// When the transaction is resolved by using an external source of truth, its status will be updated to either:
    /// - Transaction NOT completed on destination: REQUESTED, allowing other Relayers to prove, or the user
    /// to refund if the prove deadline is over.
    /// - Transaction completed on destination, but by another Relayer: RELAYER_PROVED, allowing that other Relayer
    /// to later claim the origin funds.
    /// - Transaction completed on destination, by the original Relayer: RELAYER_CLAIMED, allowing the original Relayer
    /// to claim the origin funds.
    enum BridgeStatusV2 {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED,
        DISPUTED
    }

    error FastBridge__AmountIncorrect();
    error FastBridge__ChainIncorrect();
    error FastBridge__DeadlineExceeded();
    error FastBridge__DeadlineNotExceeded();
    error FastBridge__DeadlineTooShort();
    error FastBridge__DisputePeriodNotPassed();
    error FastBridge__DisputePeriodPassed();
    error FastBridge__MsgValueIncorrect();
    error FastBridge__SenderIncorrect();
    error FastBridge__StatusIncorrect();
    error FastBridge__TransactionNotDisputed();
    error FastBridge__TransactionRelayed();
    error FastBridge__ZeroAddress();

    /// @notice Allows the Guard to pass the information about transaction's destination Relayer
    /// in order to resolve the disputed transaction.
    /// @dev destRelayer will be set to zero for transactions that were not completed on the destination chain.
    /// @param transactionId    The transaction id associated with the encoded bridge transaction to dispute.
    /// @param destRelayer      The address of the Relayer who completed the transaction on the destination chain.
    function resolve(bytes32 transactionId, address destRelayer) external;

    /// @notice Returns whether transaction has been relayed on the destination chain.
    /// @dev This function is added for backwards compatibility with FastBridgeV1.
    function bridgeRelays(bytes32 transactionId) external view returns (bool);

    /// @notice Returns the address of the Relayer who completed the transaction.
    /// Note: returns address(0) if the transaction has not been relayed.
    function getDestinationRelayer(bytes32 transactionId) external view returns (address);
}
