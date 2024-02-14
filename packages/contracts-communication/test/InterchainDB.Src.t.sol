// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, IInterchainDB, IInterchainDBEvents} from "../contracts/InterchainDB.sol";

import {InterchainModuleMock, IInterchainModule} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Unit tests for InterchainDB interactions on the source chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBSourceTest is Test, IInterchainDBEvents {
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

    function getMockEntry(
        address writer,
        uint256 nonce
    )
        internal
        pure
        returns (IInterchainDB.InterchainEntry memory entry)
    {
        return IInterchainDB.InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            srcWriter: addressToBytes32(writer),
            writerNonce: nonce,
            dataHash: getMockDataHash(writer, nonce)
        });
    }

    function getModuleCalldata(IInterchainDB.InterchainEntry memory entry) internal pure returns (bytes memory) {
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

    function assertEq(
        IInterchainDB.InterchainEntry memory entry,
        IInterchainDB.InterchainEntry memory expected
    )
        internal
    {
        assertEq(entry.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(entry.srcWriter, expected.srcWriter, "!srcWriter");
        assertEq(entry.writerNonce, expected.writerNonce, "!writerNonce");
        assertEq(entry.dataHash, expected.dataHash, "!dataHash");
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
        IInterchainDB.InterchainEntry memory entry = getMockEntry(writerF, writerNonce);
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
        IInterchainDB.InterchainEntry memory entry = getMockEntry(writerF, writerNonce);
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
        IInterchainDB.InterchainEntry memory entry = getMockEntry(writerS, writerNonce);
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
        IInterchainDB.InterchainEntry memory entry = getMockEntry(writerS, writerNonce);
        // expectCall(address callee, uint256 msgValue, bytes calldata data)
        vm.expectCall(address(moduleA), MODULE_A_FEE, getModuleCalldata(entry));
        vm.expectCall(address(moduleB), MODULE_B_FEE, getModuleCalldata(entry));
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, writerS, writerNonce, twoModules);
    }

    // ═════════════════════════════════════════ TESTS: GET INTERCHAIN FEE ═════════════════════════════════════════════

    function test_getInterchainFee_oneModule() public {
        // [moduleA]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, oneModule), MODULE_A_FEE);
    }

    function test_getInterchainFee_twoModules() public {
        // [moduleA, moduleB]
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, twoModules), MODULE_A_FEE + MODULE_B_FEE);
    }
}
