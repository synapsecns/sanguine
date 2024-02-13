// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainDB, IInterchainDB, IInterchainDBEvents} from "../contracts/InterchainDB.sol";

import {InterchainModuleMock} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

/// @notice Unit tests for InterchainDB interactions on the source chain
/// Note: we inherit from interface with the events to avoid their copy-pasting.
contract InterchainDBSourceTest is Test, IInterchainDBEvents {
    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    uint256 public constant INITIAL_WRITER_A_NONCE = 1;
    uint256 public constant INITIAL_WRITER_B_NONCE = 2;

    InterchainDB public icDB;
    InterchainModuleMock public moduleA;
    InterchainModuleMock public moduleB;

    address[] public oneModule;
    address[] public twoModules;

    address public writerA = makeAddr("Writer A");
    address public writerB = makeAddr("Writer B");

    function setUp() public {
        vm.chainId(SRC_CHAIN_ID);
        icDB = new InterchainDB();
        moduleA = new InterchainModuleMock();
        moduleB = new InterchainModuleMock();
        oneModule.push(address(moduleA));
        twoModules.push(address(moduleA));
        twoModules.push(address(moduleB));
        setupWriterNonce(writerA, INITIAL_WRITER_A_NONCE);
        setupWriterNonce(writerB, INITIAL_WRITER_B_NONCE);
    }

    function setupWriterNonce(address writer, uint256 nonce) internal {
        for (uint256 i = 0; i < nonce; i++) {
            writeEntry(writer, 0, getMockDataHash(writer, i), oneModule);
        }
    }

    function getMockDataHash(address writer, uint256 nonce) internal pure returns (bytes32) {
        return keccak256(abi.encode(writer, nonce));
    }

    /// @dev Mocks a return value of module.getModuleFee(DST_CHAIN_ID)
    function mockModuleFee(InterchainModuleMock module, uint256 feeValue) internal {
        bytes memory callData = abi.encodeCall(module.getModuleFee, (DST_CHAIN_ID));
        bytes memory returnData = abi.encode(feeValue);
        vm.mockCall(address(module), callData, returnData);
    }

    function writeEntry(address writer, uint256 msgValue, bytes32 dataHash, address[] memory modules) internal {
        deal(writer, msgValue);
        vm.prank(writer);
        icDB.writeEntry{value: msgValue}(DST_CHAIN_ID, dataHash, modules);
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
        assertEq(icDB.getWriterNonce(writerA), INITIAL_WRITER_A_NONCE);
        assertEq(icDB.getWriterNonce(writerB), INITIAL_WRITER_B_NONCE);
    }

    function test_setup_getEntry() public {
        for (uint256 i = 0; i < INITIAL_WRITER_A_NONCE; i++) {
            assertEq(
                icDB.getEntry(writerA, i),
                IInterchainDB.InterchainEntry({
                    srcChainId: SRC_CHAIN_ID,
                    srcWriter: writerA,
                    writerNonce: i,
                    dataHash: getMockDataHash(writerA, i)
                })
            );
        }
        for (uint256 i = 0; i < INITIAL_WRITER_B_NONCE; i++) {
            assertEq(
                icDB.getEntry(writerB, i),
                IInterchainDB.InterchainEntry({
                    srcChainId: SRC_CHAIN_ID,
                    srcWriter: writerB,
                    writerNonce: i,
                    dataHash: getMockDataHash(writerB, i)
                })
            );
        }
    }

    // ═══════════════════════════════════════ TESTS: USING A SINGLE MODULE ════════════════════════════════════════════

    function test_getInterchainFee_oneModule() public {
        address[] memory modules = new address[](1);
        modules[0] = address(moduleA);
        mockModuleFee(moduleA, 100);
        assertEq(icDB.getInterchainFee(DST_CHAIN_ID, modules), 100);
    }

    function test_writeEntry_writerA_oneModule_emitsEvent() public {
        bytes32 dataHash = getMockDataHash(writerA, INITIAL_WRITER_A_NONCE);
        mockModuleFee(moduleA, 100);
        vm.expectEmit(address(icDB));
        emit InterchainEntryWritten(SRC_CHAIN_ID, writerA, INITIAL_WRITER_A_NONCE, dataHash);
        writeEntry({writer: writerA, msgValue: 100, dataHash: dataHash, modules: oneModule});
    }

    function test_writeEntry_writerA_oneModule_writesEntry() public {}

    function test_writeEntry_writerA_oneModule_increasesWriterNonce() public {}

    function test_writeEntry_writerA_oneModule_callsModuleForVerification() public {}

    // ═════════════════════════════════════════ TESTS: USING TWO MODULES ══════════════════════════════════════════════
}
