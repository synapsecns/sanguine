// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../contracts/interfaces/IInterchainClientV1.sol";
import {AppConfigV1} from "../contracts/libs/AppConfig.sol";
import {InterchainBatch} from "../contracts/libs/InterchainBatch.sol";
import {InterchainEntryLib} from "../contracts/libs/InterchainEntry.sol";
import {OptionsV1} from "../contracts/libs/Options.sol";
import {VersionedPayloadLib} from "../contracts/libs/VersionedPayload.sol";

import {
    InterchainClientV1BaseTest, InterchainTransaction, InterchainTxDescriptor
} from "./InterchainClientV1.Base.t.sol";

import {NoOpHarness} from "./harnesses/NoOpHarness.sol";
import {InterchainAppMock} from "./mocks/InterchainAppMock.sol";
import {InterchainDBMock} from "./mocks/InterchainDBMock.sol";

// solhint-disable code-complexity
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
abstract contract InterchainClientV1DstTest is InterchainClientV1BaseTest {
    uint64 public constant MOCK_DB_NONCE = 444;
    uint64 public constant MOCK_ENTRY_INDEX = 0;

    uint64 public constant MOCK_LOCAL_DB_NONCE = 123;
    uint64 public constant MOCK_LOCAL_ENTRY_INDEX = 0;

    uint256 public constant BIGGER_PERIOD = 7 days;

    bytes32 public constant MOCK_SRC_SENDER = keccak256("Sender");
    bytes public constant MOCK_MESSAGE = "Hello, World!";

    uint256 public constant GUARD_DISABLED = 0;
    uint256 public constant GUARD_DEFAULT = 1;
    uint256 public constant GUARD_CUSTOM = 2;

    address public customGuard = makeAddr("Custom Guard");

    address public executor = makeAddr("Executor");

    address public dstReceiver;
    bytes32 public dstReceiverBytes32;

    address public receiverEOA = makeAddr("Receiver EOA");
    address public receiverNotICApp;

    bytes32[] public emptyProof;

    address[] public oneModuleA;
    address[] public twoModules;

    // Possible module confirmation states:
    // - Not verified: verifiedAt == 0
    // - Almost verified: verified exactly "optimistic period" ago
    // - Just verified: verified exactly "optimistic period + 1 second" ago
    // - Over verified: verified long ago
    // - Conflict: module verified a conflicting batch
    // Only "just verified" and "over verified" should be considered as verified.
    uint256 public constant INITIAL_TS = 1_704_067_200; // 2024-01-01 00:00:00 UTC

    uint256 public constant NOT_VERIFIED = 0;
    uint256 public constant OVER_VERIFIED = INITIAL_TS - BIGGER_PERIOD;
    uint256 public constant CONFLICT = type(uint256).max;

    function setUp() public override {
        vm.warp(INITIAL_TS);
        super.setUp();
        setDefaultGuard(defaultGuard);
        setLinkedClient(REMOTE_CHAIN_ID, MOCK_REMOTE_CLIENT);
        dstReceiver = address(new InterchainAppMock());
        dstReceiverBytes32 = bytes32(uint256(uint160(dstReceiver)));
        receiverNotICApp = address(new NoOpHarness());
        oneModuleA.push(icModuleA);
        twoModules.push(icModuleA);
        twoModules.push(icModuleB);
    }

    // ════════════════════════════════════════════════ MOCK TOOLS ═════════════════════════════════════════════════════

    /// @dev Override the InterchainApp's receiving config to return the given appConfig and two modules.
    function mockReceivingConfig(uint256 requiredResponses, uint256 guardFlag) internal {
        AppConfigV1 memory appConfig = getAppConfig(requiredResponses, guardFlag);
        bytes memory encodedConfig = appConfig.encodeAppConfigV1();
        vm.mockCall(
            dstReceiver, abi.encodeCall(InterchainAppMock.getReceivingConfig, ()), abi.encode(encodedConfig, twoModules)
        );
    }

    /// @dev Override the InterchainDB's verification check to return the given verifiedAt timestamp
    /// for given module and expected transaction entry.
    function mockCheckVerification(
        address dstModule,
        InterchainTxDescriptor memory desc,
        uint256 verifiedAt
    )
        internal
    {
        InterchainBatch memory batch = InterchainBatch({
            srcChainId: REMOTE_CHAIN_ID,
            dbNonce: desc.dbNonce,
            batchRoot: InterchainEntryLib.getEntryValue({srcWriter: MOCK_REMOTE_CLIENT, dataHash: desc.transactionId})
        });
        vm.mockCall(
            icDB, abi.encodeCall(InterchainDBMock.checkBatchVerification, (dstModule, batch)), abi.encode(verifiedAt)
        );
    }

    /// @dev Override the local DB's returned next entry index (both for reads and writes)
    function mockLocalNextEntryIndex(uint64 dbNonce, uint64 entryIndex) internal {
        bytes memory returnData = abi.encode(dbNonce, entryIndex);
        // Use partial calldata to override return values for calls to these functions with any arguments.
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.getNextEntryIndex.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntry.selector), returnData);
        vm.mockCall(icDB, abi.encodeWithSelector(InterchainDBMock.writeEntryWithVerification.selector), returnData);
    }

    // ═════════════════════════════════════════════════ TEST DATA ═════════════════════════════════════════════════════

    function getOptions() internal view virtual returns (OptionsV1 memory);
    function getOptimisticPeriod() internal view virtual returns (uint256);

    function getAppConfig(uint256 requiredResponses, uint256 guardFlag) internal view returns (AppConfigV1 memory) {
        return AppConfigV1({
            requiredResponses: requiredResponses,
            optimisticPeriod: getOptimisticPeriod(),
            guardFlag: guardFlag,
            guard: guardFlag == GUARD_CUSTOM ? customGuard : address(0)
        });
    }

    function almostVerTS() internal view returns (uint256) {
        return INITIAL_TS - getOptimisticPeriod();
    }

    function justVerTS() internal view returns (uint256) {
        return INITIAL_TS - getOptimisticPeriod() - 1;
    }

    function getTimestampFixture(uint256 index) internal view returns (uint256) {
        index = index % 5;
        if (index == 0) {
            return NOT_VERIFIED;
        } else if (index == 1) {
            return almostVerTS();
        } else if (index == 2) {
            return justVerTS();
        } else if (index == 3) {
            return OVER_VERIFIED;
        } else {
            return CONFLICT;
        }
    }

    function boundTimestamp(uint256 timestamp) internal view returns (uint256) {
        if (timestamp == CONFLICT) {
            return timestamp;
        } else {
            return timestamp % block.timestamp;
        }
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
        desc = getTxDescriptor(icTx);
    }

    function constructInterchainTx()
        internal
        view
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        return constructInterchainTx(getOptions().encodeOptionsV1());
    }

    function getTxDescriptor(InterchainTransaction memory icTx) internal view returns (InterchainTxDescriptor memory) {
        return InterchainTxDescriptor({
            dbNonce: icTx.dbNonce,
            entryIndex: icTx.entryIndex,
            transactionId: keccak256(getEncodedTx(icTx))
        });
    }

    // ══════════════════════════════════════════════ TEST ASSERTIONS ══════════════════════════════════════════════════

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

    function assertExecutorSaved(InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) internal view {
        assertEq(icClient.getExecutor(getEncodedTx(icTx)), executor, "!getExecutor");
        assertEq(icClient.getExecutorById(desc.transactionId), executor, "!getExecutorById");
    }

    function assertCorrectReadiness(
        InterchainTransaction memory icTx,
        IInterchainClientV1.TxReadiness expected
    )
        internal
        view
    {
        assertCorrectReadiness(icTx, expected, 0, 0);
    }

    function assertCorrectReadiness(
        InterchainTransaction memory icTx,
        IInterchainClientV1.TxReadiness expected,
        address expectedFirstArg
    )
        internal
        view
    {
        assertCorrectReadiness(icTx, expected, uint256(uint160(expectedFirstArg)), 0);
    }

    function assertCorrectReadiness(
        InterchainTransaction memory icTx,
        IInterchainClientV1.TxReadiness expected,
        uint256 expectedFirstArg
    )
        internal
        view
    {
        assertCorrectReadiness(icTx, expected, expectedFirstArg, 0);
    }

    function assertCorrectReadiness(
        InterchainTransaction memory icTx,
        IInterchainClientV1.TxReadiness expected,
        uint256 expectedFirstArg,
        uint256 expectedSecondArg
    )
        internal
        view
    {
        assertCorrectReadiness(icTx, emptyProof, expected, expectedFirstArg, expectedSecondArg);
    }

    function assertCorrectReadiness(
        InterchainTransaction memory icTx,
        bytes32[] memory proof,
        IInterchainClientV1.TxReadiness expected,
        uint256 expectedFirstArg,
        uint256 expectedSecondArg
    )
        internal
        view
    {
        (IInterchainClientV1.TxReadiness actual, bytes32 firstArg, bytes32 secondArg) =
            icClient.getTxReadinessV1(icTx, proof);
        assertEq(uint256(actual), uint256(expected));
        assertEq(firstArg, bytes32(expectedFirstArg));
        assertEq(secondArg, bytes32(expectedSecondArg));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function executeTransaction(bytes memory encodedTx, bytes32[] memory proof) internal {
        OptionsV1 memory options = getOptions();
        deal(executor, options.gasAirdrop);
        vm.prank(executor);
        icClient.interchainExecute{value: options.gasAirdrop}(options.gasLimit, encodedTx, proof);
    }

    function prepareExecuteTest(
        uint256 requiredResponses,
        uint256 guardFlag,
        uint256[] memory verificationTimes
    )
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        // Sanity check
        assert(twoModules.length == verificationTimes.length);
        (icTx, desc) = constructInterchainTx();
        mockReceivingConfig(requiredResponses, guardFlag);
        for (uint256 i = 0; i < twoModules.length; i++) {
            mockCheckVerification(twoModules[i], desc, verificationTimes[i]);
        }
    }

    function prepareAlreadyExecutedTest()
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        (icTx, desc) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: 1,
            verificationTimes: toArr(OVER_VERIFIED, OVER_VERIFIED)
        });
        bytes memory encodedTx = getEncodedTx(icTx);
        executeTransaction(encodedTx, emptyProof);
        skip(1 days);
    }

    function makeTxDescriptorExecutable(InterchainTxDescriptor memory desc) internal {
        mockReceivingConfig({requiredResponses: 1, guardFlag: 0});
        mockCheckVerification(icModuleA, desc, justVerTS());
    }

    function addGuardConflict(address guard) internal {
        (, InterchainTxDescriptor memory desc) = constructInterchainTx();
        mockCheckVerification(guard, desc, CONFLICT);
    }

    /// @dev Assuming the transaction is ready to be executed, check the happy path scenario.
    function executeAndCheck(InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) internal {
        OptionsV1 memory options = getOptions();
        expectAppReceiveCall(options);
        expectEventInterchainTransactionReceived(icTx, desc);
        bytes memory encodedTx = getEncodedTx(icTx);
        assertTrue(icClient.isExecutable(encodedTx, emptyProof));
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.Ready);
        executeTransaction(encodedTx, emptyProof);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.AlreadyExecuted, uint256(desc.transactionId));
        assertExecutorSaved(icTx, desc);
        assertEq(dstReceiver.balance, options.gasAirdrop);
    }

    /// @dev Execute the happy path scenario with the given modules and verification times.
    function checkHappyPath(uint256 required, uint256 guardFlag, uint256[] memory times) internal {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareExecuteTest(required, guardFlag, times);
        executeAndCheck(icTx, desc);
    }

    /// @dev Execute the scenario where not enough confirmations have been received.
    /// Both `isExecutable` and `interchainExecute` should revert.
    function checkNotEnoughConfirmations(
        uint256 actual,
        uint256 required,
        uint256 guardFlag,
        uint256[] memory times
    )
        internal
    {
        (InterchainTransaction memory icTx,) = prepareExecuteTest(required, guardFlag, times);
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.NotEnoughResponses, actual, required);
        expectRevertNotEnoughResponses({actual: actual, required: required});
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertNotEnoughResponses({actual: actual, required: required});
        executeTransaction(encodedTx, emptyProof);
    }

    /// @dev Execute the scenario where a batch conflict has been detected.
    /// Both `isExecutable` and `interchainExecute` should revert.
    function checkBatchConflict(address module, uint256 required, uint256 guardFlag, uint256[] memory times) internal {
        (InterchainTransaction memory icTx,) = prepareExecuteTest(required, guardFlag, times);
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.BatchConflict, module);
        expectRevertBatchConflict(module);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertBatchConflict(module);
        executeTransaction(encodedTx, emptyProof);
    }

    function checkScenario(uint256 required, uint256 guardFlag, uint256 indexA, uint256 indexB) internal {
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        uint256[] memory times = toArr(timeA, timeB);
        // Check for conflicts
        if (timeA == CONFLICT) {
            checkBatchConflict({module: icModuleA, required: required, guardFlag: guardFlag, times: times});
            return;
        }
        if (timeB == CONFLICT) {
            checkBatchConflict({module: icModuleB, required: required, guardFlag: guardFlag, times: times});
            return;
        }
        uint256 actual = 0;
        if (timeA == justVerTS() || timeA == OVER_VERIFIED) {
            actual++;
        }
        if (timeB == justVerTS() || timeB == OVER_VERIFIED) {
            actual++;
        }
        if (actual >= required) {
            checkHappyPath({required: required, guardFlag: guardFlag, times: times});
        } else {
            checkNotEnoughConfirmations({actual: actual, required: required, guardFlag: guardFlag, times: times});
        }
    }

    // ═════════════════════════════════════ APP CONFIG: 1 OUT OF 2 RESPONSES ══════════════════════════════════════════

    function test_execute_1_2_noGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: GUARD_DISABLED, indexA: indexA, indexB: indexB});
    }

    /// @dev Guard conflict should not affect the behavior if the app didn't opt-in for the guard.
    function test_execute_1_2_noGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_1_2_noGuard(indexA, indexB);
    }

    function test_execute_1_2_defaultGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: GUARD_DEFAULT, indexA: indexA, indexB: indexB});
    }

    /// @dev Guard conflict should always revert the transaction if the app opted-in for the guard.
    function test_execute_1_2_defaultGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: defaultGuard, required: 1, guardFlag: GUARD_DEFAULT, times: toArr(timeA, timeB)});
    }

    function test_execute_1_2_customGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: 2, indexA: indexA, indexB: indexB});
    }

    /// @dev Custom guard conflict should always revert the transaction if the app opted-in for the custom guard.
    function test_execute_1_2_customGuard_customGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(customGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: customGuard, required: 1, guardFlag: GUARD_CUSTOM, times: toArr(timeA, timeB)});
    }

    /// @dev Default guard conflict should not affect the behavior if the app opted-in for a custom guard.
    function test_execute_1_2_customGuard_defaultGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_1_2_customGuard(indexA, indexB);
    }

    /// @dev Default Guard conflict should be ignored if the app opted-in for a custom guard,
    /// but the custom guard conflict should still revert the transaction.
    function test_execute_1_2_customGuard_bothGuardsConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        addGuardConflict(customGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: customGuard, required: 1, guardFlag: GUARD_CUSTOM, times: toArr(timeA, timeB)});
    }

    // ═════════════════════════════════════ APP CONFIG: 2 OUT OF 2 RESPONSES ══════════════════════════════════════════

    function test_execute_2_2_noGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_DISABLED, indexA: indexA, indexB: indexB});
    }

    /// @dev Guard conflict should not affect the behavior if the app didn't opt-in for the guard.
    function test_execute_2_2_noGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_2_2_noGuard(indexA, indexB);
    }

    function test_execute_2_2_defaultGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_DEFAULT, indexA: indexA, indexB: indexB});
    }

    /// @dev Guard conflict should always revert the transaction if the app opted-in for the guard.
    function test_execute_2_2_defaultGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: defaultGuard, required: 2, guardFlag: GUARD_DEFAULT, times: toArr(timeA, timeB)});
    }

    function test_execute_2_2_customGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_CUSTOM, indexA: indexA, indexB: indexB});
    }

    /// @dev Custom guard conflict should always revert the transaction if the app opted-in for the custom guard.
    function test_execute_2_2_customGuard_customGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(customGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: customGuard, required: 2, guardFlag: GUARD_CUSTOM, times: toArr(timeA, timeB)});
    }

    /// @dev Default guard conflict should not affect the behavior if the app opted-in for a custom guard.
    function test_execute_2_2_customGuard_defaultGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_2_2_customGuard(indexA, indexB);
    }

    /// @dev Default Guard conflict should be ignored if the app opted-in for a custom guard,
    /// but the custom guard conflict should still revert the transaction.
    function test_execute_2_2_customGuard_bothGuardsConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        addGuardConflict(customGuard);
        uint256 timeA = getTimestampFixture(indexA);
        uint256 timeB = getTimestampFixture(indexB);
        checkBatchConflict({module: customGuard, required: 2, guardFlag: GUARD_CUSTOM, times: toArr(timeA, timeB)});
    }

    // ═══════════════════════════════════════════ EXECUTE: MISC REVERTS ═══════════════════════════════════════════════

    function encodeAndMakeExecutable(InterchainTransaction memory icTx) internal returns (bytes memory) {
        bytes memory encodedTx = getEncodedTx(icTx);
        makeTxDescriptorExecutable(getTxDescriptor(icTx));
        return encodedTx;
    }

    function test_execute_revert_invalidTransactionVersion(uint16 version) public {
        vm.assume(version != CLIENT_VERSION);
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        bytes memory invalidVersionTx = VersionedPayloadLib.encodeVersionedPayload(version, abi.encode(icTx));
        makeTxDescriptorExecutable(getTxDescriptor(icTx));
        expectRevertInvalidTransactionVersion(version);
        icClient.isExecutable(invalidVersionTx, emptyProof);
        expectRevertInvalidTransactionVersion(version);
        executeTransaction(invalidVersionTx, emptyProof);
    }

    function test_execute_revert_srcChainNotRemote() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.srcChainId = LOCAL_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_srcChainNotLinked() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.srcChainId = UNKNOWN_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertNoLinkedClient(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertNoLinkedClient(UNKNOWN_CHAIN_ID);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_dstChainIncorrect() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstChainId = UNKNOWN_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.IncorrectDstChainId, UNKNOWN_CHAIN_ID);
        expectRevertIncorrectDstChainId(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertIncorrectDstChainId(UNKNOWN_CHAIN_ID);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_revert_nonZeroEntryIndex() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.entryIndex = 1;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertIncorrectEntryIndex(icTx.entryIndex);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertIncorrectEntryIndex(icTx.entryIndex);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_revert_nonEmptyProof() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        bytes32[] memory proof = new bytes32[](1);
        assertCorrectReadiness(icTx, proof, IInterchainClientV1.TxReadiness.UndeterminedRevert, 0, 0);
        expectRevertIncorrectProof();
        icClient.isExecutable(encodedTx, proof);
        expectRevertIncorrectProof();
        executeTransaction(encodedTx, proof);
    }

    function test_execute_revert_emptyOptions() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx("");
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.isExecutable(encodedTx, emptyProof);
        vm.expectRevert();
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_invalidOptionsV0() public {
        bytes memory invalidOptionsV0 =
            VersionedPayloadLib.encodeVersionedPayload({version: 0, payload: abi.encode(getOptions())});
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV0);
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertIncorrectVersion(0);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertIncorrectVersion(0);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_invalidOptionsV1() public {
        // Only include a single field to make the payload invalid.
        bytes memory invalidOptionsV1 =
            VersionedPayloadLib.encodeVersionedPayload({version: 1, payload: abi.encode(getOptions().gasLimit)});
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV1);
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.isExecutable(encodedTx, emptyProof);
        vm.expectRevert();
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_receiverEOA() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstReceiver = bytes32(uint256(uint160(receiverEOA)));
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        expectRevertReceiverNotICApp(receiverEOA);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertReceiverNotICApp(receiverEOA);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_receiverNotICApp() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstReceiver = bytes32(uint256(uint160(receiverNotICApp)));
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        expectRevertReceiverNotICApp(receiverNotICApp);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertReceiverNotICApp(receiverNotICApp);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_alreadyExecuted() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareAlreadyExecutedTest();
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.AlreadyExecuted, uint256(desc.transactionId));
        expectRevertTxAlreadyExecuted(desc.transactionId);
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertTxAlreadyExecuted(desc.transactionId);
        executeTransaction(encodedTx, emptyProof);
    }

    function test_execute_revert_zeroRequiredResponses() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = constructInterchainTx();
        bytes memory encodedTx = getEncodedTx(icTx);
        mockReceivingConfig({requiredResponses: 0, guardFlag: 0});
        mockCheckVerification(icModuleA, desc, justVerTS());
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.ZeroRequiredResponses);
        expectRevertZeroRequiredResponses();
        icClient.isExecutable(encodedTx, emptyProof);
        expectRevertZeroRequiredResponses();
        executeTransaction(encodedTx, emptyProof);
    }
}
