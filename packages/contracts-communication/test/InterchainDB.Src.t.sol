// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainDB,
    InterchainEntry,
    InterchainEntryLib,
    IInterchainDB,
    InterchainDBEvents
} from "../contracts/InterchainDB.sol";

import {InterchainEntryLibHarness} from "./harnesses/InterchainEntryLibHarness.sol";
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

    InterchainEntryLibHarness public entryLibHarness;
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
        entryLibHarness = new InterchainEntryLibHarness();
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
            writeEntry(writerF, getMockDigest(i, writerF));
            initialEntries.push(getMockEntry(i, writerF));
        }
        for (uint64 i = 0; i < INITIAL_WRITER_S; ++i) {
            writeEntry(writerS, getMockDigest(INITIAL_WRITER_F + i, writerS));
            initialEntries.push(getMockEntry(INITIAL_WRITER_F + i, writerS));
        }
    }

    function getInitialEntry(uint64 dbNonce) internal view returns (InterchainEntry memory) {
        require(dbNonce < initialEntries.length, "dbNonce out of range");
        return initialEntries[dbNonce];
    }

    function getEmptyEntry(uint64 dbNonce) internal pure returns (InterchainEntry memory) {
        return InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: dbNonce, entryValue: 0});
    }

    function getMockDigest(uint64 nonce, address writer) internal pure returns (bytes32) {
        return keccak256(abi.encode("Mock data", nonce, writer));
    }

    function getMockEntry(uint64 dbNonce, address writer) internal view returns (InterchainEntry memory entry) {
        bytes32 digest = getMockDigest(dbNonce, writer);
        return InterchainEntry({
            srcChainId: SRC_CHAIN_ID,
            dbNonce: dbNonce,
            entryValue: entryLibHarness.getEntryValue(writer, digest)
        });
    }

    function getVersionedEntry(InterchainEntry memory entry) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(DB_VERSION, entryLibHarness.encodeEntry(entry));
    }

    function getModuleCalldata(InterchainEntry memory entry) internal view returns (bytes memory) {
        bytes memory versionedEntry = getVersionedEntry(entry);
        return abi.encodeCall(IInterchainModule.requestEntryVerification, (DST_CHAIN_ID, versionedEntry));
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

    function writeEntry(address writer, bytes32 digest) internal returns (uint64 dbNonce) {
        vm.prank(writer);
        return icDB.writeEntry(digest);
    }

    function requestVerification(address caller, uint256 msgValue, uint64 dbNonce, address[] memory modules) internal {
        deal(caller, msgValue);
        vm.prank(caller);
        icDB.requestEntryVerification{value: msgValue}(DST_CHAIN_ID, dbNonce, modules);
    }

    function writeEntryRequestVerification(
        uint256 msgValue,
        address writer,
        bytes32 digest,
        address[] memory modules
    )
        internal
        returns (uint64 dbNonce)
    {
        deal(writer, msgValue);
        vm.prank(writer);
        return icDB.writeEntryRequestVerification{value: msgValue}(DST_CHAIN_ID, digest, modules);
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function checkSavedEntry(uint64 dbNonce, InterchainEntry memory expected) internal view {
        assertEq(icDB.getEntry(dbNonce), expected);
        assertEq(icDB.getEntryValue(dbNonce), expected.entryValue);
    }

    function assertEq(InterchainEntry memory actual, InterchainEntry memory expected) internal pure {
        assertEq(actual.srcChainId, expected.srcChainId, "!srcChainId");
        assertEq(actual.dbNonce, expected.dbNonce, "!dbNonce");
        assertEq(actual.entryValue, expected.entryValue, "!entryValue");
    }

    function expectEventInterchainEntryWritten(
        uint64 dbNonce,
        address srcWriter,
        bytes32 digest,
        bytes32 entryValue
    )
        internal
    {
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten({
            dbNonce: dbNonce,
            srcWriter: addressToBytes32(srcWriter),
            digest: digest,
            entryValue: entryValue
        });
    }

    function expectEventEntryVerificationRequested(uint64 dbNonce, address[] memory srcModules) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerificationRequested(DST_CHAIN_ID, dbNonce, srcModules);
    }

    function expectRevertFeeAmountBelowMin(uint256 feeAmount, uint256 minRequired) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__FeeAmountBelowMin.selector, feeAmount, minRequired)
        );
    }

    function expectRevertModulesNotProvided() internal {
        vm.expectRevert(IInterchainDB.InterchainDB__ModulesNotProvided.selector);
    }

    function expectRevertChainIdNotRemote(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__ChainIdNotRemote.selector, chainId));
    }

    // ═══════════════════════════════════════════════ TESTS: SET UP ═══════════════════════════════════════════════════

    function test_setup_getters() public view {
        for (uint64 i = 0; i < INITIAL_DB_NONCE; ++i) {
            checkSavedEntry(i, getInitialEntry(i));
        }
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE);
    }

    // ══════════════════════════════════════════ TESTS: WRITING AN ENTRY ══════════════════════════════════════════════

    function test_writeEntry_writerF() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        bytes32 entryValue = getMockEntry(INITIAL_DB_NONCE, writerF).entryValue;
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerF, digest, entryValue);
        uint64 nonce = writeEntry(writerF, digest);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, getMockEntry(INITIAL_DB_NONCE, writerF));
    }

    function test_writeEntry_writerS() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerS);
        bytes32 entryValue = getMockEntry(INITIAL_DB_NONCE, writerS).entryValue;
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerS, digest, entryValue);
        uint64 nonce = writeEntry(writerS, digest);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, getMockEntry(INITIAL_DB_NONCE, writerS));
    }

    // ═══════════════════════════════════════ TESTS: REQUESTING VALIDATION ════════════════════════════════════════════

    function test_requestVerification_writerF_oneModule() public {
        uint64 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_oneModule_higherFee() public {
        uint64 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_writerF_twoModules() public {
        uint64 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerF_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 0;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_oneModule() public {
        uint64 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_oneModule_higherFee() public {
        uint64 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_writerS_twoModules() public {
        uint64 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_writerS_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 2;
        InterchainEntry memory entry = getInitialEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_nextNonce_oneModule() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_nextNonce_oneModule_higherFee() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_nextNonce_twoModules() public {
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_nextNonce_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = INITIAL_DB_NONCE;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE * 2 + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_hugeNonce_oneModule() public {
        uint64 dbNonce = 2 ** 32;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE, dbNonce, oneModule);
    }

    function test_requestVerification_hugeNonce_oneModule_higherFee() public {
        uint64 dbNonce = 2 ** 32;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, oneModule);
        requestVerification(requestCaller, MODULE_A_FEE * 2, dbNonce, oneModule);
    }

    function test_requestVerification_hugeNonce_twoModules() public {
        uint64 dbNonce = 2 ** 32;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
        requestVerification(requestCaller, MODULE_A_FEE + MODULE_B_FEE, dbNonce, twoModules);
    }

    function test_requestVerification_hugeNonce_twoModules_higherFee() public {
        // Overpaid fees should be directed to the first module
        uint64 dbNonce = 2 ** 32;
        InterchainEntry memory entry = getEmptyEntry(dbNonce);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventEntryVerificationRequested(dbNonce, twoModules);
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
        icDB.requestEntryVerification(SRC_CHAIN_ID, 0, oneModule);
    }

    // ═════════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION ═════════════════════════════════════════

    function test_writeEntryRequestVerification_writerF_oneModule() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerF, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, oneModule);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE, writerF, digest, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerF_oneModule_higherFee() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerF, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, oneModule);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE * 2, writerF, digest, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerF_twoModules() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerF, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, twoModules);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE + MODULE_B_FEE, writerF, digest, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerF_twoModules_higherFee() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerF);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerF, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, twoModules);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE * 2 + MODULE_B_FEE, writerF, digest, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerS_oneModule() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerS);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerS, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, oneModule);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE, writerS, digest, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerS_oneModule_higherFee() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerS);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerS, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, oneModule);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE * 2, writerS, digest, oneModule);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerS_twoModules() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerS);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerS, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, twoModules);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE + MODULE_B_FEE, writerS, digest, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    function test_writeEntryRequestVerification_writerS_twoModules_higherFee() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerS);
        InterchainEntry memory entry = getMockEntry(INITIAL_DB_NONCE, writerS);
        vm.expectCall({callee: address(moduleA), msgValue: MODULE_A_FEE * 2, data: getModuleCalldata(entry)});
        vm.expectCall({callee: address(moduleB), msgValue: MODULE_B_FEE, data: getModuleCalldata(entry)});
        expectEventInterchainEntryWritten(INITIAL_DB_NONCE, writerS, digest, entry.entryValue);
        expectEventEntryVerificationRequested(INITIAL_DB_NONCE, twoModules);
        uint64 nonce = writeEntryRequestVerification(MODULE_A_FEE * 2 + MODULE_B_FEE, writerS, digest, twoModules);
        assertEq(nonce, INITIAL_DB_NONCE);
        assertEq(icDB.getDBNonce(), INITIAL_DB_NONCE + 1);
        checkSavedEntry(INITIAL_DB_NONCE, entry);
    }

    // ════════════════════════════════ TESTS: WRITE + REQUEST VALIDATION (REVERTS) ════════════════════════════════════

    function test_writeEntryRequestVerification_revert_FeeAmountBelowMin_oneModule_underpaid() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        uint256 incorrectFee = MODULE_A_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE);
        writeEntryRequestVerification(incorrectFee, writerF, digest, oneModule);
    }

    function test_writeEntryRequestVerification_revert_FeeAmountBelowMin_twoModules_underpaid() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        uint256 incorrectFee = MODULE_A_FEE + MODULE_B_FEE - 1;
        expectRevertFeeAmountBelowMin(incorrectFee, MODULE_A_FEE + MODULE_B_FEE);
        writeEntryRequestVerification(incorrectFee, writerF, digest, twoModules);
    }

    function test_writeEntryRequestVerification_revert_ModulesNotProvided() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        expectRevertModulesNotProvided();
        writeEntryRequestVerification(0, writerF, digest, new address[](0));
    }

    function test_writeEntryRequestVerification_revert_ChainIdNotRemote() public {
        bytes32 digest = getMockDigest(INITIAL_DB_NONCE, writerF);
        expectRevertChainIdNotRemote(SRC_CHAIN_ID);
        vm.prank(writerF);
        icDB.writeEntryRequestVerification(SRC_CHAIN_ID, digest, oneModule);
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

    function test_getEntry_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            InterchainEntry memory entry = icDB.getEntry(nonce);
            assertEq(entry, getInitialEntry(nonce));
        }
    }

    function test_getEntry_nextNonce() public view {
        InterchainEntry memory entry = icDB.getEntry(INITIAL_DB_NONCE);
        InterchainEntry memory expected =
            InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: INITIAL_DB_NONCE, entryValue: 0});
        assertEq(entry, expected);
    }

    function test_getEntry_hugeNonce() public view {
        InterchainEntry memory entry = icDB.getEntry(2 ** 32);
        InterchainEntry memory expected = InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: 2 ** 32, entryValue: 0});
        assertEq(entry, expected);
    }

    function test_getEntryValue_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            InterchainEntry memory expectedEntry = getInitialEntry(nonce);
            assertEq(icDB.getEntryValue(nonce), expectedEntry.entryValue);
        }
    }

    function test_getEntryValue_nextNonce() public view {
        assertEq(icDB.getEntryValue(INITIAL_DB_NONCE), 0);
    }

    function test_getEntryValue_hugeNonce() public view {
        assertEq(icDB.getEntryValue(2 ** 32), 0);
    }

    function test_getEncodedEntry_finalized() public view {
        for (uint64 nonce = 0; nonce < INITIAL_DB_NONCE; ++nonce) {
            InterchainEntry memory entry = getInitialEntry(nonce);
            bytes memory encodedEntry = icDB.getEncodedEntry(nonce);
            assertEq(encodedEntry, getVersionedEntry(entry));
        }
    }

    function test_getEncodedEntry_nextNonce() public view {
        InterchainEntry memory entry =
            InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: INITIAL_DB_NONCE, entryValue: 0});
        bytes memory encodedEntry = icDB.getEncodedEntry(INITIAL_DB_NONCE);
        assertEq(encodedEntry, getVersionedEntry(entry));
    }

    function test_getEncodedEntry_hugeNonce() public view {
        InterchainEntry memory entry = InterchainEntry({srcChainId: SRC_CHAIN_ID, dbNonce: 2 ** 32, entryValue: 0});
        bytes memory encodedEntry = icDB.getEncodedEntry(2 ** 32);
        assertEq(encodedEntry, getVersionedEntry(entry));
    }
}
