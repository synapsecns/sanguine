// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {
    InterchainDB,
    InterchainBatch,
    InterchainBatchLib,
    InterchainEntry,
    IInterchainDB,
    InterchainDBEvents
} from "../contracts/InterchainDB.sol";

import {InterchainBatchLibHarness} from "./harnesses/InterchainBatchLibHarness.sol";
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

    InterchainBatchLibHarness public batchLibHarness;
    VersionedPayloadLibHarness public payloadLibHarness;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

    address public writerT = makeAddr("Test Writer");

    mapping(address module => mapping(bytes32 batchHash => uint256 timestamp)) public verifiedAt;

    function setUp() public {
        vm.chainId(DST_CHAIN_ID);
        icDB = new InterchainDB();
        moduleA = new InterchainModuleMock();
        moduleB = new InterchainModuleMock();
        batchLibHarness = new InterchainBatchLibHarness();
        payloadLibHarness = new VersionedPayloadLibHarness();
        // Format is {chainId: nonce}
        // Verify some batches with module A
        // A: {0: 0}, {0: 10}, {1: 10}
        verifyBatch(moduleA, getVersionedBatch(getMockBatch(SRC_CHAIN_ID_0, 0)));
        verifyBatch(moduleA, getVersionedBatch(getMockBatch(SRC_CHAIN_ID_0, 10)));
        verifyBatch(moduleA, getVersionedBatch(getMockBatch(SRC_CHAIN_ID_1, 10)));
        // Verify some batches with module B
        // B: {1: 0}, {1: 10}
        verifyBatch(moduleB, getVersionedBatch(getMockBatch(SRC_CHAIN_ID_1, 0)));
        verifyBatch(moduleB, getVersionedBatch(getMockBatch(SRC_CHAIN_ID_1, 10)));
    }

    function getVersionedBatch(InterchainBatch memory batch) internal view returns (bytes memory) {
        return payloadLibHarness.encodeVersionedPayload(DB_VERSION, batchLibHarness.encodeBatch(batch));
    }

    function verifyBatch(InterchainModuleMock module, bytes memory versionedBatch) internal {
        skip(1 minutes);
        verifiedAt[address(module)][keccak256(abi.encode(versionedBatch))] = block.timestamp;
        module.mockVerifyRemoteBatch(address(icDB), versionedBatch);
    }

    function introduceConflicts() public {
        // Have module A verify a different batch {1:0} (already verified by module B)
        verifyBatch(moduleA, getVersionedBatch(getFakeBatch(SRC_CHAIN_ID_1, 0)));
        // Have module B verify a different batch {0:10} (already verified by module A)
        verifyBatch(moduleB, getVersionedBatch(getFakeBatch(SRC_CHAIN_ID_0, 10)));
    }

    function introduceEmptyBatches() public {
        // Have module A verify an empty batch for batches that module B has not verified
        // {0: 20}
        verifyBatch(moduleA, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_0, 20)));
        // Have module A verify an empty batch for batches that module B has verified
        // {1: 0}
        verifyBatch(moduleA, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_1, 0)));
        // Have module B verify an empty batch for batches that module A has not verified
        // {0: 30}
        verifyBatch(moduleB, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_0, 30)));
        // Have module B verify an empty batch for batches that module A has verified
        // {0: 10}
        verifyBatch(moduleB, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_0, 10)));
    }

    function introduceEqualEmptyBatches() public {
        // {0: 20}
        verifyBatch(moduleA, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_0, 20)));
        verifyBatch(moduleB, getVersionedBatch(getEmptyBatch(SRC_CHAIN_ID_0, 20)));
    }

    // ══════════════════════════════════════════════ DATA GENERATION ══════════════════════════════════════════════════

    function getMockBatchRoot(uint256 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, getMockDataHash(nonce)));
    }

    function getMockBatch(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainBatch memory batch) {
        return InterchainBatch(srcChainId, dbNonce, getMockBatchRoot(dbNonce));
    }

    function getMockDataHash(uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce));
    }

    function getMockEntry(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            entryIndex: 0,
            srcWriter: addressToBytes32(writerT),
            dataHash: getMockDataHash(dbNonce)
        });
    }

    function getFakeBatchRoot(uint256 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, getFakeDataHash(nonce)));
    }

    function getFakeBatch(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainBatch memory batch) {
        return InterchainBatch(srcChainId, dbNonce, getFakeBatchRoot(dbNonce));
    }

    function getFakeDataHash(uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce, "Fake data"));
    }

    function getFakeEntry(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            entryIndex: 0,
            srcWriter: addressToBytes32(writerT),
            dataHash: getFakeDataHash(dbNonce)
        });
    }

    function getEmptyBatchRoot() internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, 0));
    }

    function getEmptyBatch(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainBatch memory batch) {
        return InterchainBatch({srcChainId: srcChainId, dbNonce: dbNonce, batchRoot: getEmptyBatchRoot()});
    }

    function getEmptyEntry(uint64 srcChainId, uint256 dbNonce) internal view returns (InterchainEntry memory entry) {
        return InterchainEntry({
            srcChainId: srcChainId,
            dbNonce: dbNonce,
            entryIndex: 0,
            srcWriter: addressToBytes32(writerT),
            dataHash: 0
        });
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectInitialVerificationTime(InterchainModuleMock module, InterchainEntry memory entry) internal {
        InterchainBatch memory batch =
            InterchainBatch(entry.srcChainId, entry.dbNonce, keccak256(abi.encode(entry.srcWriter, entry.dataHash)));
        bytes memory versionedBatch = getVersionedBatch(batch);
        uint256 savedVerificationTime = verifiedAt[address(module)][keccak256(abi.encode(versionedBatch))];
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
    {
        checkVerification(module, entry, new bytes32[](0), expectedVerificationTime);
    }

    function checkVerification(
        InterchainModuleMock module,
        InterchainEntry memory entry,
        bytes32[] memory proof,
        uint256 expectedVerificationTime
    )
        internal
    {
        uint256 timestamp = icDB.checkVerification(address(module), entry, proof);
        assertEq(timestamp, expectedVerificationTime);
    }

    function expectEventBatchVerified(InterchainModuleMock module, InterchainBatch memory batch) internal {
        vm.expectEmit(address(icDB));
        emit InterchainBatchVerified(address(module), batch.srcChainId, batch.dbNonce, batch.batchRoot);
    }

    function expectRevertConflictingBatches(
        InterchainModuleMock module,
        InterchainBatch memory existingBatch,
        InterchainBatch memory newBatch
    )
        internal
    {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainDB.InterchainDB__ConflictingBatches.selector,
                address(module),
                existingBatch.batchRoot,
                newBatch
            )
        );
    }

    function expectRevertInvalidBatchVersion(uint16 version) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__InvalidBatchVersion.selector, version));
    }

    function expectSameChainId(uint256 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__SameChainId.selector, chainId));
    }

    // ═════════════════════════════════════════ TESTS: VERIFYING BATCHES ══════════════════════════════════════════════

    function test_verifyBatch_new_emitsEvent() public {
        skip(1 days);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
    }

    function test_verifyBatch_new_savesVerificationTime() public {
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedBatch = getVersionedBatch(batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, entry, block.timestamp);
    }

    function test_verifyBatch_existing_diffModule_emitsEvent() public {
        // {1: 0} was already verified by module B
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
    }

    function test_verifyBatch_existing_diffModule_doesNotUpdateExistingVerificationTime() public {
        // {1: 0} was already verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        verifyBatch(moduleA, versionedBatch);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_verifyBatch_existing_diffModule_savesVerificationTime() public {
        // {1: 0} was already verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, entry, block.timestamp);
    }

    function test_verifyBatch_existing_sameModule_doesNotEmitEvent() public {
        // {0:0} was already verified by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        vm.recordLogs();
        verifyBatch(moduleA, versionedBatch);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_verifyBatch_existing_sameModule_doesNotUpdateExistingVerificationTime() public {
        // {0:0} was already verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        uint256 moduleAVerificationTime = verifiedAt[address(moduleA)][keccak256(abi.encode(versionedBatch))];
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, entry, moduleAVerificationTime);
    }

    function test_verifyBatch_conflict_diffModule_emitsEvent() public {
        // {1: 0} was already verified by module B
        InterchainBatch memory batch = getFakeBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
    }

    function test_verifyBatch_conflict_diffModule_doesNotUpdateExistingVerificationTime() public {
        // {1: 0} was already verified by module B
        InterchainBatch memory batch = getFakeBatch(SRC_CHAIN_ID_1, 0);
        InterchainEntry memory realEntry = getMockEntry(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        verifyBatch(moduleA, versionedBatch);
        assertCorrectInitialVerificationTime(moduleB, realEntry);
    }

    function test_verifyBatch_conflict_diffModule_savesVerificationTime() public {
        // {1: 0} was already verified by module B
        InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_1, 0);
        InterchainBatch memory batch = getFakeBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, entry, block.timestamp);
    }

    // ════════════════════════════════════ TESTS: VERIFYING BATCHES (REVERTS) ═════════════════════════════════════════

    function test_verifyBatch_revert_conflict_sameModule() public {
        // {0:0} was already verified by module A
        InterchainBatch memory existingBatch = getMockBatch(SRC_CHAIN_ID_0, 0);
        InterchainBatch memory conflictingBatch = getFakeBatch(SRC_CHAIN_ID_0, 0);
        bytes memory versionedConflictingBatch = getVersionedBatch(conflictingBatch);
        expectRevertConflictingBatches(moduleA, existingBatch, conflictingBatch);
        verifyBatch(moduleA, versionedConflictingBatch);
    }

    function test_verifyBatch_revert_sameChainId() public {
        // Try to verify batch coming from the same chain
        InterchainBatch memory batch = getMockBatch(DST_CHAIN_ID, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectSameChainId(DST_CHAIN_ID);
        verifyBatch(moduleA, versionedBatch);
    }

    function test_verifyBatch_revert_wrongVersion(uint16 version) public {
        vm.assume(version != DB_VERSION);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        bytes memory versionedBatch =
            payloadLibHarness.encodeVersionedPayload(version, batchLibHarness.encodeBatch(batch));
        expectRevertInvalidBatchVersion(version);
        moduleA.mockVerifyRemoteBatch(address(icDB), versionedBatch);
    }

    // ══════════════════════════════════════════ TESTS: READING BATCHES ═══════════════════════════════════════════════

    function test_checkVerification_existingA_existingB() public {
        // {1: 10} was verified by module A and module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_existingA_unknownB() public {
        // {0: 0} was verified by module A, but not by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_existingA_differentB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" batch was verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_existingA_emptyB() public {
        introduceEmptyBatches();
        // {0: 10} was verified by module A, but an "empty" batch was verified by module B
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, entry);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_unknownA_existingB() public {
        // {1: 0} was verified by module B, but not by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, entry, 0);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_unknownA_unknownB() public {
        // {0: 20} was not verified by any module
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 20);
        checkVerification(moduleA, entry, 0);
        checkVerification(moduleB, entry, 0);
    }

    function test_checkVerification_unknownA_differentB() public {
        // {1: 0} was verified by module B, but not by module A
        // Check the fake batch that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, fakeEntry, 0);
        checkVerification(moduleB, fakeEntry, 0);
    }

    function test_checkVerification_unknownA_emptyB() public {
        introduceEmptyBatches();
        // {0: 30} was not verified by module A, but an "empty" batch was verified by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 30);
        checkVerification(moduleA, emptyEntry, 0);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_differentA_existingB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" batch was verified by module B
        // Check the fake batch that A never verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeEntry, 0);
        assertCorrectInitialVerificationTime(moduleB, fakeEntry);
    }

    function test_checkVerification_differentA_unknownB() public {
        // {0: 10} was verified by module A, but not by module B
        // Check the fake batch that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeEntry, 0);
        checkVerification(moduleB, fakeEntry, 0);
    }

    function test_checkVerification_differentA_differentB() public {
        // {1: 10} was verified by module A and module B
        // Check the fake batch that neither module verified
        InterchainEntry memory fakeEntry = getFakeEntry(SRC_CHAIN_ID_1, 10);
        checkVerification(moduleA, fakeEntry, 0);
        checkVerification(moduleB, fakeEntry, 0);
    }

    function test_checkVerification_differentA_emptyB() public {
        introduceEmptyBatches();
        // {0: 10} was verified by module A, but an "empty" batch was verified by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, emptyEntry, 0);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_emptyA_existingB() public {
        introduceEmptyBatches();
        // {1: 0} was verified by module B, but an "empty" batch was verified by module A
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, entry, 0);
        assertCorrectInitialVerificationTime(moduleB, entry);
    }

    function test_checkVerification_emptyA_unknownB() public {
        introduceEmptyBatches();
        // {0: 20} was verified as "empty" by module A, but not by module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        checkVerification(moduleB, emptyEntry, 0);
    }

    function test_checkVerification_emptyA_differentB() public {
        introduceEmptyBatches();
        // {1: 0} was verified by module B, but an "empty" batch was verified by module A
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_1, 0);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        checkVerification(moduleB, emptyEntry, 0);
    }

    function test_checkVerification_emptyA_emptyB() public {
        introduceEqualEmptyBatches();
        // {0: 20} was verified as "empty" by module A and module B
        InterchainEntry memory emptyEntry = getEmptyEntry(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyEntry);
        assertCorrectInitialVerificationTime(moduleB, emptyEntry);
    }

    function test_checkVerification_nonEmptyProof() public {
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = hex"deadbeaf";
        // Should return 0 for batches that were verified or not
        checkVerification(moduleA, getMockEntry(SRC_CHAIN_ID_0, 0), proof, 0);
        checkVerification(moduleA, getFakeEntry(SRC_CHAIN_ID_0, 0), proof, 0);
        proof[0] = 0;
        checkVerification(moduleB, getMockEntry(SRC_CHAIN_ID_1, 0), proof, 0);
        checkVerification(moduleB, getFakeEntry(SRC_CHAIN_ID_1, 0), proof, 0);
    }

    function test_checkVerification_modifySrcChainId() public {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.srcChainId ^= 1;
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifyDbNonce() public {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.dbNonce ^= 1;
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifyEntryIndex() public {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.entryIndex ^= 1;
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifySrcWriter() public {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.srcWriter ^= bytes32(uint256(1));
        checkVerification(moduleA, entry, 0);
    }

    function test_checkVerification_modifyDataHash() public {
        // Valid entry
        InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, entry);
        entry.dataHash ^= bytes32(uint256(1));
        checkVerification(moduleA, entry, 0);
    }

    // ═════════════════════════════════════ TESTS: READING BATCHES (REVERTS) ══════════════════════════════════════════

    function test_checkVerification_revert_sameChainId() public {
        InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, 0);
        expectSameChainId(DST_CHAIN_ID);
        icDB.checkVerification(address(moduleA), entry, new bytes32[](0));
    }
}
