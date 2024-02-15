// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, IInterchainDB, IInterchainDBEvents} from "../contracts/InterchainDB.sol";

import {InterchainModuleMock, IInterchainModule} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Unit tests for InterchainDB interactions on the destination chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBDestinationTest is Test, IInterchainDBEvents {
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
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_0, writerF, 0));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_0, writerF, 10));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, writerF, 10));
        // writerS: {1: 0}, {1: 10}
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, writerS, 0));
        verifyEntry(moduleA, getMockEntry(SRC_CHAIN_ID_1, writerS, 10));
        // Verify some entries with module B
        // writerF: {1: 0}, {1: 10}
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, writerF, 0));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, writerF, 10));
        // writerS: {0: 0}, {0: 10}, {1: 0}
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_0, writerS, 0));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_0, writerS, 10));
        verifyEntry(moduleB, getMockEntry(SRC_CHAIN_ID_1, writerS, 0));
    }

    function verifyEntry(InterchainModuleMock module, IInterchainDB.InterchainEntry memory entry) internal {
        skip(1 minutes);
        verifiedAt[address(module)][keccak256(abi.encode(entry))] = block.timestamp;
        module.mockVerifyEntry(address(icDB), entry);
    }

    function introduceConflicts() public {
        // Have module A verify a different entry for writerS {0:10} (already verified by module B)
        verifyEntry(moduleA, getFakeEntry(SRC_CHAIN_ID_0, writerS, 10));
        // Have module B verify a different entry for writerF {0:10} (already verified by module A)
        verifyEntry(moduleB, getFakeEntry(SRC_CHAIN_ID_0, writerF, 10));
    }

    // ══════════════════════════════════════════════ DATA GENERATION ══════════════════════════════════════════════════

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    function getMockEntry(
        uint256 srcChainId,
        address writer,
        uint256 nonce
    )
        internal
        pure
        returns (IInterchainDB.InterchainEntry memory entry)
    {
        return IInterchainDB.InterchainEntry({
            srcChainId: srcChainId,
            srcWriter: addressToBytes32(writer),
            writerNonce: nonce,
            dataHash: getMockDataHash(writer, nonce)
        });
    }

    function getFakeDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce, "Fake data"));
    }

    function getFakeEntry(
        uint256 srcChainId,
        address writer,
        uint256 nonce
    )
        internal
        pure
        returns (IInterchainDB.InterchainEntry memory entry)
    {
        return IInterchainDB.InterchainEntry({
            srcChainId: srcChainId,
            srcWriter: addressToBytes32(writer),
            writerNonce: nonce,
            dataHash: getFakeDataHash(writer, nonce)
        });
    }

    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // ═══════════════════════════════════════════════ TEST HELPERS ════════════════════════════════════════════════════

    function assertCorrectVerificationTime(
        IInterchainDB.InterchainEntry memory entry,
        address module,
        uint256 moduleVerifiedAt
    )
        internal
    {
        assertEq(verifiedAt[module][keccak256(abi.encode(entry))], moduleVerifiedAt);
    }

    function expectConflictingEntries(IInterchainDB.InterchainEntry memory existingEntry, bytes32 dataHash) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IInterchainDB.InterchainDB__ConflictingEntries.selector, existingEntry, dataHash)
        );
    }

    function expectSameChainId() internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainDB.InterchainDB__SameChainId.selector));
    }

    // ═════════════════════════════════════════ TESTS: VERIFYING ENTRIES ══════════════════════════════════════════════

    function test_verifyEntry_new_emitsEvent() public {
        skip(1 days);
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerF, 20);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(
            address(moduleA), entry.srcChainId, entry.srcWriter, entry.writerNonce, entry.dataHash
        );
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_new_savesVerificationTime() public {
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerF, 20);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    function test_verifyEntry_existing_diffModule_emitsEvent() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerS, 0);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(
            address(moduleA), entry.srcChainId, entry.srcWriter, entry.writerNonce, entry.dataHash
        );
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_existing_diffModule_doesNotUpdateExistingVerificationTime() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerS, 0);
        uint256 moduleBVerifiedAt = verifiedAt[address(moduleB)][keccak256(abi.encode(entry))];
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleB), entry), moduleBVerifiedAt);
    }

    function test_verifyEntry_existing_diffModule_savesVerificationTime() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerS, 0);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    function test_verifyEntry_existing_sameModule_doesNotEmitEvent() public {
        // writerF {0:0} was already verified by module A
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerF, 0);
        vm.recordLogs();
        verifyEntry(moduleA, entry);
        assertEq(vm.getRecordedLogs().length, 0);
    }

    function test_verifyEntry_existing_sameModule_doesNotUpdateExistingVerificationTime() public {
        // writerF {0:0} was already verified by module A
        IInterchainDB.InterchainEntry memory entry = getMockEntry(SRC_CHAIN_ID_0, writerF, 0);
        uint256 moduleAVerifiedAt = verifiedAt[address(moduleA)][keccak256(abi.encode(entry))];
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), moduleAVerifiedAt);
    }

    function test_verifyEntry_conflict_diffModule_emitsEvent() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, writerS, 0);
        vm.expectEmit(address(icDB));
        emit InterchainEntryVerified(
            address(moduleA), entry.srcChainId, entry.srcWriter, entry.writerNonce, entry.dataHash
        );
        verifyEntry(moduleA, entry);
    }

    function test_verifyEntry_conflict_diffModule_doesNotUpdateExistingVerificationTime() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, writerS, 0);
        IInterchainDB.InterchainEntry memory realEntry = getMockEntry(SRC_CHAIN_ID_0, writerS, 0);
        uint256 moduleBVerifiedAt = verifiedAt[address(moduleB)][keccak256(abi.encode(realEntry))];
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleB), realEntry), moduleBVerifiedAt);
    }

    function test_verifyEntry_conflict_diffModule_savesVerificationTime() public {
        // writerS {0:0} was already verified by module B
        IInterchainDB.InterchainEntry memory entry = getFakeEntry(SRC_CHAIN_ID_0, writerS, 0);
        verifyEntry(moduleA, entry);
        assertEq(icDB.readEntry(address(moduleA), entry), block.timestamp);
    }

    // ════════════════════════════════════ TESTS: VERIFYING ENTRIES (REVERTS) ═════════════════════════════════════════

    function test_verifyEntry_revert_conflict_sameModule() public {
        // writerF {0:0} was already verified by module A
        IInterchainDB.InterchainEntry memory existingEntry = getMockEntry(SRC_CHAIN_ID_0, writerF, 0);
        IInterchainDB.InterchainEntry memory conflictingEntry = getFakeEntry(SRC_CHAIN_ID_0, writerF, 0);
        expectConflictingEntries(existingEntry, conflictingEntry.dataHash);
        verifyEntry(moduleA, conflictingEntry);
    }

    function test_verifyEntry_revert_sameChainId() public {
        // Try to verify entry coming from the same chain
        IInterchainDB.InterchainEntry memory entry = getMockEntry(DST_CHAIN_ID, writerF, 0);
        expectSameChainId();
        verifyEntry(moduleA, entry);
    }
}
