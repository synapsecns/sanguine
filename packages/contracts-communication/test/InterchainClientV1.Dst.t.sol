// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AppConfigV1} from "../contracts/libs/AppConfig.sol";
import {InterchainEntry} from "../contracts/libs/InterchainEntry.sol";
import {OptionsV1} from "../contracts/libs/Options.sol";
import {VersionedPayloadLib} from "../contracts/libs/VersionedPayload.sol";

import {
    InterchainClientV1,
    InterchainClientV1BaseTest,
    InterchainTransaction,
    InterchainTxDescriptor
} from "./InterchainClientV1.Base.t.sol";

import {InterchainAppMock} from "./mocks/InterchainAppMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering

/// @notice InterchainClientV1 destination chain tests.
/// # Happy path for interchainExecute.
/// 1. Decode the interchain transaction payload and derive the transaction ID.
/// 2. Check that the transaction has not been executed yet.
/// 3. Check that the transaction *could be* executed.
/// 3a. Check that msg.value matches the requested gasAirdrop from transaction options.
/// 3b. Construct the InterchainEntry that source InterchainClient was supposed to write for this transaction.
/// 3c. Fetch the AppConfig from the message recipient: list of modules, responses threshold and optimistic period.
/// 3d. Check that enough responses have been received for entry constructed in step 3b.
/// 4. Execute the transaction by passing the message to the recipient.
/// 5. Mark transaction as executed and emit an event.
contract InterchainClientV1DestinationTest is InterchainClientV1BaseTest {
    uint256 public constant MOCK_DB_NONCE = 444;
    uint64 public constant MOCK_ENTRY_INDEX = 4;

    uint256 public constant MOCK_LOCAL_DB_NONCE = 123;
    uint64 public constant MOCK_LOCAL_ENTRY_INDEX = 5;

    uint256 public constant MOCK_GAS_LIMIT = 100_000;
    uint256 public constant MOCK_GAS_AIRDROP = 1 ether;

    uint256 public constant MOCK_OPTIMISTIC_PERIOD = 10 minutes;
    uint256 public constant BIGGER_PERIOD = 7 days;

    bytes32 public constant MOCK_SRC_SENDER = keccak256("Sender");
    bytes public constant MOCK_MESSAGE = "Hello, World!";

    OptionsV1 public optionsAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: MOCK_GAS_AIRDROP});
    OptionsV1 public optionsNoAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: 0});

    bytes public invalidOptionsV0 = VersionedPayloadLib.encodeVersionedPayload(0, abi.encode(optionsAirdrop));
    bytes public invalidOptionsV1 = VersionedPayloadLib.encodeVersionedPayload(1, abi.encode(optionsAirdrop.gasLimit));

    AppConfigV1 public oneConfNoOP = AppConfigV1({requiredResponses: 1, optimisticPeriod: 0});
    AppConfigV1 public oneConfWithOP = AppConfigV1({requiredResponses: 1, optimisticPeriod: MOCK_OPTIMISTIC_PERIOD});
    AppConfigV1 public twoConfNoOP = AppConfigV1({requiredResponses: 2, optimisticPeriod: 0});
    AppConfigV1 public twoConfWithOP = AppConfigV1({requiredResponses: 2, optimisticPeriod: MOCK_OPTIMISTIC_PERIOD});

    address public executor = makeAddr("Executor");

    address public dstReceiver;
    bytes32 public dstReceiverBytes32;

    bytes32[] public emptyProof;

    address[] public oneModuleA;
    address[] public twoModules;

    // Possible module confirmation states:
    // - Not verified: verifiedAt == 0
    // - Almost verified: verified exactly "optimistic period" ago
    // - Just verified: verified exactly "optimistic period + 1 second" ago
    // - Over verified: verified long ago
    // Only "just verified" and "over verified" should be considered as verified.
    uint256 public constant INITIAL_TS = 1_704_067_200; // 2024-01-01 00:00:00 UTC

    uint256 public constant NOT_VERIFIED = 0;
    uint256 public constant ALMOST_VERIFIED = INITIAL_TS - MOCK_OPTIMISTIC_PERIOD;
    uint256 public constant JUST_VERIFIED = INITIAL_TS - MOCK_OPTIMISTIC_PERIOD - 1;
    uint256 public constant OVER_VERIFIED = INITIAL_TS - BIGGER_PERIOD;

    uint256 public constant JUST_NOW = INITIAL_TS - 1;

    function setUp() public override {
        vm.warp(INITIAL_TS);
        super.setUp();
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        dstReceiver = address(new InterchainAppMock());
        dstReceiverBytes32 = bytes32(uint256(uint160(dstReceiver)));
        oneModuleA.push(icModuleA);
        twoModules.push(icModuleA);
        twoModules.push(icModuleB);
    }

    /// @dev Override the InterchainApp's receiving config to return the given appConfig and modules.
    function mockReceivingConfig(AppConfigV1 memory appConfig, address[] memory modules) internal {
        bytes memory encodedConfig = appConfig.encodeAppConfigV1();
        vm.mockCall(
            dstReceiver, abi.encodeCall(InterchainAppMock.getReceivingConfig, ()), abi.encode(encodedConfig, modules)
        );
    }

    /// @dev Override the InterchainDB's verification check to return the given verifiedAt timestamp
    /// for given module and expected transaction entry.
    function mockCheckVerification(
        address dstModule,
        InterchainTxDescriptor memory desc,
        bytes32[] memory proof,
        uint256 verifiedAt
    )
        internal
    {
        InterchainEntry memory entry = InterchainEntry({
            srcChainId: REMOTE_CHAIN_ID,
            dbNonce: desc.dbNonce,
            entryIndex: desc.entryIndex,
            srcWriter: MOCK_REMOTE_CLIENT,
            dataHash: desc.transactionId
        });
        vm.mockCall(
            icDB, abi.encodeCall(InterchainDBMock.checkVerification, (dstModule, entry, proof)), abi.encode(verifiedAt)
        );
    }

    /// @dev Override the local DB's returned next entry index (both for reads and writes)
    function mockLocalNextEntryIndex(uint256 dbNonce, uint64 entryIndex) internal {
        bytes memory returnData = abi.encode(dbNonce, entryIndex);
        // Use partial calldata to override return values for calls to these functions with any arguments.
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.getNextEntryIndex.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntry.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntryWithVerification.selector), returnData);
    }

    /// @dev Constructs an interchain transaction and its descriptor for testing.
    function constructInterchainTx(bytes memory encodedOptions)
        internal
        view
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        icTx = InterchainTransaction({
            srcChainId: REMOTE_CHAIN_ID,
            srcSender: MOCK_SRC_SENDER,
            dstChainId: LOCAL_CHAIN_ID,
            dstReceiver: dstReceiverBytes32,
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            options: encodedOptions,
            message: MOCK_MESSAGE
        });
        desc = InterchainTxDescriptor({
            dbNonce: MOCK_DB_NONCE,
            entryIndex: MOCK_ENTRY_INDEX,
            transactionId: icTx.transactionId()
        });
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function expectAppReceiveCall(OptionsV1 memory options) internal {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainAppMock.appReceive,
            (REMOTE_CHAIN_ID, MOCK_SRC_SENDER, MOCK_DB_NONCE, MOCK_ENTRY_INDEX, MOCK_MESSAGE)
        );
        vm.expectCall({
            callee: dstReceiver,
            msgValue: options.gasAirdrop,
            gas: uint64(options.gasLimit),
            data: expectedCalldata,
            count: 1
        });
    }

    function assertExecutorSaved(InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) internal {
        assertEq(icClient.getExecutor(abi.encode(icTx)), executor, "!getExecutor");
        assertEq(icClient.getExecutorById(desc.transactionId), executor, "!getExecutorById");
    }

    function executeTransaction(
        InterchainTransaction memory icTx,
        OptionsV1 memory options,
        bytes32[] memory proof
    )
        internal
    {
        bytes memory encodedTx = abi.encode(icTx);
        deal(executor, options.gasAirdrop);
        vm.prank(executor);
        icClient.interchainExecute{value: options.gasAirdrop}(options.gasLimit, encodedTx, proof);
    }

    function prepareExecuteTest(
        bytes memory encodedOptions,
        AppConfigV1 memory appConfig,
        bytes32[] memory proof,
        address[] memory modules,
        uint256[] memory verificationTimes
    )
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        // Sanity check
        assert(modules.length == verificationTimes.length);
        (icTx, desc) = constructInterchainTx(encodedOptions);
        mockReceivingConfig(appConfig, modules);
        for (uint256 i = 0; i < modules.length; i++) {
            mockCheckVerification(modules[i], desc, proof, verificationTimes[i]);
        }
    }

    function prepareAlreadyExecutedTest(OptionsV1 memory options)
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        (icTx, desc) = prepareExecuteTest({
            encodedOptions: options.encodeOptionsV1(),
            appConfig: oneConfWithOP,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        executeTransaction(icTx, options, emptyProof);
        skip(1 days);
    }

    function checkHappyPathScenario(
        OptionsV1 memory options,
        InterchainTransaction memory icTx,
        InterchainTxDescriptor memory desc,
        bytes32[] memory proof
    )
        internal
    {
        expectAppReceiveCall(options);
        expectEventInterchainTransactionReceived(icTx, desc);
        assertTrue(icClient.isExecutable(abi.encode(icTx), proof));
        executeTransaction(icTx, options, proof);
        assertExecutorSaved(icTx, desc);
    }

    function checkSuccessA(
        OptionsV1 memory options,
        AppConfigV1 memory appConfig,
        uint256[] memory verificationTimes
    )
        internal
    {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareExecuteTest({
            encodedOptions: options.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: verificationTimes
        });
        checkHappyPathScenario(options, icTx, desc, emptyProof);
    }

    function checkSuccessAB(
        OptionsV1 memory options,
        AppConfigV1 memory appConfig,
        uint256[] memory verificationTimes
    )
        internal
    {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareExecuteTest({
            encodedOptions: options.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: twoModules,
            verificationTimes: verificationTimes
        });
        checkHappyPathScenario(options, icTx, desc, emptyProof);
    }

    function checkNotEnoughConfirmationsA(
        OptionsV1 memory options,
        AppConfigV1 memory appConfig,
        uint256 actualConfirmations,
        uint256[] memory verificationTimes
    )
        internal
    {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: options.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: verificationTimes
        });
        expectRevertNotEnoughResponses({actual: actualConfirmations, required: appConfig.requiredResponses});
        icClient.isExecutable(abi.encode(icTx), emptyProof);
        expectRevertNotEnoughResponses({actual: actualConfirmations, required: appConfig.requiredResponses});
        executeTransaction(icTx, options, emptyProof);
    }

    function checkNotEnoughConfirmationsAB(
        OptionsV1 memory options,
        AppConfigV1 memory appConfig,
        uint256 actualConfirmations,
        uint256[] memory verificationTimes
    )
        internal
    {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: options.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: twoModules,
            verificationTimes: verificationTimes
        });
        expectRevertNotEnoughResponses({actual: actualConfirmations, required: appConfig.requiredResponses});
        icClient.isExecutable(abi.encode(icTx), emptyProof);
        expectRevertNotEnoughResponses({actual: actualConfirmations, required: appConfig.requiredResponses});
        executeTransaction(icTx, options, emptyProof);
    }

    // ══════════════════════════════════════════ REQUIRED: 1, MODULES: A ══════════════════════════════════════════════

    function test_interchainExecute_1_A_periodNonZero_notVerifiedA_revert() public {
        checkNotEnoughConfirmationsA(optionsNoAirdrop, oneConfWithOP, 0, toArray(NOT_VERIFIED));
    }

    function test_interchainExecute_1_A_periodNonZero_almostVerifiedA_revert() public {
        checkNotEnoughConfirmationsA(optionsNoAirdrop, oneConfWithOP, 0, toArray(ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_A_periodNonZero_justVerifiedA_noAirdrop_success() public {
        checkSuccessA(optionsNoAirdrop, oneConfWithOP, toArray(JUST_VERIFIED));
    }

    function test_interchainExecute_1_A_periodNonZero_justVerifiedA_withAirdrop_success() public {
        checkSuccessA(optionsAirdrop, oneConfWithOP, toArray(JUST_VERIFIED));
    }

    function test_interchainExecute_1_A_periodNonZero_overVerifiedA_noAirdrop_success() public {
        checkSuccessA(optionsNoAirdrop, oneConfWithOP, toArray(OVER_VERIFIED));
    }

    function test_interchainExecute_1_A_periodNonZero_overVerifiedA_withAirdrop_success() public {
        checkSuccessA(optionsAirdrop, oneConfWithOP, toArray(OVER_VERIFIED));
    }

    // ═══════════════════════════════════ REQUIRED: 1, MODULES: A (PERIOD ZERO) ═══════════════════════════════════════

    function test_interchainExecute_1_A_periodZero_notVerifiedA_revert() public {
        checkNotEnoughConfirmationsA(optionsNoAirdrop, oneConfNoOP, 0, toArray(NOT_VERIFIED));
    }

    function test_interchainExecute_1_A_periodZero_almostVerifiedA_revert() public {
        checkNotEnoughConfirmationsA(optionsNoAirdrop, oneConfNoOP, 0, toArray(INITIAL_TS));
    }

    function test_interchainExecute_1_A_periodZero_justVerifiedA_noAirdrop_success() public {
        checkSuccessA(optionsNoAirdrop, oneConfNoOP, toArray(JUST_NOW));
    }

    function test_interchainExecute_1_A_periodZero_justVerifiedA_withAirdrop_success() public {
        checkSuccessA(optionsAirdrop, oneConfNoOP, toArray(JUST_NOW));
    }

    function test_interchainExecute_1_A_periodZero_overVerifiedA_noAirdrop_success() public {
        checkSuccessA(optionsNoAirdrop, oneConfNoOP, toArray(OVER_VERIFIED));
    }

    function test_interchainExecute_1_A_periodZero_overVerifiedA_withAirdrop_success() public {
        checkSuccessA(optionsAirdrop, oneConfNoOP, toArray(OVER_VERIFIED));
    }

    // ═════════════════════════════════════════ REQUIRED: 1, MODULES: A+B ═════════════════════════════════════════════

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfWithOP, 0, toArr(NOT_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfWithOP, 0, toArr(NOT_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(NOT_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(NOT_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_notVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfWithOP, 0, toArr(ALMOST_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfWithOP, 0, toArr(ALMOST_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(ALMOST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(ALMOST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(ALMOST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_almostVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(ALMOST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_notVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_notVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_almostVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_almostVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_justVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(JUST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_notVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_notVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_almostVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_almostVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodNonZero_overVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfWithOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    // ══════════════════════════════════ REQUIRED: 1, MODULES: A+B (PERIOD ZERO) ══════════════════════════════════════

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfNoOP, 0, toArr(NOT_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfNoOP, 0, toArr(NOT_VERIFIED, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(NOT_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(NOT_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_notVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfNoOP, 0, toArr(INITIAL_TS, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, oneConfNoOP, 0, toArr(INITIAL_TS, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(INITIAL_TS, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(INITIAL_TS, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(INITIAL_TS, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_almostVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(INITIAL_TS, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_notVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(JUST_NOW, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_notVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(JUST_NOW, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_almostVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(JUST_NOW, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_almostVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(JUST_NOW, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(JUST_NOW, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(JUST_NOW, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(JUST_NOW, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_justVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(JUST_NOW, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_notVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_notVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_almostVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_almostVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, INITIAL_TS));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_1_AB_periodZero_overVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, oneConfNoOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    // ═════════════════════════════════════════ REQUIRED: 2, MODULES: A+B ═════════════════════════════════════════════

    function test_interchainExecute_2_AB_periodNonZero_notVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 0, toArr(NOT_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_notVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 0, toArr(NOT_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_notVerifiedA_justVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(NOT_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_notVerifiedA_overVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_almostVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 0, toArr(ALMOST_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_almostVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 0, toArr(ALMOST_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_almostVerifiedA_justVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(ALMOST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_almostVerifiedA_overVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(ALMOST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(JUST_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(JUST_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfWithOP, toArr(JUST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfWithOP, toArr(JUST_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfWithOP, toArr(JUST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_justVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfWithOP, toArr(JUST_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfWithOP, 1, toArr(OVER_VERIFIED, ALMOST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfWithOP, toArr(OVER_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfWithOP, toArr(OVER_VERIFIED, JUST_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfWithOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodNonZero_overVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfWithOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    // ══════════════════════════════════ REQUIRED: 2, MODULES: A+B (PERIOD ZERO) ══════════════════════════════════════

    function test_interchainExecute_2_AB_periodZero_notVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 0, toArr(NOT_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_notVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 0, toArr(NOT_VERIFIED, INITIAL_TS));
    }

    function test_interchainExecute_2_AB_periodZero_notVerifiedA_justVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(NOT_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_notVerifiedA_overVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(NOT_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_almostVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 0, toArr(INITIAL_TS, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_almostVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 0, toArr(INITIAL_TS, INITIAL_TS));
    }

    function test_interchainExecute_2_AB_periodZero_almostVerifiedA_justVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(INITIAL_TS, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_almostVerifiedA_overVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(INITIAL_TS, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(JUST_NOW, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(JUST_NOW, INITIAL_TS));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfNoOP, toArr(JUST_NOW, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfNoOP, toArr(JUST_NOW, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfNoOP, toArr(JUST_NOW, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_justVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfNoOP, toArr(JUST_NOW, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_notVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(OVER_VERIFIED, NOT_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_almostVerifiedB_revert() public {
        checkNotEnoughConfirmationsAB(optionsNoAirdrop, twoConfNoOP, 1, toArr(OVER_VERIFIED, INITIAL_TS));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_justVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfNoOP, toArr(OVER_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_justVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfNoOP, toArr(OVER_VERIFIED, JUST_NOW));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_overVerifiedB_noAirdrop_success() public {
        checkSuccessAB(optionsNoAirdrop, twoConfNoOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    function test_interchainExecute_2_AB_periodZero_overVerifiedA_overVerifiedB_withAirdrop_success() public {
        checkSuccessAB(optionsAirdrop, twoConfNoOP, toArr(OVER_VERIFIED, OVER_VERIFIED));
    }

    // ═══════════════════════════════════════════ EXECUTE: MISC REVERTS ═══════════════════════════════════════════════

    function prepareExecutableTx(InterchainTransaction memory icTx) internal {
        InterchainTxDescriptor memory desc = InterchainTxDescriptor({
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            transactionId: icTx.transactionId()
        });
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        mockCheckVerification(icModuleA, desc, new bytes32[](0), JUST_VERIFIED);
    }

    function test_interchainExecute_revert_srcChainNotRemote() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.srcChainId = LOCAL_CHAIN_ID;
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_srcChainNotLinked() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.srcChainId = UNKNOWN_CHAIN_ID;
        expectRevertNoLinkedClient(UNKNOWN_CHAIN_ID);
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_dstChainIncorrect() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.dstChainId = UNKNOWN_CHAIN_ID;
        expectRevertIncorrectDstChainId(UNKNOWN_CHAIN_ID);
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_emptyOptions() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx("");
        prepareExecutableTx(icTx);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_invalidOptionsV0() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV0);
        prepareExecutableTx(icTx);
        expectRevertIncorrectVersion(0);
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_invalidOptionsV1() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV1);
        prepareExecutableTx(icTx);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_alreadyExecuted() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareAlreadyExecutedTest(optionsNoAirdrop);
        expectRevertTxAlreadyExecuted(desc.transactionId);
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_withAirdrop_zeroMsgValue() public {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsAirdrop.encodeOptionsV1(),
            appConfig: oneConfWithOP,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        uint256 requiredValue = optionsAirdrop.gasAirdrop;
        optionsAirdrop.gasAirdrop = 0;
        expectRevertIncorrectMsgValue({actual: 0, required: requiredValue});
        executeTransaction(icTx, optionsAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_withAirdrop_lowerMsgValue() public {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsAirdrop.encodeOptionsV1(),
            appConfig: oneConfWithOP,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        uint256 requiredValue = optionsAirdrop.gasAirdrop;
        optionsAirdrop.gasAirdrop = requiredValue - 1;
        expectRevertIncorrectMsgValue({actual: requiredValue - 1, required: requiredValue});
        executeTransaction(icTx, optionsAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_withAirdrop_higherMsgValue() public {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsAirdrop.encodeOptionsV1(),
            appConfig: oneConfWithOP,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        uint256 requiredValue = optionsAirdrop.gasAirdrop;
        optionsAirdrop.gasAirdrop = requiredValue + 1;
        expectRevertIncorrectMsgValue({actual: requiredValue + 1, required: requiredValue});
        executeTransaction(icTx, optionsAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_noAirdrop_nonZeroMsgValue() public {
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsNoAirdrop.encodeOptionsV1(),
            appConfig: oneConfWithOP,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        optionsNoAirdrop.gasAirdrop = MOCK_GAS_AIRDROP;
        expectRevertIncorrectMsgValue({actual: MOCK_GAS_AIRDROP, required: 0});
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_interchainExecute_revert_zeroRequiredResponses() public {
        AppConfigV1 memory appConfig = oneConfWithOP;
        appConfig.requiredResponses = 0;
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsNoAirdrop.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        expectRevertZeroRequiredResponses();
        executeTransaction(icTx, optionsNoAirdrop, emptyProof);
    }

    function test_isExecutable_revert_srcChainNotRemote() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.srcChainId = LOCAL_CHAIN_ID;
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_srcChainNotLinked() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.srcChainId = UNKNOWN_CHAIN_ID;
        expectRevertNoLinkedClient(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_dstChainIncorrect() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(optionsNoAirdrop.encodeOptionsV1());
        mockReceivingConfig(oneConfWithOP, oneModuleA);
        icTx.dstChainId = UNKNOWN_CHAIN_ID;
        expectRevertIncorrectDstChainId(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_emptyOptions() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx("");
        prepareExecutableTx(icTx);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_invalidOptionsV0() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV0);
        prepareExecutableTx(icTx);
        expectRevertIncorrectVersion(0);
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_invalidOptionsV1() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV1);
        prepareExecutableTx(icTx);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_alreadyExecuted() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareAlreadyExecutedTest(optionsNoAirdrop);
        expectRevertTxAlreadyExecuted(desc.transactionId);
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    function test_isExecutable_revert_zeroRequiredResponses() public {
        AppConfigV1 memory appConfig = oneConfWithOP;
        appConfig.requiredResponses = 0;
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            encodedOptions: optionsNoAirdrop.encodeOptionsV1(),
            appConfig: appConfig,
            proof: emptyProof,
            modules: oneModuleA,
            verificationTimes: toArray(JUST_VERIFIED)
        });
        expectRevertZeroRequiredResponses();
        icClient.isExecutable(abi.encode(icTx), emptyProof);
    }

    // ═══════════════════════════════════════ TESTS: WRITE EXECUTION PROOF ════════════════════════════════════════════

    function test_writeExecutionProof() public {
        (, InterchainTxDescriptor memory desc) = prepareAlreadyExecutedTest(optionsNoAirdrop);
        bytes32 proofHash = keccak256(abi.encode(desc.transactionId, executor));
        bytes memory expectedCalldata = abi.encodeCall(InterchainDBMock.writeEntry, (proofHash));
        mockLocalNextEntryIndex(MOCK_LOCAL_DB_NONCE, MOCK_LOCAL_ENTRY_INDEX);
        vm.expectCall({callee: icDB, data: expectedCalldata, count: 1});
        expectEventExecutionProofWritten(desc.transactionId, MOCK_LOCAL_DB_NONCE, MOCK_LOCAL_ENTRY_INDEX, executor);
        icClient.writeExecutionProof(desc.transactionId);
    }

    function test_writeExecutionProof_revert_notExecuted() public {
        bytes32 mockTxId = keccak256("mockTxId");
        expectRevertTxNotExecuted(mockTxId);
        icClient.writeExecutionProof(mockTxId);
    }
}
