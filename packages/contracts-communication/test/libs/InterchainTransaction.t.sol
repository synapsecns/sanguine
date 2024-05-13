// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainTransaction,
    InterchainTransactionLibHarness,
    ICTxHeader
} from "../harnesses/InterchainTransactionLibHarness.sol";
import {VersionedPayloadLibHarness} from "../harnesses/VersionedPayloadLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainTransactionLibTest is Test {
    InterchainTransactionLibHarness public libHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    function setUp() public {
        libHarness = new InterchainTransactionLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
    }

    function test_constructLocalTransaction(
        uint64 srcChainId,
        address srcSender,
        uint64 dstChainId,
        bytes32 dstReceiver,
        uint64 dbNonce,
        bytes memory options,
        bytes memory message
    )
        public
    {
        vm.chainId(srcChainId);
        InterchainTransaction memory icTx =
            libHarness.constructLocalTransaction(srcSender, dstChainId, dstReceiver, dbNonce, options, message);
        assertEq(icTx.srcChainId, srcChainId, "!srcChainId");
        assertEq(icTx.srcSender, bytes32(uint256(uint160(srcSender))), "!srcSender");
        assertEq(icTx.dstChainId, dstChainId, "!dstChainId");
        assertEq(icTx.dstReceiver, dstReceiver, "!dstReceiver");
        assertEq(icTx.dbNonce, dbNonce, "!dbNonce");
        assertEq(icTx.options, options, "!options");
        assertEq(icTx.message, message, "!message");
    }

    function test_encodeTransaction_roundTrip(InterchainTransaction memory icTx) public view {
        bytes memory encoded = libHarness.encodeTransaction(icTx);
        InterchainTransaction memory decoded = libHarness.decodeTransaction(encoded);
        assertEq(decoded.srcChainId, icTx.srcChainId, "!srcChainId");
        assertEq(decoded.srcSender, icTx.srcSender, "!srcSender");
        assertEq(decoded.dstChainId, icTx.dstChainId, "!dstChainId");
        assertEq(decoded.dstReceiver, icTx.dstReceiver, "!dstReceiver");
        assertEq(decoded.dbNonce, icTx.dbNonce, "!dbNonce");
        assertEq(decoded.options, icTx.options, "!options");
        assertEq(decoded.message, icTx.message, "!message");
    }

    function test_payloadSize(InterchainTransaction memory icTx) public view {
        uint256 size = libHarness.payloadSize(icTx.options.length, icTx.message.length);
        uint256 expectedSize = payloadLibHarness.encodeVersionedPayload(0, libHarness.encodeTransaction(icTx)).length;
        assertEq(size, expectedSize);
    }

    function test_payloadSize_fuzzBytesOnly(bytes memory options, bytes memory message) public view {
        InterchainTransaction memory icTx;
        icTx.options = options;
        icTx.message = message;
        test_payloadSize(icTx);
    }

    function test_encodeTxHeader_roundTrip(uint64 srcChainId, uint64 dstChainId, uint64 dbNonce) public view {
        ICTxHeader header = libHarness.encodeTxHeader(srcChainId, dstChainId, dbNonce);
        (uint64 decodedSrcChainId, uint64 decodedDstChainId, uint64 decodedDbNonce) = libHarness.decodeTxHeader(header);
        assertEq(decodedSrcChainId, srcChainId, "!srcChainId");
        assertEq(decodedDstChainId, dstChainId, "!dstChainId");
        assertEq(decodedDbNonce, dbNonce, "!dbNonce");
    }
}
