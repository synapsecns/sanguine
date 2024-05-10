// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../contracts/libs/Options.sol";

import {InterchainAppMock} from "./mocks/InterchainAppMock.sol";
import {NoOpHarness} from "./harnesses/NoOpHarness.sol";

import {AppConfigV1, InterchainClientV1BaseTest, InterchainTransaction} from "./InterchainClientV1.Base.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainClientV1GenericViewsTest is InterchainClientV1BaseTest {
    address public moduleA = makeAddr("Module A");
    address public moduleB = makeAddr("Module B");

    address public app;
    address[] public oneModule;
    address[] public twoModules;

    function setUp() public override {
        super.setUp();
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        app = address(new InterchainAppMock());
        oneModule.push(moduleA);
        twoModules.push(moduleA);
        twoModules.push(moduleB);
    }

    function assertEq(AppConfigV1 memory expected, AppConfigV1 memory actual) internal pure {
        assertEq(expected.requiredResponses, actual.requiredResponses);
        assertEq(expected.optimisticPeriod, actual.optimisticPeriod);
        assertEq(expected.guardFlag, actual.guardFlag);
        assertEq(expected.guard, actual.guard);
        assertEq(abi.encode(expected), abi.encode(actual));
    }

    function test_getLinkedClient_chainIdKnown() public view {
        assertEq(icClient.getLinkedClient(REMOTE_CHAIN_ID), MOCK_REMOTE_CLIENT);
    }

    function test_getLinkedClient_chainIdUnknown() public view {
        assertEq(icClient.getLinkedClient(UNKNOWN_CHAIN_ID), 0);
    }

    function test_getLinkedClient_revert_chainIdLocal() public {
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
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
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        icClient.getLinkedClientEVM(LOCAL_CHAIN_ID);
    }

    function test_getLinkedClientEVM_revert_clientNotEVM() public {
        expectRevertLinkedClientNotEVM(MOCK_REMOTE_CLIENT);
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

    function test_getAppReceivingConfigV1() public {
        AppConfigV1 memory appConfig =
            AppConfigV1({requiredResponses: 1, optimisticPeriod: 30, guardFlag: 2, guard: address(4)});
        mockReceivingConfig(app, appConfig, twoModules);
        (AppConfigV1 memory fetchedConfig, address[] memory fetchedModules) = icClient.getAppReceivingConfigV1(app);
        assertEq(fetchedConfig, appConfig);
        assertEq(fetchedModules, twoModules);
    }

    function test_getAppReceivingConfigV1_revert_receiverEOA() public {
        address eoa = makeAddr("EOA");
        expectRevertReceiverNotICApp(eoa);
        icClient.getAppReceivingConfigV1(eoa);
    }

    function test_getAppReceivingConfigV1_revert_receiverNotICApp() public {
        address noOp = address(new NoOpHarness());
        expectRevertReceiverNotICApp(noOp);
        icClient.getAppReceivingConfigV1(noOp);
    }
}
