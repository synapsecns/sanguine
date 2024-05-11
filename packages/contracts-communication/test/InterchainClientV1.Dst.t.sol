// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainClientV1} from "../contracts/interfaces/IInterchainClientV1.sol";
import {AppConfigV1} from "../contracts/libs/AppConfig.sol";
import {InterchainEntry, InterchainEntryLib} from "../contracts/libs/InterchainEntry.sol";
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

    uint64 public constant MOCK_LOCAL_DB_NONCE = 123;

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

    address[] public modulesA;
    address[] public modulesAB;

    // Possible module confirmation states:
    // - Not verified: verifiedAt == 0
    // - Almost verified: verified exactly "optimistic period" ago
    // - Just verified: verified exactly "optimistic period + 1 second" ago
    // - Over verified: verified long ago
    // - Conflict: module verified a conflicting entry
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
        modulesA.push(icModuleA);
        modulesAB.push(icModuleA);
        modulesAB.push(icModuleB);
    }

    // TODO: add tests for 0 and 1 modules

    // ════════════════════════════════════════════════ MOCK TOOLS ═════════════════════════════════════════════════════

    /// @dev Override the InterchainApp's receiving config to return the given appConfig and two modules.
    function mockReceivingConfig(uint256 requiredResponses, uint256 guardFlag, address[] memory modules) internal {
        AppConfigV1 memory appConfig = getAppConfig(requiredResponses, guardFlag);
        mockReceivingConfig(dstReceiver, appConfig, modules);
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
        InterchainEntry memory entry = InterchainEntry({
            srcChainId: REMOTE_CHAIN_ID,
            dbNonce: desc.dbNonce,
            entryValue: InterchainEntryLib.getEntryValue({srcWriter: MOCK_REMOTE_CLIENT, digest: desc.transactionId})
        });
        vm.mockCall(
            icDB, abi.encodeCall(InterchainDBMock.checkEntryVerification, (dstModule, entry)), abi.encode(verifiedAt)
        );
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
        return InterchainTxDescriptor({dbNonce: icTx.dbNonce, transactionId: keccak256(getEncodedTx(icTx))});
    }

    function getRequiredCount(uint256 requiredConfig, uint256 modulesCount) internal pure returns (uint256) {
        return requiredConfig != 0 ? requiredConfig : modulesCount;
    }

    // ══════════════════════════════════════════════ TEST ASSERTIONS ══════════════════════════════════════════════════

    function expectAppReceiveCall(OptionsV1 memory options) internal {
        bytes memory expectedCalldata = abi.encodeCall(
            InterchainAppMock.appReceive, (REMOTE_CHAIN_ID, MOCK_SRC_SENDER, MOCK_DB_NONCE, MOCK_MESSAGE)
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
        (IInterchainClientV1.TxReadiness actual, bytes32 firstArg, bytes32 secondArg) = icClient.getTxReadinessV1(icTx);
        assertEq(uint256(actual), uint256(expected));
        assertEq(firstArg, bytes32(expectedFirstArg));
        assertEq(secondArg, bytes32(expectedSecondArg));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function executeTransaction(bytes memory encodedTx) internal {
        OptionsV1 memory options = getOptions();
        deal(executor, options.gasAirdrop);
        vm.prank(executor);
        icClient.interchainExecute{value: options.gasAirdrop}(options.gasLimit, encodedTx);
    }

    function prepareExecuteTest(
        uint256 requiredResponses,
        uint256 guardFlag,
        address[] memory modules,
        uint256[] memory times
    )
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        // Sanity check
        assert(modules.length == times.length);
        (icTx, desc) = constructInterchainTx();
        mockReceivingConfig(requiredResponses, guardFlag, modules);
        for (uint256 i = 0; i < modules.length; i++) {
            mockCheckVerification(modules[i], desc, times[i]);
        }
    }

    function prepareAlreadyExecutedTest()
        internal
        returns (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc)
    {
        (icTx, desc) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: 1,
            modules: modulesAB,
            times: toArr(OVER_VERIFIED, OVER_VERIFIED)
        });
        bytes memory encodedTx = getEncodedTx(icTx);
        executeTransaction(encodedTx);
        skip(1 days);
    }

    function makeTxDescriptorExecutable(InterchainTxDescriptor memory desc) internal {
        mockReceivingConfig({requiredResponses: 1, guardFlag: 0, modules: modulesA});
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
        assertTrue(icClient.isExecutable(encodedTx));
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.Ready);
        executeTransaction(encodedTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.AlreadyExecuted, uint256(desc.transactionId));
        assertExecutorSaved(icTx, desc);
        assertEq(dstReceiver.balance, options.gasAirdrop);
    }

    /// @dev Execute the scenario where not enough confirmations have been received.
    /// Both `isExecutable` and `interchainExecute` should revert.
    function checkNotEnoughConfirmations(
        uint256 actual,
        uint256 required,
        InterchainTransaction memory icTx
    )
        internal
    {
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.EntryAwaitingResponses, actual, required);
        expectRevertResponsesAmountBelowMin({actual: actual, required: required});
        icClient.isExecutable(encodedTx);
        expectRevertResponsesAmountBelowMin({actual: actual, required: required});
        executeTransaction(encodedTx);
    }

    /// @dev Execute the scenario where an entry conflict has been detected.
    /// Both `isExecutable` and `interchainExecute` should revert.
    function checkEntryConflict(address module, InterchainTransaction memory icTx) internal {
        // Don't adjust required count as it's only used to set the app config
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.EntryConflict, module);
        expectRevertEntryConflict(module);
        icClient.isExecutable(encodedTx);
        expectRevertEntryConflict(module);
        executeTransaction(encodedTx);
    }

    function checkScenario(
        uint256 required,
        uint256 guardFlag,
        address[] memory modules,
        uint256[] memory timeIndices
    )
        internal
    {
        uint256[] memory times = new uint256[](timeIndices.length);
        for (uint256 i = 0; i < timeIndices.length; i++) {
            times[i] = getTimestampFixture(timeIndices[i]);
        }
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) =
            prepareExecuteTest(required, guardFlag, modules, times);
        // Check for conflicts
        for (uint256 i = 0; i < modules.length; i++) {
            if (times[i] == CONFLICT) {
                checkEntryConflict({module: modules[i], icTx: icTx});
                return;
            }
        }
        // If the guard is disabled, the last accepted timestamp is the initial timestamp - 1.
        // Otherwise, it's the "just verified" timestamp: initial timestamp - optimistic period - 1.
        uint256 lastAcceptedTS = (guardFlag == GUARD_DISABLED) ? INITIAL_TS - 1 : justVerTS();
        uint256 actual = 0;
        for (uint256 i = 0; i < modules.length; i++) {
            if (times[i] != NOT_VERIFIED && times[i] <= lastAcceptedTS) {
                actual++;
            }
        }
        // Use a default threshold equal to amount of modules if the required responses is zero
        uint256 requiredAdjusted = getRequiredCount({requiredConfig: required, modulesCount: modules.length});
        if (actual >= requiredAdjusted) {
            executeAndCheck(icTx, desc);
        } else {
            checkNotEnoughConfirmations(actual, requiredAdjusted, icTx);
        }
    }

    // ═════════════════════════════════════ APP CONFIG: 1 OUT OF 2 RESPONSES ══════════════════════════════════════════

    function test_execute_1_2_noGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: GUARD_DISABLED, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should not affect the behavior if the app didn't opt-in for the guard.
    function test_execute_1_2_noGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_1_2_noGuard(indexA, indexB);
    }

    function test_execute_1_2_defaultGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: GUARD_DEFAULT, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should always revert the transaction if the app opted-in for the guard.
    function test_execute_1_2_defaultGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 1,
            guardFlag: GUARD_DEFAULT,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: defaultGuard, icTx: icTx});
    }

    function test_execute_1_2_customGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 1, guardFlag: GUARD_CUSTOM, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Custom guard conflict should always revert the transaction if the app opted-in for the custom guard.
    function test_execute_1_2_customGuard_customGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(customGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 1,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
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
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 1,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
    }

    // ═════════════════════════════════════ APP CONFIG: 2 OUT OF 2 RESPONSES ══════════════════════════════════════════

    function test_execute_2_2_noGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_DISABLED, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should not affect the behavior if the app didn't opt-in for the guard.
    function test_execute_2_2_noGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_2_2_noGuard(indexA, indexB);
    }

    function test_execute_2_2_defaultGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_DEFAULT, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should always revert the transaction if the app opted-in for the guard.
    function test_execute_2_2_defaultGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_DEFAULT,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: defaultGuard, icTx: icTx});
    }

    function test_execute_2_2_customGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 2, guardFlag: GUARD_CUSTOM, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Custom guard conflict should always revert the transaction if the app opted-in for the custom guard.
    function test_execute_2_2_customGuard_customGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(customGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
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
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
    }

    // ════════════════════════════ APP CONFIG: 2 OUT OF 2 RESPONSES (WHEN SET TO ZERO) ════════════════════════════════

    /// @notice If an app sets the required responses to zero, it should default to the amount of trusted modules.
    function test_execute_0_2_noGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 0, guardFlag: GUARD_DISABLED, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should not affect the behavior if the app didn't opt-in for the guard.
    function test_execute_0_2_noGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_0_2_noGuard(indexA, indexB);
    }

    function test_execute_0_2_defaultGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 0, guardFlag: GUARD_DEFAULT, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Guard conflict should always revert the transaction if the app opted-in for the guard.
    function test_execute_0_2_defaultGuard_guardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_DEFAULT,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: defaultGuard, icTx: icTx});
    }

    function test_execute_0_2_customGuard(uint256 indexA, uint256 indexB) public {
        checkScenario({required: 0, guardFlag: GUARD_CUSTOM, modules: modulesAB, timeIndices: toArr(indexA, indexB)});
    }

    /// @dev Custom guard conflict should always revert the transaction if the app opted-in for the custom guard.
    function test_execute_0_2_customGuard_customGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(customGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
    }

    /// @dev Default guard conflict should not affect the behavior if the app opted-in for a custom guard.
    function test_execute_0_2_customGuard_defaultGuardConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        test_execute_0_2_customGuard(indexA, indexB);
    }

    /// @dev Default Guard conflict should be ignored if the app opted-in for a custom guard,
    /// but the custom guard conflict should still revert the transaction.
    function test_execute_0_2_customGuard_bothGuardsConflict(uint256 indexA, uint256 indexB) public {
        addGuardConflict(defaultGuard);
        addGuardConflict(customGuard);
        (InterchainTransaction memory icTx,) = prepareExecuteTest({
            requiredResponses: 2,
            guardFlag: GUARD_CUSTOM,
            modules: modulesAB,
            times: toArr(getTimestampFixture(indexA), getTimestampFixture(indexB))
        });
        checkEntryConflict({module: customGuard, icTx: icTx});
    }

    // ═══════════════════════════════════════════ EXECUTE: MISC REVERTS ═══════════════════════════════════════════════

    function encodeAndMakeExecutable(InterchainTransaction memory icTx) internal returns (bytes memory) {
        bytes memory encodedTx = getEncodedTx(icTx);
        makeTxDescriptorExecutable(getTxDescriptor(icTx));
        return encodedTx;
    }

    function test_execute_revert_TxVersionMismatch(uint16 version) public {
        vm.assume(version != CLIENT_VERSION);
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        bytes memory invalidVersionTx = VersionedPayloadLib.encodeVersionedPayload(version, abi.encode(icTx));
        makeTxDescriptorExecutable(getTxDescriptor(icTx));
        expectRevertTxVersionMismatch(version, CLIENT_VERSION);
        icClient.isExecutable(invalidVersionTx);
        expectRevertTxVersionMismatch(version, CLIENT_VERSION);
        executeTransaction(invalidVersionTx);
    }

    function test_execute_revert_srcChainNotRemote() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.srcChainId = LOCAL_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        icClient.isExecutable(encodedTx);
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        executeTransaction(encodedTx);
    }

    function test_execute_revert_srcChainNotLinked() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.srcChainId = UNKNOWN_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertChainIdNotLinked(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(encodedTx);
        expectRevertChainIdNotLinked(UNKNOWN_CHAIN_ID);
        executeTransaction(encodedTx);
    }

    function test_execute_revert_dstChainIncorrect() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstChainId = UNKNOWN_CHAIN_ID;
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.TxWrongDstChainId, UNKNOWN_CHAIN_ID);
        expectRevertDstChainIdNotLocal(UNKNOWN_CHAIN_ID);
        icClient.isExecutable(encodedTx);
        expectRevertDstChainIdNotLocal(UNKNOWN_CHAIN_ID);
        executeTransaction(encodedTx);
    }

    function test_execute_revert_emptyOptions() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx("");
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        // OptionsLib doesn't have a specific error for this case, so we expect a generic revert during decoding.
        vm.expectRevert();
        icClient.isExecutable(encodedTx);
        vm.expectRevert();
        executeTransaction(encodedTx);
    }

    function test_execute_revert_invalidOptionsV0() public {
        bytes memory invalidOptionsV0 =
            VersionedPayloadLib.encodeVersionedPayload({version: 0, payload: abi.encode(getOptions())});
        (InterchainTransaction memory icTx,) = constructInterchainTx(invalidOptionsV0);
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.UndeterminedRevert);
        expectRevertVersionInvalid(0);
        icClient.isExecutable(encodedTx);
        expectRevertVersionInvalid(0);
        executeTransaction(encodedTx);
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
        icClient.isExecutable(encodedTx);
        vm.expectRevert();
        executeTransaction(encodedTx);
    }

    function test_execute_revert_receiverEOA() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstReceiver = bytes32(uint256(uint160(receiverEOA)));
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.ReceiverNotICApp, receiverEOA);
        expectRevertReceiverNotICApp(receiverEOA);
        icClient.isExecutable(encodedTx);
        expectRevertReceiverNotICApp(receiverEOA);
        executeTransaction(encodedTx);
    }

    function test_execute_revert_receiverNotICApp() public {
        (InterchainTransaction memory icTx,) = constructInterchainTx();
        icTx.dstReceiver = bytes32(uint256(uint160(receiverNotICApp)));
        bytes memory encodedTx = encodeAndMakeExecutable(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.ReceiverNotICApp, receiverNotICApp);
        expectRevertReceiverNotICApp(receiverNotICApp);
        icClient.isExecutable(encodedTx);
        expectRevertReceiverNotICApp(receiverNotICApp);
        executeTransaction(encodedTx);
    }

    function test_execute_revert_alreadyExecuted() public {
        (InterchainTransaction memory icTx, InterchainTxDescriptor memory desc) = prepareAlreadyExecutedTest();
        bytes memory encodedTx = getEncodedTx(icTx);
        assertCorrectReadiness(icTx, IInterchainClientV1.TxReadiness.AlreadyExecuted, uint256(desc.transactionId));
        expectRevertTxAlreadyExecuted(desc.transactionId);
        icClient.isExecutable(encodedTx);
        expectRevertTxAlreadyExecuted(desc.transactionId);
        executeTransaction(encodedTx);
    }
}
