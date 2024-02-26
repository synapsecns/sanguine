// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainTransaction, InterchainTransactionLibHarness
} from "../harnesses/InterchainTransactionLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
contract InterchainTransactionLibTest is Test {
    InterchainTransactionLibHarness public libHarness;

    function setUp() public {
        libHarness = new InterchainTransactionLibHarness();
    }

    function test_constructLocalTransaction(
        uint64 srcChainId,
        address srcSender,
        uint256 dstChainId,
        bytes32 dstReceiver,
        uint256 dbNonce,
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

    function test_encodeRoundtrip(InterchainTransaction memory icTx) public {
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

    function test_transactionId(InterchainTransaction memory icTx) public {
        bytes32 expected = keccak256(abi.encode(icTx));
        assertEq(libHarness.transactionId(icTx), expected);
    }
}
