// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {BridgeTransactionV2Lib, IFastBridgeV2} from "../../contracts/libs/BridgeTransactionV2.sol";

contract BridgeTransactionV2Harness {
    function encodeV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public pure returns (bytes memory) {
        return BridgeTransactionV2Lib.encodeV2(bridgeTx);
    }

    function decodeV2(bytes memory encodedTx) public pure returns (IFastBridgeV2.BridgeTransactionV2 memory) {
        return BridgeTransactionV2Lib.decodeV2(encodedTx);
    }

    function encode(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public pure returns (bytes memory) {
        return BridgeTransactionV2Lib.encode(bridgeTx);
    }

    function decode(bytes calldata encodedTx) public pure returns (IFastBridgeV2.BridgeTransactionV2 memory) {
        return BridgeTransactionV2Lib.decode(encodedTx);
    }

    function version(bytes calldata encodedTx) public pure returns (uint16) {
        return BridgeTransactionV2Lib.version(encodedTx);
    }

    function originChainId(bytes calldata encodedTx) public pure returns (uint32) {
        return BridgeTransactionV2Lib.originChainId(encodedTx);
    }

    function destChainId(bytes calldata encodedTx) public pure returns (uint32) {
        return BridgeTransactionV2Lib.destChainId(encodedTx);
    }

    function originSender(bytes calldata encodedTx) public pure returns (address) {
        return BridgeTransactionV2Lib.originSender(encodedTx);
    }

    function destRecipient(bytes calldata encodedTx) public pure returns (address) {
        return BridgeTransactionV2Lib.destRecipient(encodedTx);
    }

    function originToken(bytes calldata encodedTx) public pure returns (address) {
        return BridgeTransactionV2Lib.originToken(encodedTx);
    }

    function destToken(bytes calldata encodedTx) public pure returns (address) {
        return BridgeTransactionV2Lib.destToken(encodedTx);
    }

    function originAmount(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.originAmount(encodedTx);
    }

    function destAmount(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.destAmount(encodedTx);
    }

    function originFeeAmount(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.originFeeAmount(encodedTx);
    }

    function callValue(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.callValue(encodedTx);
    }

    function deadline(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.deadline(encodedTx);
    }

    function nonce(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.nonce(encodedTx);
    }

    function exclusivityRelayer(bytes calldata encodedTx) public pure returns (address) {
        return BridgeTransactionV2Lib.exclusivityRelayer(encodedTx);
    }

    function exclusivityEndTime(bytes calldata encodedTx) public pure returns (uint256) {
        return BridgeTransactionV2Lib.exclusivityEndTime(encodedTx);
    }

    function callParams(bytes calldata encodedTx) public pure returns (bytes calldata) {
        return BridgeTransactionV2Lib.callParams(encodedTx);
    }
}
