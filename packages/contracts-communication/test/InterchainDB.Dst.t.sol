// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, InterchainBatch, IInterchainDB, InterchainDBEvents} from "../contracts/InterchainDB.sol";

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

    uint256 public constant BATCH_CONFLICT = type(uint256).max;

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

    function getMockBatchRoot(uint64 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, getMockDataHash(nonce)));
    }

    function getMockBatch(uint64 srcChainId, uint64 dbNonce) internal view returns (InterchainBatch memory batch) {
        return InterchainBatch(srcChainId, dbNonce, getMockBatchRoot(dbNonce));
    }

    function getMockDataHash(uint64 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce));
    }

    function getFakeBatchRoot(uint64 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, getFakeDataHash(nonce)));
    }

    function getFakeBatch(uint64 srcChainId, uint64 dbNonce) internal view returns (InterchainBatch memory batch) {
        return InterchainBatch(srcChainId, dbNonce, getFakeBatchRoot(dbNonce));
    }

    function getFakeDataHash(uint64 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(nonce, "Fake data"));
    }

    function getEmptyBatchRoot() internal view returns (bytes32) {
        return keccak256(abi.encode(writerT, 0));
    }

    function getEmptyBatch(uint64 srcChainId, uint64 dbNonce) internal pure returns (InterchainBatch memory batch) {
        return InterchainBatch({srcChainId: srcChainId, dbNonce: dbNonce, batchRoot: 0});
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectInitialVerificationTime(
        InterchainModuleMock module,
        InterchainBatch memory batch
    )
        internal
        view
    {
        bytes memory versionedBatch = getVersionedBatch(batch);
        uint256 savedVerificationTime = verifiedAt[address(module)][keccak256(abi.encode(versionedBatch))];
        // We never save 0 as a verification time during initial setup
        assertGt(savedVerificationTime, 0);
        checkVerification(module, batch, savedVerificationTime);
    }

    function checkVerification(
        InterchainModuleMock module,
        InterchainBatch memory batch,
        uint256 expectedVerificationTime
    )
        internal
        view
    {
        uint256 timestamp = icDB.checkBatchVerification(address(module), batch);
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

    function expectSameChainId(uint64 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__SameChainId.selector, chainId));
    }

    // ═════════════════════════════════════════ TESTS: VERIFYING BATCHES ══════════════════════════════════════════════

    function test_verifyBatch_newKey_nonEmptyRoot() public {
        skip(1 days);
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, batch, block.timestamp);
    }

    function test_verifyBatch_newKey_emptyRoot() public {
        skip(1 days);
        InterchainBatch memory batch = getEmptyBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, batch, block.timestamp);
    }

    function test_verifyBatch_sameKey_diffModule_prevEmptyRoot_emptyRoot() public {
        introduceEmptyBatches();
        skip(1 days);
        // {0: 30} was verified as "empty" by module B, but not by module A
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 30);
        bytes memory versionedEmptyBatch = getVersionedBatch(emptyBatch);
        expectEventBatchVerified(moduleA, emptyBatch);
        verifyBatch(moduleA, versionedEmptyBatch);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, emptyBatch, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, emptyBatch);
    }

    function test_verifyBatch_sameKey_diffModule_prevEmptyRoot_nonEmptyRoot() public {
        introduceEmptyBatches();
        skip(1 days);
        // {0: 30} was verified as "empty" by module B, but not by module A
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 30);
        InterchainBatch memory mockBatch = getMockBatch(SRC_CHAIN_ID_0, 30);
        bytes memory versionedBatch = getVersionedBatch(mockBatch);
        expectEventBatchVerified(moduleA, mockBatch);
        verifyBatch(moduleA, versionedBatch);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, mockBatch, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, emptyBatch);
    }

    function test_verifyBatch_sameKey_diffModule_prevNonEmptyRoot_emptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedEmptyBatch = getVersionedBatch(emptyBatch);
        expectEventBatchVerified(moduleA, emptyBatch);
        verifyBatch(moduleA, versionedEmptyBatch);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, emptyBatch, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_verifyBatch_sameKey_diffModule_prevNonEmptyRoot_diffNonEmptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        InterchainBatch memory conflictingBatch = getFakeBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedConflictingBatch = getVersionedBatch(conflictingBatch);
        expectEventBatchVerified(moduleA, conflictingBatch);
        verifyBatch(moduleA, versionedConflictingBatch);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, conflictingBatch, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_verifyBatch_sameKey_diffModule_prevNonEmptyRoot_sameNonEmptyRoot() public {
        skip(1 days);
        // {1: 0} was verified by module B, but not by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        bytes memory versionedBatch = getVersionedBatch(batch);
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
        // Should save the verification time for A and not overwrite the existing verification time for B
        checkVerification(moduleA, batch, block.timestamp);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_verifyBatch_sameKey_sameModule_prevEmptyRoot_emptyRoot() public {
        introduceEmptyBatches();
        skip(1 days);
        // {0: 20} was verified as "empty" by module A
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedEmptyBatch = getVersionedBatch(emptyBatch);
        // Should emit no event and not update the verification time
        uint256 moduleAVerificationTime = verifiedAt[address(moduleA)][keccak256(abi.encode(versionedEmptyBatch))];
        vm.recordLogs();
        verifyBatch(moduleA, versionedEmptyBatch);
        assertEq(vm.getRecordedLogs().length, 0);
        checkVerification(moduleA, emptyBatch, moduleAVerificationTime);
    }

    function test_verifyBatch_sameKey_sameModule_prevEmptyRoot_nonEmptyRoot() public {
        introduceEmptyBatches();
        skip(1 days);
        // {0: 20} was verified as "empty" by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 20);
        bytes memory versionedBatch = getVersionedBatch(batch);
        // Overwriting an empty batch with a non-empty one is allowed
        expectEventBatchVerified(moduleA, batch);
        verifyBatch(moduleA, versionedBatch);
        checkVerification(moduleA, batch, block.timestamp);
    }

    function test_verifyBatch_sameKey_sameModule_prevNonEmptyRoot_sameNonEmptyRoot() public {
        skip(1 days);
        // {0: 10} was verified by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 10);
        bytes memory versionedBatch = getVersionedBatch(batch);
        // Should emit no event and not update the verification time
        uint256 moduleAVerificationTime = verifiedAt[address(moduleA)][keccak256(abi.encode(versionedBatch))];
        vm.recordLogs();
        verifyBatch(moduleA, versionedBatch);
        assertEq(vm.getRecordedLogs().length, 0);
        checkVerification(moduleA, batch, moduleAVerificationTime);
    }

    // ════════════════════════════════════ TESTS: VERIFYING BATCHES (REVERTS) ═════════════════════════════════════════

    function test_verifyBatch_sameKey_sameModule_prevNonEmptyRoot_emptyRoot_revert() public {
        // {0: 0} was verified by module A
        InterchainBatch memory existingBatch = getMockBatch(SRC_CHAIN_ID_0, 0);
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 0);
        bytes memory versionedEmptyBatch = getVersionedBatch(emptyBatch);
        expectRevertConflictingBatches(moduleA, existingBatch, emptyBatch);
        verifyBatch(moduleA, versionedEmptyBatch);
    }

    function test_verifyBatch_sameKey_sameModule_prevNonEmptyRoot_diffNonEmptyRoot_revert() public {
        // {0: 0} was verified by module A
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

    function test_checkVerification_existingA_existingB() public view {
        // {1: 10} was verified by module A and module B
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 10);
        assertCorrectInitialVerificationTime(moduleA, batch);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_checkVerification_existingA_unknownB() public view {
        // {0: 0} was verified by module A, but not by module B
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, batch);
        checkVerification(moduleB, batch, 0);
    }

    function test_checkVerification_existingA_differentB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" batch was verified by module B
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, batch);
        checkVerification(moduleB, batch, BATCH_CONFLICT);
    }

    function test_checkVerification_existingA_emptyB() public {
        introduceEmptyBatches();
        // {0: 10} was verified by module A, but an "empty" batch was verified by module B
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 10);
        assertCorrectInitialVerificationTime(moduleA, batch);
        checkVerification(moduleB, batch, BATCH_CONFLICT);
    }

    function test_checkVerification_unknownA_existingB() public view {
        // {1: 0} was verified by module B, but not by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, batch, 0);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_checkVerification_unknownA_unknownB() public view {
        // {0: 20} was not verified by any module
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 20);
        checkVerification(moduleA, batch, 0);
        checkVerification(moduleB, batch, 0);
    }

    function test_checkVerification_unknownA_differentB() public view {
        // {1: 0} was verified by module B, but not by module A
        // Check the fake batch that neither module verified
        InterchainBatch memory fakeBatch = getFakeBatch(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, fakeBatch, 0);
        checkVerification(moduleB, fakeBatch, BATCH_CONFLICT);
    }

    function test_checkVerification_unknownA_emptyB() public {
        introduceEmptyBatches();
        // {0: 30} was not verified by module A, but an "empty" batch was verified by module B
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 30);
        checkVerification(moduleA, emptyBatch, 0);
        assertCorrectInitialVerificationTime(moduleB, emptyBatch);
    }

    function test_checkVerification_differentA_existingB() public {
        introduceConflicts();
        // {0: 10} was verified by module A, but a "fake" batch was verified by module B
        // Check the fake batch that A never verified
        InterchainBatch memory fakeBatch = getFakeBatch(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeBatch, BATCH_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, fakeBatch);
    }

    function test_checkVerification_differentA_unknownB() public view {
        // {0: 10} was verified by module A, but not by module B
        // Check the fake batch that neither module verified
        InterchainBatch memory fakeBatch = getFakeBatch(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, fakeBatch, BATCH_CONFLICT);
        checkVerification(moduleB, fakeBatch, 0);
    }

    function test_checkVerification_differentA_differentB() public view {
        // {1: 10} was verified by module A and module B
        // Check the fake batch that neither module verified
        InterchainBatch memory fakeBatch = getFakeBatch(SRC_CHAIN_ID_1, 10);
        checkVerification(moduleA, fakeBatch, BATCH_CONFLICT);
        checkVerification(moduleB, fakeBatch, BATCH_CONFLICT);
    }

    function test_checkVerification_differentA_emptyB() public {
        introduceEmptyBatches();
        // {0: 10} was verified by module A, but an "empty" batch was verified by module B
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 10);
        checkVerification(moduleA, emptyBatch, BATCH_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, emptyBatch);
    }

    function test_checkVerification_emptyA_existingB() public {
        introduceEmptyBatches();
        // {1: 0} was verified by module B, but an "empty" batch was verified by module A
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_1, 0);
        checkVerification(moduleA, batch, BATCH_CONFLICT);
        assertCorrectInitialVerificationTime(moduleB, batch);
    }

    function test_checkVerification_emptyA_unknownB() public {
        introduceEmptyBatches();
        // {0: 20} was verified as "empty" by module A, but not by module B
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyBatch);
        checkVerification(moduleB, emptyBatch, 0);
    }

    function test_checkVerification_emptyA_differentB() public {
        introduceEmptyBatches();
        // {1: 0} was verified by module B, but an "empty" batch was verified by module A
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_1, 0);
        assertCorrectInitialVerificationTime(moduleA, emptyBatch);
        checkVerification(moduleB, emptyBatch, BATCH_CONFLICT);
    }

    function test_checkVerification_emptyA_emptyB() public {
        introduceEqualEmptyBatches();
        // {0: 20} was verified as "empty" by module A and module B
        InterchainBatch memory emptyBatch = getEmptyBatch(SRC_CHAIN_ID_0, 20);
        assertCorrectInitialVerificationTime(moduleA, emptyBatch);
        assertCorrectInitialVerificationTime(moduleB, emptyBatch);
    }

    function test_checkVerification_modifySrcChainId() public view {
        // Valid batch
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, batch);
        batch.srcChainId ^= 1;
        checkVerification(moduleA, batch, 0);
    }

    function test_checkVerification_modifyDbNonce() public view {
        // Valid batch
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, batch);
        batch.dbNonce ^= 1;
        checkVerification(moduleA, batch, 0);
    }

    function test_checkVerification_modifyBatchRoot() public view {
        // Valid entry
        InterchainBatch memory batch = getMockBatch(SRC_CHAIN_ID_0, 0);
        assertCorrectInitialVerificationTime(moduleA, batch);
        batch.batchRoot ^= bytes32(uint256(1));
        checkVerification(moduleA, batch, BATCH_CONFLICT);
    }

    // ═════════════════════════════════════ TESTS: READING BATCHES (REVERTS) ══════════════════════════════════════════

    function test_checkVerification_revert_sameChainId() public {
        InterchainBatch memory batch = getMockBatch(DST_CHAIN_ID, 0);
        expectSameChainId(DST_CHAIN_ID);
        icDB.checkBatchVerification(address(moduleA), batch);
    }
}
