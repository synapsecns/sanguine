// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, InterchainEntry, IInterchainDB, InterchainDBEvents} from "../contracts/InterchainDB.sol";

import {InterchainEntryLibHarness} from "./harnesses/InterchainEntryLibHarness.sol";
import {VersionedPayloadLibHarness} from "./harnesses/VersionedPayloadLibHarness.sol";
import {InterchainModuleMock} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
/// @notice Unit tests for InterchainDB interactions on the destination chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBDestinationTest is Test, InterchainDBEvents {
    uint64 public constant SRC_CHAIN_ID_0 = 1337;
    uint64 public constant SRC_CHAIN_ID_1 = 1338;
    uint64 public constant DST_CHAIN_ID = 7331;

    uint16 public constant DB_VERSION = 1;

    uint256 public constant ENTRY_CONFLICT = type(uint256).max;

    InterchainEntryLibHarness public entryLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

    address public writerT = makeAddr("Test Writer");

    mapping(address module => mapping(bytes32 entryHash => uint256 timestamp)) public verifiedAt;

    function setUp() public {
        vm.chainId(DST_CHAIN_ID);
        icDB = new InterchainDB();
        moduleA = new InterchainModuleMock();
        moduleB = new InterchainModuleMock();
        entryLibHarness = new InterchainEntryLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        // Format is {chainId: nonce}
        // Verify some entries with module A
        // A: {0: 0}, {0: 10}, {1: 10}
        verifyEntry(moduleA, getVersionedEntry(getMockEntry(SRC_CHAIN_ID_0, 0)));
        verifyEntry(moduleA, getVersionedEntry(getMockEntry(SRC_CHAIN_ID_0, 10)));
        verifyEntry(moduleA, getVersionedEntry(getMockEntry(SRC_CHAIN_ID_1, 10)));
        // Verify some entries with module B
        // B: {1: 0}, {1: 10}
        verifyEntry(moduleB, getVersionedEntry(getMockEntry(SRC_CHAIN_ID_1, 0)));
        verifyEntry(moduleB, getVersionedEntry(getMockEntry(SRC_CHAIN_ID_1, 10)));
    }

    function getVersionedEntry(InterchainEntry memory entry) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(DB_VERSION, entryLibHarness.encodeEntry(entry));
    }

    function verifyEntry(InterchainModuleMock module, bytes memory versionedEntry) internal {
        skip(1 minutes);
        verifiedAt[address(module)][keccak256(abi.encode(versionedEntry))] = block.timestamp;
        module.mockVerifyRemoteEntry(address(icDB), versionedEntry);
    }

    function introduceConflicts() public {
        // Have module A verify a different entry {1:0} (already verified by module B)
        verifyEntry(moduleA, getVersionedEntry(getFakeEntry(SRC_CHAIN_ID_1, 0)));
        // Have module B verify a different entry {0:10} (already verified by module A)
        verifyEntry(moduleB, getVersionedEntry(getFakeEntry(SRC_CHAIN_ID_0, 10)));
    }

    function introduceEmptyEntries() public {
        // Have module A verify an empty entry for entries that module B has not verified
        // {0: 20}
        verifyEntry(moduleA, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_0, 20)));
        // Have module A verify an empty entry for entries that module B has verified
        // {1: 0}
        verifyEntry(moduleA, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_1, 0)));
        // Have module B verify an empty entry for entries that module A has not verified
        // {0: 30}
        verifyEntry(moduleB, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_0, 30)));
        // Have module B verify an empty entry for entries that module A has verified
        // {0: 10}
        verifyEntry(moduleB, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_0, 10)));
    }

    function introduceEqualEmptyEntries() public {
        // {0: 20}
        verifyEntry(moduleA, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_0, 20)));
        verifyEntry(moduleB, getVersionedEntry(getEmptyEntry(SRC_CHAIN_ID_0, 20)));
    }

    // ══════════════════════════════════════════════ DATA GENERATION ══════════════════════════════════════════════════

    function getMockEntryValue(uint64 nonce) internal view returns (bytes32) {
        return entryLibHarness.getEntryValue(writerT, getMockDigest(nonce));
    }

    function getMockEntry(uint64 srcChainId, uint64 dbNonce) internal view returns (InterchainEntry memory entry) {
        return InterchainEntry(srcChainId, dbNonce, getMockEntryValue(dbNonce));
    }

    function getMockDigest(uint64 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce));
    }

    function getFakeEntryValue(uint64 nonce) internal view returns (bytes32) {
        return entryLibHarness.getEntryValue(writerT, getFakeDigest(nonce));
    }

    function getFakeEntry(uint64 srcChainId, uint64 dbNonce) internal view returns (InterchainEntry memory entry) {
        return InterchainEntry(srcChainId, dbNonce, getFakeEntryValue(dbNonce));
    }

    function getFakeDigest(uint64 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce, "Fake data"));
    }

    function getEmptyEntry(uint64 srcChainId, uint64 dbNonce) internal pure returns (InterchainEntry memory entry) {
        return InterchainEntry({srcChainId: srcChainId, dbNonce: dbNonce, entryValue: 0});
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectInitialVerificationTime(
        InterchainModuleMock module,
        InterchainEntry memory entry
    )
        internal
        view
    {
        bytes memory versionedEntry = getVersionedEntry(entry);
        uint256 savedVerificationTime = verifiedAt[address(module)][keccak256(abi.encode(versionedEntry))];
        // We never save 0 as a verification time during initial setup
        assertGt(savedVerificationTime, 0);
        checkVerification(module, entry, savedVerificationTime);
    }

    function checkVerification(
        InterchainModuleMock module,
        InterchainEntry memory entry,
        uint256 expectedVerificationTime
    )
        internal
        view
    {
        uint256 timestamp = icDB.checkEntryVerification(address(module), entry);
        assertEq(timestamp, expectedVerificationTime);
    }

    function expectEventEntryVerified(InterchainModuleMock module, InterchainEntry memory entry) internal {
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(address(module), entry.srcChainId, entry.dbNonce, entry.entryValue);
    }

    function expectRevertEntryConflict(InterchainModuleMock module, InterchainEntry memory newEntry) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__EntryConflict.selector, address(module), newEntry)
        );
    }

    function expectRevertEntryVersionMismatch(uint16 version, uint16 required) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__EntryVersionMismatch.selector, version, required)
        );
    }

    function expectChainIdNotRemote(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__ChainIdNotRemote.selector, chainId));
    }

    // ═════════════════════════════════════════ TESTS: VERIFYING ENTRIES ══════════════════════════════════════════════

    function test_verifyEntry_newKey_nonEmptyRoot() public {
        skip(1 days);
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20);
        bytes memory versionedEntry = getVersionedEntry(entry);
        expectEventEntryVerified(moduleA, entry);
        verifyEntry(moduleA, versionedEntry);
        checkVerification(moduleA, entry, block.timestamp);
    }

    function test_verifyEntry_newKey_emptyRoot() public {
        skip(1 days);
        InterchainEntry memory entry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        bytes memory versionedEntry = getVersionedEntry(entry);
        expectEventEntryVerified(moduleA, entry);
        verifyEntry(moduleA, versionedEntry);
        checkVerification(moduleA, entry, block.timestamp);
    }

    function test_verifyEntry_sameKey_diffModule_prevEmptyRoot_emptyRoot() public {
        introduceEmptyEntries();
        skip(1 days);
        // {0: 30} was verified as "empty" by module B, but not by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 30);
        bytes memory versionedEmptyEntry = getVersionedEntry(emptyEntry);
        expectEventEntryVerified(moduleA, emptyEntry);
        verifyEntry(moduleA, versionedEmptyEntry);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, emptyEntry, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_verifyEntry_sameKey_diffModule_prevEmptyRoot_nonEmptyRoot() public {
        introduceEmptyEntries();
        skip(1 days);
        // {0: 30} was verified as "empty" by module B, but not by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 30);
        InterchainEntry memory mockEntry = getMockEntry(SRC_CHAIN_ID_0, 30);
        bytes memory versionedEntry = getVersionedEntry(mockEntry);
        expectEventEntryVerified(moduleA, mockEntry);
        verifyEntry(moduleA, versionedEntry);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, mockEntry, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_verifyEntry_sameKey_diffModule_prevNonEmptyRoot_emptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_1, 0);
        bytes memory versionedEmptyEntry = getVersionedEntry(emptyEntry);
        expectEventEntryVerified(moduleA, emptyEntry);
        verifyEntry(moduleA, versionedEmptyEntry);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, emptyEntry, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_verifyEntry_sameKey_diffModule_prevNonEmptyRoot_diffNonEmptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        InterchainEntry memory conflictingEntry = getFakeEntry(SRC_CHAIN_ID_1, 0);
        bytes memory versionedConflictingEntry = getVersionedEntry(conflictingEntry);
        expectEventEntryVerified(moduleA, conflictingEntry);
        verifyEntry(moduleA, versionedConflictingEntry);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, conflictingEntry, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_verifyEntry_sameKey_diffModule_prevNonEmptyRoot_sameNonEmptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        bytes memory versionedEntry = getVersionedEntry(entry);
        expectEventEntryVerified(moduleA, entry);
        verifyEntry(moduleA, versionedEntry);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, entry, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_verifyEntry_sameKey_sameModule_prevEmptyRoot_emptyRoot() public {
        introduceEmptyEntries();
        skip(1 days);
        // {0: 20} was verified as "empty" by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        bytes memory versionedEmptyEntry = getVersionedEntry(emptyEntry);
        // Should emit no event and not update the verification time
        uint256 moduleAVerificationTime = verifiedAt[address(moduleA)][keccak256(abi.encode(versionedEmptyEntry))];
        vm.recordLogs();
        verifyEntry(moduleA, versionedEmptyEntry);
        assertEq(vm.getRecordedLogs().length, 0);
        checkVerification(moduleA, emptyEntry, moduleAVerificationTime);
    }

    function test_verifyEntry_sameKey_sameModule_prevEmptyRoot_nonEmptyRoot() public {
        introduceEmptyEntries();
        skip(1 days);
        // {0: 20} was verified as "empty" by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20);
        bytes memory versionedEntry = getVersionedEntry(entry);
        // Overwriting an empty entry with a non-empty one is allowed
        expectEventEntryVerified(moduleA, entry);
        verifyEntry(moduleA, versionedEntry);
        checkVerification(moduleA, entry, block.timestamp);
    }

    function test_verifyEntry_sameKey_sameModule_prevNonEmptyRoot_sameNonEmptyRoot() public {
        skip(1 days);
        // {0: 10} was verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 10);
        bytes memory versionedEntry = getVersionedEntry(entry);
        // Should emit no event and not update the verification time
        uint256 moduleAVerificationTime = verifiedAt[address(moduleA)][keccak256(abi.encode(versionedEntry))];
        vm.recordLogs();
        verifyEntry(moduleA, versionedEntry);
        assertEq(vm.getRecordedLogs().length, 0);
        checkVerification(moduleA, entry, moduleAVerificationTime);
    }

    // ════════════════════════════════════ TESTS: VERIFYING ENTRIES (REVERTS) ═════════════════════════════════════════

    function test_verifyEntry_sameKey_sameModule_prevNonEmptyRoot_emptyRoot_revert() public {
        // {0: 0} was verified by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 0);
        bytes memory versionedEmptyEntry = getVersionedEntry(emptyEntry);
        expectRevertEntryConflict(moduleA, emptyEntry);
        verifyEntry(moduleA, versionedEmptyEntry);
    }

    function test_verifyEntry_sameKey_sameModule_prevNonEmptyRoot_diffNonEmptyRoot_revert() public {
        // {0: 0} was verified by module A
        InterchainEntry memory conflictingEntry = getFakeEntry(SRC_CHAIN_ID_0, 0);
        bytes memory versionedConflictingEntry = getVersionedEntry(conflictingEntry);
        expectRevertEntryConflict(moduleA, conflictingEntry);
        verifyEntry(moduleA, versionedConflictingEntry);
    }

    function test_verifyEntry_revert_ChainIdNotRemote() public {
        // Try to verify entry coming from the same chain
        InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, 0);
        bytes memory versionedEntry = getVersionedEntry(entry);
        expectChainIdNotRemote(DST_CHAIN_ID);
        verifyEntry(moduleA, versionedEntry);
    }

    function test_verifyEntry_revert_wrongVersion(uint16 version) public {
        vm.assume(version != DB_VERSION);
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        bytes memory versionedEntry =
            payloadLibHarness.encodeVersionedPayload(version, entryLibHarness.encodeEntry(entry));
        expectRevertEntryVersionMismatch(version, DB_VERSION);
        moduleA.mockVerifyRemoteEntry(address(icDB), versionedEntry);
    }

    // ══════════════════════════════════════════ TESTS: READING ENTRIES ═══════════════════════════════════════════════

    function test_checkVerification_existingA_existingB() public view {
        // {1: 10} was verified by module A and module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_existingA_unknownB() public view {
        // {0: 0} was verified by module A, but not by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_existingA_differentB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" entry was verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, ENTRY_CONFLICT);
    }

    function test_checkVerification_existingA_emptyB() public {
        introduceEmptyEntries();
        // {0: 10} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, ENTRY_CONFLICT);
    }

    function test_checkVerification_unknownA_existingB() public view {
        // {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, entry, 0);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_unknownA_unknownB() public view {
        // {0: 20} was not verified by any module
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20);
        checkVerification(moduleA, entry, 0);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_unknownA_differentB() public view {
        // {1: 0} was verified by module B, but not by module A
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, fakeEntry, 0);
        checkVerification(moduleB, fakeEntry, ENTRY_CONFLICT);
    }

    function test_checkVerification_unknownA_emptyB() public {
        introduceEmptyEntries();
        // {0: 30} was not verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 30);
        checkVerification(moduleA, emptyEntry, 0);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_differentA_existingB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" entry was verified by module B
        // Check the fake entry that A never verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeEntry, ENTRY_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, fakeEntry);
    }

    function test_checkVerification_differentA_unknownB() public view {
        // {0: 10} was verified by module A, but not by module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeEntry, ENTRY_CONFLICT);
        checkVerification(moduleB, fakeEntry, 0);
    }

    function test_checkVerification_differentA_differentB() public view {
        // {1: 10} was verified by module A and module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_1, 10);
        checkVerification(moduleA, fakeEntry, ENTRY_CONFLICT);
        checkVerification(moduleB, fakeEntry, ENTRY_CONFLICT);
    }

    function test_checkVerification_differentA_emptyB() public {
        introduceEmptyEntries();
        // {0: 10} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, emptyEntry, ENTRY_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_emptyA_existingB() public {
        introduceEmptyEntries();
        // {1: 0} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, entry, ENTRY_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_emptyA_unknownB() public {
        introduceEmptyEntries();
        // {0: 20} was verified as "empty" by module A, but not by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        checkVerification(moduleB, emptyEntry, 0);
    }

    function test_checkVerification_emptyA_differentB() public {
        introduceEmptyEntries();
        // {1: 0} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_1, 0);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        checkVerification(moduleB, emptyEntry, ENTRY_CONFLICT);
    }

    function test_checkVerification_emptyA_emptyB() public {
        introduceEqualEmptyEntries();
        // {0: 20} was verified as "empty" by module A and module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_modifySrcChainId() public view {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.srcChainId ^= 1;
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifyDbNonce() public view {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.dbNonce ^= 1;
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifyEntryValue() public view {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.entryValue ^= bytes32(uint256(1));
        checkVerification(moduleA, entry, ENTRY_CONFLICT);
    }

    // ═════════════════════════════════════ TESTS: READING ENTRIES (REVERTS) ══════════════════════════════════════════

    function test_checkVerification_revert_ChainIdNotRemote() public {
        InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, 0);
        expectChainIdNotRemote(DST_CHAIN_ID);
        icDB.checkEntryVerification(address(moduleA), entry);
    }
}
