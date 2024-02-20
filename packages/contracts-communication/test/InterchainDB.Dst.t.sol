// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainDB,
    InterchainEntry,
    InterchainEntryLib,
    IInterchainDB,
    InterchainDBEvents
} from "../contracts/InterchainDB.sol";

import {InterchainModuleMock, IInterchainModule} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Unit tests for InterchainDB interactions on the destination chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBDestinationTest is Test, InterchainDBEvents {
    uint256 public constant SRC_CHAIN_ID_0 = 1337;
    uint256 public constant SRC_CHAIN_ID_1 = 1338;
    uint256 public constant DST_CHAIN_ID = 7331;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

    address public writerF = makeAddr("First Writer");
    address public writerS = makeAddr("Second Writer");
    address public outsideAddress = makeAddr("Outside Address");

    mapping(address module => mapping(bytes32 entryHash => uint256 timestamp)) public verifiedAt;

    function setUp() public {
        vm.chainId(DST_CHAIN_ID);
        icDB = new InterchainDB();
        moduleA = new InterchainModuleMock();
        moduleB = new InterchainModuleMock();
        // Format is writer: {chainId: nonce}
        // Verify some entries with module A
        // writerF: {0: 0}, {0: 10}, {1: 10}
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_0, 0, writerF));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_0, 10, writerF));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, 10, writerF));
        // writerS: {1: 1}, {1: 11}
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, 1, writerS));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, 11, writerS));
        // Verify some entries with module B
        // writerF: {1: 0}, {1: 10}
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, 0, writerF));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, 10, writerF));
        // writerS: {0: 1}, {0: 11}, {1: 1}
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_0, 1, writerS));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_0, 11, writerS));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, 1, writerS));
    }

    function verifyEntry(InterchainModuleMock module, InterchainEntry memory entry) internal {
        skip(1 minutes);
        verifiedAt[address(module)][keccak256(abi.encode(entry))] = block.timestamp;
        module.mockVerifyEntry(address(icDB), entry);
    }

    function introduceConflictsSameWriter() public {
        // Have module A verify a different entry for writerS {0:11} (already verified by module B)
        verifyEntry(moduleA, getFakeEntry(SRC_CHAIN_ID_0, 11, writerS));
        // Have module B verify a different entry for writerF {0:10} (already verified by module A)
        verifyEntry(moduleB, getFakeEntry(SRC_CHAIN_ID_0, 10, writerF));
    }

    function introduceConflictsDiffWriter() public {
        // Have module A verify an entry for writerF {0: 11} (already verified by module B with writerS)
        verifyEntry(moduleA, getFakeEntry(SRC_CHAIN_ID_0, 11, writerF));
        // Have module B verify an entry for writerS {0: 10} (already verified by module A with writerF)
        verifyEntry(moduleB, getFakeEntry(SRC_CHAIN_ID_0, 10, writerS));
    }

    function introduceEmptyEntries() public {
        // Have module A verify an empty entry for entries that module B has not verified
        // writerF {0:20}
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_0, 20, writerF));
        // writerS {1:5}
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_1, 5, writerS));
        // Have module A verify an empty entry for entries that module B has verified
        // writerF {1:0}
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_1, 0, writerF));
        // writerS {0:1}
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_0, 1, writerS));
        // Have module B verify an empty entry for entries that module A has not verified
        // writerF {0:30}
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_0, 30, writerF));
        // writerS {1:21}
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_1, 21, writerS));
        // Have module B verify an empty entry for entries that module A has verified
        // writerF {0: 10}
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_0, 10, writerF));
        // writerS {1: 11}
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_1, 11, writerS));
    }

    function introduceEqualEmptyEntries() public {
        // writerF {0: 20}
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_0, 20, writerF));
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_0, 20, writerF));
        // writerS {1: 5}
        verifyEntry(moduleB, getEmptyEntry(SRC_CHAIN_ID_1, 5, writerS));
        verifyEntry(moduleA, getEmptyEntry(SRC_CHAIN_ID_1, 5, writerS));
    }

    // ══════════════════════════════════════════════ DATA GENERATION ══════════════════════════════════════════════════

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(
        uint256 srcChainId,
        uint256 dbNonce,
        address writer
    )
        internal
        pure
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            srcWriter: addressToBytes32(writer),
            dataHash: getMockDataHash(writer, dbNonce)
        });
    }

    function getFakeDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce, "Fake data"));
    }

    function getFakeEntry(
        uint256 srcChainId,
        uint256 dbNonce,
        address writer
    )
        internal
        pure
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            srcWriter: addressToBytes32(writer),
            dataHash: getFakeDataHash(writer, dbNonce)
        });
    }

    function getEmptyEntry(
        uint256 srcChainId,
        uint256 dbNonce,
        address writer
    )
        internal
        pure
        returns (InterchainEntry memory entry)
    {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            srcWriter: addressToBytes32(writer),
            dataHash: 0
        });
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectVerificationTime(
        InterchainEntry memory entry,
        address module,
        uint256 timestampToCheck
    )
        internal
    {
        assertGt(timestampToCheck, 0);
        assertEq(timestampToCheck, verifiedAt[module][keccak256(abi.encode(entry))]);
    }

    function expectConflictingEntries(InterchainEntry memory existingEntry, InterchainEntry memory newEntry) internal {
        bytes32 entryValue = InterchainEntryLib.entryValue(existingEntry);
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__ConflictingEntries.selector, entryValue, newEntry)
        );
    }

    function expectSameChainId() internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__SameChainId.selector));
    }

    // ═════════════════════════════════════════ TESTS: VERIFYING ENTRIES ══════════════════════════════════════════════

    function test_verifyEntry_new_emitsEvent() public {
        skip(1 days);
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20, writerF);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(address(moduleA), entry.srcChainId, entry.dbNonce, entry.srcWriter, entry.dataHash);
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_new_savesVerificationTime() public {
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20, writerF);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    function test_verifyEntry_existing_diffModule_emitsEvent() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(address(moduleA), entry.srcChainId, entry.dbNonce, entry.srcWriter, entry.dataHash);
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_existing_diffModule_doesNotUpdateExistingVerificationTime() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        uint256 moduleBVerifiedAt = verifiedAt[address(moduleB)][keccak256(abi.encode(entry))];
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleB), entry), moduleBVerifiedAt);
    }

    function test_verifyEntry_existing_diffModule_savesVerificationTime() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    function test_verifyEntry_existing_sameModule_doesNotEmitEvent() public {
        // writerF {0:0} was already verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0, writerF);
        vm.recordLogs();
        verifyEntry(moduleA, entry);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_verifyEntry_existing_sameModule_doesNotUpdateExistingVerificationTime() public {
        // writerF {0:0} was already verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0, writerF);
        uint256 moduleAVerifiedAt = verifiedAt[address(moduleA)][keccak256(abi.encode(entry))];
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), moduleAVerifiedAt);
    }

    function test_verifyEntry_conflict_diffModule_emitsEvent() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, 1, writerS);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(address(moduleA), entry.srcChainId, entry.dbNonce, entry.srcWriter, entry.dataHash);
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_conflict_diffModule_doesNotUpdateExistingVerificationTime() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, 1, writerS);
        InterchainEntry memory realEntry = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        verifyEntry(moduleA, entry);
        assertCorrectVerificationTime(realEntry, address(moduleB), icDB.readEntry(address(moduleB), realEntry));
    }

    function test_verifyEntry_conflict_diffModule_savesVerificationTime() public {
        // writerS {0:1} was already verified by module B
        InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, 1, writerS);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    // ════════════════════════════════════ TESTS: VERIFYING ENTRIES (REVERTS) ═════════════════════════════════════════

    function test_verifyEntry_revert_conflict_sameModule_sameWriter() public {
        // writerF {0:0} was already verified by module A
        InterchainEntry memory existingEntry = getMockEntry(SRC_CHAIN_ID_0, 0, writerF);
        InterchainEntry memory conflictingEntry = getFakeEntry(SRC_CHAIN_ID_0, 0, writerF);
        expectConflictingEntries(existingEntry, conflictingEntry);
        verifyEntry(moduleA, conflictingEntry);
    }

    function test_verifyEntry_revert_conflict_sameModule_diffWriter() public {
        // writerF {0:0} was already verified by module A
        InterchainEntry memory existingEntry = getMockEntry(SRC_CHAIN_ID_0, 0, writerF);
        InterchainEntry memory conflictingEntry = getFakeEntry(SRC_CHAIN_ID_0, 0, writerS);
        expectConflictingEntries(existingEntry, conflictingEntry);
        verifyEntry(moduleA, conflictingEntry);
    }

    function test_verifyEntry_revert_sameChainId() public {
        // Try to verify entry coming from the same chain
        InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, 0, writerF);
        expectSameChainId();
        verifyEntry(moduleA, entry);
    }

    // ══════════════════════════════════════════ TESTS: READING ENTRIES ═══════════════════════════════════════════════

    function test_readEntry_existingA_existingB() public {
        // writerF {1: 10} was verified by module A and module B
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_1, 10, writerF);
        assertCorrectVerificationTime(entryF, address(moduleA), icDB.readEntry(address(moduleA), entryF));
        assertCorrectVerificationTime(entryF, address(moduleB), icDB.readEntry(address(moduleB), entryF));
        // writerS {1: 1} was verified by module A and module B
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_1, 1, writerS);
        assertCorrectVerificationTime(entryS, address(moduleA), icDB.readEntry(address(moduleA), entryS));
        assertCorrectVerificationTime(entryS, address(moduleB), icDB.readEntry(address(moduleB), entryS));
    }

    function test_readEntry_existingA_unknownB() public {
        // writerF {0: 0} was verified by module A, but not by module B
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_0, 0, writerF);
        assertCorrectVerificationTime(entryF, address(moduleA), icDB.readEntry(address(moduleA), entryF));
        assertEq(icDB.readEntry(address(moduleB), entryF), 0);
        // writerS {1: 11} was verified by module A, but not by module B
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_1, 11, writerS);
        assertCorrectVerificationTime(entryS, address(moduleA), icDB.readEntry(address(moduleA), entryS));
        assertEq(icDB.readEntry(address(moduleB), entryS), 0);
    }

    function test_readEntry_existingA_differentB_sameWriter() public {
        introduceConflictsSameWriter();
        // writerF {0: 10} was verified by module A, but a "fake" entry was verified by module B
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertCorrectVerificationTime(entryF, address(moduleA), icDB.readEntry(address(moduleA), entryF));
        assertEq(icDB.readEntry(address(moduleB), entryF), 0);
        // writerS {0: 11} was verified by module B, but a "fake" entry was verified by module A
        InterchainEntry memory fakeEntryS = getFakeEntry(SRC_CHAIN_ID_0, 11, writerS);
        assertCorrectVerificationTime(fakeEntryS, address(moduleA), icDB.readEntry(address(moduleA), fakeEntryS));
        assertEq(icDB.readEntry(address(moduleB), fakeEntryS), 0);
    }

    function test_readEntry_existingA_differentB_diffWriter() public {
        introduceConflictsDiffWriter();
        // writerF {0: 10} was verified by module A, but a "fake" entry was verified by module B
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertCorrectVerificationTime(entryF, address(moduleA), icDB.readEntry(address(moduleA), entryF));
        assertEq(icDB.readEntry(address(moduleB), entryF), 0);
        // writerS {0: 11} was verified by module B, but a "fake" entry was verified by module A
        InterchainEntry memory fakeEntryS = getFakeEntry(SRC_CHAIN_ID_0, 11, writerF);
        assertCorrectVerificationTime(fakeEntryS, address(moduleA), icDB.readEntry(address(moduleA), fakeEntryS));
        assertEq(icDB.readEntry(address(moduleB), fakeEntryS), 0);
    }

    function test_readEntry_existingA_emptyB() public {
        introduceEmptyEntries();
        // writerF {0: 10} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertCorrectVerificationTime(entryF, address(moduleA), icDB.readEntry(address(moduleA), entryF));
        assertEq(icDB.readEntry(address(moduleB), entryF), 0);
        // writerS {1: 11} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_1, 11, writerS);
        assertCorrectVerificationTime(entryS, address(moduleA), icDB.readEntry(address(moduleA), entryS));
        assertEq(icDB.readEntry(address(moduleB), entryS), 0);
    }

    function test_readEntry_unknownA_existingB() public {
        // writerF {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_1, 0, writerF);
        assertEq(icDB.readEntry(address(moduleA), entryF), 0);
        assertCorrectVerificationTime(entryF, address(moduleB), icDB.readEntry(address(moduleB), entryF));
        // writerS {0: 1} was verified by module B, but not by module A
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        assertEq(icDB.readEntry(address(moduleA), entryS), 0);
        assertCorrectVerificationTime(entryS, address(moduleB), icDB.readEntry(address(moduleB), entryS));
    }

    function test_readEntry_unknownA_unknownB() public {
        // writerF {0: 20} was not verified by any module
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_0, 20, writerF);
        assertEq(icDB.readEntry(address(moduleA), entryF), 0);
        assertEq(icDB.readEntry(address(moduleB), entryF), 0);
        // writerS {1: 5} was not verified by any module
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_1, 5, writerS);
        assertEq(icDB.readEntry(address(moduleA), entryS), 0);
        assertEq(icDB.readEntry(address(moduleB), entryS), 0);
    }

    function test_readEntry_unknownA_differentB() public {
        // writerF {1: 0} was verified by module B, but not by module A
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryF = getFakeEntry(SRC_CHAIN_ID_1, 0, writerF);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryF), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryF), 0);
        // writerS {0: 11} was verified by module B, but not by module A
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryS = getFakeEntry(SRC_CHAIN_ID_0, 11, writerS);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryS), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryS), 0);
    }

    function test_readEntry_unknownA_emptyB() public {
        introduceEmptyEntries();
        // writerF {0: 30} was not verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntryF = getEmptyEntry(SRC_CHAIN_ID_0, 30, writerF);
        assertEq(icDB.readEntry(address(moduleA), emptyEntryF), 0);
        assertCorrectVerificationTime(emptyEntryF, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryF));
        // writerS {1: 21} was not verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntryS = getEmptyEntry(SRC_CHAIN_ID_1, 21, writerS);
        assertEq(icDB.readEntry(address(moduleA), emptyEntryS), 0);
        assertCorrectVerificationTime(emptyEntryS, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryS));
    }

    function test_readEntry_differentA_existingB_sameWriter() public {
        introduceConflictsSameWriter();
        // writerF {0: 10} was verified by module A, but a "fake" entry was verified by module B
        // Check the fake entry that A never verified
        InterchainEntry memory fakeEntryF = getFakeEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryF), 0);
        assertCorrectVerificationTime(fakeEntryF, address(moduleB), icDB.readEntry(address(moduleB), fakeEntryF));
        // writerS {0: 11} was verified by module B, but a "fake" entry was verified by module A
        // Check the real entry that A never verified
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_0, 11, writerS);
        assertEq(icDB.readEntry(address(moduleA), entryS), 0);
        assertCorrectVerificationTime(entryS, address(moduleB), icDB.readEntry(address(moduleB), entryS));
    }

    function test_readEntry_differentA_existingB_diffWriter() public {
        introduceConflictsDiffWriter();
        // writerF {0: 10} was verified by module A, but a "fake" entry was verified by module B (with writerS)
        // Check the fake entry that A never verified
        InterchainEntry memory fakeEntryF = getFakeEntry(SRC_CHAIN_ID_0, 10, writerS);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryF), 0);
        assertCorrectVerificationTime(fakeEntryF, address(moduleB), icDB.readEntry(address(moduleB), fakeEntryF));
        // writerS {0: 11} was verified by module B, but a "fake" entry was verified by module A (with writerF)
        // Check the real entry that A never verified
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_0, 11, writerS);
        assertEq(icDB.readEntry(address(moduleA), entryS), 0);
        assertCorrectVerificationTime(entryS, address(moduleB), icDB.readEntry(address(moduleB), entryS));
    }

    function test_readEntry_differentA_unknownB() public {
        // writerF {0: 10} was verified by module A, but not by module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryF = getFakeEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryF), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryF), 0);
        // writerS {1: 11} was verified by module A, but not by module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryS = getFakeEntry(SRC_CHAIN_ID_1, 11, writerS);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryS), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryS), 0);
    }

    function test_readEntry_differentA_differentB() public {
        // writerF {1: 10} was verified by module A and module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryF = getFakeEntry(SRC_CHAIN_ID_1, 10, writerF);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryF), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryF), 0);
        // writerS {1: 1} was verified by module A and module B
        // Check the fake entry that neither module verified
        InterchainEntry memory fakeEntryS = getFakeEntry(SRC_CHAIN_ID_1, 1, writerS);
        assertEq(icDB.readEntry(address(moduleA), fakeEntryS), 0);
        assertEq(icDB.readEntry(address(moduleB), fakeEntryS), 0);
    }

    function test_readEntry_differentA_emptyB() public {
        introduceEmptyEntries();
        // writerF {0: 10} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntryF = getEmptyEntry(SRC_CHAIN_ID_0, 10, writerF);
        assertEq(icDB.readEntry(address(moduleA), emptyEntryF), 0);
        assertCorrectVerificationTime(emptyEntryF, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryF));
        // writerS {1: 11} was verified by module A, but an "empty" entry was verified by module B
        InterchainEntry memory emptyEntryS = getEmptyEntry(SRC_CHAIN_ID_1, 11, writerS);
        assertEq(icDB.readEntry(address(moduleA), emptyEntryS), 0);
        assertCorrectVerificationTime(emptyEntryS, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryS));
    }

    function test_readEntry_emptyA_existingB() public {
        introduceEmptyEntries();
        // writerF {1: 0} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory entryF = getMockEntry(SRC_CHAIN_ID_1, 0, writerF);
        assertEq(icDB.readEntry(address(moduleA), entryF), 0);
        assertCorrectVerificationTime(entryF, address(moduleB), icDB.readEntry(address(moduleB), entryF));
        // writerS {0: 1} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory entryS = getMockEntry(SRC_CHAIN_ID_0, 1, writerS);
        assertEq(icDB.readEntry(address(moduleA), entryS), 0);
        assertCorrectVerificationTime(entryS, address(moduleB), icDB.readEntry(address(moduleB), entryS));
    }

    function test_readEntry_emptyA_unknownB() public {
        introduceEmptyEntries();
        // writerF {0: 20} was verified as "empty" by module A, but not by module B
        InterchainEntry memory emptyEntryF = getEmptyEntry(SRC_CHAIN_ID_0, 20, writerF);
        assertCorrectVerificationTime(emptyEntryF, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryF));
        assertEq(icDB.readEntry(address(moduleB), emptyEntryF), 0);
        // writerS {1: 5} was verified as "empty" by module A, but not by module B
        InterchainEntry memory emptyEntryS = getEmptyEntry(SRC_CHAIN_ID_1, 5, writerS);
        assertCorrectVerificationTime(emptyEntryS, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryS));
        assertEq(icDB.readEntry(address(moduleB), emptyEntryS), 0);
    }

    function test_readEntry_emptyA_differentB() public {
        introduceEmptyEntries();
        // writerF {1: 0} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory emptyEntryF = getEmptyEntry(SRC_CHAIN_ID_1, 0, writerF);
        assertCorrectVerificationTime(emptyEntryF, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryF));
        assertEq(icDB.readEntry(address(moduleB), emptyEntryF), 0);
        // writerS {0: 1} was verified by module B, but an "empty" entry was verified by module A
        InterchainEntry memory emptyEntryS = getEmptyEntry(SRC_CHAIN_ID_0, 1, writerS);
        assertCorrectVerificationTime(emptyEntryS, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryS));
        assertEq(icDB.readEntry(address(moduleB), emptyEntryS), 0);
    }

    function test_readEntry_emptyA_emptyB() public {
        introduceEqualEmptyEntries();
        // writerF {0: 20} was verified as "empty" by module A and module B
        InterchainEntry memory emptyEntryF = getEmptyEntry(SRC_CHAIN_ID_0, 20, writerF);
        assertCorrectVerificationTime(emptyEntryF, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryF));
        assertCorrectVerificationTime(emptyEntryF, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryF));
        // writerS {1: 5} was verified as "empty" by module A and module B
        InterchainEntry memory emptyEntryS = getEmptyEntry(SRC_CHAIN_ID_1, 5, writerS);
        assertCorrectVerificationTime(emptyEntryS, address(moduleA), icDB.readEntry(address(moduleA), emptyEntryS));
        assertCorrectVerificationTime(emptyEntryS, address(moduleB), icDB.readEntry(address(moduleB), emptyEntryS));
    }

    // ═════════════════════════════════════ TESTS: READING ENTRIES (REVERTS) ══════════════════════════════════════════

    function test_readEntry_revert_sameChainId() public {
        InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, 0, writerF);
        expectSameChainId();
        icDB.readEntry(address(moduleA), entry);
    }
}
