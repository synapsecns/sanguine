// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
    enum BridgeStatusV2 {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED
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
    error FastBridge__TransactionRelayed();
    error FastBridge__ZeroAddress();

    /// @notice Returns whether transaction has been relayed on the destination chain.
    /// @dev This function is added for backwards compatibility with FastBridgeV1.
    function bridgeRelays(bytes32 transactionId) external view returns (bool);

    /// @notice Returns the address of the Relayer who completed the transaction.
    /// Note: returns address(0) if the transaction has not been relayed.
    function getDestinationRelayer(bytes32 transactionId) external view returns (address);
}
