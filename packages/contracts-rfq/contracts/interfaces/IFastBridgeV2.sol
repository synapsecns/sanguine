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

    error FastBridge__MsgValueIncorrect();

    /// @notice Returns whether transaction has been relayed on the destination chain.
    /// @dev This function is added for backwards compatibility with FastBridgeV1.
    function bridgeRelays(bytes32 transactionId) external view returns (bool);

    /// @notice Returns the address of the Relayer who completed the transaction.
    /// Note: returns address(0) if the transaction has not been relayed.
    function getTransactionRelayer(bytes32 transactionId) external view returns (address);
}
