// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IFastBridge} from "./IFastBridge.sol";

interface IFastBridgeV2 is IFastBridge {
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
    /// @param quoteRelayer             Relayer that provided the quote for the transaction
    /// @param quoteExclusivitySeconds  Period of time the quote relayer is guaranteed exclusivity after user's deposit
    /// @param quoteId                  Unique quote identifier used for tracking the quote
    struct BridgeParamsV2 {
        address quoteRelayer;
        uint256 quoteExclusivitySeconds;
        bytes quoteId;
    }

    /// @notice Updated bridge transaction struct to include parameters introduced in FastBridgeV2.
    /// Note: only `exclusivityRelayer` can fill such a transaction until `exclusivityEndTime`.
    /// TODO: consider changing the encoding scheme to prevent spending extra gas on decoding.
    struct BridgeTransactionV2 {
        BridgeTransaction txV1;
        address exclusivityRelayer;
        uint256 exclusivityEndTime;
    }

    /// @notice Initiates bridge on origin chain to be relayed by off-chain relayer, with the ability
    /// to provide temporary exclusivity fill rights for the quote relayer.
    /// @param params   The parameters required to bridge
    /// @param paramsV2 The parameters for exclusivity fill rights (optional, could be left empty)
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
}
