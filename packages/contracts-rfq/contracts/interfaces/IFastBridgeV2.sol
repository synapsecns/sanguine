// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
    enum BridgeStatus {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED
    }

    struct BridgeTxDetails {
        BridgeStatus status;
        uint32 destChainId;
        uint56 proofBlockTimestamp;
        address proofRelayer;
    }

    struct BridgeRelay {
        uint48 blockNumber;
        uint48 blockTimestamp;
        address relayer;
    }

    /// @notice New params introduced in the FastBridgeV2.
    /// We are passing fields from the older BridgeParams struct outside of this struct
    /// for backwards compatibility.
    /// Note: quoteRelayer and quoteExclusivitySeconds are either both zero (indicating no exclusivity)
    /// or both non-zero (indicating exclusivity for the given period).
    /// Note: zapNative > 0 can NOT be used with destToken = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE (native token)
    /// @param quoteRelayer             Relayer that provided the quote for the transaction
    /// @param quoteExclusivitySeconds  Period of time the quote relayer is guaranteed exclusivity after user's deposit
    /// @param quoteId                  Unique quote identifier used for tracking the quote
    /// @param zapNative                ETH value to send to the recipient (if any)
    /// @param zapData                  Parameters for the Zap to the destination recipient (if any)
    struct BridgeParamsV2 {
        address quoteRelayer;
        int256 quoteExclusivitySeconds;
        bytes quoteId;
        uint256 zapNative;
        bytes zapData;
    }

    /// @notice Updated bridge transaction struct to include parameters introduced in FastBridgeV2.
    /// Note: only `exclusivityRelayer` can fill such a transaction until `exclusivityEndTime`.
    struct BridgeTransactionV2 {
        uint32 originChainId;
        uint32 destChainId;
        address originSender; // user (origin)
        address destRecipient; // user (dest)
        address originToken;
        address destToken;
        uint256 originAmount; // amount in on origin bridge less originFeeAmount
        uint256 destAmount;
        uint256 originFeeAmount;
        // Note: sendChainGas flag from V1 is deprecated
        uint256 deadline; // user specified deadline for destination relay
        uint256 nonce;
        address exclusivityRelayer;
        uint256 exclusivityEndTime;
        uint256 zapNative; // ETH value to send to the recipient (if any)
        bytes zapData; // data to pass for the Zap action (if any)
    }

    event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId);

    /// @notice Initiates bridge on origin chain to be relayed by off-chain relayer, with the ability
    /// to provide temporary exclusivity fill rights for the quote relayer.
    /// @param params   The parameters required to bridge
    /// @param paramsV2 The parameters for exclusivity fill rights (optional, can be left empty)
    function bridge(BridgeParams memory params, BridgeParamsV2 memory paramsV2) external payable;

    /// @notice Relays destination side of bridge transaction by off-chain relayer
    /// @param request The encoded bridge transaction to relay on destination chain
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function relay(bytes memory request, address relayer) external payable;

    /// @notice Provides proof on origin side that relayer provided funds on destination side of bridge transaction
    /// @param transactionId The transaction id associated with the encoded bridge transaction to prove
    /// @param destTxHash The destination tx hash proving bridge transaction was relayed
    /// @param relayer The address of the relaying entity which should have control of the origin funds when claimed
    function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) external;

    /// @notice Completes bridge transaction on origin chain by claiming originally deposited capital.
    /// @notice Can only send funds to the relayer address on the proof.
    /// @param request The encoded bridge transaction to claim on origin chain
    function claim(bytes memory request) external;
    /// @notice Checks if a transaction has been relayed
    /// @param transactionId The ID of the transaction to check
    /// @return True if the transaction has been relayed, false otherwise
    function bridgeRelays(bytes32 transactionId) external view returns (bool);

    /// @notice Returns the status of a bridge transaction
    /// @param transactionId The ID of the bridge transaction
    /// @return BridgeStatus Status of the bridge transaction
    function bridgeStatuses(bytes32 transactionId) external view returns (BridgeStatus);

    /// @notice Returns the timestamp and relayer of a bridge proof
    /// @param transactionId The ID of the bridge transaction
    /// @return timestamp The timestamp of the bridge proof
    /// @return relayer The relayer address of the bridge proof
    function bridgeProofs(bytes32 transactionId) external view returns (uint96 timestamp, address relayer);

    /// @notice Decodes bridge request into a bridge transaction V2 struct used by FastBridgeV2
    /// @param request The bridge request to decode
    function getBridgeTransactionV2(bytes memory request) external view returns (BridgeTransactionV2 memory);
}
