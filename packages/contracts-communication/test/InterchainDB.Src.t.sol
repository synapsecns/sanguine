// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainDB,
    InterchainBatch,
    InterchainEntry,
    InterchainEntryLib,
    IInterchainDB,
    InterchainDBEvents,
    BatchingV1Lib
} from "../contracts/InterchainDB.sol";

import {InterchainBatchLibHarness} from "./harnesses/InterchainBatchLibHarness.sol";
import {VersionedPayloadLibHarness} from "./harnesses/VersionedPayloadLibHarness.sol";
import {InterchainModuleMock, IInterchainModule} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable custom-errors
// solhint-disable func-name-mixedcase
// solhint-disable ordering
/// @notice Unit tests for InterchainDB interactions on the source chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBSourceTest is Test, InterchainDBEvents {
    uint64 public constant SRC_CHAIN_ID = 1337;
    uint64 public constant DST_CHAIN_ID = 7331;

    uint16 public constant DB_VERSION = 1;

    uint64 public constant INITIAL_WRITER_F = 1;
    uint64 public constant INITIAL_WRITER_S = 2;
    uint64 public constant INITIAL_DB_NONCE = INITIAL_WRITER_F + INITIAL_WRITER_S;

    uint256 public constant MODULE_A_FEE = 100;
    uint256 public constant MODULE_B_FEE = 200;

    InterchainBatchLibHarness public batchLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

    InterchainEntry[] public initialEntries;

    address[] public oneModule;
    address[] public twoModules;

    address public requestCaller = makeAddr("Request Caller");
    address public writerF = makeAddr("First Writer");
    address public writerS = makeAddr("Second Writer");
    address public notWriter = makeAddr("Not a Writer");

    function setUp() public {
        vm.chainId(SRC_CHAIN_ID);
        icDB = new InterchainDB();
        moduleA = new InterchainModuleMock();
        moduleB = new InterchainModuleMock();
        batchLibHarness = new InterchainBatchLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        oneModule.push(address(moduleA));
        twoModules.push(address(moduleA));
        twoModules.push(address(moduleB));
        initialWrites();
        mockModuleFee(moduleA, MODULE_A_FEE);
        mockModuleFee(moduleB, MODULE_B_FEE);
    }

    function initialWrites() internal {
        for (uint64 i = 0; i < INITIAL_WRITER_F; ++i) {
            InterchainEntry memory entry = getMockEntry(i, writerF);
            writeEntry(writerF, entry.dataHash);
            initialEntries.push(entry);
        }
        for (uint64 i = 0; i < INITIAL_WRITER_S; ++i) {
            InterchainEntry memory entry = getMockEntry(INITIAL_WRITER_F + i, writerS);
            writeEntry(writerS, entry.dataHash);
            initialEntries.push(entry);
        }
    }

    function getInitialEntry(uint64 dbNonce) internal view returns (InterchainEntry memory) {
        require(dbNonce < initialEntries.length, "dbNonce out of range");
        return initialEntries[dbNonce];
    }

    function getMockDataHash(address writer, uint64 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(uint64 dbNonce, address writer) internal pure returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            dbNonce: dbNonce,
            // TODO: entryIndex
            entryIndex: 0,
            srcWriter: addressToBytes32(writer),
            dataHash: getMockDataHash(writer, dbNonce)
        });
    }

    function getExpectedBatch(InterchainEntry memory entry) internal pure returns (InterchainBatch memory) {
        return InterchainBatch({
            srcChainId: entry.srcChainId,
            dbNonce: entry.dbNonce,
            batchRoot: InterchainEntryLib.entryValue(entry)
        });
    }

    function getVersionedBatch(InterchainBatch memory batch) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(DB_VERSION, batchLibHarness.encodeBatch(batch));
    }

    function getModuleCalldata(InterchainBatch memory batch) internal view returns (bytes memory) {
        bytes memory versionedBatch = getVersionedBatch(batch);
        return abi.encodeCall(IInterchainModule.requestBatchVerification, (DST_CHAIN_ID, versionedBatch));
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    /// @dev Mocks a return value of module.getModuleFee(DST_CHAIN_ID, *)
    function mockModuleFee(InterchainModuleMock module, uint256 feeValue) internal {
        // Encode partial calldata so that we can mock the return value for any dbNonce
        bytes memory callData = abi.encodeWithSelector(module.getModuleFee.selector, DST_CHAIN_ID);
        bytes memory returnData = abi.encode(feeValue);
        vm.mockCall(address(module), callData, returnData);
    }

    function writeEntry(address writer, bytes32 dataHash) internal returns (uint64 dbNonce) {
        vm.prank(writer);
        (dbNonce,) = icDB.writeEntry(dataHash);
    }

    function requestVerification(address caller, uint256 msgValue, uint64 dbNonce, address[] memory modules) internal {
        deal(caller, msgValue);
        vm.prank(caller);
        icDB.requestBatchVerification{value: msgValue}(DST_CHAIN_ID, dbNonce, modules);
    }

    function writeEntryWithVerification(
        uint256 msgValue,
        address writer,
        bytes32 dataHash,
        address[] memory modules
    )
        internal
        returns (uint64 dbNonce)
    {
        deal(writer, msgValue);
        vm.prank(writer);
        (dbNonce,) = icDB.writeEntryWithVerification{value: msgValue}(DST_CHAIN_ID, dataHash, modules);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectValue(bytes32 entryValue, InterchainEntry memory expected) internal pure {
        bytes32 expectedValue = keccak256(abi.encode(expected.srcWriter, expected.dataHash));
        assertEq(entryValue, expectedValue, "!entryValue");
    }

    function assertEq(InterchainBatch memory actual, InterchainBatch memory expected) internal pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.batchRoot, expected.batchRoot, "!batchRoot");
    }

    function expectEventInterchainEntryWritten(InterchainEntry memory entry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten({
            dbNonce: entry.dbNonce,
            entryIndex: entry.entryIndex,
            srcWriter: entry.srcWriter,
            dataHash: entry.dataHash
        });
        // In the V1 of InterchainDB, the batch is finalized immediately after writing an entry
        vm.expectEmit(address(icDB));
        emit InterchainBatchFinalized({dbNonce: entry.dbNonce, batchRoot: entry.entryValue()});
    }

    function expectEventBatchVerificationRequested(
        InterchainBatch memory batch,
        address[] memory srcModules
    )
        internal
    {
        vm.expectEmit(address(icDB));
        emit InterchainBatchVerificationRequested(DST_CHAIN_ID, batch.dbNonce, batch.batchRoot, srcModules);
    }

    function expectRevertEntryIndexOutOfRange(uint64 dbNonce, uint64 entryIndex, uint64 batchSize) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainDB.InterchainDB__EntryIndexOutOfRange.selector, dbNonce, entryIndex, batchSize
            )
        );
    }

    function expectRevertEntryIndexNotZero(uint64 entryIndex) internal {
        vm.expectRevert(abi.encodeWithSelector(BatchingV1Lib.BatchingV1__EntryIndexNotZero.selector, entryIndex));
    }

    function expectRevertFeeAmountBelowMin(uint256 feeAmount, uint256 minRequired) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__FeeAmountBelowMin.selector, feeAmount, minRequired)
        );
    }

    function expectRevertProofNotEmpty() internal {
        vm.expectRevert(BatchingV1Lib.BatchingV1__ProofNotEmpty.selector);
    }

    function expectRevertEntryRangeInvalid(uint64 dbNonce, uint64 start, uint64 end) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__EntryRangeInvalid.selector, dbNonce, start, end)
        );
    }

    function expectRevertModulesNotProvided() internal {
        vm.expectRevert(IInterchainDB.InterchainDB__ModulesNotProvided.selector);
    }

    function expectRevertChainIdNotRemote(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__ChainIdNotRemote.selector, chainId));
    }

    // ═══════════════════════════════════════════════ TESTS: SET UP ═══════════════════════════════════════════════════

    function checkCorrectDBNonceEntryIndex(uint64 expectedDBNonce) internal view {
        assertEq(icDB.getDBNonce(), expectedDBNonce);
        (uint64 dbNonce, uint64 entryIndex) = icDB.getNextEntryIndex();
        assertEq(dbNonce, expectedDBNonce);
        assertEq(entryIndex, 0);
    }

    function test_setup_getters() public view {
        for (uint64 i = 0; i < INITIAL_DB_NONCE; ++i) {
            assertCorrectValue(icDB.getEntryValue(i, 0), getInitialEntry(i));
        }
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE);
    }

    // ══════════════════════════════════════════ TESTS: WRITING AN ENTRY ══════════════════════════════════════════════

    function test_writeEntry_writerF_emitsEvent() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        expectEventInterchainEntryWritten(entry);
        writeEntry(writerF, entry.dataHash);
    }

    function test_writeEntry_writerF_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntry(writerF, dataHash);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntry_writerF_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint64 nonce = writeEntry(writerF, dataHash);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntry_writerF_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntry(writerF, entry.dataHash);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntry_writerS_emitsEvent() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        expectEventInterchainEntryWritten(entry);
        writeEntry(writerS, entry.dataHash);
    }

    function test_writeEntry_writerS_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntry(writerS, dataHash);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntry_writerS_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint64 nonce = writeEntry(writerS, dataHash);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntry_writerS_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntry(writerS, entry.dataHash);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    // ═══════════════════════════════════════ TESTS: REQUESTING VALIDATION ════════════════════════════════════════════

    function test_requestVerification_writerF_oneModule() public {
        uint64 dbNonce = 0;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_oneModule_higherFee() public {
        uint64 dbNonce = 0;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_twoModules() public {
        uint64 dbNonce = 0;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerF_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 0;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_oneModule() public {
        uint64 dbNonce = 2;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_oneModule_higherFee() public {
        uint64 dbNonce = 2;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_twoModules() public {
        uint64 dbNonce = 2;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 2;
        InterchainBatch memory batch = getExpectedBatch(getInitialEntry(dbNonce));
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_nextNonce_oneModule() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_nextNonce_oneModule_higherFee() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_nextNonce_twoModules() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_nextNonce_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_hugeNonce_oneModule() public {
        uint64 dbNonce = 2 ** 32;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_hugeNonce_oneModule_higherFee() public {
        uint64 dbNonce = 2 ** 32;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_hugeNonce_twoModules() public {
        uint64 dbNonce = 2 ** 32;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_hugeNonce_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 2 ** 32;
        InterchainBatch memory batch = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, batchRoot: 0});
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventBatchVerificationRequested(batch, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    // ══════════════════════════════════ TESTS: REQUESTING VALIDATION (REVERTS) ═══════════════════════════════════════

    function test_requestVerification_revert_FeeAmountBelowMin_oneModule_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE);
        requestVerification(requestCaller, incorrectFee, 0, oneModule);
    }

    function test_requestVerification_revert_FeeAmountBelowMin_twoModules_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        requestVerification(requestCaller, incorrectFee, 0, twoModules);
    }

    function test_requestVerification_revert_ModulesNotProvided() public {
        expectRevertModulesNotProvided();
        requestVerification(requestCaller, MODULE_A_FEE, 0, new address[](0));
    }

    function test_requestVerification_revert_ChainIdNotRemote() public {
        expectRevertChainIdNotRemote(SRC_CHAIN_ID);
        vm.prank(requestCaller);
        icDB.requestBatchVerification(SRC_CHAIN_ID, 0, oneModule);
    }

    // ═════════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION ═════════════════════════════════════════

    function test_writeEntryWithVerification_writerF_oneModule_callsModule() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        writeEntryWithVerification(MODULE_A_FEE, writerF, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_emitsEvents() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        expectEventInterchainEntryWritten(entry);
        expectEventBatchVerificationRequested(getExpectedBatch(entry), oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerF, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint64 nonce = writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerF_oneModule_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntryWithVerification(MODULE_A_FEE, writerF, entry.dataHash, oneModule);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerF_oneModule_higherFee() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventInterchainEntryWritten(entry);
        writeEntryWithVerification(MODULE_A_FEE * 2, writerF, dataHash, oneModule);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerF_twoModules_callsModules() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_emitsEvents() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        expectEventInterchainEntryWritten(entry);
        expectEventBatchVerificationRequested(getExpectedBatch(entry), twoModules);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint64 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerF_twoModules_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, entry.dataHash, twoModules);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerF_twoModules_higherFee() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventInterchainEntryWritten(entry);
        writeEntryWithVerification(MODULE_A_FEE * 2 + MODULE_B_FEE, writerF, dataHash, twoModules);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerS_oneModule_callsModule() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        writeEntryWithVerification(MODULE_A_FEE, writerS, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_emitsEvents() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        expectEventInterchainEntryWritten(entry);
        expectEventBatchVerificationRequested(getExpectedBatch(entry), oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerS, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint64 nonce = writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerS_oneModule_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntryWithVerification(MODULE_A_FEE, writerS, entry.dataHash, oneModule);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerS_oneModule_higherFee() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        expectEventInterchainEntryWritten(entry);
        writeEntryWithVerification(MODULE_A_FEE * 2, writerS, dataHash, oneModule);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerS_twoModules_callsModules() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_emitsEvents() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        expectEventInterchainEntryWritten(entry);
        expectEventBatchVerificationRequested(getExpectedBatch(entry), twoModules);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint64 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerS_twoModules_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, entry.dataHash, twoModules);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    function test_writeEntryWithVerification_writerS_twoModules_higherFee() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        InterchainBatch memory batch = getExpectedBatch(entry);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(batch)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(batch)});
        expectEventInterchainEntryWritten(entry);
        writeEntryWithVerification(MODULE_A_FEE * 2 + MODULE_B_FEE, writerS, dataHash, twoModules);
        checkCorrectDBNonceEntryIndex(INITIAL_DB_NONCE + 1);
        assertCorrectValue(icDB.getEntryValue(INITIAL_DB_NONCE, 0), entry);
    }

    // ════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION (REVERTS) ════════════════════════════════════

    function test_writeEntryWithVerification_revert_FeeAmountBelowMin_oneModule_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_revert_FeeAmountBelowMin_twoModules_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_revert_ModulesNotProvided() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        expectRevertModulesNotProvided();
        writeEntryWithVerification(0, writerF, dataHash, new address[](0));
    }

    function test_writeEntryWithVerification_revert_ChainIdNotRemote() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        expectRevertChainIdNotRemote(SRC_CHAIN_ID);
        vm.prank(writerF);
        icDB.writeEntryWithVerification(SRC_CHAIN_ID, dataHash, oneModule);
    }

    // ═════════════════════════════════════════ TESTS: GET INTERCHAIN FEE ═════════════════════════════════════════════

    function test_getInterchainFee_noModules() public {
        expectRevertModulesNotProvided();
        icDB.getInterchainFee(DST_CHAIN_ID, new address[](0));
    }

    function test_getInterchainFee_oneModule() public view {
        // [moduleA]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, oneModule), MODULE_A_FEE);
    }

    function test_getInterchainFee_twoModules() public view {
        // [moduleA, moduleB]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, twoModules), MODULE_A_FEE + MODULE_B_FEE);
    }

    // ════════════════════════════════════════ TESTS: RETRIEVING DB VALUES ════════════════════════════════════════════

    function test_getBatchLeafs_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            bytes32[] memory leafs = icDB.getBatchLeafs(nonce);
            assertEq(leafs.length, 1, "!leafs.length");
            assertEq(leafs[0], getInitialEntry(nonce).entryValue());
        }
    }

    function test_getBatchLeafs_nextNonce() public view {
        bytes32[] memory leafs = icDB.getBatchLeafs(INITIAL_DB_NONCE);
        assertEq(leafs.length, 0);
    }

    function test_getBatchLeafs_hugeNonce() public view {
        bytes32[] memory leafs = icDB.getBatchLeafs(2 ** 32);
        assertEq(leafs.length, 0);
    }

    function test_getBatchLeafsPaginated_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            bytes32[] memory leafs = icDB.getBatchLeafsPaginated(nonce, 0, 0);
            assertEq(leafs.length, 0, "!leafs.length");
            leafs = icDB.getBatchLeafsPaginated(nonce, 1, 1);
            assertEq(leafs.length, 0, "!leafs.length");
            leafs = icDB.getBatchLeafsPaginated(nonce, 0, 1);
            assertEq(leafs.length, 1, "!leafs.length");
            assertEq(leafs[0], getInitialEntry(nonce).entryValue());
        }
    }

    function test_getBatchLeafsPaginated_nextNonce() public view {
        bytes32[] memory leafs = icDB.getBatchLeafsPaginated(INITIAL_DB_NONCE, 0, 0);
        assertEq(leafs.length, 0);
    }

    function test_getBatchLeafsPaginated_hugeNonce() public view {
        bytes32[] memory leafs = icDB.getBatchLeafsPaginated(2 ** 32, 0, 0);
        assertEq(leafs.length, 0);
    }

    function test_getBatchLeafsPaginated_revert_invalidRange_finalized() public {
        expectRevertEntryRangeInvalid(1, 0, 2);
        icDB.getBatchLeafsPaginated(1, 0, 2);
        expectRevertEntryRangeInvalid(3, 1, 0);
        icDB.getBatchLeafsPaginated(3, 1, 0);
    }

    function test_getBatchLeafsPaginated_revert_invalidRange_nextNonce() public {
        expectRevertEntryRangeInvalid(INITIAL_DB_NONCE, 0, 1);
        icDB.getBatchLeafsPaginated(INITIAL_DB_NONCE, 0, 1);
        expectRevertEntryRangeInvalid(INITIAL_DB_NONCE, 1, 1);
        icDB.getBatchLeafsPaginated(INITIAL_DB_NONCE, 1, 1);
        expectRevertEntryRangeInvalid(INITIAL_DB_NONCE, 1, 0);
        icDB.getBatchLeafsPaginated(INITIAL_DB_NONCE, 1, 0);
    }

    function test_getBatchLeafsPaginated_revert_invalidRange_hugeNonce() public {
        expectRevertEntryRangeInvalid(2 ** 32, 0, 1);
        icDB.getBatchLeafsPaginated(2 ** 32, 0, 1);
        expectRevertEntryRangeInvalid(2 ** 32, 1, 1);
        icDB.getBatchLeafsPaginated(2 ** 32, 1, 1);
        expectRevertEntryRangeInvalid(2 ** 32, 1, 0);
        icDB.getBatchLeafsPaginated(2 ** 32, 1, 0);
    }

    function test_getBatchSize_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            assertEq(icDB.getBatchSize(nonce), 1, "!batchSize");
        }
    }

    function test_getBatchSize_nextNonce() public view {
        assertEq(icDB.getBatchSize(INITIAL_DB_NONCE), 0);
    }

    function test_getBatchSize_hugeNonce() public view {
        assertEq(icDB.getBatchSize(2 ** 32), 0);
    }

    function test_getBatch_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            InterchainBatch memory batch = icDB.getBatch(nonce);
            assertEq(batch, getExpectedBatch(getInitialEntry(nonce)));
        }
    }

    function test_getBatch_nextNonce() public view {
        InterchainBatch memory batch = icDB.getBatch(INITIAL_DB_NONCE);
        InterchainBatch memory expected =
            InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: INITIAL_DB_NONCE, batchRoot: 0});
        assertEq(batch, expected);
    }

    function test_getBatch_hugeNonce() public view {
        InterchainBatch memory batch = icDB.getBatch(2 ** 32);
        InterchainBatch memory expected = InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: 2 ** 32, batchRoot: 0});
        assertEq(batch, expected);
    }

    function test_getVersionedBatch_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            bytes memory versionedBatch = icDB.getVersionedBatch(nonce);
            InterchainBatch memory expectedBatch = getExpectedBatch(getInitialEntry(nonce));
            assertEq(versionedBatch, getVersionedBatch(expectedBatch));
        }
    }

    function test_getVersionedBatch_nextNonce() public view {
        bytes memory versionedBatch = icDB.getVersionedBatch(INITIAL_DB_NONCE);
        InterchainBatch memory expectedBatch =
            InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: INITIAL_DB_NONCE, batchRoot: 0});
        assertEq(versionedBatch, getVersionedBatch(expectedBatch));
    }

    function test_getVersionedBatch_hugeNonce() public view {
        bytes memory versionedBatch = icDB.getVersionedBatch(2 ** 32);
        InterchainBatch memory expectedBatch =
            InterchainBatch({srcChainId: SRC_CHAIN_ID, dbNonce: 2 ** 32, batchRoot: 0});
        assertEq(versionedBatch, getVersionedBatch(expectedBatch));
    }

    function test_getEntryValue() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            InterchainEntry memory expectedEntry = getInitialEntry(nonce);
            assertCorrectValue(icDB.getEntryValue(nonce, 0), expectedEntry);
        }
    }

    function test_getEntryValue_revert_finalizedOutOfRange() public {
        expectRevertEntryIndexOutOfRange(INITIAL_DB_NONCE - 1, 1, 1);
        icDB.getEntryValue(INITIAL_DB_NONCE - 1, 1);
    }

    function test_getEntryValue_revert_nextNonceOutOfRange() public {
        expectRevertEntryIndexOutOfRange(INITIAL_DB_NONCE, 0, 0);
        icDB.getEntryValue(INITIAL_DB_NONCE, 0);
    }

    function test_getEntryValue_revert_hugeNonceOutOfRange() public {
        expectRevertEntryIndexOutOfRange(2 ** 32, 0, 0);
        icDB.getEntryValue(2 ** 32, 0);
    }

    function test_getEntryProof_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            bytes32[] memory proof = icDB.getEntryProof(nonce, 0);
            assertEq(proof.length, 0, "!proof.length");
        }
    }

    function test_getEntryProof_revert_finalizedOutOfRange() public {
        expectRevertEntryIndexOutOfRange(INITIAL_DB_NONCE - 1, 1, 1);
        icDB.getEntryProof(INITIAL_DB_NONCE - 1, 1);
    }

    function test_getEntryProof_revert_nextNonceOutOfRange() public {
        expectRevertEntryIndexOutOfRange(INITIAL_DB_NONCE, 0, 0);
        icDB.getEntryProof(INITIAL_DB_NONCE, 0);
    }

    function test_getEntryProof_revert_hugeNonceOutOfRange() public {
        expectRevertEntryIndexOutOfRange(2 ** 32, 0, 0);
        icDB.getEntryProof(2 ** 32, 0);
    }

    // ═══════════════════════════════════════════ TESTS: GET BATCH ROOT ═══════════════════════════════════════════════

    function test_getBatchRoot(InterchainEntry memory entry) public view {
        entry.entryIndex = 0;
        bytes32 batchRoot = icDB.getBatchRoot(entry, new bytes32[](0));
        assertEq(batchRoot, InterchainEntryLib.entryValue(entry));
    }

    function test_getBatchRoot_revert_nonZeroEntryIndex() public {
        InterchainEntry memory entry = getInitialEntry(0);
        entry.entryIndex = 1;
        expectRevertEntryIndexNotZero(1);
        icDB.getBatchRoot(entry, new bytes32[](0));
    }

    function test_getBatchRoot_revert_nonEmptyProof() public {
        InterchainEntry memory entry = getInitialEntry(0);
        expectRevertProofNotEmpty();
        icDB.getBatchRoot(entry, new bytes32[](1));
    }
}
