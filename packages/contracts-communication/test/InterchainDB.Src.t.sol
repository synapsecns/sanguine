// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, InterchainEntry, IInterchainDB, InterchainDBEvents} from "../contracts/InterchainDB.sol";

import {InterchainModuleMock, IInterchainModule} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Unit tests for InterchainDB interactions on the source chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBSourceTest is Test, InterchainDBEvents {
    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    uint256 public constant INITIAL_WRITER_F = 1;
    uint256 public constant INITIAL_WRITER_S = 2;
    uint256 public constant INITIAL_DB_NONCE = INITIAL_WRITER_F + INITIAL_WRITER_S;

    uint256 public constant MODULE_A_FEE = 100;
    uint256 public constant MODULE_B_FEE = 200;

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
        oneModule.push(address(moduleA));
        twoModules.push(address(moduleA));
        twoModules.push(address(moduleB));
        initialWrites();
        mockModuleFee(moduleA, MODULE_A_FEE);
        mockModuleFee(moduleB, MODULE_B_FEE);
    }

    function initialWrites() internal {
        for (uint256 i = 0; i < INITIAL_WRITER_F; ++i) {
            InterchainEntry memory entry = getMockEntry(i, writerF);
            writeEntry(writerF, entry.dataHash);
            initialEntries.push(entry);
        }
        for (uint256 i = 0; i < INITIAL_WRITER_S; ++i) {
            InterchainEntry memory entry = getMockEntry(INITIAL_WRITER_F + i, writerS);
            writeEntry(writerS, entry.dataHash);
            initialEntries.push(entry);
        }
    }

    function getInitialEntry(uint256 dbNonce) internal view returns (InterchainEntry memory) {
        require(dbNonce < initialEntries.length, "dbNonce out of range");
        return initialEntries[dbNonce];
    }

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(uint256 dbNonce, address writer) internal pure returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            dbNonce: dbNonce,
            srcWriter: addressToBytes32(writer),
            dataHash: getMockDataHash(writer, dbNonce)
        });
    }

    function getModuleCalldata(InterchainEntry memory entry) internal pure returns (bytes memory) {
        return abi.encodeCall(IInterchainModule.requestVerification, (DST_CHAIN_ID, entry));
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    /// @dev Mocks a return value of module.getModuleFee(DST_CHAIN_ID)
    function mockModuleFee(InterchainModuleMock module, uint256 feeValue) internal {
        bytes memory callData = abi.encodeCall(module.getModuleFee, (DST_CHAIN_ID));
        bytes memory returnData = abi.encode(feeValue);
        vm.mockCall(address(module), callData, returnData);
    }

    function writeEntry(address writer, bytes32 dataHash) internal returns (uint256 dbNonce) {
        vm.prank(writer);
        dbNonce = icDB.writeEntry(dataHash);
    }

    function requestVerification(
        address caller,
        uint256 msgValue,
        uint256 dbNonce,
        address[] memory modules
    )
        internal
    {
        deal(caller, msgValue);
        vm.prank(caller);
        icDB.requestVerification{value: msgValue}(DST_CHAIN_ID, dbNonce, modules);
    }

    function writeEntryWithVerification(
        uint256 msgValue,
        address writer,
        bytes32 dataHash,
        address[] memory modules
    )
        internal
        returns (uint256 dbNonce)
    {
        deal(writer, msgValue);
        vm.prank(writer);
        dbNonce = icDB.writeEntryWithVerification{value: msgValue}(DST_CHAIN_ID, dataHash, modules);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertEq(InterchainEntry memory entry, InterchainEntry memory expected) internal {
        assertEq(entry.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(entry.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(entry.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(entry.dataHash, expected.dataHash, "!dataHash");
    }

    function expectEntryDoesNotExist(uint256 dbNonce) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__EntryDoesNotExist.selector, dbNonce));
    }

    function expectIncorrectFeeAmount(uint256 actualFee, uint256 expectedFee) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__IncorrectFeeAmount.selector, actualFee, expectedFee)
        );
    }

    function expectNoModulesSpecified() internal {
        vm.expectRevert(IInterchainDB.InterchainDB__NoModulesSpecified.selector);
    }

    function expectSameChainId() internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__SameChainId.selector));
    }

    // ═══════════════════════════════════════════════ TESTS: SET UP ═══════════════════════════════════════════════════

    function test_setup_getDBNonce() public {
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE);
    }

    function test_setup_getEntry() public {
        for (uint256 i = 0; i < INITIAL_DB_NONCE; ++i) {
            assertEq(icDB.getEntry(i), getInitialEntry(i));
        }
    }

    // ══════════════════════════════════════════ TESTS: WRITING AN ENTRY ══════════════════════════════════════════════

    function test_writeEntry_writerF_emitsEvent() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerF), dataHash);
        writeEntry(writerF, dataHash);
    }

    function test_writeEntry_writerF_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntry(writerF, dataHash);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntry_writerF_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 nonce = writeEntry(writerF, dataHash);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntry_writerF_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntry(writerF, entry.dataHash);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    function test_writeEntry_writerS_emitsEvent() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerS), dataHash);
        writeEntry(writerS, dataHash);
    }

    function test_writeEntry_writerS_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntry(writerS, dataHash);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntry_writerS_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint256 nonce = writeEntry(writerS, dataHash);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntry_writerS_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntry(writerS, entry.dataHash);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    // ═══════════════════════════════════════ TESTS: REQUESTING VALIDATION ════════════════════════════════════════════

    function test_requestVerification_writerF_oneModule_emitsEvent() public {
        uint256 dbNonce = 0;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_oneModule_callsModule() public {
        uint256 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_twoModules_emitsEvent() public {
        uint256 dbNonce = 0;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerF_twoModules_callsModules() public {
        uint256 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_oneModule_emitsEvent() public {
        uint256 dbNonce = 2;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_oneModule_callsModule() public {
        uint256 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_twoModules_emitsEvent() public {
        uint256 dbNonce = 2;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_twoModules_callsModules() public {
        uint256 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    // ══════════════════════════════════ TESTS: REQUESTING VALIDATION (REVERTS) ═══════════════════════════════════════

    function test_requestVerification_revert_entryDoesNotExist_oneModule_nextNonce() public {
        expectEntryDoesNotExist(INITIAL_DB_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE, INITIAL_DB_NONCE, oneModule);
    }

    function test_requestVerification_revert_entryDoesNotExist_oneModule_hugeNonce() public {
        expectEntryDoesNotExist(2 ** 32);
        requestVerification(requestCaller, MODULE_A_FEE, 2 ** 32, oneModule);
        expectEntryDoesNotExist(2 ** 64);
        requestVerification(requestCaller, MODULE_A_FEE, 2 ** 64, oneModule);
        expectEntryDoesNotExist(type(uint256).max);
        requestVerification(requestCaller, MODULE_A_FEE, type(uint256).max, oneModule);
    }

    function test_requestVerification_revert_entryDoesNotExist_twoModules_nextNonce() public {
        expectEntryDoesNotExist(INITIAL_DB_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, INITIAL_DB_NONCE, twoModules);
    }

    function test_requestVerification_revert_entryDoesNotExist_twoModules_hugeNonce() public {
        expectEntryDoesNotExist(2 ** 32);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, 2 ** 32, twoModules);
        expectEntryDoesNotExist(2 ** 64);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, 2 ** 64, twoModules);
        expectEntryDoesNotExist(type(uint256).max);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, type(uint256).max, twoModules);
    }

    function test_requestVerification_revert_incorrectFeeAmount_oneModule_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        requestVerification(requestCaller, incorrectFee, 0, oneModule);
    }

    function test_requestVerification_revert_incorrectFeeAmount_oneModule_overpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        requestVerification(requestCaller, incorrectFee, 0, oneModule);
    }

    function test_requestVerification_revert_incorrectFeeAmount_twoModules_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        requestVerification(requestCaller, incorrectFee, 0, twoModules);
    }

    function test_requestVerification_revert_incorrectFeeAmount_twoModules_overpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        requestVerification(requestCaller, incorrectFee, 0, twoModules);
    }

    function test_requestVerification_revert_noModulesSpecified() public {
        expectNoModulesSpecified();
        requestVerification(requestCaller, MODULE_A_FEE, 0, new address[](0));
    }

    function test_requestVerification_revert_sameChainId() public {
        expectSameChainId();
        vm.prank(requestCaller);
        icDB.requestVerification(SRC_CHAIN_ID, 0, oneModule);
    }

    // ═════════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION ═════════════════════════════════════════

    function test_writeEntryWithVerification_writerF_oneModule_callsModule() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE, writerF, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerF), dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, INITIAL_DB_NONCE, oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerF_oneModule_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntryWithVerification(MODULE_A_FEE, writerF, entry.dataHash, oneModule);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    function test_writeEntryWithVerification_writerF_twoModules_callsModules() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerF), dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, INITIAL_DB_NONCE, twoModules);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerF_twoModules_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, entry.dataHash, twoModules);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    function test_writeEntryWithVerification_writerS_oneModule_callsModule() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE, writerS, entry.dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerS), dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, INITIAL_DB_NONCE, oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerS_oneModule_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntryWithVerification(MODULE_A_FEE, writerS, entry.dataHash, oneModule);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    function test_writeEntryWithVerification_writerS_twoModules_callsModules() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, entry.dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, INITIAL_DB_NONCE, addressToBytes32(writerS), dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, INITIAL_DB_NONCE, twoModules);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_increasesDBNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_DB_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
    }

    function test_writeEntryWithVerification_writerS_twoModules_savesEntry() public {
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, entry.dataHash, twoModules);
        assertEq(icDB.getEntry(INITIAL_DB_NONCE), entry);
    }

    // ════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION (REVERTS) ════════════════════════════════════

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_oneModule_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_oneModule_overpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_twoModules_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_twoModules_overpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_revert_noModulesSpecified() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        expectNoModulesSpecified();
        writeEntryWithVerification(0, writerF, dataHash, new address[](0));
    }

    function test_writeEntryWithVerification_revert_sameChainId() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_DB_NONCE);
        expectSameChainId();
        vm.prank(writerF);
        icDB.writeEntryWithVerification(SRC_CHAIN_ID, dataHash, oneModule);
    }

    // ═════════════════════════════════════════ TESTS: GET INTERCHAIN FEE ═════════════════════════════════════════════

    function test_getInterchainFee_noModules() public {
        expectNoModulesSpecified();
        icDB.getInterchainFee(DST_CHAIN_ID, new address[](0));
    }

    function test_getInterchainFee_oneModule() public {
        // [moduleA]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, oneModule), MODULE_A_FEE);
    }

    function test_getInterchainFee_twoModules() public {
        // [moduleA, moduleB]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, twoModules), MODULE_A_FEE + MODULE_B_FEE);
    }
}
