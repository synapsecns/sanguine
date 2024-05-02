// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../contracts/libs/Options.sol";

import {InterchainClientV1BaseTest, InterchainTransaction} from "./InterchainClientV1.Base.t.sol";
import {InterchainTransactionLibHarness} from "./harnesses/InterchainTransactionLibHarness.sol";
import {VersionedPayloadLibHarness} from "./harnesses/VersionedPayloadLibHarness.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainClientV1GenericViewsTest is InterchainClientV1BaseTest {
    function setUp() public override {
        super.setUp();
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
    }

    function test_getLinkedClient_chainIdKnown() public view {
        assertEq(icClient.getLinkedClient(REMOTE_CHAIN_ID), MOCK_REMOTE_CLIENT);
    }

    function test_getLinkedClient_chainIdUnknown() public view {
        assertEq(icClient.getLinkedClient(UNKNOWN_CHAIN_ID), 0);
    }

    function test_getLinkedClient_revert_chainIdLocal() public {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        icClient.getLinkedClient(LOCAL_CHAIN_ID);
    }

    function test_getLinkedClientEVM_chainIdKnown() public {
        setLinkedClient(REMOTE_CHAIN_ID, mockRemoteClientEVMBytes32);
        assertEq(icClient.getLinkedClientEVM(REMOTE_CHAIN_ID), mockRemoteClientEVM);
    }

    function test_getLinkedClientEVM_chainIdUnknown() public view {
        assertEq(icClient.getLinkedClientEVM(UNKNOWN_CHAIN_ID), address(0));
    }

    function test_getLinkedClientEVM_revert_chainIdLocal() public {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        icClient.getLinkedClientEVM(LOCAL_CHAIN_ID);
    }

    function test_getLinkedClientEVM_revert_clientNotEVM() public {
        expectRevertNotEVMClient(MOCK_REMOTE_CLIENT);
        icClient.getLinkedClientEVM(REMOTE_CHAIN_ID);
    }

    function test_encodeTransaction(InterchainTransaction memory icTx) public view {
        bytes memory encoded = icClient.encodeTransaction(icTx);
        uint16 version = payloadLibHarness.getVersion(encoded);
        InterchainTransaction memory decoded = txLibHarness.decodeTransaction(payloadLibHarness.getPayload(encoded));
        assertEq(version, CLIENT_VERSION);
        assertEq(decoded, icTx);
    }

    function test_decodeOptions(OptionsV1 memory options) public view {
        bytes memory encoded = options.encodeOptionsV1();
        OptionsV1 memory decoded = icClient.decodeOptions(encoded);
        assertEq(decoded.gasLimit, options.gasLimit, "!gasLimit");
        assertEq(decoded.gasAirdrop, options.gasAirdrop, "!gasAirdrop");
    }
}
