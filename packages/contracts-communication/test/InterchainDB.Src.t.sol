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

    uint256 public constant INITIAL_WRITER_F_NONCE = 1;
    uint256 public constant INITIAL_WRITER_S_NONCE = 2;

    uint256 public constant MODULE_A_FEE = 100;
    uint256 public constant MODULE_B_FEE = 200;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

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
        setupWriterNonce(writerF, INITIAL_WRITER_F_NONCE);
        setupWriterNonce(writerS, INITIAL_WRITER_S_NONCE);
        mockModuleFee(moduleA, MODULE_A_FEE);
        mockModuleFee(moduleB, MODULE_B_FEE);
    }

    function setupWriterNonce(address writer, uint256 nonce) internal {
        for (uint256 i = 0; i < nonce; i++) {
            writeEntry(writer, getMockDataHash(writer, i));
        }
    }

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(address writer, uint256 nonce) internal pure returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            srcWriter: addressToBytes32(writer),
            writerNonce: nonce,
            dataHash: getMockDataHash(writer, nonce)
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

    function writeEntry(address writer, bytes32 dataHash) internal returns (uint256 writerNonce) {
        vm.prank(writer);
        writerNonce = icDB.writeEntry(dataHash);
    }

    function requestVerification(
        address caller,
        uint256 msgValue,
        address writer,
        uint256 writerNonce,
        address[] memory modules
    )
        internal
    {
        deal(caller, msgValue);
        vm.prank(caller);
        icDB.requestVerification{value: msgValue}(DST_CHAIN_ID, writer, writerNonce, modules);
    }

    function writeEntryWithVerification(
        uint256 msgValue,
        address writer,
        bytes32 dataHash,
        address[] memory modules
    )
        internal
        returns (uint256 writerNonce)
    {
        deal(writer, msgValue);
        vm.prank(writer);
        writerNonce = icDB.writeEntryWithVerification{value: msgValue}(DST_CHAIN_ID, dataHash, modules);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertEq(InterchainEntry memory entry, InterchainEntry memory expected) internal {
        assertEq(entry.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(entry.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(entry.writerNonce, expected.writerNonce, "!writerNonce");
        assertEq(entry.dataHash, expected.dataHash, "!dataHash");
    }

    function expectEntryDoesNotExist(address writer, uint256 writerNonce) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__EntryDoesNotExist.selector, writer, writerNonce)
        );
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

    function test_setup_getWriterNonce() public {
        assertEq(icDB.getWriterNonce(writerF), INITIAL_WRITER_F_NONCE);
        assertEq(icDB.getWriterNonce(writerS), INITIAL_WRITER_S_NONCE);
    }

    function test_setup_getEntry() public {
        for (uint256 i = 0; i < INITIAL_WRITER_F_NONCE; i++) {
            assertEq(icDB.getEntry(writerF, i), getMockEntry(writerF, i));
        }
        for (uint256 i = 0; i < INITIAL_WRITER_S_NONCE; i++) {
            assertEq(icDB.getEntry(writerS, i), getMockEntry(writerS, i));
        }
    }

    // ══════════════════════════════════════════ TESTS: WRITING AN ENTRY ══════════════════════════════════════════════

    function test_writeEntry_writerF_emitsEvent() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerF), INITIAL_WRITER_F_NONCE, dataHash);
        writeEntry(writerF, dataHash);
    }

    function test_writeEntry_writerF_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntry(writerF, dataHash);
        assertEq(icDB.getWriterNonce(writerF), INITIAL_WRITER_F_NONCE + 1);
    }

    function test_writeEntry_writerF_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 nonce = writeEntry(writerF, dataHash);
        assertEq(nonce, INITIAL_WRITER_F_NONCE);
    }

    function test_writeEntry_writerF_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntry(writerF, dataHash);
        assertEq(icDB.getEntry(writerF, INITIAL_WRITER_F_NONCE), getMockEntry(writerF, INITIAL_WRITER_F_NONCE));
    }

    function test_writeEntry_writerS_emitsEvent() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerS), INITIAL_WRITER_S_NONCE, dataHash);
        writeEntry(writerS, dataHash);
    }

    function test_writeEntry_writerS_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntry(writerS, dataHash);
        assertEq(icDB.getWriterNonce(writerS), INITIAL_WRITER_S_NONCE + 1);
    }

    function test_writeEntry_writerS_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        uint256 nonce = writeEntry(writerS, dataHash);
        assertEq(nonce, INITIAL_WRITER_S_NONCE);
    }

    function test_writeEntry_writerS_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntry(writerS, dataHash);
        assertEq(icDB.getEntry(writerS, INITIAL_WRITER_S_NONCE), getMockEntry(writerS, INITIAL_WRITER_S_NONCE));
    }

    // ═══════════════════════════════════════ TESTS: REQUESTING VALIDATION ════════════════════════════════════════════

    function test_requestVerification_writerF_oneModule_emitsEvent() public {
        uint256 writerNonce = 0;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerF), writerNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, writerF, writerNonce, oneModule);
    }

    function test_requestVerification_writerF_oneModule_callsModule() public {
        uint256 writerNonce = 0;
        InterchainEntry memory entry = getMockEntry(writerF, writerNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE, writerF, writerNonce, oneModule);
    }

    function test_requestVerification_writerF_twoModules_emitsEvent() public {
        uint256 writerNonce = 0;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerF), writerNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerF, writerNonce, twoModules);
    }

    function test_requestVerification_writerF_twoModules_callsModules() public {
        uint256 writerNonce = 0;
        InterchainEntry memory entry = getMockEntry(writerF, writerNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerF, writerNonce, twoModules);
    }

    function test_requestVerification_writerS_oneModule_emitsEvent() public {
        uint256 writerNonce = 1;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerS), writerNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, writerS, writerNonce, oneModule);
    }

    function test_requestVerification_writerS_oneModule_callsModule() public {
        uint256 writerNonce = 1;
        InterchainEntry memory entry = getMockEntry(writerS, writerNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE, writerS, writerNonce, oneModule);
    }

    function test_requestVerification_writerS_twoModules_emitsEvent() public {
        uint256 writerNonce = 1;
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerS), writerNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerS, writerNonce, twoModules);
    }

    function test_requestVerification_writerS_twoModules_callsModules() public {
        uint256 writerNonce = 1;
        InterchainEntry memory entry = getMockEntry(writerS, writerNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerS, writerNonce, twoModules);
    }

    // ══════════════════════════════════ TESTS: REQUESTING VALIDATION (REVERTS) ═══════════════════════════════════════

    function test_requestVerification_revert_entryDoesNotExist_oneModule_nextNonce() public {
        expectEntryDoesNotExist(writerF, INITIAL_WRITER_F_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE, writerF, INITIAL_WRITER_F_NONCE, oneModule);
        expectEntryDoesNotExist(writerS, INITIAL_WRITER_S_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE, writerS, INITIAL_WRITER_S_NONCE, oneModule);
        expectEntryDoesNotExist(notWriter, 0);
        requestVerification(requestCaller, MODULE_A_FEE, notWriter, 0, oneModule);
    }

    function test_requestVerification_revert_entryDoesNotExist_oneModule_hugeNonce() public {
        expectEntryDoesNotExist(writerF, 2 ** 32);
        requestVerification(requestCaller, MODULE_A_FEE, writerF, 2 ** 32, oneModule);
        expectEntryDoesNotExist(writerS, 2 ** 64);
        requestVerification(requestCaller, MODULE_A_FEE, writerS, 2 ** 64, oneModule);
        expectEntryDoesNotExist(notWriter, type(uint256).max);
        requestVerification(requestCaller, MODULE_A_FEE, notWriter, type(uint256).max, oneModule);
    }

    function test_requestVerification_revert_entryDoesNotExist_twoModules_nextNonce() public {
        expectEntryDoesNotExist(writerF, INITIAL_WRITER_F_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerF, INITIAL_WRITER_F_NONCE, twoModules);
        expectEntryDoesNotExist(writerS, INITIAL_WRITER_S_NONCE);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerS, INITIAL_WRITER_S_NONCE, twoModules);
        expectEntryDoesNotExist(notWriter, 0);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, notWriter, 0, twoModules);
    }

    function test_requestVerification_revert_entryDoesNotExist_twoModules_hugeNonce() public {
        expectEntryDoesNotExist(writerF, 2 ** 32);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerF, 2 ** 32, twoModules);
        expectEntryDoesNotExist(writerS, 2 ** 64);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerS, 2 ** 64, twoModules);
        expectEntryDoesNotExist(notWriter, type(uint256).max);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, notWriter, type(uint256).max, twoModules);
    }

    function test_requestVerification_revert_incorrectFeeAmount_oneModule_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        requestVerification(requestCaller, incorrectFee, writerF, 0, oneModule);
    }

    function test_requestVerification_revert_incorrectFeeAmount_oneModule_overpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        requestVerification(requestCaller, incorrectFee, writerF, 0, oneModule);
    }

    function test_requestVerification_revert_incorrectFeeAmount_twoModules_underpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        requestVerification(requestCaller, incorrectFee, writerF, 0, twoModules);
    }

    function test_requestVerification_revert_incorrectFeeAmount_twoModules_overpaid() public {
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        requestVerification(requestCaller, incorrectFee, writerF, 0, twoModules);
    }

    function test_requestVerification_revert_noModulesSpecified() public {
        expectNoModulesSpecified();
        requestVerification(requestCaller, MODULE_A_FEE, writerF, 0, new address[](0));
    }

    function test_requestVerification_revert_sameChainId() public {
        expectSameChainId();
        vm.prank(requestCaller);
        icDB.requestVerification(SRC_CHAIN_ID, writerF, 0, oneModule);
    }

    // ═════════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION ═════════════════════════════════════════

    function test_writeEntryWithVerification_writerF_oneModule_callsModule() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        InterchainEntry memory entry = getMockEntry(writerF, INITIAL_WRITER_F_NONCE);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerF), INITIAL_WRITER_F_NONCE, dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerF), INITIAL_WRITER_F_NONCE, oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerF_oneModule_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(icDB.getWriterNonce(writerF), INITIAL_WRITER_F_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(nonce, INITIAL_WRITER_F_NONCE);
    }

    function test_writeEntryWithVerification_writerF_oneModule_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerF, dataHash, oneModule);
        assertEq(icDB.getEntry(writerF, INITIAL_WRITER_F_NONCE), getMockEntry(writerF, INITIAL_WRITER_F_NONCE));
    }

    function test_writeEntryWithVerification_writerF_twoModules_callsModules() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        InterchainEntry memory entry = getMockEntry(writerF, INITIAL_WRITER_F_NONCE);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerF), INITIAL_WRITER_F_NONCE, dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(
            DST_CHAIN_ID, addressToBytes32(writerF), INITIAL_WRITER_F_NONCE, twoModules
        );
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerF_twoModules_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(icDB.getWriterNonce(writerF), INITIAL_WRITER_F_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerF_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(nonce, INITIAL_WRITER_F_NONCE);
    }

    function test_writeEntryWithVerification_writerF_twoModules_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, dataHash, twoModules);
        assertEq(icDB.getEntry(writerF, INITIAL_WRITER_F_NONCE), getMockEntry(writerF, INITIAL_WRITER_F_NONCE));
    }

    function test_writeEntryWithVerification_writerS_oneModule_callsModule() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        InterchainEntry memory entry = getMockEntry(writerS, INITIAL_WRITER_S_NONCE);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerS), INITIAL_WRITER_S_NONCE, dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(DST_CHAIN_ID, addressToBytes32(writerS), INITIAL_WRITER_S_NONCE, oneModule);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_writerS_oneModule_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(icDB.getWriterNonce(writerS), INITIAL_WRITER_S_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_oneModule_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(nonce, INITIAL_WRITER_S_NONCE);
    }

    function test_writeEntryWithVerification_writerS_oneModule_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntryWithVerification(MODULE_A_FEE, writerS, dataHash, oneModule);
        assertEq(icDB.getEntry(writerS, INITIAL_WRITER_S_NONCE), getMockEntry(writerS, INITIAL_WRITER_S_NONCE));
    }

    function test_writeEntryWithVerification_writerS_twoModules_callsModules() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        InterchainEntry memory entry = getMockEntry(writerS, INITIAL_WRITER_S_NONCE);
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_emitsEvents() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, addressToBytes32(writerS), INITIAL_WRITER_S_NONCE, dataHash);
        vm.expectEmit(address(icDB));
        emit InterchainVerificationRequested(
            DST_CHAIN_ID, addressToBytes32(writerS), INITIAL_WRITER_S_NONCE, twoModules
        );
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_writerS_twoModules_increasesWriterNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(icDB.getWriterNonce(writerS), INITIAL_WRITER_S_NONCE + 1);
    }

    function test_writeEntryWithVerification_writerS_twoModules_returnsCorrectNonce() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        uint256 nonce = writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(nonce, INITIAL_WRITER_S_NONCE);
    }

    function test_writeEntryWithVerification_writerS_twoModules_savesEntry() public {
        bytes32 dataHash = getMockDataHash(writerS, INITIAL_WRITER_S_NONCE);
        writeEntryWithVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, dataHash, twoModules);
        assertEq(icDB.getEntry(writerS, INITIAL_WRITER_S_NONCE), getMockEntry(writerS, INITIAL_WRITER_S_NONCE));
    }

    // ════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION (REVERTS) ════════════════════════════════════

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_oneModule_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_oneModule_overpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, oneModule);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_twoModules_underpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_revert_incorrectFeeAmount_twoModules_overpaid() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE + 1;
        expectIncorrectFeeAmount(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryWithVerification(incorrectFee, writerF, dataHash, twoModules);
    }

    function test_writeEntryWithVerification_revert_noModulesSpecified() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
        expectNoModulesSpecified();
        writeEntryWithVerification(0, writerF, dataHash, new address[](0));
    }

    function test_writeEntryWithVerification_revert_sameChainId() public {
        bytes32 dataHash = getMockDataHash(writerF, INITIAL_WRITER_F_NONCE);
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
